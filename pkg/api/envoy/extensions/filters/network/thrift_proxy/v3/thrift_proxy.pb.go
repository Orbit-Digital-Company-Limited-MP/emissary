// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto

package thrift_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/annotations"
	v31 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/accesslog/v3"
	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any1 "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Thrift transport types supported by Envoy.
type TransportType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which transport to use.
	// For upstream connections, the Thrift proxy will use same transport as the downstream
	// connection.
	TransportType_AUTO_TRANSPORT TransportType = 0
	// The Thrift proxy will use the Thrift framed transport.
	TransportType_FRAMED TransportType = 1
	// The Thrift proxy will use the Thrift unframed transport.
	TransportType_UNFRAMED TransportType = 2
	// The Thrift proxy will assume the client is using the Thrift header transport.
	TransportType_HEADER TransportType = 3
)

// Enum value maps for TransportType.
var (
	TransportType_name = map[int32]string{
		0: "AUTO_TRANSPORT",
		1: "FRAMED",
		2: "UNFRAMED",
		3: "HEADER",
	}
	TransportType_value = map[string]int32{
		"AUTO_TRANSPORT": 0,
		"FRAMED":         1,
		"UNFRAMED":       2,
		"HEADER":         3,
	}
)

func (x TransportType) Enum() *TransportType {
	p := new(TransportType)
	*p = x
	return p
}

func (x TransportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes[0].Descriptor()
}

func (TransportType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes[0]
}

func (x TransportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportType.Descriptor instead.
func (TransportType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{0}
}

// Thrift Protocol types supported by Envoy.
type ProtocolType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which protocol to use.
	// Note that the older, non-strict (or lax) binary protocol is not included in automatic protocol
	// detection. For upstream connections, the Thrift proxy will use the same protocol as the
	// downstream connection.
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	// The Thrift proxy will use the Thrift binary protocol.
	ProtocolType_BINARY ProtocolType = 1
	// The Thrift proxy will use Thrift non-strict binary protocol.
	ProtocolType_LAX_BINARY ProtocolType = 2
	// The Thrift proxy will use the Thrift compact protocol.
	ProtocolType_COMPACT ProtocolType = 3
	// The Thrift proxy will use the Thrift "Twitter" protocol implemented by the finagle library.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto.
	ProtocolType_TWITTER ProtocolType = 4
)

// Enum value maps for ProtocolType.
var (
	ProtocolType_name = map[int32]string{
		0: "AUTO_PROTOCOL",
		1: "BINARY",
		2: "LAX_BINARY",
		3: "COMPACT",
		4: "TWITTER",
	}
	ProtocolType_value = map[string]int32{
		"AUTO_PROTOCOL": 0,
		"BINARY":        1,
		"LAX_BINARY":    2,
		"COMPACT":       3,
		"TWITTER":       4,
	}
)

func (x ProtocolType) Enum() *ProtocolType {
	p := new(ProtocolType)
	*p = x
	return p
}

func (x ProtocolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes[1].Descriptor()
}

func (ProtocolType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes[1]
}

func (x ProtocolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolType.Descriptor instead.
func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{1}
}

type Trds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration source specifier.
	// In case of “api_config_source“ only aggregated “api_type“ is supported.
	ConfigSource *v3.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// The name of the route configuration. This allows to use different route
	// configurations. Tells which route configuration should be fetched from the configuration source.
	// Leave unspecified is also valid and means the unnamed route configuration.
	RouteConfigName string `protobuf:"bytes,2,opt,name=route_config_name,json=routeConfigName,proto3" json:"route_config_name,omitempty"`
}

func (x *Trds) Reset() {
	*x = Trds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trds) ProtoMessage() {}

func (x *Trds) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trds.ProtoReflect.Descriptor instead.
func (*Trds) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *Trds) GetConfigSource() *v3.ConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

func (x *Trds) GetRouteConfigName() string {
	if x != nil {
		return x.RouteConfigName
	}
	return ""
}

