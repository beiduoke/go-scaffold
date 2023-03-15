package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Dept is a Dept model.
type Dept struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ParentID  uint
	Sort      int32
	Remarks   string
	State     int32
	Children  []*Dept
}

// DeptRepo is a Greater repo.
type DeptRepo interface {
	Save(context.Context, *Dept) (*Dept, error)
	Update(context.Context, *Dept) (*Dept, error)
	FindByID(context.Context, uint) (*Dept, error)
	FindByName(context.Context, string) (*Dept, error)
	ListByName(context.Context, string) ([]*Dept, error)
	ListAll(context.Context) ([]*Dept, error)
	Delete(context.Context, *Dept) error
	ListPage(context.Context, *pagination.Pagination) ([]*Dept, int64)
}

type DeptUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo DeptRepo
}

// NewDeptUsecase new a Dept usecase.
func NewDeptUsecase(logger log.Logger, biz *Biz, repo DeptRepo) *DeptUsecase {
	return &DeptUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Dept, and returns the new Dept.
func (uc *DeptUsecase) Create(ctx context.Context, g *Dept) (*Dept, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定部门ID集合
func (uc *DeptUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Dept, err error) {
	// roles, _ = uc.repo.ListPage(ctx, &pagination.Pagination{Nopaging: true, Query: map[string]string{"id": true}})
	return
}

// Update 修改部门
func (uc *DeptUsecase) Update(ctx context.Context, g *Dept) error {
	uc.log.WithContext(ctx).Debugf("UpdateDept: %v", g)

	menu, _ := uc.repo.FindByID(ctx, g.ID)
	if menu == nil {
		return errors.New("部门未注册")
	}

	if menu.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("部门名已存在")
		}
	}

	// 新数据合并到源数据
	if err := mergo.Merge(menu, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	_, err := uc.repo.Update(ctx, menu)
	return err
}

// List 部门列表全部
func (uc *DeptUsecase) ListAll(ctx context.Context) ([]*Dept, int64) {
	uc.log.WithContext(ctx).Debugf("ListAll")
	return uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrderBy(map[string]bool{"id": true, "sort": true})))
}

// List 部门列表分页
func (uc *DeptUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Dept, int64) {
	uc.log.WithContext(ctx).Debugf("DeptPage")
	return uc.repo.ListPage(ctx, paging)
}

// GetID 根据角色ID部门
func (uc *DeptUsecase) GetID(ctx context.Context, g *Dept) (*Dept, error) {
	uc.log.WithContext(ctx).Debugf("GetDeptID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除部门
func (uc *DeptUsecase) Delete(ctx context.Context, g *Dept) error {
	uc.log.WithContext(ctx).Debugf("DeleteDept: %v", g)
	return uc.repo.Delete(ctx, g)
}
