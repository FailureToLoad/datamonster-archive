package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/failuretoload/datamonster/graph"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"

	"entgo.io/ent/dialect"
	"github.com/failuretoload/datamonster/config"
	"github.com/failuretoload/datamonster/ent"
	"github.com/failuretoload/datamonster/ent/migrate"
	_ "github.com/lib/pq"
)

func main() {

	client, err := ent.Open(dialect.Postgres, config.PGConn())
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if schemaErr := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); schemaErr != nil {
		log.Println("db url", config.PGConn())
		log.Fatal("opening ent client", schemaErr)
	}

	app := NewServer(client)

	app.Run()
}

type Server struct {
	Mux *chi.Mux
}

func NewServer(client *ent.Client) Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(10 * time.Second))
	router.Use(CorsHandler())
	mode := config.Mode()
	if mode != "schema" {
		router.Use(SecureOptions())
		router.Use(CacheControl)
		router.Use(clerkhttp.WithHeaderAuthorization())
		router.Use(UserIDExtractor)
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

func (s Server) Run() {
	clerk.SetKey(config.Key())
	port := "0.0.0.0:8080"
	if config.Mode() == "prod" {
		port = "0.0.0.0:80"
	}

	log.Default().Println("Starting server on ", port)
	err := http.ListenAndServe(port, s.Mux)
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
		AllowedOrigins:   []string{"https://data-monster.net", "http://datamonster-web"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           3599, // Maximum value not ignored by any of major browsers
	})
	return c.Handler
}

func UserIDExtractor(next http.Handler) http.Handler {
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
		ctx := context.WithValue(req.Context(), config.UserIDKey, userID)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func Unauthorized(w http.ResponseWriter, message *string) {
	content := "unauthorized"
	if message != nil {
		content = content + ": " + *message
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte(content))
	log.Println(content)
}
