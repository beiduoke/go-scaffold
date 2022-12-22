package biz

import (
	"context"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Authority is a Authority model.
type Authority struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	ParentID      uint
	DefaultRouter string
	Sort          int32
	State         int32
	Remarks       string
	Users         []*User
	Domains       []*Domain
	Menus         []*Menu
	Apis          []*Api
}

// AuthorityRepo is a Greater repo.
type AuthorityRepo interface {
	Save(context.Context, *Authority) (*Authority, error)
	Update(context.Context, *Authority) (*Authority, error)
	FindByID(context.Context, uint) (*Authority, error)
	FindByName(context.Context, string) (*Authority, error)
	ListByName(context.Context, string) ([]*Authority, error)
	ListByIDs(context.Context, ...uint) ([]*Authority, error)
	ListAll(context.Context) ([]*Authority, error)
	Delete(context.Context, *Authority) error
	ListPage(context.Context, pagination.PaginationHandler) ([]*Authority, int64)
	HandleMenu(context.Context, *Authority) error
	HandleApi(context.Context, *Authority) error
	ListMenuByIDs(context.Context, ...uint) ([]*Menu, error)
	ListMenuAndParentByIDs(context.Context, ...uint) ([]*Menu, error)
}

// AuthorityUsecase is a Authority usecase.
type AuthorityUsecase struct {
	biz *Biz
	log *log.Helper
}

// NewAuthorityUsecase new a Authority usecase.
func NewAuthorityUsecase(logger log.Logger, biz *Biz) *AuthorityUsecase {
	return &AuthorityUsecase{log: log.NewHelper(logger), biz: biz}
}

// Create creates a Authority, and returns the new Authority.
func (uc *AuthorityUsecase) Create(ctx context.Context, g *Authority) (*Authority, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.biz.authorityRepo.Save(ctx, g)
}

// ListByIDs 获取指定权限角色ID集合
func (uc *AuthorityUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Authority, err error) {
	authorities, _ = uc.biz.authorityRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改权限角色
func (uc *AuthorityUsecase) Update(ctx context.Context, g *Authority) error {
	uc.log.WithContext(ctx).Infof("UpdateAuthority: %v", g)

	authority, _ := uc.biz.authorityRepo.FindByID(ctx, g.ID)
	if authority == nil {
		return errors.New("权限角色未注册")
	}

	if authority.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.authorityRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("权限角色名已存在")
		}
	}

	if g.State <= 0 {
		g.State = int32(pb.AuthorityState_AUTHORITY_STATE_ACTIVE)
	}
	// 新数据合并到源数据
	if err := mergo.Merge(authority, *g, mergo.WithOverwriteWithEmptyValue); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}
	_, err := uc.biz.authorityRepo.Update(ctx, authority)
	return err
}

// List 权限角色列表全部
func (uc *AuthorityUsecase) ListAll(ctx context.Context) ([]*Authority, int64) {
	uc.log.WithContext(ctx).Infof("AuthorityList")
	return uc.biz.authorityRepo.ListPage(ctx, pagination.NewPagination())
}

// List 权限角色列表分页
func (uc *AuthorityUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Authority, int64) {
	uc.log.WithContext(ctx).Infof("AuthorityPage")
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
	return uc.biz.authorityRepo.ListPage(ctx, page)
}

// GetID 根据角色ID权限角色
func (uc *AuthorityUsecase) GetID(ctx context.Context, g *Authority) (*Authority, error) {
	uc.log.WithContext(ctx).Infof("GetAuthorityID: %v", g)
	return uc.biz.authorityRepo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除权限角色
func (uc *AuthorityUsecase) Delete(ctx context.Context, g *Authority) error {
	uc.log.WithContext(ctx).Infof("DeleteAuthority: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.authorityRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(convert.UnitToString(g.ID))
		return err
	})
}

// HandleMenu 获取权限角色菜单
func (uc *AuthorityUsecase) ListMenuByID(ctx context.Context, g *Authority) ([]*Menu, error) {
	uc.log.WithContext(ctx).Infof("ListMenuByIDs: %v", g)
	return uc.biz.authorityRepo.ListMenuByIDs(ctx, g.ID)
}

// HandleMenu 绑定菜单
func (uc *AuthorityUsecase) HandleMenu(ctx context.Context, g *Authority) error {
	uc.log.WithContext(ctx).Infof("HandleMenu: %v", g)
	return uc.biz.authorityRepo.HandleMenu(ctx, g)
}

// HandleApi 绑定接口
func (uc *AuthorityUsecase) HandleApi(ctx context.Context, g *Authority) error {
	uc.log.WithContext(ctx).Infof("HandleApi: %v", g)
	return uc.biz.authorityRepo.HandleApi(ctx, g)
}
