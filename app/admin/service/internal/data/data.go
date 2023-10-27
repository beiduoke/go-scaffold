package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	// 数据库客户端
	NewGormClient,
	// redis客户端
	NewRedisClient,
	// 搜索引擎客户端
	NewMeilisearchClient,
	// 雪花ID生成器
	NewSnowflake,
	// 数据迁移
	NewModelMigrate,
	// 数据集成器··
	NewData,
	// 事物
	NewTransaction,
	// 认证解析器
	NewSecurityUser,
	// 认证器
	NewAuthenticator,
	// casbin鉴权客户端
	NewAuthzCasbinClient,
	//  鉴权器
	NewAuthorized,
	// 数据操作
	NewDomainRepo,
	NewRoleRepo,
	NewMenuRepo,
	NewUserRepo,
	NewDeptRepo,
	NewPostRepo,
	NewDictRepo,
	NewAuthRepo,
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
	base *conf.Base
}

// Data .
type Data struct {
	cfg      *conf.Bootstrap
	log      *log.Helper
	db       *gorm.DB
	rdb      *redis.Client
	sdb      *meilisearch.Client
	sf       *snowflake.Node
	enforcer *casbin.SyncedEnforcer
}

// NewData .
func NewData(cfg *conf.Bootstrap, db *gorm.DB, rdb *redis.Client, sdb *meilisearch.Client, enforcer *casbin.SyncedEnforcer, sf *snowflake.Node, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/initialize"))
	d := &Data{db: db, rdb: rdb, sdb: sdb, log: l, sf: sf, enforcer: enforcer, cfg: cfg}
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
	return bootstrap.NewGormClient(cfg, l, models...)
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/service"))
	return bootstrap.NewRedisClient(cfg, l)
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

// NewAuthenticator 创建认证
func NewAuthenticator(cfg *conf.Bootstrap, logger log.Logger) authn.Authenticator {
	return bootstrap.NewJwtAuthenticator(cfg, logger)
}

// NewAuthzCasbinClient 创建Casbin客户端
func NewAuthzCasbinClient(cfg *conf.Bootstrap, logger log.Logger) *casbin.SyncedEnforcer {
	log.NewHelper(log.With(logger, "module", "casbin/authz/service"))
	model, adapter, watcher := bootstrap.NewAuthzCasbinModel(cfg, logger), bootstrap.NewAuthzCasbinGormAdapter(cfg, logger), bootstrap.NewAuthzCasbinWatcher(cfg, logger)
	return bootstrap.NewAuthzCasbinEnforcer(model, adapter, watcher, logger)
}

// NewAuthorized 创建鉴权
func NewAuthorized(enforcer *casbin.SyncedEnforcer, logger log.Logger) authz.Authorized {
	return bootstrap.NewAuthzCasbin(enforcer, logger)
}
