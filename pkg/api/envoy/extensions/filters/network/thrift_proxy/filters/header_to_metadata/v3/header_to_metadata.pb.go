// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/extensions/filters/network/thrift_proxy/filters/header_to_metadata/v3/header_to_metadata.proto

package header_to_metadatav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type HeaderToMetadata_ValueType int32

const (
	HeaderToMetadata_STRING HeaderToMetadata_ValueType = 0
	HeaderToMetadata_NUMBER HeaderToMetadata_ValueType = 1
	// The value is a serialized `protobuf.Value
	// <https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/struct.proto#L62>`_.
	HeaderToMetadata_PROTOBUF_VALUE HeaderToMetadata_ValueType = 2
)

// Enum value maps for HeaderToMetadata_ValueType.
var (
	HeaderToMetadata_ValueType_name = map[int32]string{
		0: "STRING",
		1: "NUMBER",
		2: "PROTOBUF_VALUE",
	}
	HeaderToMetadata_ValueType_value = map[string]int32{
		"STRING":         0,
		"NUMBER":         1,
		"PROTOBUF_VALUE": 2,
	}
)

func (x HeaderToMetadata_ValueType) Enum() *HeaderToMetadata_ValueType {
	p := new(HeaderToMetadata_ValueType)
	*p = x
	return p
}

func (x HeaderToMetadata_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeaderToMetadata_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes[0].Descriptor()
}

func (HeaderToMetadata_ValueType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes[0]
}

