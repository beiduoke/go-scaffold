// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/conf/server.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 服务器
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http      *Server_HTTP      `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`           // HTTP服务
	Grpc      *Server_GRPC      `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`           // gRPC服务
	Websocket *Server_Websocket `protobuf:"bytes,3,opt,name=websocket,proto3" json:"websocket,omitempty"` // Websocket服务
	Mqtt      *Server_Mqtt      `protobuf:"bytes,4,opt,name=mqtt,proto3" json:"mqtt,omitempty"`           // MQTT服务
	Kafka     *Server_Kafka     `protobuf:"bytes,5,opt,name=kafka,proto3" json:"kafka,omitempty"`         // Kafka服务
	Rabbitmq  *Server_RabbitMQ  `protobuf:"bytes,6,opt,name=rabbitmq,proto3" json:"rabbitmq,omitempty"`   // RabbitMQ服务
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Server) GetWebsocket() *Server_Websocket {
	if x != nil {
		return x.Websocket
	}
	return nil
}

func (x *Server) GetMqtt() *Server_Mqtt {
	if x != nil {
		return x.Mqtt
	}
	return nil
}

func (x *Server) GetKafka() *Server_Kafka {
	if x != nil {
		return x.Kafka
	}
	return nil
}

func (x *Server) GetRabbitmq() *Server_RabbitMQ {
	if x != nil {
		return x.Rabbitmq
	}
	return nil
}

// HTTP
type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`                                   // 网络
	Addr          string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                                         // 服务监听地址
	Timeout       *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`                                   // 超时时间
	Cors          *Server_HTTP_CORS    `protobuf:"bytes,4,opt,name=cors,proto3" json:"cors,omitempty"`                                         // 服务监听地址
	Middleware    *Middleware          `protobuf:"bytes,5,opt,name=middleware,proto3" json:"middleware,omitempty"`                             // 中间件
	EnableSwagger bool                 `protobuf:"varint,6,opt,name=enable_swagger,json=enableSwagger,proto3" json:"enable_swagger,omitempty"` // 启用SwaggerUI
	EnablePprof   bool                 `protobuf:"varint,7,opt,name=enable_pprof,json=enablePprof,proto3" json:"enable_pprof,omitempty"`       // 启用pprof
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetCors() *Server_HTTP_CORS {
	if x != nil {
		return x.Cors
	}
	return nil
}

func (x *Server_HTTP) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Server_HTTP) GetEnableSwagger() bool {
	if x != nil {
		return x.EnableSwagger
	}
	return false
}

func (x *Server_HTTP) GetEnablePprof() bool {
	if x != nil {
		return x.EnablePprof
	}
	return false
}

// gPRC
type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network    string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"` // 网络
	Addr       string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`       // 服务监听地址
	Timeout    *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"` // 超时时间
	Middleware *Middleware          `protobuf:"bytes,4,opt,name=middleware,proto3" json:"middleware,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_GRPC) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

// Websocket
type Server_Websocket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"` // 网络样式：http、https
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`       // 服务监听地址
	Path    string               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`       // 路径
	Codec   string               `protobuf:"bytes,4,opt,name=codec,proto3" json:"codec,omitempty"`     // 编解码器: json,xml,yaml...
	Timeout *durationpb.Duration `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"` // 超时时间
}

func (x *Server_Websocket) Reset() {
	*x = Server_Websocket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Websocket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Websocket) ProtoMessage() {}

func (x *Server_Websocket) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Websocket.ProtoReflect.Descriptor instead.
func (*Server_Websocket) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Server_Websocket) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_Websocket) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_Websocket) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Server_Websocket) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

func (x *Server_Websocket) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// MQTT
type Server_Mqtt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"` // 对端网络地址
}

func (x *Server_Mqtt) Reset() {
	*x = Server_Mqtt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Mqtt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Mqtt) ProtoMessage() {}

func (x *Server_Mqtt) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Mqtt.ProtoReflect.Descriptor instead.
func (*Server_Mqtt) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Server_Mqtt) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

// Kafka
type Server_Kafka struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"` // 对端网络地址
}

func (x *Server_Kafka) Reset() {
	*x = Server_Kafka{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Kafka) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Kafka) ProtoMessage() {}

func (x *Server_Kafka) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Kafka.ProtoReflect.Descriptor instead.
func (*Server_Kafka) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Server_Kafka) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

// RabbitMQ
type Server_RabbitMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"` // 对端网络地址
}

func (x *Server_RabbitMQ) Reset() {
	*x = Server_RabbitMQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_RabbitMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_RabbitMQ) ProtoMessage() {}

func (x *Server_RabbitMQ) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_RabbitMQ.ProtoReflect.Descriptor instead.
func (*Server_RabbitMQ) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Server_RabbitMQ) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type Server_HTTP_CORS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"` //
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"` //
	Origins []string `protobuf:"bytes,3,rep,name=origins,proto3" json:"origins,omitempty"` //
}

