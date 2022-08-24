package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	data                *Data
	log                 *log.Helper
	domain              DomainRepo
	authority           AuthorityRepo
	domainAuthorityUser DomainAuthorityUserRepo
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data:      data,
		log:       log.NewHelper(logger),
		domain:    DomainRepo{},
		authority: AuthorityRepo{},
	}
}

func (r *UserRepo) toModel(d *biz.User) *SysUser {
	if d == nil {
		return nil
	}
	domains := []SysDomain{}
	for _, v := range d.Domains {
		domains = append(domains, *r.domain.toModel(&v))
	}
	authorities := []SysAuthority{}
	for _, v := range d.Authorities {
		authorities = append(authorities, *r.authority.toModel(&v))
	}
	domainAuthorityUsers := []SysDomainAuthorityUser{}
	for _, v := range d.DomainAuthorityUsers {
		domainAuthorityUsers = append(domainAuthorityUsers, *r.domainAuthorityUser.toModel(&v))
	}
	return &SysUser{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:     d.Name,
		NickName: d.NickName,
		RealName: d.RealName,
		Password: d.Password,
		Birthday: d.Birthday,
		Gender:   d.Gender,
		Mobile:   d.Mobile,
		Email:    d.Email,
		State:    d.State,
		// DomainAuthorityUsers: domainAuthorityUsers,
	}
}

func (r *UserRepo) toBiz(d *SysUser) *biz.User {
	if d == nil {
		return nil
	}
	domains := []biz.Domain{}
	for _, v := range d.Domains {
		domains = append(domains, *r.domain.toBiz(&v))
	}
	authorities := []biz.Authority{}
	for _, v := range d.Authorities {
		authorities = append(authorities, *r.authority.toBiz(&v))
	}
	return &biz.User{
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		ID:          d.ID,
		Name:        d.Name,
		NickName:    d.NickName,
		RealName:    d.RealName,
		Password:    d.Password,
		Birthday:    d.Birthday,
		Gender:      d.Gender,
		Mobile:      d.Mobile,
		Email:       d.Email,
		State:       d.State,
		Domains:     domains,
		Authorities: authorities,
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *UserRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DB(ctx).Last(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByName(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DB(ctx).Last(&user, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}
func (r *UserRepo) FindByMobile(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DB(ctx).Last(&user, "mobile = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}
func (r *UserRepo) FindByEmail(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DB(ctx).Last(&user, "email = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}
func (r *UserRepo) ListByName(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DB(ctx).Find(&sysUsers, "name LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByMobile(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DB(ctx).Find(&sysUsers, "mobile LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByEmail(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DB(ctx).Find(&sysUsers, "email LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}
