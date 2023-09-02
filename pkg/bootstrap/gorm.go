package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/beiduoke/go-scaffold/api/common/conf"
)

// NewGormClient 创建数据库客户端
func NewGormClient(cfg *conf.Bootstrap, l *log.Helper, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Data.Database.Source), &gorm.Config{
		Logger:         gormLogger.Default,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "scaffold_", // table name prefix, table for `User` would be `t_users`
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		l.Fatalf("failed opening connection to mysql: %v", err)
	}
	if cfg.Data.Database.Migrate {
		if err := db.AutoMigrate(models...); err != nil {
			l.Fatal(err)
		}
	}
	return db
}
