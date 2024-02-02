table "bus_demos" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "名称"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_bus_demos_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_bus_demos_state" {
    columns = [column.state]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id]
  }
}
table "sys_api_operation_logs" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "ip" {
    null    = false
    type    = varchar(100)
    comment = "请求ip"
  }
  column "method" {
    null    = false
    type    = varchar(255)
    comment = "请求方法"
  }
  column "path" {
    null    = false
    type    = varchar(255)
    comment = "请求路径"
  }
  column "status" {
    null    = false
    type    = int
    comment = "请求状态"
  }
  column "latency" {
    null    = false
    type    = int
    default = 0
    comment = "延迟"
  }
  column "agent" {
    null    = false
    type    = varchar(255)
    comment = "代理"
  }
  column "error" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "错误信息"
  }
  column "body" {
    null    = true
    type    = text
    comment = "请求Body"
  }
  column "resp" {
    null    = true
    type    = text
    comment = "响应Body"
  }
  column "user_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "用户id"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id]
  }
  index "idx_sys_api_operation_logs_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "sys_casbin_rules" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "ptype" {
    null = true
    type = varchar(100)
  }
  column "v0" {
    null = true
    type = varchar(100)
  }
  column "v1" {
    null = true
    type = varchar(100)
  }
  column "v2" {
    null = true
    type = varchar(100)
  }
  column "v3" {
    null = true
    type = varchar(100)
  }
  column "v4" {
    null = true
    type = varchar(100)
  }
  column "v5" {
    null = true
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_casbin_rules" {
    unique  = true
    columns = [column.ptype, column.v0, column.v1, column.v2, column.v3, column.v4, column.v5]
  }
}
table "sys_depts" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "资源名称"
  }
  column "ancestors" {
    null    = false
    type    = varchar(100)
    default = "0"
    comment = "祖级列表"
  }
  column "parent_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "父角色ID"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "状态 0 未指定  1 启用 2 停用"
  }
  column "leader_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "负责人id"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id]
  }
  index "idx_sys_depts_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_depts_state" {
    columns = [column.state]
  }
}
table "sys_dicts" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "字典名称"
  }
  column "type" {
    null    = false
    type    = varchar(100)
    comment = "字典类型"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "字典状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_dicts_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_dicts_state" {
    columns = [column.state]
  }
}
table "sys_dict_data" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "dict_type" {
    null    = false
    type    = varchar(100)
    comment = "字典类型"
  }
  column "label" {
    null    = false
    type    = varchar(255)
    comment = "字典标签"
  }
  column "value" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "字典键值"
  }
  column "color_type" {
    null    = false
    type    = varchar(100)
    default = ""
    comment = "颜色类型"
  }
  column "css_class" {
    null    = false
    type    = varchar(100)
    default = ""
    comment = "CSS样式"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "字典状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_dict_data_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_dict_data_state" {
    columns = [column.state]
  }
}
table "sys_domains" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "领域名称"
  }
  column "parent_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "父角色ID"
  }
  column "super_user_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "超级用户ID"
  }
  column "code" {
    null    = false
    type    = varchar(100)
    comment = "领域编码"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "alias" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "领域别名"
  }
  column "logo" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "领域LOGO"
  }
  column "pic" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "领域主图"
  }
  column "keywords" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "领域关键字"
  }
  column "description" {
    null    = true
    type    = text
    comment = "描述"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_domains_code" {
    columns = [column.code]
  }
  index "idx_sys_domains_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_domains_state" {
    columns = [column.state]
  }
}
table "sys_domain_menus" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_menu_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_domain_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_menu_id, column.sys_domain_id]
  }
}
table "sys_menus" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "菜单/路由名称"
  }
  column "parent_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "父菜单ID"
  }
  column "type" {
    null    = false
    type    = bool
    default = 1
    comment = "菜单类型 0 无指定 1 目录 2 菜单 3 功能(按钮等)"
  }
  column "path" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "路由地址"
  }
  column "component" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "组件路径"
  }
  column "redirect" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "重定向地址"
  }
  column "permission" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "权限标识"
  }
  column "sort" {
    null    = false
    type    = int
    default = 10
    comment = "排序标记"
  }
  column "icon" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "图标"
  }
  column "title" {
    null    = false
    type    = varchar(255)
    comment = "菜单标题"
  }
  column "is_hidden" {
    null    = false
    type    = bool
    default = 1
    comment = "是否隐藏 0 无指定 1 是 2 否"
  }
  column "is_cache" {
    null    = false
    type    = bool
    default = 1
    comment = "是否缓存 0 无指定 1 是 2 否"
  }
  column "is_affix" {
    null    = false
    type    = bool
    default = 1
    comment = "是否固定 0 无指定 1 是 2 否"
  }
  column "link_type" {
    null    = false
    type    = bool
    default = 1
    comment = "外链类型  0 无指定 1 无 2 内嵌 3 跳转"
  }
  column "link_url" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "链接地址"
  }
  column "api_resource" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "接口资源"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_menu_type_name" {
    unique  = true
    columns = [column.name, column.type]
  }
  index "idx_sys_menus_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_menus_parent_id" {
    columns = [column.parent_id]
  }
}
table "sys_menu_buttons" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = true
    type = datetime(3)
  }
  column "menu_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "菜单ID"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "按钮关键key"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "按钮备注"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_menu_buttons_menu_id" {
    columns = [column.menu_id]
  }
}
table "sys_menu_parameters" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = true
    type = datetime(3)
  }
  column "menu_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "菜单ID"
  }
  column "type" {
    null    = false
    type    = bool
    default = 1
    comment = "地址栏携带参类型 0 未指定 1 params 2 query"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "地址栏携带参数的名称"
  }
  column "value" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "地址栏携带参数的值"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_sys_menu_parameters_menu_id" {
    columns = [column.menu_id]
  }
}
table "sys_menu_resources" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_menu_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_resource_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_menu_id, column.sys_resource_id]
  }
}
table "sys_posts" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "岗位名称"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "岗位状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "code" {
    null    = false
    type    = varchar(100)
    default = ""
    comment = "岗位编码"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id]
  }
  index "idx_sys_posts_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_posts_state" {
    columns = [column.state]
  }
}
table "sys_resources" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "资源名称"
  }
  column "path" {
    null    = false
    type    = varchar(255)
    comment = "请求路径"
  }
  column "method" {
    null    = false
    type    = varchar(255)
    default = "POST"
    comment = "方法"
  }
  column "operation" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "请求动作"
  }
  column "group" {
    null    = false
    type    = varchar(255)
    comment = "api分组"
  }
  column "description" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "api描述"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_api_path_method" {
    unique  = true
    columns = [column.path, column.method, column.operation]
  }
  index "idx_sys_resources_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "sys_roles" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "角色名称"
  }
  column "parent_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "父角色ID"
  }
  column "default_router" {
    null    = false
    type    = varchar(255)
    default = "/dashboard"
    comment = "默认路由"
  }
  column "sort" {
    null    = false
    type    = int
    default = 100
    comment = "排序"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "角色状态 0 未指定  1 启用 2 停用"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "data_scope" {
    null    = false
    type    = tinyint
    default = 1
    comment = "数据范围（0：未指定 1：本人数据权限 2：全部数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：自定部门数据权限 ）"
  }
  column "menu_check_strictly" {
    null    = false
    type    = tinyint
    default = 1
    comment = "菜单树选择项是否关联显示"
  }
  column "dept_check_strictly" {
    null    = false
    type    = tinyint
    default = 1
    comment = "部门树选择项是否关联显示"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id, column.name]
  }
  index "idx_sys_roles_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_roles_state" {
    columns = [column.state]
  }
}
table "sys_role_depts" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_dept_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_role_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_dept_id, column.sys_role_id]
  }
}
table "sys_role_menus" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_role_id" {
    null    = false
    type    = bigint
    comment = "角色ID"
  }
  column "sys_menu_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "菜单ID"
  }
  index "idx_role_menu_role_id_menu_id" {
    unique  = true
    columns = [column.sys_role_id, column.sys_menu_id]
  }
}
table "sys_role_relations" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_role_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "role_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_role_id, column.role_id]
  }
}
table "sys_role_resources" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_role_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_resource_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_role_id, column.sys_resource_id]
  }
}
table "sys_users" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    comment        = "主键ID"
    auto_increment = true
  }
  column "created_at" {
    null    = true
    type    = datetime(3)
    comment = "创建时间"
  }
  column "updated_at" {
    null    = true
    type    = datetime(3)
    comment = "修改时间"
  }
  column "deleted_at" {
    null    = true
    type    = datetime(3)
    comment = "删除时间"
  }
  column "domain_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "领域ID"
  }
  column "avatar" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "头像"
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "用户名称"
  }
  column "nick_name" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "昵称"
  }
  column "real_name" {
    null    = false
    type    = varchar(100)
    default = ""
    comment = "实名"
  }
  column "password" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "密码"
  }
  column "salt" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "密码加盐"
  }
  column "birthday" {
    null    = true
    type    = datetime
    comment = "生日"
  }
  column "gender" {
    null    = false
    type    = bool
    default = 1
    comment = "性别 0 未指定 1 男 2 女"
  }
  column "phone" {
    null    = false
    type    = varchar(20)
    default = ""
    comment = "手机号"
  }
  column "email" {
    null    = false
    type    = varchar(50)
    default = ""
    comment = "邮箱"
  }
  column "dept_id" {
    null    = false
    type    = bigint
    default = 0
    comment = "部门ID"
  }
  column "state" {
    null    = false
    type    = bool
    default = 1
    comment = "用户状态 0 未指定  1 启用 2 停用"
  }
  column "last_use_role_id" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
    comment  = "最后使用角色ID"
  }
  column "last_login_at" {
    null    = true
    type    = datetime
    comment = "最后登录时间"
  }
  column "last_login_ip" {
    null    = false
    type    = varchar(50)
    default = ""
    comment = "最后登录IP"
  }
  column "remarks" {
    null    = false
    type    = varchar(255)
    default = ""
    comment = "备注"
  }
  column "creator" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "创建者"
  }
  column "updater" {
    null    = false
    type    = varchar(64)
    default = ""
    comment = "更新者"
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_domain_id_data" {
    columns = [column.domain_id]
  }
  index "idx_sys_users_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_sys_users_state" {
    columns = [column.state]
  }
  index "idx_users_mobile_email" {
    columns = [column.phone, column.email]
  }
  index "idx_users_name_nick_name_real_name" {
    columns = [column.name, column.real_name, column.nick_name]
  }
  index "idx_users_phone_email" {
    columns = [column.phone, column.email]
  }
}
table "sys_user_posts" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_post_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_user_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_dict_id" {
    null     = true
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_post_id, column.sys_user_id]
  }
}
table "sys_user_roles" {
  schema  = schema.go_scaffold
  collate = "utf8mb4_general_ci"
  column "sys_role_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  column "sys_user_id" {
    null     = false
    type     = bigint
    unsigned = true
    comment  = "主键ID"
  }
  primary_key {
    columns = [column.sys_role_id, column.sys_user_id]
  }
}
schema "go_scaffold" {
  charset = "utf8mb4"
  collate = "utf8mb4_bin"
}
