package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
)

func (r *DeptRepo) toProto(in *ent.Dept) *v1.Dept {
	if in == nil {
		return nil
	}
	return &v1.Dept{
		CreatedAt: convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt: convert.TimeValueToString(in.UpdatedAt, time.DateTime),
		Id:        in.ID,
		Name:      in.Name,
		State:     in.State,
		Remark:    in.Remark,
		Sort:      in.Sort,
	}
}

type DeptRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeptRepo .
func NewDeptRepo(data *Data, logger log.Logger) *DeptRepo {
	return &DeptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DeptRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Dept.Query()
	if len(whereCond) > 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *DeptRepo) CreateDept(ctx context.Context, req *v1.CreateDeptRequest) (*v1.CreateDeptResponse, error) {
	builder := r.data.db.Dept.Create().SetCreatedAt(time.Now())
	builder = builder.SetName(req.Dept.GetName()).
		SetNillableState(req.Dept.State).
		SetNillableRemark(req.Dept.Remark).
		SetNillableSort(req.Dept.Sort)

	_, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreateDeptResponse{}, err
}
func (r *DeptRepo) UpdateDept(ctx context.Context, req *v1.UpdateDeptRequest) (*v1.UpdateDeptResponse, error) {

	builder := r.data.db.Dept.UpdateOneID(req.Id)
	builder = builder.SetName(req.Dept.GetName()).
		SetNillableState(req.Dept.State).
		SetNillableRemark(req.Dept.Remark).
		SetNillableSort(req.Dept.Sort)

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDeptResponse{}, err
}
func (r *DeptRepo) DeleteDept(ctx context.Context, req *v1.DeleteDeptRequest) (*v1.DeleteDeptResponse, error) {
	err := r.data.db.Dept.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteDeptResponse{}, nil
}
func (r *DeptRepo) GetDept(ctx context.Context, req *v1.GetDeptRequest) (*v1.Dept, error) {
	ret, err := r.data.db.Dept.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return r.toProto(ret), err
}

func (r *DeptRepo) ListDept(ctx context.Context, req *pagination.PagingRequest) (*v1.ListDeptResponse, error) {
	builder := r.data.db.Debug().Dept.Query()
	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.Dept, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListDeptResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
