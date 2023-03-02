package data

import (
	"context"
	"strings"

	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/pkg/auth"
	"github.com/beiduoke/go-scaffold/pkg/auth/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthenticator(ac *conf.Auth, logger log.Logger) auth.Authenticator {
	log := log.NewHelper(log.With(logger, "module", "data/authenticatorModel"))
	conf := ac.GetJwt()
	parseContextToken := func(ctx context.Context) (string, error) {
		if header, ok := transport.FromServerContext(ctx); ok {
			authorize := header.RequestHeader().Get(conf.GetHeader())
			if authorize == "" {
				return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+conf.GetScheme())
			}
			splits := strings.SplitN(authorize, " ", 2)
			if len(splits) < 2 {
				return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
			}

			if !strings.EqualFold(splits[0], conf.GetScheme()) {
				return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+conf.GetScheme())
			}
			return splits[1], nil
		}
		return "", nil
	}

	authenticator, err := jwt.NewAuthenticator(jwt.WithSecretKey(conf.GetSecretKey()), jwt.WithParseContext(parseContextToken), jwt.WithExpiresAt(conf.GetExpiresTime().AsDuration()))
	if err != nil {
		log.Fatalf("failed authenticator init fail: %v", err)
	}
	return authenticator
}
