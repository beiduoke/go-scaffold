package admin

import (
	"context"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ProfileUser 概括
func (s *AdminService) GetUserProfile(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	user, err := s.userCase.GetID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	var birthday string
	if user.Birthday != nil {
		birthday = user.Birthday.Format("2006-01-02")
	}
	return &v1.User{
		Id:        uint64(user.ID),
		Name:      user.Name,
		Avatar:    user.Avatar,
		NickName:  user.NickName,
		RealName:  user.RealName,
		Gender:    protobuf.UserGender(user.Gender),
		Birthday:  birthday,
		Mobile:    user.Mobile,
		Email:     user.Email,
		State:     protobuf.UserState(user.State),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

// ProfileUser 概括
func (s *AdminService) GetUserMenu(ctx context.Context, in *emptypb.Empty) (*v1.GetUserMenuReply, error) {
	// id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	name := "菜单"
	return &v1.GetUserMenuReply{
		Name: name,
	}, nil
}

// ListUser 列表用户
func (s *AdminService) ListUser(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, _ := s.userCase.List(ctx)
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		user := &v1.User{
			Id:        uint64(v.ID),
			Name:      v.Name,
			Avatar:    v.Avatar,
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
			user.Birthday = v.Birthday.Format("2006-01-02")
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
	v := in.GetData()
	domainIds := make([]uint, 0, len(v.GetDomainIds()))
	for _, domainId := range v.GetDomainIds() {
		domainIds = append(domainIds, uint(domainId))
	}

	domains, _ := s.domainCase.ListByIDs(ctx, domainIds...)
	var birthday *time.Time
	if v.GetBirthday() != "" {
		day, err := time.Parse("2006-01-02", v.GetBirthday())
		if err != nil {
			return nil, v1.ErrorUserCreateFail("生日格式错误")
		}
		birthday = &day
	}
	_, err := s.userCase.Create(ctx, &biz.User{
		Name:     v.GetName(),
		Avatar:   v.GetAvatar(),
		Password: v.GetPassword(),
		Gender:   int32(v.GetGender()),
		NickName: v.GetNickName(),
		RealName: v.GetRealName(),
		Birthday: birthday,
		Mobile:   v.GetMobile(),
		Email:    v.GetEmail(),
		State:    int32(v.GetState()),
		Domains:  domains,
	})
	if err != nil {
		return nil, v1.ErrorUserCreateFail("用户创建失败: %v", err.Error())
	}
	return &v1.CreateUserReply{
		Success: true,
	}, nil
}

// UpdateUser 修改用户
func (s *AdminService) UpdateUser(ctx context.Context, in *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	id := in.GetId()
	if id <= 0 {
		return nil, v1.ErrorUserIdNull("修改用户ID不能为空")
	}
	v := in.GetData()
	var birthday *time.Time
	if v.GetBirthday() != "" {
		day, err := time.Parse("2006-01-02", v.GetBirthday())
		if err != nil {
			return nil, v1.ErrorUserUpdateFail("生日格式错误")
		}
		birthday = &day
	}
	err := s.userCase.Update(ctx, &biz.User{
		ID:       uint(id),
		Name:     v.GetName(),
		Avatar:   v.GetAvatar(),
		NickName: v.GetNickName(),
		RealName: v.GetRealName(),
		Password: v.GetPassword(),
		Birthday: birthday,
		Gender:   int32(v.GetGender()),
		Mobile:   v.GetMobile(),
		Email:    v.GetEmail(),
		State:    int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorUserUpdateFail("用户修改失败 %v", err)
	}
	return &v1.UpdateUserReply{
		Success: true,
	}, nil
}

// GetUser 获取用户
func (s *AdminService) GetUser(ctx context.Context, in *v1.GetUserReq) (*v1.User, error) {
	user, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	var birthday string
	if user.Birthday != nil {
		birthday = user.Birthday.Format("2006-01-02")
	}
	return &v1.User{
		Id:        uint64(user.ID),
		Name:      user.Name,
		Avatar:    user.Avatar,
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
