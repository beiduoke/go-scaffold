package authz

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/pkg/auth/authz"
)

type Option func(*options)

type options struct {
}

func NewContext(ctx context.Context, claims *authz.AuthClaims) context.Context {
	return authz.ContextWithAuthClaims(ctx, claims)
}

func FromContext(ctx context.Context) (*authz.AuthClaims, bool) {
	return authz.AuthClaimsFromContext(ctx)
}
