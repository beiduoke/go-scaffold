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
	// SaveAuthorityUser 添加领域角色权限
	SaveAuthorityForUserInDomain(context.Context, uint /* userID */, uint /* authorityID */, uint /* domainID */) error
	// FindAuthorityForUserInDomain 获取领域用户的角色
	FindAuthoritiesForUserInDomain(context.Context, uint /* userID */, uint /* domainID */) []*Authority
	// FindAuthoritiesForUserInDomain 获取具有域内角色的用户
	FindUsersForRoleInDomain(context.Context, uint /* authorityID */, uint /* domainID */) []*User
}

// DomainUsecase is a Domain usecase.
type DomainUsecase struct {
	repo DomainRepo
	log  *log.Helper
	tm   Transaction
}

// NewDomainUsecase new a Domain usecase.
func NewDomainUsecase(repo DomainRepo, tm Transaction, logger log.Logger) *DomainUsecase {
	return &DomainUsecase{repo: repo, tm: tm, log: log.NewHelper(logger)}
}

// Create creates a Domain, and returns the new Domain.
func (uc *DomainUsecase) Create(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// GetDomainID 获取指定领域ID
func (uc *DomainUsecase) GetDomainID(ctx context.Context, domainId string) (*Domain, error) {
	return uc.repo.FindByDomainID(ctx, domainId)
}
