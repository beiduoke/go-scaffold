// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/member"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MemberCreate) SetCreatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MemberCreate) SetUpdatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableUpdatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MemberCreate) SetDeletedAt(t time.Time) *MemberCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableDeletedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetRemark sets the "remark" field.
func (mc *MemberCreate) SetRemark(s string) *MemberCreate {
	mc.mutation.SetRemark(s)
	return mc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (mc *MemberCreate) SetNillableRemark(s *string) *MemberCreate {
	if s != nil {
		mc.SetRemark(*s)
	}
	return mc
}

// SetSort sets the "sort" field.
func (mc *MemberCreate) SetSort(i int32) *MemberCreate {
	mc.mutation.SetSort(i)
	return mc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mc *MemberCreate) SetNillableSort(i *int32) *MemberCreate {
	if i != nil {
		mc.SetSort(*i)
	}
	return mc
}

// SetState sets the "state" field.
func (mc *MemberCreate) SetState(i int32) *MemberCreate {
	mc.mutation.SetState(i)
	return mc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mc *MemberCreate) SetNillableState(i *int32) *MemberCreate {
	if i != nil {
		mc.SetState(*i)
	}
	return mc
}

// SetUsername sets the "username" field.
func (mc *MemberCreate) SetUsername(s string) *MemberCreate {
	mc.mutation.SetUsername(s)
	return mc
}

// SetPassword sets the "password" field.
func (mc *MemberCreate) SetPassword(s string) *MemberCreate {
	mc.mutation.SetPassword(s)
	return mc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mc *MemberCreate) SetNillablePassword(s *string) *MemberCreate {
	if s != nil {
		mc.SetPassword(*s)
	}
	return mc
}

// SetNickname sets the "nickname" field.
func (mc *MemberCreate) SetNickname(s string) *MemberCreate {
	mc.mutation.SetNickname(s)
	return mc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mc *MemberCreate) SetNillableNickname(s *string) *MemberCreate {
	if s != nil {
		mc.SetNickname(*s)
	}
	return mc
}

// SetPhone sets the "phone" field.
func (mc *MemberCreate) SetPhone(s string) *MemberCreate {
	mc.mutation.SetPhone(s)
	return mc
}

// SetEmail sets the "email" field.
func (mc *MemberCreate) SetEmail(s string) *MemberCreate {
	mc.mutation.SetEmail(s)
	return mc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mc *MemberCreate) SetNillableEmail(s *string) *MemberCreate {
	if s != nil {
		mc.SetEmail(*s)
	}
	return mc
}

// SetAvatar sets the "avatar" field.
func (mc *MemberCreate) SetAvatar(s string) *MemberCreate {
	mc.mutation.SetAvatar(s)
	return mc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mc *MemberCreate) SetNillableAvatar(s *string) *MemberCreate {
	if s != nil {
		mc.SetAvatar(*s)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *MemberCreate) SetDescription(s string) *MemberCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MemberCreate) SetNillableDescription(s *string) *MemberCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(u uint32) *MemberCreate {
	mc.mutation.SetID(u)
	return mc
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MemberCreate) defaults() {
	if _, ok := mc.mutation.Remark(); !ok {
		v := member.DefaultRemark
		mc.mutation.SetRemark(v)
	}
	if _, ok := mc.mutation.Sort(); !ok {
		v := member.DefaultSort
		mc.mutation.SetSort(v)
	}
	if _, ok := mc.mutation.State(); !ok {
		v := member.DefaultState
		mc.mutation.SetState(v)
	}
	if _, ok := mc.mutation.Password(); !ok {
		v := member.DefaultPassword
		mc.mutation.SetPassword(v)
	}
	if _, ok := mc.mutation.Nickname(); !ok {
		v := member.DefaultNickname
		mc.mutation.SetNickname(v)
	}
	if _, ok := mc.mutation.Email(); !ok {
		v := member.DefaultEmail
		mc.mutation.SetEmail(v)
	}
	if _, ok := mc.mutation.Avatar(); !ok {
		v := member.DefaultAvatar
		mc.mutation.SetAvatar(v)
	}
	if _, ok := mc.mutation.Description(); !ok {
		v := member.DefaultDescription
		mc.mutation.SetDescription(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Member.sort"`)}
	}
	if v, ok := mc.mutation.Sort(); ok {
		if err := member.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Member.sort": %w`, err)}
		}
	}
	if _, ok := mc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Member.state"`)}
	}
	if v, ok := mc.mutation.State(); ok {
		if err := member.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Member.state": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Member.username"`)}
	}
	if v, ok := mc.mutation.Username(); ok {
		if err := member.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Member.username": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Member.password"`)}
	}
	if v, ok := mc.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "Member.nickname"`)}
	}
	if v, ok := mc.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Member.phone"`)}
	}
	if v, ok := mc.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Member.email"`)}
	}
	if v, ok := mc.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "Member.avatar"`)}
	}
	if v, ok := mc.mutation.Avatar(); ok {
		if err := member.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Member.avatar": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Member.description"`)}
	}
	if v, ok := mc.mutation.Description(); ok {
		if err := member.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Member.description": %w`, err)}
		}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := member.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Member.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(member.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := mc.mutation.Remark(); ok {
		_spec.SetField(member.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := mc.mutation.Sort(); ok {
		_spec.SetField(member.FieldSort, field.TypeInt32, value)
		_node.Sort = &value
	}
	if value, ok := mc.mutation.State(); ok {
		_spec.SetField(member.FieldState, field.TypeInt32, value)
		_node.State = &value
	}
	if value, ok := mc.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := mc.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := mc.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := mc.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := mc.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(member.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MemberCreate) OnConflict(opts ...sql.ConflictOption) *MemberUpsertOne {
	mc.conflict = opts
	return &MemberUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MemberCreate) OnConflictColumns(columns ...string) *MemberUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertOne{
		create: mc,
	}
}

type (
	// MemberUpsertOne is the builder for "upsert"-ing
	//  one Member node.
	MemberUpsertOne struct {
		create *MemberCreate
	}

	// MemberUpsert is the "OnConflict" setter.
	MemberUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *MemberUpsert) SetUpdatedAt(v time.Time) *MemberUpsert {
	u.Set(member.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MemberUpsert) UpdateUpdatedAt() *MemberUpsert {
	u.SetExcluded(member.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MemberUpsert) ClearUpdatedAt() *MemberUpsert {
	u.SetNull(member.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MemberUpsert) SetDeletedAt(v time.Time) *MemberUpsert {
	u.Set(member.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MemberUpsert) UpdateDeletedAt() *MemberUpsert {
	u.SetExcluded(member.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MemberUpsert) ClearDeletedAt() *MemberUpsert {
	u.SetNull(member.FieldDeletedAt)
	return u
}

// SetRemark sets the "remark" field.
func (u *MemberUpsert) SetRemark(v string) *MemberUpsert {
	u.Set(member.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *MemberUpsert) UpdateRemark() *MemberUpsert {
	u.SetExcluded(member.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *MemberUpsert) ClearRemark() *MemberUpsert {
	u.SetNull(member.FieldRemark)
	return u
}

// SetSort sets the "sort" field.
func (u *MemberUpsert) SetSort(v int32) *MemberUpsert {
	u.Set(member.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *MemberUpsert) UpdateSort() *MemberUpsert {
	u.SetExcluded(member.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *MemberUpsert) AddSort(v int32) *MemberUpsert {
	u.Add(member.FieldSort, v)
	return u
}

// SetState sets the "state" field.
func (u *MemberUpsert) SetState(v int32) *MemberUpsert {
	u.Set(member.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *MemberUpsert) UpdateState() *MemberUpsert {
	u.SetExcluded(member.FieldState)
	return u
}

// AddState adds v to the "state" field.
func (u *MemberUpsert) AddState(v int32) *MemberUpsert {
	u.Add(member.FieldState, v)
	return u
}

// SetPassword sets the "password" field.
func (u *MemberUpsert) SetPassword(v string) *MemberUpsert {
	u.Set(member.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MemberUpsert) UpdatePassword() *MemberUpsert {
	u.SetExcluded(member.FieldPassword)
	return u
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsert) SetNickname(v string) *MemberUpsert {
	u.Set(member.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsert) UpdateNickname() *MemberUpsert {
	u.SetExcluded(member.FieldNickname)
	return u
}

// SetPhone sets the "phone" field.
func (u *MemberUpsert) SetPhone(v string) *MemberUpsert {
	u.Set(member.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsert) UpdatePhone() *MemberUpsert {
	u.SetExcluded(member.FieldPhone)
	return u
}

// SetEmail sets the "email" field.
func (u *MemberUpsert) SetEmail(v string) *MemberUpsert {
	u.Set(member.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MemberUpsert) UpdateEmail() *MemberUpsert {
	u.SetExcluded(member.FieldEmail)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsert) SetAvatar(v string) *MemberUpsert {
	u.Set(member.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsert) UpdateAvatar() *MemberUpsert {
	u.SetExcluded(member.FieldAvatar)
	return u
}

// SetDescription sets the "description" field.
func (u *MemberUpsert) SetDescription(v string) *MemberUpsert {
	u.Set(member.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MemberUpsert) UpdateDescription() *MemberUpsert {
	u.SetExcluded(member.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertOne) UpdateNewValues() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(member.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(member.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Username(); exists {
			s.SetIgnore(member.FieldUsername)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MemberUpsertOne) Ignore() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertOne) DoNothing() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreate.OnConflict
// documentation for more info.
func (u *MemberUpsertOne) Update(set func(*MemberUpsert)) *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MemberUpsertOne) SetUpdatedAt(v time.Time) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateUpdatedAt() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MemberUpsertOne) ClearUpdatedAt() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MemberUpsertOne) SetDeletedAt(v time.Time) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateDeletedAt() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MemberUpsertOne) ClearDeletedAt() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearDeletedAt()
	})
}

// SetRemark sets the "remark" field.
func (u *MemberUpsertOne) SetRemark(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateRemark() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *MemberUpsertOne) ClearRemark() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearRemark()
	})
}

// SetSort sets the "sort" field.
func (u *MemberUpsertOne) SetSort(v int32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *MemberUpsertOne) AddSort(v int32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateSort() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateSort()
	})
}

// SetState sets the "state" field.
func (u *MemberUpsertOne) SetState(v int32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *MemberUpsertOne) AddState(v int32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateState() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateState()
	})
}

// SetPassword sets the "password" field.
func (u *MemberUpsertOne) SetPassword(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdatePassword() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePassword()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertOne) SetNickname(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// SetPhone sets the "phone" field.
func (u *MemberUpsertOne) SetPhone(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdatePhone() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePhone()
	})
}

// SetEmail sets the "email" field.
func (u *MemberUpsertOne) SetEmail(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateEmail() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateEmail()
	})
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsertOne) SetAvatar(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateAvatar() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAvatar()
	})
}

// SetDescription sets the "description" field.
func (u *MemberUpsertOne) SetDescription(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateDescription() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *MemberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MemberUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MemberUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
	conflict []sql.ConflictOption
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflict(opts ...sql.ConflictOption) *MemberUpsertBulk {
	mcb.conflict = opts
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflictColumns(columns ...string) *MemberUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// MemberUpsertBulk is the builder for "upsert"-ing
// a bulk of Member nodes.
type MemberUpsertBulk struct {
	create *MemberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertBulk) UpdateNewValues() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(member.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(member.FieldCreatedAt)
			}
			if _, exists := b.mutation.Username(); exists {
				s.SetIgnore(member.FieldUsername)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MemberUpsertBulk) Ignore() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertBulk) DoNothing() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreateBulk.OnConflict
// documentation for more info.
func (u *MemberUpsertBulk) Update(set func(*MemberUpsert)) *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MemberUpsertBulk) SetUpdatedAt(v time.Time) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateUpdatedAt() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MemberUpsertBulk) ClearUpdatedAt() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MemberUpsertBulk) SetDeletedAt(v time.Time) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateDeletedAt() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MemberUpsertBulk) ClearDeletedAt() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearDeletedAt()
	})
}

// SetRemark sets the "remark" field.
func (u *MemberUpsertBulk) SetRemark(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateRemark() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *MemberUpsertBulk) ClearRemark() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearRemark()
	})
}

// SetSort sets the "sort" field.
func (u *MemberUpsertBulk) SetSort(v int32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *MemberUpsertBulk) AddSort(v int32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateSort() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateSort()
	})
}

// SetState sets the "state" field.
func (u *MemberUpsertBulk) SetState(v int32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *MemberUpsertBulk) AddState(v int32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateState() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateState()
	})
}

// SetPassword sets the "password" field.
func (u *MemberUpsertBulk) SetPassword(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdatePassword() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePassword()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertBulk) SetNickname(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// SetPhone sets the "phone" field.
func (u *MemberUpsertBulk) SetPhone(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdatePhone() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePhone()
	})
}

// SetEmail sets the "email" field.
func (u *MemberUpsertBulk) SetEmail(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateEmail() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateEmail()
	})
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsertBulk) SetAvatar(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateAvatar() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAvatar()
	})
}

// SetDescription sets the "description" field.
func (u *MemberUpsertBulk) SetDescription(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateDescription() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *MemberUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MemberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}