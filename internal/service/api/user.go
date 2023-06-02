package api

import (
	"context"
	"strings"
	"time"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/internal/pkg/proto"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

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
		Gender:    (*protobuf.UserGender)(&data.Gender),
		Birthday:  &birthday,
		Phone:     &data.Phone,
		Email:     &data.Email,
		State:     (*protobuf.UserState)(&data.State),
	}
}

// ListUser 列表用户
func (s *ApiService) ListUser(ctx context.Context, in *v1.ListUserReq) (*v1.ListUserReply, error) {
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
	return &v1.ListUserReply{
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
func (s *ApiService) CreateUser(ctx context.Context, in *v1.CreateUserReq) (*v1.CreateUserReply, error) {
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
		Password: in.GetPassword(),
		Gender:   int32(in.GetGender()),
		NickName: in.GetNickName(),
		RealName: in.GetRealName(),
		Birthday: birthday,
		Phone:    in.GetPhone(),
		Email:    in.GetEmail(),
		State:    int32(in.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorUserCreateFail("用户创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateUserReply{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateUser 修改用户
func (s *ApiService) UpdateUser(ctx context.Context, in *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
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
		Password: v.GetPassword(),
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
	return &v1.UpdateUserReply{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetUser 获取用户
func (s *ApiService) GetUser(ctx context.Context, in *v1.GetUserReq) (*v1.User, error) {
	user, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	return TransformUser(user), nil
}

// DeleteUser 删除用户
func (s *ApiService) DeleteUser(ctx context.Context, in *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	if err := s.userCase.Delete(ctx, &biz.User{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorUserDeleteFail("用户删除失败：%v", err)
	}
	return &v1.DeleteUserReply{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

// ExistUserName 用户名是否存在
func (s *ApiService) ExistUserName(ctx context.Context, in *v1.ExistUserNameReq) (*v1.ExistUserNameReply, error) {
	user, _ := s.userCase.GetName(ctx, &biz.User{Name: in.GetName()})
	handleType, message := constant.HandleType_error.String(), "用户不存在"
	if user != nil && user.ID > 0 {
		handleType = constant.HandleType_success.String()
		message = "用户存在"
	}

	return &v1.ExistUserNameReply{
		Type:    handleType,
		Message: message,
	}, nil
}

// HandleUserRole 绑定用户权限
func (s *ApiService) HandleUserRole(ctx context.Context, in *v1.HandleUserRoleReq) (*v1.HandleUserRoleReply, error) {
	v := in.GetData()
	_, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	roleIds := make([]uint, 0, len(v.GetRoleIds()))
	for _, roleId := range v.GetRoleIds() {
		roleIds = append(roleIds, uint(roleId))
	}
	roles, _ := s.roleCase.ListByIDs(ctx, roleIds...)
	err = s.userCase.HandleRole(ctx, &biz.User{
		ID:    uint(in.GetId()),
		Roles: roles,
	})
	if err != nil {
		return nil, v1.ErrorUserHandleRoleFail("绑定用户权限失败: %v", err.Error())
	}
	return &v1.HandleUserRoleReply{
		Type:    constant.HandleType_success.String(),
		Message: "处理成功",
	}, nil
}

// GetUserInfo 用户详情
func (s *ApiService) GetUserInfo(ctx context.Context, in *emptypb.Empty) (*v1.GetUserInfoReply, error) {
	user, err := s.userCase.Info(ctx)
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	return &v1.GetUserInfoReply{
		Name:     user.Name,
		NickName: user.NickName,
		RealName: user.RealName,
		Birthday: func() string {
			if user.Birthday != nil {
				return user.Birthday.Format(time.DateOnly)
			}
			return ""
		}(),
		Gender: protobuf.UserGender(user.Gender),
		Phone:  user.Phone,
		Email:  user.Email,
		Avatar: user.Avatar,
		State:  protobuf.UserState(user.State),
		Roles: func(roles []*biz.Role) (userRoles []*v1.GetUserInfoReply_UserRole) {
			for _, v := range roles {
				userRoles = append(userRoles, &v1.GetUserInfoReply_UserRole{
					Id:            uint64(v.ID),
					Name:          v.Name,
					DefaultRouter: v.DefaultRouter,
					Sort:          v.Sort,
				})
			}
			return userRoles
		}(user.Roles),
	}, nil
}

// ListUserRole 用户角色
func (s *ApiService) ListUserRole(ctx context.Context, in *emptypb.Empty) (*v1.ListUserRoleReply, error) {
	roleModels, err := s.userCase.Roles(ctx)
	if err != nil {
		return nil, v1.ErrorUserRoleFindFail("用户角色失败 %v", err)
	}
	roles := make([]*v1.Role, 0, len(roleModels))
	for _, v := range roleModels {
		roles = append(roles, TransformRole(v))
	}
	return &v1.ListUserRoleReply{
		Items: roles,
	}, nil
}

// 获取角色菜单路由树形列表
func (s *ApiService) ListUserRoleMenuRouterTree(ctx context.Context, in *v1.ListUserRoleMenuRouterTreeReq) (*v1.ListUserRoleMenuRouterTreeReply, error) {
	roleMenus, err := s.userCase.RoleMenus(ctx)
	if err != nil {
		s.log.Debugf("用户菜单查询失败 %v", err)
	}
	treeData := make([]*v1.MenuRouter, 0)
	for _, v := range roleMenus {
		if v.Type == int32(protobuf.MenuType_MENU_TYPE_ABILITY) {
			continue
		}
		treeData = append(treeData, TransformMenuRouter(v))
	}
	return &v1.ListUserRoleMenuRouterTreeReply{
		Items: proto.ToTree(treeData, in.GetMenuParentId(), func(t *v1.MenuRouter, ts ...*v1.MenuRouter) error {
			t.Children = append(t.Children, ts...)
			if len(ts) > 0 && !ts[0].GetMeta().GetHideMenu() {
				path, child := t.Path, ts[0]
				if !strings.HasPrefix(path, "/") {
					path = "/" + path
				}
				redirect := path + "/" + child.Path
				t.Redirect = &redirect
			}
			return nil
		}),
	}, nil
}

// 获取角色权限列表
func (s *ApiService) ListUserRolePermission(ctx context.Context, in *v1.ListUserRolePermissionReq) (*v1.ListUserRolePermissionReply, error) {
	menuModels, _ := s.userCase.RolePermissions(ctx)
	return &v1.ListUserRolePermissionReply{
		Items: convert.ArrayStrUnique(menuModels),
	}, nil
}

// 获取角色菜单树形列表
func (s *ApiService) ListUserRoleMenuTree(ctx context.Context, in *v1.ListUserRoleMenuTreeReq) (*v1.ListUserRoleMenuTreeReply, error) {
	results, err := s.userCase.RoleMenus(ctx)
	if err != nil {
		s.log.Debugf("用户菜单查询失败 %v", err)
	}
	treeData := make([]*v1.Menu, 0)
	for _, v := range results {
		if v.Type == int32(protobuf.MenuType_MENU_TYPE_ABILITY) {
			continue
		}
		treeData = append(treeData, TransformMenu(v))
	}
	return &v1.ListUserRoleMenuTreeReply{
		Items: proto.ToTree(treeData, in.GetMenuParentId(), func(t *v1.Menu, ts ...*v1.Menu) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}
