package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"github.com/tx7do/go-utils/sonyflake"
)

type PlatformId struct {
	mixin.Schema
}

func (PlatformId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("platform_id").
			Comment("平台ID").
			DefaultFunc(sonyflake.GenerateSonyflake).
			Positive().
			StructTag(`json:"platform_id,omitempty"`).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
			}),
	}
}

// Indexes of the PlatformId.
func (PlatformId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("platform_id"),
	}
}
