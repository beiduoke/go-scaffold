package data

import (
	"context"
	"errors"

	"github.com/beiduoke/go-scaffold/pkg/auth/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"gorm.io/gorm"
)

func (d *Data) HasSuperAdmin(ctx context.Context) bool {
	if d.CtxDomainID(ctx) == 1 && d.CtxUserID(ctx) == 1 {
		return true
	}
	return false
}

func (d *Data) HasDomainSuperUser(ctx context.Context) bool {
	sysDomain, domainId, userId := SysDomain{}, d.CtxDomainID(ctx), d.CtxUserID(ctx)
	result := d.DB(ctx).Model(sysDomain).Debug().Select("SuperUserID").Last(&sysDomain, domainId)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && sysDomain.SuperUserID == userId {
		return true
	}
	return false
}

func (d *Data) CtxAuthUser(ctx context.Context) authn.SecurityUser {
	security, success := authn.AuthUserFromContext(ctx)
	if !success {
		return &securityUser{}
	}
	return security
}

func (d *Data) CtxDomainID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetDomain())
}

func (d *Data) CtxUserID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetUser())
}

func (d *Data) CtxRoleID(ctx context.Context) uint {
	return convert.StringToUint(d.CtxAuthUser(ctx).GetSubject())
}
