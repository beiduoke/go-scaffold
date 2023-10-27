package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
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
		sysDomainPackage := *d
		sysMenus := make([]*SysMenu, 0)
		for _, v := range g.Menus {
			sysMenus = append(sysMenus, toMenuModel(v))
		}
		err := r.data.DB(ctx).Model(&sysDomainPackage).Association("Menus").Replace(sysMenus)
		if err != nil {
			return err
		}

		return r.data.DB(ctx).Model(&sysDomainPackage).Updates(d).Error
	})
	return toDomainPackageBiz(d), err
}
func (r *DomainRepo) PackageFindByID(ctx context.Context, id uint) (*biz.DomainPackage, error) {
	return nil, nil
}
func (r *DomainRepo) PackageDelete(ctx context.Context, g *biz.DomainPackage) error {
	return nil
}
func (r *DomainRepo) PackageListAll(ctx context.Context) ([]*biz.DomainPackage, error) {
	return nil, nil
}
func (r *DomainRepo) PackageListPage(ctx context.Context, paging *pagination.Pagination) ([]*biz.DomainPackage, int64) {
	return nil, 0
}
