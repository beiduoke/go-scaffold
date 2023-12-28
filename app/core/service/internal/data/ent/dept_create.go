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
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/dept"
)

// DeptCreate is the builder for creating a Dept entity.
type DeptCreate struct {
	config
	mutation *DeptMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeptCreate) SetCreatedAt(t time.Time) *DeptCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeptCreate) SetNillableCreatedAt(t *time.Time) *DeptCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeptCreate) SetUpdatedAt(t time.Time) *DeptCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeptCreate) SetNillableUpdatedAt(t *time.Time) *DeptCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DeptCreate) SetDeletedAt(t time.Time) *DeptCreate {
	dc.mutation.SetDeletedAt(t)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DeptCreate) SetNillableDeletedAt(t *time.Time) *DeptCreate {
	if t != nil {
		dc.SetDeletedAt(*t)
	}
	return dc
}

// SetPlatformID sets the "platform_id" field.
func (dc *DeptCreate) SetPlatformID(u uint64) *DeptCreate {
	dc.mutation.SetPlatformID(u)
	return dc
}

// SetNillablePlatformID sets the "platform_id" field if the given value is not nil.
func (dc *DeptCreate) SetNillablePlatformID(u *uint64) *DeptCreate {
	if u != nil {
		dc.SetPlatformID(*u)
	}
	return dc
}

// SetSort sets the "sort" field.
func (dc *DeptCreate) SetSort(i int32) *DeptCreate {
	dc.mutation.SetSort(i)
	return dc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dc *DeptCreate) SetNillableSort(i *int32) *DeptCreate {
	if i != nil {
		dc.SetSort(*i)
	}
	return dc
}

