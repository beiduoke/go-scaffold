package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DomainAuthorityUserRepo struct {
	data      *Data
	log       *log.Helper
	domain    DomainRepo
	authority AuthorityRepo
}

// NewDomainAuthorityUserRepo .
func NewDomainAuthorityUserRepo(data *Data, logger log.Logger) *DomainAuthorityUserRepo {
	return &DomainAuthorityUserRepo{
		data:      data,
		log:       log.NewHelper(logger),
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
