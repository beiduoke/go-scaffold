package data

import (
	"time"

	"gorm.io/gorm"
)

func NewSysModelMigrate() []interface{} {
	return []interface{}{
		&SysDomain{},
		&SysUser{},
		&SysApi{},
		&SysApiOperationLog{},
		&SysMenu{},
		&SysMenuButton{},
		&SysMenuParameter{},
		&SysAuthority{},
		&SysAuthorityMenuButton{},
		&SysDomainAuthorityUser{},
	}
}

type DomainModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DomainID  string         `gorm:"type:varchar(100);column:domain_id;not null;index;comment:领域ID;"`
}

type SysDomain struct {
	gorm.Model
	DomainID           string `gorm:"type:varchar(100);column:domain_id;not null;uniqueIndex;comment:领域ID;"`
	Name               string `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	State              int32  `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	DefaultAuthorityID uint   `gorm:"type:bigint(20);column:default_authority_id;not null;index;comment:默认角色;"`
	// Users       []SysUser      `gorm:"many2many:sys_domain_users;"`
	// Authorities []SysAuthority `gorm:"many2many:sys_domain_authorities;"`
	Users                []SysUser                `gorm:"-"`
	Authorities          []SysAuthority           `gorm:"-"`
	DomainAuthorityUsers []SysDomainAuthorityUser `gorm:"foreignKey:DomainID"`
}

// User 用户
type SysUser struct {
	gorm.Model
	Name     string     `gorm:"type:varchar(255);column:name;not null;index:idx_users_name_nick_name_real_name,priority:1;comment:名称;"`
	NickName string     `gorm:"type:varchar(255);column:nick_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:3;comment:昵称;"`
	RealName string     `gorm:"type:varchar(100);column:real_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:2;comment:实名;"`
	Password string     `gorm:"type:varchar(255);column:password;not null;default:'';comment:密码;"`
	Birthday *time.Time `gorm:"type:datetime;column:birthday;comment:生日;"`
	Gender   int32      `gorm:"type:tinyint(1);column:gender;not null;default:1;comment:性别 0 未指定 1 男 2 女;"`
	Mobile   string     `gorm:"type:varchar(20);column:mobile;not null;default:'';index:idx_users_mobile_email,priority:1;comment:手机号;"`
	Email    string     `gorm:"type:varchar(50);column:email;not null;default:'';index:idx_users_mobile_email,priority:2;comment:邮箱;"`
	State    int32      `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:用户状态 0 未指定  1 启用 2 停用;"`
	// Authorities []SysAuthority `gorm:"many2many:sys_authority_users;"`
	// Domains     []SysDomain    `gorm:"many2many:sys_domain_users;"`
	Authorities          []SysAuthority           `gorm:"-"`
	Domains              []SysDomain              `gorm:"-"`
	DomainAuthorityUsers []SysDomainAuthorityUser `gorm:"foreignKey:UserID"`
}

