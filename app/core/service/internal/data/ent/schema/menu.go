package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("菜单表"),
	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinTop{},
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("菜单名称").
			Default("").
			MaxLen(32).
			NotEmpty().
			// Optional().
			Nillable(),

		field.String("title").
			Comment("菜单标题").
			Default("").
			NotEmpty().
			// Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("父级ID").
			Default(0).
			Optional().
			Nillable(),

		field.Int32("type").
			Comment("菜单类型 0 UNSPECIFIED, 目录 1 -> FOLDER, 菜单 2 -> MENU, 按钮 3 -> BUTTON").
			Default(1).
			// Optional().
			Nillable(),

		field.String("path").
			Comment("路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login").
			Default("").
			// Optional().
			Nillable(),

		field.String("component").
			Comment("前端页面组件").
			Default("").
			// Optional().
			Nillable(),

		field.String("icon").
			Comment("图标").
			Default("").
			MaxLen(128).
			// Optional().
			Nillable(),

		field.Bool("is_ext").
			Comment("是否外链").
			Default(false).
			// Optional().
			Nillable(),

		field.String("ext_url").
			Comment("外链地址").
			MaxLen(255).
			Default("").
			// Optional().
			Nillable(),

		field.Strings("permissions").
			Comment("权限代码 例如:sys:menu").
			SchemaType(map[string]string{
				dialect.MySQL:    "json",
				dialect.Postgres: "jsonb",
			}).
			Default([]string{}).
			Optional(),
		field.String("redirect").
			Comment("跳转路径").
			Default("").
			// Optional().
			Nillable(),
		field.String("current_active_menu").
			Comment("当前激活菜单").
			Default("").
			// Optional().
			Nillable(),

		field.Bool("keep_alive").
			Comment("是否缓存").
			Default(false).
			// Optional().
			Nillable(),

		field.Bool("visible").
			Comment("是否显示").
			Default(true).
			// Optional().
			Nillable(),

		field.Bool("hide_tab").
			Comment("是否显示在标签页导航").
			Default(true).
			// Optional().
			Nillable(),

		field.Bool("hide_menu").
			Comment("是否显示在菜单导航").
			Default(true).
			// Optional().
			Nillable(),

		field.Bool("hide_breadcrumb").
			Comment("是否显示在面包屑导航").
			Default(true).
			// Optional().
			Nillable(),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}

// Indexes of the Menu.
func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("phone", "authority").Unique(),
	}
}
