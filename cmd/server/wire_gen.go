// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/data"
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/beiduoke/go-scaffold/internal/server"
	"github.com/beiduoke/go-scaffold/internal/service/api"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, auth *conf.Auth, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	websocketService := websocket.NewWebsocketService(logger)
	v := data.NewModelMigrate()
	db := data.NewDB(confData, logger, v)
	client := data.NewRDB(confData, logger)
	meilisearchClient := data.NewSDB(confData, logger)
	model := data.NewAuthModel(auth, logger)
	adapter := data.NewAuthAdapter(db, auth, logger)
	watcher := data.NewWatcher(confData, logger)
	iEnforcer := data.NewAuthEnforcer(model, adapter, watcher, logger)
	node := data.NewSnowflake(logger)
	dataData, cleanup, err := data.NewData(db, client, meilisearchClient, iEnforcer, node, logger)
	if err != nil {
		return nil, nil, err
	}
	transaction := data.NewTransaction(dataData)
	menuRepo := data.NewMenuRepo(dataData, logger)
	domainRepo := data.NewDomainRepo(logger, dataData, menuRepo)
	deptRepo := data.NewDeptRepo(logger, dataData)
	roleRepo := data.NewRoleRepo(logger, dataData, menuRepo, deptRepo)
	authenticator := data.NewAuthenticator(auth, logger)
	postRepo := data.NewPostRepo(logger, dataData)
	userRepo := data.NewUserRepo(logger, dataData, auth, authenticator, domainRepo, roleRepo, postRepo, menuRepo, deptRepo)
	bizBiz := biz.NewBiz(logger, transaction, iEnforcer, domainRepo, roleRepo, userRepo)
	authUsecase := biz.NewAuthUsecase(logger, bizBiz, authenticator)
	userUsecase := biz.NewUserUsecase(logger, bizBiz, auth)
	domainUsecase := biz.NewDomainUsecase(logger, bizBiz)
	roleUsecase := biz.NewRoleUsecase(logger, bizBiz)
	menuUsecase := biz.NewMenuUsecase(logger, bizBiz, menuRepo)
	deptUsecase := biz.NewDeptUsecase(logger, bizBiz, deptRepo)
	postUsecase := biz.NewPostUsecase(logger, bizBiz, postRepo)
	apiService := api.NewApiService(logger, auth, websocketService, authUsecase, userUsecase, domainUsecase, roleUsecase, menuUsecase, deptUsecase, postUsecase)
	grpcServer := server.NewGRPCServer(confServer, auth, apiService, logger)
	authorized := data.NewAuthCasbin(logger, iEnforcer)
	securityUserCreator := data.NewSecurityUser(logger, dataData, userRepo)
	middleware := server.NewAuthMiddleware(auth, authenticator, authorized, securityUserCreator)
	serverOption := server.NewMiddleware(logger, middleware)
	httpServer := server.NewHTTPServer(confServer, apiService, serverOption)
	websocketServer := server.NewWebsocketServer(confServer, logger, websocketService)
	app := newApp(logger, grpcServer, httpServer, websocketServer)
	return app, func() {
		cleanup()
	}, nil
}
