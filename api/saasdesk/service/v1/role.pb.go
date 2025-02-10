// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: saasdesk/service/v1/role.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 角色模块
type Role struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	Id                uint64                 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DefaultRouter     *string                `protobuf:"bytes,6,opt,name=default_router,json=defaultRouter,proto3,oneof" json:"default_router,omitempty"`
	Sort              *int32                 `protobuf:"varint,7,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
	DataScope         *int32                 `protobuf:"varint,8,opt,name=data_scope,json=dataScope,proto3,oneof" json:"data_scope,omitempty"`
	MenuCheckStrictly *int32                 `protobuf:"varint,9,opt,name=menu_check_strictly,json=menuCheckStrictly,proto3,oneof" json:"menu_check_strictly,omitempty"`
	DeptCheckStrictly *int32                 `protobuf:"varint,10,opt,name=dept_check_strictly,json=deptCheckStrictly,proto3,oneof" json:"dept_check_strictly,omitempty"`
	State             *int32                 `protobuf:"varint,11,opt,name=state,proto3,oneof" json:"state,omitempty"`
	Remark            *string                `protobuf:"bytes,12,opt,name=remark,proto3,oneof" json:"remark,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Role) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDefaultRouter() string {
	if x != nil && x.DefaultRouter != nil {
		return *x.DefaultRouter
	}
	return ""
}

func (x *Role) GetSort() int32 {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return 0
}

func (x *Role) GetDataScope() int32 {
	if x != nil && x.DataScope != nil {
		return *x.DataScope
	}
	return 0
}

func (x *Role) GetMenuCheckStrictly() int32 {
	if x != nil && x.MenuCheckStrictly != nil {
		return *x.MenuCheckStrictly
	}
	return 0
}

func (x *Role) GetDeptCheckStrictly() int32 {
	if x != nil && x.DeptCheckStrictly != nil {
		return *x.DeptCheckStrictly
	}
	return 0
}

func (x *Role) GetState() int32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

func (x *Role) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{1}
}

type CreateRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoleResponse) Reset() {
	*x = CreateRoleResponse{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleResponse) ProtoMessage() {}

func (x *CreateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{2}
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{3}
}

type UpdateRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRoleResponse) Reset() {
	*x = UpdateRoleResponse{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleResponse) ProtoMessage() {}

func (x *UpdateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoleResponse) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{4}
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{5}
}

type DeleteRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoleResponse) Reset() {
	*x = DeleteRoleResponse{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleResponse) ProtoMessage() {}

func (x *DeleteRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleResponse) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{6}
}

type GetRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{7}
}

type GetRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{8}
}

type ListRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoleRequest) Reset() {
	*x = ListRoleRequest{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleRequest) ProtoMessage() {}

func (x *ListRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleRequest.ProtoReflect.Descriptor instead.
func (*ListRoleRequest) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{9}
}

type ListRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoleResponse) Reset() {
	*x = ListRoleResponse{}
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleResponse) ProtoMessage() {}

func (x *ListRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saasdesk_service_v1_role_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleResponse.ProtoReflect.Descriptor instead.
func (*ListRoleResponse) Descriptor() ([]byte, []int) {
	return file_saasdesk_service_v1_role_proto_rawDescGZIP(), []int{10}
}

var File_saasdesk_service_v1_role_proto protoreflect.FileDescriptor

var file_saasdesk_service_v1_role_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x04, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x04, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x33, 0x0a, 0x13, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05,
	0x52, 0x11, 0x6d, 0x65, 0x6e, 0x75, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x6c, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x6c, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x11, 0x64, 0x65, 0x70, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x6c, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x07, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x6c, 0x79, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x6c, 0x79, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xee,
	0x03, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x73,
	0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x26, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64,
	0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x26, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x61,
	0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64,
	0x65, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xd0, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x13, 0x53,
	0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x13, 0x53, 0x61, 0x61, 0x73, 0x64, 0x65, 0x73, 0x6b, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x53, 0x61, 0x61, 0x73, 0x64,
	0x65, 0x73, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x53, 0x61, 0x61,
	0x73, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_saasdesk_service_v1_role_proto_rawDescOnce sync.Once
	file_saasdesk_service_v1_role_proto_rawDescData []byte
)

