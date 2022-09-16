package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
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