func (x *Server_HTTP_CORS) Reset() {
	*x = Server_HTTP_CORS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP_CORS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP_CORS) ProtoMessage() {}

func (x *Server_HTTP_CORS) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP_CORS.ProtoReflect.Descriptor instead.
func (*Server_HTTP_CORS) Descriptor() ([]byte, []int) {
	return file_common_conf_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Server_HTTP_CORS) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Server_HTTP_CORS) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Server_HTTP_CORS) GetOrigins() []string {
	if x != nil {
		return x.Origins
	}
	return nil
}

var File_common_conf_server_proto protoreflect.FileDescriptor

var file_common_conf_server_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x08, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x12, 0x25, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50,
	0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x34, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x0a,
	0x04, 0x6d, 0x71, 0x74, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x52, 0x04,
	0x6d, 0x71, 0x74, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x12, 0x31,
	0x0a, 0x08, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52,
	0x61, 0x62, 0x62, 0x69, 0x74, 0x4d, 0x51, 0x52, 0x08, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d,
	0x71, 0x1a, 0xe7, 0x02, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a,
	0x04, 0x63, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x2e, 0x43,
	0x4f, 0x52, 0x53, 0x52, 0x04, 0x63, 0x6f, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52,
	0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x70, 0x72,
	0x6f, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x70, 0x72, 0x6f, 0x66, 0x1a, 0x54, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x53, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x9b, 0x01, 0x0a, 0x04,
	0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x98, 0x01, 0x0a, 0x09, 0x57, 0x65,
	0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64,
	0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x1a, 0x1a, 0x0a, 0x04, 0x4d, 0x71, 0x74, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x1a, 0x1d, 0x0a, 0x05, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x1a,
	0x20, 0x0a, 0x08, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x4d, 0x51, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x64, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72,
	0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66,
	0x66, 0x6f, 0x6c, 0x64, 0x2d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_server_proto_rawDescOnce sync.Once
	file_common_conf_server_proto_rawDescData = file_common_conf_server_proto_rawDesc
)

func file_common_conf_server_proto_rawDescGZIP() []byte {
	file_common_conf_server_proto_rawDescOnce.Do(func() {
		file_common_conf_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_server_proto_rawDescData)
	})
	return file_common_conf_server_proto_rawDescData
}

var file_common_conf_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_conf_server_proto_goTypes = []interface{}{
	(*Server)(nil),              // 0: conf.Server
	(*Server_HTTP)(nil),         // 1: conf.Server.HTTP
	(*Server_GRPC)(nil),         // 2: conf.Server.GRPC
	(*Server_Websocket)(nil),    // 3: conf.Server.Websocket
	(*Server_Mqtt)(nil),         // 4: conf.Server.Mqtt
	(*Server_Kafka)(nil),        // 5: conf.Server.Kafka
	(*Server_RabbitMQ)(nil),     // 6: conf.Server.RabbitMQ
	(*Server_HTTP_CORS)(nil),    // 7: conf.Server.HTTP.CORS
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
	(*Middleware)(nil),          // 9: conf.Middleware
}
var file_common_conf_server_proto_depIdxs = []int32{
	1,  // 0: conf.Server.http:type_name -> conf.Server.HTTP
	2,  // 1: conf.Server.grpc:type_name -> conf.Server.GRPC
	3,  // 2: conf.Server.websocket:type_name -> conf.Server.Websocket
	4,  // 3: conf.Server.mqtt:type_name -> conf.Server.Mqtt
	5,  // 4: conf.Server.kafka:type_name -> conf.Server.Kafka
	6,  // 5: conf.Server.rabbitmq:type_name -> conf.Server.RabbitMQ
	8,  // 6: conf.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	7,  // 7: conf.Server.HTTP.cors:type_name -> conf.Server.HTTP.CORS
	9,  // 8: conf.Server.HTTP.middleware:type_name -> conf.Middleware
	8,  // 9: conf.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	9,  // 10: conf.Server.GRPC.middleware:type_name -> conf.Middleware
	8,  // 11: conf.Server.Websocket.timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_common_conf_server_proto_init() }
func file_common_conf_server_proto_init() {
	if File_common_conf_server_proto != nil {
		return
	}
	file_common_conf_middleware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_conf_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_common_conf_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_common_conf_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_common_conf_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Websocket); i {
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
		file_common_conf_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Mqtt); i {
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
		file_common_conf_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Kafka); i {
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
		file_common_conf_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_RabbitMQ); i {
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
		file_common_conf_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP_CORS); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_conf_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_server_proto_goTypes,
		DependencyIndexes: file_common_conf_server_proto_depIdxs,
		MessageInfos:      file_common_conf_server_proto_msgTypes,
	}.Build()
	File_common_conf_server_proto = out.File
	file_common_conf_server_proto_rawDesc = nil
	file_common_conf_server_proto_goTypes = nil
	file_common_conf_server_proto_depIdxs = nil
}
