// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/conf/middleware.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Middleware struct {
	state                protoimpl.MessageState  `protogen:"open.v1"`
	EnableLogging        bool                    `protobuf:"varint,1,opt,name=enable_logging,json=enableLogging,proto3" json:"enable_logging,omitempty"`                        // 日志开关
	EnableRecovery       bool                    `protobuf:"varint,2,opt,name=enable_recovery,json=enableRecovery,proto3" json:"enable_recovery,omitempty"`                     // 异常恢复
	EnableTracing        bool                    `protobuf:"varint,3,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`                        // 链路追踪开关
	EnableValidate       bool                    `protobuf:"varint,4,opt,name=enable_validate,json=enableValidate,proto3" json:"enable_validate,omitempty"`                     // 参数校验开关
	EnableCircuitBreaker bool                    `protobuf:"varint,5,opt,name=enable_circuit_breaker,json=enableCircuitBreaker,proto3" json:"enable_circuit_breaker,omitempty"` // 熔断器
	Limiter              *Middleware_RateLimiter `protobuf:"bytes,6,opt,name=limiter,proto3" json:"limiter,omitempty"`
	Metrics              *Middleware_Metrics     `protobuf:"bytes,7,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Auth                 *Middleware_Auth        `protobuf:"bytes,8,opt,name=auth,proto3" json:"auth,omitempty"`
	Authorizer           *Middleware_Authorizer  `protobuf:"bytes,9,opt,name=authorizer,proto3" json:"authorizer,omitempty"`
	Localize             *Middleware_Localize    `protobuf:"bytes,10,opt,name=localize,proto3" json:"localize,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Middleware) Reset() {
	*x = Middleware{}
	mi := &file_common_conf_middleware_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware) ProtoMessage() {}

func (x *Middleware) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware.ProtoReflect.Descriptor instead.
func (*Middleware) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *Middleware) GetEnableLogging() bool {
	if x != nil {
		return x.EnableLogging
	}
	return false
}

func (x *Middleware) GetEnableRecovery() bool {
	if x != nil {
		return x.EnableRecovery
	}
	return false
}

func (x *Middleware) GetEnableTracing() bool {
	if x != nil {
		return x.EnableTracing
	}
	return false
}

func (x *Middleware) GetEnableValidate() bool {
	if x != nil {
		return x.EnableValidate
	}
	return false
}

func (x *Middleware) GetEnableCircuitBreaker() bool {
	if x != nil {
		return x.EnableCircuitBreaker
	}
	return false
}

func (x *Middleware) GetLimiter() *Middleware_RateLimiter {
	if x != nil {
		return x.Limiter
	}
	return nil
}

func (x *Middleware) GetMetrics() *Middleware_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Middleware) GetAuth() *Middleware_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Middleware) GetAuthorizer() *Middleware_Authorizer {
	if x != nil {
		return x.Authorizer
	}
	return nil
}

func (x *Middleware) GetLocalize() *Middleware_Localize {
	if x != nil {
		return x.Localize
	}
	return nil
}

// JWT校验
type Middleware_Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`                              // JWT签名的算法，支持算法：HS256
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                    // JWT 秘钥
	Header        string                 `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`                              // Header字段：Authentication
	Scheme        string                 `protobuf:"bytes,4,opt,name=scheme,proto3" json:"scheme,omitempty"`                              // token 前缀
	Multipoint    bool                   `protobuf:"varint,5,opt,name=multipoint,proto3" json:"multipoint,omitempty"`                     // 是否多设备登录
	ExpiresTime   *durationpb.Duration   `protobuf:"bytes,6,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"` // 过期时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_Auth) Reset() {
	*x = Middleware_Auth{}
	mi := &file_common_conf_middleware_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Auth) ProtoMessage() {}

func (x *Middleware_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Auth.ProtoReflect.Descriptor instead.
func (*Middleware_Auth) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Middleware_Auth) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Middleware_Auth) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Middleware_Auth) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Middleware_Auth) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Middleware_Auth) GetMultipoint() bool {
	if x != nil {
		return x.Multipoint
	}
	return false
}

func (x *Middleware_Auth) GetExpiresTime() *durationpb.Duration {
	if x != nil {
		return x.ExpiresTime
	}
	return nil
}

// 限流器
type Middleware_RateLimiter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 限流器名字，支持：bbr。
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_RateLimiter) Reset() {
	*x = Middleware_RateLimiter{}
	mi := &file_common_conf_middleware_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_RateLimiter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_RateLimiter) ProtoMessage() {}

func (x *Middleware_RateLimiter) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_RateLimiter.ProtoReflect.Descriptor instead.
func (*Middleware_RateLimiter) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Middleware_RateLimiter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 性能指标
type Middleware_Metrics struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Histogram     bool                   `protobuf:"varint,1,opt,name=histogram,proto3" json:"histogram,omitempty"` // 直方图
	Counter       bool                   `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`     // 计数器
	Gauge         bool                   `protobuf:"varint,3,opt,name=gauge,proto3" json:"gauge,omitempty"`         // 仪表盘
	Summary       bool                   `protobuf:"varint,4,opt,name=summary,proto3" json:"summary,omitempty"`     // 摘要
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_Metrics) Reset() {
	*x = Middleware_Metrics{}
	mi := &file_common_conf_middleware_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Metrics) ProtoMessage() {}

func (x *Middleware_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Metrics.ProtoReflect.Descriptor instead.
func (*Middleware_Metrics) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Middleware_Metrics) GetHistogram() bool {
	if x != nil {
		return x.Histogram
	}
	return false
}

func (x *Middleware_Metrics) GetCounter() bool {
	if x != nil {
		return x.Counter
	}
	return false
}

func (x *Middleware_Metrics) GetGauge() bool {
	if x != nil {
		return x.Gauge
	}
	return false
}

func (x *Middleware_Metrics) GetSummary() bool {
	if x != nil {
		return x.Summary
	}
	return false
}

// 语言包
type Middleware_Localize struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Default       string                 `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_Localize) Reset() {
	*x = Middleware_Localize{}
	mi := &file_common_conf_middleware_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Localize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Localize) ProtoMessage() {}

