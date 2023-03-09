package authn

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/auth/authz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

const (
	reason string = "FORBIDDEN"
)

var (
	ErrInvalidAuthUser = errors.Unauthorized(reason, "invalid auth user")
)

// Server is a server authenticator middleware.
func Server(authenticator authn.Authenticator, userCreator authn.SecurityUserCreator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claims, err := authenticator.Authenticate(ctx)
			if err != nil {
				return nil, err
			}

			ctx = authn.ContextWithAuthClaims(ctx, claims)
			if userCreator != nil {
				securityUser := userCreator(claims)
				if err := securityUser.ParseFromContext(ctx); err != nil {
					return nil, ErrInvalidAuthUser
				}
				ctx = authn.ContextWithAuthUser(ctx, securityUser)
				action := authz.Action(securityUser.GetAction())
				subject := authz.Subject(securityUser.GetSubject())
				resource := authz.Resource(securityUser.GetObject())
				project := authz.Project(securityUser.GetDomain())
				ctx = authz.ContextWithAuthClaims(ctx, &authz.AuthClaims{
					Subject:  &subject,
					Action:   &action,
					Resource: &resource,
					Project:  &project,
				})
			}

			return handler(ctx, req)
		}
	}
}
