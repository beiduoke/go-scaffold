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
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/role"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RoleCreate) SetDeletedAt(t time.Time) *RoleCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDeletedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetRemark sets the "remark" field.
func (rc *RoleCreate) SetRemark(s string) *RoleCreate {
	rc.mutation.SetRemark(s)
	return rc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (rc *RoleCreate) SetNillableRemark(s *string) *RoleCreate {
	if s != nil {
		rc.SetRemark(*s)
	}
	return rc
}

// SetSort sets the "sort" field.
func (rc *RoleCreate) SetSort(i int32) *RoleCreate {
	rc.mutation.SetSort(i)
	return rc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (rc *RoleCreate) SetNillableSort(i *int32) *RoleCreate {
	if i != nil {
		rc.SetSort(*i)
	}
	return rc
}

// SetState sets the "state" field.
func (rc *RoleCreate) SetState(i int32) *RoleCreate {
	rc.mutation.SetState(i)
	return rc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (rc *RoleCreate) SetNillableState(i *int32) *RoleCreate {
	if i != nil {
		rc.SetState(*i)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rc *RoleCreate) SetNillableName(s *string) *RoleCreate {
	if s != nil {
		rc.SetName(*s)
	}
	return rc
}

// SetDefaultRouter sets the "default_router" field.
func (rc *RoleCreate) SetDefaultRouter(s string) *RoleCreate {
	rc.mutation.SetDefaultRouter(s)
	return rc
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDefaultRouter(s *string) *RoleCreate {
	if s != nil {
		rc.SetDefaultRouter(*s)
	}
	return rc
}

// SetDataScope sets the "data_scope" field.
func (rc *RoleCreate) SetDataScope(i int32) *RoleCreate {
	rc.mutation.SetDataScope(i)
	return rc
}

// SetNillableDataScope sets the "data_scope" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDataScope(i *int32) *RoleCreate {
	if i != nil {
		rc.SetDataScope(*i)
	}
	return rc
}

// SetMenuCheckStrictly sets the "menu_check_strictly" field.
func (rc *RoleCreate) SetMenuCheckStrictly(i int32) *RoleCreate {
	rc.mutation.SetMenuCheckStrictly(i)
	return rc
}

// SetNillableMenuCheckStrictly sets the "menu_check_strictly" field if the given value is not nil.
func (rc *RoleCreate) SetNillableMenuCheckStrictly(i *int32) *RoleCreate {
	if i != nil {
		rc.SetMenuCheckStrictly(*i)
	}
	return rc
}

// SetDeptCheckStrictly sets the "dept_check_strictly" field.
func (rc *RoleCreate) SetDeptCheckStrictly(i int32) *RoleCreate {
	rc.mutation.SetDeptCheckStrictly(i)
	return rc
}

// SetNillableDeptCheckStrictly sets the "dept_check_strictly" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDeptCheckStrictly(i *int32) *RoleCreate {
	if i != nil {
		rc.SetDeptCheckStrictly(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(u uint32) *RoleCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rc *RoleCreate) AddUserIDs(ids ...uint32) *RoleCreate {
	rc.mutation.AddUserIDs(ids...)
	return rc
}

// AddUsers adds the "users" edges to the User entity.
func (rc *RoleCreate) AddUsers(u ...*User) *RoleCreate {
	ids := make([]uint32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.Remark(); !ok {
		v := role.DefaultRemark
		rc.mutation.SetRemark(v)
	}
	if _, ok := rc.mutation.Sort(); !ok {
		v := role.DefaultSort
		rc.mutation.SetSort(v)
	}
	if _, ok := rc.mutation.State(); !ok {
		v := role.DefaultState
		rc.mutation.SetState(v)
	}
	if _, ok := rc.mutation.Name(); !ok {
		v := role.DefaultName
		rc.mutation.SetName(v)
	}
	if _, ok := rc.mutation.DefaultRouter(); !ok {
		v := role.DefaultDefaultRouter
		rc.mutation.SetDefaultRouter(v)
	}
	if _, ok := rc.mutation.DataScope(); !ok {
		v := role.DefaultDataScope
		rc.mutation.SetDataScope(v)
	}
	if _, ok := rc.mutation.MenuCheckStrictly(); !ok {
		v := role.DefaultMenuCheckStrictly
		rc.mutation.SetMenuCheckStrictly(v)
	}
	if _, ok := rc.mutation.DeptCheckStrictly(); !ok {
		v := role.DefaultDeptCheckStrictly
		rc.mutation.SetDeptCheckStrictly(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Role.sort"`)}
	}
	if v, ok := rc.mutation.Sort(); ok {
		if err := role.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Role.sort": %w`, err)}
		}
	}
	if _, ok := rc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Role.state"`)}
	}
	if v, ok := rc.mutation.State(); ok {
		if err := role.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Role.state": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Role.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.DefaultRouter(); !ok {
		return &ValidationError{Name: "default_router", err: errors.New(`ent: missing required field "Role.default_router"`)}
	}
	if v, ok := rc.mutation.DefaultRouter(); ok {
		if err := role.DefaultRouterValidator(v); err != nil {
			return &ValidationError{Name: "default_router", err: fmt.Errorf(`ent: validator failed for field "Role.default_router": %w`, err)}
		}
	}
	if _, ok := rc.mutation.DataScope(); !ok {
		return &ValidationError{Name: "data_scope", err: errors.New(`ent: missing required field "Role.data_scope"`)}
	}
	if _, ok := rc.mutation.MenuCheckStrictly(); !ok {
		return &ValidationError{Name: "menu_check_strictly", err: errors.New(`ent: missing required field "Role.menu_check_strictly"`)}
	}
	if _, ok := rc.mutation.DeptCheckStrictly(); !ok {
		return &ValidationError{Name: "dept_check_strictly", err: errors.New(`ent: missing required field "Role.dept_check_strictly"`)}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := role.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Role.id": %w`, err)}
		}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := rc.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := rc.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt32, value)
		_node.Sort = &value
	}
	if value, ok := rc.mutation.State(); ok {
		_spec.SetField(role.FieldState, field.TypeInt32, value)
		_node.State = &value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := rc.mutation.DefaultRouter(); ok {
		_spec.SetField(role.FieldDefaultRouter, field.TypeString, value)
		_node.DefaultRouter = &value
	}
	if value, ok := rc.mutation.DataScope(); ok {
		_spec.SetField(role.FieldDataScope, field.TypeInt32, value)
		_node.DataScope = &value
	}
	if value, ok := rc.mutation.MenuCheckStrictly(); ok {
		_spec.SetField(role.FieldMenuCheckStrictly, field.TypeInt32, value)
		_node.MenuCheckStrictly = &value
	}
	if value, ok := rc.mutation.DeptCheckStrictly(); ok {
		_spec.SetField(role.FieldDeptCheckStrictly, field.TypeInt32, value)
		_node.DeptCheckStrictly = &value
	}
	if nodes := rc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RoleCreate) OnConflict(opts ...sql.ConflictOption) *RoleUpsertOne {
	rc.conflict = opts
	return &RoleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoleCreate) OnConflictColumns(columns ...string) *RoleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertOne{
		create: rc,
	}
}

type (
	// RoleUpsertOne is the builder for "upsert"-ing
	//  one Role node.
	RoleUpsertOne struct {
		create *RoleCreate
	}

	// RoleUpsert is the "OnConflict" setter.
	RoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsert) SetUpdatedAt(v time.Time) *RoleUpsert {
	u.Set(role.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateUpdatedAt() *RoleUpsert {
	u.SetExcluded(role.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RoleUpsert) ClearUpdatedAt() *RoleUpsert {
	u.SetNull(role.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsert) SetDeletedAt(v time.Time) *RoleUpsert {
	u.Set(role.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDeletedAt() *RoleUpsert {
	u.SetExcluded(role.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsert) ClearDeletedAt() *RoleUpsert {
	u.SetNull(role.FieldDeletedAt)
	return u
}

// SetRemark sets the "remark" field.
func (u *RoleUpsert) SetRemark(v string) *RoleUpsert {
	u.Set(role.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *RoleUpsert) UpdateRemark() *RoleUpsert {
	u.SetExcluded(role.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *RoleUpsert) ClearRemark() *RoleUpsert {
	u.SetNull(role.FieldRemark)
	return u
}

// SetSort sets the "sort" field.
func (u *RoleUpsert) SetSort(v int32) *RoleUpsert {
	u.Set(role.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsert) UpdateSort() *RoleUpsert {
	u.SetExcluded(role.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsert) AddSort(v int32) *RoleUpsert {
	u.Add(role.FieldSort, v)
	return u
}

// SetState sets the "state" field.
func (u *RoleUpsert) SetState(v int32) *RoleUpsert {
	u.Set(role.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *RoleUpsert) UpdateState() *RoleUpsert {
	u.SetExcluded(role.FieldState)
	return u
}

// AddState adds v to the "state" field.
func (u *RoleUpsert) AddState(v int32) *RoleUpsert {
	u.Add(role.FieldState, v)
	return u
}

// SetName sets the "name" field.
func (u *RoleUpsert) SetName(v string) *RoleUpsert {
	u.Set(role.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsert) UpdateName() *RoleUpsert {
	u.SetExcluded(role.FieldName)
	return u
}

// SetDefaultRouter sets the "default_router" field.
func (u *RoleUpsert) SetDefaultRouter(v string) *RoleUpsert {
	u.Set(role.FieldDefaultRouter, v)
	return u
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDefaultRouter() *RoleUpsert {
	u.SetExcluded(role.FieldDefaultRouter)
	return u
}

// SetDataScope sets the "data_scope" field.
func (u *RoleUpsert) SetDataScope(v int32) *RoleUpsert {
	u.Set(role.FieldDataScope, v)
	return u
}

// UpdateDataScope sets the "data_scope" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDataScope() *RoleUpsert {
	u.SetExcluded(role.FieldDataScope)
	return u
}

// AddDataScope adds v to the "data_scope" field.
func (u *RoleUpsert) AddDataScope(v int32) *RoleUpsert {
	u.Add(role.FieldDataScope, v)
	return u
}

// SetMenuCheckStrictly sets the "menu_check_strictly" field.
func (u *RoleUpsert) SetMenuCheckStrictly(v int32) *RoleUpsert {
	u.Set(role.FieldMenuCheckStrictly, v)
	return u
}

// UpdateMenuCheckStrictly sets the "menu_check_strictly" field to the value that was provided on create.
func (u *RoleUpsert) UpdateMenuCheckStrictly() *RoleUpsert {
	u.SetExcluded(role.FieldMenuCheckStrictly)
	return u
}

// AddMenuCheckStrictly adds v to the "menu_check_strictly" field.
func (u *RoleUpsert) AddMenuCheckStrictly(v int32) *RoleUpsert {
	u.Add(role.FieldMenuCheckStrictly, v)
	return u
}

// SetDeptCheckStrictly sets the "dept_check_strictly" field.
func (u *RoleUpsert) SetDeptCheckStrictly(v int32) *RoleUpsert {
	u.Set(role.FieldDeptCheckStrictly, v)
	return u
}

// UpdateDeptCheckStrictly sets the "dept_check_strictly" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDeptCheckStrictly() *RoleUpsert {
	u.SetExcluded(role.FieldDeptCheckStrictly)
	return u
}

// AddDeptCheckStrictly adds v to the "dept_check_strictly" field.
func (u *RoleUpsert) AddDeptCheckStrictly(v int32) *RoleUpsert {
	u.Add(role.FieldDeptCheckStrictly, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertOne) UpdateNewValues() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(role.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(role.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUpsertOne) Ignore() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertOne) DoNothing() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreate.OnConflict
// documentation for more info.
func (u *RoleUpsertOne) Update(set func(*RoleUpsert)) *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertOne) SetUpdatedAt(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateUpdatedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RoleUpsertOne) ClearUpdatedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertOne) SetDeletedAt(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDeletedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsertOne) ClearDeletedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetRemark sets the "remark" field.
func (u *RoleUpsertOne) SetRemark(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateRemark() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *RoleUpsertOne) ClearRemark() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearRemark()
	})
}

// SetSort sets the "sort" field.
func (u *RoleUpsertOne) SetSort(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsertOne) AddSort(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateSort() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateSort()
	})
}

// SetState sets the "state" field.
func (u *RoleUpsertOne) SetState(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *RoleUpsertOne) AddState(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateState() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateState()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertOne) SetName(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateName() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetDefaultRouter sets the "default_router" field.
func (u *RoleUpsertOne) SetDefaultRouter(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDefaultRouter(v)
	})
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDefaultRouter() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDefaultRouter()
	})
}

// SetDataScope sets the "data_scope" field.
func (u *RoleUpsertOne) SetDataScope(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDataScope(v)
	})
}

// AddDataScope adds v to the "data_scope" field.
func (u *RoleUpsertOne) AddDataScope(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddDataScope(v)
	})
}

// UpdateDataScope sets the "data_scope" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDataScope() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDataScope()
	})
}

