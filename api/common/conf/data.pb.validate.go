// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/data.proto

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

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDatabase()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Database",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Database",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDatabase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Database",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRedis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRedis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Redis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMongodb()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Mongodb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Mongodb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMongodb()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Mongodb",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetElasticSearch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "ElasticSearch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "ElasticSearch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetElasticSearch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "ElasticSearch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClickhouse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Clickhouse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Clickhouse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClickhouse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Clickhouse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetInfluxdb()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Influxdb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Influxdb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfluxdb()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Influxdb",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDoris()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Doris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Doris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDoris()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Doris",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKafka()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Kafka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Kafka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKafka()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Kafka",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMeilisearch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Meilisearch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Meilisearch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMeilisearch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Meilisearch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on Data_Database with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Database) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Database with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_DatabaseMultiError, or
// nil if none found.
func (m *Data_Database) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Database) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Driver

	// no validation rules for Source

	// no validation rules for Migrate

	// no validation rules for Debug

	// no validation rules for MaxIdleConnections

	// no validation rules for MaxOpenConnections

	if all {
		switch v := interface{}(m.GetConnectionMaxLifetime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_DatabaseValidationError{
					field:  "ConnectionMaxLifetime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_DatabaseValidationError{
					field:  "ConnectionMaxLifetime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConnectionMaxLifetime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_DatabaseValidationError{
				field:  "ConnectionMaxLifetime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Data_DatabaseMultiError(errors)
	}

	return nil
}

// Data_DatabaseMultiError is an error wrapping multiple validation errors
// returned by Data_Database.ValidateAll() if the designated constraints
// aren't met.
type Data_DatabaseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_DatabaseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_DatabaseMultiError) AllErrors() []error { return m }

// Data_DatabaseValidationError is the validation error returned by
// Data_Database.Validate if the designated constraints aren't met.
type Data_DatabaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_DatabaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_DatabaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_DatabaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_DatabaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_DatabaseValidationError) ErrorName() string { return "Data_DatabaseValidationError" }

// Error satisfies the builtin error interface
func (e Data_DatabaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Database.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_DatabaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_DatabaseValidationError{}

// Validate checks the field values on Data_Redis with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Redis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Redis with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_RedisMultiError, or
// nil if none found.
func (m *Data_Redis) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Redis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Addr

	// no validation rules for Password

	// no validation rules for Db

	if all {
		switch v := interface{}(m.GetDialTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDialTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "DialTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReadTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "ReadTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_RedisValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_RedisValidationError{
				field:  "WriteTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EnableTracing

	// no validation rules for EnableMetrics

	if len(errors) > 0 {
		return Data_RedisMultiError(errors)
	}

	return nil
}

// Data_RedisMultiError is an error wrapping multiple validation errors
// returned by Data_Redis.ValidateAll() if the designated constraints aren't met.
type Data_RedisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_RedisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_RedisMultiError) AllErrors() []error { return m }

// Data_RedisValidationError is the validation error returned by
// Data_Redis.Validate if the designated constraints aren't met.
type Data_RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_RedisValidationError) ErrorName() string { return "Data_RedisValidationError" }

// Error satisfies the builtin error interface
func (e Data_RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Redis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_RedisValidationError{}

// Validate checks the field values on Data_MongoDB with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_MongoDB) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_MongoDB with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_MongoDBMultiError, or
// nil if none found.
func (m *Data_MongoDB) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_MongoDB) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return Data_MongoDBMultiError(errors)
	}

	return nil
}

// Data_MongoDBMultiError is an error wrapping multiple validation errors
// returned by Data_MongoDB.ValidateAll() if the designated constraints aren't met.
type Data_MongoDBMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_MongoDBMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_MongoDBMultiError) AllErrors() []error { return m }

// Data_MongoDBValidationError is the validation error returned by
// Data_MongoDB.Validate if the designated constraints aren't met.
type Data_MongoDBValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_MongoDBValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_MongoDBValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_MongoDBValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_MongoDBValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_MongoDBValidationError) ErrorName() string { return "Data_MongoDBValidationError" }

// Error satisfies the builtin error interface
func (e Data_MongoDBValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_MongoDB.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_MongoDBValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_MongoDBValidationError{}

// Validate checks the field values on Data_ClickHouse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Data_ClickHouse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_ClickHouse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Data_ClickHouseMultiError, or nil if none found.
func (m *Data_ClickHouse) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_ClickHouse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return Data_ClickHouseMultiError(errors)
	}

	return nil
}

// Data_ClickHouseMultiError is an error wrapping multiple validation errors
// returned by Data_ClickHouse.ValidateAll() if the designated constraints
// aren't met.
type Data_ClickHouseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_ClickHouseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_ClickHouseMultiError) AllErrors() []error { return m }

// Data_ClickHouseValidationError is the validation error returned by
// Data_ClickHouse.Validate if the designated constraints aren't met.
type Data_ClickHouseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_ClickHouseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_ClickHouseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_ClickHouseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_ClickHouseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_ClickHouseValidationError) ErrorName() string { return "Data_ClickHouseValidationError" }

