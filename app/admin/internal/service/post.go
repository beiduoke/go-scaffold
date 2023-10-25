package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServiceServer = (*AdminService)(nil)

func TransformPost(data *biz.Post) *v1.Post {
	return &v1.Post{
		Id:        uint64(data.ID),
		Name:      data.Name,
		Code:      &data.Code,
		Sort:      &data.Sort,
		State:     &data.State,
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Remarks:   &data.Remarks,
	}
}

// ListPost 列表-岗位
func (s *AdminService) ListPost(ctx context.Context, in *v1.ListPostRequest) (*v1.ListPostResponse, error) {
	results, total := s.postCase.ListPage(ctx, pagination.NewPagination(
		pagination.WithPage(in.GetPage()),
		pagination.WithPageSize(in.GetPageSize()),
		pagination.WithQuery(map[string]interface{}{
			// "name": in.GetName(),
		}),
	))
	items := make([]*v1.Post, 0, len(results))
	for _, v := range results {
		items = append(items, TransformPost(v))
	}
	return &v1.ListPostResponse{
		Total: total,
		Items: items,
	}, nil
}

// CreatePost 创建岗位
func (s *AdminService) CreatePost(ctx context.Context, in *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	user, err := s.postCase.Create(ctx, &biz.Post{
		Name:    in.GetName(),
		State:   int32(in.GetState()),
		Sort:    in.GetSort(),
		Remarks: in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorPostCreateFail("岗位创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&v1.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreatePostResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdatePost 修改岗位
func (s *AdminService) UpdatePost(ctx context.Context, in *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	v := in.GetData()
	err := s.postCase.Update(ctx, &biz.Post{
		ID:      uint(in.GetId()),
		Name:    v.GetName(),
		Sort:    v.GetSort(),
		State:   int32(v.GetState()),
		Remarks: v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorPostUpdateFail("岗位修改失败: %v", err.Error())
	}
	return &v1.UpdatePostResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdatePostState 修改岗位-状态
func (s *AdminService) UpdatePostState(ctx context.Context, in *v1.UpdatePostStateRequest) (*v1.UpdatePostStateResponse, error) {
	v := in.GetData()
	err := s.postCase.UpdateState(ctx, &biz.Post{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorPostUpdateFail("岗位状态修改失败: %v", err.Error())
	}
	return &v1.UpdatePostStateResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetPost 获取岗位
func (s *AdminService) GetPost(ctx context.Context, in *v1.GetPostRequest) (*v1.Post, error) {
	post, err := s.postCase.GetID(ctx, &biz.Post{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorPostNotFound("岗位未找到")
	}
	return TransformPost(post), nil
}

// DeletePost 删除岗位
func (s *AdminService) DeletePost(ctx context.Context, in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	if err := s.postCase.Delete(ctx, &biz.Post{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorPostDeleteFail("岗位删除失败：%v", err)
	}
	return &v1.DeletePostResponse{
		Message: "删除成功",
	}, nil
}
