package data

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
)

// NewEntClient .
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *ent.Client {
	l := log.NewHelper(log.With(logger, "module", "ent/data/service"))

	drv, err := sql.Open(cfg.Data.Database.Driver, cfg.Data.Database.Source)
	if err != nil {
		l.Fatalf("failed opening connection to %s: %v", cfg.Data.Database.Driver, err)
		return nil
	}

	{
		db := drv.DB()
		// 连接池中最多保留的空闲连接数量
		db.SetMaxIdleConns(int(cfg.Data.Database.MaxIdleConnections))
		// 连接池在同一时间打开连接的最大数量
		db.SetMaxOpenConns(int(cfg.Data.Database.MaxOpenConnections))
		// 连接可重用的最大时间长度
		db.SetConnMaxLifetime(cfg.Data.Database.ConnectionMaxLifetime.AsDuration())
	}

	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)

	// 运行数据库迁移工具
	if cfg.Data.Database.Migrate {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return client
}
