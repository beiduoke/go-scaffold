package server

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/service/v1"
	"github.com/beiduoke/go-scaffold/api/common/conf"
	authnM "github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/authn"
	authzM "github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/authz"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/localize"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/multipoint"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/middleware/signout"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/service"
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
	NewHttpMiddleware,
	// 认证中间件
	NewAuthMiddleware,
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	whiteList[v1.OperationAuthServiceRegister] = struct{}{}
	whiteList[v1.OperationAuthServiceLoginByPassword] = struct{}{}
	whiteList[v1.OperationAuthServiceLoginBySms] = struct{}{}
	whiteList[v1.OperationAuthServiceLoginByEmail] = struct{}{}
	whiteList[v1.OperationDomainServiceGetDomainCode] = struct{}{}
	whiteList[v1.OperationDomainServiceGetDomainName] = struct{}{}
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

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	// service
	authSrv *service.AuthService,
	userSrv *service.UserService,
	roleSrv *service.RoleService,
	domainSrv *service.DomainService,
	deptSrv *service.DeptService,
	menuSrv *service.MenuService,
	postSrv *service.PostService,
	dictSrv *service.DictService,
	ms []middleware.Middleware) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, ms...)

	openAPIhandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(false)))
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterAuthServiceHTTPServer(srv, authSrv)
	v1.RegisterUserServiceHTTPServer(srv, userSrv)
	v1.RegisterRoleServiceHTTPServer(srv, roleSrv)
	v1.RegisterDomainServiceHTTPServer(srv, domainSrv)
	v1.RegisterDeptServiceHTTPServer(srv, deptSrv)
	v1.RegisterMenuServiceHTTPServer(srv, menuSrv)
	v1.RegisterPostServiceHTTPServer(srv, postSrv)
	v1.RegisterDictServiceHTTPServer(srv, dictSrv)

	return srv
}
