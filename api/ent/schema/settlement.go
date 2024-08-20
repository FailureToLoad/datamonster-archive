package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Settlement holds the schema definition for the Settlement entity.
type Settlement struct {
	ent.Schema
}

// Fields of the Settlement.
func (Settlement) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").NotEmpty().Annotations(entgql.OrderField("OWNER")),
		field.String("name").MaxLen(50).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Int("survivalLimit").Min(0).Default(0).Annotations(entgql.OrderField("SURVIVAL_LIMIT")),
		field.Int("departingSurvival").Min(0).Default(0).Annotations(entgql.OrderField("DEPARTING_SURVIVAL")),
		field.Int("collectiveCognition").Min(0).Max(50).Default(0).Annotations(entgql.OrderField("COLLECTIVE_COGNITION")),
		field.Int("currentYear").Min(0).Max(35).Default(0).Annotations(entgql.OrderField("CURRENT_YEAR")),
	}
}

// Edges of the Settlement.
func (Settlement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("population", Survivor.Type),
	}
}

func (Settlement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
