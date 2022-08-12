package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type DomainRepo struct {
	data *Data
	log  *log.Helper
}

// NewDomainRepo .
func NewDomainRepo(data *Data, logger log.Logger) biz.DomainRepo {
	return &DomainRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DomainRepo) toModel(d *biz.Domain) *SysDomain {
	if d == nil {
		return nil
	}
	return &SysDomain{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		DomainID: d.DomainID,
		Name:     d.Name,
		State:    d.State,
	}
}

func (r *DomainRepo) toBiz(d *SysDomain) *biz.Domain {
	if d == nil {
		return nil
	}
	return &biz.Domain{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		DomainID:  d.DomainID,
		Name:      d.Name,
		State:     d.State,
	}
}

func (r *DomainRepo) Save(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	d := r.toModel(g)
	id := r.data.sf.Generate()
	// g.DomainID = base64.StdEncoding.EncodeToString([]byte(id.String()))
	d.DomainID = id.String()
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *DomainRepo) Update(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	return g, nil
}

func (r *DomainRepo) FindByID(context.Context, int64) (*biz.Domain, error) {
	return nil, nil
}

func (r *DomainRepo) FindByDomainID(ctx context.Context, domainId string) (*biz.Domain, error) {
	sysDomain := &SysDomain{}
	result := r.data.DB(ctx).Debug().Last(sysDomain, "domain_id", domainId)
	return r.toBiz(sysDomain), result.Error
}

func (r *DomainRepo) ListByName(context.Context, string) ([]*biz.Domain, error) {
	return nil, nil
}

func (r *DomainRepo) ListAll(ctx context.Context) ([]*biz.Domain, error) {
	return nil, nil
}

func (r *DomainRepo) FindInDomainID(ctx context.Context, domainIds ...string) ([]*biz.Domain, error) {
	sysDomains, bizDomains := []*SysDomain{}, []*biz.Domain{}
	result := r.data.DB(ctx).Debug().Where("domain_id", domainIds).Find(sysDomains)

	for _, v := range sysDomains {
		bizDomains = append(bizDomains, r.toBiz(v))
	}
	return bizDomains, result.Error
}