// SetRemark sets the "remark" field.
func (dc *DeptCreate) SetRemark(s string) *DeptCreate {
	dc.mutation.SetRemark(s)
	return dc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dc *DeptCreate) SetNillableRemark(s *string) *DeptCreate {
	if s != nil {
		dc.SetRemark(*s)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DeptCreate) SetStatus(d dept.Status) *DeptCreate {
	dc.mutation.SetStatus(d)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DeptCreate) SetNillableStatus(d *dept.Status) *DeptCreate {
	if d != nil {
		dc.SetStatus(*d)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DeptCreate) SetName(s string) *DeptCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DeptCreate) SetNillableName(s *string) *DeptCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetParentID sets the "parent_id" field.
func (dc *DeptCreate) SetParentID(i int32) *DeptCreate {
	dc.mutation.SetParentID(i)
	return dc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (dc *DeptCreate) SetNillableParentID(i *int32) *DeptCreate {
	if i != nil {
		dc.SetParentID(*i)
	}
	return dc
}

// SetAncestors sets the "ancestors" field.
func (dc *DeptCreate) SetAncestors(i []int) *DeptCreate {
	dc.mutation.SetAncestors(i)
	return dc
}

// SetID sets the "id" field.
func (dc *DeptCreate) SetID(u uint32) *DeptCreate {
	dc.mutation.SetID(u)
	return dc
}

// Mutation returns the DeptMutation object of the builder.
func (dc *DeptCreate) Mutation() *DeptMutation {
	return dc.mutation
}

// Save creates the Dept in the database.
func (dc *DeptCreate) Save(ctx context.Context) (*Dept, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeptCreate) SaveX(ctx context.Context) *Dept {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeptCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeptCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeptCreate) defaults() {
	if _, ok := dc.mutation.PlatformID(); !ok {
		v := dept.DefaultPlatformID()
		dc.mutation.SetPlatformID(v)
	}
	if _, ok := dc.mutation.Sort(); !ok {
		v := dept.DefaultSort
		dc.mutation.SetSort(v)
	}
	if _, ok := dc.mutation.Remark(); !ok {
		v := dept.DefaultRemark
		dc.mutation.SetRemark(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := dept.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.ParentID(); !ok {
		v := dept.DefaultParentID
		dc.mutation.SetParentID(v)
	}
	if _, ok := dc.mutation.Ancestors(); !ok {
		v := dept.DefaultAncestors
		dc.mutation.SetAncestors(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeptCreate) check() error {
	if _, ok := dc.mutation.PlatformID(); !ok {
		return &ValidationError{Name: "platform_id", err: errors.New(`ent: missing required field "Dept.platform_id"`)}
	}
	if v, ok := dc.mutation.PlatformID(); ok {
		if err := dept.PlatformIDValidator(v); err != nil {
			return &ValidationError{Name: "platform_id", err: fmt.Errorf(`ent: validator failed for field "Dept.platform_id": %w`, err)}
		}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := dept.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Dept.status": %w`, err)}
		}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := dept.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Dept.name": %w`, err)}
		}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := dept.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Dept.id": %w`, err)}
		}
	}
	return nil
}

func (dc *DeptCreate) sqlSave(ctx context.Context) (*Dept, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeptCreate) createSpec() (*Dept, *sqlgraph.CreateSpec) {
	var (
		_node = &Dept{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dept.Table, sqlgraph.NewFieldSpec(dept.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(dept.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(dept.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(dept.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := dc.mutation.PlatformID(); ok {
		_spec.SetField(dept.FieldPlatformID, field.TypeUint64, value)
		_node.PlatformID = value
	}
	if value, ok := dc.mutation.Sort(); ok {
		_spec.SetField(dept.FieldSort, field.TypeInt32, value)
		_node.Sort = &value
	}
	if value, ok := dc.mutation.Remark(); ok {
		_spec.SetField(dept.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(dept.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(dept.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := dc.mutation.ParentID(); ok {
		_spec.SetField(dept.FieldParentID, field.TypeInt32, value)
		_node.ParentID = &value
	}
	if value, ok := dc.mutation.Ancestors(); ok {
		_spec.SetField(dept.FieldAncestors, field.TypeJSON, value)
		_node.Ancestors = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Dept.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dc *DeptCreate) OnConflict(opts ...sql.ConflictOption) *DeptUpsertOne {
	dc.conflict = opts
	return &DeptUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Dept.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DeptCreate) OnConflictColumns(columns ...string) *DeptUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DeptUpsertOne{
		create: dc,
	}
}

type (
	// DeptUpsertOne is the builder for "upsert"-ing
	//  one Dept node.
	DeptUpsertOne struct {
		create *DeptCreate
	}

	// DeptUpsert is the "OnConflict" setter.
	DeptUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptUpsert) SetUpdatedAt(v time.Time) *DeptUpsert {
	u.Set(dept.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptUpsert) UpdateUpdatedAt() *DeptUpsert {
	u.SetExcluded(dept.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeptUpsert) ClearUpdatedAt() *DeptUpsert {
	u.SetNull(dept.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeptUpsert) SetDeletedAt(v time.Time) *DeptUpsert {
	u.Set(dept.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeptUpsert) UpdateDeletedAt() *DeptUpsert {
	u.SetExcluded(dept.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeptUpsert) ClearDeletedAt() *DeptUpsert {
	u.SetNull(dept.FieldDeletedAt)
	return u
}

// SetPlatformID sets the "platform_id" field.
func (u *DeptUpsert) SetPlatformID(v uint64) *DeptUpsert {
	u.Set(dept.FieldPlatformID, v)
	return u
}

// UpdatePlatformID sets the "platform_id" field to the value that was provided on create.
func (u *DeptUpsert) UpdatePlatformID() *DeptUpsert {
	u.SetExcluded(dept.FieldPlatformID)
	return u
}

// AddPlatformID adds v to the "platform_id" field.
func (u *DeptUpsert) AddPlatformID(v uint64) *DeptUpsert {
	u.Add(dept.FieldPlatformID, v)
	return u
}

// SetSort sets the "sort" field.
func (u *DeptUpsert) SetSort(v int32) *DeptUpsert {
	u.Set(dept.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DeptUpsert) UpdateSort() *DeptUpsert {
	u.SetExcluded(dept.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *DeptUpsert) AddSort(v int32) *DeptUpsert {
	u.Add(dept.FieldSort, v)
	return u
}

// ClearSort clears the value of the "sort" field.
func (u *DeptUpsert) ClearSort() *DeptUpsert {
	u.SetNull(dept.FieldSort)
	return u
}

// SetRemark sets the "remark" field.
func (u *DeptUpsert) SetRemark(v string) *DeptUpsert {
	u.Set(dept.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DeptUpsert) UpdateRemark() *DeptUpsert {
	u.SetExcluded(dept.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *DeptUpsert) ClearRemark() *DeptUpsert {
	u.SetNull(dept.FieldRemark)
	return u
}

// SetStatus sets the "status" field.
func (u *DeptUpsert) SetStatus(v dept.Status) *DeptUpsert {
	u.Set(dept.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *DeptUpsert) UpdateStatus() *DeptUpsert {
	u.SetExcluded(dept.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *DeptUpsert) ClearStatus() *DeptUpsert {
	u.SetNull(dept.FieldStatus)
	return u
}

// SetName sets the "name" field.
func (u *DeptUpsert) SetName(v string) *DeptUpsert {
	u.Set(dept.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeptUpsert) UpdateName() *DeptUpsert {
	u.SetExcluded(dept.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *DeptUpsert) ClearName() *DeptUpsert {
	u.SetNull(dept.FieldName)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *DeptUpsert) SetParentID(v int32) *DeptUpsert {
	u.Set(dept.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DeptUpsert) UpdateParentID() *DeptUpsert {
	u.SetExcluded(dept.FieldParentID)
	return u
}

// AddParentID adds v to the "parent_id" field.
func (u *DeptUpsert) AddParentID(v int32) *DeptUpsert {
	u.Add(dept.FieldParentID, v)
	return u
}

// ClearParentID clears the value of the "parent_id" field.
func (u *DeptUpsert) ClearParentID() *DeptUpsert {
	u.SetNull(dept.FieldParentID)
	return u
}

// SetAncestors sets the "ancestors" field.
func (u *DeptUpsert) SetAncestors(v []int) *DeptUpsert {
	u.Set(dept.FieldAncestors, v)
	return u
}

// UpdateAncestors sets the "ancestors" field to the value that was provided on create.
func (u *DeptUpsert) UpdateAncestors() *DeptUpsert {
	u.SetExcluded(dept.FieldAncestors)
	return u
}

// ClearAncestors clears the value of the "ancestors" field.
func (u *DeptUpsert) ClearAncestors() *DeptUpsert {
	u.SetNull(dept.FieldAncestors)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Dept.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dept.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeptUpsertOne) UpdateNewValues() *DeptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(dept.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(dept.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Dept.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeptUpsertOne) Ignore() *DeptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeptUpsertOne) DoNothing() *DeptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeptCreate.OnConflict
// documentation for more info.
func (u *DeptUpsertOne) Update(set func(*DeptUpsert)) *DeptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptUpsertOne) SetUpdatedAt(v time.Time) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateUpdatedAt() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeptUpsertOne) ClearUpdatedAt() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeptUpsertOne) SetDeletedAt(v time.Time) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateDeletedAt() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeptUpsertOne) ClearDeletedAt() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearDeletedAt()
	})
}

// SetPlatformID sets the "platform_id" field.
func (u *DeptUpsertOne) SetPlatformID(v uint64) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetPlatformID(v)
	})
}

// AddPlatformID adds v to the "platform_id" field.
func (u *DeptUpsertOne) AddPlatformID(v uint64) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.AddPlatformID(v)
	})
}

// UpdatePlatformID sets the "platform_id" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdatePlatformID() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdatePlatformID()
	})
}

// SetSort sets the "sort" field.
func (u *DeptUpsertOne) SetSort(v int32) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *DeptUpsertOne) AddSort(v int32) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateSort() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateSort()
	})
}

// ClearSort clears the value of the "sort" field.
func (u *DeptUpsertOne) ClearSort() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearSort()
	})
}

// SetRemark sets the "remark" field.
func (u *DeptUpsertOne) SetRemark(v string) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateRemark() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DeptUpsertOne) ClearRemark() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearRemark()
	})
}

