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

func (r *PostRepo) toProto(in *ent.Post) *v1.Post {
	if in == nil {
		return nil
	}
	return &v1.Post{
		CreatedAt: convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt: convert.TimeValueToString(in.UpdatedAt, time.DateTime),
		Id:        in.ID,
		Name:      in.Name,
		State:     in.State,
		Remark:    in.Remark,
		Sort:      in.Sort,
	}
}

type PostRepo struct {
	data *Data
	log  *log.Helper
}

// NewPostRepo .
func NewPostRepo(data *Data, logger log.Logger) *PostRepo {
	return &PostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PostRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Post.Query()
	if len(whereCond) > 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *PostRepo) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	builder := r.data.db.Post.Create().SetCreatedAt(time.Now())
	builder = builder.SetName(req.Post.GetName()).
		SetNillableState(req.Post.State).
		SetNillableRemark(req.Post.Remark).
		SetNillableSort(req.Post.Sort)

	_, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreatePostResponse{}, err
}
func (r *PostRepo) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {

	builder := r.data.db.Post.UpdateOneID(req.Id)
	builder = builder.SetName(req.Post.GetName()).
		SetNillableState(req.Post.State).
		SetNillableRemark(req.Post.Remark).
		SetNillableSort(req.Post.Sort)

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UpdatePostResponse{}, err
}
func (r *PostRepo) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	err := r.data.db.Post.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeletePostResponse{}, nil
}
func (r *PostRepo) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.Post, error) {
	ret, err := r.data.db.Post.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return r.toProto(ret), err
}

func (r *PostRepo) ListPost(ctx context.Context, req *pagination.PagingRequest) (*v1.ListPostResponse, error) {
	builder := r.data.db.Debug().Post.Query()
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

	items := make([]*v1.Post, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListPostResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
