package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/interface/v1"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	coreV1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/pkg/middleware/authn"
	"github.com/beiduoke/go-scaffold/pkg/util/trans"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UserServiceHTTPServer

	uc  coreV1.UserServiceClient
	log *log.Helper
}

func NewUserService(logger log.Logger, uc coreV1.UserServiceClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*coreV1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *coreV1.GetUserRequest) (*coreV1.User, error) {
	return s.uc.GetUser(ctx, req)
}

func (s *UserService) GetUserByName(ctx context.Context, req *coreV1.GetUserByNameRequest) (*coreV1.User, error) {
	return s.uc.GetUserByName(ctx, req)
}

func (s *UserService) CreateUser(ctx context.Context, req *coreV1.CreateUserRequest) (*coreV1.CreateUserResponse, error) {
	// return s.uc.CreateUser(ctx, req)
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId
	req.User.CreatorId = trans.Uint32(authInfo.UserId)
	if req.User.Authority == nil {
		req.User.Authority = (*int32)(coreV1.Authority_AUTHORITY_CUSTOMER_USER.Enum())
	}

	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *coreV1.UpdateUserRequest) (*coreV1.UpdateUserResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, v1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *coreV1.DeleteUserRequest) (*coreV1.DeleteUserResponse, error) {
	authInfo, err := authn.AuthFromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, v1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteUser(ctx, req)
}

func (s *UserService) UserExists(ctx context.Context, req *coreV1.UserExistsRequest) (*coreV1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
