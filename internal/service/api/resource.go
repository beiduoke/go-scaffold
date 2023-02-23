package api

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

func TransformResource(data *biz.Resource) *v1.Resource {
	return &v1.Resource{
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

// ListResource 列表资源
func (s *ApiService) ListResource(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.resourceCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformResource(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: total,
		Items: items,
	}, nil
}

// ListResourceGroup 列表资源-分组
func (s *ApiService) ListResourceGroup(ctx context.Context, in *protobuf.PagingReq) (*v1.ListResourceGroupReply, error) {
	results, total := s.resourceCase.ListAllGroup(ctx)
	return &v1.ListResourceGroupReply{
		Total: &total,
		Items: results,
	}, nil
}

// CreateResource 创建资源
func (s *ApiService) CreateResource(ctx context.Context, in *v1.CreateResourceReq) (*v1.CreateResourceReply, error) {
	user, err := s.resourceCase.Create(ctx, &biz.Resource{
		Name:        in.GetName(),
		Path:        in.GetPath(),
		Method:      in.GetMethod(),
		Group:       in.GetGroup(),
		Description: in.GetDescription(),
		Operation:   in.GetOperation(),
	})
	if err != nil {
		return nil, v1.ErrorResourceCreateFail("资源创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateResourceReply{
		Success: true,
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateResource 修改资源
func (s *ApiService) UpdateResource(ctx context.Context, in *v1.UpdateResourceReq) (*v1.UpdateResourceReply, error) {
	v := in.GetData()
	err := s.resourceCase.Update(ctx, &biz.Resource{
		ID:          uint(in.GetId()),
		Name:        v.GetName(),
		Path:        v.GetPath(),
		Method:      v.GetMethod(),
		Group:       v.GetGroup(),
		Description: v.GetDescription(),
		Operation:   v.GetOperation(),
	})
	if err != nil {
		return nil, v1.ErrorResourceUpdateFail("资源修改失败: %v", err.Error())
	}
	return &v1.UpdateResourceReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetResource 获取资源
func (s *ApiService) GetResource(ctx context.Context, in *v1.GetResourceReq) (*v1.Resource, error) {
	api, err := s.resourceCase.GetID(ctx, &biz.Resource{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorResourceNotFound("资源未找到")
	}
	return TransformResource(api), nil
}

// DeleteResource 删除资源
func (s *ApiService) DeleteResource(ctx context.Context, in *v1.DeleteResourceReq) (*v1.DeleteResourceReply, error) {
	if err := s.resourceCase.Delete(ctx, &biz.Resource{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorResourceDeleteFail("资源删除失败：%v", err)
	}
	return &v1.DeleteResourceReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
