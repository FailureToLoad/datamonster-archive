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

var (
	ready = false
)

func main() {

	log.Println("opening db connection")
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

	mode := config.Mode()
	if mode == "schema" {
		router.Handle("/",
			playground.Handler("Datamonster.gql", "/graphql"),
		)
	}
	router.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/startup", func(w http.ResponseWriter, r *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("api is not ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	router.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("api is not ready"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	router.Mount("/graphql", graphqlRouter(mode, srv))

	return Server{
		Mux: router,
	}
}

func graphqlRouter(mode string, srv *handler.Server) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(CorsHandler())
	if mode != "schema" {
		r.Use(SecureOptions())
		r.Use(CacheControl)
		r.Use(clerkhttp.WithHeaderAuthorization())
		r.Use(UserIDExtractor)
	}
	r.Handle("/", srv)
	return r
}

func (s Server) Handle(route string, handler http.Handler) {
	s.Mux.Handle(route, handler)
}

func (s Server) Run() {
	clerk.SetKey(config.Key())
	ready = true
	log.Default().Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", s.Mux)
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
