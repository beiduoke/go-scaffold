package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Sort struct {
	mixin.Schema
}

func (Sort) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("sort").
			Comment("排序").
			Default(100).
			NonNegative().
			SchemaType(map[string]string{
				dialect.MySQL: "int",
			}).
			Nillable(),
	}
}
