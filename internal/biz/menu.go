package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Menu is a Menu model.
type Menu struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Type       int32
	ParentID   uint
	Path       string
	Hidden     int32
	Component  string
	Permission string
	Sort       int32
	Icon       string
	Title      string
	KeepAlive  int32
	BaseMenu   int32
	CloseTab   int32
	ExtType    int32
	Children   []*Menu
	Parameters []*MenuParameter
	Buttons    []*MenuButton
}

type MenuParameter struct {
	ID    uint
	Type  int32
	Name  string
	Value string
}

type MenuButton struct {
	ID      uint
	Name    string
	Remarks string
}

// MenuRepo is a Greater repo.
type MenuRepo interface {
	Save(context.Context, *Menu) (*Menu, error)
	Update(context.Context, *Menu) (*Menu, error)
	FindByID(context.Context, uint) (*Menu, error)
	FindByName(context.Context, string) (*Menu, error)
	ListByName(context.Context, string) ([]*Menu, error)
	ListAll(context.Context) ([]*Menu, error)
	Delete(context.Context, *Menu) error
	ListPage(context.Context, pagination.PaginationHandler) ([]*Menu, int64)
}

type MenuUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo MenuRepo
}

// NewMenuUsecase new a Menu usecase.
func NewMenuUsecase(logger log.Logger, biz *Biz, repo MenuRepo) *MenuUsecase {
	return &MenuUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Menu, and returns the new Menu.
func (uc *MenuUsecase) Create(ctx context.Context, g *Menu) (*Menu, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定菜单ID集合
func (uc *MenuUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Menu, err error) {
	authorities, _ = uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改菜单
func (uc *MenuUsecase) Update(ctx context.Context, g *Menu) error {
	uc.log.WithContext(ctx).Infof("UpdateMenu: %v", g)

	menu, _ := uc.repo.FindByID(ctx, g.ID)
	if menu == nil {
		return errors.New("菜单未注册")
	}

	if menu.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("菜单名已存在")
		}
	}

	// 清理原始数据
	menu.Parameters = []*MenuParameter{}
	menu.Buttons = []*MenuButton{}
	menu.Component = ""
	// 新数据合并到源数据
	if err := mergo.Merge(menu, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	fmt.Println(menu.Component, "384837373737373737337373737")
	_, err := uc.repo.Update(ctx, menu)
	return err
}

// List 菜单列表全部
func (uc *MenuUsecase) ListAll(ctx context.Context) ([]*Menu, int64) {
	uc.log.WithContext(ctx).Infof("MenuList")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// List 菜单列表分页
func (uc *MenuUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Menu, int64) {
	uc.log.WithContext(ctx).Infof("MenuPage")
	conditions := []pagination.Condition{}
	for k, v := range query {
		conditions = append(conditions, pagination.Condition{Query: k, Args: []interface{}{v}})
	}
	orders := []pagination.Order{}
	for k, v := range order {
		orders = append(orders, pagination.Order{Column: k, Desc: v})
	}

	page := pagination.NewPagination(
		pagination.WithPageNum(pageNum),
		pagination.WithPageSize(pageSize),
		pagination.WithConditions(conditions...),
		pagination.WithOrders(orders...),
	)
	return uc.repo.ListPage(ctx, page)
}

// GetTree 获取菜单树形
func (uc *MenuUsecase) GetTree(ctx context.Context, id uint) []*Menu {
	uc.log.WithContext(ctx).Infof("GetTree")
	menus, _ := uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrder("sort", false)))
	return menus
}

// GetID 根据角色ID菜单
func (uc *MenuUsecase) GetID(ctx context.Context, g *Menu) (*Menu, error) {
	uc.log.WithContext(ctx).Infof("GetMenuID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除菜单
func (uc *MenuUsecase) Delete(ctx context.Context, g *Menu) error {
	uc.log.WithContext(ctx).Infof("DeleteMenu: %v", g)
	return uc.repo.Delete(ctx, g)
}
