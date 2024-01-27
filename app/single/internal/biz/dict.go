package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold-single/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// Dict is a Dict model.
type Dict struct {
	CreatedAt time.Time `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	ID        uint      `json:"id,omitempty" form:"id,omitempty"`
	Type      string    `json:"type,omitempty" form:"type,omitempty"`
	Remarks   string    `json:"remarks,omitempty" form:"remarks,omitempty"`
	Name      string    `json:"name,omitempty" form:"name,omitempty"`
	Sort      int32     `json:"sort,omitempty" form:"sort,omitempty"`
	State     int32     `json:"state,omitempty" form:"state,omitempty"`
}

// DictData is a DictData model.
type DictData struct {
	CreatedAt time.Time `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	ID        uint      `json:"id,omitempty" form:"id,omitempty"`
	Label     string    `json:"label,omitempty" form:"label,omitempty"`
	Value     string    `json:"value,omitempty" form:"value,omitempty"`
	ColorType string    `json:"colorType,omitempty" form:"colorType,omitempty"`
	CssClass  string    `json:"cssClass,omitempty" form:"cssClass,omitempty"`
	DictType  string    `json:"dictType,omitempty" form:"dictType,omitempty"`
	Remarks   string    `json:"remarks,omitempty" form:"remarks,omitempty"`
	Sort      int32     `json:"sort,omitempty" form:"sort,omitempty"`
	State     int32     `json:"state,omitempty" form:"state,omitempty"`
}

// DictRepo is a Greater repo.
type DictRepo interface {
	Save(context.Context, *Dict) (*Dict, error)
	Update(context.Context, *Dict) (*Dict, error)
	FindByID(context.Context, uint) (*Dict, error)
	FindByType(context.Context, string) (*Dict, error)
	FindByName(context.Context, string) (*Dict, error)
	ListByIDs(context.Context, ...uint) ([]*Dict, error)
	ListByName(context.Context, string) ([]*Dict, error)
	Delete(context.Context, *Dict) error
	ListAll(context.Context) ([]*Dict, error)
	ListPage(context.Context, *pagination.Pagination) ([]*Dict, int64)

	// 字典数据
	DataSave(context.Context, *DictData) (*DictData, error)
	DataUpdate(context.Context, *DictData) (*DictData, error)
	DataFindByID(context.Context, uint) (*DictData, error)
	DataFindByDictType(context.Context, string) (*DictData, error)
	DataFindByLabel(context.Context, string) (*DictData, error)
	DataListByIDs(context.Context, ...uint) ([]*DictData, error)
	DataListByLabel(context.Context, string) ([]*DictData, error)
	DataDelete(context.Context, *DictData) error
	DataListAll(context.Context) ([]*DictData, error)
	DataListPage(context.Context, *pagination.Pagination) ([]*DictData, int64)
}

// DictUsecase is a Dict usecase.
type DictUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo DictRepo
}

// NewDictUsecase new a Dict usecase.
func NewDictUsecase(logger log.Logger, biz *Biz, repo DictRepo) *DictUsecase {
	return &DictUsecase{biz: biz, log: log.NewHelper(logger), repo: repo}
}

// Create creates a Dict, and returns the new Dict.
func (uc *DictUsecase) Create(ctx context.Context, g *Dict) (*Dict, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		_, err := uc.repo.Save(ctx, g)
		return err
	})
	return g, err
}

// ListByIDs 获取指定字典ID集合
func (uc *DictUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Dict, err error) {
	// roles, _ = uc.repo.ListPage(ctx, noop.NewPagination(noop.WithNopaging(), noop.WithCondition("id in ?", id)))
	return
}

// Update 修改字典
func (uc *DictUsecase) Update(ctx context.Context, g *Dict) error {
	uc.log.WithContext(ctx).Debugf("UpdateDict: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("字典未创建")
	}

	if post.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("字典名已存在")
		}
	}

	fmt.Printf("%s", g.Remarks)
	_, err := uc.repo.Update(ctx, g)
	return err
}

