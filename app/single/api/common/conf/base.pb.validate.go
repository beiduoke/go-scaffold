// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/base.proto

package conf

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

// Validate checks the field values on Base with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Base) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Base with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BaseMultiError, or nil if none found.
func (m *Base) ValidateAll() error {
	return m.validate(true)
}

func (m *Base) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAdmin()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BaseValidationError{
					field:  "Admin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BaseValidationError{
					field:  "Admin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAdmin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BaseValidationError{
				field:  "Admin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDomain()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BaseValidationError{
					field:  "Domain",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BaseValidationError{
					field:  "Domain",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BaseValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BaseMultiError(errors)
	}

	return nil
}

// BaseMultiError is an error wrapping multiple validation errors returned by
// Base.ValidateAll() if the designated constraints aren't met.
type BaseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BaseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BaseMultiError) AllErrors() []error { return m }

// BaseValidationError is the validation error returned by Base.Validate if the
// designated constraints aren't met.
type BaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BaseValidationError) ErrorName() string { return "BaseValidationError" }

// Error satisfies the builtin error interface
func (e BaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BaseValidationError{}

// Validate checks the field values on Base_Admin with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Base_Admin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Base_Admin with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Base_AdminMultiError, or
// nil if none found.
func (m *Base_Admin) ValidateAll() error {
	return m.validate(true)
}

func (m *Base_Admin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DomainId

	// no validation rules for UserId

	if len(errors) > 0 {
		return Base_AdminMultiError(errors)
	}

	return nil
}

// Base_AdminMultiError is an error wrapping multiple validation errors
// returned by Base_Admin.ValidateAll() if the designated constraints aren't met.
type Base_AdminMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Base_AdminMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Base_AdminMultiError) AllErrors() []error { return m }

// Base_AdminValidationError is the validation error returned by
// Base_Admin.Validate if the designated constraints aren't met.
type Base_AdminValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Base_AdminValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Base_AdminValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Base_AdminValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Base_AdminValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Base_AdminValidationError) ErrorName() string { return "Base_AdminValidationError" }

// Error satisfies the builtin error interface
func (e Base_AdminValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBase_Admin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Base_AdminValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Base_AdminValidationError{}

// Validate checks the field values on Base_Domain with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Base_Domain) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Base_Domain with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Base_DomainMultiError, or
// nil if none found.
func (m *Base_Domain) ValidateAll() error {
	return m.validate(true)
}

func (m *Base_Domain) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Base_DomainMultiError(errors)
	}

	return nil
}

// Base_DomainMultiError is an error wrapping multiple validation errors
// returned by Base_Domain.ValidateAll() if the designated constraints aren't met.
type Base_DomainMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Base_DomainMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Base_DomainMultiError) AllErrors() []error { return m }

// Base_DomainValidationError is the validation error returned by
// Base_Domain.Validate if the designated constraints aren't met.
type Base_DomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Base_DomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Base_DomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Base_DomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Base_DomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Base_DomainValidationError) ErrorName() string { return "Base_DomainValidationError" }

// Error satisfies the builtin error interface
func (e Base_DomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBase_Domain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Base_DomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Base_DomainValidationError{}
