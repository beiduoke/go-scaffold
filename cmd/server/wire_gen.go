// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/bedoke/go-scaffold/internal/data"
	"github.com/bedoke/go-scaffold/internal/server"
	"github.com/bedoke/go-scaffold/internal/service"
	"github.com/bedoke/go-scaffold/internal/service/admin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, auth *conf.Auth, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	v := data.NewModelMigrate()
	db := data.NewDB(confData, logger, v)
	client := data.NewRDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	userUsecase := biz.NewUserUsecase(userRepo, transaction, logger)
	adminService := admin.NewAdminService(logger, userUsecase)
	grpcServer := server.NewGRPCServer(confServer, auth, adminService, logger)
	model := data.NewAuthModel(auth, logger)
	adapter := data.NewAuthGormAdapter(dataData, logger)
	middleware := server.NewAuthMiddleware(auth, model, adapter)
	serverOption := server.NewMiddleware(logger, middleware)
	httpServer := server.NewHTTPServer(confServer, adminService, serverOption)
	websocketService := service.NewWebsocketService(logger)
	websocketServer := server.NewWebsocketServer(confServer, logger, websocketService)
	app := newApp(logger, grpcServer, httpServer, websocketServer)
	return app, func() {
		cleanup()
	}, nil
}
