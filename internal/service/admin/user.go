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
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
		Id:          uint64(data.ID),
		Name:        data.Name,
		Avatar:      data.Avatar,
		NickName:    data.NickName,
		RealName:    data.RealName,
		Gender:      protobuf.UserGender(data.Gender),
		Birthday:    birthday,
		Mobile:      data.Mobile,
		Email:       data.Email,
		State:       protobuf.UserState(data.State),
		Authorities: make([]*v1.Authority, 0, 0),
	}
}

// ProfileUser 概括
func (s *AdminService) GetUserProfile(ctx context.Context, in *emptypb.Empty) (*v1.User, error) {
	id := convert.StringToUint(authz.ParseFromContext(ctx).GetUser())
	user, err := s.userCase.GetID(ctx, &biz.User{ID: id})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户查询失败 %v", err)
	}
	result := TransformUser(user)
	for _, v := range user.Authorities {
		result.Authorities = append(result.Authorities, TransformAuthority(v))
	}
	return result, nil
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
