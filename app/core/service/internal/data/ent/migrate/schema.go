// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DeptsColumns holds the columns for the "depts" table.
	DeptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString, Size: 128, Comment: "名称"},
		{Name: "ancestors", Type: field.TypeJSON, Nullable: true, Comment: "祖级列表"},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "父级ID", Default: 0, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// DeptsTable holds the schema information for the "depts" table.
	DeptsTable = &schema.Table{
		Name:       "depts",
		Comment:    "部门表",
		Columns:    DeptsColumns,
		PrimaryKey: []*schema.Column{DeptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "depts_depts_children",
				Columns:    []*schema.Column{DeptsColumns[9]},
				RefColumns: []*schema.Column{DeptsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50, Comment: "会员名"},
		{Name: "password", Type: field.TypeString, Size: 255, Comment: "密码", Default: ""},
		{Name: "nickname", Type: field.TypeString, Size: 128, Comment: "昵称", Default: ""},
		{Name: "phone", Type: field.TypeString, Unique: true, Size: 20, Comment: "手机号"},
		{Name: "email", Type: field.TypeString, Size: 127, Comment: "电子邮箱", Default: ""},
		{Name: "avatar", Type: field.TypeString, Size: 500, Comment: "头像", Default: ""},
		{Name: "description", Type: field.TypeString, Size: 1023, Comment: "个人说明", Default: ""},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:       "members",
		Comment:    "会员表",
		Columns:    MembersColumns,
		PrimaryKey: []*schema.Column{MembersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "member_phone",
				Unique:  true,
				Columns: []*schema.Column{MembersColumns[10]},
			},
		},
	}
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString, Size: 32, Comment: "菜单名称", Default: ""},
		{Name: "title", Type: field.TypeString, Comment: "菜单标题", Default: ""},
		{Name: "type", Type: field.TypeInt32, Comment: "菜单类型 0 UNSPECIFIED, 目录 1 -> FOLDER, 菜单 2 -> MENU, 按钮 3 -> BUTTON", Default: 1},
		{Name: "path", Type: field.TypeString, Comment: "路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login", Default: ""},
		{Name: "component", Type: field.TypeString, Comment: "前端页面组件", Default: ""},
		{Name: "icon", Type: field.TypeString, Size: 128, Comment: "图标", Default: ""},
		{Name: "is_ext", Type: field.TypeBool, Comment: "是否外链", Default: false},
		{Name: "ext_url", Type: field.TypeString, Size: 255, Comment: "外链地址", Default: ""},
		{Name: "permissions", Type: field.TypeJSON, Nullable: true, Comment: "权限代码 例如:sys:menu", SchemaType: map[string]string{"mysql": "json", "postgres": "jsonb"}},
		{Name: "redirect", Type: field.TypeString, Comment: "跳转路径", Default: ""},
		{Name: "current_active_menu", Type: field.TypeString, Comment: "当前激活菜单", Default: ""},
		{Name: "keep_alive", Type: field.TypeBool, Comment: "是否缓存", Default: false},
		{Name: "visible", Type: field.TypeBool, Comment: "是否显示", Default: true},
		{Name: "hide_tab", Type: field.TypeBool, Comment: "是否显示在标签页导航", Default: true},
		{Name: "hide_menu", Type: field.TypeBool, Comment: "是否显示在菜单导航", Default: true},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Comment: "是否显示在面包屑导航", Default: true},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "父级ID", Default: 0, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Comment:    "菜单表",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menus_menus_children",
				Columns:    []*schema.Column{MenusColumns[23]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "名称"},
		{Name: "user_posts", Type: field.TypeUint32, Nullable: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Comment:    "岗位表",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString, Size: 128, Comment: "名称"},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Comment:    "角色表",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Comment:    "租户表",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "sort", Type: field.TypeInt32, Comment: "排序", Default: 100, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "state", Type: field.TypeInt32, Comment: "状态 0 UNSPECIFIED 开启 1 -> ACTIVE 关闭 2 -> INACTIVE, 禁用 3 -> BANNED", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 50, Comment: "用户名"},
		{Name: "password", Type: field.TypeString, Size: 255, Comment: "密码", Default: ""},
		{Name: "nick_name", Type: field.TypeString, Size: 128, Comment: "昵称", Default: ""},
		{Name: "real_name", Type: field.TypeString, Size: 128, Comment: "昵称", Default: ""},
		{Name: "phone", Type: field.TypeString, Unique: true, Size: 20, Comment: "手机号"},
		{Name: "email", Type: field.TypeString, Size: 127, Comment: "电子邮箱", Default: ""},
		{Name: "birthday", Type: field.TypeTime, Comment: "生日", SchemaType: map[string]string{"mysql": "date", "postgres": "date"}},
		{Name: "gender", Type: field.TypeInt32, Comment: "性别 0 UNSPECIFIED, 1 -> MAN, 2 -> WOMAN", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)", "postgres": "tinyint(2)"}},
		{Name: "avatar", Type: field.TypeString, Size: 500, Comment: "头像", Default: ""},
		{Name: "description", Type: field.TypeString, Size: 1023, Comment: "个人说明", Default: ""},
		{Name: "authority", Type: field.TypeInt32, Comment: "授权 0 UNSPECIFIED, 1 -> SYS_ADMIN, 2 -> CUSTOMER_USER, 3 -> GUEST_USER, 4 -> REFRESH_TOKEN", Default: 1, SchemaType: map[string]string{"mysql": "tinyint(2)", "postgres": "tinyint(2)"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户表",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_phone_name",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[11], UsersColumns[7]},
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint32, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "role_id", Type: field.TypeUint32, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DeptsTable,
		MembersTable,
		MenusTable,
		PostsTable,
		RolesTable,
		TenantsTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	DeptsTable.ForeignKeys[0].RefTable = DeptsTable
	DeptsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	MembersTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	MenusTable.ForeignKeys[0].RefTable = MenusTable
	MenusTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	RolesTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	TenantsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}
