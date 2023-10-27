package biz

import (
	"context"
	"time"

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
	CreatedAt   time.Time      `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	ID          uint           `json:"id,omitempty" form:"id,omitempty"`
	ParentID    uint           `json:"parentId,omitempty" form:"parentId,omitempty"`
	PackageID   uint           `json:"packageId,omitempty" form:"packageId,omitempty"`
	Name        string         `json:"name,omitempty" form:"name,omitempty"`
	Code        string         `json:"code,omitempty" form:"code,omitempty"`
	Alias       string         `json:"alias,omitempty" form:"alias,omitempty"`
	Keywords    string         `json:"keywords,omitempty" form:"keywords,omitempty"`
	Logo        string         `json:"logo,omitempty" form:"logo,omitempty"`
	Pic         string         `json:"pic,omitempty" form:"pic,omitempty"`
	Description string         `json:"description,omitempty" form:"description,omitempty"`
	Sort        int32          `json:"sort,omitempty" form:"sort,omitempty"`
	State       int32          `json:"state,omitempty" form:"state,omitempty"`
	Remarks     string         `json:"remarks,omitempty" form:"remarks,omitempty"`
	Role        *Role          `json:"role,omitempty" form:"role,omitempty"`
	Menus       []*Menu        `json:"menus,omitempty" form:"menus,omitempty"`
	Package     *DomainPackage `json:"package,omitempty" form:"package,omitempty"`
}

func (g Domain) GetID() string {
	return convert.UnitToString(g.ID)
}

// DomainPackage is a Post model.
type DomainPackage struct {
	CreatedAt time.Time `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	ID        uint      `json:"id,omitempty" form:"id,omitempty"`
	Menus     []*Menu   `json:"menus,omitempty" form:"menus,omitempty"`
	Remarks   string    `json:"remarks,omitempty" form:"remarks,omitempty"`
	Name      string    `json:"name,omitempty" form:"name,omitempty"`
	Sort      int32     `json:"sort,omitempty" form:"sort,omitempty"`
	State     int32     `json:"state,omitempty" form:"state,omitempty"`
}

func (g DomainPackage) GetID() string {
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

	// 租户套餐
	PackageSave(context.Context, *DomainPackage) (*DomainPackage, error)
	PackageUpdate(context.Context, *DomainPackage) (*DomainPackage, error)
	PackageFindByID(context.Context, uint) (*DomainPackage, error)
	PackageDelete(context.Context, *DomainPackage) error
	PackageListAll(context.Context) ([]*DomainPackage, error)
	PackageListPage(context.Context, *pagination.Pagination) ([]*DomainPackage, int64)
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

// ListByIDs 获取指定租户ID集合
func (uc *DomainUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Domain, err error) {
	// roles, _ = uc.biz.domainRepo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改租户
func (uc *DomainUsecase) Update(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("UpdateDomain: %v", g)

	domain, _ := uc.biz.domainRepo.FindByID(ctx, g.ID)
	if domain == nil {
		return errors.New("租户未注册")
	}

	if domain.Name != g.Name && g.Name != "" {
		name, _ := uc.biz.domainRepo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("租户名已存在")
		}
	}
	// 新数据合并到源数据
	if err := mergo.Merge(domain, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}

	_, err := uc.biz.domainRepo.Update(ctx, domain)
	return err
}

// UpdateState 修改租户状态
func (uc *DomainUsecase) UpdateState(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("UpdateDomainState: %v", g)

	domain, _ := uc.biz.domainRepo.FindByID(ctx, g.ID)
	if domain == nil {
		return errors.New("租户不存在")
	}

	domain.State = g.State
	_, err := uc.biz.domainRepo.Update(ctx, domain)
	return err
}

// List 租户列表全部
func (uc *DomainUsecase) ListAll(ctx context.Context) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Debugf("ListAll")
	return uc.biz.domainRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 租户列表分页
func (uc *DomainUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Domain, int64) {
	uc.log.WithContext(ctx).Debugf("DomainPage")
	return uc.biz.domainRepo.ListPage(ctx, paging)
}

// GetID 根据租户ID租户
func (uc *DomainUsecase) GetID(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Debugf("GetDomainID: %v", g)
	return uc.biz.domainRepo.FindByID(ctx, g.ID)
}

// GetID 根据租户Code租户
func (uc *DomainUsecase) GetCode(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Debugf("GetDomainCode: %v", g)
	return uc.biz.domainRepo.FindByCode(ctx, g.Code)
}

// GetID 根据租户Name租户
func (uc *DomainUsecase) GetName(ctx context.Context, g *Domain) (*Domain, error) {
	uc.log.WithContext(ctx).Debugf("GetDomainName: %v", g)
	return uc.biz.domainRepo.FindByName(ctx, g.Name)
}

// Delete 根据租户ID删除租户
func (uc *DomainUsecase) Delete(ctx context.Context, g *Domain) error {
	uc.log.WithContext(ctx).Debugf("DeleteDomain: %v", g)
	return uc.biz.domainRepo.Delete(ctx, g)
}

// HandleMenu 绑定菜单
func (uc *DomainUsecase) HandleMenu(ctx context.Context, g *Domain) error {
	// uc.log.WithContext(ctx).Debugf("HandleMenu: %v", g)
	return uc.biz.domainRepo.HandleMenu(ctx, g)
}

// HandleMenu 获取租户菜单
func (uc *DomainUsecase) ListMenuByID(ctx context.Context, g *Domain) ([]*Menu, error) {
	uc.log.WithContext(ctx).Debugf("ListMenuByIDs: %v", g)
	return uc.biz.domainRepo.ListMenuByIDs(ctx, g.ID)
}

// GetTree 获取租户树形
func (uc *DomainUsecase) GetTree(ctx context.Context, id uint) []*Domain {
	uc.log.WithContext(ctx).Debugf("GetTree")
	menus, _ := uc.biz.domainRepo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
	return menus
}
