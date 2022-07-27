package data

import (
	"time"

	"gorm.io/gorm"
)

var migrates = []interface{}{
	&User{},
	&Role{},
	&UserRole{},
}

type User struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	Email  string `gorm:"type:varchar(255);column:email;not null;default:'';comment:邮箱;"`
	Switch int32  `gorm:"type:tinyint(1);column:switch;not null;default:1;comment:开关 0 -> 未指定  1 -> 启用 2 -> 停用;"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
}

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UserID    uint `gorm:"type:bigint(20);column:user_id;not null;uniqueIndex:idx_user_roles_user_id_role_id;comment:用户ID;"`
	RoleID    uint `gorm:"type:bigint(20);column:role_id;not null;uniqueIndex:idx_user_roles_user_id_role_id;comment:角色ID;"`
}
