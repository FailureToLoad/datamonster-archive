package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/failuretoload/datamonster/config"
	"github.com/failuretoload/datamonster/ent"
	"github.com/failuretoload/datamonster/graph"
	"github.com/go-chi/cors"

	"github.com/unrolled/secure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Mux *chi.Mux
}

func NewServer(client *ent.Client, clientURI string) Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(10 * time.Second))
	router.Use(CorsHandler(clientURI))
	mode := config.Mode()
	if mode != "schema" {
		router.Use(SecureOptions())
		router.Use(CacheControl)
		router.Use(clerkhttp.WithHeaderAuthorization())
		router.Use(UserIdExtractor)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	router.Handle("/graphql", srv)
	if mode == "schema" {
		router.Handle("/",
			playground.Handler("Datamonster.gql", "/graphql"),
		)
	}
	return Server{
		Mux: router,
	}
}

func (s Server) Handle(route string, handler http.Handler) {
	s.Mux.Handle(route, handler)
}

func (s Server) Run(authKey string) {
	clerk.SetKey(authKey)
	log.Default().Println("Starting server on port 8080")
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

func CorsHandler(clientURI string) func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{clientURI},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           3599, // Maximum value not ignored by any of major browsers
	})
	return c.Handler
}

func UserIdExtractor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		claims, ok := clerk.SessionClaimsFromContext(req.Context())
		if !ok {
			Unauthorized(rw, nil)
			return
		}
		userID := claims.RegisteredClaims.Subject
		if userID == "" {
			reason := "user id not found"
			Unauthorized(rw, &reason)
			return
		}
		ctx := context.WithValue(req.Context(), UserIdKey, userID)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
