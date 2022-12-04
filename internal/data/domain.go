package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DomainRepo struct {
	data *Data
	log  *log.Helper
}

// NewDomainRepo .
func NewDomainRepo(logger log.Logger, data *Data) biz.DomainRepo {
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
		ID:                 d.ID,
		CreatedAt:          d.CreatedAt,
		UpdatedAt:          d.UpdatedAt,
		Code:               d.Code,
		Name:               d.Name,
		State:              d.State,
		DefaultAuthorityID: d.DefaultAuthorityID,
		ParentID:           d.ParentID,
		Sort:               d.Sort,
	}
}

func (r *DomainRepo) toBiz(d *SysDomain) *biz.Domain {
	if d == nil {
		return nil
	}
	return &biz.Domain{
		CreatedAt:          d.CreatedAt,
		UpdatedAt:          d.UpdatedAt,
		ID:                 d.ID,
		Code:               d.Code,
		Name:               d.Name,
		State:              d.State,
		DefaultAuthorityID: d.DefaultAuthorityID,
	}
}

func (r *DomainRepo) Save(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	d := r.toModel(g)
	sfId := r.data.sf.Generate()
	// g.DomainID = base64.StdEncoding.EncodeToString([]byte(id.String()))
	d.Code = sfId.String()
	result := r.data.DB(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *DomainRepo) Update(ctx context.Context, g *biz.Domain) (*biz.Domain, error) {
	d := r.toModel(g)
	result := r.data.DB(ctx).Model(d).Omit("Code").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *DomainRepo) FindByID(ctx context.Context, id uint) (*biz.Domain, error) {
	domain := SysDomain{}
	result := r.data.DB(ctx).Last(&domain, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DomainRepo) FindByName(ctx context.Context, s string) (*biz.Domain, error) {
	domain := SysDomain{}
	result := r.data.DB(ctx).Last(&domain, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&domain), nil
}

func (r *DomainRepo) FindByCode(ctx context.Context, code string) (*biz.Domain, error) {
	sysDomain := &SysDomain{}
	result := r.data.DB(ctx).Last(sysDomain, "code", code)
	return r.toBiz(sysDomain), result.Error
}

func (r *DomainRepo) ListByIDs(ctx context.Context, id ...uint) (domains []*biz.Domain, err error) {
	db := r.data.DB(ctx).Model(&SysDomain{})
	sysDomains := []*SysDomain{}

	err = db.Find(&sysDomains).Error
	if err != nil {
		return domains, err
	}
	for _, v := range sysDomains {
		domains = append(domains, r.toBiz(v))
	}
	return
}

func (r *DomainRepo) ListByName(ctx context.Context, name string) ([]*biz.Domain, error) {
	sysDomains, bizDomains := []*SysDomain{}, []*biz.Domain{}
	result := r.data.DB(ctx).Find(&sysDomains, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysDomains {
		bizDomains = append(bizDomains, r.toBiz(v))
	}
	return bizDomains, nil
}

func (r *DomainRepo) Delete(ctx context.Context, g *biz.Domain) error {
	return r.data.DB(ctx).Delete(r.toModel(g)).Error
}

func (r *DomainRepo) ListAll(ctx context.Context) ([]*biz.Domain, error) {
	return nil, nil
}

func (r *DomainRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (domains []*biz.Domain, total int64) {
	db := r.data.DB(ctx).Model(&SysDomain{})
	sysDomains := []*SysDomain{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}

	if !handler.GetNopaging() {
		db = db.Count(&total).Offset(handler.GetPageOffset())
	}

	result := db.Limit(int(handler.GetPageSize())).Find(&sysDomains)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysDomains {
		domains = append(domains, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(domains))
	}

	return domains, total
}
