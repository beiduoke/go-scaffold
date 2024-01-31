package bootstrap

import (
	"context"
	"strings"
	"time"

	"github.com/beiduoke/go-scaffold/api/common/conf"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	jwtHeaderField  = "Authorization"
	jwtHeaderScheme = "Bearer"
	jwtExpire       = time.Hour * 24 * 7
)

// NewJwtAuthenticator Jwt创建认证器
func NewJwtAuthenticator(cfg *conf.Bootstrap, logger log.Logger) authn.Authenticator {
	log := log.NewHelper(log.With(logger, "module", "jwt/data/authenticator"))
	conf := cfg.Server.Http.Middleware.Auth
	if conf.GetHeader() == "" {
		conf.Header = jwtHeaderField
	}
	if conf.GetScheme() == "" {
		conf.Scheme = jwtHeaderScheme
	}
	if conf.GetExpiresTime() == nil {
		conf.ExpiresTime = durationpb.New(jwtExpire)
	}
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
