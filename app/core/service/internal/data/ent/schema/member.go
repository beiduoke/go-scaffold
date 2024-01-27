package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("会员表"),
	}
}

// Mixin of the Member.
func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinTop{},
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("会员名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Match(regexp.MustCompile("^[a-zA-Z0-9_]{4,16}$")),

		field.String("password").
			Comment("密码").
			Default("").
			MaxLen(255).
			NotEmpty(),

		field.String("nickname").
			Comment("昵称").
			Default("").
			MaxLen(128).
			NotEmpty(),

		field.String("phone").
			Comment("手机号").
			Unique().
			MaxLen(20).
			NotEmpty().
			Match(regexp.MustCompile("^1[3-9]{10}$")),
		field.String("email").
			Comment("电子邮箱").
			MaxLen(127).
			Default("").
			NotEmpty(),

		field.String("avatar").
			Comment("头像").
			MaxLen(500).
			Default("").
			NotEmpty(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Default("").
			NotEmpty(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}

// Indexes of the Member.
func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone").Unique(),
	}
}

// Policy defines the privacy policy of the Member.
func (Member) Policy() ent.Policy {
	// Privacy policy defined in the BaseMixin and TenantMixin.
	return nil
}
