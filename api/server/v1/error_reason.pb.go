// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: server/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	ErrorReason_UNSPECIFIED               ErrorReason = 0
	ErrorReason_SYSTEM_NOT_FOUND          ErrorReason = 1
	ErrorReason_USER_NOT_FOUND            ErrorReason = 101
	ErrorReason_USER_LOGIN_FAIL           ErrorReason = 102
	ErrorReason_USER_REGISTER_FAIL        ErrorReason = 103
	ErrorReason_USER_CREATE_FAIL          ErrorReason = 104
	ErrorReason_USER_ID_NULL              ErrorReason = 105
	ErrorReason_USER_UPDATE_FAIL          ErrorReason = 106
	ErrorReason_USER_DELETE_FAIL          ErrorReason = 107
	ErrorReason_USER_HANDLE_FAIL          ErrorReason = 108
	ErrorReason_USER_HANDLE_ROLE_FAIL     ErrorReason = 109
	ErrorReason_USER_DOMAIN_FIND_FAIL     ErrorReason = 110
	ErrorReason_USER_ROLE_FIND_FAIL       ErrorReason = 111
	ErrorReason_DOMAIN_NOT_FOUND          ErrorReason = 201
	ErrorReason_DOMAIN_CREATE_FAIL        ErrorReason = 202
	ErrorReason_DOMAIN_UPDATE_FAIL        ErrorReason = 203
	ErrorReason_DOMAIN_DELETE_FAIL        ErrorReason = 204
	ErrorReason_DOMAIN_HANDLE_MENU_FAIL   ErrorReason = 205
	ErrorReason_ROLE_NOT_FOUND            ErrorReason = 301
	ErrorReason_ROLE_CREATE_FAIL          ErrorReason = 302
	ErrorReason_ROLE_UPDATE_FAIL          ErrorReason = 303
	ErrorReason_ROLE_DELETE_FAIL          ErrorReason = 304
	ErrorReason_ROLE_HANDLE_MENU_FAIL     ErrorReason = 305
	ErrorReason_ROLE_HANDLE_RESOURCE_FAIL ErrorReason = 306
	ErrorReason_MENU_NOT_FOUND            ErrorReason = 401
	ErrorReason_MENU_CREATE_FAIL          ErrorReason = 402
	ErrorReason_MENU_UPDATE_FAIL          ErrorReason = 403
	ErrorReason_MENU_DELETE_FAIL          ErrorReason = 404
	ErrorReason_RESOURCE_NOT_FOUND        ErrorReason = 501
	ErrorReason_RESOURCE_CREATE_FAIL      ErrorReason = 502
	ErrorReason_RESOURCE_UPDATE_FAIL      ErrorReason = 503
	ErrorReason_RESOURCE_DELETE_FAIL      ErrorReason = 504
	ErrorReason_DEPT_NOT_FOUND            ErrorReason = 601
	ErrorReason_DEPT_CREATE_FAIL          ErrorReason = 602
	ErrorReason_DEPT_UPDATE_FAIL          ErrorReason = 603
	ErrorReason_DEPT_DELETE_FAIL          ErrorReason = 604
	ErrorReason_POST_NOT_FOUND            ErrorReason = 701
	ErrorReason_POST_CREATE_FAIL          ErrorReason = 702
	ErrorReason_POST_UPDATE_FAIL          ErrorReason = 703
	ErrorReason_POST_DELETE_FAIL          ErrorReason = 704
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "UNSPECIFIED",
		1:   "SYSTEM_NOT_FOUND",
		101: "USER_NOT_FOUND",
		102: "USER_LOGIN_FAIL",
		103: "USER_REGISTER_FAIL",
		104: "USER_CREATE_FAIL",
		105: "USER_ID_NULL",
		106: "USER_UPDATE_FAIL",
		107: "USER_DELETE_FAIL",
		108: "USER_HANDLE_FAIL",
		109: "USER_HANDLE_ROLE_FAIL",
		110: "USER_DOMAIN_FIND_FAIL",
		111: "USER_ROLE_FIND_FAIL",
		201: "DOMAIN_NOT_FOUND",
		202: "DOMAIN_CREATE_FAIL",
		203: "DOMAIN_UPDATE_FAIL",
		204: "DOMAIN_DELETE_FAIL",
		205: "DOMAIN_HANDLE_MENU_FAIL",
		301: "ROLE_NOT_FOUND",
		302: "ROLE_CREATE_FAIL",
		303: "ROLE_UPDATE_FAIL",
		304: "ROLE_DELETE_FAIL",
		305: "ROLE_HANDLE_MENU_FAIL",
		306: "ROLE_HANDLE_RESOURCE_FAIL",
		401: "MENU_NOT_FOUND",
		402: "MENU_CREATE_FAIL",
		403: "MENU_UPDATE_FAIL",
		404: "MENU_DELETE_FAIL",
		501: "RESOURCE_NOT_FOUND",
		502: "RESOURCE_CREATE_FAIL",
		503: "RESOURCE_UPDATE_FAIL",
		504: "RESOURCE_DELETE_FAIL",
		601: "DEPT_NOT_FOUND",
		602: "DEPT_CREATE_FAIL",
		603: "DEPT_UPDATE_FAIL",
		604: "DEPT_DELETE_FAIL",
		701: "POST_NOT_FOUND",
		702: "POST_CREATE_FAIL",
		703: "POST_UPDATE_FAIL",
		704: "POST_DELETE_FAIL",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":               0,
		"SYSTEM_NOT_FOUND":          1,
		"USER_NOT_FOUND":            101,
		"USER_LOGIN_FAIL":           102,
		"USER_REGISTER_FAIL":        103,
		"USER_CREATE_FAIL":          104,
		"USER_ID_NULL":              105,
		"USER_UPDATE_FAIL":          106,
		"USER_DELETE_FAIL":          107,
		"USER_HANDLE_FAIL":          108,
		"USER_HANDLE_ROLE_FAIL":     109,
		"USER_DOMAIN_FIND_FAIL":     110,
		"USER_ROLE_FIND_FAIL":       111,
		"DOMAIN_NOT_FOUND":          201,
		"DOMAIN_CREATE_FAIL":        202,
		"DOMAIN_UPDATE_FAIL":        203,
		"DOMAIN_DELETE_FAIL":        204,
		"DOMAIN_HANDLE_MENU_FAIL":   205,
		"ROLE_NOT_FOUND":            301,
		"ROLE_CREATE_FAIL":          302,
		"ROLE_UPDATE_FAIL":          303,
		"ROLE_DELETE_FAIL":          304,
		"ROLE_HANDLE_MENU_FAIL":     305,
		"ROLE_HANDLE_RESOURCE_FAIL": 306,
		"MENU_NOT_FOUND":            401,
		"MENU_CREATE_FAIL":          402,
		"MENU_UPDATE_FAIL":          403,
		"MENU_DELETE_FAIL":          404,
		"RESOURCE_NOT_FOUND":        501,
		"RESOURCE_CREATE_FAIL":      502,
		"RESOURCE_UPDATE_FAIL":      503,
		"RESOURCE_DELETE_FAIL":      504,
		"DEPT_NOT_FOUND":            601,
		"DEPT_CREATE_FAIL":          602,
		"DEPT_UPDATE_FAIL":          603,
		"DEPT_DELETE_FAIL":          604,
		"POST_NOT_FOUND":            701,
		"POST_CREATE_FAIL":          702,
		"POST_UPDATE_FAIL":          703,
		"POST_DELETE_FAIL":          704,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_server_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_server_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_server_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_server_v1_error_reason_proto protoreflect.FileDescriptor

