package authn

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/pkg/auth/authn"
)

type ContextWithToken func(context.Context, string) context.Context

type Option func(*options)

type options struct {
	claims               authn.AuthClaims
	contextWithTokenFunc ContextWithToken
}

func WithAuthClaims(claims authn.AuthClaims) Option {
	return func(o *options) {
		o.claims = claims
	}
}

func WithContextToken(f ContextWithToken) Option {
	return func(o *options) {
		o.contextWithTokenFunc = f
	}
}

func NewContext(ctx context.Context, claims *authn.AuthClaims) context.Context {
	return authn.ContextWithAuthClaims(ctx, claims)
}

func FromContext(ctx context.Context) (*authn.AuthClaims, bool) {
	return authn.AuthClaimsFromContext(ctx)
}
