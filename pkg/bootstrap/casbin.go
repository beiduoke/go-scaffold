package bootstrap

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	authzCasbin "github.com/beiduoke/go-scaffold/pkg/auth/authz/casbin"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/constant"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	entrapper "github.com/casbin/ent-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

const modelText = `
[request_definition]
r = sub, obj, act, dom

[policy_definition]
p = sub, obj, act, dom

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act || (r.sub == "1" && r.dom == "1")
`

// NewAuthzCasbinModel 模型
func NewAuthzCasbinModel(cfg *conf.Bootstrap, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbinModel"))
	authz := cfg.Server.Http.Middleware.Authorizer.GetCasbin()
	m, err := model.NewModelFromString(modelText)
	if authz.GetModelPath() != "" {
		m, err = model.NewModelFromFile(authz.GetModelPath())
	}
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}

// NewAuthzCasbinGormAdapter 适配器
func NewAuthzCasbinGormAdapter(cfg *conf.Bootstrap, logger log.Logger) (adapter persist.Adapter) {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbinGormAdapter"))
	// gormadapter "github.com/casbin/gorm-adapter/v3"
	adapter, err := gormadapter.NewAdapterByDBUseTableName(NewGormClient(cfg, log), "sys", "casbin_rules")
	log.Info("initialization gorm adapter ")
	if err != nil {
		log.Fatalf("failed gorm casbin adapters connection %v", err)
	}
	return adapter
}

// NewAuthzCasbinEntAdapter 适配器
func NewAuthzCasbinEntAdapter(cfg *conf.Bootstrap, logger log.Logger) (adapter persist.Adapter) {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbinEntAdapter"))
	adapter, err := entrapper.NewAdapter(cfg.Data.Database.Driver, cfg.Data.Database.Source)
	log.Info("initialization ent adapter ")
	if err != nil {
		log.Fatalf("failed gorm casbin adapters connection %v", err)
	}
	return adapter
}

// NewAuthzCasbinWatcher 监视器
func NewAuthzCasbinWatcher(cfg *conf.Bootstrap, logger log.Logger) persist.Watcher {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbinWatcher"))
	// rediswatcher "github.com/casbin/redis-watcher/v2"
	w, err := rediswatcher.NewWatcher(cfg.Data.GetRedis().GetAddr(), rediswatcher.WatcherOptions{
		Options: redis.Options{
			Network:  cfg.Data.GetRedis().GetNetwork(),
			Password: cfg.Data.GetRedis().GetPassword(),
		},
		Channel: "/casbin",
		// Only exists in test, generally be true
		IgnoreSelf: false,
	})
	if err != nil {
		log.Fatalf("failed casbin redis watch %v", err)
	}

	_ = w.SetUpdateCallback(func(s string) {
		log.Infof("casbin redis watcher info %v", s)
	})

	return w
}

// NewAuthzCasbinEnforcer 执行器
func NewAuthzCasbinEnforcer(model model.Model, adapter persist.Adapter, watcher persist.Watcher, logger log.Logger) *stdcasbin.SyncedEnforcer {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbinEnforcer"))
	// enforcer, err := stdcasbin.NewEnforcer(model, adapter)
	// enforcer, err := stdcasbin.NewCachedEnforcer(model, adapter)
	enforcer, err := stdcasbin.NewSyncedEnforcer(model, adapter)
	if err != nil {
		log.Fatalf("failed casbin enforcer %v", err)
	}
	enforcer.SetFieldIndex("p", constant.DomainIndex, 3)
	err = enforcer.SetWatcher(watcher)
	if err != nil {
		log.Fatalf("failed casbin watcher %v", err)
	}
	return enforcer
}

// NewAuthzCasbin casbin认证
func NewAuthzCasbin(enforcer *stdcasbin.SyncedEnforcer, logger log.Logger) authz.Authorized {
	log := log.NewHelper(log.With(logger, "module", "casbin/data/authCasbin"))
	authorized, err := authzCasbin.NewAuthorized(context.Background(), authzCasbin.WithEnforcer(enforcer))
	if err != nil {
		log.Fatalf("failed casbin authorized %v", err)
	}
	return authorized
}
