package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/failuretoload/datamonster/ent"
)

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}
