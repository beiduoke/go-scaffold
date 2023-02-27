package auth

import (
	"time"

	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/pkg/errors"
)

type ClaimsOption func(*ClaimsOptions)

type ClaimsOptions struct {
	securityKey string
	expiresAt   time.Time
}

func newClaimsOptions(opts ...ClaimsOption) *ClaimsOptions {
	co := &ClaimsOptions{
		expiresAt: time.Now().Add(time.Hour * 24),
	}
	for _, o := range opts {
		o(co)
	}
	return co
}

func WidthAuthExpiresAt(d time.Duration) ClaimsOption {
	return func(ac *ClaimsOptions) {
		ac.expiresAt = time.Now().Add(d)
	}
}

func WidthAuthSecurityKey(s string) ClaimsOption {
	return func(ac *ClaimsOptions) {
		ac.securityKey = s
	}
}

var _ AuthClaims = (*claims)(nil)

type claims struct {
	options *ClaimsOptions
	token   string
}

func NewAuthClaims(opts ...ClaimsOption) AuthClaims {
	c := &claims{
		options: newClaimsOptions(),
	}
	for _, opt := range opts {
		opt(c.options)
	}

	return c
}

func (ac claims) Token() string {
	return ac.token
}

func (ac claims) ExpiresAt() time.Time {
	return ac.options.expiresAt
}

func (c *claims) CreateToken(user string) error {
	options := c.options
	securityUser := myAuthz.NewSecurityUserData(
		myAuthz.WithUser(user),
		myAuthz.WithExpires(options.expiresAt),
	)
	if c.token = securityUser.CreateAccessJwtToken([]byte(options.securityKey)); c.token == "" {
		return errors.New("token生成失败")
	}
	return nil
}
