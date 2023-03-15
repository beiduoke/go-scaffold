package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type ResourceRepo struct {
	data *Data
	log  *log.Helper
}

// NewResourceRepo .
func NewResourceRepo(logger log.Logger, data *Data) biz.ResourceRepo {
	return &ResourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ResourceRepo) toModel(d *biz.Resource) *SysResource {
	if d == nil {
		return nil
	}
	if d.Operation == "" {
		d.Operation = d.Path
	}
	return &SysResource{
		Model: Model{
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

func (r *ResourceRepo) toBiz(d *SysResource) *biz.Resource {
	if d == nil {
		return nil
	}
	return &biz.Resource{
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

func (r *ResourceRepo) Save(ctx context.Context, g *biz.Resource) (*biz.Resource, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *ResourceRepo) Update(ctx context.Context, g *biz.Resource) (*biz.Resource, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Debug().Select("*").Omit("CreatedAt").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *ResourceRepo) FindByName(ctx context.Context, s string) (*biz.Resource, error) {
	api := SysResource{}
	result := r.data.DB(ctx).Last(&api, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&api), nil
}

func (r *ResourceRepo) FindByID(ctx context.Context, id uint) (*biz.Resource, error) {
	api := SysResource{}
	result := r.data.DB(ctx).Last(&api, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&api), nil
}

func (r *ResourceRepo) ListByIDs(ctx context.Context, id ...uint) (resources []*biz.Resource, err error) {
	db := r.data.DB(ctx).Model(&SysResource{})
	sysResources := []*SysResource{}

	err = db.Find(&sysResources).Error
	if err != nil {
		return resources, err
	}
	for _, v := range sysResources {
		resources = append(resources, r.toBiz(v))
	}
	return
}

func (r *ResourceRepo) ListByName(ctx context.Context, name string) ([]*biz.Resource, error) {
	sysResources, bizResources := []*SysResource{}, []*biz.Resource{}
	result := r.data.DB(ctx).Find(&sysResources, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysResources {
		bizResources = append(bizResources, r.toBiz(v))
	}
	return bizResources, nil
}

func (r *ResourceRepo) Delete(ctx context.Context, g *biz.Resource) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *ResourceRepo) ListAll(ctx context.Context) ([]*biz.Resource, error) {
	return nil, nil
}

func (r *ResourceRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (resources []*biz.Resource, total int64) {
	db := r.data.DB(ctx).Model(&SysResource{}).Debug()
	sysResources := []*SysResource{}
	// 查询条件
	for k, v := range paging.Query {
		db = db.Where(k, v)
	}
	// 排序
	for k, v := range paging.OrderBy {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: k}, Desc: v})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysResources)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysResources {
		resources = append(resources, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(resources))
	}

	return resources, total
}

func (r *ResourceRepo) ListAllGroup(ctx context.Context) ([]string, error) {
	db := r.data.DB(ctx).Model(&SysResource{})
	sysGroups := []string{}
	result := db.Group("group").Pluck("group", &sysGroups)
	return sysGroups, result.Error
}
