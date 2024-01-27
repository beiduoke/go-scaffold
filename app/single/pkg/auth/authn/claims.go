package authn

import (
	"bytes"
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type ctxKey string

var (
	authClaimsContextKey = ctxKey("authn-claims")
	authUserContextKey   = ctxKey("authn-user")
)

type ScopeSet map[string]bool

// AuthClaims contains claims that are included in OIDC standard claims.
// See https://openid.net/specs/openid-connect-core-1_0.html#IDToken
type AuthClaims struct {
	Subject string

	// Scopes see: https://datatracker.ietf.org/doc/html/rfc6749#section-3.3
	Scopes ScopeSet
}

// ContextWithAuthClaims injects the provided AuthClaims into the parent context.
func ContextWithAuthClaims(parent context.Context, claims *AuthClaims) context.Context {
	return context.WithValue(parent, authClaimsContextKey, claims)
}

// AuthClaimsFromContext extracts the AuthClaims from the provided ctx (if any).
func AuthClaimsFromContext(ctx context.Context) (*AuthClaims, bool) {
	claims, ok := ctx.Value(authClaimsContextKey).(*AuthClaims)
	if !ok {
		return nil, false
	}

	return claims, true
}

// ContextWithAuthClaims injects the provided AuthClaims into the parent context.
func ContextWithAuthUser(parent context.Context, user SecurityUser) context.Context {
	return context.WithValue(parent, authUserContextKey, user)
}

// AuthUserFromContext extracts the AuthUser from the provided ctx (if any).
func AuthUserFromContext(ctx context.Context) (SecurityUser, bool) {
	claims, ok := ctx.Value(authUserContextKey).(SecurityUser)
	if !ok {
		return nil, false
	}

	return claims, true
}

func AuthClaimsToJwtClaims(raw AuthClaims) jwt.Claims {
	claims := jwt.MapClaims{
		"sub": raw.Subject,
	}

	var buffer bytes.Buffer
	count := len(raw.Scopes)
	idx := 0
	for scope := range raw.Scopes {
		buffer.WriteString(scope)
		if idx != count-1 {
			buffer.WriteString(" ")
		}
		idx++
	}
	str := buffer.String()
	if len(str) > 0 {
		claims["scope"] = buffer.String()
	}

	return claims
}

func MapClaimsToAuthClaims(rawClaims jwt.MapClaims) (*AuthClaims, error) {
	// optional subject
	var subject = ""
	if subjectClaim, ok := rawClaims["sub"]; ok {
		if subject, ok = subjectClaim.(string); !ok {
			return nil, ErrInvalidSubject
		}
	}

	claims := &AuthClaims{
		Subject: subject,
		Scopes:  make(ScopeSet),
	}

	// optional scopes
	if scopeKey, ok := rawClaims["scope"]; ok {
		if scope, ok := scopeKey.(string); ok {
			scopes := strings.Split(scope, " ")
			for _, s := range scopes {
				claims.Scopes[s] = true
			}
		}
	}

	return claims, nil
}

func JwtClaimsToAuthClaims(rawClaims jwt.Claims) (*AuthClaims, error) {
	claims, ok := rawClaims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}
	return MapClaimsToAuthClaims(claims)
}
