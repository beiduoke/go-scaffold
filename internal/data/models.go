package data

import (
	"time"

	"gorm.io/gorm"
)

func NewSysModelMigrate() []interface{} {
	return []interface{}{
		&SysDomain{},
		&SysUser{},
		&SysResource{},
		&SysApiOperationLog{},
		&SysMenu{},
		&SysMenuButton{},
		&SysMenuParameter{},
		&SysRole{},
		&SysRoleMenu{},
		&SysDepartment{},
		&SysPost{},
	}
}

type Model struct {
	ID        uint           `gorm:"primarykey;comment:主键ID;"`
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"comment:修改时间;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间;"`
}
type DomainModel struct {
	ID        uint           `gorm:"primarykey;comment:主键ID;"`
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"comment:修改时间;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间;"`
	DomainID  uint           `gorm:"type:bigint(20);column:domain_id;not null;default:0;index:idx_domain_id_data;comment:领域ID;"`
}

type SysDomain struct {
	Model
	Code          string     `gorm:"type:varchar(100);column:code;not null;index;comment:领域编码;"`
	Name          string     `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	ParentID      uint       `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	Sort          int32      `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State         int32      `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	DefaultRoleID uint       `gorm:"type:bigint(20);column:default_role_id;not null;index;comment:默认角色;"`
	Remarks       string     `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Menus         []SysMenu  `gorm:"many2many:sys_domain_menus;"`
	Parent        *SysDomain `gorm:"foreignKey:ParentID"`
	Users         []SysUser  `gorm:"-"`
	Roles         []SysRole  `gorm:"-"`
}

// User 用户
type SysUser struct {
	Model
	Avatar        string      `gorm:"type:varchar(255);column:avatar;not null;default:'';comment:头像;"`
	Name          string      `gorm:"type:varchar(255);column:name;not null;index:idx_users_name_nick_name_real_name,priority:1;comment:名称;"`
	NickName      string      `gorm:"type:varchar(255);column:nick_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:3;comment:昵称;"`
	RealName      string      `gorm:"type:varchar(100);column:real_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:2;comment:实名;"`
	Password      string      `gorm:"type:varchar(255);column:password;not null;default:'';comment:密码;"`
	Birthday      *time.Time  `gorm:"type:datetime;column:birthday;comment:生日;"`
	Gender        int32       `gorm:"type:tinyint(1);column:gender;not null;default:1;comment:性别 0 未指定 1 男 2 女;"`
	Mobile        string      `gorm:"type:varchar(20);column:mobile;not null;default:'';index:idx_users_mobile_email,priority:1;comment:手机号;"`
	Email         string      `gorm:"type:varchar(50);column:email;not null;default:'';index:idx_users_mobile_email,priority:2;comment:邮箱;"`
	State         int32       `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:用户状态 0 未指定  1 启用 2 停用;"`
	LastUseDomain uint        `gorm:"type:bigint(20);column:last_use_domain;not null;default:0;comment:最后使用租户"`
	LastLoginAt   *time.Time  `gorm:"type:datetime;column:last_login_at;comment:最后登录时间"`
	LastLoginIP   string      `gorm:"type:varchar(50);column:last_login_ip;not null;default:'';comment:最后登录IP"`
	Remarks       string      `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Posts         []SysPost   `gorm:"many2many:sys_user_posts;"`
	Roles         []SysRole   `gorm:"-"`
	Domains       []SysDomain `gorm:"-"`
}

// Role 角色
type SysRole struct {
	DomainModel
	Name          string          `gorm:"type:varchar(255);column:name;not null;index:idx_domain_id_data;comment:角色名称;"`
	ParentID      uint            `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	DefaultRouter string          `gorm:"type:varchar(255);column:default_router;not null;default:'/dashboard';comment:默认路由;"`
	Sort          int32           `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State         int32           `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:角色状态 0 未指定  1 启用 2 停用;"`
	Remarks       string          `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Parent        *SysRole        `gorm:"foreignKey:ParentID"`
	Roles         []SysRole       `gorm:"many2many:sys_role_relations"`
	Menus         []SysMenu       `gorm:"many2many:sys_role_menus;"`
	Resources     []SysResource   `gorm:"many2many:sys_role_resources;"`
	Departments   []SysDepartment `gorm:"many2many:sys_role_departments;"`
	Users         []SysUser       `gorm:"-"`
	Domains       []SysDomain     `gorm:"-"`
}

// Resource api资源
type SysResource struct {
	Model
	Name        string    `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	Path        string    `gorm:"type:varchar(255);column:path;not null;uniqueIndex:idx_api_path_method;comment:请求路径"`
	Method      string    `gorm:"type:varchar(255);column:method;not null;default:POST;uniqueIndex:idx_api_path_method;comment:方法"`
	Operation   string    `gorm:"type:varchar(255);column:operation;not null;default:'';uniqueIndex:idx_api_path_method;comment:请求动作"`
	Group       string    `gorm:"type:varchar(255);column:group;not null;comment:api分组"`
	Description string    `gorm:"type:varchar(255);column:description;not null;default:'';comment:api描述"`
	Remarks     string    `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Roles       []SysRole `gorm:"many2many:sys_role_resources;"`
	Menus       []SysMenu `gorm:"many2many:sys_menu_resources;"`
}

// SysMenu 菜单
type SysMenu struct {
	Model
	Name       string             `gorm:"type:varchar(255);column:name;not null;default:'';uniqueIndex:idx_menu_type_name;comment:路由名称;"`
	Type       int32              `gorm:"type:tinyint(1);column:type;not null;default:1;uniqueIndex:idx_menu_type_name;comment:菜单类型 0 无指定 1 目录 2 菜单 3 功能(按钮等);"`
	ParentID   uint               `gorm:"type:bigint(20);column:parent_id;not null;default:0;index;comment:父菜单ID"`
	Path       string             `gorm:"type:varchar(255);column:path;not null;default:'';comment:路由path"`
	Hidden     int32              `gorm:"type:tinyint(1);column:hidden;not null;default:1;comment:隐藏 0 无指定 1 是 2 否"`
	Component  string             `gorm:"type:varchar(255);column:component;not null;default:'';comment:对应前端文件路径"`
	Permission string             `gorm:"type:varchar(255);column:permission;not null;default:'';comment:权限标识"`
	Sort       int32              `gorm:"type:int(10);column:sort;not null;default:10;comment:排序标记"`
	Meta       SysMeta            `gorm:"embedded;comment:附加属性"`
	Remarks    string             `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Roles      []SysRole          `gorm:"many2many:sys_role_menus;"`
	Domains    []SysDomain        `gorm:"many2many:sys_domain_menus;"`
	Resources  []SysResource      `gorm:"many2many:sys_menu_resources;"`
	Parameters []SysMenuParameter `gorm:"foreignKey:MenuID;"`
	Buttons    []SysMenuButton    `gorm:"foreignKey:MenuID;"`
}

