// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/server.proto

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

// Validate checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServerMultiError, or nil if none found.
func (m *Server) ValidateAll() error {
	return m.validate(true)
}

func (m *Server) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHttp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGrpc()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Grpc",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Grpc",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpc()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Grpc",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWebsocket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Websocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Websocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWebsocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Websocket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMqtt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Mqtt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Mqtt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMqtt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Mqtt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKafka()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Kafka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Kafka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKafka()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Kafka",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRabbitmq()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Rabbitmq",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Rabbitmq",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRabbitmq()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Rabbitmq",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServerMultiError(errors)
	}

	return nil
}

// ServerMultiError is an error wrapping multiple validation errors returned by
// Server.ValidateAll() if the designated constraints aren't met.
type ServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerMultiError) AllErrors() []error { return m }

// ServerValidationError is the validation error returned by Server.Validate if
// the designated constraints aren't met.
type ServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerValidationError) ErrorName() string { return "ServerValidationError" }

// Error satisfies the builtin error interface
func (e ServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerValidationError{}

// Validate checks the field values on Server_HTTP with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_HTTP) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_HTTP with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_HTTPMultiError, or
// nil if none found.
func (m *Server_HTTP) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_HTTP) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Addr

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_HTTPValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCors()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Cors",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Cors",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCors()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_HTTPValidationError{
				field:  "Cors",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMiddleware()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Middleware",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_HTTPValidationError{
					field:  "Middleware",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMiddleware()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_HTTPValidationError{
				field:  "Middleware",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Server_HTTPMultiError(errors)
	}

	return nil
}

// Server_HTTPMultiError is an error wrapping multiple validation errors
// returned by Server_HTTP.ValidateAll() if the designated constraints aren't met.
type Server_HTTPMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_HTTPMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_HTTPMultiError) AllErrors() []error { return m }

// Server_HTTPValidationError is the validation error returned by
// Server_HTTP.Validate if the designated constraints aren't met.
type Server_HTTPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_HTTPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_HTTPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_HTTPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_HTTPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_HTTPValidationError) ErrorName() string { return "Server_HTTPValidationError" }

// Error satisfies the builtin error interface
func (e Server_HTTPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_HTTP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_HTTPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_HTTPValidationError{}

// Validate checks the field values on Server_GRPC with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_GRPC) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_GRPC with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_GRPCMultiError, or
// nil if none found.
func (m *Server_GRPC) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_GRPC) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Addr

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_GRPCValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_GRPCValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_GRPCValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMiddleware()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_GRPCValidationError{
					field:  "Middleware",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_GRPCValidationError{
					field:  "Middleware",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMiddleware()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_GRPCValidationError{
				field:  "Middleware",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Server_GRPCMultiError(errors)
	}

	return nil
}

// Server_GRPCMultiError is an error wrapping multiple validation errors
// returned by Server_GRPC.ValidateAll() if the designated constraints aren't met.
type Server_GRPCMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_GRPCMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_GRPCMultiError) AllErrors() []error { return m }

// Server_GRPCValidationError is the validation error returned by
// Server_GRPC.Validate if the designated constraints aren't met.
type Server_GRPCValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_GRPCValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_GRPCValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_GRPCValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_GRPCValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_GRPCValidationError) ErrorName() string { return "Server_GRPCValidationError" }

// Error satisfies the builtin error interface
func (e Server_GRPCValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_GRPC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_GRPCValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_GRPCValidationError{}

// Validate checks the field values on Server_Websocket with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Server_Websocket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_Websocket with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Server_WebsocketMultiError, or nil if none found.
func (m *Server_Websocket) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_Websocket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Addr

	// no validation rules for Path

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Server_WebsocketValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Server_WebsocketValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_WebsocketValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Server_WebsocketMultiError(errors)
	}

	return nil
}

// Server_WebsocketMultiError is an error wrapping multiple validation errors
// returned by Server_Websocket.ValidateAll() if the designated constraints
// aren't met.
type Server_WebsocketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_WebsocketMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_WebsocketMultiError) AllErrors() []error { return m }

// Server_WebsocketValidationError is the validation error returned by
// Server_Websocket.Validate if the designated constraints aren't met.
type Server_WebsocketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_WebsocketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_WebsocketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_WebsocketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_WebsocketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_WebsocketValidationError) ErrorName() string { return "Server_WebsocketValidationError" }

