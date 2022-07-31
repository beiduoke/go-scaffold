package data

import (
	"context"

	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	redisadapter "github.com/mlsen/casbin-redis-adapter/v2"
)

func NewAuthModel(ac *conf.Auth, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "data/authModel"))
	m, err := model.NewModelFromFile(ac.Casbin.ModelPath)
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}

func NewAuthAdapter(d *Data, ac *conf.Auth, logger log.Logger) persist.Adapter {
	log := log.NewHelper(log.With(logger, "module", "data/authAdapter"))
	gormAdapter, err := gormadapter.NewAdapterByDBUseTableName(d.DB(context.Background()), "sys", "casbin_rules")
	if err != nil {
		log.Fatalf("failed gorm casbin adapters connection %v", err)
	}
	// 优先使用gorm进行存储
	return gormAdapter
	// redis 适配器
	fileAdapter := fileAdapter.NewAdapter(ac.Casbin.PolicyPath)
	// redis 适配器
	redisAdapter := redisadapter.NewFromClient(d.rdb)
	return &AuthAdapterRepo{
		redisAdapter: redisAdapter,
		gormAdapter:  gormAdapter,
		fileAdapter:  fileAdapter,
	}
}

// 代理模式进行多个适配器相同操作
type AuthAdapterRepo struct {
	redisAdapter persist.Adapter
	gormAdapter  persist.Adapter
	fileAdapter  persist.Adapter
}

// LoadPolicy loads all policy rules from the storage.
func (auth *AuthAdapterRepo) LoadPolicy(model model.Model) error {
	return auth.gormAdapter.LoadPolicy(model)
}

// SavePolicy saves all policy rules to the storage.
func (auth *AuthAdapterRepo) SavePolicy(model model.Model) error {
	return auth.gormAdapter.SavePolicy(model)
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (auth *AuthAdapterRepo) AddPolicy(sec string, ptype string, rule []string) error {
	return auth.gormAdapter.AddPolicy(sec, ptype, rule)
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (auth *AuthAdapterRepo) RemovePolicy(sec string, ptype string, rule []string) error {
	return auth.gormAdapter.RemovePolicy(sec, ptype, rule)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (auth *AuthAdapterRepo) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return auth.gormAdapter.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...)
}
