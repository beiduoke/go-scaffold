// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: admin/v1/error_reason.proto

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
	ErrorReason_UNSPECIFIED                ErrorReason = 0
	ErrorReason_SYSTEM_NOT_FOUND           ErrorReason = 1
	ErrorReason_USER_NOT_FOUND             ErrorReason = 101
	ErrorReason_USER_LOGIN_FAIL            ErrorReason = 102
	ErrorReason_USER_REGISTER_FAIL         ErrorReason = 103
	ErrorReason_USER_CREATE_FAIL           ErrorReason = 104
	ErrorReason_USER_ID_NULL               ErrorReason = 105
	ErrorReason_USER_UPDATE_FAIL           ErrorReason = 106
	ErrorReason_USER_DELETE_FAIL           ErrorReason = 107
	ErrorReason_USER_HANDLE_DOMAIN_FAIL    ErrorReason = 108
	ErrorReason_USER_HANDLE_AUTHORITY_FAIL ErrorReason = 109
	ErrorReason_DOMAIN_NOT_FOUND           ErrorReason = 201
	ErrorReason_MENU_NOT_FOUND             ErrorReason = 301
	ErrorReason_API_NOT_FOUND              ErrorReason = 401
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
		108: "USER_HANDLE_DOMAIN_FAIL",
		109: "USER_HANDLE_AUTHORITY_FAIL",
		201: "DOMAIN_NOT_FOUND",
		301: "MENU_NOT_FOUND",
		401: "API_NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":                0,
		"SYSTEM_NOT_FOUND":           1,
		"USER_NOT_FOUND":             101,
		"USER_LOGIN_FAIL":            102,
		"USER_REGISTER_FAIL":         103,
		"USER_CREATE_FAIL":           104,
		"USER_ID_NULL":               105,
		"USER_UPDATE_FAIL":           106,
		"USER_DELETE_FAIL":           107,
		"USER_HANDLE_DOMAIN_FAIL":    108,
		"USER_HANDLE_AUTHORITY_FAIL": 109,
		"DOMAIN_NOT_FOUND":           201,
		"MENU_NOT_FOUND":             301,
		"API_NOT_FOUND":              401,
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
	return file_admin_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_admin_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_admin_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_admin_v1_error_reason_proto protoreflect.FileDescriptor

var file_admin_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd2, 0x02, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x10, 0x66, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x67, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x68,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x69, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6a, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6b, 0x12, 0x1b,
	0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x44, 0x4f,
	0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6c, 0x12, 0x1e, 0x0a, 0x1a, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6d, 0x12, 0x15, 0x0a, 0x10, 0x44,
	0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0xc9, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x45, 0x4e, 0x55, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0xad, 0x02, 0x12, 0x12, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x91, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0x90,
	0x03, 0x42, 0x4a, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64,
	0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0xa2, 0x02, 0x0a, 0x41, 0x50, 0x49, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_error_reason_proto_rawDescOnce sync.Once
	file_admin_v1_error_reason_proto_rawDescData = file_admin_v1_error_reason_proto_rawDesc
)

func file_admin_v1_error_reason_proto_rawDescGZIP() []byte {
	file_admin_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_admin_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_error_reason_proto_rawDescData)
	})
	return file_admin_v1_error_reason_proto_rawDescData
}

var file_admin_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: admin.v1.ErrorReason
}
var file_admin_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_v1_error_reason_proto_init() }
func file_admin_v1_error_reason_proto_init() {
	if File_admin_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_admin_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_admin_v1_error_reason_proto_enumTypes,
	}.Build()
	File_admin_v1_error_reason_proto = out.File
	file_admin_v1_error_reason_proto_rawDesc = nil
	file_admin_v1_error_reason_proto_goTypes = nil
	file_admin_v1_error_reason_proto_depIdxs = nil
}
