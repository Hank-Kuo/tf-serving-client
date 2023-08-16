// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow_serving/config/logging_config.proto

package config

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

// Attributes of requests that can be optionally sampled.
// Note: Enabling more attributes will increase logging storage requirements.
type SamplingConfig_Attributes int32

const (
	SamplingConfig_ATTR_DEFAULT             SamplingConfig_Attributes = 0
	SamplingConfig_ATTR_REQUEST_ORIGIN      SamplingConfig_Attributes = 1
	SamplingConfig_ATTR_REQUEST_CRITICALITY SamplingConfig_Attributes = 2
)

// Enum value maps for SamplingConfig_Attributes.
var (
	SamplingConfig_Attributes_name = map[int32]string{
		0: "ATTR_DEFAULT",
		1: "ATTR_REQUEST_ORIGIN",
		2: "ATTR_REQUEST_CRITICALITY",
	}
	SamplingConfig_Attributes_value = map[string]int32{
		"ATTR_DEFAULT":             0,
		"ATTR_REQUEST_ORIGIN":      1,
		"ATTR_REQUEST_CRITICALITY": 2,
	}
)

func (x SamplingConfig_Attributes) Enum() *SamplingConfig_Attributes {
	p := new(SamplingConfig_Attributes)
	*p = x
	return p
}

func (x SamplingConfig_Attributes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamplingConfig_Attributes) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_serving_config_logging_config_proto_enumTypes[0].Descriptor()
}

func (SamplingConfig_Attributes) Type() protoreflect.EnumType {
	return &file_tensorflow_serving_config_logging_config_proto_enumTypes[0]
}

func (x SamplingConfig_Attributes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamplingConfig_Attributes.Descriptor instead.
func (SamplingConfig_Attributes) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_logging_config_proto_rawDescGZIP(), []int{0, 0}
}

type SamplingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requests will be logged uniformly at random with this probability.
	// Valid range: [0, 1.0].
	SamplingRate float64 `protobuf:"fixed64,1,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	// Bitwise OR of above attributes
	Attributes int32 `protobuf:"varint,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *SamplingConfig) Reset() {
	*x = SamplingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_logging_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplingConfig) ProtoMessage() {}

func (x *SamplingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_logging_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplingConfig.ProtoReflect.Descriptor instead.
func (*SamplingConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_logging_config_proto_rawDescGZIP(), []int{0}
}

func (x *SamplingConfig) GetSamplingRate() float64 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *SamplingConfig) GetAttributes() int32 {
	if x != nil {
		return x.Attributes
	}
	return 0
}

// Configuration for logging query/responses.
type LoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogCollectorConfig *LogCollectorConfig `protobuf:"bytes,1,opt,name=log_collector_config,json=logCollectorConfig,proto3" json:"log_collector_config,omitempty"`
	SamplingConfig     *SamplingConfig     `protobuf:"bytes,2,opt,name=sampling_config,json=samplingConfig,proto3" json:"sampling_config,omitempty"`
}

func (x *LoggingConfig) Reset() {
	*x = LoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_logging_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingConfig) ProtoMessage() {}

func (x *LoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_logging_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingConfig.ProtoReflect.Descriptor instead.
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_logging_config_proto_rawDescGZIP(), []int{1}
}

func (x *LoggingConfig) GetLogCollectorConfig() *LogCollectorConfig {
	if x != nil {
		return x.LogCollectorConfig
	}
	return nil
}

func (x *LoggingConfig) GetSamplingConfig() *SamplingConfig {
	if x != nil {
		return x.SamplingConfig
	}
	return nil
}

var File_tensorflow_serving_config_logging_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_logging_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x1a, 0x34, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x22, 0x55, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x54, 0x54, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x52, 0x49, 0x54,
	0x49, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x02, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x14, 0x6c,
	0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0xde, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x12, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54,
	0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_logging_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_logging_config_proto_rawDescData = file_tensorflow_serving_config_logging_config_proto_rawDesc
)

func file_tensorflow_serving_config_logging_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_logging_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_logging_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_logging_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_logging_config_proto_rawDescData
}

var file_tensorflow_serving_config_logging_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_serving_config_logging_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_serving_config_logging_config_proto_goTypes = []interface{}{
	(SamplingConfig_Attributes)(0), // 0: tensorflow.serving.SamplingConfig.Attributes
	(*SamplingConfig)(nil),         // 1: tensorflow.serving.SamplingConfig
	(*LoggingConfig)(nil),          // 2: tensorflow.serving.LoggingConfig
	(*LogCollectorConfig)(nil),     // 3: tensorflow.serving.LogCollectorConfig
}
var file_tensorflow_serving_config_logging_config_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.LoggingConfig.log_collector_config:type_name -> tensorflow.serving.LogCollectorConfig
	1, // 1: tensorflow.serving.LoggingConfig.sampling_config:type_name -> tensorflow.serving.SamplingConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_logging_config_proto_init() }
func file_tensorflow_serving_config_logging_config_proto_init() {
	if File_tensorflow_serving_config_logging_config_proto != nil {
		return
	}
	file_tensorflow_serving_config_log_collector_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_logging_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplingConfig); i {
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
		file_tensorflow_serving_config_logging_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingConfig); i {
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
			RawDescriptor: file_tensorflow_serving_config_logging_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_logging_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_logging_config_proto_depIdxs,
		EnumInfos:         file_tensorflow_serving_config_logging_config_proto_enumTypes,
		MessageInfos:      file_tensorflow_serving_config_logging_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_logging_config_proto = out.File
	file_tensorflow_serving_config_logging_config_proto_rawDesc = nil
	file_tensorflow_serving_config_logging_config_proto_goTypes = nil
	file_tensorflow_serving_config_logging_config_proto_depIdxs = nil
}