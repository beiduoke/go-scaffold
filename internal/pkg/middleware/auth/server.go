package auth

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth"
	"github.com/go-kratos/kratos/v2/middleware"
)

// Server is a server authenticator middleware.
func Server(authenticator auth.Authenticator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claims, err := authenticator.Authenticate(ctx)
			if err != nil {
				return nil, err
			}

			ctx = auth.ContextWithAuthClaims(ctx, claims)

			return handler(ctx, req)
		}
	}
}