// SysMeta 元数据
type SysMeta struct {
	Icon      string `gorm:"type:varchar(255);column:icon;not null;default:'';comment:图标"`
	Title     string `gorm:"type:varchar(255);column:title;not null;comment:菜单标题"`
	KeepAlive int32  `gorm:"type:tinyint(1);column:keep_alive;not null;default:1;comment:是否缓存 0 无指定 1 是 2 否"`
	BaseMenu  int32  `gorm:"type:tinyint(1);column:base_menu;not null;default:1;comment:基础菜单 0 无指定 1 是 2 否"`
	CloseTab  int32  `gorm:"type:tinyint(1);column:close_tab;not null;default:1;comment:自动关闭TAB  0 无指定 1 是  2 否"`
	ExtType   int32  `gorm:"type:tinyint(1);column:ext_type;not null;default:1;comment:外链类型  0 无指定 1 无外链 2 新窗口 3 内框架"`
}

// SysRoleMenu 角色菜单-Many2Many 替换
type SysRoleMenu struct {
	CreatedAt     time.Time
	RoleID        uint    `gorm:"type:bigint(20);column:sys_role_id;not null;uniqueIndex:idx_role_menu_role_id_menu_id;comment:角色ID"`
	MenuID        uint    `gorm:"type:bigint(20);column:sys_menu_id;not null;uniqueIndex:idx_role_menu_role_id_menu_id;comment:菜单ID"`
	MenuButton    string  `gorm:"type:json;column:sys_menu_button;comment:菜单按钮"`
	MenuParameter string  `gorm:"type:json;column:sys_menu_parameter;comment:菜单参数"`
	Menu          SysMenu `gorm:"foreignKey:MenuID"`
}

