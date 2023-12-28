package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type State struct {
	mixin.Schema
}

func (State) Fields() []ent.Field {
	return []ent.Field{
		field.Int("state").
			Comment("状态").
			Default(1).
			Optional().
			Nillable(),
	}
}
