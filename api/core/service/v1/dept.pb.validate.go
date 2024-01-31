// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/service/v1/dept.proto

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

// Validate checks the field values on Dept with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Dept) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dept with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DeptMultiError, or nil if none found.
func (m *Dept) ValidateAll() error {
	return m.validate(true)
}

func (m *Dept) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if m.CreatedAt != nil {
		// no validation rules for CreatedAt
	}

	if m.UpdatedAt != nil {
		// no validation rules for UpdatedAt
	}

	if m.Sort != nil {
		// no validation rules for Sort
	}

	if m.State != nil {
		// no validation rules for State
	}

	if m.Remarks != nil {
		// no validation rules for Remarks
	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.LeaderId != nil {
		// no validation rules for LeaderId
	}

	if len(errors) > 0 {
		return DeptMultiError(errors)
	}

	return nil
}

// DeptMultiError is an error wrapping multiple validation errors returned by
// Dept.ValidateAll() if the designated constraints aren't met.
type DeptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeptMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeptMultiError) AllErrors() []error { return m }

// DeptValidationError is the validation error returned by Dept.Validate if the
// designated constraints aren't met.
type DeptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeptValidationError) ErrorName() string { return "DeptValidationError" }

// Error satisfies the builtin error interface
func (e DeptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDept.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeptValidationError{}

// Validate checks the field values on CreateDeptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateDeptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDeptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDeptRequestMultiError, or nil if none found.
func (m *CreateDeptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDeptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDept()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateDeptRequestValidationError{
					field:  "Dept",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateDeptRequestValidationError{
					field:  "Dept",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDept()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateDeptRequestValidationError{
				field:  "Dept",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return CreateDeptRequestMultiError(errors)
	}

	return nil
}

// CreateDeptRequestMultiError is an error wrapping multiple validation errors
// returned by CreateDeptRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateDeptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDeptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDeptRequestMultiError) AllErrors() []error { return m }

// CreateDeptRequestValidationError is the validation error returned by
// CreateDeptRequest.Validate if the designated constraints aren't met.
type CreateDeptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDeptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDeptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDeptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDeptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDeptRequestValidationError) ErrorName() string {
	return "CreateDeptRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDeptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDeptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDeptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDeptRequestValidationError{}

// Validate checks the field values on CreateDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateDeptResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDeptResponseMultiError, or nil if none found.
func (m *CreateDeptResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDeptResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateDeptResponseMultiError(errors)
	}

	return nil
}

// CreateDeptResponseMultiError is an error wrapping multiple validation errors
// returned by CreateDeptResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateDeptResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDeptResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDeptResponseMultiError) AllErrors() []error { return m }

// CreateDeptResponseValidationError is the validation error returned by
// CreateDeptResponse.Validate if the designated constraints aren't met.
type CreateDeptResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDeptResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDeptResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDeptResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDeptResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDeptResponseValidationError) ErrorName() string {
	return "CreateDeptResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDeptResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDeptResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDeptResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDeptResponseValidationError{}

// Validate checks the field values on UpdateDeptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateDeptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDeptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDeptRequestMultiError, or nil if none found.
func (m *UpdateDeptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDeptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetDept()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateDeptRequestValidationError{
					field:  "Dept",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateDeptRequestValidationError{
					field:  "Dept",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDept()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateDeptRequestValidationError{
				field:  "Dept",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return UpdateDeptRequestMultiError(errors)
	}

	return nil
}

// UpdateDeptRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateDeptRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateDeptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDeptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDeptRequestMultiError) AllErrors() []error { return m }

// UpdateDeptRequestValidationError is the validation error returned by
// UpdateDeptRequest.Validate if the designated constraints aren't met.
type UpdateDeptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDeptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDeptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDeptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDeptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDeptRequestValidationError) ErrorName() string {
	return "UpdateDeptRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDeptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDeptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDeptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDeptRequestValidationError{}

// Validate checks the field values on UpdateDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateDeptResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDeptResponseMultiError, or nil if none found.
func (m *UpdateDeptResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDeptResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateDeptResponseMultiError(errors)
	}

	return nil
}

// UpdateDeptResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateDeptResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateDeptResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDeptResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDeptResponseMultiError) AllErrors() []error { return m }

// UpdateDeptResponseValidationError is the validation error returned by
// UpdateDeptResponse.Validate if the designated constraints aren't met.
type UpdateDeptResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDeptResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDeptResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDeptResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDeptResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDeptResponseValidationError) ErrorName() string {
	return "UpdateDeptResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDeptResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDeptResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDeptResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDeptResponseValidationError{}

// Validate checks the field values on DeleteDeptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteDeptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDeptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDeptRequestMultiError, or nil if none found.
func (m *DeleteDeptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDeptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return DeleteDeptRequestMultiError(errors)
	}

	return nil
}

// DeleteDeptRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteDeptRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteDeptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDeptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDeptRequestMultiError) AllErrors() []error { return m }

// DeleteDeptRequestValidationError is the validation error returned by
// DeleteDeptRequest.Validate if the designated constraints aren't met.
type DeleteDeptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDeptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDeptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDeptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDeptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDeptRequestValidationError) ErrorName() string {
	return "DeleteDeptRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDeptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDeptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDeptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDeptRequestValidationError{}

// Validate checks the field values on DeleteDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteDeptResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDeptResponseMultiError, or nil if none found.
func (m *DeleteDeptResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDeptResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteDeptResponseMultiError(errors)
	}

	return nil
}

// DeleteDeptResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteDeptResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteDeptResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDeptResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDeptResponseMultiError) AllErrors() []error { return m }

// DeleteDeptResponseValidationError is the validation error returned by
// DeleteDeptResponse.Validate if the designated constraints aren't met.
type DeleteDeptResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDeptResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDeptResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDeptResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDeptResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDeptResponseValidationError) ErrorName() string {
	return "DeleteDeptResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDeptResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDeptResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDeptResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDeptResponseValidationError{}

// Validate checks the field values on GetDeptRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetDeptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetDeptRequestMultiError,
// or nil if none found.
func (m *GetDeptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetDeptRequestMultiError(errors)
	}

	return nil
}