// Error satisfies the builtin error interface
func (e Server_WebsocketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_Websocket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_WebsocketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_WebsocketValidationError{}

// Validate checks the field values on Server_Mqtt with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_Mqtt) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_Mqtt with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_MqttMultiError, or
// nil if none found.
func (m *Server_Mqtt) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_Mqtt) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	if len(errors) > 0 {
		return Server_MqttMultiError(errors)
	}

	return nil
}

// Server_MqttMultiError is an error wrapping multiple validation errors
// returned by Server_Mqtt.ValidateAll() if the designated constraints aren't met.
type Server_MqttMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_MqttMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_MqttMultiError) AllErrors() []error { return m }

// Server_MqttValidationError is the validation error returned by
// Server_Mqtt.Validate if the designated constraints aren't met.
type Server_MqttValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_MqttValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_MqttValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_MqttValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_MqttValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_MqttValidationError) ErrorName() string { return "Server_MqttValidationError" }

// Error satisfies the builtin error interface
func (e Server_MqttValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_Mqtt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_MqttValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_MqttValidationError{}

// Validate checks the field values on Server_Kafka with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_Kafka) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_Kafka with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_KafkaMultiError, or
// nil if none found.
func (m *Server_Kafka) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_Kafka) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Server_KafkaMultiError(errors)
	}

	return nil
}

// Server_KafkaMultiError is an error wrapping multiple validation errors
// returned by Server_Kafka.ValidateAll() if the designated constraints aren't met.
type Server_KafkaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_KafkaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_KafkaMultiError) AllErrors() []error { return m }

// Server_KafkaValidationError is the validation error returned by
// Server_Kafka.Validate if the designated constraints aren't met.
type Server_KafkaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_KafkaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_KafkaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_KafkaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_KafkaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_KafkaValidationError) ErrorName() string { return "Server_KafkaValidationError" }

// Error satisfies the builtin error interface
func (e Server_KafkaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_Kafka.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_KafkaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_KafkaValidationError{}

// Validate checks the field values on Server_RabbitMQ with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Server_RabbitMQ) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_RabbitMQ with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Server_RabbitMQMultiError, or nil if none found.
func (m *Server_RabbitMQ) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_RabbitMQ) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Server_RabbitMQMultiError(errors)
	}

	return nil
}

// Server_RabbitMQMultiError is an error wrapping multiple validation errors
// returned by Server_RabbitMQ.ValidateAll() if the designated constraints
// aren't met.
type Server_RabbitMQMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_RabbitMQMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_RabbitMQMultiError) AllErrors() []error { return m }

// Server_RabbitMQValidationError is the validation error returned by
// Server_RabbitMQ.Validate if the designated constraints aren't met.
type Server_RabbitMQValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_RabbitMQValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_RabbitMQValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_RabbitMQValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_RabbitMQValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_RabbitMQValidationError) ErrorName() string { return "Server_RabbitMQValidationError" }

// Error satisfies the builtin error interface
func (e Server_RabbitMQValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_RabbitMQ.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_RabbitMQValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_RabbitMQValidationError{}

// Validate checks the field values on Server_HTTP_CORS with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Server_HTTP_CORS) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_HTTP_CORS with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Server_HTTP_CORSMultiError, or nil if none found.
func (m *Server_HTTP_CORS) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_HTTP_CORS) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Server_HTTP_CORSMultiError(errors)
	}

	return nil
}

// Server_HTTP_CORSMultiError is an error wrapping multiple validation errors
// returned by Server_HTTP_CORS.ValidateAll() if the designated constraints
// aren't met.
type Server_HTTP_CORSMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_HTTP_CORSMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_HTTP_CORSMultiError) AllErrors() []error { return m }

// Server_HTTP_CORSValidationError is the validation error returned by
// Server_HTTP_CORS.Validate if the designated constraints aren't met.
type Server_HTTP_CORSValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_HTTP_CORSValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_HTTP_CORSValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_HTTP_CORSValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_HTTP_CORSValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_HTTP_CORSValidationError) ErrorName() string { return "Server_HTTP_CORSValidationError" }

// Error satisfies the builtin error interface
func (e Server_HTTP_CORSValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_HTTP_CORS.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_HTTP_CORSValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_HTTP_CORSValidationError{}