// SetStatus sets the "status" field.
func (u *DeptUpsertOne) SetStatus(v dept.Status) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateStatus() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *DeptUpsertOne) ClearStatus() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearStatus()
	})
}

// SetName sets the "name" field.
func (u *DeptUpsertOne) SetName(v string) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateName() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeptUpsertOne) ClearName() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *DeptUpsertOne) SetParentID(v int32) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetParentID(v)
	})
}

// AddParentID adds v to the "parent_id" field.
func (u *DeptUpsertOne) AddParentID(v int32) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.AddParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateParentID() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *DeptUpsertOne) ClearParentID() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearParentID()
	})
}

// SetAncestors sets the "ancestors" field.
func (u *DeptUpsertOne) SetAncestors(v []int) *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.SetAncestors(v)
	})
}

// UpdateAncestors sets the "ancestors" field to the value that was provided on create.
func (u *DeptUpsertOne) UpdateAncestors() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateAncestors()
	})
}

// ClearAncestors clears the value of the "ancestors" field.
func (u *DeptUpsertOne) ClearAncestors() *DeptUpsertOne {
	return u.Update(func(s *DeptUpsert) {
		s.ClearAncestors()
	})
}

// Exec executes the query.
func (u *DeptUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeptCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeptUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeptUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeptUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeptCreateBulk is the builder for creating many Dept entities in bulk.
type DeptCreateBulk struct {
	config
	err      error
	builders []*DeptCreate
	conflict []sql.ConflictOption
}

// Save creates the Dept entities in the database.
func (dcb *DeptCreateBulk) Save(ctx context.Context) ([]*Dept, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dept, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeptMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeptCreateBulk) SaveX(ctx context.Context) []*Dept {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeptCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeptCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Dept.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dcb *DeptCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeptUpsertBulk {
	dcb.conflict = opts
	return &DeptUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Dept.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DeptCreateBulk) OnConflictColumns(columns ...string) *DeptUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DeptUpsertBulk{
		create: dcb,
	}
}

// DeptUpsertBulk is the builder for "upsert"-ing
// a bulk of Dept nodes.
type DeptUpsertBulk struct {
	create *DeptCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Dept.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dept.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeptUpsertBulk) UpdateNewValues() *DeptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(dept.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(dept.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Dept.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeptUpsertBulk) Ignore() *DeptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeptUpsertBulk) DoNothing() *DeptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeptCreateBulk.OnConflict
// documentation for more info.
func (u *DeptUpsertBulk) Update(set func(*DeptUpsert)) *DeptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptUpsertBulk) SetUpdatedAt(v time.Time) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateUpdatedAt() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeptUpsertBulk) ClearUpdatedAt() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeptUpsertBulk) SetDeletedAt(v time.Time) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateDeletedAt() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeptUpsertBulk) ClearDeletedAt() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearDeletedAt()
	})
}