// Authority 角色
type SysAuthority struct {
	gorm.Model
	Name          string          `gorm:"type:varchar(255);column:name;not null;comment:角色名称;"`
	ParentID      uint            `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	DefaultRouter string          `gorm:"type:varchar(255);column:default_router;not null;default:'dashboard';comment:默认路由;"`
	Parent        *SysAuthority   `gorm:"foreignKey:ParentID"`
	Authorities   []SysAuthority  `gorm:"many2many:sys_authority_relations"`
	Menus         []SysMenu       `gorm:"many2many:sys_authority_menus;"`
	MenuButtons   []SysMenuButton `gorm:"many2many:sys_authority_menu_buttons;"`
	Apis          []SysApi        `gorm:"many2many:sys_authority_apis;"`
	// Users         []SysUser       `gorm:"many2many:sys_authority_users;"`
	// Domains       []SysDomain     `gorm:"many2many:sys_domain_authorities;"`
	Users                []SysUser                `gorm:"-"`
	Domains              []SysDomain              `gorm:"-"`
	DomainAuthorityUsers []SysDomainAuthorityUser `gorm:"foreignKey:AuthorityID"`
}

type SysDomainAuthorityUser struct {
	DomainID    uint `gorm:"type:bigint(20);column:domain_id;not null;uniqueIndex:idx_sys_domain_authority_users_domain_id_authority_id_user_id;comment:领域ID"`
	AuthorityID uint `gorm:"type:bigint(20);column:authority_id;not null;uniqueIndex:idx_sys_domain_authority_users_domain_id_authority_id_user_id;comment:角色ID"`
	UserID      uint `gorm:"type:bigint(20);column:user_id;not null;uniqueIndex:idx_sys_domain_authority_users_domain_id_authority_id_user_id;comment:用户ID"`
	CreatedAt   time.Time
	Domain      SysDomain    `gorm:"foreignKey:DomainID"`
	Authority   SysAuthority `gorm:"foreignKey:AuthorityID"`
	User        SysUser      `gorm:"foreignKey:UserID"`
}

// Api api接口
type SysApi struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	Path        string         `gorm:"type:varchar(255);column:path;not null;uniqueIndex:idx_api_path_method;comment:api路径"`
	Method      string         `gorm:"type:varchar(255);column:method;not null;default:POST;uniqueIndex:idx_api_path_method;comment:方法"`
	Group       string         `gorm:"type:varchar(255);column:group;not null;comment:api分组"`
	Description string         `gorm:"type:varchar(255);column:description;not null;default:'';comment:api描述"`
	Authorities []SysAuthority `gorm:"many2many:sys_authority_apis;"`
}

// SysMenu 菜单
type SysMenu struct {
	gorm.Model
	Name           string             `gorm:"type:varchar(255);column:name;not null;comment:路由名称;"`
	ParentID       uint               `gorm:"type:bigint(20);column:parent_id;not null;default:0;index;comment:父菜单ID"`
	Path           string             `gorm:"type:varchar(255);column:path;not null;comment:路由path"`
	Hidden         int32              `gorm:"type:tinyint(1);column:hidden;not null;default:1;comment:隐藏 0 无指定 1 是 2 否"`
	Component      string             `gorm:"type:varchar(255);column:component;not null;comment:对应前端文件路径"`
	Sort           int                `gorm:"type:int(10);column:sort;not null;default:10;comment:排序标记"`
	Meta           SysMeta            `gorm:"embedded;comment:附加属性"`
	Authorities    []SysAuthority     `gorm:"many2many:sys_authority_menus;"`
	MenuParameters []SysMenuParameter `gorm:"foreignKey:MenuID;"`
	MenuButtons    []SysMenuButton    `gorm:"foreignKey:MenuID;"`
}

// SysMeta 元数据
type SysMeta struct {
	Icon      string `gorm:"type:varchar(255);column:icon;not null;default:'';comment:图标"`
	Title     string `gorm:"type:varchar(255);column:title;not null;comment:标题"`
	KeepAlive int32  `gorm:"type:tinyint(1);column:keep_alive;not null;default:1;comment:是否缓存 0 无指定 1 是 2 否"`
	BaseMenu  int32  `gorm:"type:tinyint(1);column:base_menu;not null;default:1;comment:基础菜单 0 无指定 1 是 2 否"`
	CloseTab  int32  `gorm:"type:tinyint(1);column:close_tab;not null;default:1;comment:自动关闭TAB  0 无指定 1 是  2 否"`
}

// SysAuthorityMenu 角色菜单
type SysAuthorityMenu struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	AuthorityID    uint               `gorm:"type:bigint(20);column:authority_id;not null;uniqueIndex:idx_authority_menu_authority_id_menu_id;comment:角色ID"`
	MenuID         uint               `gorm:"type:bigint(20);column:menu_id;not null;uniqueIndex:idx_authority_menu_authority_id_menu_id;comment:菜单ID"`
	Menu           SysMenu            `gorm:"foreignKey:MenuID"`
	MenuParameters []SysMenuParameter `gorm:"foreignKey:MenuID"`
}

// SysAuthorityMenuButton 角色菜单按钮-自定义关联表
type SysAuthorityMenuButton struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	AuthorityID  uint          `gorm:"type:bigint(20);column:authority_id;not null;uniqueIndex:idx_authority_menu_authority_id_menu_id_menu_button_id;comment:角色ID"`
	MenuID       uint          `gorm:"type:bigint(20);column:menu_id;not null;uniqueIndex:idx_authority_menu_authority_id_menu_id_menu_button_id;comment:菜单ID"`
	MenuButtonID uint          `gorm:"type:bigint(20);column:menu_button_id;not null;uniqueIndex:idx_authority_menu_authority_id_menu_id_menu_button_id;comment:菜单按钮ID"`
	Menu         SysMenu       `gorm:"foreignKey:MenuID"`
	MenuButton   SysMenuButton `gorm:"foreignKey:MenuButtonID"`
}

// SysMenuButton 菜单按钮
type SysMenuButton struct {
	gorm.Model
	MenuID      uint           `gorm:"type:bigint(20);column:menu_id;not null;comment:菜单ID"`
	Name        string         `gorm:"type:varchar(255);column:name;not null;comment:按钮关键key;"`
	Remarks     string         `gorm:"type:varchar(255);column:remarks;not null;comment:按钮备注;"`
	Menu        SysMenu        `gorm:"foreignKey:MenuID;"`
	Authorities []SysAuthority `gorm:"many2many:sys_authority_menu_buttons;"`
}

// SysMenuParameter 菜单参数
type SysMenuParameter struct {
	gorm.Model
	MenuID uint   `gorm:"type:bigint(20);column:menu_id;not null;comment:菜单ID"`
	Type   int32  `gorm:"type:tinyint(1);column:type;not null;default:1;comment:地址栏携带参类型 0 未指定 1 params 2 query"`
	Key    string `gorm:"type:varchar(255);column:key;not null;default:'';comment:地址栏携带参数的key"`
	Value  string `gorm:"type:varchar(255);column:value;not null;default:'';comment:地址栏携带参数的值"`
}

// ApiOperationLog API 请求日志
type SysApiOperationLog struct {
	gorm.Model
	IP      string        `gorm:"type:varchar(100);column:ip;not null;comment:请求ip"`
	Method  string        `gorm:"type:varchar(255);column:method;not null;comment:请求方法"`
	Path    string        `gorm:"type:varchar(255);column:path;not null;comment:请求路径"`
	Status  int           `gorm:"type:int(10);column:status;not null;comment:请求状态"`
	Latency time.Duration `gorm:"type:int(10);column:latency;not null;default:0;comment:延迟"`
	Agent   string        `gorm:"type:varchar(255);column:agent;not null;comment:代理"`
	Error   string        `gorm:"type:varchar(255);column:error;not null;default:'';comment:错误信息"`
	Body    string        `gorm:"type:text;column:body;comment:请求Body"`
	Resp    string        `gorm:"type:text;column:resp;comment:响应Body"`
	UserID  uint          `gorm:"type:bigint(20);column:user_id;not null;comment:用户id"`
	User    SysUser       `gorm:"foreignKey:UserID;"`
}

// Jwt 黑名单
type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;column:jwt;comment:jwt;"`
}
