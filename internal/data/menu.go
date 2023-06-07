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
		Parameters:  make([]SysMenuParameter, 0, len(d.Parameters)),
		Buttons:     make([]SysMenuButton, 0, len(d.Buttons)),
		ApiResource: d.ApiResource,
	}
	for _, v := range d.Parameters {
		sysData.Parameters = append(sysData.Parameters, SysMenuParameter{
			Type:  v.Type,
			Name:  v.Name,
			Value: v.Value,
		})
	}
	for _, v := range d.Buttons {
		sysData.Buttons = append(sysData.Buttons, SysMenuButton{
			Name:    v.Name,
			Remarks: v.Remarks,
		})
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
		Parameters:  make([]*biz.MenuParameter, 0, len(d.Parameters)),
		Buttons:     make([]*biz.MenuButton, 0, len(d.Buttons)),
		ApiResource: d.ApiResource,
	}
	for _, v := range d.Parameters {
		data.Parameters = append(data.Parameters, &biz.MenuParameter{
			Type:  v.Type,
			Name:  v.Name,
			Value: v.Value,
		})
	}
	for _, v := range d.Buttons {
		data.Buttons = append(data.Buttons, &biz.MenuButton{
			Name:    v.Name,
			Remarks: v.Remarks,
		})
	}
	return data
}

func (r *MenuRepo) Save(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	if result.Error == nil {
		result.Error = r.SetCache(ctx, d)
	}
	return r.toBiz(d), result.Error
}

func (r *MenuRepo) Update(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(&SysMenu{}).Not(g.ID).Where("name", g.Name).Pluck("id", nil)
	if result.RowsAffected > 0 {
		return nil, errors.New("duplicate name")
	}
	// 一对多关联，删除原始按钮
	if err := r.data.DB(ctx).Model(&SysMenuButton{}).Delete(&SysMenuButton{}, "menu_id", g.ID).Error; err != nil {
		return nil, err
	}
	// 一对多关联，删除原始参数
	if err := r.data.DB(ctx).Model(&SysMenuParameter{}).Delete(&SysMenuParameter{}, "menu_id", g.ID).Error; err != nil {
		return nil, err
	}
	result = r.data.DB(ctx).Model(d).Omit("CreatedAt").Updates(d)
	if result.Error == nil {
		result.Error = r.SetCache(ctx, d)
	}
	return r.toBiz(d), result.Error
}

func (r *MenuRepo) FindByName(ctx context.Context, s string) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DB(ctx).Preload("Parameters").Preload("Buttons").Last(&menu, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&menu), nil
}

func (r *MenuRepo) FindByID(ctx context.Context, id uint) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DB(ctx).Preload("Parameters").Preload("Buttons").Last(&menu, id)
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
	result := r.data.DB(ctx).Select("Parameters", "Buttons").Delete(r.toModel(g))
	if err := result.Error; err != nil {
		return err
	}
	return r.DeleteCache(ctx, g.GetID())
}

func (r *MenuRepo) ListAll(ctx context.Context) (menus []*biz.Menu, err error) {
	for _, v := range r.ListAllCache(ctx) {
		menus = append(menus, r.toBiz(v))
	}
	return
}

func (r *MenuRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (menus []*biz.Menu, total int64) {
	db := r.data.DB(ctx).Model(&SysMenu{}).Debug()
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
		for _, id := range ids {
			if _, o := mid[v.ID]; v.ID == id && !o {
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
