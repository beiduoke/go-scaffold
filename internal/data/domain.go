package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type DomainRepo struct {
	data     *Data
	log      *log.Helper
	enforcer stdcasbin.IEnforcer
}

// NewDomainRepo .
func NewDomainRepo(data *Data, enforcer stdcasbin.IEnforcer, logger log.Logger) biz.DomainRepo {
	return &DomainRepo{
		data:     data,
		log:      log.NewHelper(logger),
		enforcer: enforcer,
	}
}

func (r *DomainRepo) toModel(d *biz.Domain) *SysDomain {
	if d == nil {
		return nil
	}
	return &SysDomain{
		DomainModel: DomainModel{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
			DomainID:  d.DomainID,
		},
		Name:               d.Name,
		State:              d.State,
		DefaultAuthorityID: d.DefaultAuthorityID,
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
		DomainID:           d.DomainID,
		Name:               d.Name,
		State:              d.State,
		DefaultAuthorityID: d.DefaultAuthorityID,
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

	if !handler.GetNopaging() {
		total = int64(len(domains))
	}

	return domains, total
}

func (r *DomainRepo) FindInDomainID(ctx context.Context, domainIds ...string) ([]*biz.Domain, error) {
	sysDomains, bizDomains := []*SysDomain{}, []*biz.Domain{}
	result := r.data.DB(ctx).Debug().Where("domain_id", domainIds).Find(sysDomains)

	for _, v := range sysDomains {
		bizDomains = append(bizDomains, r.toBiz(v))
	}
	return bizDomains, result.Error
}

func (r *DomainRepo) SaveAuthorityForUserInDomain(ctx context.Context, userID, authorityID, domainID uint) error {
	success, err := r.enforcer.AddRoleForUserInDomain(convert.UnitToString(userID), convert.UnitToString(authorityID), convert.UnitToString(domainID))
	if !success {
		r.log.Warnf("域内为用户添加角色 %v", success)
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *DomainRepo) FindAuthoritiesForUserInDomain(ctx context.Context, userID, domainID uint) (authorities []*biz.Authority) {
	roles := r.enforcer.GetRolesForUserInDomain(convert.UnitToString(userID), convert.UnitToString(domainID))
	authorityIDs := []uint{}
	for _, v := range roles {
		authorityIDs = append(authorityIDs, convert.StringToUint(v))
	}
	sysAuthorities := []*SysAuthority{}
	_ = r.data.DB(ctx).Find(&sysAuthorities, authorityIDs)

	for _, v := range sysAuthorities {
		authorities = append(authorities, &biz.Authority{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return
}

// FindUsersForRoleInDomain 获取具有域内角色的用户
func (r *DomainRepo) FindUsersForRoleInDomain(ctx context.Context, authorityID, domainID uint) (users []*biz.User) {
	roles := r.enforcer.GetUsersForRoleInDomain(convert.UnitToString(authorityID), convert.UnitToString(domainID))
	userIDs := []uint{}
	for _, v := range roles {
		userIDs = append(userIDs, convert.StringToUint(v))
	}
	sysUsers := []*SysUser{}
	_ = r.data.DB(ctx).Find(&sysUsers, userIDs)

	userRepo := UserRepo{}
	for _, v := range sysUsers {
		users = append(users, userRepo.toBiz(v))
	}

	return
}
