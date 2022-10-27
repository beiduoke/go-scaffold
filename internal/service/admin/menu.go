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

// ListMenu 列表菜单
func (s *AdminService) ListMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.menuCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		user := &v1.Menu{
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
			Id:        uint64(v.ID),
			Name:      v.Name,
			Path:      v.Path,
			ParentId:  uint64(v.ParentID),
			Hidden:    protobuf.MenuHidden(v.Hidden),
			Component: v.Component,
			Sort:      v.Sort,
			Icon:      v.Icon,
			Title:     v.Title,
			KeepAlive: protobuf.MenuKeepAlive(v.KeepAlive),
			BaseMenu:  protobuf.MenuBaseMenu(v.BaseMenu),
			CloseTab:  protobuf.MenuCloseTab(v.CloseTab),
		}
		item, _ := anypb.New(user)
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateMenu 创建菜单
func (s *AdminService) CreateMenu(ctx context.Context, in *v1.CreateMenuReq) (*v1.CreateMenuReply, error) {
	user, err := s.menuCase.Create(ctx, &biz.Menu{
		Name:      in.GetName(),
		Path:      in.GetPath(),
		ParentID:  uint(in.GetParentId()),
		Hidden:    int32(in.GetHidden()),
		Component: in.GetComponent(),
		Sort:      in.GetSort(),
		Icon:      in.GetIcon(),
		Title:     in.GetTitle(),
		KeepAlive: int32(in.GetKeepAlive()),
		BaseMenu:  int32(in.GetBaseMenu()),
		CloseTab:  int32(in.GetCloseTab()),
	})
	if err != nil {
		return nil, v1.ErrorMenuCreateFail("菜单创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateMenuReply{
		Success: true,
		Message: "创建成功",
		Data:    data,
	}, nil
}

// UpdateMenu 创建菜单
func (s *AdminService) UpdateMenu(ctx context.Context, in *v1.UpdateMenuReq) (*v1.UpdateMenuReply, error) {
	v := in.GetData()
	err := s.menuCase.Update(ctx, &biz.Menu{
		ID:        uint(in.GetId()),
		Name:      v.GetName(),
		Path:      v.GetPath(),
		ParentID:  uint(v.GetParentId()),
		Hidden:    int32(v.GetHidden()),
		Component: v.GetComponent(),
		Sort:      v.GetSort(),
		Icon:      v.GetIcon(),
		Title:     v.GetTitle(),
		KeepAlive: int32(v.GetKeepAlive()),
		BaseMenu:  int32(v.GetBaseMenu()),
		CloseTab:  int32(v.GetCloseTab()),
	})
	if err != nil {
		return nil, v1.ErrorMenuUpdateFail("菜单创建失败: %v", err.Error())
	}
	return &v1.UpdateMenuReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetMenu 获取菜单
func (s *AdminService) GetMenu(ctx context.Context, in *v1.GetMenuReq) (*v1.Menu, error) {
	menu, err := s.menuCase.GetID(ctx, &biz.Menu{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorMenuNotFound("菜单未找到")
	}
	return &v1.Menu{
		CreatedAt: timestamppb.New(menu.CreatedAt),
		UpdatedAt: timestamppb.New(menu.UpdatedAt),
		Id:        uint64(menu.ID),
		Name:      menu.Name,
		Path:      menu.Path,
		ParentId:  uint64(menu.ParentID),
		Hidden:    protobuf.MenuHidden(menu.Hidden),
		Component: menu.Component,
		Sort:      menu.Sort,
		Icon:      menu.Icon,
		Title:     menu.Title,
		KeepAlive: protobuf.MenuKeepAlive(menu.KeepAlive),
		BaseMenu:  protobuf.MenuBaseMenu(menu.BaseMenu),
		CloseTab:  protobuf.MenuCloseTab(menu.CloseTab),
	}, nil
}

// DeleteMenu 删除菜单
func (s *AdminService) DeleteMenu(ctx context.Context, in *v1.DeleteMenuReq) (*v1.DeleteMenuReply, error) {
	if err := s.menuCase.Delete(ctx, &biz.Menu{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorMenuDeleteFail("菜单删除失败：%v", err)
	}
	return &v1.DeleteMenuReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
