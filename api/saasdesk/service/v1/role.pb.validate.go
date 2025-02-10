// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: saasdesk/service/v1/role.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Role) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Role with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RoleMultiError, or nil if none found.
func (m *Role) ValidateAll() error {
	return m.validate(true)
}

func (m *Role) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoleValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UpdatedAt != nil {

		if all {
			switch v := interface{}(m.GetUpdatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoleValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DefaultRouter != nil {
		// no validation rules for DefaultRouter
	}

	if m.Sort != nil {
		// no validation rules for Sort
	}

	if m.DataScope != nil {
		// no validation rules for DataScope
	}

	if m.MenuCheckStrictly != nil {
		// no validation rules for MenuCheckStrictly
	}

	if m.DeptCheckStrictly != nil {
		// no validation rules for DeptCheckStrictly
	}

	if m.State != nil {
		// no validation rules for State
	}

	if m.Remark != nil {
		// no validation rules for Remark
	}

	if len(errors) > 0 {
		return RoleMultiError(errors)
	}

	return nil
}

// RoleMultiError is an error wrapping multiple validation errors returned by
// Role.ValidateAll() if the designated constraints aren't met.
type RoleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleMultiError) AllErrors() []error { return m }

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on CreateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleRequestMultiError, or nil if none found.
func (m *CreateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateRoleRequestMultiError(errors)
	}

	return nil
}

// CreateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleRequestMultiError) AllErrors() []error { return m }

// CreateRoleRequestValidationError is the validation error returned by
// CreateRoleRequest.Validate if the designated constraints aren't met.
type CreateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleRequestValidationError) ErrorName() string {
	return "CreateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleRequestValidationError{}

// Validate checks the field values on CreateRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleResponseMultiError, or nil if none found.
func (m *CreateRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateRoleResponseMultiError(errors)
	}

	return nil
}

// CreateRoleResponseMultiError is an error wrapping multiple validation errors
// returned by CreateRoleResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleResponseMultiError) AllErrors() []error { return m }

// CreateRoleResponseValidationError is the validation error returned by
// CreateRoleResponse.Validate if the designated constraints aren't met.
type CreateRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleResponseValidationError) ErrorName() string {
	return "CreateRoleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleResponseValidationError{}

// Validate checks the field values on UpdateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleRequestMultiError, or nil if none found.
func (m *UpdateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRoleRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleRequestMultiError) AllErrors() []error { return m }

// UpdateRoleRequestValidationError is the validation error returned by
// UpdateRoleRequest.Validate if the designated constraints aren't met.
type UpdateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleRequestValidationError) ErrorName() string {
	return "UpdateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleRequestValidationError{}

// Validate checks the field values on UpdateRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleResponseMultiError, or nil if none found.
func (m *UpdateRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRoleResponseMultiError(errors)
	}

	return nil
}

// UpdateRoleResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleResponseMultiError) AllErrors() []error { return m }

// UpdateRoleResponseValidationError is the validation error returned by
// UpdateRoleResponse.Validate if the designated constraints aren't met.
type UpdateRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleResponseValidationError) ErrorName() string {
	return "UpdateRoleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleResponseValidationError{}

// Validate checks the field values on DeleteRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleRequestMultiError, or nil if none found.
func (m *DeleteRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRoleRequestMultiError(errors)
	}

	return nil
}

// DeleteRoleRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleRequestMultiError) AllErrors() []error { return m }

// DeleteRoleRequestValidationError is the validation error returned by
// DeleteRoleRequest.Validate if the designated constraints aren't met.
type DeleteRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleRequestValidationError) ErrorName() string {
	return "DeleteRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleRequestValidationError{}

// Validate checks the field values on DeleteRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleResponseMultiError, or nil if none found.
func (m *DeleteRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRoleResponseMultiError(errors)
	}

	return nil
}

// DeleteRoleResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteRoleResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleResponseMultiError) AllErrors() []error { return m }

// DeleteRoleResponseValidationError is the validation error returned by
// DeleteRoleResponse.Validate if the designated constraints aren't met.
type DeleteRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleResponseValidationError) ErrorName() string {
	return "DeleteRoleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleResponseValidationError{}

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRoleRequestMultiError,
// or nil if none found.
func (m *GetRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRoleRequestMultiError(errors)
	}

	return nil
}

// GetRoleRequestMultiError is an error wrapping multiple validation errors
// returned by GetRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleRequestMultiError) AllErrors() []error { return m }

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on GetRoleResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRoleResponseMultiError, or nil if none found.
func (m *GetRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRoleResponseMultiError(errors)
	}

	return nil
}

// GetRoleResponseMultiError is an error wrapping multiple validation errors
// returned by GetRoleResponse.ValidateAll() if the designated constraints
// aren't met.
type GetRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleResponseMultiError) AllErrors() []error { return m }

// GetRoleResponseValidationError is the validation error returned by
// GetRoleResponse.Validate if the designated constraints aren't met.
type GetRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleResponseValidationError) ErrorName() string { return "GetRoleResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleResponseValidationError{}

// Validate checks the field values on ListRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoleRequestMultiError, or nil if none found.
func (m *ListRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRoleRequestMultiError(errors)
	}

	return nil
}

// ListRoleRequestMultiError is an error wrapping multiple validation errors
// returned by ListRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type ListRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoleRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoleRequestMultiError) AllErrors() []error { return m }

// ListRoleRequestValidationError is the validation error returned by
// ListRoleRequest.Validate if the designated constraints aren't met.
type ListRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleRequestValidationError) ErrorName() string { return "ListRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleRequestValidationError{}

// Validate checks the field values on ListRoleResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoleResponseMultiError, or nil if none found.
func (m *ListRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRoleResponseMultiError(errors)
	}

	return nil
}

// ListRoleResponseMultiError is an error wrapping multiple validation errors
// returned by ListRoleResponse.ValidateAll() if the designated constraints
// aren't met.
type ListRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoleResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoleResponseMultiError) AllErrors() []error { return m }

// ListRoleResponseValidationError is the validation error returned by
// ListRoleResponse.Validate if the designated constraints aren't met.
type ListRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleResponseValidationError) ErrorName() string { return "ListRoleResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleResponseValidationError{}
