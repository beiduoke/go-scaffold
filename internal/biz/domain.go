package biz

import (
	"context"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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
	DomainID           string
	Name               string
	State              int32
	DefaultAuthorityID uint
}

// DomainRepo is a Greater repo.
type DomainRepo interface {
	Save(context.Context, *Domain) (*Domain, error)
	Update(context.Context, *Domain) (*Domain, error)
	FindByID(context.Context, uint) (*Domain, error)
	FindByDomainID(context.Context, string) (*Domain, error)
	ListByName(context.Context, string) ([]*Domain, error)
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
	return uc.biz.domainRepo.Save(ctx, g)
}

// GetDomainID 获取指定领域ID
func (uc *DomainUsecase) GetByDomainID(ctx context.Context, domainId string) (*Domain, error) {
	return uc.biz.domainRepo.FindByDomainID(ctx, domainId)
}

// GetDomainInID 获取指定领域ID集合
func (uc *DomainUsecase) ListByIDs(ctx context.Context, id ...uint) (domains []*Domain, err error) {
	domains, _ = uc.biz.domainRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}
