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

func TransformDept(data *biz.Dept) *v1.Dept {
	return &v1.Dept{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		ParentId:  uint64(data.ParentID),
		Sort:      data.Sort,
		State:     protobuf.DeptState(data.State),
		Remarks:   data.Remarks,
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
func (s *AdminService) ListDeptTree(ctx context.Context, in *v1.ListDeptTreeReq) (*v1.ListDeptTreeReply, error) {
	results, _ := s.deptCase.ListAll(ctx)
	return &v1.ListDeptTreeReply{
		Items: TreeDept(results, uint(in.GetId())),
	}, nil
}

// // ListAllDeptTree 列表全部部门-树形
// func (s *AdminService) ListAllDeptTree(ctx context.Context, in *emptypb.Empty) (*v1.ListDeptTreeReply, error) {
// 	results, total := s.deptCase.ListAll(ctx)
// 	return &v1.ListDeptTreeReply{
// 		Items: TreeDept(results, 0),
// 		Total: total,
// 	}, nil
// }

// ListDept 列表-部门角色
func (s *AdminService) ListDept(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.deptCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformDept(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: total,
		Items: items,
	}, nil
}

// CreateDept 创建部门角色
func (s *AdminService) CreateDept(ctx context.Context, in *v1.CreateDeptReq) (*v1.CreateDeptReply, error) {
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
		Success: true,
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDept 创建部门角色
func (s *AdminService) UpdateDept(ctx context.Context, in *v1.UpdateDeptReq) (*v1.UpdateDeptReply, error) {
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
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetDept 获取部门角色
func (s *AdminService) GetDept(ctx context.Context, in *v1.GetDeptReq) (*v1.Dept, error) {
	dept, err := s.deptCase.GetID(ctx, &biz.Dept{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDeptNotFound("部门角色未找到")
	}
	return TransformDept(dept), nil
}

// DeleteDept 删除部门角色
func (s *AdminService) DeleteDept(ctx context.Context, in *v1.DeleteDeptReq) (*v1.DeleteDeptReply, error) {
	if err := s.deptCase.Delete(ctx, &biz.Dept{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDeptDeleteFail("部门角色删除失败：%v", err)
	}
	return &v1.DeleteDeptReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
