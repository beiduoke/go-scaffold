package data

import (
	"context"
	"encoding/json"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type AuthorityRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthorityRepo .
func NewAuthorityRepo(logger log.Logger, data *Data) biz.AuthorityRepo {
	return &AuthorityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthorityRepo) toModel(d *biz.Authority) *SysAuthority {
	if d == nil {
		return nil
	}
	sysData := &SysAuthority{
		Name:          d.Name,
		DefaultRouter: d.DefaultRouter,
		State:         d.State,
		ParentID:      d.ParentID,
	}

	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func (r *AuthorityRepo) toBiz(d *SysAuthority) *biz.Authority {
	if d == nil {
		return nil
	}
	return &biz.Authority{
		ID:            d.ID,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
		Name:          d.Name,
		ParentID:      d.ParentID,
		DefaultRouter: d.DefaultRouter,
		Sort:          d.Sort,
		State:         d.State,
	}
}

func (r *AuthorityRepo) Save(ctx context.Context, g *biz.Authority) (*biz.Authority, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *AuthorityRepo) Update(ctx context.Context, g *biz.Authority) (*biz.Authority, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Updates(d)
	return r.toBiz(d), result.Error
}

func (r *AuthorityRepo) FindByName(ctx context.Context, s string) (*biz.Authority, error) {
	authority := SysAuthority{}
	result := r.data.DBD(ctx).Last(&authority, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&authority), nil
}

func (r *AuthorityRepo) FindByID(ctx context.Context, id uint) (*biz.Authority, error) {
	authority := SysAuthority{}
	result := r.data.DBD(ctx).Last(&authority, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&authority), nil
}

func (r *AuthorityRepo) ListByIDs(ctx context.Context, id ...uint) (authorities []*biz.Authority, err error) {
	db := r.data.DBD(ctx).Model(&SysAuthority{})
	sysAuthorities := []*SysAuthority{}

	err = db.Find(&sysAuthorities).Error
	if err != nil {
		return authorities, err
	}
	for _, v := range sysAuthorities {
		authorities = append(authorities, r.toBiz(v))
	}
	return
}

func (r *AuthorityRepo) ListByName(ctx context.Context, name string) ([]*biz.Authority, error) {
	sysAuthorities, bizAuthorities := []*SysAuthority{}, []*biz.Authority{}
	result := r.data.DBD(ctx).Find(&sysAuthorities, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysAuthorities {
		bizAuthorities = append(bizAuthorities, r.toBiz(v))
	}
	return bizAuthorities, nil
}

func (r *AuthorityRepo) Delete(ctx context.Context, g *biz.Authority) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
}

func (r *AuthorityRepo) ListAll(ctx context.Context) ([]*biz.Authority, error) {
	return nil, nil
}

func (r *AuthorityRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (authorities []*biz.Authority, total int64) {
	db := r.data.DBD(ctx).Model(&SysAuthority{}).Debug()
	sysAuthorities := []*SysAuthority{}
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

	result := db.Limit(int(handler.GetPageSize())).Find(&sysAuthorities)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysAuthorities {
		authorities = append(authorities, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(authorities))
	}

	return authorities, total
}

func (r *AuthorityRepo) HandleMenu(ctx context.Context, g *biz.Authority) error {
	sysAuthority := r.toModel(g)
	err := r.data.DB(ctx).Debug().Model(sysAuthority).Association("Menus").Clear()
	if err != nil {
		return err
	}
	sysAuthorityMenus := []SysAuthorityMenu{}
	for _, v := range g.Menus {
		menuButtons := make([]uint, len(v.Buttons))
		for _, m := range v.Buttons {
			menuButtons = append(menuButtons, m.ID)
		}
		buttons, _ := json.Marshal(menuButtons)
		menuParameters := make([]uint, len(v.Parameters))
		for _, m := range v.Parameters {
			menuParameters = append(menuParameters, m.ID)
		}
		parameters, _ := json.Marshal(menuParameters)
		sysAuthorityMenus = append(sysAuthorityMenus, SysAuthorityMenu{
			AuthorityID:   g.ID,
			MenuID:        v.ID,
			MenuButton:    string(buttons),
			MenuParameter: string(parameters),
		})
	}
	return r.data.DB(ctx).Model(&SysAuthorityMenu{}).CreateInBatches(&sysAuthorityMenus, len(g.Menus)).Error
}

func (r *AuthorityRepo) HandleApi(ctx context.Context, g *biz.Authority) error {
	sysAuthority := r.toModel(g)

	apiRepo := ApiRepo{}
	for _, v := range g.Apis {
		sysAuthority.Apis = append(sysAuthority.Apis, *apiRepo.toModel(v))
	}

	return r.data.DB(ctx).Model(sysAuthority).Debug().Association("Api").Replace(g.Apis)
}