// SysRoleMenuButton 角色菜单按钮-自定义关联表-未用
type SysRoleMenuButton struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	RoleID       uint          `gorm:"type:bigint(20);column:role_id;not null;uniqueIndex:idx_role_menu_role_id_menu_id_menu_button_id;comment:角色ID"`
	MenuID       uint          `gorm:"type:bigint(20);column:menu_id;not null;uniqueIndex:idx_role_menu_role_id_menu_id_menu_button_id;comment:菜单ID"`
	MenuButtonID uint          `gorm:"type:bigint(20);column:menu_button_id;not null;uniqueIndex:idx_role_menu_role_id_menu_id_menu_button_id;comment:菜单按钮ID"`
	Menu         SysMenu       `gorm:"foreignKey:MenuID"`
	MenuButton   SysMenuButton `gorm:"foreignKey:MenuButtonID"`
}

// SysMenuButton 菜单按钮
type SysMenuButton struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	MenuID    uint    `gorm:"type:bigint(20);column:menu_id;not null;index;comment:菜单ID"`
	Name      string  `gorm:"type:varchar(255);column:name;not null;comment:按钮关键key;"`
	Remarks   string  `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:按钮备注;"`
	Menu      SysMenu `gorm:"foreignKey:MenuID;"`
}

// SysMenuParameter 菜单参数
type SysMenuParameter struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	MenuID    uint    `gorm:"type:bigint(20);column:menu_id;not null;index;comment:菜单ID"`
	Type      int32   `gorm:"type:tinyint(1);column:type;not null;default:1;comment:地址栏携带参类型 0 未指定 1 params 2 query"`
	Name      string  `gorm:"type:varchar(255);column:name;not null;default:'';comment:地址栏携带参数的名称"`
	Value     string  `gorm:"type:varchar(255);column:value;not null;default:'';comment:地址栏携带参数的值"`
	Menu      SysMenu `gorm:"foreignKey:MenuID;"`
}

// SysDepartment 部门
type SysDepartment struct {
	DomainModel
	Name     string    `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	ParentID uint      `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	Sort     int32     `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	Remarks  string    `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	State    int32     `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	UserID   uint      `gorm:"type:bigint(20);column:user_id;not null;default:0;comment:用户id"`
	User     SysUser   `gorm:"foreignKey:UserID;"`
	Roles    []SysRole `gorm:"many2many:sys_role_departments;"`
	// User     SysUser `gorm:"many2many:sys_dept_users;"`
}

// Post 职位
type SysPost struct {
	DomainModel
	Name    string    `gorm:"type:varchar(255);column:name;not null;comment:职位名称;"`
	Code    string    `gorm:"type:varchar(100);column:code;not null;default:'';comment:职位编码;"`
	Sort    int32     `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State   int32     `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:职位状态 0 未指定  1 启用 2 停用;"`
	Remarks string    `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Users   []SysUser `gorm:"many2many:sys_user_posts;"`
}

// ApiOperationLog API 请求日志
type SysApiOperationLog struct {
	DomainModel
	IP      string        `gorm:"type:varchar(100);column:ip;not null;comment:请求ip"`
	Method  string        `gorm:"type:varchar(255);column:method;not null;comment:请求方法"`
	Path    string        `gorm:"type:varchar(255);column:path;not null;comment:请求路径"`
	Status  int           `gorm:"type:int(10);column:status;not null;comment:请求状态"`
	Latency time.Duration `gorm:"type:int(10);column:latency;not null;default:0;comment:延迟"`
	Agent   string        `gorm:"type:varchar(255);column:agent;not null;comment:代理"`
	Error   string        `gorm:"type:varchar(255);column:error;not null;default:'';comment:错误信息"`
	Body    string        `gorm:"type:text;column:body;comment:请求Body"`
	Resp    string        `gorm:"type:text;column:resp;comment:响应Body"`
	UserID  uint          `gorm:"type:bigint(20);column:user_id;not null;default:0;comment:用户id"`
	User    SysUser       `gorm:"foreignKey:UserID;"`
}

// Jwt 黑名单
type JwtBlacklist struct {
	Model
	UserID uint   `gorm:"type:bigint(20);column:user_id;not null;default:0;comment:用户id"`
	Jwt    string `gorm:"type:text;column:jwt;comment:jwt;"`
}
