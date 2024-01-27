// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/data"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/server"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, registrar registry.Registrar, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	client := data.NewEntClient(bootstrap, logger)
	redisClient := data.NewRedisClient(bootstrap, logger)
	meilisearchClient := data.NewMeilisearchClient(bootstrap, logger)
	syncedEnforcer := data.NewAuthzCasbinClient(bootstrap, logger)
	node := data.NewSnowflake(logger)
	dataData, cleanup, err := data.NewData(logger, bootstrap, client, redisClient, meilisearchClient, syncedEnforcer, node)
	if err != nil {
		return nil, nil, err
	}
	authRepo := data.NewAuthRepo(dataData, logger)
	authService := service.NewAuthService(logger, authRepo)
	userRepo := data.NewUserRepo(dataData, logger)
	userService := service.NewUserService(logger, userRepo)
	roleRepo := data.NewRoleRepo(dataData, logger)
	roleService := service.NewRoleService(logger, roleRepo)
	grpcServer := server.NewGRPCServer(bootstrap, logger, authService, userService, roleService)
	authenticator := data.NewAuthenticator(bootstrap, logger)
	authorized := data.NewAuthorized(syncedEnforcer, logger)
	securityUserCreator := data.NewSecurityUser(logger, dataData)
	httpServer := server.NewHTTPServer(bootstrap, logger, authenticator, authorized, securityUserCreator, authService, userService, roleService)
	app := newApp(logger, registrar, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}