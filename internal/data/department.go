package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DepartmentRepo struct {
	data *Data
	log  *log.Helper
}

// NewDepartmentRepo .
func NewDepartmentRepo(logger log.Logger, data *Data) biz.DepartmentRepo {
	return &DepartmentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DepartmentRepo) toModel(d *biz.Department) *SysDepartment {
	if d == nil {
		return nil
	}
	sysData := &SysDepartment{
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

func (r *DepartmentRepo) toBiz(d *SysDepartment) *biz.Department {
	if d == nil {
		return nil
	}
	return &biz.Department{
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

func (r *DepartmentRepo) Save(ctx context.Context, g *biz.Department) (*biz.Department, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *DepartmentRepo) Update(ctx context.Context, g *biz.Department) (*biz.Department, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *DepartmentRepo) FindByID(ctx context.Context, id uint) (*biz.Department, error) {
	domain := SysDepartment{}
	result := r.data.DB(ctx).Last(&domain, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DepartmentRepo) FindByName(ctx context.Context, s string) (*biz.Department, error) {
	domain := SysDepartment{}
	result := r.data.DB(ctx).Last(&domain, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DepartmentRepo) ListByIDs(ctx context.Context, id ...uint) (domains []*biz.Department, err error) {
	db := r.data.DB(ctx).Model(&SysDepartment{})
	sysDepartments := []*SysDepartment{}

	err = db.Find(&sysDepartments).Error
	if err != nil {
		return domains, err
	}
	for _, v := range sysDepartments {
		domains = append(domains, r.toBiz(v))
	}
	return
}

func (r *DepartmentRepo) ListByName(ctx context.Context, name string) ([]*biz.Department, error) {
	sysDepartments, bizDepartments := []*SysDepartment{}, []*biz.Department{}
	result := r.data.DB(ctx).Find(&sysDepartments, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDepartments {
		bizDepartments = append(bizDepartments, r.toBiz(v))
	}
	return bizDepartments, nil
}

func (r *DepartmentRepo) Delete(ctx context.Context, g *biz.Department) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *DepartmentRepo) ListAll(ctx context.Context) ([]*biz.Department, error) {
	return nil, nil
}

func (r *DepartmentRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (domains []*biz.Department, total int64) {
	db := r.data.DB(ctx).Model(&SysDepartment{})
	sysDepartments := []*SysDepartment{}
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

	result := db.Limit(int(handler.GetPageSize())).Find(&sysDepartments)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDepartments {
		domains = append(domains, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(domains))
	}

	return domains, total
}
