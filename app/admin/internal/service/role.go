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

func TransformRole(data *biz.Role) *v1.Role {
	menuIds := make([]uint64, len(data.Menus))
	for _, v := range data.Menus {
		menuIds = append(menuIds, uint64(v.ID))
	}
	return &v1.Role{
		CreatedAt:     timestamppb.New(data.CreatedAt),
		UpdatedAt:     timestamppb.New(data.UpdatedAt),
		Id:            uint64(data.ID),
		Name:          data.Name,
		ParentId:      uint64(data.ParentID),
		DefaultRouter: &data.DefaultRouter,
		Sort:          &data.Sort,
		State:         &data.State,
		Remarks:       &data.Remarks,
		MenuIds:       menuIds,
	}
}

// ListRole 列表-角色
func (s *AdminService) ListRole(ctx context.Context, in *v1.ListRoleRequest) (*v1.ListRoleResponse, error) {
	results, total := s.roleCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListRoleResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.Role) *v1.Role {
			return TransformRole(v)
		}),
	}, nil
}

// CreateRole 创建角色
func (s *AdminService) CreateRole(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error) {
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
	if len(in.GetMenuIds()) > 0 {
		// 同步角色菜单操作
		if _, err = s.HandleRoleMenu(ctx, &v1.HandleRoleMenuRequest{
			Id:   uint64(role.ID),
			Data: &v1.HandleRoleMenuRequest_Data{MenuIds: in.GetMenuIds()},
		}); err != nil {
			return nil, err
		}
	}

	data, _ := anypb.New(&v1.Result{
		Id: uint64(role.ID),
	})
	return &v1.CreateRoleResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateRole 修改角色
func (s *AdminService) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error) {
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
	if _, err = s.HandleRoleMenu(ctx, &v1.HandleRoleMenuRequest{
		Id:   in.GetId(),
		Data: &v1.HandleRoleMenuRequest_Data{MenuIds: v.GetMenuIds()},
	}); err != nil {
		return nil, err
	}
	return &v1.UpdateRoleResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateRoleState 修改角色-状态
func (s *AdminService) UpdateRoleState(ctx context.Context, in *v1.UpdateRoleStateRequest) (*v1.UpdateRoleStateResponse, error) {
	v := in.GetData()
	err := s.roleCase.UpdateState(ctx, &biz.Role{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorRoleUpdateFail("租户状态修改失败: %v", err.Error())
	}
	return &v1.UpdateRoleStateResponse{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetRole 获取角色
func (s *AdminService) GetRole(ctx context.Context, in *v1.GetRoleRequest) (*v1.Role, error) {
	role, err := s.roleCase.GetID(ctx, &biz.Role{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorRoleNotFound("角色未找到")
	}
	return TransformRole(role), nil
}

// DeleteRole 删除角色
func (s *AdminService) DeleteRole(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error) {
	if err := s.roleCase.Delete(ctx, &biz.Role{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorRoleDeleteFail("角色删除失败：%v", err)
	}
	return &v1.DeleteRoleResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ListRoleMenu 列表-指定ID角色菜单
func (s *AdminService) ListRoleMenu(ctx context.Context, in *v1.ListRoleMenuRequest) (*v1.ListRoleMenuResponse, error) {
	id := in.GetId()
	menus, _ := s.roleCase.ListMenuByID(ctx, &biz.Role{ID: uint(id)})
	return &v1.ListRoleMenuResponse{Items: convert.ArrayToAny(menus, func(t *biz.Menu) *v1.Menu {
		return TransformMenu(t)
	})}, nil
}

// HandleRoleMenu 处理指定ID角色菜单
func (s *AdminService) HandleRoleMenu(ctx context.Context, in *v1.HandleRoleMenuRequest) (*v1.HandleRoleMenuResponse, error) {
	data := in.GetData()
	menus := make([]*biz.Menu, 0, len(data.GetMenuIds()))
	for _, v := range data.GetMenuIds() {
		menus = append(menus, &biz.Menu{
			ID: uint(v),
		})
	}
	if err := s.roleCase.HandleMenu(ctx, &biz.Role{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorRoleHandleMenuFail("角色菜单处理失败：%v", err)
	}
	return &v1.HandleRoleMenuResponse{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}

// ListRoleDept 列表-获取指定ID角色部门
func (s *AdminService) ListRoleDept(ctx context.Context, in *v1.ListRoleDeptRequest) (*v1.ListRoleDeptResponse, error) {
	id := in.GetId()
	menus, _ := s.roleCase.ListDeptByID(ctx, &biz.Role{ID: uint(id)})
	return &v1.ListRoleDeptResponse{Items: convert.ArrayToAny(menus, func(t *biz.Dept) *v1.Dept {
		return TransformDept(t)
	})}, nil
}

// GetRoleDataScope 获取指定ID角色数据范围
func (s *AdminService) GetRoleDataScope(ctx context.Context, in *v1.GetRoleDataScopeRequest) (*v1.GetRoleDataScopeResponse, error) {
	id := in.GetId()
	role, _ := s.roleCase.GetDataScopeByID(ctx, &biz.Role{ID: uint(id)})
	deptCustoms := make([]uint64, 0, len(role.Depts))
	for _, v := range role.Depts {
		deptCustoms = append(deptCustoms, uint64(v.ID))
	}
	return &v1.GetRoleDataScopeResponse{
		Scope:             role.DataScope,
		DeptCheckStrictly: &role.DeptCheckStrictly,
		DeptCustoms:       deptCustoms,
	}, nil
}

// HandleRoleDataScope 处理角色数据
func (s *AdminService) HandleRoleDataScope(ctx context.Context, in *v1.HandleRoleDataScopeRequest) (*v1.HandleRoleDataScopeResponse, error) {
	var depts []*biz.Dept
	if in.Data.GetScope() == int32(v1.RoleDataScope_ROLE_DATA_SCOPE_DEPT_CUSTOM) {
		inDeptCustoms := in.GetData().GetDeptCustoms()
		deptIds := make([]uint, 0, len(inDeptCustoms))
		for _, v := range inDeptCustoms {
			deptIds = append(deptIds, uint(v))
		}
		depts, _ = s.deptCase.ListByIDs(ctx, deptIds...)
	}
	if err := s.roleCase.HandleDataScope(ctx, &biz.Role{ID: uint(in.GetId()), DataScope: int32(in.Data.GetScope()), Depts: depts}); err != nil {
		return nil, v1.ErrorRoleHandleDeptFail("角色资源处理失败：%v", err)
	}
	return &v1.HandleRoleDataScopeResponse{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}
