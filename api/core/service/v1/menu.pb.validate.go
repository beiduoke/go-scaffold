// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/service/v1/menu.proto

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

// Validate checks the field values on Menu with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Menu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Menu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MenuMultiError, or nil if none found.
func (m *Menu) ValidateAll() error {
	return m.validate(true)
}

func (m *Menu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.CreatedAt != nil {
		// no validation rules for CreatedAt
	}

	if m.UpdatedAt != nil {
		// no validation rules for UpdatedAt
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Sort != nil {
		// no validation rules for Sort
	}

	if m.State != nil {
		// no validation rules for State
	}

	if m.Remark != nil {
		// no validation rules for Remark
	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if m.Component != nil {
		// no validation rules for Component
	}

	if m.Icon != nil {
		// no validation rules for Icon
	}

	if m.IsExt != nil {
		// no validation rules for IsExt
	}

	if m.ExtUrl != nil {
		// no validation rules for ExtUrl
	}

	if m.Permissions != nil {
		// no validation rules for Permissions
	}

	if m.Redirect != nil {
		// no validation rules for Redirect
	}

	if m.CurrentActiveMenu != nil {
		// no validation rules for CurrentActiveMenu
	}

	if m.KeepAlive != nil {
		// no validation rules for KeepAlive
	}

	if m.Visible != nil {
		// no validation rules for Visible
	}

	if m.HideTab != nil {
		// no validation rules for HideTab
	}

	if m.HideMenu != nil {
		// no validation rules for HideMenu
	}

	if m.HideBreadcrumb != nil {
		// no validation rules for HideBreadcrumb
	}

	if len(errors) > 0 {
		return MenuMultiError(errors)
	}

	return nil
}

// MenuMultiError is an error wrapping multiple validation errors returned by
// Menu.ValidateAll() if the designated constraints aren't met.
type MenuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuMultiError) AllErrors() []error { return m }

// MenuValidationError is the validation error returned by Menu.Validate if the
// designated constraints aren't met.
type MenuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuValidationError) ErrorName() string { return "MenuValidationError" }

// Error satisfies the builtin error interface
func (e MenuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuValidationError{}

// Validate checks the field values on CreateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuRequestMultiError, or nil if none found.
func (m *CreateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMenu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMenu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMenuRequestValidationError{
				field:  "Menu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return CreateMenuRequestMultiError(errors)
	}

	return nil
}

// CreateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by CreateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuRequestMultiError) AllErrors() []error { return m }

// CreateMenuRequestValidationError is the validation error returned by
// CreateMenuRequest.Validate if the designated constraints aren't met.
type CreateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuRequestValidationError) ErrorName() string {
	return "CreateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuRequestValidationError{}

// Validate checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuResponseMultiError, or nil if none found.
func (m *CreateMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateMenuResponseMultiError(errors)
	}

	return nil
}

// CreateMenuResponseMultiError is an error wrapping multiple validation errors
// returned by CreateMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuResponseMultiError) AllErrors() []error { return m }

// CreateMenuResponseValidationError is the validation error returned by
// CreateMenuResponse.Validate if the designated constraints aren't met.
type CreateMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuResponseValidationError) ErrorName() string {
	return "CreateMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuResponseValidationError{}

// Validate checks the field values on UpdateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMenuRequestMultiError, or nil if none found.
func (m *UpdateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetMenu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateMenuRequestValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMenu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateMenuRequestValidationError{
				field:  "Menu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return UpdateMenuRequestMultiError(errors)
	}

	return nil
}

// UpdateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMenuRequestMultiError) AllErrors() []error { return m }

// UpdateMenuRequestValidationError is the validation error returned by
// UpdateMenuRequest.Validate if the designated constraints aren't met.
type UpdateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMenuRequestValidationError) ErrorName() string {
	return "UpdateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMenuRequestValidationError{}

// Validate checks the field values on UpdateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMenuResponseMultiError, or nil if none found.
func (m *UpdateMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateMenuResponseMultiError(errors)
	}

	return nil
}

// UpdateMenuResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMenuResponseMultiError) AllErrors() []error { return m }

// UpdateMenuResponseValidationError is the validation error returned by
// UpdateMenuResponse.Validate if the designated constraints aren't met.
type UpdateMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMenuResponseValidationError) ErrorName() string {
	return "UpdateMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMenuResponseValidationError{}

// Validate checks the field values on DeleteMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMenuRequestMultiError, or nil if none found.
func (m *DeleteMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return DeleteMenuRequestMultiError(errors)
	}

	return nil
}

// DeleteMenuRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMenuRequestMultiError) AllErrors() []error { return m }

// DeleteMenuRequestValidationError is the validation error returned by
// DeleteMenuRequest.Validate if the designated constraints aren't met.
type DeleteMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMenuRequestValidationError) ErrorName() string {
	return "DeleteMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMenuRequestValidationError{}

// Validate checks the field values on DeleteMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMenuResponseMultiError, or nil if none found.
func (m *DeleteMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteMenuResponseMultiError(errors)
	}

	return nil
}

// DeleteMenuResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMenuResponseMultiError) AllErrors() []error { return m }

// DeleteMenuResponseValidationError is the validation error returned by
// DeleteMenuResponse.Validate if the designated constraints aren't met.
type DeleteMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMenuResponseValidationError) ErrorName() string {
	return "DeleteMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMenuResponseValidationError{}

// Validate checks the field values on GetMenuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetMenuRequestMultiError,
// or nil if none found.
func (m *GetMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetMenuRequestMultiError(errors)
	}

	return nil
}

// GetMenuRequestMultiError is an error wrapping multiple validation errors
// returned by GetMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuRequestMultiError) AllErrors() []error { return m }

// GetMenuRequestValidationError is the validation error returned by
// GetMenuRequest.Validate if the designated constraints aren't met.
type GetMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuRequestValidationError) ErrorName() string { return "GetMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuRequestValidationError{}

// Validate checks the field values on GetMenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMenuResponseMultiError, or nil if none found.
func (m *GetMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetMenuResponseMultiError(errors)
	}

	return nil
}

// GetMenuResponseMultiError is an error wrapping multiple validation errors
// returned by GetMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type GetMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuResponseMultiError) AllErrors() []error { return m }

// GetMenuResponseValidationError is the validation error returned by
// GetMenuResponse.Validate if the designated constraints aren't met.
type GetMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuResponseValidationError) ErrorName() string { return "GetMenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuResponseValidationError{}

// Validate checks the field values on ListMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMenuRequestMultiError, or nil if none found.
func (m *ListMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListMenuRequestMultiError(errors)
	}

	return nil
}

// ListMenuRequestMultiError is an error wrapping multiple validation errors
// returned by ListMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type ListMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMenuRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMenuRequestMultiError) AllErrors() []error { return m }

// ListMenuRequestValidationError is the validation error returned by
// ListMenuRequest.Validate if the designated constraints aren't met.
type ListMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMenuRequestValidationError) ErrorName() string { return "ListMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMenuRequestValidationError{}

// Validate checks the field values on ListMenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMenuResponseMultiError, or nil if none found.
func (m *ListMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMenuResponse) validate(all bool) error {
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
					errors = append(errors, ListMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListMenuResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMenuResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListMenuResponseMultiError(errors)
	}

	return nil
}

// ListMenuResponseMultiError is an error wrapping multiple validation errors
// returned by ListMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type ListMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMenuResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMenuResponseMultiError) AllErrors() []error { return m }

// ListMenuResponseValidationError is the validation error returned by
// ListMenuResponse.Validate if the designated constraints aren't met.
type ListMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMenuResponseValidationError) ErrorName() string { return "ListMenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMenuResponseValidationError{}
