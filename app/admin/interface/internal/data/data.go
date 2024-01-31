package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz/noop"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/beiduoke/go-scaffold/pkg/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	// 服务注册发现
	NewDiscovery,
	NewData,

	NewRedisClient,

	// 认证
	NewAuthenticator,
	NewAuthTokenRepo,
	NewSecurityUser,
	NewAuthorized,

	NewUserServiceClient,
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client
}

// NewData .
func NewData(rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		rdb: rdb,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return bootstrap.NewRedisClient(cfg.Data)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) coreV1.UserServiceClient {
	return coreV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewAuthorizerServiceClient(r registry.Discovery, c *conf.Bootstrap) coreV1.UserServiceClient {
	return coreV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

// NewAuthenticator 创建认证
func NewAuthenticator(cfg *conf.Bootstrap, logger log.Logger) authn.Authenticator {
	return bootstrap.NewJwtAuthenticator(cfg, logger)
}

// NewAuthorized 创建鉴权
func NewAuthorized(logger log.Logger) authz.Authorized {
	return noop.State{}
}