// UpdateState 修改字典状态
func (uc *DictUsecase) UpdateState(ctx context.Context, g *Dict) error {
	uc.log.WithContext(ctx).Debugf("UpdateDictState: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("字典不存在")
	}

	post.State = g.State
	_, err := uc.repo.Update(ctx, post)
	return err
}

// List 字典列表全部
func (uc *DictUsecase) ListAll(ctx context.Context) ([]*Dict, int64) {
	uc.log.WithContext(ctx).Debugf("DictList")
	return uc.repo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 字典列表分页
func (uc *DictUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Dict, int64) {
	uc.log.WithContext(ctx).Debugf("DictPage")
	return uc.repo.ListPage(ctx, paging)
}

// GetID 根据角色ID字典
func (uc *DictUsecase) GetID(ctx context.Context, g *Dict) (*Dict, error) {
	uc.log.WithContext(ctx).Debugf("GetDictID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除字典
func (uc *DictUsecase) Delete(ctx context.Context, g *Dict) error {
	uc.log.WithContext(ctx).Debugf("DeleteDict: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		return uc.repo.Delete(ctx, g)
	})
}

// Create creates a Dict, and returns the new Dict.
func (uc *DictUsecase) DataCreate(ctx context.Context, g *DictData) (*DictData, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Label)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		_, err := uc.repo.DataSave(ctx, g)
		return err
	})
	return g, err
}

// ListByIDs 获取指定字典ID集合
func (uc *DictUsecase) DataListByIDs(ctx context.Context, id ...uint) (r []*DictData, err error) {
	return
}

// Update 修改字典
func (uc *DictUsecase) DataUpdate(ctx context.Context, g *DictData) error {
	uc.log.WithContext(ctx).Debugf("UpdateDict: %v", g)

	v, _ := uc.repo.DataFindByID(ctx, g.ID)
	if v == nil {
		return errors.New("字典未创建")
	}

	if v.Label != g.Label && g.Label != "" {
		name, _ := uc.repo.DataFindByLabel(ctx, g.Label)
		if name != nil {
			return errors.New("字典名已存在")
		}
	}

	fmt.Printf("%s", g.Remarks)
	_, err := uc.repo.DataUpdate(ctx, g)
	return err
}

// UpdateState 修改字典状态
func (uc *DictUsecase) DataUpdateState(ctx context.Context, g *DictData) error {
	uc.log.WithContext(ctx).Debugf("UpdateDictDataState: %v", g)

	post, _ := uc.repo.DataFindByID(ctx, g.ID)
	if post == nil {
		return errors.New("字典不存在")
	}

	post.State = g.State
	_, err := uc.repo.DataUpdate(ctx, post)
	return err
}

// List 字典列表全部
func (uc *DictUsecase) DataListAll(ctx context.Context) ([]*DictData, int64) {
	uc.log.WithContext(ctx).Debugf("DictDataList")
	return uc.repo.DataListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 字典列表分页
func (uc *DictUsecase) DataListPage(ctx context.Context, paging *pagination.Pagination) ([]*DictData, int64) {
	uc.log.WithContext(ctx).Debugf("DictDataPage")
	return uc.repo.DataListPage(ctx, paging)
}

// GetID 根据角色ID字典
func (uc *DictUsecase) DataGetID(ctx context.Context, g *DictData) (*DictData, error) {
	uc.log.WithContext(ctx).Debugf("GetDictDataID: %v", g)
	return uc.repo.DataFindByID(ctx, g.ID)
}

// Delete 根据角色ID删除字典
func (uc *DictUsecase) DataDelete(ctx context.Context, g *DictData) error {
	uc.log.WithContext(ctx).Debugf("DeleteDictData: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		return uc.repo.DataDelete(ctx, g)
	})
}
