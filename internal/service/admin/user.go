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

func TransformUser(data *biz.User) *v1.User {
	var birthday string
	if data.Birthday != nil {
		birthday = data.Birthday.Format("2006-01-02")
	}
	return &v1.User{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		Avatar:    data.Avatar,
		NickName:  data.NickName,
		RealName:  data.RealName,
		Gender:    protobuf.UserGender(data.Gender),
		Birthday:  birthday,
		Phone:     data.Phone,
		Email:     data.Email,
		State:     protobuf.UserState(data.State),
	}
}

// ListUser 列表用户
func (s *AdminService) ListUser(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.userCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformUser(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: total,
		Items: items,
	}, nil
}

// CreateUser 创建用户
func (s *AdminService) CreateUser(ctx context.Context, in *v1.CreateUserReq) (*v1.CreateUserReply, error) {
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
		Success: true,
		Message: "创建成功",
		Result:  data,
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
		Phone:    v.GetPhone(),
		Email:    v.GetEmail(),
		State:    int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorUserUpdateFail("用户修改失败 %v", err)
	}
	return &v1.UpdateUserReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetUser 获取用户
func (s *AdminService) GetUser(ctx context.Context, in *v1.GetUserReq) (*v1.User, error) {
	user, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	return TransformUser(user), nil
}