// [#next-free-field: 11]
type ThriftProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Supplies the type of transport that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_TRANSPORT<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`.
	Transport TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_PROTOCOL<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`.
	Protocol ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	// It is invalid to define both “route_config“ and “trds“.
	RouteConfig *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// Use xDS to fetch the route configuration. It is invalid to define both “route_config“ and “trds“.
	Trds *Trds `protobuf:"bytes,8,opt,name=trds,proto3" json:"trds,omitempty"`
	// A list of individual Thrift filters that make up the filter chain for requests made to the
	// Thrift proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no thrift_filters are specified, a default Thrift router filter
	// (“envoy.filters.thrift.router“) is used.
	// [#extension-category: envoy.thrift_proxy.filters]
	ThriftFilters []*ThriftFilter `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	// If set to true, Envoy will try to skip decode data after metadata in the Thrift message.
	// This mode will only work if the upstream and downstream protocols are the same and the transports
	// are Framed or Header, and the protocol is not Twitter. Otherwise Envoy will
	// fallback to decode the data.
	PayloadPassthrough bool `protobuf:"varint,6,opt,name=payload_passthrough,json=payloadPassthrough,proto3" json:"payload_passthrough,omitempty"`
	// Optional maximum requests for a single downstream connection. If not specified, there is no limit.
	MaxRequestsPerConnection *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_requests_per_connection,json=maxRequestsPerConnection,proto3" json:"max_requests_per_connection,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by Thrift proxy.
	AccessLog []*v31.AccessLog `protobuf:"bytes,9,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// If set to true, Envoy will preserve the case of Thrift header keys instead of serializing them to
	// lower case as per the default behavior. Note that NUL, CR and LF characters will also be preserved
	// as mandated by the Thrift spec.
	//
	// More info: https://github.com/apache/thrift/commit/e165fa3c85d00cb984f4d9635ed60909a1266ce1.
	HeaderKeysPreserveCase bool `protobuf:"varint,10,opt,name=header_keys_preserve_case,json=headerKeysPreserveCase,proto3" json:"header_keys_preserve_case,omitempty"`
}

func (x *ThriftProxy) Reset() {
	*x = ThriftProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThriftProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThriftProxy) ProtoMessage() {}

func (x *ThriftProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThriftProxy.ProtoReflect.Descriptor instead.
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *ThriftProxy) GetTransport() TransportType {
	if x != nil {
		return x.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (x *ThriftProxy) GetProtocol() ProtocolType {
	if x != nil {
		return x.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (x *ThriftProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if x != nil {
		return x.RouteConfig
	}
	return nil
}

func (x *ThriftProxy) GetTrds() *Trds {
	if x != nil {
		return x.Trds
	}
	return nil
}

func (x *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if x != nil {
		return x.ThriftFilters
	}
	return nil
}

func (x *ThriftProxy) GetPayloadPassthrough() bool {
	if x != nil {
		return x.PayloadPassthrough
	}
	return false
}

func (x *ThriftProxy) GetMaxRequestsPerConnection() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxRequestsPerConnection
	}
	return nil
}

func (x *ThriftProxy) GetAccessLog() []*v31.AccessLog {
	if x != nil {
		return x.AccessLog
	}
	return nil
}

func (x *ThriftProxy) GetHeaderKeysPreserveCase() bool {
	if x != nil {
		return x.HeaderKeysPreserveCase
	}
	return false
}

// ThriftFilter configures a Thrift filter.
type ThriftFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(zuercher): Auto generate the following list]
	// * :ref:`envoy.filters.thrift.router <config_thrift_filters_router>`
	// * :ref:`envoy.filters.thrift.rate_limit <config_thrift_filters_rate_limit>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated. See the supported
	// filters for further documentation.
	//
	// Types that are assignable to ConfigType:
	//
	//	*ThriftFilter_TypedConfig
	ConfigType isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
}

func (x *ThriftFilter) Reset() {
	*x = ThriftFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThriftFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThriftFilter) ProtoMessage() {}

