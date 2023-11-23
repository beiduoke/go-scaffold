package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/service/v1"
	"github.com/beiduoke/go-scaffold/api/common"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/beiduoke/go-scaffold/pkg/util/trans"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.DomainServiceServer = (*DomainService)(nil)

// Service is a  service.
type DomainService struct {
	v1.UnimplementedDomainServiceServer
	log        *log.Helper
	domainCase *biz.DomainUsecase
}

// NewService new a  service.
func NewDomainService(
	logger log.Logger,
	domainCase *biz.DomainUsecase,
) *DomainService {
	l := log.NewHelper(log.With(logger, "module", "domain/service/admin-service"))
	return &DomainService{
		log:        l,
		domainCase: domainCase,
	}
}

func TransformDomain(data *biz.Domain) *v1.Domain {
	return &v1.Domain{
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
		Id:          uint64(data.ID),
		Name:        data.Name,
		ParentId:    uint64(data.ParentID),
		Code:        trans.String(data.Code),
		Sort:        trans.Int32(data.Sort),
		Alias:       trans.String(data.Alias),
		Logo:        trans.String(data.Logo),
		Pic:         trans.String(data.Pic),
		Keywords:    trans.String(data.Keywords),
		Description: trans.String(data.Description),
		State:       trans.Int32(data.State),
		Remarks:     trans.String(data.Remarks),
		Children:    make([]*v1.Domain, 0),
		PackageId:   uint64(data.PackageID),
		Package:     TransformDomainPackage(data.Package),
	}
}

// GetTreeDomain 列表部门-树形
func (s *DomainService) ListDomainTree(ctx context.Context, in *v1.ListDomainTreeRequest) (*v1.ListDomainTreeResponse, error) {
	results, _ := s.domainCase.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrderBy(map[string]bool{"sort": true})))
	treeData := make([]*v1.Domain, 0)
	for _, v := range results {
		treeData = append(treeData, TransformDomain(v))
	}

	return &v1.ListDomainTreeResponse{
		Items: convert.ToTree(treeData, in.GetId(), func(t *v1.Domain, ts ...*v1.Domain) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}

// ListDomain 列表-租户
func (s *DomainService) ListDomain(ctx context.Context, in *v1.ListDomainRequest) (*v1.ListDomainResponse, error) {
	results, total := s.domainCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDomainResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(t *biz.Domain) *v1.Domain {
			return TransformDomain(t)
		}),
	}, nil
}

