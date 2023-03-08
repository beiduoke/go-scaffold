package data

import (
	"context"
	"strings"

	"github.com/beiduoke/go-scaffold/internal/conf"
	auth "github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthenticator(ac *conf.Auth, logger log.Logger) auth.Authenticator {
	log := log.NewHelper(log.With(logger, "module", "data/authenticatorModel"))
	conf := ac.GetJwt()
	authenticator, err := jwt.NewAuthenticator(
		jwt.WithSecretKey(conf.GetSecretKey()),
		jwt.WithParseContext(parseContextTokenCall(conf.GetHeader(), conf.GetScheme())),
		jwt.WithExpiresAt(conf.GetExpiresTime().AsDuration()),
	)
	if err != nil {
		log.Fatalf("failed authenticator init fail: %v", err)
	}
	return authenticator
}

func parseContextTokenCall(headerField string, scheme string) func(context.Context) (string, error) {
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

// func AuthClaimsFromUser(ctx context.Context) (*SysUser, error) {
// 	claims, success := authn.AuthClaimsFromContext(ctx)
// 	if success {
// 		return nil, errors.New("Auth Claims from fail")
// 	}

// 	return nil, nil
// }
