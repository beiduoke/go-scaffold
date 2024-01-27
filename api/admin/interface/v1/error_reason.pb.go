// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: admin/interface/v1/error_reason.proto

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
	ErrorReason_BAD_REQUEST           ErrorReason = 0  // 400
	ErrorReason_NOT_LOGGED_IN         ErrorReason = 1  // 401
	ErrorReason_ACCESS_FORBIDDEN      ErrorReason = 2  // 403
	ErrorReason_RESOURCE_NOT_FOUND    ErrorReason = 3  // 404
	ErrorReason_METHOD_NOT_ALLOWED    ErrorReason = 4  // 405
	ErrorReason_REQUEST_TIMEOUT       ErrorReason = 5  // 408
	ErrorReason_INTERNAL_SERVER_ERROR ErrorReason = 10 // 500
	ErrorReason_NOT_IMPLEMENTED       ErrorReason = 11 // 501
	ErrorReason_NETWORK_ERROR         ErrorReason = 12 // 502
	ErrorReason_SERVICE_UNAVAILABLE   ErrorReason = 13 // 503
	ErrorReason_NETWORK_TIMEOUT       ErrorReason = 14 // 504
	ErrorReason_REQUEST_NOT_SUPPORT   ErrorReason = 15 // 505
	// 认证相关
	ErrorReason_AUTH_TOKEN_EXPIRED   ErrorReason = 30 // token过期
	ErrorReason_AUTH_INVALID_TOKEN   ErrorReason = 31 // token无效
	ErrorReason_AUTH_TOKEN_NOT_EXIST ErrorReason = 32 // token不存在
	// 用户相关错误
	ErrorReason_USER_NOT_FOUND          ErrorReason = 50
	ErrorReason_USER_NOT_EXIST          ErrorReason = 51 // 用户不存在
	ErrorReason_USER_INCORRECT_PASSWORD ErrorReason = 52 // 密码错误
	ErrorReason_USER_FREEZE             ErrorReason = 53 // 用户冻结
	ErrorReason_USER_INVALID_ID         ErrorReason = 54 // 用户ID无效
	ErrorReason_USER_INVALID_PASSWORD   ErrorReason = 55 // 密码无效
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "BAD_REQUEST",
		1:  "NOT_LOGGED_IN",
		2:  "ACCESS_FORBIDDEN",
		3:  "RESOURCE_NOT_FOUND",
		4:  "METHOD_NOT_ALLOWED",
		5:  "REQUEST_TIMEOUT",
		10: "INTERNAL_SERVER_ERROR",
		11: "NOT_IMPLEMENTED",
		12: "NETWORK_ERROR",
		13: "SERVICE_UNAVAILABLE",
		14: "NETWORK_TIMEOUT",
		15: "REQUEST_NOT_SUPPORT",
		30: "AUTH_TOKEN_EXPIRED",
		31: "AUTH_INVALID_TOKEN",
		32: "AUTH_TOKEN_NOT_EXIST",
		50: "USER_NOT_FOUND",
		51: "USER_NOT_EXIST",
		52: "USER_INCORRECT_PASSWORD",
		53: "USER_FREEZE",
		54: "USER_INVALID_ID",
		55: "USER_INVALID_PASSWORD",
	}
	ErrorReason_value = map[string]int32{
		"BAD_REQUEST":             0,
		"NOT_LOGGED_IN":           1,
		"ACCESS_FORBIDDEN":        2,
		"RESOURCE_NOT_FOUND":      3,
		"METHOD_NOT_ALLOWED":      4,
		"REQUEST_TIMEOUT":         5,
		"INTERNAL_SERVER_ERROR":   10,
		"NOT_IMPLEMENTED":         11,
		"NETWORK_ERROR":           12,
		"SERVICE_UNAVAILABLE":     13,
		"NETWORK_TIMEOUT":         14,
		"REQUEST_NOT_SUPPORT":     15,
		"AUTH_TOKEN_EXPIRED":      30,
		"AUTH_INVALID_TOKEN":      31,
		"AUTH_TOKEN_NOT_EXIST":    32,
		"USER_NOT_FOUND":          50,
		"USER_NOT_EXIST":          51,
		"USER_INCORRECT_PASSWORD": 52,
		"USER_FREEZE":             53,
		"USER_INVALID_ID":         54,
		"USER_INVALID_PASSWORD":   55,
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
	return file_admin_interface_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_admin_interface_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_admin_interface_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_admin_interface_v1_error_reason_proto protoreflect.FileDescriptor

var file_admin_interface_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xe1, 0x04, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x15, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c,
	0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x1a, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49,
	0x44, 0x44, 0x45, 0x4e, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1c, 0x0a, 0x12,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x1a, 0x04, 0xa8,
	0x45, 0x98, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0a, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12,
	0x17, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x0d, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45,
	0xf8, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x0f, 0x1a, 0x04, 0xa8, 0x45, 0xf9,
	0x03, 0x12, 0x1b, 0x0a, 0x12, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x1e, 0x1a, 0x03, 0xa8, 0x45, 0x67, 0x12, 0x1b,
	0x0a, 0x12, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x1f, 0x1a, 0x03, 0xa8, 0x45, 0x68, 0x12, 0x1d, 0x0a, 0x14, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58,
	0x49, 0x53, 0x54, 0x10, 0x20, 0x1a, 0x03, 0xa8, 0x45, 0x69, 0x12, 0x17, 0x0a, 0x0e, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x32, 0x1a, 0x03,
	0xa8, 0x45, 0x6e, 0x12, 0x17, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x33, 0x1a, 0x03, 0xa8, 0x45, 0x70, 0x12, 0x20, 0x0a, 0x17,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x50,
	0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x34, 0x1a, 0x03, 0xa8, 0x45, 0x71, 0x12, 0x14,
	0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x5a, 0x45, 0x10, 0x35, 0x1a,
	0x03, 0xa8, 0x45, 0x72, 0x12, 0x18, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x36, 0x1a, 0x03, 0xa8, 0x45, 0x73, 0x12, 0x1e,
	0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50,
	0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x37, 0x1a, 0x03, 0xa8, 0x45, 0x74, 0x1a, 0x04,
	0xa0, 0x45, 0xf4, 0x03, 0x42, 0x6c, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0xa2, 0x02, 0x0d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x61, 0x70, 0x69,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_interface_v1_error_reason_proto_rawDescOnce sync.Once
	file_admin_interface_v1_error_reason_proto_rawDescData = file_admin_interface_v1_error_reason_proto_rawDesc
)

func file_admin_interface_v1_error_reason_proto_rawDescGZIP() []byte {
	file_admin_interface_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_admin_interface_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_interface_v1_error_reason_proto_rawDescData)
	})
	return file_admin_interface_v1_error_reason_proto_rawDescData
}

var file_admin_interface_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_interface_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: admin.interface.v1.ErrorReason
}
var file_admin_interface_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_interface_v1_error_reason_proto_init() }
func file_admin_interface_v1_error_reason_proto_init() {
	if File_admin_interface_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_interface_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_interface_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_admin_interface_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_admin_interface_v1_error_reason_proto_enumTypes,
	}.Build()
	File_admin_interface_v1_error_reason_proto = out.File
	file_admin_interface_v1_error_reason_proto_rawDesc = nil
	file_admin_interface_v1_error_reason_proto_goTypes = nil
	file_admin_interface_v1_error_reason_proto_depIdxs = nil
}
