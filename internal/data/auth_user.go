package data

import (
	"context"
	"errors"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

type AuthRole struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Sort          int32  `json:"sort" form:"sort"`
	DefaultRouter string `json:"defaultRouter" form:"defaultRouter"`
}

type AuthDept struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type AuthUser struct {
	ID            uint       `json:"id" form:"id"`
	DomainID      uint       `json:"domainId" form:"domainId"`
	Name          string     `json:"name" form:"name"`
	NickName      string     `json:"nickName" form:"nickName"`
	RealName      string     `json:"realName" form:"realName"`
	Avatar        string     `json:"avatar" form:"avatar"`
	Birthday      *time.Time `json:"birthday" form:"birthday"`
	Gender        int32      `json:"gender" form:"gender"`
	Phone         string     `json:"phone" form:"phone"`
	Email         string     `json:"email" form:"email"`
	State         int32      `json:"state" form:"state"`
	Remarks       string     `json:"remarks" form:"remarks"`
	DeptId        uint       `json:"deptId" form:"deptId"`
	LastUseRoleID uint       `json:"lastUseRoleId" form:"lastUseRoleId"`
	LastLoginAt   *time.Time `json:"lastLoginAt" form:"lastLoginAt"`
	LastUseRole   *AuthRole  `json:"lastUseRole" form:"lastUseRole"`
	Roles         []AuthRole `json:"roles" form:"roles"`
	Dept          *AuthDept  `json:"dept" form:"dept"`
}

var _ authn.SecurityUser = (*securityUser)(nil)

func NewSecurityUser(logger log.Logger, data *Data, userRepo biz.UserRepo) authn.SecurityUserCreator {
	log := log.NewHelper(log.With(logger, "module", "data/securityUserCreator"))
	return func(ac *authn.AuthClaims) authn.SecurityUser {
		if ac == nil {
			log.Error("auth claims creator fail ac == nil")
		}
		return &securityUser{options: Options{data: data, log: log, userRepo: userRepo.(*UserRepo), authClaims: ac}}
	}
}

type Options struct {
	data       *Data
	log        *log.Helper
	authClaims *authn.AuthClaims
	userRepo   *UserRepo
}

type securityUser struct {
	options Options
	// 用户
	user uint
	// 域/租户
	domain uint
	// 角色/主题
	subject uint
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
		// if header.Kind() == transport.KindHTTP {
		// 	if ht, ok := header.(http.Transporter); ok {
		// 		su.Object = ht.Request().URL.Object
		// 		su.Action = ht.Request().Action
		// 	}
		// }
	} else {
		return errors.New("parse from request header")
	}
	authUser, err := su.options.userRepo.GetLoginIDCache(ctx, su.options.authClaims.Subject)
	if err != nil {
		su.options.log.Errorf("result auth user fail: %v", err)
		return err
	}
	su.user = authUser.ID
	su.domain = authUser.DomainID
	su.subject = authUser.LastUseRoleID
	return nil
}

// GetSubject returns the subject of the token.
func (su *securityUser) GetSubject() string {
	return convert.UnitToString(su.subject)
}

// GetObject returns the object of the token.
func (su *securityUser) GetObject() string {
	return su.object
}

// GetAction returns the action of the token.
func (su *securityUser) GetAction() string {
	return su.action
}

// GetDomain returns the domain of the token.
func (su *securityUser) GetDomain() string {
	return convert.UnitToString(su.domain)
}

// GetID returns the user of the token.
func (su *securityUser) GetUser() string {
	return convert.UnitToString(su.user)
}
