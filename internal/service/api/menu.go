package api

import (
	"context"
	"strings"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/proto"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

// 将菜单转换路由树形结构
func TransformMenuRouter(menu *biz.Menu) *v1.MenuRouter {
	meta := &v1.MenuRouter_Meta{
		// 路由title  一般必填
		Title: menu.Title,
		// 图标，也是菜单图标
		Icon: &menu.Icon,
		// 菜单排序，只对第一级有效
		OrderNo: &menu.Sort,
	}

	// 是否隐藏
	hidden := (menu.IsHidden == int32(protobuf.MenuHidden_MENU_HIDDEN_YES))
	if hidden {
		meta.HideMenu = &hidden
	}

	// 当前激活的菜单。用于配置详情页时左侧激活的菜单路径
	if parent := menu.Parent; parent != nil && hidden {
		currentActiveMenu := parent.Path
		for {
			if !strings.HasPrefix(currentActiveMenu, "/") {
				currentActiveMenu = "/" + currentActiveMenu
			}
			parent = parent.Parent
			if parent == nil {
				break
			}
			currentActiveMenu = parent.Path + currentActiveMenu
		}
		meta.CurrentActiveMenu = &currentActiveMenu
	}
	// 忽略缓存
	if cache := menu.IsCache == int32(protobuf.MenuCache_MENU_CACHE_NO); cache {
		meta.IgnoreKeepAlive = &cache
	}
	// 实体组件路径
	component := menu.Component
	if menu.LinkType == int32(protobuf.MenuLinkType_MENU_LINK_TYPE_IFRAME) {
		meta.FrameSrc = &menu.LinkUrl
	}
	id := uint64(menu.ID)
	parentId := uint64(menu.ParentID)
	return &v1.MenuRouter{
		Name:      menu.Name,
		Path:      menu.Path,
		Component: component,
		Redirect:  menu.Redirect,
		Meta:      meta,
		Children:  make([]*v1.MenuRouter, 0),
		Id:        &id,
		ParentId:  &parentId,
	}
}

func TransformMenu(data *biz.Menu) *v1.Menu {
	return &v1.Menu{
		CreatedAt:  timestamppb.New(data.CreatedAt),
		UpdatedAt:  timestamppb.New(data.UpdatedAt),
		Id:         uint64(data.ID),
		Name:       data.Name,
		Type:       protobuf.MenuType(data.Type),
		ParentId:   uint64(data.ParentID),
		Path:       data.Path,
		IsHidden:   protobuf.MenuHidden(data.IsHidden),
		Component:  data.Component,
		Permission: data.Permission,
		Sort:       data.Sort,
		Icon:       data.Icon,
		Title:      data.Title,
		IsCache:    protobuf.MenuCache(data.IsCache),
		IsAffix:    protobuf.MenuAffix(data.IsAffix),
		LinkType:   protobuf.MenuLinkType(data.LinkType),
		LinkUrl:    data.LinkUrl,
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

// ListMenuTree 列表菜单-树形
func (s *ApiService) ListMenuTree(ctx context.Context, in *v1.ListMenuTreeReq) (*v1.ListMenuTreeReply, error) {
	results, total := s.menuCase.ListAll(ctx)
	treeData := make([]*v1.Menu, 0)
	for _, v := range results {
		treeData = append(treeData, TransformMenu(v))
	}
	return &v1.ListMenuTreeReply{
		Items: proto.ToTree(treeData, in.GetId(), func(t *v1.Menu, ts ...*v1.Menu) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
		Total: &total,
	}, nil
}

// ListMenu 列表菜单
func (s *ApiService) ListMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.menuCase.ListPage(ctx, &pagination.Pagination{Page: in.GetPage(), PageSize: in.GetPageSize(), Nopaging: in.GetNopaging(), Query: in.GetQuery(), OrderBy: in.GetOrderBy()})
	return &protobuf.PagingReply{
		Total: total,
		Items: proto.ToAny(results, func(t *biz.Menu) protoreflect.ProtoMessage {
			return TransformMenu(t)
		}),
	}, nil
}

// CreateMenu 创建菜单
func (s *ApiService) CreateMenu(ctx context.Context, in *v1.CreateMenuReq) (*v1.CreateMenuReply, error) {
	parameters, buttons := make([]*biz.MenuParameter, 0, len(in.GetParameters())), make([]*biz.MenuButton, 0, len(in.GetButtons()))
	for _, v := range in.GetParameters() {
		parameters = append(parameters, &biz.MenuParameter{
			Type:  int32(v.GetType()),
			Name:  v.GetName(),
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
		Type:       int32(in.GetType()),
		Path:       in.GetPath(),
		ParentID:   uint(in.GetParentId()),
		Component:  in.GetComponent(),
		Permission: in.GetPermission(),
		Sort:       in.GetSort(),
		Icon:       in.GetIcon(),
		Title:      in.GetTitle(),
		IsHidden:   int32(in.GetIsHidden()),
		IsCache:    int32(in.GetIsCache()),
		IsAffix:    int32(in.GetIsAffix()),
		LinkType:   int32(in.GetLinkType()),
		LinkUrl:    in.GetLinkUrl(),
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
		Result:  data,
	}, nil
}

// UpdateMenu 创建菜单
func (s *ApiService) UpdateMenu(ctx context.Context, in *v1.UpdateMenuReq) (*v1.UpdateMenuReply, error) {
	v := in.GetData()

	parameters, buttons := make([]*biz.MenuParameter, 0, len(v.GetParameters())), make([]*biz.MenuButton, 0, len(v.GetButtons()))
	for _, v := range v.GetParameters() {
		parameters = append(parameters, &biz.MenuParameter{
			Type:  int32(v.GetType()),
			Name:  v.GetName(),
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
		Type:       int32(v.Type),
		Path:       v.GetPath(),
		ParentID:   uint(v.GetParentId()),
		Component:  v.GetComponent(),
		Permission: v.GetPermission(),
		Sort:       v.GetSort(),
		Icon:       v.GetIcon(),
		Title:      v.GetTitle(),
		IsHidden:   int32(v.GetIsHidden()),
		IsCache:    int32(v.GetIsCache()),
		IsAffix:    int32(v.GetIsAffix()),
		LinkType:   int32(v.GetLinkType()),
		LinkUrl:    v.GetLinkUrl(),
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
func (s *ApiService) GetMenu(ctx context.Context, in *v1.GetMenuReq) (*v1.Menu, error) {
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
			Name:  v.Name,
			Value: v.Value,
		})
	}
	return m, nil
}

// DeleteMenu 删除菜单
func (s *ApiService) DeleteMenu(ctx context.Context, in *v1.DeleteMenuReq) (*v1.DeleteMenuReply, error) {
	if err := s.menuCase.Delete(ctx, &biz.Menu{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorMenuDeleteFail("菜单删除失败：%v", err)
	}
	return &v1.DeleteMenuReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
