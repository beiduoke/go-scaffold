package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
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
	return &SysMenu{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:      d.Name,
		ParentID:  d.ParentID,
		Path:      d.Path,
		Hidden:    d.Hidden,
		Component: d.Component,
		Sort:      d.Sort,
		Meta: SysMeta{
			Icon:      d.Icon,
			Title:     d.Title,
			KeepAlive: d.KeepAlive,
			BaseMenu:  d.BaseMenu,
			CloseTab:  d.CloseTab,
		},
	}
}

func (r *MenuRepo) toBiz(d *SysMenu) *biz.Menu {
	if d == nil {
		return nil
	}
	return &biz.Menu{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Name:      d.Name,
		ParentID:  d.ParentID,
		Path:      d.Path,
		Hidden:    d.Hidden,
		Component: d.Component,
		Sort:      d.Sort,
		Icon:      d.Meta.Icon,
		Title:     d.Meta.Title,
		KeepAlive: d.Meta.KeepAlive,
		BaseMenu:  d.Meta.BaseMenu,
		CloseTab:  d.Meta.CloseTab,
	}
}

func (r *MenuRepo) Save(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *MenuRepo) Update(ctx context.Context, g *biz.Menu) (*biz.Menu, error) {
	return g, nil
}

func (r *MenuRepo) FindByID(context.Context, int64) (*biz.Menu, error) {
	return nil, nil
}

func (r *MenuRepo) ListByName(context.Context, string) ([]*biz.Menu, error) {
	return nil, nil
}

func (r *MenuRepo) ListAll(context.Context) ([]*biz.Menu, error) {
	return nil, nil
}

func (r *MenuRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (menus []*biz.Menu, total int64) {
	db := r.data.DB(ctx).Model(&SysMenu{})
	sysMenus := []*SysMenu{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}
	result := db.Count(&total).Offset(handler.GetPageOffset()).Limit(int(handler.GetPageSize())).Find(&sysMenus)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysMenus {
		menus = append(menus, r.toBiz(v))
	}
	return menus, total
}
