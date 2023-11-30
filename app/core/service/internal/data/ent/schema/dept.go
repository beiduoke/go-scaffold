package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Dept holds the schema definition for the Dept entity.
type Dept struct {
	ent.Schema
}

func (Dept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("部门表"),
	}
}

// Mixin of the Dept.
func (Dept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SnowflackId{},
		mixin.TimeAt{},
	}
}

// Fields of the Dept.
func (Dept) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名称").
			MaxLen(128).
			Optional().
			Nillable(),
	}
}

// Edges of the Dept.
func (Dept) Edges() []ent.Edge {
	return nil
}

// Indexes of the Dept.
func (Dept) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("phone", "authority").Unique(),
	}
}
