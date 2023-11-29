package server

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	v1 "github.com/beiduoke/go-scaffold/api/saasdesk/service/v1"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/cmd/server/assets"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/service"
	authn "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	authnM "github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	authzM "github.com/beiduoke/go-scaffold/pkg/middleware/authz"
	"github.com/beiduoke/go-scaffold/pkg/middleware/localize"
	"github.com/beiduoke/go-scaffold/pkg/middleware/multipoint"
	"github.com/beiduoke/go-scaffold/pkg/middleware/signout"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authn.Authenticator, authorized authz.Authorized, creator authn.SecurityUserCreator,
	// service
	authSrv *service.AuthService,
	userSrv *service.UserService,
	roleSrv *service.RoleService,
) *http.Server {

	// Auth 认证鉴权相关中间件
	authMiddleware := NewAuthMiddleware(authenticator, authorized, creator)
	ms := NewHttpMiddleware(logger, authMiddleware)
	srv := bootstrap.CreateHttpServer(cfg, ms...)

	if cfg.GetServer().GetHttp().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos-saas"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}
	openAPIhandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(false)))
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterAuthServiceHTTPServer(srv, authSrv)
	v1.RegisterUserServiceHTTPServer(srv, userSrv)
	v1.RegisterRoleServiceHTTPServer(srv, roleSrv)

	return srv
}

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	whiteList[v1.OperationAuthServiceRegister] = struct{}{}
	whiteList[v1.OperationAuthServiceLoginByPassword] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewAuthMiddleware(authenticator authn.Authenticator, authorized authz.Authorized, creator authn.SecurityUserCreator) middleware.Middleware {
	// jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, jwtV4.RegisteredClaims{})
	return selector.Server(
		// 认证
		authnM.Server(authenticator, creator),
		// 鉴权
		authzM.Server(authorized),
		// 多地登录
		multipoint.Server(),
		// 下线判断
		signout.Server(),
	).
		Match(NewWhiteListMatcher()).Build()
}

// NewHttpMiddleware 创建中间件
func NewHttpMiddleware(logger log.Logger, middle middleware.Middleware) []middleware.Middleware {
	var ms = []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(logger),
		tracing.Server(),
		localize.I18N(),
		validate.Validator(),
	}
	ms = append(ms, middle)
	return ms
}
