// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/conf/logger.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// 日志
type Logger struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Zap           *Logger_Zap            `protobuf:"bytes,2,opt,name=zap,proto3" json:"zap,omitempty"`
	Logrus        *Logger_Logrus         `protobuf:"bytes,3,opt,name=logrus,proto3" json:"logrus,omitempty"`
	Fluent        *Logger_Fluent         `protobuf:"bytes,4,opt,name=fluent,proto3" json:"fluent,omitempty"`
	Aliyun        *Logger_Aliyun         `protobuf:"bytes,5,opt,name=aliyun,proto3" json:"aliyun,omitempty"`
	Tencent       *Logger_Tencent        `protobuf:"bytes,6,opt,name=tencent,proto3" json:"tencent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logger) Reset() {
	*x = Logger{}
	mi := &file_common_conf_logger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Logger) GetZap() *Logger_Zap {
	if x != nil {
		return x.Zap
	}
	return nil
}

func (x *Logger) GetLogrus() *Logger_Logrus {
	if x != nil {
		return x.Logrus
	}
	return nil
}

func (x *Logger) GetFluent() *Logger_Fluent {
	if x != nil {
		return x.Fluent
	}
	return nil
}

func (x *Logger) GetAliyun() *Logger_Aliyun {
	if x != nil {
		return x.Aliyun
	}
	return nil
}

func (x *Logger) GetTencent() *Logger_Tencent {
	if x != nil {
		return x.Tencent
	}
	return nil
}

// Zap
type Logger_Zap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`                        //
	Level         string                 `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`                              //
	MaxSize       int32                  `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`          //
	MaxAge        int32                  `protobuf:"varint,4,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`             //
	MaxBackups    int32                  `protobuf:"varint,5,opt,name=max_backups,json=maxBackups,proto3" json:"max_backups,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logger_Zap) Reset() {
	*x = Logger_Zap{}
	mi := &file_common_conf_logger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger_Zap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Zap) ProtoMessage() {}

func (x *Logger_Zap) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Zap.ProtoReflect.Descriptor instead.
func (*Logger_Zap) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Logger_Zap) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Logger_Zap) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Zap) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Logger_Zap) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Logger_Zap) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

// logrus
type Logger_Logrus struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Level            string                 `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`                                                // 日志等级
	Formatter        string                 `protobuf:"bytes,2,opt,name=formatter,proto3" json:"formatter,omitempty"`                                        // 输出格式：text, json.
	TimestampFormat  string                 `protobuf:"bytes,3,opt,name=timestamp_format,json=timestampFormat,proto3" json:"timestamp_format,omitempty"`     // 定义时间戳格式，例如："2006-01-02 15:04:05"
	DisableColors    bool                   `protobuf:"varint,4,opt,name=disable_colors,json=disableColors,proto3" json:"disable_colors,omitempty"`          // 不需要彩色日志
	DisableTimestamp bool                   `protobuf:"varint,5,opt,name=disable_timestamp,json=disableTimestamp,proto3" json:"disable_timestamp,omitempty"` // 不需要时间戳
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Logger_Logrus) Reset() {
	*x = Logger_Logrus{}
	mi := &file_common_conf_logger_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger_Logrus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Logrus) ProtoMessage() {}

func (x *Logger_Logrus) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Logrus.ProtoReflect.Descriptor instead.
func (*Logger_Logrus) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Logger_Logrus) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Logger_Logrus) GetFormatter() string {
	if x != nil {
		return x.Formatter
	}
	return ""
}

func (x *Logger_Logrus) GetTimestampFormat() string {
	if x != nil {
		return x.TimestampFormat
	}
	return ""
}

func (x *Logger_Logrus) GetDisableColors() bool {
	if x != nil {
		return x.DisableColors
	}
	return false
}

func (x *Logger_Logrus) GetDisableTimestamp() bool {
	if x != nil {
		return x.DisableTimestamp
	}
	return false
}

// Fluent
type Logger_Fluent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"` // 公网接入地址
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logger_Fluent) Reset() {
	*x = Logger_Fluent{}
	mi := &file_common_conf_logger_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger_Fluent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Fluent) ProtoMessage() {}

