package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
)

// Menu is a Menu model.
type Menu struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ParentID  uint
	Path      string
	Hidden    int32
	Component string
	Sort      int
	Icon      string
	Title     string
	KeepAlive int32
	BaseMenu  int32
	CloseTab  int32
}

// MenuRepo is a Greater repo.
type MenuRepo interface {
	Save(context.Context, *Menu) (*Menu, error)
	Update(context.Context, *Menu) (*Menu, error)
	FindByID(context.Context, uint) (*Menu, error)
	ListByName(context.Context, string) ([]*Menu, error)
	ListAll(context.Context) ([]*Menu, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*Menu, int64)
}

// MenuUsecase is a Menu usecase.
type MenuUsecase struct {
	*Biz
	repo MenuRepo
}

// NewMenuUsecase new a Menu usecase.
func NewMenuUsecase(biz *Biz, repo MenuRepo) *MenuUsecase {
	return &MenuUsecase{repo: repo, Biz: biz}
}

// Create creates a Menu, and returns the new Menu.
func (uc *MenuUsecase) Create(ctx context.Context, g *Menu) (*Menu, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