func (x HeaderToMetadata_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeaderToMetadata_ValueType.Descriptor instead.
func (HeaderToMetadata_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

// ValueEncode defines the encoding algorithm.
type HeaderToMetadata_ValueEncode int32

const (
	// The value is not encoded.
	HeaderToMetadata_NONE HeaderToMetadata_ValueEncode = 0
	// The value is encoded in `Base64 <https://tools.ietf.org/html/rfc4648#section-4>`_.
	// Note: this is mostly used for STRING and PROTOBUF_VALUE to escape the
	// non-ASCII characters in the header.
	HeaderToMetadata_BASE64 HeaderToMetadata_ValueEncode = 1
)

// Enum value maps for HeaderToMetadata_ValueEncode.
var (
	HeaderToMetadata_ValueEncode_name = map[int32]string{
		0: "NONE",
		1: "BASE64",
	}
	HeaderToMetadata_ValueEncode_value = map[string]int32{
		"NONE":   0,
		"BASE64": 1,
	}
)

func (x HeaderToMetadata_ValueEncode) Enum() *HeaderToMetadata_ValueEncode {
	p := new(HeaderToMetadata_ValueEncode)
	*p = x
	return p
}

func (x HeaderToMetadata_ValueEncode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeaderToMetadata_ValueEncode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes[1].Descriptor()
}

func (HeaderToMetadata_ValueEncode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes[1]
}

func (x HeaderToMetadata_ValueEncode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeaderToMetadata_ValueEncode.Descriptor instead.
func (HeaderToMetadata_ValueEncode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP(), []int{0, 1}
}

type HeaderToMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules to apply to requests.
	RequestRules []*HeaderToMetadata_Rule `protobuf:"bytes,1,rep,name=request_rules,json=requestRules,proto3" json:"request_rules,omitempty"`
}

func (x *HeaderToMetadata) Reset() {
	*x = HeaderToMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderToMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderToMetadata) ProtoMessage() {}

func (x *HeaderToMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderToMetadata.ProtoReflect.Descriptor instead.
func (*HeaderToMetadata) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderToMetadata) GetRequestRules() []*HeaderToMetadata_Rule {
	if x != nil {
		return x.RequestRules
	}
	return nil
}

// [#next-free-field: 7]
type HeaderToMetadata_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to ValueType:
	//	*HeaderToMetadata_KeyValuePair_Value
	//	*HeaderToMetadata_KeyValuePair_RegexValueRewrite
	ValueType isHeaderToMetadata_KeyValuePair_ValueType `protobuf_oneof:"value_type"`
	// The value's type — defaults to string.
	Type HeaderToMetadata_ValueType `protobuf:"varint,5,opt,name=type,proto3,enum=envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata_ValueType" json:"type,omitempty"`
	// How is the value encoded, default is NONE (not encoded).
	// The value will be decoded accordingly before storing to metadata.
	Encode HeaderToMetadata_ValueEncode `protobuf:"varint,6,opt,name=encode,proto3,enum=envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata_ValueEncode" json:"encode,omitempty"`
}

func (x *HeaderToMetadata_KeyValuePair) Reset() {
	*x = HeaderToMetadata_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderToMetadata_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderToMetadata_KeyValuePair) ProtoMessage() {}

func (x *HeaderToMetadata_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderToMetadata_KeyValuePair.ProtoReflect.Descriptor instead.
func (*HeaderToMetadata_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HeaderToMetadata_KeyValuePair) GetMetadataNamespace() string {
	if x != nil {
		return x.MetadataNamespace
	}
	return ""
}

func (x *HeaderToMetadata_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *HeaderToMetadata_KeyValuePair) GetValueType() isHeaderToMetadata_KeyValuePair_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (x *HeaderToMetadata_KeyValuePair) GetValue() string {
	if x, ok := x.GetValueType().(*HeaderToMetadata_KeyValuePair_Value); ok {
		return x.Value
	}
	return ""
}

func (x *HeaderToMetadata_KeyValuePair) GetRegexValueRewrite() *v3.RegexMatchAndSubstitute {
	if x, ok := x.GetValueType().(*HeaderToMetadata_KeyValuePair_RegexValueRewrite); ok {
		return x.RegexValueRewrite
	}
	return nil
}

func (x *HeaderToMetadata_KeyValuePair) GetType() HeaderToMetadata_ValueType {
	if x != nil {
		return x.Type
	}
	return HeaderToMetadata_STRING
}

func (x *HeaderToMetadata_KeyValuePair) GetEncode() HeaderToMetadata_ValueEncode {
	if x != nil {
		return x.Encode
	}
	return HeaderToMetadata_NONE
}

type isHeaderToMetadata_KeyValuePair_ValueType interface {
	isHeaderToMetadata_KeyValuePair_ValueType()
}

type HeaderToMetadata_KeyValuePair_Value struct {
	// The value to pair with the given key.
	//
	// When used for on_present case, if value is non-empty it'll be used instead
	// of the header value. If both are empty, no metadata is added.
	//
	// When used for on_missing case, a non-empty value must be provided otherwise
	// no metadata is added.
	Value string `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

type HeaderToMetadata_KeyValuePair_RegexValueRewrite struct {
	// If present, the header's value will be matched and substituted with this.
	// If there is no match or substitution, the header value
	// is used as-is.
	//
	// This is only used for on_present.
	//
	// Note: if the ``value`` field is non-empty this field should be empty.
	RegexValueRewrite *v3.RegexMatchAndSubstitute `protobuf:"bytes,4,opt,name=regex_value_rewrite,json=regexValueRewrite,proto3,oneof"`
}

func (*HeaderToMetadata_KeyValuePair_Value) isHeaderToMetadata_KeyValuePair_ValueType() {}

func (*HeaderToMetadata_KeyValuePair_RegexValueRewrite) isHeaderToMetadata_KeyValuePair_ValueType() {}

// A Rule defines what metadata to apply when a header is present or missing.
type HeaderToMetadata_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies that a match will be performed on the value of a header.
	//
	// The header to be extracted.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// If the header is present, apply this metadata KeyValuePair.
	//
	// If the value in the KeyValuePair is non-empty, it'll be used instead
	// of the header value.
	OnPresent *HeaderToMetadata_KeyValuePair `protobuf:"bytes,2,opt,name=on_present,json=onPresent,proto3" json:"on_present,omitempty"`
	// If the header is not present, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set, since it'll be used in lieu
	// of the missing header value.
	OnMissing *HeaderToMetadata_KeyValuePair `protobuf:"bytes,3,opt,name=on_missing,json=onMissing,proto3" json:"on_missing,omitempty"`
	// Whether or not to remove the header after a rule is applied.
	//
	// This prevents headers from leaking.
	Remove bool `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (x *HeaderToMetadata_Rule) Reset() {
	*x = HeaderToMetadata_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderToMetadata_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderToMetadata_Rule) ProtoMessage() {}

func (x *HeaderToMetadata_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderToMetadata_Rule.ProtoReflect.Descriptor instead.
func (*HeaderToMetadata_Rule) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HeaderToMetadata_Rule) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *HeaderToMetadata_Rule) GetOnPresent() *HeaderToMetadata_KeyValuePair {
	if x != nil {
		return x.OnPresent
	}
	return nil
}

func (x *HeaderToMetadata_Rule) GetOnMissing() *HeaderToMetadata_KeyValuePair {
	if x != nil {
		return x.OnMissing
	}
	return nil
}

func (x *HeaderToMetadata_Rule) GetRemove() bool {
	if x != nil {
		return x.Remove
	}
	return false
}

var File_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDesc = []byte{
	0x0a, 0x64, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x33, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3,
	0x08, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x62, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0xec, 0x03, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6e, 0x64, 0x53, 0x75, 0x62,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x85, 0x01, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x67, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x52, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xdd, 0x02, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x01, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x09, 0x6f, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x42, 0x55, 0x46, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x22,
	0x23, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x41, 0x53, 0x45,
	0x36, 0x34, 0x10, 0x01, 0x42, 0x86, 0x02, 0x0a, 0x59, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x33, 0x42, 0x15, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x87, 0x01, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33,
	0x3b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescData = file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDesc
)

func file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDescData
}

var file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_goTypes = []interface{}{
	(HeaderToMetadata_ValueType)(0),       // 0: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.ValueType
	(HeaderToMetadata_ValueEncode)(0),     // 1: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.ValueEncode
	(*HeaderToMetadata)(nil),              // 2: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata
	(*HeaderToMetadata_KeyValuePair)(nil), // 3: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair
	(*HeaderToMetadata_Rule)(nil),         // 4: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.Rule
	(*v3.RegexMatchAndSubstitute)(nil),    // 5: envoy.type.matcher.v3.RegexMatchAndSubstitute
}
var file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.request_rules:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.Rule
	5, // 1: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair.regex_value_rewrite:type_name -> envoy.type.matcher.v3.RegexMatchAndSubstitute
	0, // 2: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair.type:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.ValueType
	1, // 3: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair.encode:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.ValueEncode
	3, // 4: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.Rule.on_present:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair
	3, // 5: envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.Rule.on_missing:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.header_to_metadata.v3.HeaderToMetadata.KeyValuePair
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_init()
}
func file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_init() {
	if File_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderToMetadata); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderToMetadata_KeyValuePair); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderToMetadata_Rule); i {
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
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HeaderToMetadata_KeyValuePair_Value)(nil),
		(*HeaderToMetadata_KeyValuePair_RegexValueRewrite)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto = out.File
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_rawDesc = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_goTypes = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_header_to_metadata_v3_header_to_metadata_proto_depIdxs = nil
}
