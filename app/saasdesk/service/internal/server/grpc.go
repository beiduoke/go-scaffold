package server

import (
	"github.com/beiduoke/go-scaffold/api/common/conf"
	v1 "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/service"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/go-kratos/kratos/v2/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	cfg *conf.Bootstrap, logger log.Logger,
	// service
	authSrv *service.AuthService,
	userSrv *service.UserService,
	roleSrv *service.RoleService,
) *grpc.Server {
	var ms = []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		validate.Validator(),
	}
	srv := bootstrap.CreateGrpcServer(cfg, ms...)
	v1.RegisterAuthServiceServer(srv, authSrv)
	v1.RegisterUserServiceServer(srv, userSrv)
	v1.RegisterRoleServiceServer(srv, roleSrv)
	return srv
}