func (x *Logger_Fluent) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Fluent.ProtoReflect.Descriptor instead.
func (*Logger_Fluent) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Logger_Fluent) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// 阿里云
type Logger_Aliyun struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`                             // 公网接入地址
	Project       string                 `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`                               //
	AccessKey     string                 `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`          // 访问密钥ID
	AccessSecret  string                 `protobuf:"bytes,4,opt,name=access_secret,json=accessSecret,proto3" json:"access_secret,omitempty"` // 访问密钥
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logger_Aliyun) Reset() {
	*x = Logger_Aliyun{}
	mi := &file_common_conf_logger_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger_Aliyun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Aliyun) ProtoMessage() {}

func (x *Logger_Aliyun) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Aliyun.ProtoReflect.Descriptor instead.
func (*Logger_Aliyun) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Logger_Aliyun) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Logger_Aliyun) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Logger_Aliyun) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Logger_Aliyun) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

// 腾讯
type Logger_Tencent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`                             // 公网接入地址
	TopicId       string                 `protobuf:"bytes,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`                //
	AccessKey     string                 `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`          // 访问密钥ID
	AccessSecret  string                 `protobuf:"bytes,4,opt,name=access_secret,json=accessSecret,proto3" json:"access_secret,omitempty"` // 访问密钥
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logger_Tencent) Reset() {
	*x = Logger_Tencent{}
	mi := &file_common_conf_logger_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logger_Tencent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger_Tencent) ProtoMessage() {}

func (x *Logger_Tencent) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_logger_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger_Tencent.ProtoReflect.Descriptor instead.
func (*Logger_Tencent) Descriptor() ([]byte, []int) {
	return file_common_conf_logger_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Logger_Tencent) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Logger_Tencent) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

func (x *Logger_Tencent) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Logger_Tencent) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

var File_common_conf_logger_proto protoreflect.FileDescriptor

var file_common_conf_logger_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66,
	0x22, 0xf6, 0x06, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x03, 0x7a, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x5a, 0x61, 0x70, 0x52, 0x03,
	0x7a, 0x61, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x72, 0x75, 0x73, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x72, 0x75, 0x73,
	0x12, 0x2b, 0x0a, 0x06, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x46,
	0x6c, 0x75, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a,
	0x06, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x69, 0x79,
	0x75, 0x6e, 0x52, 0x06, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x1a, 0x8c, 0x01, 0x0a, 0x03, 0x5a,
	0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x1a, 0xbb, 0x01, 0x0a, 0x06, 0x4c, 0x6f,
	0x67, 0x72, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x24, 0x0a, 0x06, 0x46, 0x6c, 0x75, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x82, 0x01,
	0x0a, 0x06, 0x41, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x1a, 0x84, 0x01, 0x0a, 0x07, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x7d, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0b, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61,
	0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02,
	0x10, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_conf_logger_proto_rawDescOnce sync.Once
	file_common_conf_logger_proto_rawDescData []byte
)

func file_common_conf_logger_proto_rawDescGZIP() []byte {
	file_common_conf_logger_proto_rawDescOnce.Do(func() {
		file_common_conf_logger_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_conf_logger_proto_rawDesc), len(file_common_conf_logger_proto_rawDesc)))
	})
	return file_common_conf_logger_proto_rawDescData
}

var file_common_conf_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_conf_logger_proto_goTypes = []any{
	(*Logger)(nil),         // 0: conf.Logger
	(*Logger_Zap)(nil),     // 1: conf.Logger.Zap
	(*Logger_Logrus)(nil),  // 2: conf.Logger.Logrus
	(*Logger_Fluent)(nil),  // 3: conf.Logger.Fluent
	(*Logger_Aliyun)(nil),  // 4: conf.Logger.Aliyun
	(*Logger_Tencent)(nil), // 5: conf.Logger.Tencent
}
var file_common_conf_logger_proto_depIdxs = []int32{
	1, // 0: conf.Logger.zap:type_name -> conf.Logger.Zap
	2, // 1: conf.Logger.logrus:type_name -> conf.Logger.Logrus
	3, // 2: conf.Logger.fluent:type_name -> conf.Logger.Fluent
	4, // 3: conf.Logger.aliyun:type_name -> conf.Logger.Aliyun
	5, // 4: conf.Logger.tencent:type_name -> conf.Logger.Tencent
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_conf_logger_proto_init() }
func file_common_conf_logger_proto_init() {
	if File_common_conf_logger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_conf_logger_proto_rawDesc), len(file_common_conf_logger_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_logger_proto_goTypes,
		DependencyIndexes: file_common_conf_logger_proto_depIdxs,
		MessageInfos:      file_common_conf_logger_proto_msgTypes,
	}.Build()
	File_common_conf_logger_proto = out.File
	file_common_conf_logger_proto_goTypes = nil
	file_common_conf_logger_proto_depIdxs = nil
}
