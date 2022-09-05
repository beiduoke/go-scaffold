package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRDB, NewTransaction, NewModelMigrate, NewSnowflake, NewAuthModel, NewAuthAdapter, NewWatcher, NewAuthEnforcer, NewDomainRepo, NewAuthorityRepo, NewMenuRepo, NewApiRepo, NewUserRepo)

// NewModelMigrate 数据模型迁移
func NewModelMigrate() []interface{} {
	migrates := NewSysModelMigrate()
	// migrates = append(migrates, NewWebModelMigrate()...)
	return migrates
}

// NewTransaction 事务
func NewTransaction(d *Data) biz.Transaction {
	return d
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

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
	sf  *snowflake.Node
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, sf *snowflake.Node, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{db: db, rdb: rdb, log: l, sf: sf}
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
		log.Fatalf("init redis connection failed %v", err)
	}
	return rdb
}
