package data

import (
	"context"
	"errors"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper
}

// NewMenuRepo .
func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &MenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MenuRepo) toModel(d *biz.Menu) *SysMenu {
	if d == nil {
		return nil
	}
	sysData := &SysMenu{
		Name:        d.Name,
		Type:        d.Type,
		ParentID:    d.ParentID,
		Path:        d.Path,
		Component:   d.Component,
		Permission:  d.Permission,
		Sort:        d.Sort,
		Icon:        d.Icon,
		Title:       d.Title,
		IsHidden:    d.IsHidden,
		IsCache:     d.IsCache,
		IsAffix:     d.IsAffix,
		LinkType:    d.LinkType,
		LinkUrl:     d.LinkUrl,
		ApiResource: d.ApiResource,
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func (r *MenuRepo) toBiz(d *SysMenu) *biz.Menu {
	if d == nil {
		return nil
	}
	data := &biz.Menu{
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		ID:          d.ID,
		Name:        d.Name,
		Type:        d.Type,
		ParentID:    d.ParentID,
		Path:        d.Path,
		Component:   d.Component,
		Permission:  d.Permission,
		Sort:        d.Sort,
		Icon:        d.Icon,
		Title:       d.Title,
		IsHidden:    d.IsHidden,
		IsCache:     d.IsCache,
		IsAffix:     d.IsAffix,
		LinkType:    d.LinkType,
		LinkUrl:     d.LinkUrl,
		ApiResource: d.ApiResource,
	}
	return data
}

func (r *MenuRepo) Save(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	bizMenu := r.toBiz(d)
	if result.Error == nil {
		result.Error = r.SetCache(ctx, bizMenu)
	}
	return bizMenu, result.Error
}

func (r *MenuRepo) Update(ctx context.Context, g *biz.Menu) (bizMenu *biz.Menu, err error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(&SysMenu{}).Not(g.ID).Where("name", g.Name).Pluck("id", nil)
	if result.RowsAffected > 0 {
		return nil, errors.New("duplicate name")
	}

	// 判断是否更新角色策略
	beforeMenu := SysMenu{}
	result = r.data.DB(ctx).Select("ApiResource").Last(&beforeMenu, g.ID)
	if result.RowsAffected > 0 && beforeMenu.ApiResource != d.ApiResource {
		resultErr := r.data.RoleUpdatePolicyResource(ctx, beforeMenu.ApiResource, d.ApiResource)
		if resultErr != nil {
			r.log.Error(resultErr)
		}
	}

	err = r.data.InTx(ctx, func(ctx context.Context) error {
		result := r.data.DB(ctx).Model(d).Omit("CreatedAt").Updates(d)
		if result.Error != nil {
			return err
		}
		bizMenu := r.toBiz(d)
		return r.SetCache(ctx, bizMenu)
	})

	return bizMenu, err
}

func (r *MenuRepo) FindByName(ctx context.Context, s string) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DB(ctx).Last(&menu, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&menu), nil
}

func (r *MenuRepo) FindByID(ctx context.Context, id uint) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DB(ctx).Last(&menu, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&menu), nil
}

func (r *MenuRepo) ListByIDs(ctx context.Context, id ...uint) (menus []*biz.Menu, err error) {
	db := r.data.DB(ctx).Model(&SysMenu{})
	sysMenus := []*SysMenu{}

	err = db.Find(&sysMenus, id).Error
	if err != nil {
		return menus, err
	}
	for _, v := range sysMenus {
		menus = append(menus, r.toBiz(v))
	}
	return
}

func (r *MenuRepo) ListByName(ctx context.Context, name string) ([]*biz.Menu, error) {
	sysMenus, bizMenus := []*SysMenu{}, []*biz.Menu{}
	result := r.data.DB(ctx).Find(&sysMenus, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysMenus {
		bizMenus = append(bizMenus, r.toBiz(v))
	}
	return bizMenus, nil
}

func (r *MenuRepo) Delete(ctx context.Context, g *biz.Menu) error {
	result := r.data.DB(ctx).Delete(r.toModel(g))
	if err := result.Error; err != nil {
		return err
	}
	// 判断是否删除角色策略
	sysMenu := SysMenu{}
	result = r.data.DB(ctx).Unscoped().Select("ApiResource").Last(&sysMenu, g.ID)
	if result.RowsAffected > 0 && sysMenu.ApiResource != "" {
		resultErr := r.data.RoleDeletePolicyResource(ctx, sysMenu.ApiResource)
		if resultErr != nil {
			r.log.Error(resultErr)
		}
	}
	return r.DeleteCache(ctx, g.GetID())
}

func (r *MenuRepo) ListAll(ctx context.Context) (menus []*biz.Menu, err error) {
	return r.ListAllCache(ctx), nil
}

func (r *MenuRepo) ListAllIDs(ctx context.Context) []uint {
	bizAllMenus := r.ListAllCache(ctx)
	menuIds := make([]uint, 0, len(bizAllMenus))
	for _, v := range bizAllMenus {
		menuIds = append(menuIds, v.ID)
	}
	return menuIds
}

func (r *MenuRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (menus []*biz.Menu, total int64) {
	db := r.data.DB(ctx).Model(&SysMenu{})
	sysMenus := []*SysMenu{}
	// 查询条件
	if name, ok := paging.Query["name"].(string); ok && name != "" {
		if name != "" {
			db = db.Where("name LIKE ?", name+"%")
		}
	}
	if title, ok := paging.Query["title"].(string); ok && title != "" {
		if title != "" {
			db = db.Where("title LIKE ?", title+"%")
		}
	}
	// 排序
	if sortBy, ok := paging.OrderBy["sort"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "sort"}, Desc: sortBy})
	}

	if createdBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: createdBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysMenus)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysMenus {
		menus = append(menus, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(menus))
	}

	return menus, total
}

// 根据ID递归查询父级菜单
func menuRecursiveParent(menus []*biz.Menu, ids ...uint) []*biz.Menu {
	result, mid := []*biz.Menu{}, map[uint]uint{}
	for _, v := range menus {
		for _, p := range menus {
			if v.ParentID == p.ID {
				v.Parent = p
				break
			}
		}
		_, ok := mid[v.ID]
		if ok {
			continue
		}
		for _, id := range ids {
			if v.ID == id {
				mid[v.ID] = v.ID
				result = append(result, v)
				for _, m := range menuRecursiveParent(menus, v.ParentID) {
					if _, ok := mid[m.ID]; !ok {
						mid[m.ID] = m.ID
						result = append(result, m)
					}
				}
			}
		}
	}
	return result
}
