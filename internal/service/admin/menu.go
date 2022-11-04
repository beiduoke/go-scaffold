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

func TransformMenu(data *biz.Menu) *v1.Menu {
	return &v1.Menu{
		CreatedAt:  timestamppb.New(data.CreatedAt),
		UpdatedAt:  timestamppb.New(data.UpdatedAt),
		Id:         uint64(data.ID),
		Name:       data.Name,
		Path:       data.Path,
		ParentId:   uint64(data.ParentID),
		Hidden:     protobuf.MenuHidden(data.Hidden),
		Component:  data.Component,
		Sort:       data.Sort,
		Icon:       data.Icon,
		Title:      data.Title,
		KeepAlive:  protobuf.MenuKeepAlive(data.KeepAlive),
		BaseMenu:   protobuf.MenuBaseMenu(data.BaseMenu),
		CloseTab:   protobuf.MenuCloseTab(data.CloseTab),
		Children:   make([]*v1.Menu, 0),
		Parameters: make([]*v1.MenuParameter, 0),
		Buttons:    make([]*v1.MenuButton, 0),
	}
}

// TreeMenu 菜单树形
func TreeMenu(menus []*biz.Menu, pid uint) []*v1.Menu {
	list := make([]*v1.Menu, 0)
	for _, menu := range menus {
		if menu.ParentID == pid {
			m := TransformMenu(menu)
			m.Children = append(m.Children, TreeMenu(menus, menu.ID)...)
			list = append(list, m)
		}
	}
	return list
}

// GetTreeMenu 列表菜单-树形
func (s *AdminService) GetMenuTree(ctx context.Context, in *v1.GetMenuTreeReq) (*v1.GetMenuTreeReply, error) {
	results := s.menuCase.GetTree(ctx, uint(in.GetParentId()))
	return &v1.GetMenuTreeReply{
		Items: TreeMenu(results, uint(in.GetParentId())),
	}, nil
}

// ListMenu 列表菜单
func (s *AdminService) ListMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.menuCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformMenu(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateMenu 创建菜单
func (s *AdminService) CreateMenu(ctx context.Context, in *v1.CreateMenuReq) (*v1.CreateMenuReply, error) {
	parameters, buttons := make([]*biz.MenuParameter, 0, len(in.GetParameters())), make([]*biz.MenuButton, 0, len(in.GetButtons()))
	for _, v := range in.GetParameters() {
		parameters = append(parameters, &biz.MenuParameter{
			Type:  int32(v.GetType()),
			Key:   v.GetKey(),
			Value: v.GetValue(),
		})
	}
	for _, v := range in.GetButtons() {
		buttons = append(buttons, &biz.MenuButton{
			Name:    v.GetName(),
			Remarks: v.GetRemarks(),
		})
	}

	user, err := s.menuCase.Create(ctx, &biz.Menu{
		Name:       in.GetName(),
		Path:       in.GetPath(),
		ParentID:   uint(in.GetParentId()),
		Hidden:     int32(in.GetHidden()),
		Component:  in.GetComponent(),
		Sort:       in.GetSort(),
		Icon:       in.GetIcon(),
		Title:      in.GetTitle(),
		KeepAlive:  int32(in.GetKeepAlive()),
		BaseMenu:   int32(in.GetBaseMenu()),
		CloseTab:   int32(in.GetCloseTab()),
		Parameters: parameters,
		Buttons:    buttons,
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

	parameters, buttons := make([]*biz.MenuParameter, 0, len(v.GetParameters())), make([]*biz.MenuButton, 0, len(v.GetButtons()))
	for _, v := range v.GetParameters() {
		parameters = append(parameters, &biz.MenuParameter{
			Type:  int32(v.GetType()),
			Key:   v.GetKey(),
			Value: v.GetValue(),
		})
	}
	for _, v := range v.GetButtons() {
		buttons = append(buttons, &biz.MenuButton{
			Name:    v.GetName(),
			Remarks: v.GetRemarks(),
		})
	}
	err := s.menuCase.Update(ctx, &biz.Menu{
		ID:         uint(in.GetId()),
		Name:       v.GetName(),
		Path:       v.GetPath(),
		ParentID:   uint(v.GetParentId()),
		Hidden:     int32(v.GetHidden()),
		Component:  v.GetComponent(),
		Sort:       v.GetSort(),
		Icon:       v.GetIcon(),
		Title:      v.GetTitle(),
		KeepAlive:  int32(v.GetKeepAlive()),
		BaseMenu:   int32(v.GetBaseMenu()),
		CloseTab:   int32(v.GetCloseTab()),
		Parameters: parameters,
		Buttons:    buttons,
	})
	if err != nil {
		return nil, v1.ErrorMenuUpdateFail("菜单修改失败: %v", err.Error())
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
	m := TransformMenu(menu)
	for _, v := range menu.Buttons {
		m.Buttons = append(m.Buttons, &v1.MenuButton{
			Name:    v.Name,
			Remarks: v.Remarks,
		})
	}
	for _, v := range menu.Parameters {
		m.Parameters = append(m.Parameters, &v1.MenuParameter{
			Type:  protobuf.MenuParameterType(v.Type),
			Key:   v.Key,
			Value: v.Value,
		})
	}
	return m, nil
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
