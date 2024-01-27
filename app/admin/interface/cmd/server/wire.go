//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/admin/interface/internal/data"
	"github.com/beiduoke/go-scaffold/app/admin/interface/internal/server"
	"github.com/beiduoke/go-scaffold/app/admin/interface/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
