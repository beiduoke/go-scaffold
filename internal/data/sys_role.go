package data

import (
	"context"

	"github.com/bedoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RoleRepo) Save(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (r *RoleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (r *RoleRepo) FindByID(context.Context, int64) (*biz.Role, error) {
	return nil, nil
}

func (r *RoleRepo) ListByName(context.Context, string) ([]*biz.Role, error) {
	return nil, nil
}

func (r *RoleRepo) ListAll(context.Context) ([]*biz.Role, error) {
	return nil, nil
}
