package api

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/protobuf"
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/internal/pkg/proto"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.ApiServer = (*ApiService)(nil)

func TransformDept(data *biz.Dept) *v1.Dept {
	return &v1.Dept{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		ParentId:  uint64(data.ParentID),
		Sort:      data.Sort,
		Remarks:   data.Remarks,
		State:     protobuf.DeptState(data.State),
		Children:  []*v1.Dept{},
	}
}

// TreeMenu 部门树形
func TreeDept(depts []*biz.Dept, pid uint) []*v1.Dept {
	list := make([]*v1.Dept, 0)
	for _, dept := range depts {
		if dept.ParentID == pid {
			m := TransformDept(dept)
			m.Children = append(m.Children, TreeDept(depts, dept.ID)...)
			list = append(list, m)
		}
	}
	return list
}

// GetTreeDept 列表部门-树形
func (s *ApiService) ListDeptTree(ctx context.Context, in *v1.ListDeptTreeReq) (*v1.ListDeptTreeReply, error) {
	results, _ := s.deptCase.ListAll(ctx)
	treeData := make([]*v1.Dept, 0)
	for _, v := range results {
		treeData = append(treeData, TransformDept(v))
	}
	return &v1.ListDeptTreeReply{
		Items: proto.ToTree(treeData, in.GetId(), func(t *v1.Dept, ts ...*v1.Dept) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}

// ListDept 列表-部门角色
func (s *ApiService) ListDept(ctx context.Context, in *v1.ListDeptReq) (*v1.ListDeptReply, error) {
	results, total := s.deptCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDeptReply{
		Total: total,
		Items: proto.ToAny(results, func(t *biz.Dept) protoreflect.ProtoMessage {
			return TransformDept(t)
		}),
	}, nil
}

// CreateDept 创建部门角色
func (s *ApiService) CreateDept(ctx context.Context, in *v1.CreateDeptReq) (*v1.CreateDeptReply, error) {
	user, err := s.deptCase.Create(ctx, &biz.Dept{
		Name:     in.GetName(),
		ParentID: uint(in.GetParentId()),
		State:    int32(in.GetState()),
		Remarks:  in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDeptCreateFail("部门角色创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateDeptReply{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDept 创建部门角色
func (s *ApiService) UpdateDept(ctx context.Context, in *v1.UpdateDeptReq) (*v1.UpdateDeptReply, error) {
	v := in.GetData()
	err := s.deptCase.Update(ctx, &biz.Dept{
		ID:       uint(in.GetId()),
		Name:     v.GetName(),
		ParentID: uint(v.GetParentId()),
		Sort:     v.GetSort(),
		State:    int32(v.GetState()),
		Remarks:  v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDeptUpdateFail("部门角色创建失败: %v", err.Error())
	}
	return &v1.UpdateDeptReply{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetDept 获取部门角色
func (s *ApiService) GetDept(ctx context.Context, in *v1.GetDeptReq) (*v1.Dept, error) {
	dept, err := s.deptCase.GetID(ctx, &biz.Dept{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDeptNotFound("部门角色未找到")
	}
	return TransformDept(dept), nil
}

// DeleteDept 删除部门角色
func (s *ApiService) DeleteDept(ctx context.Context, in *v1.DeleteDeptReq) (*v1.DeleteDeptReply, error) {
	if err := s.deptCase.Delete(ctx, &biz.Dept{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDeptDeleteFail("部门角色删除失败：%v", err)
	}
	return &v1.DeleteDeptReply{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}
