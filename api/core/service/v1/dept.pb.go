// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: core/service/v1/dept.proto

package v1

import (
	pagination "github.com/beiduoke/go-scaffold/api/common/pagination"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// 部门模块
type Dept struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CreatedAt     *string                `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt     *string                `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	Id            uint32                 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Sort          *int32                 `protobuf:"varint,5,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
	State         *int32                 `protobuf:"varint,6,opt,name=state,proto3,oneof" json:"state,omitempty"`
	Remark        *string                `protobuf:"bytes,7,opt,name=remark,proto3,oneof" json:"remark,omitempty"`
	ParentId      *uint32                `protobuf:"varint,8,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	LeaderId      *uint32                `protobuf:"varint,9,opt,name=leader_id,json=leaderId,proto3,oneof" json:"leader_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dept) Reset() {
	*x = Dept{}
	mi := &file_core_service_v1_dept_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dept) ProtoMessage() {}

func (x *Dept) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dept.ProtoReflect.Descriptor instead.
func (*Dept) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{0}
}

func (x *Dept) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *Dept) GetUpdatedAt() string {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return ""
}

func (x *Dept) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dept) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Dept) GetSort() int32 {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return 0
}

func (x *Dept) GetState() int32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

func (x *Dept) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *Dept) GetParentId() uint32 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Dept) GetLeaderId() uint32 {
	if x != nil && x.LeaderId != nil {
		return *x.LeaderId
	}
	return 0
}

type CreateDeptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dept          *Dept                  `protobuf:"bytes,1,opt,name=dept,proto3" json:"dept,omitempty"`
	OperatorId    uint32                 `protobuf:"varint,2,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDeptRequest) Reset() {
	*x = CreateDeptRequest{}
	mi := &file_core_service_v1_dept_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeptRequest) ProtoMessage() {}

func (x *CreateDeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeptRequest.ProtoReflect.Descriptor instead.
func (*CreateDeptRequest) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeptRequest) GetDept() *Dept {
	if x != nil {
		return x.Dept
	}
	return nil
}

func (x *CreateDeptRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

type CreateDeptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDeptResponse) Reset() {
	*x = CreateDeptResponse{}
	mi := &file_core_service_v1_dept_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeptResponse) ProtoMessage() {}

func (x *CreateDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeptResponse.ProtoReflect.Descriptor instead.
func (*CreateDeptResponse) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{2}
}

type UpdateDeptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Dept          *Dept                  `protobuf:"bytes,2,opt,name=dept,proto3" json:"dept,omitempty"`
	OperatorId    uint32                 `protobuf:"varint,3,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDeptRequest) Reset() {
	*x = UpdateDeptRequest{}
	mi := &file_core_service_v1_dept_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeptRequest) ProtoMessage() {}

func (x *UpdateDeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeptRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeptRequest) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDeptRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDeptRequest) GetDept() *Dept {
	if x != nil {
		return x.Dept
	}
	return nil
}

func (x *UpdateDeptRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

type UpdateDeptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDeptResponse) Reset() {
	*x = UpdateDeptResponse{}
	mi := &file_core_service_v1_dept_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeptResponse) ProtoMessage() {}

func (x *UpdateDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeptResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeptResponse) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{4}
}

type DeleteDeptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId    uint32                 `protobuf:"varint,2,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDeptRequest) Reset() {
	*x = DeleteDeptRequest{}
	mi := &file_core_service_v1_dept_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeptRequest) ProtoMessage() {}

func (x *DeleteDeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeptRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeptRequest) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDeptRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteDeptRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

type DeleteDeptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDeptResponse) Reset() {
	*x = DeleteDeptResponse{}
	mi := &file_core_service_v1_dept_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeptResponse) ProtoMessage() {}

func (x *DeleteDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeptResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeptResponse) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{6}
}

type GetDeptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDeptRequest) Reset() {
	*x = GetDeptRequest{}
	mi := &file_core_service_v1_dept_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeptRequest) ProtoMessage() {}

func (x *GetDeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeptRequest.ProtoReflect.Descriptor instead.
func (*GetDeptRequest) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{7}
}

func (x *GetDeptRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDeptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDeptResponse) Reset() {
	*x = GetDeptResponse{}
	mi := &file_core_service_v1_dept_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeptResponse) ProtoMessage() {}

func (x *GetDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeptResponse.ProtoReflect.Descriptor instead.
func (*GetDeptResponse) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{8}
}

type ListDeptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDeptRequest) Reset() {
	*x = ListDeptRequest{}
	mi := &file_core_service_v1_dept_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeptRequest) ProtoMessage() {}

func (x *ListDeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeptRequest.ProtoReflect.Descriptor instead.
func (*ListDeptRequest) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{9}
}

type ListDeptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Dept                `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDeptResponse) Reset() {
	*x = ListDeptResponse{}
	mi := &file_core_service_v1_dept_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeptResponse) ProtoMessage() {}

