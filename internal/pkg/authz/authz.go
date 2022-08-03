package authz

import (
	"context"
	"errors"
	"fmt"

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
	AuthorityId string
	Domain      string
	ID          string
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
		str, ok := claims.(jwtV4.MapClaims)[ClaimAuthorityId]
		if ok {
			su.AuthorityId = str.(string)
		}
		str, ok = claims.(jwtV4.MapClaims)[Domain]
		if ok {
			su.Domain = str.(string)
		}
		str, ok = claims.(jwtV4.MapClaims)[ID]
		if ok {
			su.ID = str.(string)
		}
		str, ok = claims.(jwtV4.MapClaims)[ExpiresAt]
		if ok {
			su.ExpiresAt = str.(int64)
		}
	} else {
		return errors.New("jwt claim missing")
	}

	if header, ok := transport.FromServerContext(ctx); ok {
		su.Path = header.Operation()
		su.Method = "*"
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
			ExpiresAt:        float64(su.ExpiresAt),
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

	strAuthorityId, ok := mc[ClaimAuthorityId]
	if ok {
		su.AuthorityId = strAuthorityId.(string)
	}
	strExpiresAt, ok := mc[ExpiresAt]
	if ok {
		// expiresAtFloat := strconv.FormatFloat(strExpiresAt.(float64), 'f', 0, 64)
		su.ExpiresAt = strExpiresAt.(int64)
	}
	strDomain, ok := mc[Domain]
	if ok {
		su.Domain = strDomain.(string)
	}
	strId, ok := mc[ID]
	if ok {
		su.ID = strId.(string)
	}

	return nil
}
