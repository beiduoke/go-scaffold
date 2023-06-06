package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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
	Redirect   string
	Component  string
	Permission string
	Sort       int32
	Icon       string
	Title      string
	IsHidden   int32
	IsCache    int32
	IsAffix    int32
	LinkType   int32
	LinkUrl    string
	Parent     *Menu
	Children   []*Menu
	Parameters []*MenuParameter
	Buttons    []*MenuButton
}

func (g Menu) GetID() string {
	return convert.UnitToString(g.ID)
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
	ListPage(context.Context, *pagination.Pagination) ([]*Menu, int64)
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
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定菜单ID集合
func (uc *MenuUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Menu, err error) {
	// roles, _ = uc.repo.ListPage(ctx, noop.NewPagination(noop.WithNopaging(), noop.WithCondition("id in ?", id)))
	return
}

// Update 修改菜单
func (uc *MenuUsecase) Update(ctx context.Context, g *Menu) error {
	uc.log.WithContext(ctx).Debugf("UpdateMenu: %v", g)
	_, err := uc.repo.Update(ctx, g)
	return err
}

// List 菜单列表全部
func (uc *MenuUsecase) ListAll(ctx context.Context) ([]*Menu, int64) {
	uc.log.WithContext(ctx).Debugf("MenuList")
	return uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrderBy(map[string]bool{"id": true, "sort": true})))
}

// List 菜单列表分页
func (uc *MenuUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Menu, int64) {
	uc.log.WithContext(ctx).Debugf("MenuPage")
	return uc.repo.ListPage(ctx, paging)
}

// GetID 根据角色ID菜单
func (uc *MenuUsecase) GetID(ctx context.Context, g *Menu) (*Menu, error) {
	uc.log.WithContext(ctx).Debugf("GetMenuID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除菜单
func (uc *MenuUsecase) Delete(ctx context.Context, g *Menu) error {
	uc.log.WithContext(ctx).Debugf("DeleteMenu: %v", g)
	return uc.repo.Delete(ctx, g)
}
