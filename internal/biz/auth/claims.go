package auth

import (
	"time"

	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/pkg/errors"
)

type ClaimsOptions func(*Claims)

type Claims struct {
	User      string
	Domain    string
	ExpiresAt time.Time
	token     string
}

func (ac Claims) Token() string {
	return ac.token
}

func WidthAuthExpiresAt(d time.Duration) ClaimsOptions {
	return func(ac *Claims) {
		ac.ExpiresAt = time.Now().Add(d)
	}
}

func WidthAuthUser(u string) ClaimsOptions {
	return func(ac *Claims) {
		ac.User = u
	}
}

func NewAuthClaims(opts ...ClaimsOptions) *Claims {
	claims := &Claims{
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	for _, opt := range opts {
		opt(claims)
	}

	return claims
}

func (c *Claims) CreateToken(key string) error {
	securityUser := myAuthz.NewSecurityUserData(
		myAuthz.WithUser(c.User),
		myAuthz.WithDomain(c.Domain),
		myAuthz.WithExpires(c.ExpiresAt),
	)
	if c.token = securityUser.CreateAccessJwtToken([]byte(key)); c.token == "" {
		return errors.New("token生成失败")
	}
	return nil
}
