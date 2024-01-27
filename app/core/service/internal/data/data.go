package data

import (
	"context"
	"fmt"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/bwmarrin/snowflake"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	// 数据实例化
	NewData,
	// 数据库客户端
	NewEntClient,
	// redis客户端
	NewRedisClient,
	// 搜索引擎客户端
	NewMeilisearchClient,
	// 雪花ID生成器
	NewSnowflake,
	// 系统方法
	// 认证
	NewAuthRepo,
	// 用户
	NewUserRepo,
	// 角色
	NewRoleRepo,
)

// Data .
type Data struct {
	log *log.Helper
	cfg *conf.Bootstrap
	db  *ent.Client
	rdb *redis.Client
	sdb *meilisearch.Client
	sf  *snowflake.Node
}

// NewData .
func NewData(logger log.Logger, cfg *conf.Bootstrap, db *ent.Client, rdb *redis.Client, sdb *meilisearch.Client, sf *snowflake.Node) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{db: db, rdb: rdb, sdb: sdb, log: l, sf: sf, cfg: cfg}
	return d, func() {
		l.Info("closing db")
		if err := db.Close(); err != nil {
			log.Error(err)
		}
		l.Info("closing rdb")
		if err := rdb.Close(); err != nil {
			log.Error(err)
		}
		// l.Info("closing sdb")
	}, nil
}

// InTx 执行事务
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := ent.TxFromContext(ctx)
	if tx != nil {
		return fn(ctx)
	}

	tx, err := d.db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err = fn(ent.NewTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return err
}

func (d *Data) DB(ctx context.Context) *ent.Client {
	tx := ent.TxFromContext(ctx)
	if tx != nil {
		return tx.Client()
	}
	return d.db
}

// NewSnowflake 生成雪花算法id
func NewSnowflake(logger log.Logger) *snowflake.Node {
	l := log.NewHelper(log.With(logger, "module", "snowflake/initialize"))
	sf, err := snowflake.NewNode(1)
	if err != nil {
		l.Fatal("snowflake no init")
	}
	l.Infof("init snowflake ID：%s", sf.Generate())
	return sf
}

// NewGormClient 创建数据库客户端
func NewGormClient(cfg *conf.Bootstrap, logger log.Logger, models []interface{}) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "gorm/data/service"))
	return bootstrap.NewGormClient(cfg, l, models...)
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	// l := log.NewHelper(log.With(logger, "module", "redis/data/service"))
	return bootstrap.NewRedisClient(cfg.GetData())
}

// NewMeilisearchClient 创建Meilisearch客户端
func NewMeilisearchClient(cfg *conf.Bootstrap, logger log.Logger) *meilisearch.Client {
	l := log.NewHelper(log.With(logger, "module", "meilisearch/data/service"))
	return bootstrap.NewMeilisearchClient(cfg, l)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap, logger log.Logger) registry.Discovery {
	log.NewHelper(log.With(logger, "module", "discovery/data/service"))
	return bootstrap.NewConsulRegistry(cfg.Registry)
}
