package data

import (
	"github.com/beiduoke/go-scaffold/internal/conf"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func NewAuthModel(ac *conf.Auth, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "data/authModel"))
	m, err := model.NewModelFromFile(ac.Casbin.ModelPath)
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}

func NewAuthAdapter(db *gorm.DB, ac *conf.Auth, logger log.Logger) (adapter persist.Adapter) {
	log := log.NewHelper(log.With(logger, "module", "data/authAdapter"))
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

func NewWatcher(conf *conf.Data, logger log.Logger) persist.Watcher {
	log := log.NewHelper(log.With(logger, "module", "data/authWatcher"))
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

func NewAuthEnforcer(model model.Model, adapter persist.Adapter, watcher persist.Watcher, logger log.Logger) stdcasbin.IEnforcer {
	log := log.NewHelper(log.With(logger, "module", "data/authEnforcer"))
	// enforcer, err := stdcasbin.NewEnforcer(model, adapter)
	// enforcer, err := stdcasbin.NewCachedEnforcer(model, adapter)
	enforcer, err := stdcasbin.NewSyncedEnforcer(model, adapter)
	if err != nil {
		log.Fatalf("failed casbin enforcer %v", err)
	}
	err = enforcer.SetWatcher(watcher)
	if err != nil {
		log.Fatalf("failed casbin watcher %v", err)
	}

	return enforcer
}
