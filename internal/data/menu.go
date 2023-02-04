package data

import (
	"context"
	"encoding/json"
	"errors"
	"sort"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
		Name:       d.Name,
		Type:       d.Type,
		ParentID:   d.ParentID,
		Path:       d.Path,
		Hidden:     d.Hidden,
		Component:  d.Component,
		Permission: d.Permission,
		Sort:       d.Sort,
		Meta: SysMeta{
			Icon:      d.Icon,
			Title:     d.Title,
			KeepAlive: d.KeepAlive,
			BaseMenu:  d.BaseMenu,
			CloseTab:  d.CloseTab,
			ExtType:   d.ExtType,
		},
		Parameters: make([]SysMenuParameter, 0, len(d.Parameters)),
		Buttons:    make([]SysMenuButton, 0, len(d.Buttons)),
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
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		ID:         d.ID,
		Name:       d.Name,
		Type:       d.Type,
		ParentID:   d.ParentID,
		Path:       d.Path,
		Hidden:     d.Hidden,
		Component:  d.Component,
		Permission: d.Permission,
		Sort:       d.Sort,
		Icon:       d.Meta.Icon,
		Title:      d.Meta.Title,
		KeepAlive:  d.Meta.KeepAlive,
		BaseMenu:   d.Meta.BaseMenu,
		CloseTab:   d.Meta.CloseTab,
		ExtType:    d.Meta.ExtType,
		Parameters: make([]*biz.MenuParameter, 0, len(d.Parameters)),
		Buttons:    make([]*biz.MenuButton, 0, len(d.Buttons)),
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
		r.setCache(ctx, d)
	}
	return r.toBiz(d), result.Error
}

func (r *MenuRepo) Update(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	d := r.toModel(g)

	// 一对多关联，删除原始按钮
	if err := r.data.DB(ctx).Model(&SysMenuButton{}).Delete(&SysMenuButton{}, "menu_id", g.ID).Error; err != nil {
		return nil, err
	}

	// 一对多关联，删除原始参数
	if err := r.data.DB(ctx).Model(&SysMenuParameter{}).Delete(&SysMenuParameter{}, "menu_id", g.ID).Error; err != nil {
		return nil, err
	}

	result := r.data.DBD(ctx).Model(d).Updates(d)

	if result.Error == nil {
		r.setCache(ctx, d)
	}
	return r.toBiz(d), result.Error
}

func (r *MenuRepo) FindByName(ctx context.Context, s string) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DBD(ctx).Preload("Parameters").Preload("Buttons").Last(&menu, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&menu), nil
}

func (r *MenuRepo) FindByID(ctx context.Context, id uint) (*biz.Menu, error) {
	menu := SysMenu{}
	result := r.data.DBD(ctx).Preload("Parameters").Preload("Buttons").Last(&menu, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&menu), nil
}

func (r *MenuRepo) ListByIDs(ctx context.Context, id ...uint) (menus []*biz.Menu, err error) {
	db := r.data.DBD(ctx).Model(&SysMenu{})
	sysMenus := []*SysMenu{}

	err = db.Find(&sysMenus).Error
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
	result := r.data.DBD(ctx).Find(&sysMenus, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysMenus {
		bizMenus = append(bizMenus, r.toBiz(v))
	}
	return bizMenus, nil
}

func (r *MenuRepo) Delete(ctx context.Context, g *biz.Menu) error {
	result := r.data.DBD(ctx).Select("Parameters", "Buttons").Delete(r.toModel(g))
	if err := result.Error; err != nil {
		return err
	}
	return r.data.rdb.HDel(ctx, cacheMenuKey, convert.UnitToString(g.ID)).Err()
}

func (r *MenuRepo) ListAll(ctx context.Context) (menus []*biz.Menu, err error) {
	for _, v := range r.allCache(ctx) {
		menus = append(menus, r.toBiz(v))
	}
	return
}

func (r *MenuRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (menus []*biz.Menu, total int64) {
	db := r.data.DB(ctx).Model(&SysMenu{}).Debug()
	sysMenus := []*SysMenu{}
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

	r.data.DomainID(ctx)
	r.data.db.Model(&SysDomain{})
	// db.Where()

	result := db.Limit(int(handler.GetPageSize())).Find(&sysMenus)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysMenus {
		menus = append(menus, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(menus))
	}

	return menus, total
}

// setCache 设置菜单缓存
func (r *MenuRepo) setCache(ctx context.Context, g *SysMenu) error {
	dataStr, err := json.Marshal(g)
	if err != nil {
		r.log.Errorf("菜单缓存失败 %v", err)
		return err
	}
	return r.data.rdb.HSet(ctx, cacheMenuKey, convert.UnitToString(g.ID), dataStr).Err()
}

// getCache 获取菜单缓存
func (r *MenuRepo) getCache(ctx context.Context, key string) (sysMenu *SysMenu) {
	dataStr, err := r.data.rdb.HGet(ctx, cacheMenuKey, key).Result()
	if err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(dataStr), &sysMenu); err != nil {
		r.log.Errorf("缓存反序列化失败 %v", err)
	}
	return sysMenu
}

func (r *MenuRepo) allCache(ctx context.Context) (menus []*SysMenu) {
	if l, _ := r.data.rdb.HLen(ctx, cacheMenuKey).Result(); l > 0 {
		menuMap, _ := r.data.rdb.HGetAll(ctx, cacheMenuKey).Result()
		for _, v := range menuMap {
			sysMenu := SysMenu{}
			err := json.Unmarshal([]byte(v), &sysMenu)
			if err != nil {
				r.log.Errorf("菜单缓存反序列失败 %v", err)
				continue
			}
			menus = append(menus, &sysMenu)
		}
	} else {
		result := r.data.DB(ctx).Find(&menus)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.log.Errorf("菜单查询失败 %v", result.Error)
			return nil
		}
		menuMap := make(map[string]interface{})
		for _, v := range menus {
			menuStr, err := json.Marshal(v)
			if err != nil {
				r.log.Errorf("菜单缓存序列化失败 %v", err)
				continue
			}
			menuMap[convert.UnitToString(v.ID)] = string(menuStr)
		}
		if err := r.data.rdb.HSet(ctx, cacheMenuKey, menuMap).Err(); err != nil {
			r.log.Errorf("菜单缓存失败 %v", err)
		}
	}
	// 根据序号进行排序
	sort.SliceStable(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	return menus
}

// 根据ID递归查询父级菜单
func menuRecursiveParent(menus []*biz.Menu, ids ...uint) []*biz.Menu {
	result, mid := []*biz.Menu{}, map[uint]uint{}
	for _, v := range menus {
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
	// 根据序号进行排序
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})
	return result
}
