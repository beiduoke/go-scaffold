package authz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"

	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
)

const (
	reason string = "FORBIDDEN"
)

var (
	ErrUnauthorized  = errors.Forbidden(reason, "unauthorized access")
	ErrMissingClaims = errors.Forbidden(reason, "missing authz claims")
	ErrInvalidClaims = errors.Forbidden(reason, "invalid authz claims")
)

func Server(authorizer authz.Authorizer, opts ...Option) middleware.Middleware {
	o := &options{}

	for _, opt := range opts {
		opt(o)
	}

	if authorizer == nil {
		return nil
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				allowed bool
				err     error
			)

			claims, ok := authz.AuthClaimsFromContext(ctx)
			if !ok {
				return nil, ErrMissingClaims
			}

			if claims.Subject == nil || claims.Action == nil || claims.Resource == nil {
				return nil, ErrInvalidClaims
			}

			allowed, err = authorizer.IsAuthorized(ctx, *claims.Subject, *claims.Action, *claims.Resource, *claims.Project)
			if err != nil {
				return nil, err
			}
			if !allowed {
				return nil, ErrUnauthorized
			}

			return handler(ctx, req)
		}
	}
}
