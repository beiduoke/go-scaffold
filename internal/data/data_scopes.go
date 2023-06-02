package data

import (
	"gorm.io/gorm"
)

// DBScopesDomain 获取租户
func DBScopesDomain(id ...uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("domain_id", id)
	}
}

// DBScopesUser 获取用户
func DBScopesUser(id ...uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id", id)
	}
}

// DBScopesDept 获取部门
func DBScopesDept(id ...uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("dept_id", id)
	}
}

func DBScopesOmitUpdate(args ...string) func(db *gorm.DB) *gorm.DB {
	args = append(args, []string{"ID", "CreatedAt", "DomainID"}...)
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(args...)
	}
}
