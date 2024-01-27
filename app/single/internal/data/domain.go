package data

import (
	"context"

	"github.com/beiduoke/go-scaffold-single/internal/biz"
	"github.com/beiduoke/go-scaffold-single/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DomainRepo struct {
	data *Data
	log  *log.Helper
	menu MenuRepo
}

// NewDomainRepo .
func NewDomainRepo(logger log.Logger, data *Data, menuRepo biz.MenuRepo) biz.DomainRepo {
	return &DomainRepo{
		data: data,
		log:  log.NewHelper(logger),
		menu: *(menuRepo.(*MenuRepo)),
	}
}

func toDomainModel(d *biz.Domain) *SysDomain {
	if d == nil {
		return nil
	}
	sysData := &SysDomain{
		Code:        d.Code,
		Name:        d.Name,
		State:       d.State,
		ParentID:    d.ParentID,
		Sort:        d.Sort,
		Alias:       d.Alias,
		Keywords:    d.Keywords,
		Logo:        d.Logo,
		Pic:         d.Pic,
		Description: d.Description,
		Remarks:     d.Remarks,
		PackageID:   d.PackageID,
		Package:     toDomainPackageModel(d.Package),
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func toDomainBiz(d *SysDomain) *biz.Domain {
	if d == nil {
		return nil
	}
	return &biz.Domain{
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		ID:          d.ID,
		Code:        d.Code,
		Name:        d.Name,
		Sort:        d.Sort,
		State:       d.State,
		ParentID:    d.ParentID,
		Alias:       d.Alias,
		Keywords:    d.Keywords,
		Logo:        d.Logo,
		Pic:         d.Pic,
		Description: d.Description,
		Remarks:     d.Remarks,
		PackageID:   d.PackageID,
		Package:     toDomainPackageBiz(d.Package),
	}
}

func (r *DomainRepo) Save(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	d := toDomainModel(g)
	sfId := r.data.sf.Generate()
	// g.DomainID = base64.StdEncoding.EncodeToString([]byte(id.String()))
	d.Code = sfId.String()
	result := r.data.DB(ctx).Create(d)
	return toDomainBiz(d), result.Error
}

func (r *DomainRepo) Update(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	d := toDomainModel(g)
	result := r.data.DB(ctx).Model(d).Omit("Code").Updates(d)
	return toDomainBiz(d), result.Error
}

func (r *DomainRepo) FindByID(ctx context.Context, id uint) (*biz.Domain, error) {
	domain := SysDomain{}
	result := r.data.DB(ctx).Preload("Package").Last(&domain, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDomainBiz(&domain), nil
}

func (r *DomainRepo) FindByName(ctx context.Context, s string) (*biz.Domain, error) {
	domain := SysDomain{}
	result := r.data.DB(ctx).Preload("Package").Last(&domain, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDomainBiz(&domain), nil
}

func (r *DomainRepo) FindByCode(ctx context.Context, code string) (*biz.Domain, error) {
	sysDomain := &SysDomain{}
	result := r.data.DB(ctx).Preload("Package").Last(sysDomain, "code", code)
	return toDomainBiz(sysDomain), result.Error
}

func (r *DomainRepo) ListByIDs(ctx context.Context, id ...uint) (domains []*biz.Domain, err error) {
	db := r.data.DB(ctx).Model(&SysDomain{})
	sysDomains := []*SysDomain{}

	err = db.Preload("Package").Find(&sysDomains).Error
	if err != nil {
		return domains, err
	}
	for _, v := range sysDomains {
		domains = append(domains, toDomainBiz(v))
	}
	return
}

func (r *DomainRepo) ListByName(ctx context.Context, name string) ([]*biz.Domain, error) {
	sysDomains, bizDomains := []*SysDomain{}, []*biz.Domain{}
	result := r.data.DB(ctx).Preload("Package").Find(&sysDomains, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDomains {
		bizDomains = append(bizDomains, toDomainBiz(v))
	}
	return bizDomains, nil
}

func (r *DomainRepo) Delete(ctx context.Context, g *biz.Domain) error {
	return r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Delete(toDomainModel(g))
		if err := result.Error; err != nil {
			return err
		}
		_, err := r.data.enforcer.DeleteDomains(g.GetID())
		return err
	})
}

func (r *DomainRepo) ListAll(ctx context.Context) ([]*biz.Domain, error) {
	return nil, nil
}

func (r *DomainRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (domains []*biz.Domain, total int64) {
	db := r.data.DB(ctx).Model(&SysDomain{}).Debug()
	sysDomains := []*SysDomain{}
	// 查询条件
	if name, ok := paging.Query["name"].(string); ok {
		db = db.Where("name LIKE ?", name+"%")
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

	result := db.Limit(int(paging.PageSize)).Preload("Package").Find(&sysDomains)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDomains {
		domains = append(domains, toDomainBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(domains))
	}

	return domains, total
}

func (r *DomainRepo) HandleMenu(ctx context.Context, g *biz.Domain) error {
	sysDomain, sysMenus := SysDomain{}, make([]*SysMenu, 0)
	for _, v := range g.Menus {
		sysMenus = append(sysMenus, toMenuModel(v))
	}
	sysDomain.ID = g.ID
	return r.data.DB(ctx).Model(&sysDomain).Association("Menus").Replace(sysMenus)
}

func (r *DomainRepo) ListMenuIDByIDs(ctx context.Context, ids ...uint) []uint {
	var domainMenuIds []uint
	result := r.data.DB(ctx).Table("sys_domain_menus").Where("sys_domain_id", ids).Pluck("sys_menu_id", &domainMenuIds)
	err := result.Error
	if err != nil {
		r.log.Error(err)
		return nil
	}
	return domainMenuIds
}

// 获取指定权限菜单列表
func (r *DomainRepo) ListMenuByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	domainMenuIds := r.ListMenuIDByIDs(ctx, ids...)
	bizAllMenus, err := r.menu.ListAll(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	bizMenus := make([]*biz.Menu, 0, len(domainMenuIds))
	for _, menu := range bizAllMenus {
		for _, menuID := range domainMenuIds {
			if menuID == menu.ID {
				bizMenus = append(bizMenus, menu)
				continue
			}
		}
	}
	return bizMenus, nil
}

// ListMenuAndParentByIDs
func (r *DomainRepo) ListMenuAndParentByIDs(ctx context.Context, ids ...uint) ([]*biz.Menu, error) {
	bizAllMenus, _ := r.menu.ListAll(ctx)
	return menuRecursiveParent(bizAllMenus, r.ListMenuIDByIDs(ctx, ids...)...), nil
}
