package data

import (
	"time"

	"gorm.io/gorm"
)

func NewSysModelMigrate() []interface{} {
	return []interface{}{
		&SysApiOperationLog{},
		&SysDomain{},
		&SysDomainPackage{},
		&SysUser{},
		&SysMenu{},
		&SysRole{},
		&SysDept{},
		&SysPost{},
		&SysDict{},
		&SysDictData{},
	}
}

type Model struct {
	ID        uint           `gorm:"primarykey;comment:主键ID;"`
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	Creator   string         `gorm:"type:varchar(64);column:creator;not null;default:'';comment:创建者;"`
	UpdatedAt time.Time      `gorm:"comment:修改时间;"`
	Updater   string         `gorm:"type:varchar(64);column:updater;not null;default:'';comment:更新者;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间;"`
}

type DomainModel struct {
	ID        uint           `gorm:"primarykey;comment:主键ID;"`
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	Creator   string         `gorm:"type:varchar(64);column:creator;not null;default:'';comment:创建者;"`
	UpdatedAt time.Time      `gorm:"comment:修改时间;"`
	Updater   string         `gorm:"type:varchar(64);column:updater;not null;default:'';comment:更新者;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间;"`
	DomainID  uint           `gorm:"type:bigint(20);column:domain_id;not null;default:0;index:idx_domain_id_data;comment:租户ID;"`
	Domain    *SysDomain     `gorm:"foreignKey:DomainID"`
}

// func (dm *DomainModel) BeforeUpdate(tx *gorm.DB) (err error) {
// 	// tx.Statement.Context
// 	// if dm.Updater == "admin" {
// 	// 	return errors.New("admin user not allowed to update")
// 	// }
// 	fmt.Println("更新前")
// 	return
// }

// func (dm *DomainModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	// if dm.Creator == "admin" {
// 	// 	return errors.New("invalid role")
// 	// }
// 	fmt.Println("创建前")
// 	return
// }

// SysDomain 租户
type SysDomain struct {
	Model
	ParentID    uint              `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	SuperUserID uint              `gorm:"type:bigint(20);column:super_user_id;not null;default:0;comment:超级用户ID"`
	PackageID   uint              `gorm:"type:bigint(20);column:package_id;not null;comment:租户套餐ID"`
	Name        string            `gorm:"type:varchar(255);column:name;not null;comment:租户名称;"`
	Code        string            `gorm:"type:varchar(100);column:code;not null;index;comment:租户编码;"`
	Alias       string            `gorm:"type:varchar(255);column:alias;not null;default:'';comment:租户别名;"`
	Keywords    string            `gorm:"type:varchar(255);column:keywords;not null;default:'';comment:租户关键字;"`
	Logo        string            `gorm:"type:varchar(255);column:logo;not null;default:'';comment:租户LOGO;"`
	Pic         string            `gorm:"type:varchar(255);column:pic;not null;default:'';comment:租户主图;"`
	Description string            `gorm:"type:text;column:description;comment:描述"`
	Sort        int32             `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State       int32             `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	Remarks     string            `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Menus       []SysMenu         `gorm:"many2many:sys_domain_menus;"`
	SuperUser   *SysUser          `gorm:"foreignKey:SuperUserID"`
	Parent      *SysDomain        `gorm:"foreignKey:ParentID"`
	Package     *SysDomainPackage `gorm:"foreignKey:PackageID"`
	Users       []SysUser         `gorm:"-"`
	Roles       []SysRole         `gorm:"-"`
}

