package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServiceServer = (*AdminService)(nil)

func TransformDomain(data *biz.Domain) *v1.Domain {
	return &v1.Domain{
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
		Id:          uint64(data.ID),
		Name:        data.Name,
		ParentId:    uint64(data.ParentID),
		Code:        &data.Code,
		Sort:        &data.Sort,
		Alias:       &data.Alias,
		Logo:        &data.Logo,
		Pic:         &data.Pic,
		Keywords:    &data.Keywords,
		Description: &data.Description,
		State:       &data.State,
		Remarks:     &data.Remarks,
		Children:    make([]*v1.Domain, 0),
	}
}

// GetTreeDomain 列表部门-树形
func (s *AdminService) ListDomainTree(ctx context.Context, in *v1.ListDomainTreeRequest) (*v1.ListDomainTreeResponse, error) {
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
func (s *AdminService) ListDomain(ctx context.Context, in *v1.ListDomainRequest) (*v1.ListDomainResponse, error) {
	results, total := s.domainCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDomainResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(t *biz.Domain) *v1.Domain {
			return TransformDomain(t)
		}),
	}, nil
}

// CreateDomain 创建租户
func (s *AdminService) CreateDomain(ctx context.Context, in *v1.CreateDomainRequest) (*v1.CreateDomainResponse, error) {
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
	})
	if err != nil {
		return nil, v1.ErrorDomainCreateFail("租户创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&v1.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreateDomainResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDomain 修改租户
func (s *AdminService) UpdateDomain(ctx context.Context, in *v1.UpdateDomainRequest) (*v1.UpdateDomainResponse, error) {
	v := in.GetData()
	err := s.domainCase.Update(ctx, &biz.Domain{
		ID:       uint(in.GetId()),
		Name:     v.GetName(),
		ParentID: uint(v.GetParentId()),
		Sort:     v.GetSort(),
		State:    int32(v.GetState()),
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
func (s *AdminService) UpdateDomainState(ctx context.Context, in *v1.UpdateDomainStateRequest) (*v1.UpdateDomainStateResponse, error) {
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
func (s *AdminService) GetDomain(ctx context.Context, in *v1.GetDomainRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetID(ctx, &biz.Domain{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// GetDomainCode 获取租户
func (s *AdminService) GetDomainCode(ctx context.Context, in *v1.GetDomainCodeRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetCode(ctx, &biz.Domain{Code: in.GetCode()})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// GetDomainName 获取租户
func (s *AdminService) GetDomainName(ctx context.Context, in *v1.GetDomainNameRequest) (*v1.Domain, error) {
	domain, err := s.domainCase.GetName(ctx, &biz.Domain{Name: in.GetName()})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("租户未找到")
	}
	return TransformDomain(domain), nil
}

// DeleteDomain 删除租户
func (s *AdminService) DeleteDomain(ctx context.Context, in *v1.DeleteDomainRequest) (*v1.DeleteDomainResponse, error) {
	if err := s.domainCase.Delete(ctx, &biz.Domain{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDomainDeleteFail("租户删除失败：%v", err)
	}
	return &v1.DeleteDomainResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ListDomainMenu 获取租户菜单
func (s *AdminService) ListDomainMenu(ctx context.Context, in *v1.ListDomainMenuRequest) (*v1.ListDomainMenuResponse, error) {
	id := in.GetId()
	menus, _ := s.domainCase.ListMenuByID(ctx, &biz.Domain{ID: uint(id)})
	total := int64(len(menus))
	return &v1.ListDomainMenuResponse{Items: convert.ArrayToAny(menus, func(t *biz.Menu) *v1.Menu {
		return TransformMenu(t)
	}), Total: &total}, nil
}

// HandleDomainMenu 处理租户菜单
func (s *AdminService) HandleDomainMenu(ctx context.Context, in *v1.HandleDomainMenuRequest) (*v1.HandleDomainMenuResponse, error) {
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
