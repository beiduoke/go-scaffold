package admin

import (
	"context"
	"fmt"
	"strconv"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ListApi 列表接口
func (s *AdminService) ListApi(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return &protobuf.PagingReply{
		Total: 0,
		Items: []*anypb.Any{},
	}, nil
}

// CreateApi 创建接口
func (s *AdminService) CreateApi(ctx context.Context, in *v1.CreateApiReq) (*v1.CreateApiReply, error) {
	_, err := s.apiCase.Create(ctx, &biz.Api{
		Name:        in.GetName(),
		Path:        in.GetPath(),
		Method:      in.GetMethod(),
		Group:       in.GetGroup(),
		Description: in.GetDescription(),
	})
	if err != nil {
		return nil, errors.BadRequest("ErrorReason_API_NOT_FOUND", "接口创建失败")
	}
	return &v1.CreateApiReply{
		Success: true,
	}, nil
}

// GetApi 获取接口
func (s *AdminService) GetApi(ctx context.Context, in *v1.GetApiReq) (*v1.Api, error) {
	api, err := s.apiCase.GetID(ctx, &biz.Api{ID: convert.StringToUint(strconv.FormatUint(in.GetId(), 10))})
	if err != nil {
		return nil, errors.BadRequest("ErrorReason_API_NOT_FOUND", "接口查询失败")
	}
	fmt.Println(api)
	return &v1.Api{
		CreatedAt:   timestamppb.New(api.CreatedAt),
		UpdatedAt:   timestamppb.New(api.UpdatedAt),
		Id:          uint64(api.ID),
		Name:        api.Name,
		Path:        api.Path,
		Method:      api.Method,
		Group:       api.Group,
		Description: api.Description,
	}, nil
}

// UpdateApi 修改接口
func (s *AdminService) UpdateApi(ctx context.Context, in *v1.UpdateApiReq) (*v1.UpdateApiReply, error) {
	return &v1.UpdateApiReply{
		Success: true,
	}, nil
}

// DeleteApi 删除接口
func (s *AdminService) DeleteApi(ctx context.Context, in *v1.DeleteApiReq) (*v1.DeleteApiReply, error) {
	return &v1.DeleteApiReply{
		Success: true,
	}, nil
}