// User 用户
type SysUser struct {
	DomainModel
	Name          string     `gorm:"type:varchar(255);column:name;not null;index:idx_users_name_nick_name_real_name,priority:1;comment:用户名称;"`
	NickName      string     `gorm:"type:varchar(255);column:nick_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:3;comment:昵称;"`
	RealName      string     `gorm:"type:varchar(100);column:real_name;not null;default:'';index:idx_users_name_nick_name_real_name,priority:2;comment:实名;"`
	Avatar        string     `gorm:"type:varchar(255);column:avatar;not null;default:'';comment:头像;"`
	Password      string     `gorm:"type:varchar(255);column:password;not null;default:'';comment:密码;"`
	PasswordSalt  string     `gorm:"type:varchar(255);column:salt;not null;default:'';comment:密码加盐;"`
	Birthday      *time.Time `gorm:"type:datetime;column:birthday;comment:生日;"`
	Gender        int32      `gorm:"type:tinyint(1);column:gender;not null;default:1;comment:性别 0 未指定 1 男 2 女;"`
	Phone         string     `gorm:"type:varchar(20);column:phone;not null;default:'';index:idx_users_phone_email,priority:1;comment:手机号;"`
	Email         string     `gorm:"type:varchar(50);column:email;not null;default:'';index:idx_users_phone_email,priority:2;comment:邮箱;"`
	DeptID        uint       `gorm:"type:bigint(20);column:dept_id;not null;default:0;comment:部门ID"`
	State         int32      `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:用户状态 0 未指定  1 启用 2 停用;"`
	LastUseRoleID uint       `gorm:"type:bigint(20);column:last_use_role_id;not null;default:0;comment:最后使用角色ID"`
	LastLoginAt   *time.Time `gorm:"type:datetime;column:last_login_at;comment:最后登录时间"`
	LastLoginIP   string     `gorm:"type:varchar(50);column:last_login_ip;not null;default:'';comment:最后登录IP"`
	Remarks       string     `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Posts         []SysPost  `gorm:"many2many:sys_user_posts;"`
	Roles         []SysRole  `gorm:"many2many:sys_user_roles;"`
	Dept          *SysDept   `gorm:"foreignKey:DeptID"`
	LastUseRole   *SysRole   `gorm:"foreignKey:LastUseRoleID"`
}

// Role 角色
type SysRole struct {
	DomainModel
	Name              string      `gorm:"type:varchar(255);column:name;not null;index:idx_domain_id_data;comment:角色名称;"`
	ParentID          uint        `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	DefaultRouter     string      `gorm:"type:varchar(255);column:default_router;not null;default:'/dashboard';comment:默认路由;"`
	Sort              int32       `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	DataScope         int32       `gorm:"type:tinyint(2);column:data_scope;not null;default:1;comment:数据范围（0：未指定 1：本人数据权限 2：全部数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：自定部门数据权限 ）;"`
	MenuCheckStrictly int32       `gorm:"type:tinyint(2);column:menu_check_strictly;not null;default:1;comment:菜单树选择项是否关联显示;"`
	DeptCheckStrictly int32       `gorm:"type:tinyint(2);column:dept_check_strictly;not null;default:1;comment:部门树选择项是否关联显示;"`
	State             int32       `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:角色状态 0 未指定  1 启用 2 停用;"`
	Remarks           string      `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Parent            *SysRole    `gorm:"foreignKey:ParentID"`
	Roles             []SysRole   `gorm:"many2many:sys_role_relations"`
	Menus             []SysMenu   `gorm:"many2many:sys_role_menus;"`
	Depts             []SysDept   `gorm:"many2many:sys_role_depts;joinForeignKey:SysRoleID"`
	Users             []SysUser   `gorm:"many2many:sys_user_roles;"`
	Domains           []SysDomain `gorm:"-"`
}

// SysMenu 菜单
type SysMenu struct {
	Model
	ParentID       uint               `gorm:"type:bigint(20);column:parent_id;not null;default:0;index;comment:父菜单ID"`
	Type           int32              `gorm:"type:tinyint(1);column:type;not null;default:1;uniqueIndex:idx_menu_type_name;comment:菜单类型 0 无指定 1 目录 2 菜单 3 功能(按钮等);"`
	Title          string             `gorm:"type:varchar(255);column:title;not null;comment:菜单标题"`
	Name           string             `gorm:"type:varchar(255);column:name;not null;default:'';uniqueIndex:idx_menu_type_name;comment:菜单/路由名称;"`
	Path           string             `gorm:"type:varchar(255);column:path;not null;default:'';comment:路由地址"`
	Redirect       string             `gorm:"type:varchar(255);column:redirect;not null;default:'';comment:重定向地址"`
	Component      string             `gorm:"type:varchar(255);column:component;not null;default:'';comment:组件路径"`
	Permission     string             `gorm:"type:varchar(255);column:permission;not null;default:'';comment:权限标识"`
	Icon           string             `gorm:"type:varchar(255);column:icon;not null;default:'';comment:图标"`
	Sort           int32              `gorm:"type:int(10);column:sort;not null;default:10;comment:排序标记"`
	IsHidden       int32              `gorm:"type:tinyint(1);column:is_hidden;not null;default:1;comment:是否隐藏 0 无指定 1 是 2 否"`
	IsCache        int32              `gorm:"type:tinyint(1);column:is_cache;not null;default:1;comment:是否缓存 0 无指定 1 是 2 否"`
	IsAffix        int32              `gorm:"type:tinyint(1);column:is_affix;not null;default:1;comment:是否固定 0 无指定 1 是 2 否"`
	LinkType       int32              `gorm:"type:tinyint(1);column:link_type;not null;default:1;comment:外链类型  0 无指定 1 无 2 内嵌 3 跳转"`
	LinkUrl        string             `gorm:"type:varchar(255);column:link_url;not null;default:'';comment:链接地址"`
	ApiResource    string             `gorm:"type:varchar(255);column:api_resource;not null;default:'';comment:接口资源"`
	Remarks        string             `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Roles          []SysRole          `gorm:"many2many:sys_role_menus;"`
	DomainPackages []SysDomainPackage `gorm:"many2many:sys_domain_package_menus;"`
}

