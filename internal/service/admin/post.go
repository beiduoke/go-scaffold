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

// ListPost 列表-职位
func (s *AdminService) ListPost(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.postCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformPost(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreatePost 创建职位
func (s *AdminService) CreatePost(ctx context.Context, in *v1.CreatePostReq) (*v1.CreatePostReply, error) {
	user, err := s.postCase.Create(ctx, &biz.Post{
		Name:    in.GetName(),
		State:   int32(in.GetState()),
		Sort:    in.GetSort(),
		Remarks: in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorPostCreateFail("职位创建失败: %v", err.Error())
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

// UpdatePost 修改职位
func (s *AdminService) UpdatePost(ctx context.Context, in *v1.UpdatePostReq) (*v1.UpdatePostReply, error) {
	v := in.GetData()
	err := s.postCase.Update(ctx, &biz.Post{
		ID:      uint(in.GetId()),
		Name:    v.GetName(),
		Sort:    v.GetSort(),
		State:   int32(v.GetState()),
		Remarks: v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorPostUpdateFail("职位修改失败: %v", err.Error())
	}
	return &v1.UpdatePostReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// UpdatePostState 修改职位-状态
func (s *AdminService) UpdatePostState(ctx context.Context, in *v1.UpdatePostStateReq) (*v1.UpdatePostStateReply, error) {
	v := in.GetData()
	err := s.postCase.UpdateState(ctx, &biz.Post{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorPostUpdateFail("职位状态修改失败: %v", err.Error())
	}
	return &v1.UpdatePostStateReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetPost 获取职位
func (s *AdminService) GetPost(ctx context.Context, in *v1.GetPostReq) (*v1.Post, error) {
	post, err := s.postCase.GetID(ctx, &biz.Post{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorPostNotFound("职位未找到")
	}
	return TransformPost(post), nil
}

// DeletePost 删除职位
func (s *AdminService) DeletePost(ctx context.Context, in *v1.DeletePostReq) (*v1.DeletePostReply, error) {
	if err := s.postCase.Delete(ctx, &biz.Post{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorPostDeleteFail("职位删除失败：%v", err)
	}
	return &v1.DeletePostReply{
		Message: "删除成功",
	}, nil
}