// SetPlatformID sets the "platform_id" field.
func (u *DeptUpsertBulk) SetPlatformID(v uint64) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetPlatformID(v)
	})
}

// AddPlatformID adds v to the "platform_id" field.
func (u *DeptUpsertBulk) AddPlatformID(v uint64) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.AddPlatformID(v)
	})
}

// UpdatePlatformID sets the "platform_id" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdatePlatformID() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdatePlatformID()
	})
}

// SetSort sets the "sort" field.
func (u *DeptUpsertBulk) SetSort(v int32) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *DeptUpsertBulk) AddSort(v int32) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateSort() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateSort()
	})
}

// ClearSort clears the value of the "sort" field.
func (u *DeptUpsertBulk) ClearSort() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearSort()
	})
}

// SetRemark sets the "remark" field.
func (u *DeptUpsertBulk) SetRemark(v string) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateRemark() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DeptUpsertBulk) ClearRemark() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearRemark()
	})
}

// SetStatus sets the "status" field.
func (u *DeptUpsertBulk) SetStatus(v dept.Status) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateStatus() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *DeptUpsertBulk) ClearStatus() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearStatus()
	})
}

// SetName sets the "name" field.
func (u *DeptUpsertBulk) SetName(v string) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateName() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeptUpsertBulk) ClearName() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *DeptUpsertBulk) SetParentID(v int32) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetParentID(v)
	})
}

// AddParentID adds v to the "parent_id" field.
func (u *DeptUpsertBulk) AddParentID(v int32) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.AddParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateParentID() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *DeptUpsertBulk) ClearParentID() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearParentID()
	})
}

// SetAncestors sets the "ancestors" field.
func (u *DeptUpsertBulk) SetAncestors(v []int) *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.SetAncestors(v)
	})
}

// UpdateAncestors sets the "ancestors" field to the value that was provided on create.
func (u *DeptUpsertBulk) UpdateAncestors() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.UpdateAncestors()
	})
}

// ClearAncestors clears the value of the "ancestors" field.
func (u *DeptUpsertBulk) ClearAncestors() *DeptUpsertBulk {
	return u.Update(func(s *DeptUpsert) {
		s.ClearAncestors()
	})
}

// Exec executes the query.
func (u *DeptUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeptCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeptCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeptUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
