package server

import (
	"context"

	serverv1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/common/conf"
	authnM "github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/authn"
	authzM "github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/authz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/localize"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/multipoint"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/middleware/signout"
	"github.com/beiduoke/go-scaffold/app/admin/internal/service"
	authn "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/beiduoke/go-scaffold/pkg/bootstrap"

	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"

	// fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"

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
)

// ProviderSet is server providers.
var ProviderHttp = wire.NewSet(
	NewMiddleware,
	// 认证中间件
	NewAuthMiddleware,
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.server.v1.Api/Register"] = struct{}{}
	whiteList["/api.server.v1.Api/Login"] = struct{}{}
	whiteList["/api.server.v1.Api/SmsLogin"] = struct{}{}
	whiteList["/api.server.v1.Api/EmailLogin"] = struct{}{}
	whiteList["/api.server.v1.Api/GetDomainCode"] = struct{}{}
	whiteList["/api.server.v1.Api/GetDomainName"] = struct{}{}
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

// NewMiddleware 创建中间件
func NewMiddleware(logger log.Logger, middle middleware.Middleware) []middleware.Middleware {
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

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	as *service.AdminService, ms []middleware.Middleware) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, ms...)

	openAPIhandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(false)))
	srv.HandlePrefix("/q/", openAPIhandler)

	serverv1.RegisterAdminServiceHTTPServer(srv, as)
	return srv
}
