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

func TransformRole(data *biz.Role) *v1.Role {
	return &v1.Role{
		CreatedAt:     timestamppb.New(data.CreatedAt),
		UpdatedAt:     timestamppb.New(data.UpdatedAt),
		Id:            uint64(data.ID),
		Name:          data.Name,
		ParentId:      uint64(data.ParentID),
		DefaultRouter: data.DefaultRouter,
		Sort:          data.Sort,
		State:         protobuf.RoleState(data.State),
		Remarks:       data.Remarks,
	}
}

// ListRole 列表-角色
func (s *AdminService) ListRole(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.roleCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformRole(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: total,
		Items: items,
	}, nil
}

// CreateRole 创建角色
func (s *AdminService) CreateRole(ctx context.Context, in *v1.CreateRoleReq) (*v1.CreateRoleReply, error) {
	user, err := s.roleCase.Create(ctx, &biz.Role{
		Name:          in.GetName(),
		ParentID:      uint(in.GetParentId()),
		DefaultRouter: in.GetDefaultRouter(),
		State:         int32(in.GetState()),
		Remarks:       in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorRoleCreateFail("角色创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateRoleReply{
		Success: true,
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateRole 创建角色
func (s *AdminService) UpdateRole(ctx context.Context, in *v1.UpdateRoleReq) (*v1.UpdateRoleReply, error) {
	v := in.GetData()

	err := s.roleCase.Update(ctx, &biz.Role{
		ID:            uint(in.GetId()),
		Name:          v.GetName(),
		ParentID:      uint(v.GetParentId()),
		DefaultRouter: v.GetDefaultRouter(),
		Sort:          v.GetSort(),
		State:         int32(v.GetState()),
		Remarks:       v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorRoleUpdateFail("角色创建失败: %v", err.Error())
	}
	return &v1.UpdateRoleReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// UpdateRoleState 修改角色-状态
func (s *AdminService) UpdateRoleState(ctx context.Context, in *v1.UpdateRoleStateReq) (*v1.UpdateRoleStateReply, error) {
	v := in.GetData()
	err := s.roleCase.UpdateState(ctx, &biz.Role{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域状态修改失败: %v", err.Error())
	}
	return &v1.UpdateRoleStateReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetRole 获取角色
func (s *AdminService) GetRole(ctx context.Context, in *v1.GetRoleReq) (*v1.Role, error) {
	role, err := s.roleCase.GetID(ctx, &biz.Role{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorRoleNotFound("角色未找到")
	}
	return TransformRole(role), nil
}

// DeleteRole 删除角色
func (s *AdminService) DeleteRole(ctx context.Context, in *v1.DeleteRoleReq) (*v1.DeleteRoleReply, error) {
	if err := s.roleCase.Delete(ctx, &biz.Role{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorRoleDeleteFail("角色删除失败：%v", err)
	}
	return &v1.DeleteRoleReply{
		Success: true,
		Message: "删除成功",
	}, nil
}

// ListRoleMenu 获取角色菜单
func (s *AdminService) ListRoleMenu(ctx context.Context, in *v1.ListRoleMenuReq) (*v1.ListRoleMenuReply, error) {
	id := in.GetId()
	menus, _ := s.roleCase.ListMenuByID(ctx, &biz.Role{ID: uint(id)})
	items := make([]*v1.Menu, 0, len(menus))
	for _, v := range menus {
		items = append(items, TransformMenu(v))
	}
	return &v1.ListRoleMenuReply{Items: items, Total: int64(len(items))}, nil
}

// HandleRoleMenu 处理角色菜单
func (s *AdminService) HandleRoleMenu(ctx context.Context, in *v1.HandleRoleMenuReq) (*v1.HandleRoleMenuReply, error) {
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
	if err := s.roleCase.HandleMenu(ctx, &biz.Role{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorRoleHandleMenuFail("角色菜单处理失败：%v", err)
	}
	return &v1.HandleRoleMenuReply{
		Success: true,
		Message: "处理成功",
	}, nil
}

// HandleRoleResource 处理角色资源
func (s *AdminService) HandleRoleResource(ctx context.Context, in *v1.HandleRoleResourceReq) (*v1.HandleRoleResourceReply, error) {
	inResourceIds := in.GetData().GetResourceIds()
	apiIds := make([]uint, 0, len(inResourceIds))
	for _, v := range inResourceIds {
		apiIds = append(apiIds, uint(v))
	}
	resources, err := s.resourceCase.ListByIDs(ctx, apiIds...)
	if err != nil {
		return nil, v1.ErrorRoleHandleResourceFail("角色资源查询失败")
	}
	if err := s.roleCase.HandleResource(ctx, &biz.Role{ID: uint(in.GetId()), Resources: resources}); err != nil {
		return nil, v1.ErrorRoleHandleResourceFail("角色资源处理失败：%v", err)
	}
	return &v1.HandleRoleResourceReply{
		Success: true,
		Message: "处理成功",
	}, nil
}
