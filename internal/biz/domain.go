package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Domain is a Domain model.
type Domain struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint
	DomainID  string
	Name      string
	State     int32
}

// DomainRepo is a Greater repo.
type DomainRepo interface {
	Save(context.Context, *Domain) (*Domain, error)
	Update(context.Context, *Domain) (*Domain, error)
	FindByID(context.Context, int64) (*Domain, error)
	FindByDomainID(context.Context, string) (*Domain, error)
	ListByName(context.Context, string) ([]*Domain, error)
	ListAll(context.Context) ([]*Domain, error)
	FindInDomainID(context.Context, ...string) ([]*Domain, error)
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

// CreateDomain creates a Domain, and returns the new Domain.
func (uc *DomainUsecase) CreateDomain(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Infof("CreateDomain: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// CreateDomain creates a Domain, and returns the new Domain.
func (uc *DomainUsecase) FindByDomainID(ctx context.Context, domainId string) (*Domain, error) {
	return uc.repo.FindByDomainID(ctx, domainId)
}
