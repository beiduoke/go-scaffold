package service

import (
	"context"
	"time"

	v1 "github.com/beiduoke/go-scaffold/api/admin/service/v1"
	"github.com/beiduoke/go-scaffold/api/common"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.UserServiceServer = (*UserService)(nil)

// Service is a  service.
type UserService struct {
	v1.UnimplementedUserServiceServer
	log      *log.Helper
	userCase *biz.UserUsecase
}

// NewService new a  service.
func NewUserService(
	logger log.Logger,
	userCase *biz.UserUsecase,
) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service"))
	return &UserService{
		log:      l,
		userCase: userCase,
	}
}

func TransformUser(data *biz.User) *v1.User {
	var birthday string
	if data.Birthday != nil {
		birthday = data.Birthday.Format(time.DateOnly)
	}
	return &v1.User{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		Avatar:    &data.Avatar,
		NickName:  &data.NickName,
		RealName:  &data.RealName,
		Gender:    &data.Gender,
		Birthday:  &birthday,
		Phone:     &data.Phone,
		Email:     &data.Email,
		State:     &data.State,
	}
}

// ListUser 列表用户
func (s *UserService) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	results, total := s.userCase.ListPage(ctx, pagination.NewPagination(
		pagination.WithPage(in.GetPage()),
		pagination.WithPageSize(in.GetPageSize()),
		pagination.WithQuery(map[string]interface{}{
			"name":     in.GetName(),
			"nickName": in.GetNickName(),
			"deptId":   in.GetDeptId(),
			// 预加载查询
			"preloadPosts": true,
			"preloadRoles": true,
			"preloadDept":  true,
		}),
		pagination.WithOrderBy(map[string]bool{"id": true}),
	))
	return &v1.ListUserResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.User) *v1.User {
			user := TransformUser(v)
			if v.DeptID > 0 {
				deptId := uint64(v.DeptID)
				user.DeptId = &deptId
				user.Dept = TransformDept(v.Dept)
			}
			for _, p := range v.Roles {
				user.RoleIds = append(user.RoleIds, uint64(p.ID))
				user.Roles = append(user.Roles, &v1.Role{Id: uint64(p.ID), Name: p.Name})
			}
			for _, p := range v.Posts {
				user.PostIds = append(user.PostIds, uint64(p.ID))
				user.Posts = append(user.Posts, &v1.Post{Id: uint64(p.ID), Name: p.Name})
			}
			return user
		}),
	}, nil
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	var birthday *time.Time
	if in.GetBirthday() != "" {
		day, err := time.Parse("2006-01-02", in.GetBirthday())
		if err != nil {
			return nil, v1.ErrorUserCreateFail("生日格式错误")
		}
		birthday = &day
	}
	user, err := s.userCase.Create(ctx, &biz.User{
		Name:     in.GetName(),
		Avatar:   in.GetAvatar(),
		NickName: in.GetNickName(),
		RealName: in.GetRealName(),
		Password: in.GetPassword(),
		Birthday: birthday,
		Gender:   int32(in.GetGender()),
		Phone:    in.GetPhone(),
		Email:    in.GetEmail(),
		State:    int32(in.GetState()),
		DeptID:   uint(in.GetDeptId()),
		Roles: func(roleIds []uint64) (bizRoles []*biz.Role) {
			for _, v := range roleIds {
				bizRoles = append(bizRoles, &biz.Role{ID: uint(v)})
			}
			return bizRoles
		}(in.GetRoleIds()),
		Posts: func(postIds []uint64) (bizPosts []*biz.Post) {
			for _, v := range postIds {
				bizPosts = append(bizPosts, &biz.Post{ID: uint(v)})
			}
			return bizPosts
		}(in.GetPostIds()),
	})
	if err != nil {
		return nil, v1.ErrorUserCreateFail("用户创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&common.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreateUserResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateUser 修改用户
func (s *UserService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	id := in.GetId()
	if id <= 0 {
		return nil, v1.ErrorUserIdNull("修改用户ID不能为空")
	}
	v := in.GetData()
	bizUser := biz.User{
		ID:       uint(id),
		Name:     v.GetName(),
		Avatar:   v.GetAvatar(),
		NickName: v.GetNickName(),
		RealName: v.GetRealName(),
		Gender:   int32(v.GetGender()),
		Phone:    v.GetPhone(),
		Email:    v.GetEmail(),
		State:    int32(v.GetState()),
		DeptID:   uint(v.GetDeptId()),
		Roles: func(roleIds []uint64) (bizRoles []*biz.Role) {
			for _, v := range roleIds {
				bizRoles = append(bizRoles, &biz.Role{ID: uint(v)})
			}
			return bizRoles
		}(v.GetRoleIds()),
		Posts: func(postIds []uint64) (bizPosts []*biz.Post) {
			for _, v := range postIds {
				bizPosts = append(bizPosts, &biz.Post{ID: uint(v)})
			}
			return bizPosts
		}(v.GetPostIds()),
	}
	if v.GetBirthday() != "" {
		birthday, err := time.Parse("2006-01-02", v.GetBirthday())
		if err != nil {
			return nil, v1.ErrorUserUpdateFail("生日格式错误")
		}
		bizUser.Birthday = &birthday
	}
	err := s.userCase.Update(ctx, &bizUser)
	if err != nil {
		return nil, v1.ErrorUserUpdateFail("用户修改失败 %v", err)
	}
	return &v1.UpdateUserResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetUser 获取用户
func (s *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.User, error) {
	user, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	return TransformUser(user), nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	if err := s.userCase.Delete(ctx, &biz.User{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorUserDeleteFail("用户删除失败：%v", err)
	}
	return &v1.DeleteUserResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ExistUserByName 用户名是否存在
func (s *UserService) ExistUserByName(ctx context.Context, in *v1.ExistUserByNameRequest) (*v1.ExistUserByNameResponse, error) {
	user, _ := s.userCase.GetName(ctx, &biz.User{Name: in.GetName()})
	handleType, message := constant.HandleType_error.String(), "用户不存在"
	if user != nil && user.ID > 0 {
		handleType = constant.HandleType_success.String()
		message = "用户存在"
	}

	return &v1.ExistUserByNameResponse{
		Type:    handleType,
		Message: message,
	}, nil
}