func file_saasdesk_service_v1_role_proto_rawDescGZIP() []byte {
	file_saasdesk_service_v1_role_proto_rawDescOnce.Do(func() {
		file_saasdesk_service_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_saasdesk_service_v1_role_proto_rawDesc), len(file_saasdesk_service_v1_role_proto_rawDesc)))
	})
	return file_saasdesk_service_v1_role_proto_rawDescData
}

var file_saasdesk_service_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_saasdesk_service_v1_role_proto_goTypes = []any{
	(*Role)(nil),                  // 0: saasdesk.service.v1.Role
	(*CreateRoleRequest)(nil),     // 1: saasdesk.service.v1.CreateRoleRequest
	(*CreateRoleResponse)(nil),    // 2: saasdesk.service.v1.CreateRoleResponse
	(*UpdateRoleRequest)(nil),     // 3: saasdesk.service.v1.UpdateRoleRequest
	(*UpdateRoleResponse)(nil),    // 4: saasdesk.service.v1.UpdateRoleResponse
	(*DeleteRoleRequest)(nil),     // 5: saasdesk.service.v1.DeleteRoleRequest
	(*DeleteRoleResponse)(nil),    // 6: saasdesk.service.v1.DeleteRoleResponse
	(*GetRoleRequest)(nil),        // 7: saasdesk.service.v1.GetRoleRequest
	(*GetRoleResponse)(nil),       // 8: saasdesk.service.v1.GetRoleResponse
	(*ListRoleRequest)(nil),       // 9: saasdesk.service.v1.ListRoleRequest
	(*ListRoleResponse)(nil),      // 10: saasdesk.service.v1.ListRoleResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_saasdesk_service_v1_role_proto_depIdxs = []int32{
	11, // 0: saasdesk.service.v1.Role.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: saasdesk.service.v1.Role.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: saasdesk.service.v1.RoleService.CreateRole:input_type -> saasdesk.service.v1.CreateRoleRequest
	3,  // 3: saasdesk.service.v1.RoleService.UpdateRole:input_type -> saasdesk.service.v1.UpdateRoleRequest
	5,  // 4: saasdesk.service.v1.RoleService.DeleteRole:input_type -> saasdesk.service.v1.DeleteRoleRequest
	7,  // 5: saasdesk.service.v1.RoleService.GetRole:input_type -> saasdesk.service.v1.GetRoleRequest
	9,  // 6: saasdesk.service.v1.RoleService.ListRole:input_type -> saasdesk.service.v1.ListRoleRequest
	2,  // 7: saasdesk.service.v1.RoleService.CreateRole:output_type -> saasdesk.service.v1.CreateRoleResponse
	4,  // 8: saasdesk.service.v1.RoleService.UpdateRole:output_type -> saasdesk.service.v1.UpdateRoleResponse
	6,  // 9: saasdesk.service.v1.RoleService.DeleteRole:output_type -> saasdesk.service.v1.DeleteRoleResponse
	8,  // 10: saasdesk.service.v1.RoleService.GetRole:output_type -> saasdesk.service.v1.GetRoleResponse
	10, // 11: saasdesk.service.v1.RoleService.ListRole:output_type -> saasdesk.service.v1.ListRoleResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_saasdesk_service_v1_role_proto_init() }
func file_saasdesk_service_v1_role_proto_init() {
	if File_saasdesk_service_v1_role_proto != nil {
		return
	}
	file_saasdesk_service_v1_role_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_saasdesk_service_v1_role_proto_rawDesc), len(file_saasdesk_service_v1_role_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saasdesk_service_v1_role_proto_goTypes,
		DependencyIndexes: file_saasdesk_service_v1_role_proto_depIdxs,
		MessageInfos:      file_saasdesk_service_v1_role_proto_msgTypes,
	}.Build()
	File_saasdesk_service_v1_role_proto = out.File
	file_saasdesk_service_v1_role_proto_goTypes = nil
	file_saasdesk_service_v1_role_proto_depIdxs = nil
}
