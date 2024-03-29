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

func (r *RoleRepo) toProto(in *ent.Role) *v1.Role {
	if in == nil {
		return nil
	}
	return &v1.Role{
		CreatedAt:         convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt:         convert.TimeValueToString(in.UpdatedAt, time.DateTime),
		Id:                in.ID,
		Name:              in.Name,
		DefaultRouter:     in.DefaultRouter,
		Sort:              in.Sort,
		DataScope:         in.DataScope,
		MenuCheckStrictly: in.MenuCheckStrictly,
		DeptCheckStrictly: in.DeptCheckStrictly,
		State:             in.State,
		Remark:            in.Remark,
	}
}

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RoleRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Role.Query()
	if len(whereCond) > 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *RoleRepo) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error) {
	builder := r.data.db.Role.Create().SetCreatedAt(time.Now())
	builder = builder.SetName(req.Role.GetName())

	_, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreateRoleResponse{}, err
}
func (r *RoleRepo) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error) {

	builder := r.data.db.Role.UpdateOneID(req.Id)
	builder = builder.SetNillableSort(req.Role.Sort).
		SetNillableRemark(req.Role.Remark)

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRoleResponse{}, err
}
func (r *RoleRepo) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error) {
	err := r.data.db.Role.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRoleResponse{}, nil
}
func (r *RoleRepo) GetRole(ctx context.Context, req *v1.GetRoleRequest) (*v1.Role, error) {
	ret, err := r.data.db.Role.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return r.toProto(ret), err
}

func (r *RoleRepo) ListRole(ctx context.Context, req *pagination.PagingRequest) (*v1.ListRoleResponse, error) {
	builder := r.data.db.Debug().Role.Query()
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

	items := make([]*v1.Role, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListRoleResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
