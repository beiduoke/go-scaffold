package bootstrap

import (
	"context"
	"strings"

	"github.com/beiduoke/go-scaffold-single/api/common/conf"
	"github.com/beiduoke/go-scaffold-single/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold-single/pkg/auth/authn/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewJwtAuthenticator Jwt创建认证器
func NewJwtAuthenticator(cfg *conf.Bootstrap, logger log.Logger) authn.Authenticator {
	log := log.NewHelper(log.With(logger, "module", "jwt/data/authenticator"))
	conf := cfg.Server.Http.Middleware.Auth
	authenticator, err := jwt.NewAuthenticator(
		jwt.WithSecretKey(conf.GetKey()),
		jwt.WithParseContext(jwtParseContextTokenCall(conf.GetHeader(), conf.GetScheme())),
		jwt.WithExpiresAt(conf.GetExpiresTime().AsDuration()),
	)
	if err != nil {
		log.Fatalf("failed authenticator init fail: %v", err)
	}
	return authenticator
}

func jwtParseContextTokenCall(headerField string, scheme string) func(context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		if header, ok := transport.FromServerContext(ctx); ok {
			authorize := header.RequestHeader().Get(headerField)
			if authorize == "" {
				return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+scheme)
			}
			splits := strings.SplitN(authorize, " ", 2)
			if len(splits) < 2 {
				return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
			}

			if !strings.EqualFold(splits[0], scheme) {
				return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+scheme)
			}
			return splits[1], nil
		}
		return "", nil
	}
}
