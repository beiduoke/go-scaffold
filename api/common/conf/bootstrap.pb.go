// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/conf/bootstrap.proto

package conf

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

// 引导信息
type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server       `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Client   *Client       `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data     *Data         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Trace    *Tracer       `protobuf:"bytes,4,opt,name=trace,proto3" json:"trace,omitempty"`
	Logger   *Logger       `protobuf:"bytes,5,opt,name=logger,proto3" json:"logger,omitempty"`
	Registry *Registry     `protobuf:"bytes,6,opt,name=registry,proto3" json:"registry,omitempty"`
	Config   *RemoteConfig `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	Oss      *OSS          `protobuf:"bytes,8,opt,name=oss,proto3" json:"oss,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_common_conf_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetTrace() *Tracer {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetLogger() *Logger {
	if x != nil {
		return x.Logger
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *Bootstrap) GetConfig() *RemoteConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Bootstrap) GetOss() *OSS {
	if x != nil {
		return x.Oss
	}
	return nil
}

var File_common_conf_bootstrap_proto protoreflect.FileDescriptor

var file_common_conf_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x6f, 0x6e, 0x66, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6f, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x24,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x24,
	0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x03,
	0x6f, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x4f, 0x53, 0x53, 0x52, 0x03, 0x6f, 0x73, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_bootstrap_proto_rawDescOnce sync.Once
	file_common_conf_bootstrap_proto_rawDescData = file_common_conf_bootstrap_proto_rawDesc
)

func file_common_conf_bootstrap_proto_rawDescGZIP() []byte {
	file_common_conf_bootstrap_proto_rawDescOnce.Do(func() {
		file_common_conf_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_bootstrap_proto_rawDescData)
	})
	return file_common_conf_bootstrap_proto_rawDescData
}

var file_common_conf_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_conf_bootstrap_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),    // 0: conf.Bootstrap
	(*Server)(nil),       // 1: conf.Server
	(*Client)(nil),       // 2: conf.Client
	(*Data)(nil),         // 3: conf.Data
	(*Tracer)(nil),       // 4: conf.Tracer
	(*Logger)(nil),       // 5: conf.Logger
	(*Registry)(nil),     // 6: conf.Registry
	(*RemoteConfig)(nil), // 7: conf.RemoteConfig
	(*OSS)(nil),          // 8: conf.OSS
}
var file_common_conf_bootstrap_proto_depIdxs = []int32{
	1, // 0: conf.Bootstrap.server:type_name -> conf.Server
	2, // 1: conf.Bootstrap.client:type_name -> conf.Client
	3, // 2: conf.Bootstrap.data:type_name -> conf.Data
	4, // 3: conf.Bootstrap.trace:type_name -> conf.Tracer
	5, // 4: conf.Bootstrap.logger:type_name -> conf.Logger
	6, // 5: conf.Bootstrap.registry:type_name -> conf.Registry
	7, // 6: conf.Bootstrap.config:type_name -> conf.RemoteConfig
	8, // 7: conf.Bootstrap.oss:type_name -> conf.OSS
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_common_conf_bootstrap_proto_init() }
func file_common_conf_bootstrap_proto_init() {
	if File_common_conf_bootstrap_proto != nil {
		return
	}
	file_common_conf_tracer_proto_init()
	file_common_conf_data_proto_init()
	file_common_conf_server_proto_init()
	file_common_conf_client_proto_init()
	file_common_conf_logger_proto_init()
	file_common_conf_registry_proto_init()
	file_common_conf_oss_proto_init()
	file_common_conf_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_conf_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
			RawDescriptor: file_common_conf_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_bootstrap_proto_goTypes,
		DependencyIndexes: file_common_conf_bootstrap_proto_depIdxs,
		MessageInfos:      file_common_conf_bootstrap_proto_msgTypes,
	}.Build()
	File_common_conf_bootstrap_proto = out.File
	file_common_conf_bootstrap_proto_rawDesc = nil
	file_common_conf_bootstrap_proto_goTypes = nil
	file_common_conf_bootstrap_proto_depIdxs = nil
}