// GetDeptRequestMultiError is an error wrapping multiple validation errors
// returned by GetDeptRequest.ValidateAll() if the designated constraints
// aren't met.
type GetDeptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeptRequestMultiError) AllErrors() []error { return m }

// GetDeptRequestValidationError is the validation error returned by
// GetDeptRequest.Validate if the designated constraints aren't met.
type GetDeptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeptRequestValidationError) ErrorName() string { return "GetDeptRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetDeptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeptRequestValidationError{}

// Validate checks the field values on GetDeptResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetDeptResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeptResponseMultiError, or nil if none found.
func (m *GetDeptResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeptResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetDeptResponseMultiError(errors)
	}

	return nil
}

// GetDeptResponseMultiError is an error wrapping multiple validation errors
// returned by GetDeptResponse.ValidateAll() if the designated constraints
// aren't met.
type GetDeptResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeptResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeptResponseMultiError) AllErrors() []error { return m }

// GetDeptResponseValidationError is the validation error returned by
// GetDeptResponse.Validate if the designated constraints aren't met.
type GetDeptResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeptResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeptResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeptResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeptResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeptResponseValidationError) ErrorName() string { return "GetDeptResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetDeptResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeptResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeptResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeptResponseValidationError{}

// Validate checks the field values on ListDeptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListDeptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDeptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDeptRequestMultiError, or nil if none found.
func (m *ListDeptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDeptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListDeptRequestMultiError(errors)
	}

	return nil
}

// ListDeptRequestMultiError is an error wrapping multiple validation errors
// returned by ListDeptRequest.ValidateAll() if the designated constraints
// aren't met.
type ListDeptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDeptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDeptRequestMultiError) AllErrors() []error { return m }

// ListDeptRequestValidationError is the validation error returned by
// ListDeptRequest.Validate if the designated constraints aren't met.
type ListDeptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDeptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDeptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDeptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDeptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDeptRequestValidationError) ErrorName() string { return "ListDeptRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListDeptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDeptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDeptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDeptRequestValidationError{}

// Validate checks the field values on ListDeptResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListDeptResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDeptResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDeptResponseMultiError, or nil if none found.
func (m *ListDeptResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDeptResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListDeptResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListDeptResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDeptResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListDeptResponseMultiError(errors)
	}

	return nil
}

// ListDeptResponseMultiError is an error wrapping multiple validation errors
// returned by ListDeptResponse.ValidateAll() if the designated constraints
// aren't met.
type ListDeptResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDeptResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDeptResponseMultiError) AllErrors() []error { return m }

// ListDeptResponseValidationError is the validation error returned by
// ListDeptResponse.Validate if the designated constraints aren't met.
type ListDeptResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDeptResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDeptResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDeptResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDeptResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDeptResponseValidationError) ErrorName() string { return "ListDeptResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListDeptResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDeptResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDeptResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDeptResponseValidationError{}
