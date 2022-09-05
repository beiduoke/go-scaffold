package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// Login 登录
func (s *AdminService) ProfileUser(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	name := "jayden"
	return &v1.User{
		Name: &name,
	}, nil
}
