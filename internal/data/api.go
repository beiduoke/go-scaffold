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
func NewApiRepo(logger log.Logger, data *Data) biz.ApiRepo {
	return &ApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ApiRepo) toModel(d *biz.Api) *SysApi {
	if d == nil {
		return nil
	}
	if d.Operation == "" {
		d.Operation = d.Path
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
		Operation:   d.Operation,
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
		Operation:   d.Operation,
		Group:       d.Group,
		Description: d.Description,
	}
}

func (r *ApiRepo) Save(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *ApiRepo) Update(ctx context.Context, g *biz.Api) (*biz.Api, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *ApiRepo) FindByName(ctx context.Context, s string) (*biz.Api, error) {
	api := SysApi{}
	result := r.data.DB(ctx).Last(&api, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&api), nil
}

func (r *ApiRepo) FindByID(ctx context.Context, id uint) (*biz.Api, error) {
	api := SysApi{}
	result := r.data.DB(ctx).Last(&api, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&api), nil
}

func (r *ApiRepo) ListByIDs(ctx context.Context, id ...uint) (apis []*biz.Api, err error) {
	db := r.data.DB(ctx).Model(&SysApi{})
	sysApis := []*SysApi{}

	err = db.Find(&sysApis).Error
	if err != nil {
		return apis, err
	}
	for _, v := range sysApis {
		apis = append(apis, r.toBiz(v))
	}
	return
}

func (r *ApiRepo) ListByName(ctx context.Context, name string) ([]*biz.Api, error) {
	sysApis, bizApis := []*SysApi{}, []*biz.Api{}
	result := r.data.DB(ctx).Find(&sysApis, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysApis {
		bizApis = append(bizApis, r.toBiz(v))
	}
	return bizApis, nil
}

func (r *ApiRepo) Delete(ctx context.Context, g *biz.Api) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *ApiRepo) ListAll(ctx context.Context) ([]*biz.Api, error) {
	return nil, nil
}

func (r *ApiRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (apis []*biz.Api, total int64) {
	db := r.data.DB(ctx).Model(&SysApi{})
	sysApis := []*SysApi{}
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

	result := db.Limit(int(handler.GetPageSize())).Find(&sysApis)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysApis {
		apis = append(apis, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(apis))
	}

	return apis, total
}
