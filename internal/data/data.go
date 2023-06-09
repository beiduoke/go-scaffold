package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
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
	NewDB,
	NewRDB,
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
	enforcer casbin.IEnforcer
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, sdb *meilisearch.Client, enforcer casbin.IEnforcer, sf *snowflake.Node, logger log.Logger, systemConf *conf.System) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{db: db, rdb: rdb, sdb: sdb, log: l, sf: sf, enforcer: enforcer, conf: ConfigOptions{system: systemConf}}
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

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger, migrates []interface{}) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger:         gormLogger.Default.LogMode(gormLogger.LogLevel(conf.Database.LogLevel)),
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "scaffold_", // table name prefix, table for `User` would be `t_users`
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(migrates...); err != nil {
		log.Fatal(err)
	}
	return db
}

func NewRDB(conf *conf.Data, logger log.Logger) *redis.Client {
	log := log.NewHelper(log.With(logger, "module", "data/redis"))
	rdb := redis.NewClient(&redis.Options{
		Network:      conf.GetRedis().GetNetwork(),
		Addr:         conf.GetRedis().GetAddr(),
		Password:     conf.GetRedis().GetPassword(),
		ReadTimeout:  conf.GetRedis().ReadTimeout.AsDuration(),
		WriteTimeout: conf.GetRedis().WriteTimeout.AsDuration(),
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("failed opening connection to redis %v", err)
	}
	return rdb
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
