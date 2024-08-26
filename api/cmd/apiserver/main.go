package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/failuretoload/datamonster/config"
	"github.com/failuretoload/datamonster/ent"
	"github.com/failuretoload/datamonster/ent/migrate"
	"github.com/failuretoload/datamonster/server"
	_ "github.com/lib/pq"
)

var settings config.Settings

func init() {
	settings = config.GetSettings()
}

func main() {

	client, err := ent.Open(dialect.Postgres, settings.ConnString)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if schemaErr := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); schemaErr != nil {
		log.Println("db url", settings.ConnString)
		log.Fatal("opening ent client", schemaErr)
	}

	app := server.NewServer(client, settings.ClientUri)

	app.Run(settings.AuthKey)
}
