package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

func TransformApi(data *biz.Api) *v1.Api {
	return &v1.Api{
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
		Id:          uint64(data.ID),
		Name:        data.Name,
		Path:        data.Path,
		Method:      data.Method,
		Group:       data.Group,
		Description: data.Description,
		Operation:   data.Operation,
	}
}

// ListApi 列表接口
func (s *AdminService) ListApi(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.apiCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformApi(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateApi 创建接口
func (s *AdminService) CreateApi(ctx context.Context, in *v1.CreateApiReq) (*v1.CreateApiReply, error) {
	user, err := s.apiCase.Create(ctx, &biz.Api{
		Name:        in.GetName(),
		Path:        in.GetPath(),
		Method:      in.GetMethod(),
		Group:       in.GetGroup(),
		Description: in.GetDescription(),
		Operation:   in.GetOperation(),
	})
	if err != nil {
		return nil, v1.ErrorApiCreateFail("接口创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateApiReply{
		Success: true,
		Message: "创建成功",
		Data:    data,
	}, nil
}

// UpdateApi 创建接口
func (s *AdminService) UpdateApi(ctx context.Context, in *v1.UpdateApiReq) (*v1.UpdateApiReply, error) {
	v := in.GetData()
	err := s.apiCase.Update(ctx, &biz.Api{
		ID:          uint(in.GetId()),
		Name:        v.GetName(),
		Path:        v.GetPath(),
		Method:      v.GetMethod(),
		Group:       v.GetGroup(),
		Description: v.GetDescription(),
		Operation:   v.GetOperation(),
	})
	if err != nil {
		return nil, v1.ErrorApiUpdateFail("接口创建失败: %v", err.Error())
	}
	return &v1.UpdateApiReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetApi 获取接口
func (s *AdminService) GetApi(ctx context.Context, in *v1.GetApiReq) (*v1.Api, error) {
	api, err := s.apiCase.GetID(ctx, &biz.Api{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorApiNotFound("接口未找到")
	}
	return TransformApi(api), nil
}

// DeleteApi 删除接口
func (s *AdminService) DeleteApi(ctx context.Context, in *v1.DeleteApiReq) (*v1.DeleteApiReply, error) {
	if err := s.apiCase.Delete(ctx, &biz.Api{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorApiDeleteFail("接口删除失败：%v", err)
	}
	return &v1.DeleteApiReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
