package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"gorm.io/gorm/clause"
)

func toDomainPackageModel(d *biz.DomainPackage) *SysDomainPackage {
	if d == nil {
		return nil
	}
	sysDomainPackage := &SysDomainPackage{
		Name:    d.Name,
		Sort:    d.Sort,
		State:   d.State,
		Remarks: d.Remarks,
		Menus:   []SysMenu{},
	}
	for _, v := range d.Menus {
		sysDomainPackage.Menus = append(sysDomainPackage.Menus, *toMenuModel(v))
	}
	sysDomainPackage.ID = d.ID
	sysDomainPackage.CreatedAt = d.CreatedAt
	sysDomainPackage.CreatedAt = d.UpdatedAt
	return sysDomainPackage
}

func toDomainPackageBiz(d *SysDomainPackage) *biz.DomainPackage {
	if d == nil {
		return nil
	}
	bizDomainPackage := &biz.DomainPackage{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Name:      d.Name,
		Sort:      d.Sort,
		State:     d.State,
		Remarks:   d.Remarks,
	}

	for _, v := range d.Menus {
		bizDomainPackage.Menus = append(bizDomainPackage.Menus, toMenuBiz(&v))
	}

	return bizDomainPackage
}

func (r *DomainRepo) PackageSave(ctx context.Context, g *biz.DomainPackage) (*biz.DomainPackage, error) {
	d := toDomainPackageModel(g)
	result := r.data.DB(ctx).Create(d)
	return toDomainPackageBiz(d), result.Error
}

func (r *DomainRepo) PackageUpdate(ctx context.Context, g *biz.DomainPackage) (*biz.DomainPackage, error) {
	d := toDomainPackageModel(g)
	err := r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Model(d).Omit(clause.Associations).Updates(d)
		if err := result.Error; err != nil {
			return err
		}

		return r.data.DB(ctx).Model(d).Association("Menus").Replace(d.Menus)
	})
	return toDomainPackageBiz(d), err
}

func (r *DomainRepo) PackageFindByID(ctx context.Context, id uint) (*biz.DomainPackage, error) {
	dp := SysDomainPackage{}
	result := r.data.DB(ctx).Last(&dp, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDomainPackageBiz(&dp), nil
}

func (r *DomainRepo) PackageFindByName(ctx context.Context, s string) (*biz.DomainPackage, error) {
	dp := SysDomainPackage{}
	result := r.data.DB(ctx).Last(&dp, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDomainPackageBiz(&dp), nil
}

func (r *DomainRepo) PackageListByIDs(ctx context.Context, id ...uint) (dps []*biz.DomainPackage, err error) {
	db := r.data.DB(ctx).Model(&SysDomainPackage{})
	sysDomainPackages := []*SysDomainPackage{}

	err = db.Find(&sysDomainPackages).Error
	if err != nil {
		return dps, err
	}
	for _, v := range sysDomainPackages {
		dps = append(dps, toDomainPackageBiz(v))
	}
	return
}

func (r *DomainRepo) PackageListByName(ctx context.Context, name string) ([]*biz.DomainPackage, error) {
	sysDomainPackages, bizDomainPackages := []*SysDomainPackage{}, []*biz.DomainPackage{}
	result := r.data.DB(ctx).Find(&sysDomainPackages, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDomainPackages {
		bizDomainPackages = append(bizDomainPackages, toDomainPackageBiz(v))
	}
	return bizDomainPackages, nil
}

func (r *DomainRepo) PackageDelete(ctx context.Context, g *biz.DomainPackage) error {
	return r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Delete(toDomainPackageModel(g))
		if err := result.Error; err != nil {
			return err
		}
		_, err := r.data.enforcer.DeleteDomains(g.GetID())
		return err
	})
}

func (r *DomainRepo) PackageListAll(ctx context.Context) ([]*biz.DomainPackage, error) {
	return nil, nil
}

func (r *DomainRepo) PackageListPage(ctx context.Context, paging *pagination.Pagination) (dps []*biz.DomainPackage, total int64) {
	db := r.data.DB(ctx).Model(&SysDomainPackage{})
	sysDomainPackages := []*SysDomainPackage{}
	// 查询条件
	if name, ok := paging.Query["name"].(string); ok && name != "" {
		db = db.Where("name LIKE ?", name+"%")
	}
	if ids, ok := paging.Query["ids"].([]uint); ok && len(ids) > 0 {
		db = db.Where("id", ids)
	}
	if state, ok := paging.Query["state"].(uint); ok && state > 0 {
		db = db.Where("state", state)
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

	result := db.Limit(int(paging.PageSize)).Preload("Menus").Find(&sysDomainPackages)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDomainPackages {
		dps = append(dps, toDomainPackageBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(dps))
	}

	return dps, total
}