func (x *ThriftFilter) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThriftFilter.ProtoReflect.Descriptor instead.
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ThriftFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *ThriftFilter) GetTypedConfig() *any1.Any {
	if x, ok := x.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any1.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

// ThriftProtocolOptions specifies Thrift upstream protocol options. This object is used in
// in
// :ref:`typed_extension_protocol_options<envoy_v3_api_field_config.cluster.v3.Cluster.typed_extension_protocol_options>`,
// keyed by the name “envoy.filters.network.thrift_proxy“.
type ThriftProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Supplies the type of transport that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_TRANSPORT<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`,
	// which is the default, causes the proxy to use the same transport as the downstream connection.
	Transport TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_PROTOCOL<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`,
	// which is the default, causes the proxy to use the same protocol as the downstream connection.
	Protocol ProtocolType `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
}

func (x *ThriftProtocolOptions) Reset() {
	*x = ThriftProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThriftProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThriftProtocolOptions) ProtoMessage() {}

func (x *ThriftProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThriftProtocolOptions.ProtoReflect.Descriptor instead.
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *ThriftProtocolOptions) GetTransport() TransportType {
	if x != nil {
		return x.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (x *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if x != nil {
		return x.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

var File_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x76, 0x33, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x04, 0x54, 0x72, 0x64, 0x73, 0x12, 0x51,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x07,
	0x0a, 0x0b, 0x54, 0x68, 0x72, 0x69, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x67, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x64, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x17, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x11, 0x12, 0x0f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0b, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x63, 0x0a, 0x04, 0x74, 0x72, 0x64,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x64, 0x73, 0x42,
	0x17, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x11, 0x12, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x72, 0x64, 0x73, 0x12, 0x65,
	0x0a, 0x0e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x68, 0x72, 0x69, 0x66, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x73, 0x73, 0x74,
	0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x12, 0x5b, 0x0a, 0x1b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f,
	0x67, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x39, 0x0a, 0x19, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x3a, 0x44, 0x9a, 0xc5, 0x88, 0x1e, 0x3f, 0x0a, 0x3d, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x54, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3a, 0x45, 0x9a, 0xc5, 0x88, 0x1e, 0x40, 0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x68, 0x72,
	0x69, 0x66, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb6, 0x02, 0x0a, 0x15, 0x54, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x67, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x64, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a,
	0x4e, 0x9a, 0xc5, 0x88, 0x1e, 0x49, 0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x69, 0x66, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a,
	0x49, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x0e, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f,
	0x52, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x64, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55,
	0x54, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x41, 0x58,
	0x5f, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x43, 0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45,
	0x52, 0x10, 0x04, 0x1a, 0x0b, 0x8a, 0xf4, 0x9b, 0xb3, 0x05, 0x03, 0x33, 0x2e, 0x30, 0x08, 0x01,
	0x42, 0xc4, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3e, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x54, 0x68, 0x72,
	0x69, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescData = file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_goTypes = []interface{}{
	(TransportType)(0),            // 0: envoy.extensions.filters.network.thrift_proxy.v3.TransportType
	(ProtocolType)(0),             // 1: envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType
	(*Trds)(nil),                  // 2: envoy.extensions.filters.network.thrift_proxy.v3.Trds
	(*ThriftProxy)(nil),           // 3: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy
	(*ThriftFilter)(nil),          // 4: envoy.extensions.filters.network.thrift_proxy.v3.ThriftFilter
	(*ThriftProtocolOptions)(nil), // 5: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions
	(*v3.ConfigSource)(nil),       // 6: envoy.config.core.v3.ConfigSource
	(*RouteConfiguration)(nil),    // 7: envoy.extensions.filters.network.thrift_proxy.v3.RouteConfiguration
	(*wrappers.UInt32Value)(nil),  // 8: google.protobuf.UInt32Value
	(*v31.AccessLog)(nil),         // 9: envoy.config.accesslog.v3.AccessLog
	(*any1.Any)(nil),              // 10: google.protobuf.Any
}
var file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_depIdxs = []int32{
	6,  // 0: envoy.extensions.filters.network.thrift_proxy.v3.Trds.config_source:type_name -> envoy.config.core.v3.ConfigSource
	0,  // 1: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.transport:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.TransportType
	1,  // 2: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.protocol:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType
	7,  // 3: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.route_config:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.RouteConfiguration
	2,  // 4: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.trds:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.Trds
	4,  // 5: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.thrift_filters:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.ThriftFilter
	8,  // 6: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.max_requests_per_connection:type_name -> google.protobuf.UInt32Value
	9,  // 7: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy.access_log:type_name -> envoy.config.accesslog.v3.AccessLog
	10, // 8: envoy.extensions.filters.network.thrift_proxy.v3.ThriftFilter.typed_config:type_name -> google.protobuf.Any
	0,  // 9: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions.transport:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.TransportType
	1,  // 10: envoy.extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions.protocol:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_init() }
func file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_init() {
	if File_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto != nil {
		return
	}
	file_envoy_extensions_filters_network_thrift_proxy_v3_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trds); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThriftProxy); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThriftFilter); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThriftProtocolOptions); i {
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
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ThriftFilter_TypedConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto = out.File
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_network_thrift_proxy_v3_thrift_proxy_proto_depIdxs = nil
}
