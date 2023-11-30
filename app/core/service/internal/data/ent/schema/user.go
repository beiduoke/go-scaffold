package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户表"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SnowflackId{},
		mixin.TimeAt{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("用户名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),

		field.String("password").
			Comment("密码").
			MaxLen(255).
			Optional().
			Nillable().
			NotEmpty(),

		field.String("nickname").
			Comment("昵称").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("phone").
			Comment("手机号").
			Unique().
			MaxLen(20).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^1[3-9]{10}$")),
		field.String("email").
			Comment("电子邮箱").
			MaxLen(127).
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("头像").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Enum("authority").
			Comment("授权").
			Optional().
			Nillable().
			//SchemaType(map[string]string{
			//	dialect.MySQL:    "authority",
			//	dialect.Postgres: "authority",
			//}).
			Values(
				"SYS_ADMIN",
				"CUSTOMER_USER",
				"GUEST_USER",
				"REFRESH_TOKEN",
			).
			Default("CUSTOMER_USER"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone", "authority").Unique(),
	}
}