var file_server_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc8, 0x07, 0x0a, 0x0b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x66, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x67, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x68, 0x12,
	0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10,
	0x69, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6a, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6b, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x10, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44,
	0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6d, 0x12, 0x19,
	0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x46, 0x49,
	0x4e, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x10, 0x6f, 0x12, 0x15, 0x0a, 0x10, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xc9, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x4f, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0xca, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xcb, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x44,
	0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x10, 0xcc, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x48,
	0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x4e, 0x55, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0xcd, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0xad, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xae, 0x02, 0x12, 0x15,
	0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0xaf, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xb0, 0x02, 0x12, 0x1a, 0x0a, 0x15,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x4e, 0x55,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xb1, 0x02, 0x12, 0x1e, 0x0a, 0x19, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xb2, 0x02, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x45, 0x4e, 0x55,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x91, 0x03, 0x12, 0x15, 0x0a,
	0x10, 0x4d, 0x45, 0x4e, 0x55, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x10, 0x92, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x45, 0x4e, 0x55, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x93, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x4d,
	0x45, 0x4e, 0x55, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x94, 0x03, 0x12, 0x17, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xf5, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0xf6, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xf7,
	0x03, 0x12, 0x19, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xf8, 0x03, 0x12, 0x13, 0x0a, 0x0e,
	0x44, 0x45, 0x50, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd9,
	0x04, 0x12, 0x15, 0x0a, 0x10, 0x44, 0x45, 0x50, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xda, 0x04, 0x12, 0x15, 0x0a, 0x10, 0x44, 0x45, 0x50, 0x54,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xdb, 0x04, 0x12,
	0x15, 0x0a, 0x10, 0x44, 0x45, 0x50, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0xdc, 0x04, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xbd, 0x05, 0x12, 0x15, 0x0a, 0x10, 0x50,
	0x4f, 0x53, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0xbe, 0x05, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xbf, 0x05, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x4f, 0x53,
	0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xc0, 0x05,
	0x1a, 0x04, 0xa0, 0x45, 0x90, 0x03, 0x42, 0x4c, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66,
	0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x0d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x61,
	0x70, 0x69, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_v1_error_reason_proto_rawDescOnce sync.Once
	file_server_v1_error_reason_proto_rawDescData = file_server_v1_error_reason_proto_rawDesc
)

func file_server_v1_error_reason_proto_rawDescGZIP() []byte {
	file_server_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_server_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_v1_error_reason_proto_rawDescData)
	})
	return file_server_v1_error_reason_proto_rawDescData
}

var file_server_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.v1.ErrorReason
}
var file_server_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_v1_error_reason_proto_init() }
func file_server_v1_error_reason_proto_init() {
	if File_server_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_server_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_server_v1_error_reason_proto_enumTypes,
	}.Build()
	File_server_v1_error_reason_proto = out.File
	file_server_v1_error_reason_proto_rawDesc = nil
	file_server_v1_error_reason_proto_goTypes = nil
	file_server_v1_error_reason_proto_depIdxs = nil
}