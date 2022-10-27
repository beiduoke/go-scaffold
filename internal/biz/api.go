package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Api is a Api model.
type Api struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Path        string
	Method      string
	Group       string
	Description string
}

// ApiRepo is a Greater repo.
type ApiRepo interface {
	Save(context.Context, *Api) (*Api, error)
	Update(context.Context, *Api) (*Api, error)
	FindByID(context.Context, uint) (*Api, error)
	FindByName(context.Context, string) (*Api, error)
	ListByName(context.Context, string) ([]*Api, error)
	ListAll(context.Context) ([]*Api, error)
	Delete(context.Context, *Api) error
	ListPage(context.Context, pagination.PaginationHandler) ([]*Api, int64)
}

// ApiUsecase is a Api usecase.
type ApiUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo ApiRepo
}

// NewApiUsecase new a Api usecase.
func NewApiUsecase(logger log.Logger, biz *Biz, repo ApiRepo) *ApiUsecase {
	return &ApiUsecase{log: log.NewHelper(logger), repo: repo, biz: biz}
}

// Create creates a Api, and returns the new Api.
func (uc *ApiUsecase) Create(ctx context.Context, g *Api) (*Api, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// ListByIDs 获取指定接口ID集合
func (uc *ApiUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Api, err error) {
	authorities, _ = uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改接口
func (uc *ApiUsecase) Update(ctx context.Context, g *Api) error {
	uc.log.WithContext(ctx).Infof("UpdateApi: %v", g)

	api, _ := uc.repo.FindByID(ctx, g.ID)
	if api == nil {
		return errors.New("接口未注册")
	}

	if api.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("接口名已存在")
		}
	}

	// 新数据合并到源数据
	if err := mergo.Merge(api, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	_, err := uc.repo.Update(ctx, api)
	return err
}

// List 接口列表全部
func (uc *ApiUsecase) ListAll(ctx context.Context) ([]*Api, int64) {
	uc.log.WithContext(ctx).Infof("ApiList")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// List 接口列表分页
func (uc *ApiUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Api, int64) {
	uc.log.WithContext(ctx).Infof("ApiPage")
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

// GetID 根据角色ID接口
func (uc *ApiUsecase) GetID(ctx context.Context, g *Api) (*Api, error) {
	uc.log.WithContext(ctx).Infof("GetApiID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除接口
func (uc *ApiUsecase) Delete(ctx context.Context, g *Api) error {
	uc.log.WithContext(ctx).Infof("DeleteApi: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.repo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(convert.UnitToString(g.ID))
		return err
	})
}
