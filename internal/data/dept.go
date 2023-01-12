package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
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

func (r *DeptRepo) Save(ctx context.Context, g *biz.Dept) (*biz.Dept, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *DeptRepo) Update(ctx context.Context, g *biz.Dept) (*biz.Dept, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *DeptRepo) FindByID(ctx context.Context, id uint) (*biz.Dept, error) {
	domain := SysDept{}
	result := r.data.DB(ctx).Last(&domain, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DeptRepo) FindByName(ctx context.Context, s string) (*biz.Dept, error) {
	domain := SysDept{}
	result := r.data.DB(ctx).Last(&domain, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DeptRepo) ListByIDs(ctx context.Context, id ...uint) (domains []*biz.Dept, err error) {
	db := r.data.DB(ctx).Model(&SysDept{})
	sysDepts := []*SysDept{}

	err = db.Find(&sysDepts).Error
	if err != nil {
		return domains, err
	}
	for _, v := range sysDepts {
		domains = append(domains, r.toBiz(v))
	}
	return
}

func (r *DeptRepo) ListByName(ctx context.Context, name string) ([]*biz.Dept, error) {
	sysDepts, bizDepts := []*SysDept{}, []*biz.Dept{}
	result := r.data.DB(ctx).Find(&sysDepts, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDepts {
		bizDepts = append(bizDepts, r.toBiz(v))
	}
	return bizDepts, nil
}

func (r *DeptRepo) Delete(ctx context.Context, g *biz.Dept) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *DeptRepo) ListAll(ctx context.Context) ([]*biz.Dept, error) {
	return nil, nil
}

func (r *DeptRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (domains []*biz.Dept, total int64) {
	db := r.data.DB(ctx).Model(&SysDept{})
	sysDepts := []*SysDept{}
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

	result := db.Limit(int(handler.GetPageSize())).Find(&sysDepts)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDepts {
		domains = append(domains, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(domains))
	}

	return domains, total
}