// CreateDomain 创建租户
func (s *DomainService) CreateDomain(ctx context.Context, in *v1.CreateDomainRequest) (*v1.CreateDomainResponse, error) {
	user, err := s.domainCase.Create(ctx, &biz.Domain{
		ParentID:    uint(in.GetParentId()),
		Name:        in.GetName(),
		Alias:       in.GetAlias(),
		Keywords:    in.GetKeywords(),
		Logo:        in.GetLogo(),
		Pic:         in.GetPic(),
		Description: in.GetDescription(),
		Sort:        in.GetSort(),
		State:       int32(in.GetState()),
		Remarks:     in.GetRemarks(),
		PackageID:   uint(in.GetPackageId()),
	})
	if err != nil {
		return nil, v1.ErrorDomainCreateFail("租户创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&common.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreateDomainResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDomain 修改租户
func (s *DomainService) UpdateDomain(ctx context.Context, in *v1.UpdateDomainRequest) (*v1.UpdateDomainResponse, error) {
	v := in.GetData()
	err := s.domainCase.Update(ctx, &biz.Domain{
		ID:          uint(in.GetId()),
		ParentID:    uint(v.GetParentId()),
		Name:        v.GetName(),
		Alias:       v.GetAlias(),
		Keywords:    v.GetKeywords(),
		Logo:        v.GetLogo(),
		Pic:         v.GetPic(),
		Description: v.GetDescription(),
		Sort:        v.GetSort(),
		State:       int32(v.GetState()),
		Remarks:     v.GetRemarks(),
		PackageID:   uint(v.GetPackageId()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("租户修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDomainState 修改租户-状态
func (s *DomainService) UpdateDomainState(ctx context.Context, in *v1.UpdateDomainStateRequest) (*v1.UpdateDomainStateResponse, error) {
	v := in.GetData()
	err := s.domainCase.UpdateState(ctx, &biz.Domain{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("租户状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainStateResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetDomain 获取租户
func (s *DomainService) GetDomain(ctx context.Context, in *v1.GetDomainRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetID(ctx, &biz.Domain{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// GetDomainCode 获取租户
func (s *DomainService) GetDomainCode(ctx context.Context, in *v1.GetDomainCodeRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetCode(ctx, &biz.Domain{Code: in.GetCode()})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// GetDomainName 获取租户
func (s *DomainService) GetDomainName(ctx context.Context, in *v1.GetDomainNameRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetName(ctx, &biz.Domain{Name: in.GetName()})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// DeleteDomain 删除租户
func (s *DomainService) DeleteDomain(ctx context.Context, in *v1.DeleteDomainRequest) (*v1.DeleteDomainResponse, error) {
	if err := s.domainCase.Delete(ctx, &biz.Domain{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDomainDeleteFail("租户删除失败：%v", err)
	}
	return &v1.DeleteDomainResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ListDomainMenu 获取租户菜单
func (s *DomainService) ListDomainMenu(ctx context.Context, in *v1.ListDomainMenuRequest) (*v1.ListDomainMenuResponse, error) {
	id := in.GetId()
	menus, _ := s.domainCase.ListMenuByID(ctx, &biz.Domain{ID: uint(id)})
	total := int64(len(menus))
	return &v1.ListDomainMenuResponse{Items: convert.ArrayToAny(menus, func(t *biz.Menu) *v1.Menu {
		return TransformMenu(t)
	}), Total: &total}, nil
}

// HandleDomainMenu 处理租户菜单
func (s *DomainService) HandleDomainMenu(ctx context.Context, in *v1.HandleDomainMenuRequest) (*v1.HandleDomainMenuResponse, error) {
	var menus []*biz.Menu
	data := in.GetData()
	for _, v := range data.GetMenuIds() {
		menus = append(menus, &biz.Menu{
			ID: uint(v),
		})
	}
	if err := s.domainCase.HandleMenu(ctx, &biz.Domain{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorDomainHandleMenuFail("租户菜单处理失败：%v", err)
	}
	return &v1.HandleDomainMenuResponse{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}

func TransformDomainPackage(data *biz.DomainPackage) *v1.DomainPackage {
	if data == nil {
		return nil
	}
	return &v1.DomainPackage{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Sort:      &data.Sort,
		State:     &data.State,
		Remarks:   &data.Remarks,
		Name:      data.Name,
		MenuIds: func() (ids []uint64) {
			for _, menu := range data.Menus {
				ids = append(ids, uint64(menu.ID))
			}
			return ids
		}(),
	}
}

// ListDomainPackage 列表-租户套餐
func (s *DomainService) ListDomainPackage(ctx context.Context, in *v1.ListDomainPackageRequest) (*v1.ListDomainPackageResponse, error) {
	results, total := s.domainCase.PackageListPage(ctx, pagination.NewPagination(
		pagination.WithPage(in.GetPage()),
		pagination.WithPageSize(in.GetPageSize()),
		pagination.WithQuery(map[string]interface{}{
			"name":  in.GetName(),
			"state": in.GetState(),
		}),
	))
	return &v1.ListDomainPackageResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(t *biz.DomainPackage) *v1.DomainPackage {
			return TransformDomainPackage(t)
		}),
	}, nil
}

// CreateDomainPackage 创建租户套餐
func (s *DomainService) CreateDomainPackage(ctx context.Context, in *v1.CreateDomainPackageRequest) (*v1.CreateDomainPackageResponse, error) {
	user, err := s.domainCase.PackageCreate(ctx, &biz.DomainPackage{
		Name:    in.GetName(),
		Sort:    in.GetSort(),
		State:   int32(in.GetState()),
		Remarks: in.GetRemarks(),
		Menus: func() (ms []*biz.Menu) {
			for _, id := range in.GetMenuIds() {
				ms = append(ms, &biz.Menu{ID: uint(id)})
			}
			return ms
		}(),
	})
	if err != nil {
		return nil, v1.ErrorDomainCreateFail("租户套餐创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&common.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreateDomainPackageResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDomainPackage 修改租户套餐
func (s *DomainService) UpdateDomainPackage(ctx context.Context, in *v1.UpdateDomainPackageRequest) (*v1.UpdateDomainPackageResponse, error) {
	v := in.GetData()
	err := s.domainCase.PackageUpdate(ctx, &biz.DomainPackage{
		ID:      uint(in.GetId()),
		Name:    v.GetName(),
		Sort:    v.GetSort(),
		State:   int32(v.GetState()),
		Remarks: v.GetRemarks(),
		Menus: func() (ms []*biz.Menu) {
			for _, id := range in.Data.GetMenuIds() {
				ms = append(ms, &biz.Menu{ID: uint(id)})
			}
			return ms
		}(),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("租户套餐修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainPackageResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDomainPackageState 修改租户套餐-状态
func (s *DomainService) UpdateDomainPackageState(ctx context.Context, in *v1.UpdateDomainPackageStateRequest) (*v1.UpdateDomainPackageStateResponse, error) {
	v := in.GetData()
	err := s.domainCase.PackageUpdateState(ctx, &biz.DomainPackage{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("租户套餐状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainPackageStateResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// DeleteDomainPackage 删除租户套餐
func (s *DomainService) DeleteDomainPackage(ctx context.Context, in *v1.DeleteDomainPackageRequest) (*v1.DeleteDomainPackageResponse, error) {
	if err := s.domainCase.Delete(ctx, &biz.Domain{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDomainDeleteFail("租户套餐删除失败：%v", err)
	}
	return &v1.DeleteDomainPackageResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}