func (x *ListDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_v1_dept_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeptResponse.ProtoReflect.Descriptor instead.
func (*ListDeptResponse) Descriptor() ([]byte, []int) {
	return file_core_service_v1_dept_proto_rawDescGZIP(), []int{10}
}

func (x *ListDeptResponse) GetItems() []*Dept {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListDeptResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_core_service_v1_dept_proto protoreflect.FileDescriptor

var file_core_service_v1_dept_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x04, 0x44, 0x65, 0x70, 0x74,
	0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x2a, 0xba, 0x47, 0x27, 0x92, 0x02,
	0x24, 0xe7, 0x8a, 0xb6, 0xe6, 0x80, 0x81, 0x20, 0x31, 0x20, 0xe6, 0xbf, 0x80, 0xe6, 0xb4, 0xbb,
	0x20, 0x32, 0x20, 0xe6, 0x9c, 0xaa, 0xe6, 0xbf, 0x80, 0xe6, 0xb4, 0xbb, 0x20, 0x33, 0x20, 0xe7,
	0xa6, 0x81, 0xe7, 0x94, 0xa8, 0x48, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x05, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x06, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x04, 0x64,
	0x65, 0x70, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x44, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x9f, 0x03, 0x0a,
	0x0b, 0x44, 0x65, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x74,
	0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb6,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66,
	0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53,
	0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_core_service_v1_dept_proto_rawDescOnce sync.Once
	file_core_service_v1_dept_proto_rawDescData []byte
)

func file_core_service_v1_dept_proto_rawDescGZIP() []byte {
	file_core_service_v1_dept_proto_rawDescOnce.Do(func() {
		file_core_service_v1_dept_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_service_v1_dept_proto_rawDesc), len(file_core_service_v1_dept_proto_rawDesc)))
	})
	return file_core_service_v1_dept_proto_rawDescData
}

var file_core_service_v1_dept_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_core_service_v1_dept_proto_goTypes = []any{
	(*Dept)(nil),                     // 0: core.service.v1.Dept
	(*CreateDeptRequest)(nil),        // 1: core.service.v1.CreateDeptRequest
	(*CreateDeptResponse)(nil),       // 2: core.service.v1.CreateDeptResponse
	(*UpdateDeptRequest)(nil),        // 3: core.service.v1.UpdateDeptRequest
	(*UpdateDeptResponse)(nil),       // 4: core.service.v1.UpdateDeptResponse
	(*DeleteDeptRequest)(nil),        // 5: core.service.v1.DeleteDeptRequest
	(*DeleteDeptResponse)(nil),       // 6: core.service.v1.DeleteDeptResponse
	(*GetDeptRequest)(nil),           // 7: core.service.v1.GetDeptRequest
	(*GetDeptResponse)(nil),          // 8: core.service.v1.GetDeptResponse
	(*ListDeptRequest)(nil),          // 9: core.service.v1.ListDeptRequest
	(*ListDeptResponse)(nil),         // 10: core.service.v1.ListDeptResponse
	(*pagination.PagingRequest)(nil), // 11: pagination.PagingRequest
}
var file_core_service_v1_dept_proto_depIdxs = []int32{
	0,  // 0: core.service.v1.CreateDeptRequest.dept:type_name -> core.service.v1.Dept
	0,  // 1: core.service.v1.UpdateDeptRequest.dept:type_name -> core.service.v1.Dept
	0,  // 2: core.service.v1.ListDeptResponse.items:type_name -> core.service.v1.Dept
	1,  // 3: core.service.v1.DeptService.CreateDept:input_type -> core.service.v1.CreateDeptRequest
	3,  // 4: core.service.v1.DeptService.UpdateDept:input_type -> core.service.v1.UpdateDeptRequest
	5,  // 5: core.service.v1.DeptService.DeleteDept:input_type -> core.service.v1.DeleteDeptRequest
	7,  // 6: core.service.v1.DeptService.GetDept:input_type -> core.service.v1.GetDeptRequest
	11, // 7: core.service.v1.DeptService.ListDept:input_type -> pagination.PagingRequest
	2,  // 8: core.service.v1.DeptService.CreateDept:output_type -> core.service.v1.CreateDeptResponse
	4,  // 9: core.service.v1.DeptService.UpdateDept:output_type -> core.service.v1.UpdateDeptResponse
	6,  // 10: core.service.v1.DeptService.DeleteDept:output_type -> core.service.v1.DeleteDeptResponse
	0,  // 11: core.service.v1.DeptService.GetDept:output_type -> core.service.v1.Dept
	10, // 12: core.service.v1.DeptService.ListDept:output_type -> core.service.v1.ListDeptResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_core_service_v1_dept_proto_init() }
func file_core_service_v1_dept_proto_init() {
	if File_core_service_v1_dept_proto != nil {
		return
	}
	file_core_service_v1_dept_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_service_v1_dept_proto_rawDesc), len(file_core_service_v1_dept_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_service_v1_dept_proto_goTypes,
		DependencyIndexes: file_core_service_v1_dept_proto_depIdxs,
		MessageInfos:      file_core_service_v1_dept_proto_msgTypes,
	}.Build()
	File_core_service_v1_dept_proto = out.File
	file_core_service_v1_dept_proto_goTypes = nil
	file_core_service_v1_dept_proto_depIdxs = nil
}