// SysDept 部门
type SysDept struct {
	DomainModel
	Name      string    `gorm:"type:varchar(255);column:name;not null;comment:资源名称;"`
	Ancestors string    `gorm:"type:varchar(100);column:ancestors;not null;default:'0';comment:祖级列表;"`
	ParentID  uint      `gorm:"type:bigint(20);column:parent_id;not null;default:0;comment:父角色ID"`
	Sort      int32     `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	Remarks   string    `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	State     int32     `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	LeaderID  uint      `gorm:"type:bigint(20);column:leader_id;not null;default:0;comment:负责人id"`
	Leader    *SysUser  `gorm:"foreignKey:LeaderID;"`
	Roles     []SysRole `gorm:"many2many:sys_role_depts"`
}

// Post 岗位
type SysPost struct {
	DomainModel
	Name    string    `gorm:"type:varchar(255);column:name;not null;comment:岗位名称;"`
	Code    string    `gorm:"type:varchar(100);column:code;not null;default:'';comment:岗位编码;"`
	Sort    int32     `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State   int32     `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:岗位状态 0 未指定  1 启用 2 停用;"`
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
	User    *SysUser      `gorm:"foreignKey:UserID;"`
}

// Jwt 黑名单
type JwtBlacklist struct {
	Model
	UserID uint   `gorm:"type:bigint(20);column:user_id;not null;default:0;comment:用户id"`
	Jwt    string `gorm:"type:text;column:jwt;comment:jwt;"`
}

// Dict 字典
type SysDict struct {
	Model
	Name        string        `gorm:"type:varchar(255);column:name;not null;comment:字典名称;"`
	Type        string        `gorm:"type:varchar(100);column:type;not null;comment:字典类型;"`
	Sort        int32         `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State       int32         `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:字典状态 0 未指定  1 启用 2 停用;"`
	Remarks     string        `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	SysDictData []SysDictData `gorm:"foreignKey:DictType;references:Type"`
}

// DictData 字典数据
type SysDictData struct {
	Model
	DictType  string   `gorm:"type:varchar(100);column:dict_type;not null;comment:字典类型;"`
	Label     string   `gorm:"type:varchar(255);column:label;not null;comment:字典标签;"`
	Value     string   `gorm:"type:varchar(255);column:value;not null;default:'';comment:字典键值;"`
	ColorType string   `gorm:"type:varchar(100);column:color_type;not null;default:'';comment:颜色类型;"`
	CssClass  string   `gorm:"type:varchar(100);column:css_class;not null;default:'';comment:CSS样式;"`
	Sort      int32    `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State     int32    `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:字典状态 0 未指定  1 启用 2 停用;"`
	Remarks   string   `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	SysDict   *SysDict `gorm:"foreignKey:Type;references:DictType"`
}

// DomainPackage 租户套餐
type SysDomainPackage struct {
	Model
	Name    string    `gorm:"type:varchar(255);column:name;not null;comment:套餐名称;"`
	Sort    int32     `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State   int32     `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:套餐状态 0 未指定  1 启用 2 停用;"`
	Remarks string    `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
	Menus   []SysMenu `gorm:"many2many:sys_domain_package_menus;"`
}