// Error satisfies the builtin error interface
func (e Data_ClickHouseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_ClickHouse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_ClickHouseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_ClickHouseValidationError{}

// Validate checks the field values on Data_InfluxDB with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_InfluxDB) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_InfluxDB with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_InfluxDBMultiError, or
// nil if none found.
func (m *Data_InfluxDB) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_InfluxDB) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Token

	// no validation rules for Orgnization

	// no validation rules for Bucket

	if len(errors) > 0 {
		return Data_InfluxDBMultiError(errors)
	}

	return nil
}

// Data_InfluxDBMultiError is an error wrapping multiple validation errors
// returned by Data_InfluxDB.ValidateAll() if the designated constraints
// aren't met.
type Data_InfluxDBMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_InfluxDBMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_InfluxDBMultiError) AllErrors() []error { return m }

// Data_InfluxDBValidationError is the validation error returned by
// Data_InfluxDB.Validate if the designated constraints aren't met.
type Data_InfluxDBValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_InfluxDBValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_InfluxDBValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_InfluxDBValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_InfluxDBValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_InfluxDBValidationError) ErrorName() string { return "Data_InfluxDBValidationError" }

// Error satisfies the builtin error interface
func (e Data_InfluxDBValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_InfluxDB.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_InfluxDBValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_InfluxDBValidationError{}

// Validate checks the field values on Data_Kafka with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Kafka) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Kafka with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_KafkaMultiError, or
// nil if none found.
func (m *Data_Kafka) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Kafka) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Codec

	if len(errors) > 0 {
		return Data_KafkaMultiError(errors)
	}

	return nil
}

// Data_KafkaMultiError is an error wrapping multiple validation errors
// returned by Data_Kafka.ValidateAll() if the designated constraints aren't met.
type Data_KafkaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_KafkaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_KafkaMultiError) AllErrors() []error { return m }

// Data_KafkaValidationError is the validation error returned by
// Data_Kafka.Validate if the designated constraints aren't met.
type Data_KafkaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_KafkaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_KafkaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_KafkaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_KafkaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_KafkaValidationError) ErrorName() string { return "Data_KafkaValidationError" }

// Error satisfies the builtin error interface
func (e Data_KafkaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Kafka.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_KafkaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_KafkaValidationError{}

// Validate checks the field values on Data_Doris with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Doris) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Doris with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_DorisMultiError, or
// nil if none found.
func (m *Data_Doris) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Doris) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return Data_DorisMultiError(errors)
	}

	return nil
}

// Data_DorisMultiError is an error wrapping multiple validation errors
// returned by Data_Doris.ValidateAll() if the designated constraints aren't met.
type Data_DorisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_DorisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_DorisMultiError) AllErrors() []error { return m }

// Data_DorisValidationError is the validation error returned by
// Data_Doris.Validate if the designated constraints aren't met.
type Data_DorisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_DorisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_DorisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_DorisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_DorisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_DorisValidationError) ErrorName() string { return "Data_DorisValidationError" }

// Error satisfies the builtin error interface
func (e Data_DorisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Doris.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_DorisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_DorisValidationError{}

// Validate checks the field values on Data_ElasticSearch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Data_ElasticSearch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_ElasticSearch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Data_ElasticSearchMultiError, or nil if none found.
func (m *Data_ElasticSearch) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_ElasticSearch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return Data_ElasticSearchMultiError(errors)
	}

	return nil
}

// Data_ElasticSearchMultiError is an error wrapping multiple validation errors
// returned by Data_ElasticSearch.ValidateAll() if the designated constraints
// aren't met.
type Data_ElasticSearchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_ElasticSearchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_ElasticSearchMultiError) AllErrors() []error { return m }

// Data_ElasticSearchValidationError is the validation error returned by
// Data_ElasticSearch.Validate if the designated constraints aren't met.
type Data_ElasticSearchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_ElasticSearchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_ElasticSearchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_ElasticSearchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_ElasticSearchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_ElasticSearchValidationError) ErrorName() string {
	return "Data_ElasticSearchValidationError"
}

// Error satisfies the builtin error interface
func (e Data_ElasticSearchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_ElasticSearch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_ElasticSearchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_ElasticSearchValidationError{}

// Validate checks the field values on Data_Meilisearch with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Data_Meilisearch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Meilisearch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Data_MeilisearchMultiError, or nil if none found.
func (m *Data_Meilisearch) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Meilisearch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Host

	// no validation rules for ApiKey

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_MeilisearchValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_MeilisearchValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_MeilisearchValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Data_MeilisearchMultiError(errors)
	}

	return nil
}

// Data_MeilisearchMultiError is an error wrapping multiple validation errors
// returned by Data_Meilisearch.ValidateAll() if the designated constraints
// aren't met.
type Data_MeilisearchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_MeilisearchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_MeilisearchMultiError) AllErrors() []error { return m }

// Data_MeilisearchValidationError is the validation error returned by
// Data_Meilisearch.Validate if the designated constraints aren't met.
type Data_MeilisearchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_MeilisearchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_MeilisearchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_MeilisearchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_MeilisearchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_MeilisearchValidationError) ErrorName() string { return "Data_MeilisearchValidationError" }

// Error satisfies the builtin error interface
func (e Data_MeilisearchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Meilisearch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_MeilisearchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_MeilisearchValidationError{}
