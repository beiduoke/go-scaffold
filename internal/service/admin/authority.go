package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ListAuthority 列表-授权
func (s *AdminService) ListAuthority(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return &protobuf.PagingReply{
		Total: 0,
		Items: []*anypb.Any{},
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
