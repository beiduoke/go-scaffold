package data

import (
	"context"
	"errors"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

var _ authn.SecurityUser = (*securityUser)(nil)

func NewSecurityUser(logger log.Logger, data *Data, authToken *AuthTokenRepo) authn.SecurityUserCreator {
	log := log.NewHelper(log.With(logger, "module", "data/securityUserCreator"))
	return func(ac *authn.AuthClaims) authn.SecurityUser {
		if ac == nil {
			log.Error("auth claims creator fail ac == nil")
		}
		return &securityUser{options: SecurityOptions{data: data, log: log, authClaims: ac, authToken: authToken}}
	}
}

type SecurityOptions struct {
	data       *Data
	log        *log.Helper
	authClaims *authn.AuthClaims
	authToken  *AuthTokenRepo
}

type securityUser struct {
	options SecurityOptions
	// 用户
	user uint32
	// 域/租户
	domain uint32
	// 角色/主题
	subject uint32
	// 资源/路由
	object string
	// 方法
	action string
}

// ParseFromContext parses the user from the context.
func (su *securityUser) ParseFromContext(ctx context.Context) error {
	if header, ok := transport.FromServerContext(ctx); ok {
		su.object = header.Operation()
		su.action = "*"
		if header.Kind() == transport.KindHTTP {
			// if ht, ok := header.(http.Transporter); ok {
			// 	su.object = ht.Request().URL.Object
			// 	su.action = ht.Request().Action
			// }
		}
	} else {
		return errors.New("parse from request header")
	}

	user := convert.StringToUnit32(su.options.authClaims.Subject)
	authToken := su.options.authToken.GetToken(ctx, user)
	if authToken == "" {
		err := errors.New("result auth user fail: auth token null")
		su.options.log.Error(err)
		return err
	}
	su.user = user
	// su.domain = authToken.DomainID
	// su.subject = authToken.LastUseRoleID
	return nil
}

// GetObject returns the object of the token.
func (su *securityUser) GetObject() string {
	return su.object
}

// GetAction returns the action of the token.
func (su *securityUser) GetAction() string {
	return su.action
}

// GetSubject returns the subject of the token.
func (su *securityUser) GetSubject() string {
	return convert.Unit32ToString(su.subject)
}

// GetDomain returns the domain of the token.
func (su *securityUser) GetDomain() string {
	return convert.Unit32ToString(su.domain)
}

// // GetID returns the user of the token.
func (su *securityUser) GetUser() string {
	return convert.Unit32ToString(su.user)
}
