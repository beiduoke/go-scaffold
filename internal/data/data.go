package data

import (
	"context"

	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRDB, NewTransaction, NewModelMigrate, NewAuthModel, NewAuthAdapter, NewUserRepo, NewAuthorityRepo)

func NewModelMigrate() []interface{} {
	migrates := NewSysModelMigrate()
	// migrates = append(migrates, NewWebModelMigrate()...)
	return migrates
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
}

type contextTxKey struct{}

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

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{
		db:  db,
		rdb: rdb,
		log: l,
	}
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
