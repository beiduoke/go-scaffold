package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
)

// HasSystemSuperAdmin 是否为系统超级管理员
func (d *Data) HasSystemSuperAdmin(ctx context.Context) bool {
	superAdminConfig := d.cfg.Base.GetAdmin()
	if d.CtxDomainID(ctx) == uint(superAdminConfig.GetDomainId()) &&
		d.CtxUserID(ctx) == uint(superAdminConfig.GetUserId()) {
		return true
	}
	return false
}

// HasDomainSuperUser 是否为租户超级管理员
func (d *Data) HasDomainSuperUser(ctx context.Context) bool {
	sysDomain, userId := SysDomain{}, d.CtxUserID(ctx)
	result := d.DBD(ctx).Model(sysDomain).Select("SuperUserID").Last(&sysDomain)
	if result.RowsAffected > 0 && sysDomain.SuperUserID == userId {
		return true
	}
	return false
}

// CtxAuthUser 获取上下文中用户信息（该信息是从redis读取）
func (d *Data) CtxAuthUser(ctx context.Context) authn.SecurityUser {
	security, success := authn.AuthUserFromContext(ctx)
	if !success {
		return &securityUser{}
	}
	return security
}

// CtxDomainID 获取上下文中用户租户ID
func (d *Data) CtxDomainID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetDomain())
}

// CtxUserID 获取上下文中用户ID
func (d *Data) CtxUserID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetUser())
}

// CtxRoleID 获取上下文中用户角色ID
func (d *Data) CtxRoleID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetSubject())
}

// CtxAuthClaimSubject 获取上下文中的认证主题信息（UUID）用户指纹
func (d *Data) CtxAuthClaimSubject(ctx context.Context) string {
	claims, success := authn.AuthClaimsFromContext(ctx)
	if !success {
		return ""
	}
	return claims.Subject
}

// CtxAuthClaimScopes 获取上下文中的认证范围信息
func (d *Data) CtxAuthClaimScopes(ctx context.Context) map[string]bool {
	claims, success := authn.AuthClaimsFromContext(ctx)
	if !success {
		return map[string]bool{}
	}
	return claims.Scopes
}
