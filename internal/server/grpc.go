package server

import (
	adminv1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	webv1 "github.com/beiduoke/go-scaffold/api/web/v1"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/service/admin"
	"github.com/beiduoke/go-scaffold/internal/service/web"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, as *admin.AdminService, bs *web.WebService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	adminv1.RegisterAdminServer(srv, as)
	webv1.RegisterWebServer(srv, bs)
	return srv
}
