package authn

import (
	"context"
	"strconv"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
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

type Result struct {
	UserId   uint32
	UserName string
}

func AuthFromContext(ctx context.Context) (*Result, error) {
	claims, ok := authn.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return nil, ErrExtractSubjectFailed
	}

	return &Result{
		UserId:   uint32(userId),
		UserName: "",
	}, nil
}
