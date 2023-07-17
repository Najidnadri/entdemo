package schema

import (
	"entdemo/ent/privacy"
	mixins "entdemo/ent/schema/mixins"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
		mixins.IdMixin{},
		mixins.CommonAnnotationsMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("description").
			Annotations(
				entgql.OrderField("DESCRIPTION"),
			),
		field.Float("price").
			Positive().
			Annotations(
				entgql.OrderField("PRICE"),
			),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}

func (Product) Privacy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			// rules.DenyIfNoViewerr(),
			privacy.AlwaysDenyRule(),
		},
		Mutation: privacy.MutationPolicy{
			// rules.DenyIfNoViewerr(),
			privacy.AlwaysDenyRule(),
		},
	}
}
