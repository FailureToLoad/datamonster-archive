package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/failuretoload/datamonster/ent"
	"github.com/failuretoload/datamonster/ent/migrate"
	"github.com/failuretoload/datamonster/graph"
	"github.com/failuretoload/datamonster/helpers"
	"github.com/failuretoload/datamonster/server"
	_ "github.com/lib/pq"
	"log"
)

var (
	app server.Server
)

func init() {
	app = server.NewServer()
}

func main() {
	// Create ent.Client and run the schema migration.
	conn := helpers.SafeGetEnv("GQL_DEV")
	client, err := ent.Open(dialect.Postgres, conn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if schemaErr := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); schemaErr != nil {
		log.Fatal("opening ent client", schemaErr)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	appEnv := helpers.SafeGetEnv("APP_ENV")
	if appEnv == "schema" {
		app.Handle("/",
			playground.Handler("Datamonster", "/graphql"),
		)
	}

	app.Handle("/graphql", srv)
	app.Run()
}
