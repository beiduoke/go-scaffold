package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// Resource is a Resource model.
type Resource struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Path        string
	Method      string
	Operation   string
	Group       string
	Description string
}

// ResourceRepo is a Greater repo.
type ResourceRepo interface {
	Save(context.Context, *Resource) (*Resource, error)
	Update(context.Context, *Resource) (*Resource, error)
	FindByID(context.Context, uint) (*Resource, error)
	FindByName(context.Context, string) (*Resource, error)
	ListByName(context.Context, string) ([]*Resource, error)
	ListAll(context.Context) ([]*Resource, error)
	ListAllGroup(context.Context) ([]string, error)
	Delete(context.Context, *Resource) error
	ListPage(context.Context, *pagination.Pagination) ([]*Resource, int64)
}

// ResourceUsecase is a Resource usecase.
type ResourceUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo ResourceRepo
}

// NewResourceUsecase new a Resource usecase.
func NewResourceUsecase(logger log.Logger, biz *Biz, repo ResourceRepo) *ResourceUsecase {
	return &ResourceUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Resource, and returns the new Resource.
func (uc *ResourceUsecase) Create(ctx context.Context, g *Resource) (*Resource, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定资源ID集合
func (uc *ResourceUsecase) ListByIDs(ctx context.Context, id ...uint) (resources []*Resource, err error) {
	// uc.repo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
	return
}

// Update 修改资源
func (uc *ResourceUsecase) Update(ctx context.Context, g *Resource) error {
	uc.log.WithContext(ctx).Debugf("UpdateResource: %v", g)

	api, _ := uc.repo.FindByID(ctx, g.ID)
	if api == nil {
		return errors.New("资源未注册")
	}

	if api.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("资源名已存在")
		}
	}

	_, err := uc.repo.Update(ctx, g)
	return err
}

// ListAll 资源列表全部
func (uc *ResourceUsecase) ListAll(ctx context.Context) ([]*Resource, int64) {
	uc.log.WithContext(ctx).Debugf("ResourceList")
	return uc.repo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// ListAllGroup 资源列表全部分组
func (uc *ResourceUsecase) ListAllGroup(ctx context.Context) ([]string, int64) {
	uc.log.WithContext(ctx).Debugf("ResourceListGroup")
	groups, _ := uc.repo.ListAllGroup(ctx)
	return groups, int64(len(groups))
}

// List 资源列表分页
func (uc *ResourceUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Resource, int64) {
	uc.log.WithContext(ctx).Debugf("ResourcePage")
	return uc.repo.ListPage(ctx, paging)
}

// GetID 根据角色ID资源
func (uc *ResourceUsecase) GetID(ctx context.Context, g *Resource) (*Resource, error) {
	uc.log.WithContext(ctx).Debugf("GetResourceID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除资源
func (uc *ResourceUsecase) Delete(ctx context.Context, g *Resource) error {
	uc.log.WithContext(ctx).Debugf("DeleteResource: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.repo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(convert.UnitToString(g.ID))
		return err
	})
}