// SetMenuCheckStrictly sets the "menu_check_strictly" field.
func (u *RoleUpsertOne) SetMenuCheckStrictly(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetMenuCheckStrictly(v)
	})
}

// AddMenuCheckStrictly adds v to the "menu_check_strictly" field.
func (u *RoleUpsertOne) AddMenuCheckStrictly(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddMenuCheckStrictly(v)
	})
}

// UpdateMenuCheckStrictly sets the "menu_check_strictly" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateMenuCheckStrictly() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateMenuCheckStrictly()
	})
}

// SetDeptCheckStrictly sets the "dept_check_strictly" field.
func (u *RoleUpsertOne) SetDeptCheckStrictly(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeptCheckStrictly(v)
	})
}

// AddDeptCheckStrictly adds v to the "dept_check_strictly" field.
func (u *RoleUpsertOne) AddDeptCheckStrictly(v int32) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeptCheckStrictly(v)
	})
}

// UpdateDeptCheckStrictly sets the "dept_check_strictly" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDeptCheckStrictly() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeptCheckStrictly()
	})
}

// Exec executes the query.
func (u *RoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoleUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoleUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	err      error
	builders []*RoleCreate
	conflict []sql.ConflictOption
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUpsertBulk {
	rcb.conflict = opts
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflictColumns(columns ...string) *RoleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// RoleUpsertBulk is the builder for "upsert"-ing
// a bulk of Role nodes.
type RoleUpsertBulk struct {
	create *RoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertBulk) UpdateNewValues() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(role.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(role.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUpsertBulk) Ignore() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertBulk) DoNothing() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUpsertBulk) Update(set func(*RoleUpsert)) *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertBulk) SetUpdatedAt(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateUpdatedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RoleUpsertBulk) ClearUpdatedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertBulk) SetDeletedAt(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDeletedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsertBulk) ClearDeletedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetRemark sets the "remark" field.
