package data

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/internal/biz"
	"github.com/beiduoke/go-scaffold-single/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DictRepo struct {
	data *Data
	log  *log.Helper
}

// NewDictRepo .
func NewDictRepo(logger log.Logger, data *Data) biz.DictRepo {
	return &DictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DictRepo) toModel(d *biz.Dict) *SysDict {
	if d == nil {
		return nil
	}
	sysData := &SysDict{
		Type:    d.Type,
		Name:    d.Name,
		State:   d.State,
		Remarks: d.Remarks,
		Sort:    d.Sort,
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func (r *DictRepo) toBiz(d *SysDict) *biz.Dict {
	if d == nil {
		return nil
	}
	return &biz.Dict{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Type:      d.Type,
		Name:      d.Name,
		Sort:      d.Sort,
		State:     d.State,
		Remarks:   d.Remarks,
	}
}

func (r *DictRepo) Save(ctx context.Context, g *biz.Dict) (*biz.Dict, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *DictRepo) Update(ctx context.Context, g *biz.Dict) (*biz.Dict, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Omit("Type").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *DictRepo) FindByID(ctx context.Context, id uint) (*biz.Dict, error) {
	dict := SysDict{}
	result := r.data.DB(ctx).Last(&dict, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&dict), nil
}

func (r *DictRepo) FindByName(ctx context.Context, s string) (*biz.Dict, error) {
	dict := SysDict{}
	result := r.data.DB(ctx).Last(&dict, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&dict), nil
}

func (r *DictRepo) FindByType(ctx context.Context, tp string) (*biz.Dict, error) {
	sysDict := &SysDict{}
	result := r.data.DB(ctx).Last(sysDict, "type", tp)
	return r.toBiz(sysDict), result.Error
}

func (r *DictRepo) ListByIDs(ctx context.Context, id ...uint) (dicts []*biz.Dict, err error) {
	db := r.data.DB(ctx).Model(&SysDict{})
	sysDicts := []*SysDict{}

	err = db.Find(&sysDicts).Error
	if err != nil {
		return dicts, err
	}
	for _, v := range sysDicts {
		dicts = append(dicts, r.toBiz(v))
	}
	return
}

func (r *DictRepo) ListByName(ctx context.Context, name string) ([]*biz.Dict, error) {
	sysDicts, bizDicts := []*SysDict{}, []*biz.Dict{}
	result := r.data.DB(ctx).Find(&sysDicts, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDicts {
		bizDicts = append(bizDicts, r.toBiz(v))
	}
	return bizDicts, nil
}

func (r *DictRepo) Delete(ctx context.Context, g *biz.Dict) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *DictRepo) ListAll(ctx context.Context) ([]*biz.Dict, error) {
	return nil, nil
}

func (r *DictRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (dicts []*biz.Dict, total int64) {
	db := r.data.DB(ctx).Model(&SysDict{})
	sysDicts := []*SysDict{}

	// 查询条件
	if name, ok := paging.Query["name"].(string); ok {
		db = db.Where("name LIKE ?", name+"%")
	}
	if t, ok := paging.Query["type"].(string); ok {
		db = db.Where("type LIKE ?", t+"%")
	}
	if ids, ok := paging.Query["ids"]; ok {
		db = db.Where("id", ids)
	}

	// 排序

	if sortBy, ok := paging.OrderBy["sort"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "sort"}, Desc: sortBy})
	}

	if orderBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: orderBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysDicts)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDicts {
		dicts = append(dicts, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(dicts))
	}

	return dicts, total
}
