package main

import (
	"context"
	"github.com/failuretoload/datamonster/auth"
	"github.com/failuretoload/datamonster/settlement"
	postgres "github.com/failuretoload/datamonster/store/postgres"
	"github.com/failuretoload/datamonster/survivor"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

var (
	ready                = false
	connPool             *pgxpool.Pool
	app                  Server
	appContext           context.Context
	settlementController *settlement.Controller
	survivorController   *survivor.Controller
	authController       *auth.Controller
)

func main() {
	appContext = context.Background()

	connPool = postgres.InitConnPool(appContext)
	defer connPool.Close()

	settlementController = settlement.NewController(connPool)
	survivorController = survivor.NewController(connPool)
	authController = auth.NewController()

	app = NewServer(settlementController, survivorController, authController)
	app.Run()
}

type Server struct {
	Mux *chi.Mux
}

func NewServer(settlements *settlement.Controller,
	survivors *survivor.Controller,
	googleAuth *auth.Controller) Server {
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

	router.Mount("/auth", authRoutes(googleAuth))
	router.Mount("/api", protectedRoutes(settlements, survivors))
	return Server{
		Mux: router,
	}

}

func authRoutes(googleAuth *auth.Controller) http.Handler {
	r := chi.NewRouter()
	r.Use(CorsHandler())
	r.Post("/callback", googleAuth.HandleGoogleAuth)
	r.Get("/validate", googleAuth.ValidateToken)
	return r
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
	r.Use(auth.GoogleAuthHandler)
	settlement.RegisterRoutes(r)
	survivor.RegisterRoutes(r)
	return r
}

func (s Server) Handle(route string, handler http.Handler) {
	s.Mux.Handle(route, handler)
}

func (s Server) Run() {
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
