// Code generated by ent, DO NOT EDIT.

package role

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDefaultRouter holds the string denoting the default_router field in the database.
	FieldDefaultRouter = "default_router"
	// FieldDataScope holds the string denoting the data_scope field in the database.
	FieldDataScope = "data_scope"
	// FieldMenuCheckStrictly holds the string denoting the menu_check_strictly field in the database.
	FieldMenuCheckStrictly = "menu_check_strictly"
	// FieldDeptCheckStrictly holds the string denoting the dept_check_strictly field in the database.
	FieldDeptCheckStrictly = "dept_check_strictly"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_roles"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldRemark,
	FieldSort,
	FieldState,
	FieldName,
	FieldDefaultRouter,
	FieldDataScope,
	FieldMenuCheckStrictly,
	FieldDeptCheckStrictly,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// SortValidator is a validator for the "sort" field. It is called by the builders before save.
	SortValidator func(int32) error
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState int32
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(int32) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDefaultRouter holds the default value on creation for the "default_router" field.
	DefaultDefaultRouter string
	// DefaultRouterValidator is a validator for the "default_router" field. It is called by the builders before save.
	DefaultRouterValidator func(string) error
	// DefaultDataScope holds the default value on creation for the "data_scope" field.
	DefaultDataScope int32
	// DefaultMenuCheckStrictly holds the default value on creation for the "menu_check_strictly" field.
	DefaultMenuCheckStrictly int32
	// DefaultDeptCheckStrictly holds the default value on creation for the "dept_check_strictly" field.
	DefaultDeptCheckStrictly int32
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// OrderOption defines the ordering options for the Role queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDefaultRouter orders the results by the default_router field.
func ByDefaultRouter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultRouter, opts...).ToFunc()
}

// ByDataScope orders the results by the data_scope field.
func ByDataScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataScope, opts...).ToFunc()
}

// ByMenuCheckStrictly orders the results by the menu_check_strictly field.
func ByMenuCheckStrictly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuCheckStrictly, opts...).ToFunc()
}

// ByDeptCheckStrictly orders the results by the dept_check_strictly field.
func ByDeptCheckStrictly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeptCheckStrictly, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
