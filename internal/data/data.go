package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	// 基础配置
	NewGormClient,
	NewRedisClient,
	NewSDB,
	NewSnowflake,
	NewModelMigrate,
	NewData,
	NewTransaction,
	// 认证
	NewAuthenticator,
	NewSecurityUser,
	// 鉴权
	NewAuthModel,
	NewAuthAdapter,
	NewWatcher,
	NewAuthEnforcer,
	NewAuthCasbin,
	// 数据操作
	NewDomainRepo,
	NewRoleRepo,
	NewMenuRepo,
	NewUserRepo,
	NewDeptRepo,
	NewPostRepo,
	NewDictRepo,
)

// NewModelMigrate 数据模型迁移
func NewModelMigrate() []interface{} {
	migrates := NewSysModelMigrate()
	return append(migrates, NewBusModelMigrate()...)
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

// NewTransaction 事务
func NewTransaction(d *Data) biz.Transaction {
	return d
}

type ConfigOptions struct {
	system *conf.System
}

// Data .
type Data struct {
	conf     ConfigOptions
	log      *log.Helper
	db       *gorm.DB
	rdb      *redis.Client
	sdb      *meilisearch.Client
	sf       *snowflake.Node
	enforcer *casbin.SyncedEnforcer
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, sdb *meilisearch.Client, enforcer casbin.IEnforcer, sf *snowflake.Node, logger log.Logger, systemConf *conf.System) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{db: db, rdb: rdb, sdb: sdb, log: l, sf: sf, enforcer: enforcer.(*casbin.SyncedEnforcer), conf: ConfigOptions{system: systemConf}}
	return d, func() {
		l.Info("closing db")
		sql, err := db.DB()
		if err != nil {
			l.Errorf("close db failed %v", err)
			return
		}
		if err := sql.Close(); err != nil {
			log.Error(err)
		}
		l.Info("closing rdb")
		rdb.Close()
		l.Info("closing sdb")
	}, nil
}

type contextTxKey struct{}

// InTx 执行事务
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func (d *Data) DBD(ctx context.Context) *gorm.DB {
	var db *gorm.DB
	if tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB); ok {
		db = tx
	} else {
		db = d.db
	}
	if domainId := d.CtxDomainID(ctx); domainId > 0 {
		db = db.Scopes(DBScopesDomain(domainId))
	}
	return db
}

// NewGormClient 创建数据库客户端
func NewGormClient(cfg *conf.Bootstrap, logger log.Logger, models []interface{}) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "gorm/data/service"))
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

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/service"))
	return bootstrap.NewRedisClient(cfg, l)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

func NewSDB(conf *conf.Data, logger log.Logger) *meilisearch.Client {
	return nil
	log := log.NewHelper(log.With(logger, "module", "data/meilisearch"))
	sdb := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    conf.GetMeilisearch().GetHost(),
		APIKey:  conf.GetMeilisearch().ApiKey,
		Timeout: conf.Meilisearch.Timeout.AsDuration(),
	})
	_, err := sdb.Health()
	if err != nil {
		log.Fatalf("failed opening connection to redis %v", err)
	}
	return sdb
}
