package authz

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/beiduoke/go-scaffold/pkg/authz"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

var _ authz.SecurityUser = (*SecurityUser)(nil)

const (
	ClaimAuthorityId = "authorityId"
	Domain           = "domain"
	ID               = "id"
	ExpiresAt        = "exp"
)

type Option func(*SecurityUser)

func WithID(id string) Option {
	return func(o *SecurityUser) {
		o.ID = id
	}
}

func WithAuthorityId(authorityId string) Option {
	return func(o *SecurityUser) {
		o.AuthorityId = authorityId
	}
}

func WithDomain(domain string) Option {
	return func(o *SecurityUser) {
		o.Domain = domain
	}
}
func WithExpiresAt(expiresAt int64) Option {
	return func(o *SecurityUser) {
		o.ExpiresAt = expiresAt
	}
}

type SecurityUser struct {
	Path        string
	Method      string
	ID          string
	Domain      string
	AuthorityId string
	ExpiresAt   int64
}

func NewSecurityUserData(opts ...Option) *SecurityUser {
	s := &SecurityUser{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func NewSecurityUser() authz.SecurityUser {
	return &SecurityUser{}
}

func (su *SecurityUser) ParseFromContext(ctx context.Context) error {
	if claims, ok := jwt.FromContext(ctx); ok {
		err := su.ParseAccessJwtToken(claims)
		if !errors.Is(err, nil) {
			return err
		}
	} else {
		return errors.New("jwt claim missing")
	}

	if header, ok := transport.FromServerContext(ctx); ok {
		su.Path = header.Operation()
		su.Method = "*"
		// if header.Kind() == transport.KindHTTP {
		// 	if ht, ok := header.(http.Transporter); ok {
		// 		su.Path = ht.Request().URL.Path
		// 		su.Method = ht.Request().Method
		// 	}
		// }
	} else {
		return errors.New("jwt claim missing")
	}
	return nil
}

func (su *SecurityUser) GetSubject() string {
	return su.AuthorityId
}

func (su *SecurityUser) GetObject() string {
	return su.Path
}

func (su *SecurityUser) GetAction() string {
	return su.Method
}

func (su *SecurityUser) GetDomain() string {
	return su.Domain
}

func (su *SecurityUser) CreateAccessJwtToken(secretKey []byte) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			ClaimAuthorityId: su.AuthorityId,
			Domain:           su.Domain,
			ID:               su.ID,
			ExpiresAt:        su.ExpiresAt,
		})

	signedToken, err := claims.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return signedToken
}

func (su *SecurityUser) ParseAccessJwtTokenFromContext(ctx context.Context) error {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		fmt.Println("ParseAccessJwtTokenFromContext 1")
		return errors.New("no jwt token in context")
	}
	if err := su.ParseAccessJwtToken(claims); err != nil {
		fmt.Println("ParseAccessJwtTokenFromContext 2")
		return err
	}
	return nil
}

func (su *SecurityUser) ParseAccessJwtTokenFromString(token string, secretKey []byte) error {
	parseAuth, err := jwtV4.Parse(token, func(*jwtV4.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	claims, ok := parseAuth.Claims.(jwtV4.MapClaims)
	if !ok {
		return errors.New("no jwt token in context")
	}

	if err := su.ParseAccessJwtToken(claims); err != nil {
		return err
	}

	return nil
}

func (su *SecurityUser) ParseAccessJwtToken(claims jwtV4.Claims) error {
	if claims == nil {
		return errors.New("claims is nil")
	}

	mc, ok := claims.(jwtV4.MapClaims)
	if !ok {
		return errors.New("claims is not map claims")
	}
	// 用户ID
	str, ok := mc[ID]
	if ok {
		su.ID = str.(string)
	}
	// 权限
	str, ok = mc[ClaimAuthorityId]
	if ok {
		su.AuthorityId = str.(string)
	}
	// 领域
	str, ok = mc[Domain]
	if ok {
		su.Domain = str.(string)
	}
	// 过期时间
	str, ok = mc[ExpiresAt]
	if ok {
		switch exp := str.(type) {
		case string:
			su.ExpiresAt, _ = strconv.ParseInt(exp, 10, 64)
		case float64:
			su.ExpiresAt, _ = strconv.ParseInt(strconv.FormatFloat(exp, 'f', 0, 64), 10, 64)
		}
	}
	return nil
}
