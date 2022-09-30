// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNSPECIFIED.String() && e.Code == 400
}

func ErrorUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 400
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserLoginFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_LOGIN_FAIL.String() && e.Code == 400
}

func ErrorUserLoginFail(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_USER_LOGIN_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsUserRegisterFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_REGISTER_FAIL.String() && e.Code == 400
}

func ErrorUserRegisterFail(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_USER_REGISTER_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsDomainNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DOMAIN_NOT_FOUND.String() && e.Code == 400
}

func ErrorDomainNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DOMAIN_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsMenuNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MENU_NOT_FOUND.String() && e.Code == 400
}

func ErrorMenuNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_MENU_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsApiNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_API_NOT_FOUND.String() && e.Code == 400
}

func ErrorApiNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_API_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}
