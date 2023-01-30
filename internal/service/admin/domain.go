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

func TransformDomain(data *biz.Domain) *v1.Domain {
	return &v1.Domain{
		Id:        uint64(data.ID),
		Name:      data.Name,
		ParentId:  uint64(data.ParentID),
		Sort:      int32(data.Sort),
		State:     protobuf.DomainState(data.State),
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
	}
}

// TreeMenu 部门树形
func TreeDomain(domains []*biz.Domain, pid uint) []*v1.Domain {
	list := make([]*v1.Domain, 0)
	for _, domain := range domains {
		if domain.ParentID == pid {
			m := TransformDomain(domain)
			m.Children = append(m.Children, TreeDomain(domains, domain.ID)...)
			list = append(list, m)
		}
	}
	return list
}

// GetTreeDomain 列表部门-树形
func (s *AdminService) ListDomainTree(ctx context.Context, in *v1.ListDomainTreeReq) (*v1.ListDomainTreeReply, error) {
	results := s.domainCase.GetTree(ctx, uint(in.GetId()))
	return &v1.ListDomainTreeReply{
		Items: TreeDomain(results, uint(in.GetId())),
	}, nil
}

// ListDomain 列表-领域
func (s *AdminService) ListDomain(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.domainCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformDomain(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateDomain 创建领域
func (s *AdminService) CreateDomain(ctx context.Context, in *v1.CreateDomainReq) (*v1.CreateDomainReply, error) {
	user, err := s.domainCase.Create(ctx, &biz.Domain{
		Name:          in.GetName(),
		ParentID:      uint(in.GetParentId()),
		DefaultRoleID: uint(in.GetDefaultRoleId()),
		State:         int32(in.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainCreateFail("领域创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateDomainReply{
		Success: true,
		Message: "创建成功",
		Data:    data,
	}, nil
}

// UpdateDomain 修改领域
func (s *AdminService) UpdateDomain(ctx context.Context, in *v1.UpdateDomainReq) (*v1.UpdateDomainReply, error) {
	v := in.GetData()
	err := s.domainCase.Update(ctx, &biz.Domain{
		ID:            uint(in.GetId()),
		Name:          v.GetName(),
		ParentID:      uint(v.GetParentId()),
		DefaultRoleID: uint(v.GetDefaultRoleId()),
		Sort:          v.GetSort(),
		State:         int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// UpdateDomainState 修改领域-状态
func (s *AdminService) UpdateDomainState(ctx context.Context, in *v1.UpdateDomainStateReq) (*v1.UpdateDomainStateReply, error) {
	v := in.GetData()
	err := s.domainCase.UpdateState(ctx, &biz.Domain{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDomainStateReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetDomain 获取领域
func (s *AdminService) GetDomain(ctx context.Context, in *v1.GetDomainReq) (*v1.Domain, error) {
	domain, err := s.domainCase.GetID(ctx, &biz.Domain{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("领域未找到")
	}
	return TransformDomain(domain), nil
}

// DeleteDomain 删除领域
func (s *AdminService) DeleteDomain(ctx context.Context, in *v1.DeleteDomainReq) (*v1.DeleteDomainReply, error) {
	if err := s.domainCase.Delete(ctx, &biz.Domain{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDomainDeleteFail("领域删除失败：%v", err)
	}
	return &v1.DeleteDomainReply{
		Success: true,
		Message: "删除成功",
	}, nil
}

// ListDomainMenu 获取领域菜单
func (s *AdminService) ListDomainMenu(ctx context.Context, in *v1.ListDomainMenuReq) (*v1.ListDomainMenuReply, error) {
	id := in.GetId()
	menus, _ := s.domainCase.ListMenuByID(ctx, &biz.Domain{ID: uint(id)})
	items := make([]*v1.Menu, 0, len(menus))
	for _, v := range menus {
		items = append(items, TransformMenu(v))
	}
	return &v1.ListDomainMenuReply{Items: items, Total: int32(len(items))}, nil
}

// HandleDomainMenu 处理领域菜单
func (s *AdminService) HandleDomainMenu(ctx context.Context, in *v1.HandleDomainMenuReq) (*v1.HandleDomainMenuReply, error) {
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
	if err := s.domainCase.HandleMenu(ctx, &biz.Domain{ID: uint(in.GetId()), Menus: menus}); err != nil {
		return nil, v1.ErrorDomainHandleMenuFail("领域菜单处理失败：%v", err)
	}
	return &v1.HandleDomainMenuReply{
		Success: true,
		Message: "处理成功",
	}, nil
}
