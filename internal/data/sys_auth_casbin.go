package data

import (
	"context"

	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAuthFileAdapter(ac *conf.Auth, logger log.Logger) persist.Adapter {
	log := log.NewHelper(log.With(logger, "module", "data/authFileAdapter"))
	a := fileAdapter.NewAdapter(ac.Casbin.PolicyPath)
	log.Infof("casbin file adapters connection")
	return a
}

func NewAuthGormAdapter(d *Data, logger log.Logger) persist.Adapter {
	log := log.NewHelper(log.With(logger, "module", "data/authGormAdapter"))
	a, err := gormadapter.NewAdapterByDBUseTableName(d.DB(context.Background()), "sys", "casbin_rules")
	if err != nil {
		log.Fatalf("failed casbin adapters connection %v", err)
	}
	return a
}

func NewAuthModel(ac *conf.Auth, logger log.Logger) model.Model {
	log := log.NewHelper(log.With(logger, "module", "data/authModel"))
	m, err := model.NewModelFromFile(ac.Casbin.ModelPath)
	if err != nil {
		log.Fatalf("failed casbin model connection %v", err)
	}
	return m
}
