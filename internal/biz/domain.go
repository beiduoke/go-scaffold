package biz

import (
	"context"
	"strconv"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/pkg/authz"
	casbinM "github.com/beiduoke/go-scaffold/pkg/authz/casbin"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

type DomainAuthorityUser struct {
	DomainID    uint
	AuthorityID uint
	UserID      uint
	CreatedAt   time.Time
}

// Domain is a Domain model.
type Domain struct {
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ID                 uint
	Code               string
	ParentID           uint
	Name               string
	Sort               int32
	State              int32
	DefaultAuthorityID uint
	Authority          *Authority
}

// DomainRepo is a Greater repo.
type DomainRepo interface {
	Save(context.Context, *Domain) (*Domain, error)
	Update(context.Context, *Domain) (*Domain, error)
	FindByID(context.Context, uint) (*Domain, error)
	FindByCode(context.Context, string) (*Domain, error)
	FindByName(context.Context, string) (*Domain, error)
	ListByName(context.Context, string) ([]*Domain, error)
	Delete(context.Context, *Domain) error
	ListAll(context.Context) ([]*Domain, error)
	ListPage(context.Context, pagination.PaginationHandler) ([]*Domain, int64)
	FindInDomainID(context.Context, ...string) ([]*Domain, error)

	// 领域权限
	// stdcasbin.IEnforcer
	// SaveAuthorityUser 添加领域角色权限
	// SaveAuthorityForUserInDomain(context.Context, uint /* userID */, uint /* authorityID */, uint /* domainID */) error
	// GetAuthorityForUserInDomain 获取领域用户的角色
	// GetAuthoritiesForUserInDomain(context.Context, uint /* userID */, uint /* domainID */) []*Authority
	// GetAuthoritiesForUserInDomain 获取具有域内角色的用户
	// GetUsersForRoleInDomain(context.Context, uint /* authorityID */, uint /* domainID */) []*User
	// DeleteRoleForUserInDomain 域内删除用户的角色域内删除用户的角色
	// DeleteRoleForUserInDomain(context.Context, uint /* userID */, uint /* domainID */) error
}

// DomainUsecase is a Domain usecase.
type DomainUsecase struct {
	biz *Biz
	log *log.Helper
}

// NewDomainUsecase new a Domain usecase.
func NewDomainUsecase(logger log.Logger, biz *Biz) *DomainUsecase {
	return &DomainUsecase{biz: biz, log: log.NewHelper(logger)}
}

// Create creates a Domain, and returns the new Domain.
func (uc *DomainUsecase) Create(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		domain, err := uc.biz.domainRepo.Save(ctx, g)
		if err != nil {
			return err
		}
		authCtx := context.WithValue(ctx, casbinM.SecurityUserContextKey, authz.SecurityUser{Domain: strconv.Itoa(int(domain.ID))})
		authority, err := uc.biz.authorityRepo.Save(authCtx, &Authority{
			Name: "default",
		})
		if err != nil {
			return err
		}
		domain.DefaultAuthorityID = authority.ID
		g, err = uc.biz.domainRepo.Update(ctx, domain)
		return err
	})
	return g, err
}

// ListByIDs 获取指定领域ID集合
func (uc *DomainUsecase) ListByIDs(ctx context.Context, id ...uint) (authorities []*Domain, err error) {
	authorities, _ = uc.biz.domainRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改领域
func (uc *DomainUsecase) Update(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Infof("UpdateDomain: %v", g)

	domain, _ := uc.biz.domainRepo.FindByID(ctx, g.ID)
	if domain == nil {
		return errors.New("领域未注册")
	}

	if domain.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.domainRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("领域名已存在")
		}
	}

	if g.State <= 0 {
		g.State = int32(pb.DomainState_DOMAIN_STATE_ACTIVE)
	}

	// 新数据合并到源数据
	if err := mergo.Merge(domain, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}

	_, err := uc.biz.domainRepo.Update(ctx, domain)
	return err
}

// List 领域列表全部
func (uc *DomainUsecase) ListAll(ctx context.Context) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Infof("DomainList")
	return uc.biz.domainRepo.ListPage(ctx, pagination.NewPagination())
}

// List 领域列表分页
func (uc *DomainUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Infof("DomainPage")
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
	return uc.biz.domainRepo.ListPage(ctx, page)
}

// GetID 根据角色ID领域
func (uc *DomainUsecase) GetID(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Infof("GetDomainID: %v", g)
	return uc.biz.domainRepo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除领域
func (uc *DomainUsecase) Delete(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Infof("DeleteDomain: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.domainRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(convert.UnitToString(g.ID))
		return err
	})
}
