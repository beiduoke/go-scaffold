package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ProfileUser 概括
func (s *AdminService) GetUserProfile(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	name := "jayden"
	return &v1.User{
		Name: &name,
	}, nil
}

// ProfileUser 概括
func (s *AdminService) GetUserMenu(ctx context.Context, in *emptypb.Empty) (*v1.GetUserMenuReply, error) {
	name := "菜单"
	return &v1.GetUserMenuReply{
		Name: name,
	}, nil
}

// ListUser 列表用户
func (s *AdminService) ListUser(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return &protobuf.PagingReply{
		Total: 0,
		Items: []*anypb.Any{},
	}, nil
}

// CreateUser 创建用户
func (s *AdminService) CreateUser(ctx context.Context, in *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	return &v1.CreateUserReply{
		Success: true,
	}, nil
}

// UpdateUser 修改用户
func (s *AdminService) UpdateUser(ctx context.Context, in *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{
		Success: true,
	}, nil
}

// GetUser 获取用户
func (s *AdminService) GetUser(ctx context.Context, in *v1.GetUserReq) (*v1.User, error) {
	id := in.GetId()
	return &v1.User{
		Id: &id,
	}, nil
}

// DeleteUser 删除用户
func (s *AdminService) DeleteUser(ctx context.Context, in *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{
		Success: true,
	}, nil
}
