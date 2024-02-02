package server

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/app/admin/interface/internal/service"
	"github.com/beiduoke/go-scaffold/app/core/service/cmd/server/assets"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"

	authnM "github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	authzM "github.com/beiduoke/go-scaffold/pkg/middleware/authz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
)

// NewWhiteListMatcher 创建jwt白名单
func newHttpWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[v1.OperationAuthServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newHttpMiddleware(authenticator authn.Authenticator, authorized authz.Authorized, creator authn.SecurityUserCreator, logger log.Logger) []middleware.Middleware {
	var ms = []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(logger),
		tracing.Server(),
		validate.Validator(),
	}
	ms = append(ms, selector.Server(
		// 认证
		authnM.Server(authenticator, creator),
		// 鉴权
		authzM.Server(authorized),
	).Match(newHttpWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authn.Authenticator, authorized authz.Authorized, creator authn.SecurityUserCreator,
	authnSvc *service.AuthService,
	userSvc *service.UserService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, newHttpMiddleware(authenticator, authorized, creator, logger)...)

	v1.RegisterAuthServiceHTTPServer(srv, authnSvc)
	v1.RegisterUserServiceHTTPServer(srv, userSvc)

	if cfg.GetServer().GetHttp().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Front Service"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
