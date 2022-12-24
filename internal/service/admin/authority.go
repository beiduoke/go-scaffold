package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

func TransformAuthority(data *biz.Authority) *v1.Authority {
	return &v1.Authority{
		CreatedAt:     timestamppb.New(data.CreatedAt),
		UpdatedAt:     timestamppb.New(data.UpdatedAt),
		Id:            uint64(data.ID),
		Name:          data.Name,
		ParentId:      uint64(data.ParentID),
		DefaultRouter: data.DefaultRouter,
		Sort:          data.Sort,
		State:         protobuf.AuthorityState(data.State),
		Remarks:       data.Remarks,
	}
}

// ListAuthority 列表-权限角色
func (s *AdminService) ListAuthority(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.authorityCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformAuthority(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateAuthority 创建权限角色
func (s *AdminService) CreateAuthority(ctx context.Context, in *v1.CreateAuthorityReq) (*v1.CreateAuthorityReply, error) {
	user, err := s.authorityCase.Create(ctx, &biz.Authority{
		Name:          in.GetName(),
		ParentID:      uint(in.GetParentId()),
		DefaultRouter: in.GetDefaultRouter(),
		State:         int32(in.GetState()),
		Remarks:       in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorAuthorityCreateFail("权限角色创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateAuthorityReply{
		Success: true,
		Message: "创建成功",
		Data:    data,
	}, nil
}

// UpdateAuthority 创建权限角色
func (s *AdminService) UpdateAuthority(ctx context.Context, in *v1.UpdateAuthorityReq) (*v1.UpdateAuthorityReply, error) {
	v := in.GetData()

	err := s.authorityCase.Update(ctx, &biz.Authority{
		ID:            uint(in.GetId()),
		Name:          v.GetName(),
		ParentID:      uint(v.GetParentId()),
		DefaultRouter: v.GetDefaultRouter(),
		Sort:          v.GetSort(),
		State:         int32(v.GetState()),
		Remarks:       v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorAuthorityUpdateFail("权限角色创建失败: %v", err.Error())
	}
	return &v1.UpdateAuthorityReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// UpdateAuthorityState 修改权限角色-状态
func (s *AdminService) UpdateAuthorityState(ctx context.Context, in *v1.UpdateAuthorityStateReq) (*v1.UpdateAuthorityStateReply, error) {
	v := in.GetData()
	err := s.authorityCase.UpdateState(ctx, &biz.Authority{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域状态修改失败: %v", err.Error())
	}
	return &v1.UpdateAuthorityStateReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetAuthority 获取权限角色
func (s *AdminService) GetAuthority(ctx context.Context, in *v1.GetAuthorityReq) (*v1.Authority, error) {
	authority, err := s.authorityCase.GetID(ctx, &biz.Authority{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorAuthorityNotFound("权限角色未找到")
	}
	return TransformAuthority(authority), nil
}

// DeleteAuthority 删除权限角色
func (s *AdminService) DeleteAuthority(ctx context.Context, in *v1.DeleteAuthorityReq) (*v1.DeleteAuthorityReply, error) {
	if err := s.authorityCase.Delete(ctx, &biz.Authority{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorAuthorityDeleteFail("权限角色删除失败：%v", err)
	}
	return &v1.DeleteAuthorityReply{
		Success: true,
		Message: "删除成功",
	}, nil
}

// ListAuthorityMenu 获取权限角色菜单
func (s *AdminService) ListAuthorityMenu(ctx context.Context, in *v1.ListAuthorityMenuReq) (*v1.ListAuthorityMenuReply, error) {
	id := in.GetId()
	menus, _ := s.authorityCase.ListMenuByID(ctx, &biz.Authority{ID: uint(id)})
	items := make([]*v1.Menu, 0, len(menus))
	for _, v := range menus {
		items = append(items, TransformMenu(v))
	}
	return &v1.ListAuthorityMenuReply{Items: items, Total: int32(len(items))}, nil
}

// HandleAuthorityMenu 处理权限角色菜单
func (s *AdminService) HandleAuthorityMenu(ctx context.Context, in *v1.HandleAuthorityMenuReq) (*v1.HandleAuthorityMenuReply, error) {
	var menus []*biz.Menu
	data := in.GetData()
	for _, v := range data.GetMenus() {
		parameters, buttons := make([]*biz.MenuParameter, 0, len(v.GetMenuParameterIds())), make([]*biz.MenuButton, 0, len(v.GetMenuButtonIds()))
		for _, v := range v.GetMenuParameterIds() {
			parameters = append(parameters, &biz.MenuParameter{ID: uint(v)})
		}
		for _, v := range v.GetMenuButtonIds() {
			buttons = append(buttons, &biz.MenuButton{ID: uint(v)})
		}
		menus = append(menus, &biz.Menu{
			ID:         uint(v.GetId()),
			Parameters: parameters,
			Buttons:    buttons,
		})
	}
	if err := s.authorityCase.HandleMenu(ctx, &biz.Authority{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorAuthorityHandleMenuFail("权限角色菜单处理失败：%v", err)
	}
	return &v1.HandleAuthorityMenuReply{
		Success: true,
		Message: "处理成功",
	}, nil
}

// HandleAuthorityResource 处理权限角色资源
func (s *AdminService) HandleAuthorityResource(ctx context.Context, in *v1.HandleAuthorityResourceReq) (*v1.HandleAuthorityResourceReply, error) {
	inResourceIds := in.GetData().GetResourceIds()
	apiIds := make([]uint, 0, len(inResourceIds))
	for _, v := range inResourceIds {
		apiIds = append(apiIds, uint(v))
	}
	resources, err := s.resourceCase.ListByIDs(ctx, apiIds...)
	if err != nil {
		return nil, v1.ErrorAuthorityHandleResourceFail("权限角色资源查询失败")
	}
	if err := s.authorityCase.HandleResource(ctx, &biz.Authority{ID: uint(in.GetId()), Resources: resources}); err != nil {
		return nil, v1.ErrorAuthorityHandleResourceFail("权限角色资源处理失败：%v", err)
	}
	return &v1.HandleAuthorityResourceReply{
		Success: true,
		Message: "处理成功",
	}, nil
}
