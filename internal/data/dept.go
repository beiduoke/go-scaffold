package data

import (
	"context"
	"strings"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DeptRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeptRepo .
func NewDeptRepo(logger log.Logger, data *Data) biz.DeptRepo {
	return &DeptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DeptRepo) toModel(d *biz.Dept) *SysDept {
	if d == nil {
		return nil
	}
	sysData := &SysDept{
		Name:     d.Name,
		ParentID: d.ParentID,
		Sort:     d.Sort,
		State:    d.State,
		Remarks:  d.Remarks,
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt

	return sysData
}

func (r *DeptRepo) toBiz(d *SysDept) *biz.Dept {
	if d == nil {
		return nil
	}
	return &biz.Dept{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Name:      d.Name,
		ParentID:  d.ParentID,
		Sort:      d.Sort,
		State:     d.State,
		Remarks:   d.Remarks,
	}
}

func (r *DeptRepo) LinkedAncestors(list []*biz.Dept, pid uint) (depts []*biz.Dept) {
	if pid == 0 {
		return append(depts, &biz.Dept{})
	}
	for _, v := range list {
		if v.ID == pid {
			depts = append(depts, append(r.LinkedAncestors(list, v.ParentID), v)...)
			break
		}
	}
	return
}

func (r *DeptRepo) FilterAncestors(ctx context.Context, g *biz.Dept) string {
	deptAlls, _ := r.ListAll(ctx)
	deptIdsStr := make([]string, 0, len(deptAlls))
	for _, v := range r.LinkedAncestors(deptAlls, g.ParentID) {
		deptIdsStr = append(deptIdsStr, convert.UnitToString(v.ID))
	}
	return strings.Join(deptIdsStr, ",")
}

func (r *DeptRepo) Save(ctx context.Context, g *biz.Dept) (*biz.Dept, error) {
	d := r.toModel(g)
	d.Ancestors = r.FilterAncestors(ctx, g)
	d.DomainID = r.data.DomainID(ctx)
	result := r.data.DB(ctx).Debug().Create(d)
	return r.toBiz(d), result.Error
}

func (r *DeptRepo) Update(ctx context.Context, g *biz.Dept) (*biz.Dept, error) {
	d := r.toModel(g)
	d.Ancestors = r.FilterAncestors(ctx, g)
	result := r.data.DBD(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *DeptRepo) FindByID(ctx context.Context, id uint) (*biz.Dept, error) {
	dept := SysDept{}
	result := r.data.DBD(ctx).Last(&dept, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&dept), nil
}

func (r *DeptRepo) FindByName(ctx context.Context, s string) (*biz.Dept, error) {
	dept := SysDept{}
	result := r.data.DBD(ctx).Last(&dept, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&dept), nil
}

func (r *DeptRepo) ListByIDs(ctx context.Context, id ...uint) (depts []*biz.Dept, err error) {
	db := r.data.DBD(ctx).Model(&SysDept{})
	sysDepts := []*SysDept{}

	err = db.Find(&sysDepts).Error
	if err != nil {
		return depts, err
	}
	for _, v := range sysDepts {
		depts = append(depts, r.toBiz(v))
	}
	return
}

func (r *DeptRepo) ListByName(ctx context.Context, name string) ([]*biz.Dept, error) {
	sysDepts, bizDepts := []*SysDept{}, []*biz.Dept{}
	result := r.data.DBD(ctx).Find(&sysDepts, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDepts {
		bizDepts = append(bizDepts, r.toBiz(v))
	}
	return bizDepts, nil
}

func (r *DeptRepo) Delete(ctx context.Context, g *biz.Dept) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
}

func (r *DeptRepo) ListAll(ctx context.Context) ([]*biz.Dept, error) {
	db := r.data.DBD(ctx).Model(&SysDept{})
	sysDepts, bizDepts := []*SysDept{}, []*biz.Dept{}

	err := db.Find(&sysDepts).Error
	for _, v := range sysDepts {
		bizDepts = append(bizDepts, r.toBiz(v))
	}
	return bizDepts, err
}

func (r *DeptRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (depts []*biz.Dept, total int64) {
	db := r.data.DBD(ctx).Model(&SysDept{}).Debug()
	sysDepts := []*SysDept{}
	// 查询条件
	if paging.Query != nil {
		if name, ok := paging.Query["name"].(string); ok {
			db = db.Where("name LIKE ?", name+"%")
		}
	}

	// 排序
	if paging.OrderBy != nil {
		if orderBy, ok := paging.OrderBy["createdAt"]; ok {
			db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: orderBy})
		}

		if idBy, ok := paging.OrderBy["id"]; ok {
			db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
		}
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysDepts)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDepts {
		depts = append(depts, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(depts))
	}

	return depts, total
}
