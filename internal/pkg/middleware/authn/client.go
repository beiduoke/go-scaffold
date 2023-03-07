package authn

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

// Client is a client authenticator middleware.
func Client(authenticator authn.Authenticator, opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	token, err := authenticator.CreateIdentity(context.Background(), o.claims)
	if err != nil {
		log.Errorf("authenticator middleware create token failed: %s", err.Error())
	}

	if o.contextWithTokenFunc == nil {
		log.Errorf("authenticator middleware context with token failed: %s")
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(o.contextWithTokenFunc(ctx, token), req)
		}
	}
}
