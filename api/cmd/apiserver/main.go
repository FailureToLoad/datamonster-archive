package main

import (
	"context"
	"fmt"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/failuretoload/datamonster/request"
	"github.com/failuretoload/datamonster/response"
	"github.com/failuretoload/datamonster/settlement"
	postgres "github.com/failuretoload/datamonster/store/postgres"
	"github.com/failuretoload/datamonster/survivor"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"

	"github.com/failuretoload/datamonster/config"
)

var (
	ready                = false
	connPool             *pgxpool.Pool
	app                  Server
	appContext           context.Context
	settlementController *settlement.Controller
	survivorController   *survivor.Controller
)

func main() {
	appContext = context.Background()

	connPool = postgres.InitConnPool(appContext)
	defer connPool.Close()

	settlementController = settlement.NewController(connPool)
	survivorController = survivor.NewController(connPool)

	app = NewServer(settlementController, survivorController)
	app.Run()
}

type Server struct {
	Mux *chi.Mux
}

func NewServer(settlements *settlement.Controller, survivors *survivor.Controller) Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(10 * time.Second))

	router.Get("/heartbeat", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/startup", func(w http.ResponseWriter, _ *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("api is not ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	router.Get("/ready", func(w http.ResponseWriter, _ *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("api is not ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	router.Mount("/api", protectedRoutes(settlements, survivors))
	return Server{
		Mux: router,
	}
}

func protectedRoutes(settlement *settlement.Controller, survivor *survivor.Controller) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(CorsHandler())
	r.Use(SecureOptions())
	r.Use(CacheControl)
	r.Use(clerkhttp.WithHeaderAuthorization())
	r.Use(UserIDExtractor)
	settlement.RegisterRoutes(r)
	survivor.RegisterRoutes(r)
	return r
}

func (s Server) Handle(route string, handler http.Handler) {
	s.Mux.Handle(route, handler)
}

func (s Server) Run() {
	clerk.SetKey(config.Key())
	ready = true
	log.Default().Println("Starting server on 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", s.Mux)
	if err != nil {
		log.Default().Fatal(err)
	}
}

func SecureOptions() func(http.Handler) http.Handler {
	options := secure.Options{
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ForceSTSHeader:        true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		CustomBrowserXssValue: "0",
		ContentSecurityPolicy: "default-src 'self', frame-ancestors 'none'",
	}
	return secure.New(options).Handler
}

func CacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		headers := rw.Header()
		headers.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		headers.Set("Pragma", "no-cache")
		headers.Set("Expires", "0")
		next.ServeHTTP(rw, req)
	})
}

func CorsHandler() func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           3599, // Maximum value not ignored by any of major browsers
	})
	return c.Handler
}

func UserIDExtractor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		claims, ok := clerk.SessionClaimsFromContext(req.Context())
		if !ok {
			response.Unauthorized(ctx, rw, fmt.Errorf("missing claims"))
			return
		}
		userID := claims.RegisteredClaims.Subject
		if userID == "" {
			response.Unauthorized(ctx, rw, fmt.Errorf("user id not found"))
			return
		}
		ctx = context.WithValue(req.Context(), request.UserIDKey, userID)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
