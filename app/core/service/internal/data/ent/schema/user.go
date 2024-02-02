package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		MixinTop{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("用户名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Match(regexp.MustCompile("^[a-zA-Z0-9_]{4,16}$")).
			Nillable(),

		field.String("password").
			Comment("密码").
			Default("").
			MaxLen(255).
			Nillable(),

		field.String("nick_name").
			Comment("昵称").
			Default("").
			MaxLen(128).
			Nillable(),

		field.String("real_name").
			Comment("昵称").
			Default("").
			MaxLen(128).
			Nillable(),

		field.String("phone").
			Comment("手机号").
			Unique().
			MaxLen(20).
			Match(regexp.MustCompile("^1[0-9]{10}$")).
			Nillable(),

		field.String("email").
			Comment("电子邮箱").
			MaxLen(127).
			Default("").
			Nillable(),

		field.Time("birthday").
			Comment("生日").
			Default(time.Now()).
			SchemaType(map[string]string{
				dialect.MySQL:    "date", // Override MySQL.
				dialect.Postgres: "date", // Override Postgres.
			}).
			Nillable(),

		field.Int32("gender").
			Comment("性别 0 UNSPECIFIED, 1 -> MAN, 2 -> WOMAN").
			// Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "tinyint(2)",
				dialect.Postgres: "tinyint(2)",
			}).
			Default(1).
			Nillable(),

		field.String("avatar").
			Comment("头像").
			MaxLen(500).
			Default("").
			Nillable(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Default("").
			Nillable(),

		field.Int32("authority").
			Comment("授权 0 UNSPECIFIED, 1 -> SYS_ADMIN, 2 -> CUSTOMER_USER, 3 -> GUEST_USER, 4 -> REFRESH_TOKEN").
			// Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "tinyint(2)",
				dialect.Postgres: "tinyint(2)",
			}).
			Default(1).
			Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
		edge.To("posts", Post.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone", "name").Unique(),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	// Privacy policy defined in the BaseMixin and TenantMixin.
	return nil
}