// DeleteUser 删除用户
func (s *AdminService) DeleteUser(ctx context.Context, in *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	if err := s.userCase.Delete(ctx, &biz.User{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorUserDeleteFail("用户删除失败：%v", err)
	}
	return &v1.DeleteUserReply{
		Success: true,
		Message: "删除成功",
	}, nil
}

// ExistUserName 用户名是否存在
func (s *AdminService) ExistUserName(ctx context.Context, in *v1.ExistUserNameReq) (*v1.ExistUserNameReply, error) {
	user, _ := s.userCase.GetName(ctx, &biz.User{Name: in.GetName()})
	exist, message := false, "用户不存在"
	if user != nil && user.ID > 0 {
		exist = true
		message = "用户存在"
	}

	return &v1.ExistUserNameReply{
		Success: exist,
		Message: message,
	}, nil
}

// HandleUserDomain 绑定用户领域
func (s *AdminService) HandleUserDomain(ctx context.Context, in *v1.HandleUserDomainReq) (*v1.HandleUserDomainReply, error) {
	v := in.GetData()
	_, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	domainIds := make([]uint, 0, len(v.GetDomainIds()))
	for _, domainId := range v.GetDomainIds() {
		domainIds = append(domainIds, uint(domainId))
	}

	domains, _ := s.domainCase.ListByIDs(ctx, domainIds...)
	err = s.userCase.HandleDomain(ctx, &biz.User{
		ID:      uint(in.GetId()),
		Domains: domains,
	})
	if err != nil {
		return nil, v1.ErrorUserHandleDomainFail("绑定用户领域失败: %v", err.Error())
	}
	return &v1.HandleUserDomainReply{
		Success: true,
		Message: "处理成功",
	}, nil
}

// HandleUserRole 绑定用户权限
func (s *AdminService) HandleUserDomainRole(ctx context.Context, in *v1.HandleUserDomainRoleReq) (*v1.HandleUserDomainRoleReply, error) {
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
	err = s.userCase.HandleDomainRole(ctx, &biz.User{
		ID:    uint(in.GetId()),
		Roles: roles,
	}, uint(v.GetDomainId()))
	if err != nil {
		return nil, v1.ErrorUserHandleDomainRoleFail("绑定用户权限失败: %v", err.Error())
	}
	return &v1.HandleUserDomainRoleReply{
		Success: true,
		Message: "处理成功",
	}, nil
}

// GetUserInfo 用户	详情
func (s *AdminService) GetUserInfo(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	user, err := s.userCase.GetID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	return TransformUser(user), nil
}

// GetUserProfile 用户概括
func (s *AdminService) GetUserProfile(ctx context.Context, in *emptypb.Empty) (*v1.GetUserProfileReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	user, err := s.userCase.GetID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	roleResult, err := s.ListUserRole(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserProfileReply{
		User:  TransformUser(user),
		Roles: roleResult.GetItems(),
	}, nil
}

// ListUserDomain 用户领域
func (s *AdminService) ListUserDomain(ctx context.Context, in *emptypb.Empty) (*v1.ListUserDomainReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	domainModels, err := s.userCase.ListDomainAll(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserDomainFindFail("用户领域失败 %v", err)
	}
	domains := make([]*v1.Domain, 0, len(domainModels))
	for _, v := range domainModels {
		domains = append(domains, TransformDomain(v))
	}
	return &v1.ListUserDomainReply{
		Items: domains,
	}, nil
}

// ListUserRole 用户角色
func (s *AdminService) ListUserRole(ctx context.Context, in *emptypb.Empty) (*v1.ListUserRoleReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	domainId := convert.StringToUint(authz.ParseFromContext(ctx).GetDomain())
	roleModels, err := s.userCase.ListRoleAll(ctx, &biz.User{ID: id, Domains: []*biz.Domain{{ID: domainId}}})
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

// ListUserMenu 用户菜单列表
func (s *AdminService) ListUserMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	// id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	name := "菜单"
	println(name)
	return &protobuf.PagingReply{}, nil
}

// 将用户菜单转换树形结构
func (s *AdminService) UserMenuToReplyMenu(menu *biz.Menu) *v1.MenuRouter {

	meta := &v1.MenuRouter_Meta{
		// 路由title  一般必填
		Title: menu.Title,
		// 图标，也是菜单图标
		Icon: &menu.Icon,
		// 菜单排序，只对第一级有效
		OrderNo: &menu.Sort,
	}

	// 忽略缓存
	if cache := menu.IsCache == int32(protobuf.MenuCache_MENU_CACHE_NO); cache {
		meta.IgnoreKeepAlive = &cache
	}
	if hidden := (menu.IsHidden == int32(protobuf.MenuHidden_MENU_HIDDEN_YES)); hidden {
		meta.HideMenu = &hidden
	}
	component := menu.Component
	if menu.LinkType == int32(protobuf.MenuLinkType_MENU_Link_TYPE_IFRAME) {
		meta.FrameSrc = &menu.LinkUrl
	}

	return &v1.MenuRouter{
		Name:      menu.Name,
		Path:      menu.Path,
		Component: component,
		Children:  make([]*v1.MenuRouter, 0),
		Meta:      meta,
		Redirect:  menu.Redirect,
	}
}

// 将用户菜单转换树形结构
func (s *AdminService) UserMenuTransformTree(menus []*biz.Menu, parentID uint) []*v1.MenuRouter {
	list := make([]*v1.MenuRouter, 0)
	for _, menu := range menus {
		if menu.Type == int32(protobuf.MenuType_MENU_TYPE_ABILITY) {
			continue
		}
		if menu.ParentID == parentID {
			m := s.UserMenuToReplyMenu(menu)
			m.Children = append(m.Children, s.UserMenuTransformTree(menus, menu.ID)...)
			list = append(list, m)
		}
	}
	return list
}

// 获取角色菜单树形列表
func (s *AdminService) ListUserRoleMenuTree(ctx context.Context, in *v1.ListUserRoleMenuTreeReq) (*v1.ListUserRoleMenuTreeReply, error) {
	var roles []*biz.Role
	if roleId := in.GetRoleId(); roleId > 0 {
		roles = append(roles, &biz.Role{ID: uint(roleId)})
	}
	menuModels, err := s.userCase.ListRoleMenu(ctx, &biz.User{ID: convert.StringToUint(authz.ParseFromContext(ctx).GetUser()), Roles: roles})
	if err != nil {
		s.log.Debugf("用户菜单查询失败 %v", err)
	}
	return &v1.ListUserRoleMenuTreeReply{
		Items: s.UserMenuTransformTree(menuModels, 0),
	}, nil
}

// 获取角色权限列表
func (s *AdminService) ListUserRolePermission(ctx context.Context, in *v1.ListUserRolePermissionReq) (*v1.ListUserRolePermissionReply, error) {
	var roles []*biz.Role
	if roleId := in.GetRoleId(); roleId > 0 {
		roles = append(roles, &biz.Role{ID: uint(roleId)})
	}
	menuModels, _ := s.userCase.ListRoleMenu(ctx, &biz.User{ID: convert.StringToUint(authz.ParseFromContext(ctx).GetUser()), Roles: roles})
	perms := make([]string, 0)
	for _, v := range menuModels {
		perms = append(perms, v.Permission)
	}
	perms = convert.ArrayStrUnique(perms)
	return &v1.ListUserRolePermissionReply{
		Items: perms,
	}, nil
}
