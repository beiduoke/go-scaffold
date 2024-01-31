package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"
	"github.com/beiduoke/go-scaffold/pkg/util/crypto"
	paging "github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/beiduoke/go-scaffold/pkg/util/trans"
	"github.com/go-kratos/kratos/v2/log"
)

func (r *UserRepo) toProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	var (
		createdAt, updatedAt, birthday string
	)
	if in.CreatedAt != nil {
		createdAt = in.CreatedAt.Format(time.DateTime)
	}
	if in.UpdatedAt != nil {
		updatedAt = in.UpdatedAt.Format(time.DateTime)
	}
	if in.Birthday != nil {
		birthday = in.Birthday.Format(time.DateOnly)
	}

	return &v1.User{
		Id:        in.ID,
		UserName:  in.UserName,
		NickName:  in.NickName,
		RealName:  in.RealName,
		Birthday:  &birthday,
		Gender:    in.Gender,
		Phone:     in.Phone,
		Email:     in.Email,
		Avatar:    in.Avatar,
		State:     in.State,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
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
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	pass, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}
	builder := r.data.db.User.Create().
		SetUserName(req.User.UserName).
		SetNillableNickName(req.User.NickName).
		SetPhone(trans.StringValue(req.User.Phone)).
		SetNillableEmail(req.User.Email).
		SetPassword(pass).
		SetCreatedAt(time.Now())

	_, err = builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreateUserResponse{}, err
}
func (r *UserRepo) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return &v1.UpdateUserResponse{}, nil
}
func (r *UserRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return &v1.DeleteUserResponse{}, nil
}
func (r *UserRepo) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	return &v1.User{}, nil
}

func (r *UserRepo) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Debug().User.Query()

	builder = builder.Offset(paging.GetPageOffset(*req.Page, *req.PageSize)).Limit(int(req.GetPageSize()))
	// user.By
	// builder = builder.Order(paging.GetSorting(req.Sorting))

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count := 0
	// count, err := r.Count(ctx, whereSelectors)
	// if err != nil {
	// 	return nil, err
	// }

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
func (r *UserRepo) GetUserByUserName(ctx context.Context, req *v1.GetUserByUserNameRequest) (*v1.User, error) {
	ret, err := r.data.db.User.Query().
		Where(user.UserNameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())
		return nil, err
	}

	return r.toProto(ret), err
}
func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	return &v1.VerifyPasswordResponse{}, nil
}
func (r *UserRepo) UserExists(ctx context.Context, req *v1.UserExistsRequest) (*v1.UserExistsResponse, error) {
	return &v1.UserExistsResponse{}, nil
}
