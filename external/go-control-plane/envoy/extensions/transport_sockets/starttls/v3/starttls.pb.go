// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: envoy/extensions/transport_sockets/starttls/v3/starttls.proto

package envoy_extensions_transport_sockets_starttls_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/raw_buffer/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for a downstream StartTls transport socket.
// StartTls transport socket wraps two sockets:
// * raw_buffer socket which is used at the beginning of the session
// * TLS socket used when a protocol negotiates a switch to encrypted traffic.
type StartTlsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (optional) Configuration for clear-text socket used at the beginning of the session.
	CleartextSocketConfig *v3.RawBuffer `protobuf:"bytes,1,opt,name=cleartext_socket_config,json=cleartextSocketConfig,proto3" json:"cleartext_socket_config,omitempty"`
	// Configuration for a downstream TLS socket.
	TlsSocketConfig *v31.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_socket_config,json=tlsSocketConfig,proto3" json:"tls_socket_config,omitempty"`
}

func (x *StartTlsConfig) Reset() {
	*x = StartTlsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTlsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTlsConfig) ProtoMessage() {}

func (x *StartTlsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTlsConfig.ProtoReflect.Descriptor instead.
func (*StartTlsConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescGZIP(), []int{0}
}

func (x *StartTlsConfig) GetCleartextSocketConfig() *v3.RawBuffer {
	if x != nil {
		return x.CleartextSocketConfig
	}
	return nil
}

func (x *StartTlsConfig) GetTlsSocketConfig() *v31.DownstreamTlsContext {
	if x != nil {
		return x.TlsSocketConfig
	}
	return nil
}

// Configuration for an upstream StartTls transport socket.
// StartTls transport socket wraps two sockets:
// * raw_buffer socket which is used at the beginning of the session
// * TLS socket used when a protocol negotiates a switch to encrypted traffic.
type UpstreamStartTlsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (optional) Configuration for clear-text socket used at the beginning of the session.
	CleartextSocketConfig *v3.RawBuffer `protobuf:"bytes,1,opt,name=cleartext_socket_config,json=cleartextSocketConfig,proto3" json:"cleartext_socket_config,omitempty"`
	// Configuration for an upstream TLS socket.
	TlsSocketConfig *v31.UpstreamTlsContext `protobuf:"bytes,2,opt,name=tls_socket_config,json=tlsSocketConfig,proto3" json:"tls_socket_config,omitempty"`
}

func (x *UpstreamStartTlsConfig) Reset() {
	*x = UpstreamStartTlsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamStartTlsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamStartTlsConfig) ProtoMessage() {}

func (x *UpstreamStartTlsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamStartTlsConfig.ProtoReflect.Descriptor instead.
func (*UpstreamStartTlsConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescGZIP(), []int{1}
}

func (x *UpstreamStartTlsConfig) GetCleartextSocketConfig() *v3.RawBuffer {
	if x != nil {
		return x.CleartextSocketConfig
	}
	return nil
}

func (x *UpstreamStartTlsConfig) GetTlsSocketConfig() *v31.UpstreamTlsContext {
	if x != nil {
		return x.TlsSocketConfig
	}
	return nil
}

var File_envoy_extensions_transport_sockets_starttls_v3_starttls_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x1a,
	0x41, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2f, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x73, 0x0a, 0x17, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x77, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x15, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x74, 0x65, 0x78, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x75, 0x0a, 0x11, 0x74, 0x6c, 0x73, 0x5f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x74,
	0x6c, 0x73, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x82,
	0x02, 0x0a, 0x16, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x73, 0x0a, 0x17, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x72, 0x61, 0x77, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61,
	0x77, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x52, 0x15, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x73,
	0x0a, 0x11, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74,
	0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0f, 0x74, 0x6c, 0x73, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x57, 0x0a, 0x3c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x6c, 0x73,
	0x2e, 0x76, 0x33, 0x42, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x74, 0x6c, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescData = file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_goTypes = []interface{}{
	(*StartTlsConfig)(nil),           // 0: envoy.extensions.transport_sockets.starttls.v3.StartTlsConfig
	(*UpstreamStartTlsConfig)(nil),   // 1: envoy.extensions.transport_sockets.starttls.v3.UpstreamStartTlsConfig
	(*v3.RawBuffer)(nil),             // 2: envoy.extensions.transport_sockets.raw_buffer.v3.RawBuffer
	(*v31.DownstreamTlsContext)(nil), // 3: envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
	(*v31.UpstreamTlsContext)(nil),   // 4: envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
}
var file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.transport_sockets.starttls.v3.StartTlsConfig.cleartext_socket_config:type_name -> envoy.extensions.transport_sockets.raw_buffer.v3.RawBuffer
	3, // 1: envoy.extensions.transport_sockets.starttls.v3.StartTlsConfig.tls_socket_config:type_name -> envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
	2, // 2: envoy.extensions.transport_sockets.starttls.v3.UpstreamStartTlsConfig.cleartext_socket_config:type_name -> envoy.extensions.transport_sockets.raw_buffer.v3.RawBuffer
	4, // 3: envoy.extensions.transport_sockets.starttls.v3.UpstreamStartTlsConfig.tls_socket_config:type_name -> envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_init() }
func file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_init() {
	if File_envoy_extensions_transport_sockets_starttls_v3_starttls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTlsConfig); i {
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
		file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamStartTlsConfig); i {
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
			RawDescriptor: file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_starttls_v3_starttls_proto = out.File
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_starttls_v3_starttls_proto_depIdxs = nil
}
