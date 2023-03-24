package api

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

func TransformPost(data *biz.Post) *v1.Post {
	return &v1.Post{
		Id:        uint64(data.ID),
		Name:      data.Name,
		Code:      data.Code,
		Sort:      int32(data.Sort),
		State:     protobuf.PostState(data.State),
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Remarks:   data.Remarks,
	}
}

// ListPost 列表-岗位
func (s *ApiService) ListPost(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.postCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize()), pagination.WithQuery(pagination.QueryUnmarshal(in.GetQuery())), pagination.WithOrderBy(in.GetOrderBy())))
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformPost(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: total,
		Items: items,
	}, nil
}

// CreatePost 创建岗位
func (s *ApiService) CreatePost(ctx context.Context, in *v1.CreatePostReq) (*v1.CreatePostReply, error) {
	user, err := s.postCase.Create(ctx, &biz.Post{
		Name:    in.GetName(),
		State:   int32(in.GetState()),
		Sort:    in.GetSort(),
		Remarks: in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorPostCreateFail("岗位创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreatePostReply{
		Success: true,
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdatePost 修改岗位
func (s *ApiService) UpdatePost(ctx context.Context, in *v1.UpdatePostReq) (*v1.UpdatePostReply, error) {
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
	return &v1.UpdatePostReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// UpdatePostState 修改岗位-状态
func (s *ApiService) UpdatePostState(ctx context.Context, in *v1.UpdatePostStateReq) (*v1.UpdatePostStateReply, error) {
	v := in.GetData()
	err := s.postCase.UpdateState(ctx, &biz.Post{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorPostUpdateFail("岗位状态修改失败: %v", err.Error())
	}
	return &v1.UpdatePostStateReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetPost 获取岗位
func (s *ApiService) GetPost(ctx context.Context, in *v1.GetPostReq) (*v1.Post, error) {
	post, err := s.postCase.GetID(ctx, &biz.Post{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorPostNotFound("岗位未找到")
	}
	return TransformPost(post), nil
}

// DeletePost 删除岗位
func (s *ApiService) DeletePost(ctx context.Context, in *v1.DeletePostReq) (*v1.DeletePostReply, error) {
	if err := s.postCase.Delete(ctx, &biz.Post{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorPostDeleteFail("岗位删除失败：%v", err)
	}
	return &v1.DeletePostReply{
		Message: "删除成功",
	}, nil
}
