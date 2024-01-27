package bootstrap

import (
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	midRateLimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/gorilla/handlers"

	"github.com/beiduoke/go-scaffold-single/api/common/conf"
)

// CreateHttpServer 创建HTTP服务端
func CreateHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosHttp.Server {
	var opts = []kratosHttp.ServerOption{
		kratosHttp.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Cors.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Cors.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Cors.Origins),
		)),
	}

	var ms []middleware.Middleware
	if cfg.Server != nil && cfg.Server.Http != nil && cfg.Server.Http.Middleware != nil {
		if cfg.Server.Http.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Http.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Http.Middleware.GetEnableValidate() {
			ms = append(ms, validate.Validator())
		}
		if cfg.Server.Http.Middleware.GetEnableCircuitBreaker() {
		}
		if cfg.Server.Http.Middleware.Limiter != nil {
			var limiter ratelimit.Limiter
			switch cfg.Server.Http.Middleware.Limiter.GetName() {
			case "bbr":
				limiter = bbr.NewLimiter()
			}
			ms = append(ms, midRateLimit.Server(midRateLimit.WithLimiter(limiter)))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, kratosHttp.Middleware(ms...))

	if cfg.Server.Http.Network != "" {
		opts = append(opts, kratosHttp.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, kratosHttp.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, kratosHttp.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}

	return kratosHttp.NewServer(opts...)
}
