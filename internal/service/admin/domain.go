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

// ListDomain 列表-领域
func (s *AdminService) ListDomain(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.domainCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		user := &v1.Domain{
			Id:        uint64(v.ID),
			Name:      v.Name,
			ParentId:  uint64(v.ParentID),
			State:     protobuf.DomainState(v.State),
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

// CreateDomain 创建领域
func (s *AdminService) CreateDomain(ctx context.Context, in *v1.CreateDomainReq) (*v1.CreateDomainReply, error) {
	user, err := s.domainCase.Create(ctx, &biz.Domain{
		Name:               in.GetName(),
		ParentID:           uint(in.GetParentId()),
		DefaultAuthorityID: uint(in.GetDefaultAuthorityId()),
		State:              int32(in.GetState()),
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

// UpdateDomain 创建领域
func (s *AdminService) UpdateDomain(ctx context.Context, in *v1.UpdateDomainReq) (*v1.UpdateDomainReply, error) {
	v := in.GetData()
	err := s.domainCase.Update(ctx, &biz.Domain{
		ID:                 uint(in.GetId()),
		Name:               v.GetName(),
		ParentID:           uint(v.GetParentId()),
		DefaultAuthorityID: uint(v.GetDefaultAuthorityId()),
		Sort:               v.GetSort(),
		State:              int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDomainUpdateFail("领域创建失败: %v", err.Error())
	}
	return &v1.UpdateDomainReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetDomain 获取领域
func (s *AdminService) GetDomain(ctx context.Context, in *v1.GetDomainReq) (*v1.Domain, error) {
	authority, err := s.domainCase.GetID(ctx, &biz.Domain{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDomainNotFound("领域未找到")
	}
	return &v1.Domain{
		Id:        uint64(authority.ID),
		Name:      authority.Name,
		ParentId:  uint64(authority.ParentID),
		Sort:      int32(authority.Sort),
		State:     protobuf.DomainState(authority.State),
		CreatedAt: timestamppb.New(authority.CreatedAt),
		UpdatedAt: timestamppb.New(authority.UpdatedAt),
	}, nil
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
