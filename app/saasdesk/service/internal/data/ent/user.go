// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/beiduoke/go-scaffold/app/saasdesk/service/internal/data/ent/user"
)

// 用户表
type User struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// 创建时间
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 用户名
	Username *string `json:"username,omitempty"`
	// 密码
	Password *string `json:"password,omitempty"`
	// 昵称
	Nickname *string `json:"nickname,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 电子邮箱
	Email *string `json:"email,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`
	// 个人说明
	Description *string `json:"description,omitempty"`
	// 授权
	Authority    *user.Authority `json:"authority,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldPassword, user.FieldNickname, user.FieldPhone, user.FieldEmail, user.FieldAvatar, user.FieldDescription, user.FieldAuthority:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint64(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = new(time.Time)
				*u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = new(time.Time)
				*u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = new(string)
				*u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = new(string)
				*u.Password = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = new(string)
				*u.Nickname = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = new(string)
				*u.Phone = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = new(string)
				*u.Email = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = new(string)
				*u.Avatar = value.String
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = new(string)
				*u.Description = value.String
			}
		case user.FieldAuthority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authority", values[i])
			} else if value.Valid {
				u.Authority = new(user.Authority)
				*u.Authority = user.Authority(value.String)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	if v := u.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.Username; v != nil {
		builder.WriteString("username=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Nickname; v != nil {
		builder.WriteString("nickname=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Authority; v != nil {
		builder.WriteString("authority=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User