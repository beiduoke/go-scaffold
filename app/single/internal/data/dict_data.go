package data

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/internal/biz"
	"github.com/beiduoke/go-scaffold-single/pkg/util/pagination"
	"gorm.io/gorm/clause"
)

func toDictDataModel(d *biz.DictData) *SysDictData {
	if d == nil {
		return nil
	}
	sysData := &SysDictData{
		Label:     d.Label,
		Value:     d.Value,
		ColorType: d.ColorType,
		CssClass:  d.CssClass,
		DictType:  d.DictType,
		State:     d.State,
		Remarks:   d.Remarks,
		Sort:      d.Sort,
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func toDictDataBiz(d *SysDictData) *biz.DictData {
	if d == nil {
		return nil
	}
	return &biz.DictData{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Label:     d.Label,
		Value:     d.Value,
		ColorType: d.ColorType,
		CssClass:  d.CssClass,
		DictType:  d.DictType,
		Remarks:   d.Remarks,
		Sort:      d.Sort,
		State:     d.State,
	}
}

func (r *DictRepo) DataSave(ctx context.Context, g *biz.DictData) (*biz.DictData, error) {
	sysDictData := toDictDataModel(g)
	result := r.data.DB(ctx).Create(sysDictData)
	return toDictDataBiz(sysDictData), result.Error
}

func (r *DictRepo) DataUpdate(ctx context.Context, g *biz.DictData) (*biz.DictData, error) {
	sysDictData := toDictDataModel(g)
	result := r.data.DB(ctx).Model(sysDictData).Omit("DictType").Updates(sysDictData)
	return toDictDataBiz(sysDictData), result.Error
}

func (r *DictRepo) DataFindByID(ctx context.Context, id uint) (*biz.DictData, error) {
	sysDictData := SysDictData{}
	result := r.data.DB(ctx).Last(&sysDictData, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDictDataBiz(&sysDictData), nil
}

func (r *DictRepo) DataFindByLabel(ctx context.Context, label string) (*biz.DictData, error) {
	sysDictData := SysDictData{}
	result := r.data.DB(ctx).Last(&sysDictData, "label = ?", label)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDictDataBiz(&sysDictData), nil
}

func (r *DictRepo) DataFindByDictType(ctx context.Context, dictType string) (*biz.DictData, error) {
	sysDictData := &SysDictData{}
	result := r.data.DB(ctx).Last(sysDictData, "dict_type", dictType)
	return toDictDataBiz(sysDictData), result.Error
}

func (r *DictRepo) DataListByIDs(ctx context.Context, id ...uint) (dicts []*biz.DictData, err error) {
	db := r.data.DB(ctx).Model(&SysDictData{})
	sysDictDatas := []*SysDictData{}

	err = db.Find(&sysDictDatas).Error
	if err != nil {
		return dicts, err
	}
	for _, v := range sysDictDatas {
		dicts = append(dicts, toDictDataBiz(v))
	}
	return
}

func (r *DictRepo) DataListByLabel(ctx context.Context, label string) ([]*biz.DictData, error) {
	sysDictDatas, bizDictDatas := []*SysDictData{}, []*biz.DictData{}
	result := r.data.DB(ctx).Find(&sysDictDatas, "label LIKE ?", "%"+label)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDictDatas {
		bizDictDatas = append(bizDictDatas, toDictDataBiz(v))
	}
	return bizDictDatas, nil
}

func (r *DictRepo) DataDelete(ctx context.Context, g *biz.DictData) error {
	return r.data.DB(ctx).Delete(toDictDataModel(g)).Error
}

func (r *DictRepo) DataListAll(ctx context.Context) ([]*biz.DictData, error) {
	return nil, nil
}

func (r *DictRepo) DataListPage(ctx context.Context, paging *pagination.Pagination) (dicts []*biz.DictData, total int64) {
	db := r.data.DB(ctx).Model(&SysDictData{})
	sysDictDatas := []*SysDictData{}

	// 查询条件
	if label, ok := paging.Query["label"].(string); ok {
		db = db.Where("label LIKE ?", label+"%")
	}
	if dictType, ok := paging.Query["dictType"].(string); ok && dictType != "" {
		db = db.Where("dict_type", dictType)
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

	result := db.Limit(int(paging.PageSize)).Find(&sysDictDatas)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDictDatas {
		dicts = append(dicts, toDictDataBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(dicts))
	}

	return dicts, total
}
