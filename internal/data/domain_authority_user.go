package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DomainAuthorityUserRepo struct {
	data      *Data
	log       *log.Helper
	enforcer  stdcasbin.IEnforcer
	domain    DomainRepo
	authority AuthorityRepo
}

// NewDomainAuthorityUserRepo .
func NewDomainAuthorityUserRepo(data *Data, enforcer stdcasbin.IEnforcer, logger log.Logger) *DomainAuthorityUserRepo {
	return &DomainAuthorityUserRepo{
		data:      data,
		log:       log.NewHelper(logger),
		enforcer:  enforcer,
		domain:    DomainRepo{},
		authority: AuthorityRepo{},
	}
}

func (r *DomainAuthorityUserRepo) toModel(d *biz.DomainAuthorityUser) *SysDomainAuthorityUser {
	if d == nil {
		return nil
	}
	return &SysDomainAuthorityUser{
		DomainID:    d.DomainID,
		AuthorityID: d.AuthorityID,
		UserID:      d.UserID,
		CreatedAt:   d.CreatedAt,
	}
}

func (r *DomainAuthorityUserRepo) toBiz(d *SysDomainAuthorityUser) *biz.DomainAuthorityUser {
	if d == nil {
		return nil
	}
	return &biz.DomainAuthorityUser{
		DomainID:    d.DomainID,
		AuthorityID: d.AuthorityID,
		UserID:      d.UserID,
		CreatedAt:   d.CreatedAt,
	}
}

func (r *DomainAuthorityUserRepo) Save(ctx context.Context, g *biz.DomainAuthorityUser) (*biz.DomainAuthorityUser, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *DomainAuthorityUserRepo) Update(ctx context.Context, g *biz.DomainAuthorityUser) (*biz.DomainAuthorityUser, error) {
	return g, nil
}

func (r *DomainAuthorityUserRepo) FindByID(ctx context.Context, id int64) (*biz.DomainAuthorityUser, error) {
	user := SysDomainAuthorityUser{}
	result := r.data.DB(ctx).Last(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *DomainAuthorityUserRepo) ListAll(ctx context.Context) ([]*biz.DomainAuthorityUser, error) {

	return nil, nil
}

func (r *DomainAuthorityUserRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (domainAuthorityUsers []*biz.DomainAuthorityUser, total int64) {
	db := r.data.DB(ctx).Model(&SysDomainAuthorityUser{})
	sysDomainAuthorityUsers := []*SysDomainAuthorityUser{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}
	result := db.Count(&total).Offset(handler.GetPageOffset()).Limit(int(handler.GetPageSize())).Find(&sysDomainAuthorityUsers)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDomainAuthorityUsers {
		domainAuthorityUsers = append(domainAuthorityUsers, r.toBiz(v))
	}
	return domainAuthorityUsers, total
}
