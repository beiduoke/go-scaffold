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
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks     []Hook
	mutation  *MemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MemberUpdate) SetUpdatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableUpdatedAt(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetUpdatedAt(*t)
	}
	return mu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mu *MemberUpdate) ClearUpdatedAt() *MemberUpdate {
	mu.mutation.ClearUpdatedAt()
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MemberUpdate) SetDeletedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableDeletedAt(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MemberUpdate) ClearDeletedAt() *MemberUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// SetRemark sets the "remark" field.
func (mu *MemberUpdate) SetRemark(s string) *MemberUpdate {
	mu.mutation.SetRemark(s)
	return mu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableRemark(s *string) *MemberUpdate {
	if s != nil {
		mu.SetRemark(*s)
	}
	return mu
}

// ClearRemark clears the value of the "remark" field.
func (mu *MemberUpdate) ClearRemark() *MemberUpdate {
	mu.mutation.ClearRemark()
	return mu
}

// SetSort sets the "sort" field.
func (mu *MemberUpdate) SetSort(i int32) *MemberUpdate {
	mu.mutation.ResetSort()
	mu.mutation.SetSort(i)
	return mu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableSort(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetSort(*i)
	}
	return mu
}

// AddSort adds i to the "sort" field.
func (mu *MemberUpdate) AddSort(i int32) *MemberUpdate {
	mu.mutation.AddSort(i)
	return mu
}

// SetState sets the "state" field.
func (mu *MemberUpdate) SetState(i int32) *MemberUpdate {
	mu.mutation.ResetState()
	mu.mutation.SetState(i)
	return mu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableState(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetState(*i)
	}
	return mu
}

// AddState adds i to the "state" field.
func (mu *MemberUpdate) AddState(i int32) *MemberUpdate {
	mu.mutation.AddState(i)
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePassword(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPassword(*s)
	}
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickname(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickname(*s)
	}
	return mu
}

// SetPhone sets the "phone" field.
func (mu *MemberUpdate) SetPhone(s string) *MemberUpdate {
	mu.mutation.SetPhone(s)
	return mu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePhone(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPhone(*s)
	}
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmail(s *string) *MemberUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MemberUpdate) SetAvatar(s string) *MemberUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAvatar(s *string) *MemberUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// SetDescription sets the "description" field.
func (mu *MemberUpdate) SetDescription(s string) *MemberUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableDescription(s *string) *MemberUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Sort(); ok {
		if err := member.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Member.sort": %w`, err)}
		}
	}
	if v, ok := mu.mutation.State(); ok {
		if err := member.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Member.state": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Avatar(); ok {
		if err := member.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Member.avatar": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Description(); ok {
		if err := member.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Member.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MemberUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUint32))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mu.mutation.CreatedAtCleared() {
		_spec.ClearField(member.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if mu.mutation.UpdatedAtCleared() {
		_spec.ClearField(member.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(member.FieldDeletedAt, field.TypeTime, value)
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.ClearField(member.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.Remark(); ok {
		_spec.SetField(member.FieldRemark, field.TypeString, value)
	}
	if mu.mutation.RemarkCleared() {
		_spec.ClearField(member.FieldRemark, field.TypeString)
	}
	if value, ok := mu.mutation.Sort(); ok {
		_spec.SetField(member.FieldSort, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedSort(); ok {
		_spec.AddField(member.FieldSort, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.State(); ok {
		_spec.SetField(member.FieldState, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedState(); ok {
		_spec.AddField(member.FieldState, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := mu.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(member.FieldDescription, field.TypeString, value)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MemberUpdateOne) SetUpdatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableUpdatedAt(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetUpdatedAt(*t)
	}
	return muo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (muo *MemberUpdateOne) ClearUpdatedAt() *MemberUpdateOne {
	muo.mutation.ClearUpdatedAt()
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MemberUpdateOne) SetDeletedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableDeletedAt(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MemberUpdateOne) ClearDeletedAt() *MemberUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// SetRemark sets the "remark" field.
func (muo *MemberUpdateOne) SetRemark(s string) *MemberUpdateOne {
	muo.mutation.SetRemark(s)
	return muo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRemark(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetRemark(*s)
	}
	return muo
}

// ClearRemark clears the value of the "remark" field.
func (muo *MemberUpdateOne) ClearRemark() *MemberUpdateOne {
	muo.mutation.ClearRemark()
	return muo
}

// SetSort sets the "sort" field.
func (muo *MemberUpdateOne) SetSort(i int32) *MemberUpdateOne {
	muo.mutation.ResetSort()
	muo.mutation.SetSort(i)
	return muo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableSort(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetSort(*i)
	}
	return muo
}

// AddSort adds i to the "sort" field.
func (muo *MemberUpdateOne) AddSort(i int32) *MemberUpdateOne {
	muo.mutation.AddSort(i)
	return muo
}

// SetState sets the "state" field.
func (muo *MemberUpdateOne) SetState(i int32) *MemberUpdateOne {
	muo.mutation.ResetState()
	muo.mutation.SetState(i)
	return muo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableState(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetState(*i)
	}
	return muo
}

// AddState adds i to the "state" field.
func (muo *MemberUpdateOne) AddState(i int32) *MemberUpdateOne {
	muo.mutation.AddState(i)
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePassword(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPassword(*s)
	}
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickname(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickname(*s)
	}
	return muo
}

// SetPhone sets the "phone" field.
func (muo *MemberUpdateOne) SetPhone(s string) *MemberUpdateOne {
	muo.mutation.SetPhone(s)
	return muo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePhone(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPhone(*s)
	}
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmail(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MemberUpdateOne) SetAvatar(s string) *MemberUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAvatar(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// SetDescription sets the "description" field.
func (muo *MemberUpdateOne) SetDescription(s string) *MemberUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableDescription(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Sort(); ok {
		if err := member.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Member.sort": %w`, err)}
		}
	}
	if v, ok := muo.mutation.State(); ok {
		if err := member.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Member.state": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Avatar(); ok {
		if err := member.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Member.avatar": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Description(); ok {
		if err := member.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Member.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MemberUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUint32))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if muo.mutation.CreatedAtCleared() {
		_spec.ClearField(member.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if muo.mutation.UpdatedAtCleared() {
		_spec.ClearField(member.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(member.FieldDeletedAt, field.TypeTime, value)
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.ClearField(member.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.Remark(); ok {
		_spec.SetField(member.FieldRemark, field.TypeString, value)
	}
	if muo.mutation.RemarkCleared() {
		_spec.ClearField(member.FieldRemark, field.TypeString)
	}
	if value, ok := muo.mutation.Sort(); ok {
		_spec.SetField(member.FieldSort, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedSort(); ok {
		_spec.AddField(member.FieldSort, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.State(); ok {
		_spec.SetField(member.FieldState, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedState(); ok {
		_spec.AddField(member.FieldState, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := muo.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(member.FieldDescription, field.TypeString, value)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
