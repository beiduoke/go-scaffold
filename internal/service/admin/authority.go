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
	}
}

// ListAuthority 列表-权限角色
func (s *AdminService) ListAuthority(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.authorityCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		user := &v1.Authority{
			Id:        uint64(v.ID),
			Name:      v.Name,
			ParentId:  uint64(v.ParentID),
			State:     protobuf.AuthorityState(v.State),
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		}
		item, _ := anypb.New(user)
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
	})
	if err != nil {
		return nil, v1.ErrorAuthorityUpdateFail("权限角色创建失败: %v", err.Error())
	}
	return &v1.UpdateAuthorityReply{
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
	return &v1.Authority{
		Id:        uint64(authority.ID),
		Name:      authority.Name,
		ParentId:  uint64(authority.ParentID),
		Sort:      int32(authority.Sort),
		State:     protobuf.AuthorityState(authority.State),
		CreatedAt: timestamppb.New(authority.CreatedAt),
		UpdatedAt: timestamppb.New(authority.UpdatedAt),
	}, nil
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
			ID:         uint(v.GetMenuId()),
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

// HandleAuthorityApi 处理权限角色接口
func (s *AdminService) HandleAuthorityApi(ctx context.Context, in *v1.HandleAuthorityApiReq) (*v1.HandleAuthorityApiReply, error) {
	inApiIds := in.GetData().GetApiIds()
	apiIds := make([]uint, 0, len(inApiIds))
	for _, v := range inApiIds {
		apiIds = append(apiIds, uint(v))
	}
	apis, err := s.apiCase.ListByIDs(ctx, apiIds...)
	if err != nil {
		return nil, v1.ErrorAuthorityHandleApiFail("权限角色接口查询失败")
	}
	if err := s.authorityCase.HandleApi(ctx, &biz.Authority{ID: uint(in.GetId()), Apis: apis}); err != nil {
		return nil, v1.ErrorAuthorityHandleApiFail("权限角色接口处理失败：%v", err)
	}
	return &v1.HandleAuthorityApiReply{
		Success: true,
		Message: "处理成功",
	}, nil
}
