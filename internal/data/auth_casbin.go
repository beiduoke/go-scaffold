package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	authzCasbin "github.com/beiduoke/go-scaffold/pkg/auth/authz/casbin"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/constant"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
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

// NewAuthModel 模型
func NewAuthModel(ac *conf.Auth, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinModel"))
	m, err := model.NewModelFromString(modelText)
	if ac.Casbin.GetModelPath() != "" {
		m, err = model.NewModelFromFile(ac.Casbin.GetModelPath())
	}
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}

// NewAuthAdapter 适配器
func NewAuthAdapter(db *gorm.DB, ac *conf.Auth, logger log.Logger) (adapter persist.Adapter) {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinAdapter"))
	// gormadapter "github.com/casbin/gorm-adapter/v3"
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "sys", "casbin_rules")
	log.Info("initialization gorm adapter ")
	if err != nil {
		log.Fatalf("failed gorm casbin adapters connection %v", err)
	}
	// 优先使用gorm进行存储
	// file 适配器
	// fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	// adapter = fileadapter.NewAdapter(ac.Casbin.PolicyPath)
	// log.Info("initialization file adapter ")
	return adapter
}

// NewWatcher 监视器
func NewWatcher(conf *conf.Data, logger log.Logger) persist.Watcher {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinWatcher"))
	// rediswatcher "github.com/casbin/redis-watcher/v2"
	w, err := rediswatcher.NewWatcher(conf.GetRedis().GetAddr(), rediswatcher.WatcherOptions{
		Options: redis.Options{
			Network:  conf.GetRedis().GetNetwork(),
			Password: conf.GetRedis().GetPassword(),
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

// NewAuthEnforcer 执行器
func NewAuthEnforcer(model model.Model, adapter persist.Adapter, watcher persist.Watcher, logger log.Logger) stdcasbin.IEnforcer {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinEnforcer"))
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

// NewAuthCasbin casbin认证
func NewAuthCasbin(logger log.Logger, enforcer stdcasbin.IEnforcer) authz.Authorized {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbin"))
	authorized, err := authzCasbin.NewAuthorized(context.Background(), authzCasbin.WithEnforcer(enforcer))
	if err != nil {
		log.Fatalf("failed casbin authorized %v", err)
	}
	return authorized
}
