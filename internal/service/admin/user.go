package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ProfileUser 概括
func (s *AdminService) GetUserProfile(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	name := "jayden"
	return &v1.User{
		Name: name,
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
	results, _ := s.userCase.ListUser(ctx)
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		user := &v1.User{
			Id:        uint64(v.ID),
			Name:      v.Name,
			NickName:  v.NickName,
			RealName:  v.RealName,
			Gender:    protobuf.UserGender(v.Gender),
			Mobile:    v.Mobile,
			Email:     v.Email,
			State:     protobuf.UserState(v.State),
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		}
		if v.Birthday != nil {
			user.Birthday = timestamppb.New(*v.Birthday)
		}
		item, _ := anypb.New(user)
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: 0,
		Items: items,
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
	user, err := s.userCase.GetUserByID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	var birthday *timestamppb.Timestamp
	if user.Birthday != nil {
		birthday = timestamppb.New(*user.Birthday)
	}
	return &v1.User{
		Id:        uint64(user.ID),
		Name:      user.Name,
		NickName:  user.NickName,
		RealName:  user.RealName,
		Birthday:  birthday,
		Gender:    protobuf.UserGender(user.Gender),
		Mobile:    user.Mobile,
		Email:     user.Email,
		State:     protobuf.UserState(user.State),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

// DeleteUser 删除用户
func (s *AdminService) DeleteUser(ctx context.Context, in *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{
		Success: true,
	}, nil
}
