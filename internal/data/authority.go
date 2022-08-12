package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type AuthorityRepo struct {
	data   *Data
	log    *log.Helper
	policy persist.Adapter
}

// NewAuthorityRepo .
func NewAuthorityRepo(data *Data, policy persist.Adapter, logger log.Logger) biz.AuthorityRepo {
	return &AuthorityRepo{
		data:   data,
		log:    log.NewHelper(logger),
		policy: policy,
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
		Name: d.Name,
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
	return g, nil
}

func (r *AuthorityRepo) Update(ctx context.Context, g *biz.Authority) (*biz.Authority, error) {
	return g, nil
}

func (r *AuthorityRepo) FindByID(context.Context, int64) (*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListByName(context.Context, string) ([]*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListAll(context.Context) ([]*biz.Authority, error) {
	return nil, nil
}
