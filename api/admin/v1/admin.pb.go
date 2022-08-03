// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.0
// source: admin/v1/admin.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint64  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	UserName   *string  `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3,oneof" json:"user_name,omitempty"`
	Password   *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
	NickName   *string  `protobuf:"bytes,4,opt,name=nick_name,json=nickName,proto3,oneof" json:"nick_name,omitempty"`
	Email      *string  `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Authoritys []string `protobuf:"bytes,6,rep,name=authoritys,proto3" json:"authoritys,omitempty"`
	// Deprecated: Do not use.
	Token *string `protobuf:"bytes,7,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *User) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *User) GetAuthoritys() []string {
	if x != nil {
		return x.Authoritys
	}
	return nil
}

// Deprecated: Do not use.
func (x *User) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*User `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserReply) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListUserReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 登录
type NamePasswordLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *NamePasswordLoginReq) Reset() {
	*x = NamePasswordLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamePasswordLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamePasswordLoginReq) ProtoMessage() {}

func (x *NamePasswordLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamePasswordLoginReq.ProtoReflect.Descriptor instead.
func (*NamePasswordLoginReq) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *NamePasswordLoginReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamePasswordLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 回应 - 登录
type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReply) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

// 请求 - 登出
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 回应 - 登出
type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type WebsocketProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *WebsocketProto) Reset() {
	*x = WebsocketProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsocketProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsocketProto) ProtoMessage() {}

func (x *WebsocketProto) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsocketProto.ProtoReflect.Descriptor instead.
func (*WebsocketProto) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{9}
}

func (x *WebsocketProto) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *WebsocketProto) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type WebsocketMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Sender    string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *WebsocketMessage) Reset() {
	*x = WebsocketMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsocketMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsocketMessage) ProtoMessage() {}

func (x *WebsocketMessage) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsocketMessage.ProtoReflect.Descriptor instead.
func (*WebsocketMessage) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{10}
}

func (x *WebsocketMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *WebsocketMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *WebsocketMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_admin_v1_admin_proto protoreflect.FileDescriptor

var file_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x9e,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x48,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x4f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x5c, 0x0a, 0x14, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x0a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x06, 0x18, 0x19, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5f,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x1b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x0b, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a,
	0x0e, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x62, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xa1, 0x03, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x5a, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7a,
	0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a,
	0x01, 0x2a, 0x12, 0x58, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x5a, 0x0a, 0x17,
	0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_admin_proto_rawDescOnce sync.Once
	file_admin_v1_admin_proto_rawDescData = file_admin_v1_admin_proto_rawDesc
)

func file_admin_v1_admin_proto_rawDescGZIP() []byte {
	file_admin_v1_admin_proto_rawDescOnce.Do(func() {
		file_admin_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_admin_proto_rawDescData)
	})
	return file_admin_v1_admin_proto_rawDescData
}

var file_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_admin_v1_admin_proto_goTypes = []interface{}{
	(*Content)(nil),               // 0: api.admin.v1.Content
	(*User)(nil),                  // 1: api.admin.v1.User
	(*ListUserReply)(nil),         // 2: api.admin.v1.ListUserReply
	(*NamePasswordLoginReq)(nil),  // 3: api.admin.v1.NamePasswordLoginReq
	(*LoginReply)(nil),            // 4: api.admin.v1.LoginReply
	(*LogoutReq)(nil),             // 5: api.admin.v1.LogoutReq
	(*LogoutReply)(nil),           // 6: api.admin.v1.LogoutReply
	(*RegisterReq)(nil),           // 7: api.admin.v1.RegisterReq
	(*RegisterReply)(nil),         // 8: api.admin.v1.RegisterReply
	(*WebsocketProto)(nil),        // 9: api.admin.v1.WebsocketProto
	(*WebsocketMessage)(nil),      // 10: api.admin.v1.WebsocketMessage
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_admin_v1_admin_proto_depIdxs = []int32{
	1,  // 0: api.admin.v1.ListUserReply.items:type_name -> api.admin.v1.User
	11, // 1: api.admin.v1.LoginReply.expire_time:type_name -> google.protobuf.Timestamp
	12, // 2: api.admin.v1.Admin.Logout:input_type -> google.protobuf.Empty
	3,  // 3: api.admin.v1.Admin.NamePasswordLogin:input_type -> api.admin.v1.NamePasswordLoginReq
	7,  // 4: api.admin.v1.Admin.Register:input_type -> api.admin.v1.RegisterReq
	12, // 5: api.admin.v1.Admin.ListUser:input_type -> google.protobuf.Empty
	12, // 6: api.admin.v1.Admin.Logout:output_type -> google.protobuf.Empty
	4,  // 7: api.admin.v1.Admin.NamePasswordLogin:output_type -> api.admin.v1.LoginReply
	8,  // 8: api.admin.v1.Admin.Register:output_type -> api.admin.v1.RegisterReply
	2,  // 9: api.admin.v1.Admin.ListUser:output_type -> api.admin.v1.ListUserReply
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_admin_v1_admin_proto_init() }
func file_admin_v1_admin_proto_init() {
	if File_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_admin_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
		file_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamePasswordLoginReq); i {
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
		file_admin_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_admin_v1_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_admin_v1_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
		file_admin_v1_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_admin_v1_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_admin_v1_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsocketProto); i {
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
		file_admin_v1_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsocketMessage); i {
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
	file_admin_v1_admin_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_admin_proto_goTypes,
		DependencyIndexes: file_admin_v1_admin_proto_depIdxs,
		MessageInfos:      file_admin_v1_admin_proto_msgTypes,
	}.Build()
	File_admin_v1_admin_proto = out.File
	file_admin_v1_admin_proto_rawDesc = nil
	file_admin_v1_admin_proto_goTypes = nil
	file_admin_v1_admin_proto_depIdxs = nil
}
