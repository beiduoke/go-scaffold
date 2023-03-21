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

type DomainRoleUser struct {
	DomainID  uint
	RoleID    uint
	UserID    uint
	CreatedAt time.Time
}

// Domain is a Domain model.
type Domain struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          uint
	ParentID    uint
	Name        string
	Code        string
	Title       string
	Keywords    string
	Logo        string
	Pic         string
	Description string
	Sort        int32
	State       int32
	Remarks     string
	Role        *Role
	Menus       []*Menu
}

func (g Domain) GetID() string {
	return convert.UnitToString(g.ID)
}

// DomainRepo is a Greater repo.
type DomainRepo interface {
	Save(context.Context, *Domain) (*Domain, error)
	Update(context.Context, *Domain) (*Domain, error)
	FindByID(context.Context, uint) (*Domain, error)
	FindByCode(context.Context, string) (*Domain, error)
	FindByName(context.Context, string) (*Domain, error)
	ListByIDs(context.Context, ...uint) ([]*Domain, error)
	ListByName(context.Context, string) ([]*Domain, error)
	Delete(context.Context, *Domain) error
	ListAll(context.Context) ([]*Domain, error)
	ListPage(context.Context, *pagination.Pagination) ([]*Domain, int64)
	ListMenuByIDs(context.Context, ...uint) ([]*Menu, error)
	HandleMenu(context.Context, *Domain) error

	// 领域权限
	// stdcasbin.IEnforcer
	// SaveRoleUser 添加领域领域权限
	// SaveRoleForUserInDomain(context.Context, uint /* userID */, uint /* roleID */, uint /* domainID */) error
	// GetRoleForUserInDomain 获取领域用户的领域
	// GetRolesForUserInDomain(context.Context, uint /* userID */, uint /* domainID */) []*Role
	// GetRolesForUserInDomain 获取具有域内领域的用户
	// GetUsersForRoleInDomain(context.Context, uint /* roleID */, uint /* domainID */) []*User
	// DeleteRoleForUserInDomain 域内删除用户的领域域内删除用户的领域
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
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		domain, err := uc.biz.domainRepo.Save(ctx, g)
		if err != nil {
			return err
		}

		g, err = uc.biz.domainRepo.Update(ctx, domain)
		return err
	})
	return g, err
}

// ListByIDs 获取指定领域ID集合
func (uc *DomainUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Domain, err error) {
	// roles, _ = uc.biz.domainRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改领域
func (uc *DomainUsecase) Update(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("UpdateDomain: %v", g)

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

// UpdateState 修改领域状态
func (uc *DomainUsecase) UpdateState(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("UpdateDomainState: %v", g)

	domain, _ := uc.biz.domainRepo.FindByID(ctx, g.ID)
	if domain == nil {
		return errors.New("领域不存在")
	}

	if g.State <= 0 {
		g.State = int32(pb.DomainState_DOMAIN_STATE_ACTIVE)
	}

	domain.State = g.State
	_, err := uc.biz.domainRepo.Update(ctx, domain)
	return err
}

// List 领域列表全部
func (uc *DomainUsecase) ListAll(ctx context.Context) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Debugf("ListAll")
	return uc.biz.domainRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 领域列表分页
func (uc *DomainUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Debugf("DomainPage")
	return uc.biz.domainRepo.ListPage(ctx, paging)
}

// GetID 根据领域ID领域
func (uc *DomainUsecase) GetID(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Debugf("GetDomainID: %v", g)
	return uc.biz.domainRepo.FindByID(ctx, g.ID)
}

// GetID 根据领域Code领域
func (uc *DomainUsecase) GetCode(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Debugf("GetDomainCode: %v", g)
	return uc.biz.domainRepo.FindByCode(ctx, g.Code)
}

// Delete 根据领域ID删除领域
func (uc *DomainUsecase) Delete(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("DeleteDomain: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.biz.domainRepo.Delete(ctx, g); err != nil {
			return err
		}
		_, err := uc.biz.enforcer.DeleteRole(g.GetID())
		return err
	})
}

// HandleMenu 绑定菜单
func (uc *DomainUsecase) HandleMenu(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("HandleMenu: %v", g)
	return uc.biz.domainRepo.HandleMenu(ctx, g)
}

// HandleMenu 获取领域菜单
func (uc *DomainUsecase) ListMenuByID(ctx context.Context, g *Domain) ([]*Menu, error) {
	uc.log.WithContext(ctx).Debugf("ListMenuByIDs: %v", g)
	return uc.biz.domainRepo.ListMenuByIDs(ctx, g.ID)
}

// GetTree 获取领域树形
func (uc *DomainUsecase) GetTree(ctx context.Context, id uint) []*Domain {
	uc.log.WithContext(ctx).Debugf("GetTree")
	menus, _ := uc.biz.domainRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
	return menus
}
