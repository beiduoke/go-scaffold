package server

import (
	"context"
	"fmt"

	adminv1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	webv1 "github.com/beiduoke/go-scaffold/api/web/v1"
	"github.com/beiduoke/go-scaffold/internal/conf"
	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/beiduoke/go-scaffold/internal/pkg/middleware/localize"
	"github.com/beiduoke/go-scaffold/internal/service/admin"
	"github.com/beiduoke/go-scaffold/internal/service/web"
	casbinM "github.com/beiduoke/go-scaffold/pkg/authz/casbin"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	// gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"

	// fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"

	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

// ProviderSet is server providers.
var ProviderHttp = wire.NewSet(NewMiddleware, NewAuthMiddleware)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.admin.v1.Admin/Login"] = struct{}{}
	whiteList["/api.admin.v1.Admin/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewAuthMiddleware(ac *conf.Auth, m model.Model, policy persist.Adapter, enforcer stdcasbin.IEnforcer) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			func(token *jwtV4.Token) (interface{}, error) {
				fmt.Println(token)
				return []byte(ac.ApiKey), nil
			},
			jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
			jwt.WithClaims(func() jwtV4.Claims {
				return nil
			}),
		),
		casbinM.Server(
			casbinM.WithDomainSupport(),
			// enforcer 指定则无需传 model 以及 policy
			// casbinM.WithCasbinModel(m),
			// casbinM.WithCasbinPolicy(policy),
			casbinM.WithCasbinEnforcer(enforcer),
			casbinM.WithSecurityUserCreator(myAuthz.NewSecurityUser),
		),
	).
		Match(NewWhiteListMatcher()).Build()
}

// NewMiddleware 创建中间件
func NewMiddleware(logger log.Logger, auth middleware.Middleware) http.ServerOption {
	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		validate.Validator(),
		localize.I18N(),
		auth,
	)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, as *admin.AdminService, bs *web.WebService, middleware http.ServerOption) *http.Server {
	var opts = []http.ServerOption{
		middleware,
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	openAPIhandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(true)))
	srv.HandlePrefix("/q/", openAPIhandler)

	adminv1.RegisterAdminHTTPServer(srv, as)
	webv1.RegisterWebHTTPServer(srv, bs)
	return srv
}
