package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApiRepo struct {
	data *Data
	log  *log.Helper
}

// NewApiRepo .
func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &ApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ApiRepo) toModel(d *biz.Api) *SysApi {
	if d == nil {
		return nil
	}

	return &SysApi{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:        d.Name,
		Path:        d.Path,
		Method:      d.Method,
		Group:       d.Group,
		Description: d.Description,
	}
}

func (r *ApiRepo) toBiz(d *SysApi) *biz.Api {
	if d == nil {
		return nil
	}
	return &biz.Api{
		ID:          d.ID,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		Name:        d.Name,
		Path:        d.Path,
		Method:      d.Method,
		Group:       d.Group,
		Description: d.Description,
	}
}

func (r *ApiRepo) Save(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	return g, nil
}

func (r *ApiRepo) Update(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	return g, nil
}

func (r *ApiRepo) FindByID(ctx context.Context, id int64) (*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListByName(ctx context.Context, name string) ([]*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListAll(ctx context.Context) ([]*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (apis []*biz.Api, total int64) {
	db := r.data.DB(ctx).Model(&SysApi{})
	sysApi := []*SysApi{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}
	result := db.Count(&total).Offset(handler.GetPageOffset()).Limit(int(handler.GetPageSize())).Find(&sysApi)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysApi {
		apis = append(apis, r.toBiz(v))
	}
	return apis, total
}
