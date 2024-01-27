// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/pagination/pagination.proto

package pagination

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 分页通用请求
type PagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前页码
	Page *int32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// 每页的行数
	PageSize *int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// 与过滤参数
	Query *string `protobuf:"bytes,3,opt,name=query,proto3,oneof" json:"query,omitempty"`
	// 或过滤参数
	OrQuery *string `protobuf:"bytes,4,opt,name=or_query,json=or,proto3,oneof" json:"or_query,omitempty"`
	// 排序条件
	OrderBy []string `protobuf:"bytes,5,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// 是否不分页
	NoPaging *bool `protobuf:"varint,6,opt,name=no_paging,json=nopaging,proto3,oneof" json:"no_paging,omitempty"`
	// 字段掩码
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,7,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *PagingRequest) Reset() {
	*x = PagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pagination_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingRequest) ProtoMessage() {}

func (x *PagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingRequest.ProtoReflect.Descriptor instead.
func (*PagingRequest) Descriptor() ([]byte, []int) {
	return file_common_pagination_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PagingRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *PagingRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *PagingRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *PagingRequest) GetOrQuery() string {
	if x != nil && x.OrQuery != nil {
		return *x.OrQuery
	}
	return ""
}

func (x *PagingRequest) GetOrderBy() []string {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *PagingRequest) GetNoPaging() bool {
	if x != nil && x.NoPaging != nil {
		return *x.NoPaging
	}
	return false
}

func (x *PagingRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

// 分页通用结果
type PagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 分页数据
	Items []*anypb.Any `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PagingResponse) Reset() {
	*x = PagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pagination_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingResponse) ProtoMessage() {}

func (x *PagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingResponse.ProtoReflect.Descriptor instead.
func (*PagingResponse) Descriptor() ([]byte, []int) {
	return file_common_pagination_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PagingResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagingResponse) GetItems() []*anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_common_pagination_pagination_proto protoreflect.FileDescriptor

var file_common_pagination_pagination_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x05, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x1e, 0xba, 0x47, 0x1b, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xf0, 0x3f, 0x92, 0x02, 0x0c, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe9, 0xa1, 0xb5,
	0xe7, 0xa0, 0x81, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x24, 0xba, 0x47, 0x21, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x24, 0x40, 0x92, 0x02, 0x12, 0xe6, 0xaf, 0x8f, 0xe4, 0xb8, 0x80, 0xe9, 0xa1, 0xb5, 0xe7, 0x9a,
	0x84, 0xe8, 0xa1, 0x8c, 0xe6, 0x95, 0xb0, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xba, 0x47, 0x33, 0x3a, 0x1f, 0x12, 0x1d, 0x7b, 0x22,
	0x6b, 0x65, 0x79, 0x31, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x22, 0x6b, 0x65,
	0x79, 0x32, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0x92, 0x02, 0x0f, 0xe4, 0xb8,
	0x8e, 0xe8, 0xbf, 0x87, 0xe6, 0xbb, 0xa4, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x48, 0x02, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x08, 0x6f, 0x72, 0x5f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xba, 0x47, 0x33,
	0x3a, 0x1f, 0x12, 0x1d, 0x7b, 0x22, 0x6b, 0x65, 0x79, 0x31, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c,
	0x31, 0x22, 0x2c, 0x22, 0x6b, 0x65, 0x79, 0x32, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x32, 0x22,
	0x7d, 0x92, 0x02, 0x0f, 0xe6, 0x88, 0x96, 0xe8, 0xbf, 0x87, 0xe6, 0xbb, 0xa4, 0xe5, 0x8f, 0x82,
	0xe6, 0x95, 0xb0, 0x48, 0x03, 0x52, 0x02, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x75, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x5a,
	0xba, 0x47, 0x57, 0x3a, 0x13, 0x12, 0x11, 0x7b, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x20,
	0x22, 0x2d, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0x92, 0x02, 0x3f, 0xe6, 0x8e, 0x92, 0xe5, 0xba,
	0x8f, 0xe6, 0x9d, 0xa1, 0xe4, 0xbb, 0xb6, 0xef, 0xbc, 0x8c, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5,
	0xe5, 0x90, 0x8d, 0xe5, 0x89, 0x8d, 0xe5, 0x8a, 0xa0, 0x27, 0x2d, 0x27, 0xe4, 0xb8, 0xba, 0xe9,
	0x99, 0x8d, 0xe5, 0xba, 0x8f, 0xef, 0xbc, 0x8c, 0xe5, 0x90, 0xa6, 0xe5, 0x88, 0x99, 0xe4, 0xb8,
	0xba, 0xe5, 0x8d, 0x87, 0xe5, 0xba, 0x8f, 0xe3, 0x80, 0x82, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x15, 0xba, 0x47, 0x12, 0x92, 0x02, 0x0f, 0xe6, 0x98,
	0xaf, 0xe5, 0x90, 0xa6, 0xe4, 0xb8, 0x8d, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0x48, 0x04, 0x52,
	0x08, 0x6e, 0x6f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x8c, 0x01, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x51, 0xba,
	0x47, 0x4e, 0x3a, 0x16, 0x12, 0x14, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x92, 0x02, 0x33, 0xe5, 0xad, 0x97,
	0xe6, 0xae, 0xb5, 0xe6, 0x8e, 0xa9, 0xe7, 0xa0, 0x81, 0xef, 0xbc, 0x8c, 0xe5, 0xa6, 0x82, 0xe6,
	0x9e, 0x9c, 0xe4, 0xb8, 0xba, 0xe7, 0xa9, 0xba, 0xe5, 0x88, 0x99, 0xe9, 0x80, 0x89, 0xe4, 0xb8,
	0xad, 0xe6, 0x89, 0x80, 0xe6, 0x9c, 0x89, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe3, 0x80, 0x82,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x6f,
	0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f,
	0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2d, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pagination_pagination_proto_rawDescOnce sync.Once
	file_common_pagination_pagination_proto_rawDescData = file_common_pagination_pagination_proto_rawDesc
)

func file_common_pagination_pagination_proto_rawDescGZIP() []byte {
	file_common_pagination_pagination_proto_rawDescOnce.Do(func() {
		file_common_pagination_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pagination_pagination_proto_rawDescData)
	})
	return file_common_pagination_pagination_proto_rawDescData
}

var file_common_pagination_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_pagination_pagination_proto_goTypes = []interface{}{
	(*PagingRequest)(nil),         // 0: pagination.PagingRequest
	(*PagingResponse)(nil),        // 1: pagination.PagingResponse
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
}
var file_common_pagination_pagination_proto_depIdxs = []int32{
	2, // 0: pagination.PagingRequest.field_mask:type_name -> google.protobuf.FieldMask
	3, // 1: pagination.PagingResponse.items:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_pagination_pagination_proto_init() }
func file_common_pagination_pagination_proto_init() {
	if File_common_pagination_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_pagination_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_pagination_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_pagination_pagination_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_pagination_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pagination_pagination_proto_goTypes,
		DependencyIndexes: file_common_pagination_pagination_proto_depIdxs,
		MessageInfos:      file_common_pagination_pagination_proto_msgTypes,
	}.Build()
	File_common_pagination_pagination_proto = out.File
	file_common_pagination_pagination_proto_rawDesc = nil
	file_common_pagination_pagination_proto_goTypes = nil
	file_common_pagination_pagination_proto_depIdxs = nil
}
