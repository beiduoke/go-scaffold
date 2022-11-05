package data

import (
	"context"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const cacheKeyToken string = "%d-%d"

type UserRepo struct {
	data      *Data
	log       *log.Helper
	domain    DomainRepo
	authority AuthorityRepo
}

// NewUserRepo .
func NewUserRepo(logger log.Logger, data *Data) biz.UserRepo {
	return &UserRepo{
		data:      data,
		log:       log.NewHelper(logger),
		authority: AuthorityRepo{},
	}
}

func (r *UserRepo) toModel(d *biz.User) *SysUser {
	if d == nil {
		return nil
	}
	domains := []SysDomain{}
	for _, v := range d.Domains {
		domains = append(domains, *r.domain.toModel(v))
	}
	authorities := []SysAuthority{}
	for _, v := range d.Authorities {
		authorities = append(authorities, *r.authority.toModel(v))
	}
	return &SysUser{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:        d.Name,
		Avatar:      d.Avatar,
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

func (r *UserRepo) toBiz(d *SysUser) *biz.User {
	if d == nil {
		return nil
	}
	domains := []*biz.Domain{}
	for _, v := range d.Domains {
		domains = append(domains, r.domain.toBiz(&v))
	}
	authorities := []*biz.Authority{}
	for _, v := range d.Authorities {
		authorities = append(authorities, r.authority.toBiz(&v))
	}
	return &biz.User{
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		ID:          d.ID,
		Avatar:      d.Avatar,
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
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*biz.User, error) {
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

func (r *UserRepo) Delete(ctx context.Context, g *biz.User) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *UserRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (users []*biz.User, total int64) {
	db := r.data.DB(ctx).Model(&SysUser{})
	sysUsers := []*SysUser{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}
	if !handler.GetNopaging() {
		db = db.Count(&total).Offset(handler.GetPageOffset())
	}
	result := db.Limit(int(handler.GetPageSize())).Find(&sysUsers)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysUsers {
		users = append(users, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(users))
	}

	return users, total
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

func (r *UserRepo) SetTokenCache(ctx context.Context, claims biz.AuthClaims) error {
	key := fmt.Sprintf(cacheKeyToken, claims.ID, claims.Domain)
	result := r.data.rdb.Set(ctx, key, claims.Token, time.Until(*claims.ExpiresAt))
	return result.Err()
}

func (r *UserRepo) GetTokenCache(ctx context.Context, claims biz.AuthClaims) error {
	key := fmt.Sprintf(cacheKeyToken, claims.ID, claims.Domain)
	result := r.data.rdb.Get(ctx, key)
	return result.Err()
}

// HandleDomain 绑定领域
func (r *UserRepo) HandleDomain(ctx context.Context, g *biz.User) error {
	for _, domain := range g.Domains {
		fmt.Println(convert.UnitToString(g.ID), convert.UnitToString(domain.DefaultAuthorityID), convert.UnitToString(domain.ID))
		if _, err := r.data.enforcer.AddRoleForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(domain.DefaultAuthorityID), convert.UnitToString(domain.ID)); err != nil {
			r.log.Errorf("领域绑定失败 %v", err)
		}
	}
	return nil
}

// HandleDomainAuthority 绑定领域权限
func (r *UserRepo) HandleDomainAuthority(ctx context.Context, g *biz.User) error {
	domainId := r.data.Domain(ctx)
	for _, v := range g.Authorities {
		if _, err := r.data.enforcer.AddRoleForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(v.ID), domainId); err != nil {
			r.log.Errorf("领域权限绑定失败 %v", err)
		}
	}

	return nil
}
