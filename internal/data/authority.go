package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorityRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthorityRepo .
func NewAuthorityRepo(logger log.Logger, data *Data) biz.AuthorityRepo {
	return &AuthorityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthorityRepo) toModel(d *biz.Authority) *SysAuthority {
	if d == nil {
		return nil
	}
	return &SysAuthority{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:          d.Name,
		DefaultRouter: d.DefaultRouter,
		State:         d.State,
		ParentID:      d.ParentID,
	}
}

func (r *AuthorityRepo) toBiz(d *SysAuthority) *biz.Authority {
	if d == nil {
		return nil
	}
	return &biz.Authority{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Name:      d.Name,
	}
}

func (r *AuthorityRepo) Save(ctx context.Context, g *biz.Authority) (*biz.Authority, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *AuthorityRepo) Update(ctx context.Context, g *biz.Authority) (*biz.Authority, error) {
	return g, nil
}

func (r *AuthorityRepo) FindByName(ctx context.Context, s string) (*biz.Authority, error) {
	authority := SysAuthority{}
	result := r.data.DB(ctx).Last(&authority, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&authority), nil
}

func (r *AuthorityRepo) FindByID(ctx context.Context, id uint) (*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListByIDs(ctx context.Context, id ...uint) (authorities []*biz.Authority, err error) {
	db := r.data.DB(ctx).Model(&SysDomain{})
	sysAuthorities := []*SysAuthority{}

	err = db.Find(&sysAuthorities).Error
	if err != nil {
		return authorities, err
	}
	for _, v := range sysAuthorities {
		authorities = append(authorities, r.toBiz(v))
	}
	return
}

func (r *AuthorityRepo) ListByName(ctx context.Context, name string) ([]*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListAll(ctx context.Context) ([]*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (authorities []*biz.Authority, total int64) {
	db := r.data.DB(ctx).Model(&SysAuthority{})
	sysAuthorities := []*SysAuthority{}
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

	result := db.Limit(int(handler.GetPageSize())).Find(&sysAuthorities)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysAuthorities {
		authorities = append(authorities, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(authorities))
	}

	return authorities, total
}
