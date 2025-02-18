// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/logger.proto

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

// Validate checks the field values on Logger with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LoggerMultiError, or nil if none found.
func (m *Logger) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetZap()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Zap",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Zap",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetZap()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggerValidationError{
				field:  "Zap",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLogrus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Logrus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Logrus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLogrus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggerValidationError{
				field:  "Logrus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFluent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Fluent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Fluent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFluent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggerValidationError{
				field:  "Fluent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAliyun()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Aliyun",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Aliyun",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAliyun()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggerValidationError{
				field:  "Aliyun",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTencent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Tencent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoggerValidationError{
					field:  "Tencent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTencent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggerValidationError{
				field:  "Tencent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoggerMultiError(errors)
	}

	return nil
}

// LoggerMultiError is an error wrapping multiple validation errors returned by
// Logger.ValidateAll() if the designated constraints aren't met.
type LoggerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoggerMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoggerMultiError) AllErrors() []error { return m }

// LoggerValidationError is the validation error returned by Logger.Validate if
// the designated constraints aren't met.
type LoggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoggerValidationError) ErrorName() string { return "LoggerValidationError" }

// Error satisfies the builtin error interface
func (e LoggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoggerValidationError{}

// Validate checks the field values on Logger_Zap with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger_Zap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger_Zap with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Logger_ZapMultiError, or
// nil if none found.
func (m *Logger_Zap) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger_Zap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Filename

	// no validation rules for Level

	// no validation rules for MaxSize

	// no validation rules for MaxAge

	// no validation rules for MaxBackups

	if len(errors) > 0 {
		return Logger_ZapMultiError(errors)
	}

	return nil
}

// Logger_ZapMultiError is an error wrapping multiple validation errors
// returned by Logger_Zap.ValidateAll() if the designated constraints aren't met.
type Logger_ZapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Logger_ZapMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Logger_ZapMultiError) AllErrors() []error { return m }

// Logger_ZapValidationError is the validation error returned by
// Logger_Zap.Validate if the designated constraints aren't met.
type Logger_ZapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Logger_ZapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Logger_ZapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Logger_ZapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Logger_ZapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Logger_ZapValidationError) ErrorName() string { return "Logger_ZapValidationError" }

// Error satisfies the builtin error interface
func (e Logger_ZapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger_Zap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Logger_ZapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Logger_ZapValidationError{}

// Validate checks the field values on Logger_Logrus with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger_Logrus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger_Logrus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Logger_LogrusMultiError, or
// nil if none found.
func (m *Logger_Logrus) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger_Logrus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Level

	// no validation rules for Formatter

	// no validation rules for TimestampFormat

	// no validation rules for DisableColors

	// no validation rules for DisableTimestamp

	if len(errors) > 0 {
		return Logger_LogrusMultiError(errors)
	}

	return nil
}

// Logger_LogrusMultiError is an error wrapping multiple validation errors
// returned by Logger_Logrus.ValidateAll() if the designated constraints
// aren't met.
type Logger_LogrusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Logger_LogrusMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Logger_LogrusMultiError) AllErrors() []error { return m }

// Logger_LogrusValidationError is the validation error returned by
// Logger_Logrus.Validate if the designated constraints aren't met.
type Logger_LogrusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Logger_LogrusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Logger_LogrusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Logger_LogrusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Logger_LogrusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Logger_LogrusValidationError) ErrorName() string { return "Logger_LogrusValidationError" }

// Error satisfies the builtin error interface
func (e Logger_LogrusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger_Logrus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Logger_LogrusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Logger_LogrusValidationError{}

// Validate checks the field values on Logger_Fluent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger_Fluent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger_Fluent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Logger_FluentMultiError, or
// nil if none found.
func (m *Logger_Fluent) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger_Fluent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Endpoint

	if len(errors) > 0 {
		return Logger_FluentMultiError(errors)
	}

	return nil
}

// Logger_FluentMultiError is an error wrapping multiple validation errors
// returned by Logger_Fluent.ValidateAll() if the designated constraints
// aren't met.
type Logger_FluentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Logger_FluentMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Logger_FluentMultiError) AllErrors() []error { return m }

// Logger_FluentValidationError is the validation error returned by
// Logger_Fluent.Validate if the designated constraints aren't met.
type Logger_FluentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Logger_FluentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Logger_FluentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Logger_FluentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Logger_FluentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Logger_FluentValidationError) ErrorName() string { return "Logger_FluentValidationError" }

// Error satisfies the builtin error interface
func (e Logger_FluentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger_Fluent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Logger_FluentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Logger_FluentValidationError{}

// Validate checks the field values on Logger_Aliyun with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger_Aliyun) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger_Aliyun with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Logger_AliyunMultiError, or
// nil if none found.
func (m *Logger_Aliyun) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger_Aliyun) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Endpoint

	// no validation rules for Project

	// no validation rules for AccessKey

	// no validation rules for AccessSecret

	if len(errors) > 0 {
		return Logger_AliyunMultiError(errors)
	}

	return nil
}

// Logger_AliyunMultiError is an error wrapping multiple validation errors
// returned by Logger_Aliyun.ValidateAll() if the designated constraints
// aren't met.
type Logger_AliyunMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Logger_AliyunMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Logger_AliyunMultiError) AllErrors() []error { return m }

// Logger_AliyunValidationError is the validation error returned by
// Logger_Aliyun.Validate if the designated constraints aren't met.
type Logger_AliyunValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Logger_AliyunValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Logger_AliyunValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Logger_AliyunValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Logger_AliyunValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Logger_AliyunValidationError) ErrorName() string { return "Logger_AliyunValidationError" }

// Error satisfies the builtin error interface
func (e Logger_AliyunValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger_Aliyun.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Logger_AliyunValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Logger_AliyunValidationError{}

// Validate checks the field values on Logger_Tencent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Logger_Tencent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Logger_Tencent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Logger_TencentMultiError,
// or nil if none found.
func (m *Logger_Tencent) ValidateAll() error {
	return m.validate(true)
}

func (m *Logger_Tencent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Endpoint

	// no validation rules for TopicId

	// no validation rules for AccessKey

	// no validation rules for AccessSecret

	if len(errors) > 0 {
		return Logger_TencentMultiError(errors)
	}

	return nil
}

// Logger_TencentMultiError is an error wrapping multiple validation errors
// returned by Logger_Tencent.ValidateAll() if the designated constraints
// aren't met.
type Logger_TencentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Logger_TencentMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Logger_TencentMultiError) AllErrors() []error { return m }

// Logger_TencentValidationError is the validation error returned by
// Logger_Tencent.Validate if the designated constraints aren't met.
type Logger_TencentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Logger_TencentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Logger_TencentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Logger_TencentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Logger_TencentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Logger_TencentValidationError) ErrorName() string { return "Logger_TencentValidationError" }

// Error satisfies the builtin error interface
func (e Logger_TencentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger_Tencent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Logger_TencentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Logger_TencentValidationError{}
