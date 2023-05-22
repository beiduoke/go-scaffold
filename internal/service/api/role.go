package api

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/internal/pkg/proto"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

func TransformRole(data *biz.Role) *v1.Role {
	menus := make([]uint64, len(data.Menus))
	for _, v := range data.Menus {
		menus = append(menus, uint64(v.ID))
	}
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
		Menus:         menus,
	}
}

// ListRole 列表-角色
func (s *ApiService) ListRole(ctx context.Context, in *v1.ListRoleReq) (*v1.ListRoleReply, error) {
	results, total := s.roleCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformRole(v))
		items = append(items, item)
	}
	return &v1.ListRoleReply{
		Total: total,
		Items: items,
	}, nil
}

// CreateRole 创建角色
func (s *ApiService) CreateRole(ctx context.Context, in *v1.CreateRoleReq) (*v1.CreateRoleReply, error) {
	role, err := s.roleCase.Create(ctx, &biz.Role{
		Name:              in.GetName(),
		ParentID:          uint(in.GetParentId()),
		DefaultRouter:     in.GetDefaultRouter(),
		Sort:              in.GetSort(),
		DataScope:         int32(in.GetDataScope()),
		MenuCheckStrictly: int32(in.GetMenuCheckStrictly()),
		DeptCheckStrictly: int32(in.GetDeptCheckStrictly()),
		State:             int32(in.GetState()),
		Remarks:           in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorRoleCreateFail("角色创建失败: %v", err.Error())
	}
	if len(in.GetMenus()) > 0 {
		// 同步角色菜单操作
		if _, err = s.HandleRoleMenu(ctx, &v1.HandleRoleMenuReq{
			Id:   uint64(role.ID),
			Data: &v1.HandleRoleMenuReq_Data{Menus: in.GetMenus()},
		}); err != nil {
			return nil, err
		}
	}

	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(role.ID),
	})
	return &v1.CreateRoleReply{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateRole 修改角色
func (s *ApiService) UpdateRole(ctx context.Context, in *v1.UpdateRoleReq) (*v1.UpdateRoleReply, error) {
	v := in.GetData()
	err := s.roleCase.Update(ctx, &biz.Role{
		ID:                uint(in.GetId()),
		Name:              v.GetName(),
		ParentID:          uint(v.GetParentId()),
		DefaultRouter:     v.GetDefaultRouter(),
		Sort:              v.GetSort(),
		DataScope:         int32(v.GetDataScope()),
		MenuCheckStrictly: int32(v.GetMenuCheckStrictly()),
		DeptCheckStrictly: int32(v.GetDeptCheckStrictly()),
		State:             int32(v.GetState()),
		Remarks:           v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorRoleUpdateFail("角色修改失败: %v", err.Error())
	}
	// 同步角色菜单操作
	if _, err = s.HandleRoleMenu(ctx, &v1.HandleRoleMenuReq{
		Id:   in.GetId(),
		Data: &v1.HandleRoleMenuReq_Data{Menus: v.GetMenus()},
	}); err != nil {
		return nil, err
	}
	return &v1.UpdateRoleReply{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateRoleState 修改角色-状态
func (s *ApiService) UpdateRoleState(ctx context.Context, in *v1.UpdateRoleStateReq) (*v1.UpdateRoleStateReply, error) {
	v := in.GetData()
	err := s.roleCase.UpdateState(ctx, &biz.Role{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域状态修改失败: %v", err.Error())
	}
	return &v1.UpdateRoleStateReply{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetRole 获取角色
func (s *ApiService) GetRole(ctx context.Context, in *v1.GetRoleReq) (*v1.Role, error) {
	role, err := s.roleCase.GetID(ctx, &biz.Role{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorRoleNotFound("角色未找到")
	}
	return TransformRole(role), nil
}

// DeleteRole 删除角色
func (s *ApiService) DeleteRole(ctx context.Context, in *v1.DeleteRoleReq) (*v1.DeleteRoleReply, error) {
	if err := s.roleCase.Delete(ctx, &biz.Role{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorRoleDeleteFail("角色删除失败：%v", err)
	}
	return &v1.DeleteRoleReply{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ListRoleMenu 列表-指定ID角色菜单
func (s *ApiService) ListRoleMenu(ctx context.Context, in *v1.ListRoleMenuReq) (*v1.ListRoleMenuReply, error) {
	id := in.GetId()
	menus, _ := s.roleCase.ListMenuByID(ctx, &biz.Role{ID: uint(id)})
	return &v1.ListRoleMenuReply{Items: proto.ToAny(menus, func(t *biz.Menu) protoreflect.ProtoMessage {
		return TransformMenu(t)
	})}, nil
}

// HandleRoleMenu 处理指定ID角色菜单
func (s *ApiService) HandleRoleMenu(ctx context.Context, in *v1.HandleRoleMenuReq) (*v1.HandleRoleMenuReply, error) {
	data := in.GetData()
	menus := make([]*biz.Menu, len(data.GetMenus()))
	for _, v := range data.GetMenus() {
		// 暂不使用扩展菜单权限按钮以及参数配置
		// parameters, buttons := make([]*biz.MenuParameter, 0, len(v.GetMenuParameterIds())), make([]*biz.MenuButton, 0, len(v.GetMenuButtonIds()))
		// for _, v := range v.GetMenuParameterIds() {
		// 	parameters = append(parameters, &biz.MenuParameter{ID: uint(v)})
		// }
		// for _, v := range v.GetMenuButtonIds() {
		// 	buttons = append(buttons, &biz.MenuButton{ID: uint(v)})
		// }
		menus = append(menus, &biz.Menu{
			ID: uint(v),
			// Parameters: parameters,
			// Buttons:    buttons,
		})
	}
	if err := s.roleCase.HandleMenu(ctx, &biz.Role{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorRoleHandleMenuFail("角色菜单处理失败：%v", err)
	}
	return &v1.HandleRoleMenuReply{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}

// ListRoleResource 列表-指定ID角色资源
func (s *ApiService) ListRoleResource(ctx context.Context, in *v1.ListRoleResourceReq) (*v1.ListRoleResourceReply, error) {
	id := in.GetId()
	resources, _ := s.roleCase.ListResourceByID(ctx, &biz.Role{ID: uint(id)})
	return &v1.ListRoleResourceReply{Items: proto.ToAny(resources, func(t *biz.Resource) protoreflect.ProtoMessage {
		return TransformResource(t)
	})}, nil
}

// HandleRoleResource 处理指定ID角色资源
func (s *ApiService) HandleRoleResource(ctx context.Context, in *v1.HandleRoleResourceReq) (*v1.HandleRoleResourceReply, error) {
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
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}

// ListRoleDept 列表-获取指定ID角色部门
func (s *ApiService) ListRoleDept(ctx context.Context, in *v1.ListRoleDeptReq) (*v1.ListRoleDeptReply, error) {
	id := in.GetId()
	menus, _ := s.roleCase.ListDeptByID(ctx, &biz.Role{ID: uint(id)})
	return &v1.ListRoleDeptReply{Items: proto.ToAny(menus, func(t *biz.Dept) protoreflect.ProtoMessage {
		return TransformDept(t)
	})}, nil
}

// GetRoleDataScope 获取指定ID角色数据范围
func (s *ApiService) GetRoleDataScope(ctx context.Context, in *v1.GetRoleDataScopeReq) (*v1.GetRoleDataScopeReply, error) {
	id := in.GetId()
	role, _ := s.roleCase.GetDataScopeByID(ctx, &biz.Role{ID: uint(id)})
	deptCustoms := make([]uint64, len(role.Depts))
	for _, v := range role.Depts {
		deptCustoms = append(deptCustoms, uint64(v.ID))
	}
	return &v1.GetRoleDataScopeReply{
		Scope:             protobuf.RoleDataScope(role.DataScope),
		DeptCheckStrictly: (*protobuf.RoleDeptCheckStrictly)(&role.DeptCheckStrictly),
		DeptCustoms:       deptCustoms,
	}, nil
}

// HandleRoleDataScope 处理角色数据
func (s *ApiService) HandleRoleDataScope(ctx context.Context, in *v1.HandleRoleDataScopeReq) (*v1.HandleRoleDataScopeReply, error) {
	inDeptCustoms := in.GetData().GetDeptCustoms()
	deptIds := make([]uint, 0, len(inDeptCustoms))
	for _, v := range inDeptCustoms {
		deptIds = append(deptIds, uint(v))
	}
	depts, err := s.deptCase.ListByIDs(ctx, deptIds...)
	if err != nil {
		return nil, v1.ErrorRoleHandleDeptFail("角色资源查询失败")
	}
	if err := s.roleCase.HandleDataScope(ctx, &biz.Role{ID: uint(in.GetId()), DataScope: int32(in.Data.GetScope()), Depts: depts}); err != nil {
		return nil, v1.ErrorRoleHandleDeptFail("角色资源处理失败：%v", err)
	}
	return &v1.HandleRoleDataScopeReply{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}
