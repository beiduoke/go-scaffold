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
		Mobile:    data.Mobile,
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
		Total: int32(total),
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
		Mobile:   in.GetMobile(),
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
		Data:    data,
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

// HandleUserAuthority 绑定用户权限
func (s *AdminService) HandleUserDomainAuthority(ctx context.Context, in *v1.HandleUserDomainAuthorityReq) (*v1.HandleUserDomainAuthorityReply, error) {
	v := in.GetData()
	_, err := s.userCase.GetID(ctx, &biz.User{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	authorityIds := make([]uint, 0, len(v.GetAuthorityIds()))
	for _, authorityId := range v.GetAuthorityIds() {
		authorityIds = append(authorityIds, uint(authorityId))
	}
	authorities, _ := s.authorityCase.ListByIDs(ctx, authorityIds...)
	err = s.userCase.HandleDomainAuthority(ctx, &biz.User{
		ID:          uint(in.GetId()),
		Authorities: authorities,
	}, uint(v.GetDomainId()))
	if err != nil {
		return nil, v1.ErrorUserHandleDomainAuthorityFail("绑定用户权限失败: %v", err.Error())
	}
	return &v1.HandleUserDomainAuthorityReply{
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
	authorityResult, err := s.ListUserAuthority(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserProfileReply{
		User:        TransformUser(user),
		Authorities: authorityResult.GetItems(),
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

// ListUserAuthority 用户权限角色
func (s *AdminService) ListUserAuthority(ctx context.Context, in *emptypb.Empty) (*v1.ListUserAuthorityReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	domainId := convert.StringToUint(authz.ParseFromContext(ctx).GetDomain())
	authorityModels, err := s.userCase.ListAuthorityAll(ctx, &biz.User{ID: id, Domains: []*biz.Domain{{ID: domainId}}})
	if err != nil {
		return nil, v1.ErrorUserAuthorityFindFail("用户权限角色失败 %v", err)
	}
	authorities := make([]*v1.Authority, 0, len(authorityModels))
	for _, v := range authorityModels {
		authorities = append(authorities, TransformAuthority(v))
	}
	return &v1.ListUserAuthorityReply{
		Items: authorities,
	}, nil
}

// ProfileUser 概括
func (s *AdminService) ListUserMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	// id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	name := "菜单"
	println(name)
	return &protobuf.PagingReply{}, nil
}

// ProfileUser 概括
func (s *AdminService) ListUserMenuTree(ctx context.Context, in *emptypb.Empty) (*v1.UserMenuTreeReply, error) {
	// id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	name := "菜单"
	println(name)
	return &v1.UserMenuTreeReply{}, nil
}

// 将用户菜单转换树形结构
func (s *AdminService) UserMenuToReplyMenu(menu *biz.Menu) *v1.ListUserAuthorityMenuTreeReply_Menu {
	return &v1.ListUserAuthorityMenuTreeReply_Menu{
		Name:      menu.Name,
		Path:      menu.Path,
		Component: menu.Component,
		Children:  make([]*v1.ListUserAuthorityMenuTreeReply_Menu, 0),
		Meta: &v1.ListUserAuthorityMenuTreeReply_MenuMeta{
			// 路由title  一般必填
			Title: menu.Title,
			// 动态路由可打开Tab页数
			DynamicLevel: 100,
			// 动态路由的实际Path, 即去除路由的动态部分;
			// RealPath: menu.RealPath,
			// 是否忽略KeepAlive缓存
			IgnoreKeepAlive: menu.KeepAlive == int32(protobuf.MenuKeepAlive_MENU_KEEP_ALIVE_NO),
			// 是否固定标签
			Affix: false,
			// 图标，也是菜单图标
			Icon: menu.Icon,
			// 内嵌iframe的地址
			// FrameSrc: menu.FrameSrc,
			// 指定该路由切换的动画名
			// TransitionName: menu.TransitionName,
			// 隐藏该路由在面包屑上面的显示
			HideBreadcrumb: true,
			// 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
			CarryParam: true,
			// 隐藏所有子菜单
			HideChildrenInMenu: false,
			// 当前激活的菜单。用于配置详情页时左侧激活的菜单路径
			// CurrentActiveMenu: menu.CurrentActiveMenu,
			// 当前路由不再标签页显示
			HideTab: false,
			// 当前路由不再菜单显示
			HideMenu: (menu.Hidden == int32(protobuf.MenuHidden_MENU_HIDDEN_YES)),
			// 菜单排序，只对第一级有效
			OrderNo: menu.Sort,
			// 忽略路由。用于在ROUTE_MAPPING以及BACK权限模式下，生成对应的菜单而忽略路由。2.5.3以上版本有效
			IgnoreRoute: false,
			// 是否在子级菜单的完整path中忽略本级path。2.5.3以上版本有效
			HidePathForChildren: false,
		},
	}
}

// 将用户菜单转换树形结构
func (s *AdminService) UserMenuTransformTree(menus []*biz.Menu, parentID uint) []*v1.ListUserAuthorityMenuTreeReply_Menu {
	list := make([]*v1.ListUserAuthorityMenuTreeReply_Menu, 0)
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

// 获取权限角色菜单树形列表
func (s *AdminService) ListUserAuthorityMenuTree(ctx context.Context, in *v1.ListUserAuthorityMenuTreeReq) (*v1.ListUserAuthorityMenuTreeReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	authorityIds, err := s.userCase.ListAuthorityID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserAuthorityFindFail("用户权限角色查询失败 %v", err)
	}
	authorityModels := make([]*biz.Authority, 0)
	if in.GetAuthorityId() > 0 {
		authId := uint(in.GetAuthorityId())
		isExit := false
		for _, v := range authorityIds {
			if v == authId {
				isExit = true
				authorityModels = []*biz.Authority{{ID: v}}
				break
			}
		}
		if !isExit {
			return nil, v1.ErrorUserAuthorityFindFail("用户权限角色不存在")
		}
	}
	for _, v := range authorityIds {
		authorityModels = append(authorityModels, &biz.Authority{ID: v})
	}
	menuModels, _ := s.userCase.ListAuthorityMenuAll(ctx, &biz.User{ID: id, Authorities: authorityModels})
	return &v1.ListUserAuthorityMenuTreeReply{
		Items: s.UserMenuTransformTree(menuModels, 0),
	}, nil
}

// 获取权限角色权限列表
func (s *AdminService) ListUserAuthorityPermission(ctx context.Context, in *v1.ListUserAuthorityPermissionReq) (*v1.ListUserAuthorityPermissionReply, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	authorityIds, err := s.userCase.ListAuthorityID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserAuthorityFindFail("用户权限角色查询失败 %v", err)
	}
	authorityModels := make([]*biz.Authority, 0)
	if in.GetAuthorityId() > 0 {
		authId := uint(in.GetAuthorityId())
		isExit := false
		for _, v := range authorityIds {
			if v == authId {
				isExit = true
				authorityModels = []*biz.Authority{{ID: v}}
				break
			}
		}
		if !isExit {
			return nil, v1.ErrorUserAuthorityFindFail("用户权限角色不存在")
		}
	}
	for _, v := range authorityIds {
		authorityModels = append(authorityModels, &biz.Authority{ID: v})
	}
	menuModels, _ := s.userCase.ListAuthorityMenuAll(ctx, &biz.User{ID: id, Authorities: authorityModels})
	perms := make([]string, 0)
	for _, v := range menuModels {
		perms = append(perms, v.Permission)
	}
	perms = convert.ArrayStrUnique(perms)
	return &v1.ListUserAuthorityPermissionReply{
		Items: perms,
	}, nil
}
