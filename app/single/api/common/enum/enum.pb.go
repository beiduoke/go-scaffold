// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/enum/enum.proto

package enum

import (
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

// 启用状态
type Enable int32

const (
	Enable_ENABLE_UNSPECIFIED Enable = 0
	Enable_ENABLE_YES         Enable = 1
	Enable_ENABLE_NO          Enable = 2
)

// Enum value maps for Enable.
var (
	Enable_name = map[int32]string{
		0: "ENABLE_UNSPECIFIED",
		1: "ENABLE_YES",
		2: "ENABLE_NO",
	}
	Enable_value = map[string]int32{
		"ENABLE_UNSPECIFIED": 0,
		"ENABLE_YES":         1,
		"ENABLE_NO":          2,
	}
)

func (x Enable) Enum() *Enable {
	p := new(Enable)
	*p = x
	return p
}

func (x Enable) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enable) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_enum_proto_enumTypes[0].Descriptor()
}

func (Enable) Type() protoreflect.EnumType {
	return &file_common_enum_enum_proto_enumTypes[0]
}

func (x Enable) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enable.Descriptor instead.
func (Enable) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_enum_proto_rawDescGZIP(), []int{0}
}

// 通用排序
type Sort int32

const (
	Sort_SORT_UNSPECIFIED Sort = 0
	Sort_SORT_DESC        Sort = 1
	Sort_SORT_ASC         Sort = 2
)

// Enum value maps for Sort.
var (
	Sort_name = map[int32]string{
		0: "SORT_UNSPECIFIED",
		1: "SORT_DESC",
		2: "SORT_ASC",
	}
	Sort_value = map[string]int32{
		"SORT_UNSPECIFIED": 0,
		"SORT_DESC":        1,
		"SORT_ASC":         2,
	}
)

func (x Sort) Enum() *Sort {
	p := new(Sort)
	*p = x
	return p
}

func (x Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_enum_proto_enumTypes[1].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_common_enum_enum_proto_enumTypes[1]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_enum_proto_rawDescGZIP(), []int{1}
}

// 通用状态
type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	State_STATE_ACTIVE      State = 1
	State_STATE_INACTIVE    State = 2
	State_STATE_BANNED      State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACTIVE",
		2: "STATE_INACTIVE",
		3: "STATE_BANNED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_ACTIVE":      1,
		"STATE_INACTIVE":    2,
		"STATE_BANNED":      3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_enum_proto_enumTypes[2].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_common_enum_enum_proto_enumTypes[2]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_enum_proto_rawDescGZIP(), []int{2}
}

// 开关
type Switch int32

const (
	Switch_SWITCH_UNSPECIFIED Switch = 0
	Switch_SWITCH_OPEN        Switch = 1
	Switch_SWITCH_CLOSE       Switch = 2
)

// Enum value maps for Switch.
var (
	Switch_name = map[int32]string{
		0: "SWITCH_UNSPECIFIED",
		1: "SWITCH_OPEN",
		2: "SWITCH_CLOSE",
	}
	Switch_value = map[string]int32{
		"SWITCH_UNSPECIFIED": 0,
		"SWITCH_OPEN":        1,
		"SWITCH_CLOSE":       2,
	}
)

func (x Switch) Enum() *Switch {
	p := new(Switch)
	*p = x
	return p
}

func (x Switch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Switch) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_enum_proto_enumTypes[3].Descriptor()
}

func (Switch) Type() protoreflect.EnumType {
	return &file_common_enum_enum_proto_enumTypes[3]
}

func (x Switch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Switch.Descriptor instead.
func (Switch) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_enum_proto_rawDescGZIP(), []int{3}
}

var File_common_enum_enum_proto protoreflect.FileDescriptor

var file_common_enum_enum_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x2a, 0x3f,
	0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x59, 0x45, 0x53, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x10, 0x02, 0x2a,
	0x39, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f, 0x52, 0x54, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x02, 0x2a, 0x56, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x4e, 0x45, 0x44,
	0x10, 0x03, 0x2a, 0x43, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x42, 0x65, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x62,
	0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x42, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65,
	0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f,
	0x6c, 0x64, 0x2d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enum_enum_proto_rawDescOnce sync.Once
	file_common_enum_enum_proto_rawDescData = file_common_enum_enum_proto_rawDesc
)

func file_common_enum_enum_proto_rawDescGZIP() []byte {
	file_common_enum_enum_proto_rawDescOnce.Do(func() {
		file_common_enum_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enum_enum_proto_rawDescData)
	})
	return file_common_enum_enum_proto_rawDescData
}

var file_common_enum_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_enum_enum_proto_goTypes = []interface{}{
	(Enable)(0), // 0: enum.Enable
	(Sort)(0),   // 1: enum.Sort
	(State)(0),  // 2: enum.State
	(Switch)(0), // 3: enum.Switch
}
var file_common_enum_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enum_enum_proto_init() }
func file_common_enum_enum_proto_init() {
	if File_common_enum_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enum_enum_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enum_enum_proto_goTypes,
		DependencyIndexes: file_common_enum_enum_proto_depIdxs,
		EnumInfos:         file_common_enum_enum_proto_enumTypes,
	}.Build()
	File_common_enum_enum_proto = out.File
	file_common_enum_enum_proto_rawDesc = nil
	file_common_enum_enum_proto_goTypes = nil
	file_common_enum_enum_proto_depIdxs = nil
}
