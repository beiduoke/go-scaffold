package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServiceServer = (*AdminService)(nil)

func TransformDept(data *biz.Dept) *v1.Dept {
	if data == nil {
		return nil
	}
	pid := uint64(data.ParentID)
	return &v1.Dept{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		ParentId:  &pid,
		Sort:      &data.Sort,
		Remarks:   &data.Remarks,
		State:     &data.State,
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
func (s *AdminService) ListDeptTree(ctx context.Context, in *v1.ListDeptTreeRequest) (*v1.ListDeptTreeResponse, error) {
	results, _ := s.deptCase.ListAll(ctx)
	treeData := make([]*v1.Dept, 0)
	for _, v := range results {
		treeData = append(treeData, TransformDept(v))
	}
	return &v1.ListDeptTreeResponse{
		Items: convert.ToTree(treeData, in.GetId(), func(t *v1.Dept, ts ...*v1.Dept) error {
			t.Children = append(t.Children, ts...)
			return nil
		}),
	}, nil
}

// ListDept 列表-部门角色
func (s *AdminService) ListDept(ctx context.Context, in *v1.ListDeptRequest) (*v1.ListDeptResponse, error) {
	results, total := s.deptCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDeptResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(t *biz.Dept) *v1.Dept {
			return TransformDept(t)
		}),
	}, nil
}

// CreateDept 创建部门角色
func (s *AdminService) CreateDept(ctx context.Context, in *v1.CreateDeptRequest) (*v1.CreateDeptResponse, error) {
	user, err := s.deptCase.Create(ctx, &biz.Dept{
		Name:     in.GetName(),
		ParentID: uint(in.GetParentId()),
		State:    int32(in.GetState()),
		Remarks:  in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDeptCreateFail("部门角色创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&v1.Result{
		Id: uint64(user.ID),
	})
	return &v1.CreateDeptResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDept 创建部门角色
func (s *AdminService) UpdateDept(ctx context.Context, in *v1.UpdateDeptRequest) (*v1.UpdateDeptResponse, error) {
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
	return &v1.UpdateDeptResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDeptState 修改部门-状态
func (s *AdminService) UpdateDeptState(ctx context.Context, in *v1.UpdateDeptStateRequest) (*v1.UpdateDeptStateResponse, error) {
	v := in.GetData()
	err := s.deptCase.UpdateState(ctx, &biz.Dept{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDeptUpdateFail("岗位状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDeptStateResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// GetDept 获取部门角色
func (s *AdminService) GetDept(ctx context.Context, in *v1.GetDeptRequest) (*v1.Dept, error) {
	dept, err := s.deptCase.GetID(ctx, &biz.Dept{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDeptNotFound("部门角色未找到")
	}
	return TransformDept(dept), nil
}

// DeleteDept 删除部门角色
func (s *AdminService) DeleteDept(ctx context.Context, in *v1.DeleteDeptRequest) (*v1.DeleteDeptResponse, error) {
	if err := s.deptCase.Delete(ctx, &biz.Dept{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDeptDeleteFail("部门角色删除失败：%v", err)
	}
	return &v1.DeleteDeptResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}
