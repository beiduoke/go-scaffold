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
	"github.com/beiduoke/go-scaffold/pkg/util/crypto"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
)

func (r *UserRepo) toProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:        in.ID,
		Name:      in.Name,
		NickName:  in.NickName,
		RealName:  in.RealName,
		Gender:    in.Gender,
		Phone:     in.Phone,
		Email:     in.Email,
		Avatar:    in.Avatar,
		State:     in.State,
		CreatedAt: convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt: convert.TimeValueToString(in.UpdatedAt, time.DateTime),
		Birthday:  convert.TimeValueToString(in.Birthday, time.DateOnly),
	}
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.User.Query()
	if len(whereCond) > 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	builder := r.data.db.User.Create().SetCreatedAt(time.Now())
	if req.User.Password != nil {
		pass, err := crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return nil, err
		}
		builder = builder.SetPassword(pass)
	}
	builder = builder.SetName(req.User.GetName()).
		SetPhone(req.User.GetPhone()).
		SetNillableNickName(req.User.NickName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableAvatar(req.User.Avatar).
		SetNillableDescription(req.User.Description).
		SetNillableAuthority(req.User.Authority).
		SetNillableBirthday(convert.StringValueToTime(req.User.Birthday, time.DateOnly))

	_, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreateUserResponse{}, err
}
func (r *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {

	builder := r.data.db.User.UpdateOneID(req.Id)
	if req.User.Password != nil {
		pass, err := crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return nil, err
		}
		builder = builder.SetPassword(pass)
	}
	builder = builder.SetPhone(req.User.GetPhone()).
		SetNillableNickName(req.User.NickName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableAvatar(req.User.Avatar).
		SetNillableDescription(req.User.Description).
		SetNillableAuthority(req.User.Authority).
		SetNillableBirthday(convert.StringValueToTime(req.User.Birthday, time.DateOnly))

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserResponse{}, err
}
func (r *UserRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	err := r.data.db.User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserResponse{}, nil
}
func (r *UserRepo) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	ret, err := r.data.db.User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return r.toProto(ret), err
}

func (r *UserRepo) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Debug().User.Query()
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

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
func (r *UserRepo) GetUserByName(ctx context.Context, req *v1.GetUserByNameRequest) (*v1.User, error) {
	ret, err := r.data.db.User.Query().
		Where(user.NameEQ(req.GetName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, err
	}

	return r.toProto(ret), err
}
func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	res, err := r.data.db.User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.NameEQ(req.GetName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *res.Password)
	if !bMatched {
		return nil, v1.ErrorUserNotFound("密码错误")
	}

	return &v1.VerifyPasswordResponse{}, nil
}
func (r *UserRepo) UserExists(ctx context.Context, req *v1.UserExistsRequest) (*v1.UserExistsResponse, error) {
	_, err := r.data.db.User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.NameEQ(req.GetName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, v1.ErrorUserNotFound("用户未找到")
	}
	return &v1.UserExistsResponse{Exist: true}, nil
}
