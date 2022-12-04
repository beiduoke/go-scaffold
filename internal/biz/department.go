package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Department is a Department model.
type Department struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ParentID  uint
	Sort      int32
	Remarks   string
	State     int32
	Children  []*Department
}

// DepartmentRepo is a Greater repo.
type DepartmentRepo interface {
	Save(context.Context, *Department) (*Department, error)
	Update(context.Context, *Department) (*Department, error)
	FindByID(context.Context, uint) (*Department, error)
	FindByName(context.Context, string) (*Department, error)
	ListByName(context.Context, string) ([]*Department, error)
	ListAll(context.Context) ([]*Department, error)
	Delete(context.Context, *Department) error
	ListPage(context.Context, pagination.PaginationHandler) ([]*Department, int64)
}

type DepartmentUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo DepartmentRepo
}

// NewDepartmentUsecase new a Department usecase.
func NewDepartmentUsecase(logger log.Logger, biz *Biz, repo DepartmentRepo) *DepartmentUsecase {
	return &DepartmentUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Department, and returns the new Department.
func (uc *DepartmentUsecase) Create(ctx context.Context, g *Department) (*Department, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定部门ID集合
func (uc *DepartmentUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Department, err error) {
	authorities, _ = uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改部门
func (uc *DepartmentUsecase) Update(ctx context.Context, g *Department) error {
	uc.log.WithContext(ctx).Infof("UpdateDepartment: %v", g)

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
func (uc *DepartmentUsecase) ListAll(ctx context.Context) ([]*Department, int64) {
	uc.log.WithContext(ctx).Infof("DepartmentList")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// List 部门列表分页
func (uc *DepartmentUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Department, int64) {
	uc.log.WithContext(ctx).Infof("DepartmentPage")
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
func (uc *DepartmentUsecase) GetTree(ctx context.Context, id uint) []*Department {
	uc.log.WithContext(ctx).Infof("GetTree")
	menus, _ := uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrder("sort", false)))
	return menus
}

// GetID 根据角色ID部门
func (uc *DepartmentUsecase) GetID(ctx context.Context, g *Department) (*Department, error) {
	uc.log.WithContext(ctx).Infof("GetDepartmentID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除部门
func (uc *DepartmentUsecase) Delete(ctx context.Context, g *Department) error {
	uc.log.WithContext(ctx).Infof("DeleteDepartment: %v", g)
	return uc.repo.Delete(ctx, g)
}
