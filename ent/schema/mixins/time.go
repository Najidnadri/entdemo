package mixins

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("dateCreated").
			Immutable().
			Annotations(
				entgql.OrderField("DATECREATED"),
			).
			Default(time.Now),
		field.Time("dateUpdated").
			Default(time.Now).
			Annotations(
				entgql.OrderField("DATEUPDATED"),
			).
			UpdateDefault(time.Now),
	}
}
