package auth

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth"
)

type ContextWithToken func(context.Context, string) context.Context

type Option func(*options)

type options struct {
	claims               auth.AuthClaims
	contextWithTokenFunc ContextWithToken
}

func WithAuthClaims(claims auth.AuthClaims) Option {
	return func(o *options) {
		o.claims = claims
	}
}

func WithContextToken(f ContextWithToken) Option {
	return func(o *options) {
		o.contextWithTokenFunc = f
	}
}

func NewContext(ctx context.Context, claims *auth.AuthClaims) context.Context {
	return auth.ContextWithAuthClaims(ctx, claims)
}

func FromContext(ctx context.Context) (*auth.AuthClaims, bool) {
	return auth.AuthClaimsFromContext(ctx)
}