func (u *RoleUpsertBulk) SetRemark(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateRemark() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *RoleUpsertBulk) ClearRemark() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearRemark()
	})
}

// SetSort sets the "sort" field.
func (u *RoleUpsertBulk) SetSort(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsertBulk) AddSort(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateSort() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateSort()
	})
}

// SetState sets the "state" field.
func (u *RoleUpsertBulk) SetState(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *RoleUpsertBulk) AddState(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateState() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateState()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertBulk) SetName(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateName() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetDefaultRouter sets the "default_router" field.
func (u *RoleUpsertBulk) SetDefaultRouter(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDefaultRouter(v)
	})
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDefaultRouter() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDefaultRouter()
	})
}

// SetDataScope sets the "data_scope" field.
func (u *RoleUpsertBulk) SetDataScope(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDataScope(v)
	})
}

// AddDataScope adds v to the "data_scope" field.
func (u *RoleUpsertBulk) AddDataScope(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddDataScope(v)
	})
}

// UpdateDataScope sets the "data_scope" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDataScope() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDataScope()
	})
}

// SetMenuCheckStrictly sets the "menu_check_strictly" field.
func (u *RoleUpsertBulk) SetMenuCheckStrictly(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetMenuCheckStrictly(v)
	})
}

// AddMenuCheckStrictly adds v to the "menu_check_strictly" field.
func (u *RoleUpsertBulk) AddMenuCheckStrictly(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddMenuCheckStrictly(v)
	})
}

// UpdateMenuCheckStrictly sets the "menu_check_strictly" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateMenuCheckStrictly() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateMenuCheckStrictly()
	})
}

// SetDeptCheckStrictly sets the "dept_check_strictly" field.
func (u *RoleUpsertBulk) SetDeptCheckStrictly(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeptCheckStrictly(v)
	})
}

// AddDeptCheckStrictly adds v to the "dept_check_strictly" field.
func (u *RoleUpsertBulk) AddDeptCheckStrictly(v int32) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeptCheckStrictly(v)
	})
}

// UpdateDeptCheckStrictly sets the "dept_check_strictly" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDeptCheckStrictly() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeptCheckStrictly()
	})
}

// Exec executes the query.
func (u *RoleUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
