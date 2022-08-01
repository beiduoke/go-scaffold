//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/bedoke/go-scaffold/internal/data"
	"github.com/bedoke/go-scaffold/internal/server"
	"github.com/bedoke/go-scaffold/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Auth, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
