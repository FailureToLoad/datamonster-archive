package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Survivor holds the schema definition for the Survivor entity.
type Survivor struct {
	ent.Schema
}

func (Survivor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).MinLen(1).Annotations(entgql.OrderField("NAME")),
		field.Int("born").NonNegative().Default(0).Max(35).Annotations(entgql.OrderField("BORN")),
		field.Enum("gender").NamedValues("male", "M", "female", "F").Default("F").Annotations(entgql.OrderField("GENDER")),
		field.Int("huntxp").NonNegative().Max(16).Default(0).Annotations(entgql.OrderField("HUNTXP")),
		field.Int("survival").NonNegative().Max(50).Default(0).Annotations(entgql.OrderField("SURVVAL")),
		field.Int("movement").Min(-20).Max(20).Default(5).Annotations(entgql.OrderField("MOVEMENT")),
		field.Int("accuracy").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("ACCURACY")),
		field.Int("strength").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("STRENGTH")),
		field.Int("evasion").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("EVASION")),
		field.Int("luck").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("LUCK")),
		field.Int("speed").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("SPEED")),
		field.Int("systemicpressure").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("SYSTEMICPRESSURE")),
		field.Int("torment").Min(-20).Max(20).Default(0).Annotations(entgql.OrderField("TORMENT")),
		field.Int("insanity").Min(0).Default(0).Annotations(entgql.OrderField("INSANITY")),
		field.Int("lumi").Min(0).Default(0).Annotations(entgql.OrderField("LUMI")),
		field.Int("courage").Min(0).Max(9).Default(0).Annotations(entgql.OrderField("CURRENCY")),
		field.Int("understanding").Min(0).Max(9).Default(0).Annotations(entgql.OrderField("UNDERSTANDING")),
		field.Int("settlement_id").Optional().Annotations(entgql.OrderField("SETTLEMENTID")),
	}
}

func (Survivor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("settlement", Settlement.Type).
			Ref("population").
			Unique().Field("settlement_id"),
	}
}

func (Survivor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
