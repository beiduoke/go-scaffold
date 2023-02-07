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
	ListPage(context.Context, pagination.PaginationHandler) ([]*Dept, int64)
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
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定部门ID集合
func (uc *DeptUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Dept, err error) {
	roles, _ = uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改部门
func (uc *DeptUsecase) Update(ctx context.Context, g *Dept) error {
	uc.log.WithContext(ctx).Infof("UpdateDept: %v", g)

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
	uc.log.WithContext(ctx).Infof("DeptList")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// List 部门列表分页
func (uc *DeptUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Dept, int64) {
	uc.log.WithContext(ctx).Infof("DeptPage")
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

// GetTree 获取部门树形
func (uc *DeptUsecase) GetTree(ctx context.Context, id uint) []*Dept {
	uc.log.WithContext(ctx).Infof("GetTree")
	menus, _ := uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrder("sort", false)))
	return menus
}

// GetID 根据角色ID部门
func (uc *DeptUsecase) GetID(ctx context.Context, g *Dept) (*Dept, error) {
	uc.log.WithContext(ctx).Infof("GetDeptID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除部门
func (uc *DeptUsecase) Delete(ctx context.Context, g *Dept) error {
	uc.log.WithContext(ctx).Infof("DeleteDept: %v", g)
	return uc.repo.Delete(ctx, g)
}