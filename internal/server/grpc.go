package server

import (
	"github.com/beiduoke/go-scaffold/api/common/conf"
	serverv1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/service/api"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, as *api.ApiService, logger log.Logger) *grpc.Server {
	var ms = []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		validate.Validator(),
	}
	srv := bootstrap.CreateGrpcServer(cfg, ms...)
	serverv1.RegisterApiServer(srv, as)
	return srv
}
