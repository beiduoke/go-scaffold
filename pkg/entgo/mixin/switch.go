package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Switch struct {
	mixin.Schema
}

func (Switch) Fields() []ent.Field {
	return []ent.Field{
		field.Int("switch").
			Comment("开关").
			Default(1).
			Optional().
			Nillable(),
	}
}
