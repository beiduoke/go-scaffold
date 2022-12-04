package authz

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/beiduoke/go-scaffold/pkg/authz"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

var _ authz.SecurityUser = (*SecurityUser)(nil)

const (
	User                = "user"
	Expires             = "exp"
	ClaimSubject        = "subject"
	ClaimDomain         = "domain"
	HeaderDomainIDKey   = "X-Domain-ID"
	HeaderDomainCodeKey = "X-Domain-Code"
)

type Option func(*SecurityUser)

// WithUser 设置用户
func WithUser(user string) Option {
	return func(o *SecurityUser) {
		o.User = user
	}
}

// WithSubject 设置角色
func WithSubject(subject string) Option {
	return func(o *SecurityUser) {
		o.Subject = subject
	}
}

// WithDomain 设置域/租户
func WithDomain(domain string) Option {
	return func(o *SecurityUser) {
		o.Domain = domain
	}
}

// WithExpires 设置过期时间
func WithExpires(expires time.Time) Option {
	return func(o *SecurityUser) {
		o.Expires = expires.Unix()
	}
}

type SecurityUser struct {
	// 用户
	User string `json:"user,omitempty" form:"user"`
	// 域/租户
	Domain string `json:"domain,omitempty" form:"domain"`
	// 角色
	Subject string `json:"sub,omitempty," form:"subject"`
	// 资源
	Object string `json:"object,omitempty" form:"object"`
	// 方法
	Action string `json:"action,omitempty" form:"action"`
	// 过期
	Expires int64 `json:"expires,omitempty" form:"expires"`
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
		su.Object = header.Operation()
		su.Action = "*"
		if domainId := header.RequestHeader().Get(HeaderDomainIDKey); domainId != "" {
			su.Domain = domainId
		} else if domainCode := header.RequestHeader().Get(HeaderDomainCodeKey); domainCode != "" {
			su.Domain = domainCode
		}
		// if header.Kind() == transport.KindHTTP {
		// 	if ht, ok := header.(http.Transporter); ok {
		// 		su.Object = ht.Request().URL.Object
		// 		su.Action = ht.Request().Action
		// 	}
		// }
	} else {
		return errors.New("jwt claim missing")
	}
	return nil
}

// GetUser 用户
func (su *SecurityUser) GetUser() string {
	return su.User
}

// GetSubject 角色
func (su *SecurityUser) GetSubject() string {
	return su.Subject
}

// GetObject 资源
func (su *SecurityUser) GetObject() string {
	return su.Object
}

// GetAction 方法
func (su *SecurityUser) GetAction() string {
	return su.Action
}

// GetDomain 域/租户
func (su *SecurityUser) GetDomain() string {
	return su.Domain
}

func (su *SecurityUser) CreateAccessJwtToken(secretKey []byte) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			User:         su.User,
			Expires:      su.Expires,
			ClaimSubject: su.Subject,
			ClaimDomain:  su.Domain,
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
	// 用户User
	str, ok := mc[User]
	if ok {
		su.User = str.(string)
	}
	// 权限
	str, ok = mc[ClaimSubject]
	if ok {
		su.Subject = str.(string)
	}
	// 领域
	str, ok = mc[ClaimDomain]
	if ok {
		su.Domain = str.(string)
	}
	// 过期时间
	str, ok = mc[Expires]
	if ok {
		switch exp := str.(type) {
		case string:
			su.Expires, _ = strconv.ParseInt(exp, 10, 64)
		case float64:
			su.Expires, _ = strconv.ParseInt(strconv.FormatFloat(exp, 'f', 0, 64), 10, 64)
		}
	}
	return nil
}

func ParseFromContext(ctx context.Context) authz.SecurityUser {
	securityUser := NewSecurityUser()
	if securityUser.ParseFromContext(ctx) != nil {
		return &SecurityUser{}
	}
	return securityUser
}