func (x *Middleware_Localize) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Localize.ProtoReflect.Descriptor instead.
func (*Middleware_Localize) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Middleware_Localize) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

// 鉴权
type Middleware_Authorizer struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Type          string                        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Casbin        *Middleware_Authorizer_Casbin `protobuf:"bytes,2,opt,name=casbin,proto3" json:"casbin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_Authorizer) Reset() {
	*x = Middleware_Authorizer{}
	mi := &file_common_conf_middleware_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Authorizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Authorizer) ProtoMessage() {}

func (x *Middleware_Authorizer) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Authorizer.ProtoReflect.Descriptor instead.
func (*Middleware_Authorizer) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Middleware_Authorizer) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Middleware_Authorizer) GetCasbin() *Middleware_Authorizer_Casbin {
	if x != nil {
		return x.Casbin
	}
	return nil
}

type Middleware_Authorizer_Casbin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ModelPath     string                 `protobuf:"bytes,1,opt,name=model_path,json=modelPath,proto3" json:"model_path,omitempty"`
	PolicyPath    string                 `protobuf:"bytes,2,opt,name=policy_path,json=policyPath,proto3" json:"policy_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Middleware_Authorizer_Casbin) Reset() {
	*x = Middleware_Authorizer_Casbin{}
	mi := &file_common_conf_middleware_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Authorizer_Casbin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Authorizer_Casbin) ProtoMessage() {}

func (x *Middleware_Authorizer_Casbin) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Authorizer_Casbin.ProtoReflect.Descriptor instead.
func (*Middleware_Authorizer_Casbin) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 4, 0}
}

func (x *Middleware_Authorizer_Casbin) GetModelPath() string {
	if x != nil {
		return x.ModelPath
	}
	return ""
}

func (x *Middleware_Authorizer_Casbin) GetPolicyPath() string {
	if x != nil {
		return x.PolicyPath
	}
	return ""
}

var File_common_conf_middleware_proto protoreflect.FileDescriptor

var file_common_conf_middleware_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x08, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x69,
	0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x52, 0x07, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x72, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x35, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x1a, 0xbe, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x21, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x71, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x24, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x1a, 0xa6, 0x01, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x52, 0x06, 0x63, 0x61, 0x73, 0x62, 0x69, 0x6e,
	0x1a, 0x48, 0x0a, 0x06, 0x43, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x81, 0x01, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0f, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x04,
	0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_conf_middleware_proto_rawDescOnce sync.Once
	file_common_conf_middleware_proto_rawDescData []byte
)

func file_common_conf_middleware_proto_rawDescGZIP() []byte {
	file_common_conf_middleware_proto_rawDescOnce.Do(func() {
		file_common_conf_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_conf_middleware_proto_rawDesc), len(file_common_conf_middleware_proto_rawDesc)))
	})
	return file_common_conf_middleware_proto_rawDescData
}

var file_common_conf_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_conf_middleware_proto_goTypes = []any{
	(*Middleware)(nil),                   // 0: conf.Middleware
	(*Middleware_Auth)(nil),              // 1: conf.Middleware.Auth
	(*Middleware_RateLimiter)(nil),       // 2: conf.Middleware.RateLimiter
	(*Middleware_Metrics)(nil),           // 3: conf.Middleware.Metrics
	(*Middleware_Localize)(nil),          // 4: conf.Middleware.Localize
	(*Middleware_Authorizer)(nil),        // 5: conf.Middleware.Authorizer
	(*Middleware_Authorizer_Casbin)(nil), // 6: conf.Middleware.Authorizer.Casbin
	(*durationpb.Duration)(nil),          // 7: google.protobuf.Duration
}
var file_common_conf_middleware_proto_depIdxs = []int32{
	2, // 0: conf.Middleware.limiter:type_name -> conf.Middleware.RateLimiter
	3, // 1: conf.Middleware.metrics:type_name -> conf.Middleware.Metrics
	1, // 2: conf.Middleware.auth:type_name -> conf.Middleware.Auth
	5, // 3: conf.Middleware.authorizer:type_name -> conf.Middleware.Authorizer
	4, // 4: conf.Middleware.localize:type_name -> conf.Middleware.Localize
	7, // 5: conf.Middleware.Auth.expires_time:type_name -> google.protobuf.Duration
	6, // 6: conf.Middleware.Authorizer.casbin:type_name -> conf.Middleware.Authorizer.Casbin
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_common_conf_middleware_proto_init() }
func file_common_conf_middleware_proto_init() {
	if File_common_conf_middleware_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_conf_middleware_proto_rawDesc), len(file_common_conf_middleware_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_middleware_proto_goTypes,
		DependencyIndexes: file_common_conf_middleware_proto_depIdxs,
		MessageInfos:      file_common_conf_middleware_proto_msgTypes,
	}.Build()
	File_common_conf_middleware_proto = out.File
	file_common_conf_middleware_proto_goTypes = nil
	file_common_conf_middleware_proto_depIdxs = nil
}
