package data

import (
	"context"
	"errors"

	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/auth"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func NewAuthModel(ac *conf.Auth, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinModel"))
	m, err := model.NewModelFromFile(ac.Casbin.ModelPath)
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}

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

func NewAuthEnforcer(model model.Model, adapter persist.Adapter, watcher persist.Watcher, logger log.Logger) stdcasbin.IEnforcer {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbinEnforcer"))
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

type authCasbin struct {
	logger    *log.Helper
	model     model.Model
	adapter   persist.Adapter
	watcher   persist.Watcher
	iEnforcer stdcasbin.IEnforcer
}

func NewAuthCasbin(logger log.Logger, model model.Model, adapter persist.Adapter, watcher persist.Watcher) *authCasbin {
	log := log.NewHelper(log.With(logger, "module", "data/authCasbin"))
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
	return &authCasbin{
		logger:    log,
		model:     model,
		adapter:   adapter,
		watcher:   watcher,
		iEnforcer: enforcer,
	}
}

var _ auth.SecurityUser = (*securityUser)(nil)

type securityUser struct {
	options auth.Options
	// 用户
	user string
	// 域/租户
	domain string
	// 角色
	subject string
	// 资源
	object string
	// 方法
	action string
}

type Option func(*auth.Options)

func NewSecurityUserCreator() auth.SecurityUserCreator {
	return func() auth.SecurityUser {
		return &securityUser{}
	}
}

func NewSecurityUser() auth.SecurityUser {
	return &securityUser{}
}

func ParseFromContext(ctx context.Context) auth.SecurityUser {
	newSecurityUser := NewSecurityUser()
	if newSecurityUser.ParseFromContext(ctx) != nil {
		return &securityUser{}
	}
	return newSecurityUser
}

const (
	HeaderDomainCodeKey = "X-Domain-Code"
)

// ParseFromContext parses the user from the context.
func (su *securityUser) ParseFromContext(ctx context.Context) error {
	if header, ok := transport.FromServerContext(ctx); ok {
		su.object = header.Operation()
		su.action = "*"
		if domainCode := header.RequestHeader().Get(HeaderDomainCodeKey); domainCode != "" {
			su.domain = domainCode
		}
		// if header.Kind() == transport.KindHTTP {
		// 	if ht, ok := header.(http.Transporter); ok {
		// 		su.Object = ht.Request().URL.Object
		// 		su.Action = ht.Request().Action
		// 	}
		// }
	} else {
		return errors.New("jwt claim missing")
	}
	return nil
}

// GetSubject returns the subject of the token.
func (su *securityUser) GetSubject() string {
	return su.subject
}

// GetObject returns the object of the token.
func (su *securityUser) GetObject() string {
	return su.object
}

// GetAction returns the action of the token.
func (su *securityUser) GetAction() string {
	return su.action
}

// GetDomain returns the domain of the token.
func (su *securityUser) GetDomain() string {
	return su.domain
}

// GetID returns the user of the token.
func (su *securityUser) GetUser() string {
	return su.user
}
