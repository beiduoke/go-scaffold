package api

import (
	"context"

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

func TransformDomain(data *biz.Domain) *v1.Domain {
	return &v1.Domain{
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
		Id:          uint64(data.ID),
		Name:        data.Name,
		ParentId:    uint64(data.ParentID),
		Code:        data.Code,
		Sort:        int32(data.Sort),
		Title:       data.Title,
		Logo:        data.Logo,
		Pic:         data.Pic,
		Keywords:    data.Keywords,
		Description: data.Description,
		State:       protobuf.DomainState(data.State),
		Remarks:     data.Remarks,
		Children:    make([]*v1.Domain, 0),
	}
}

// GetTreeDomain 列表部门-树形
func (s *ApiService) ListDomainTree(ctx context.Context, in *v1.ListDomainTreeReq) (*v1.ListDomainTreeReply, error) {
	results, _ := s.domainCase.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithOrderBy(map[string]bool{"sort": true})))

	treeData := make([]*v1.Domain, 0)
	for _, v := range results {
		treeData = append(treeData, TransformDomain(v))
	}

	return &v1.ListDomainTreeReply{
		Items: proto.ToTree(treeData, in.GetId(), func(t *v1.Domain, ts ...*v1.Domain) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}

// ListDomain 列表-领域
func (s *ApiService) ListDomain(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.domainCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize()), pagination.WithQuery(pagination.QueryUnmarshal(in.GetQuery())), pagination.WithOrderBy(in.GetOrderBy())))
	return &protobuf.PagingReply{
		Total: total,
		Items: proto.ToAny(results, func(t *biz.Domain) protoreflect.ProtoMessage {
			return TransformDomain(t)
		}),
	}, nil
}

// CreateDomain 创建领域
func (s *ApiService) CreateDomain(ctx context.Context, in *v1.CreateDomainReq) (*v1.CreateDomainReply, error) {
	user, err := s.domainCase.Create(ctx, &biz.Domain{
		ParentID:    uint(in.GetParentId()),
		Name:        in.GetName(),
		Title:       in.GetTitle(),
		Keywords:    in.GetKeywords(),
		Logo:        in.GetLogo(),
		Pic:         in.GetPic(),
		Description: in.GetDescription(),
		Sort:        in.GetSort(),
		State:       int32(in.GetState()),
		Remarks:     in.GetRemarks(),
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
		Result:  data,
	}, nil
}

// UpdateDomain 修改领域
func (s *ApiService) UpdateDomain(ctx context.Context, in *v1.UpdateDomainReq) (*v1.UpdateDomainReply, error) {
	v := in.GetData()
	err := s.domainCase.Update(ctx, &biz.Domain{
		ID:       uint(in.GetId()),
		Name:     v.GetName(),
		ParentID: uint(v.GetParentId()),
		Sort:     v.GetSort(),
		State:    int32(v.GetState()),
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
func (s *ApiService) UpdateDomainState(ctx context.Context, in *v1.UpdateDomainStateReq) (*v1.UpdateDomainStateReply, error) {
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
func (s *ApiService) GetDomain(ctx context.Context, in *v1.GetDomainReq) (*v1.Domain, error) {
	domain, err := s.domainCase.GetID(ctx, &biz.Domain{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("领域未找到")
	}
	return TransformDomain(domain), nil
}

// DeleteDomain 删除领域
func (s *ApiService) DeleteDomain(ctx context.Context, in *v1.DeleteDomainReq) (*v1.DeleteDomainReply, error) {
	if err := s.domainCase.Delete(ctx, &biz.Domain{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDomainDeleteFail("领域删除失败：%v", err)
	}
	return &v1.DeleteDomainReply{
		Success: true,
		Message: "删除成功",
	}, nil
}

// ListDomainMenu 获取领域菜单
func (s *ApiService) ListDomainMenu(ctx context.Context, in *v1.ListDomainMenuReq) (*v1.ListDomainMenuReply, error) {
	id := in.GetId()
	menus, _ := s.domainCase.ListMenuByID(ctx, &biz.Domain{ID: uint(id)})
	total := int64(len(menus))
	return &v1.ListDomainMenuReply{Items: proto.ToAny(menus, func(t *biz.Menu) protoreflect.ProtoMessage {
		return TransformMenu(t)
	}), Total: &total}, nil
}

// HandleDomainMenu 处理领域菜单
func (s *ApiService) HandleDomainMenu(ctx context.Context, in *v1.HandleDomainMenuReq) (*v1.HandleDomainMenuReply, error) {
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
