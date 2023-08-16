// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/service/cpu/backend_config.proto

package cpu

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

// These enum needs to be mapped to oneDNN enum for post_op algorithm.
// TODO(intel-tf): Add kinds supported by oneDNN.
type OneDnnMatMulConfig_FusionKind int32

const (
	OneDnnMatMulConfig_UNDEFINED OneDnnMatMulConfig_FusionKind = 0
	OneDnnMatMulConfig_BIAS      OneDnnMatMulConfig_FusionKind = 1
	OneDnnMatMulConfig_RELU      OneDnnMatMulConfig_FusionKind = 2
	OneDnnMatMulConfig_TANH      OneDnnMatMulConfig_FusionKind = 3
	OneDnnMatMulConfig_GELU_ERF  OneDnnMatMulConfig_FusionKind = 4
	OneDnnMatMulConfig_GELU_TANH OneDnnMatMulConfig_FusionKind = 5
)

// Enum value maps for OneDnnMatMulConfig_FusionKind.
var (
	OneDnnMatMulConfig_FusionKind_name = map[int32]string{
		0: "UNDEFINED",
		1: "BIAS",
		2: "RELU",
		3: "TANH",
		4: "GELU_ERF",
		5: "GELU_TANH",
	}
	OneDnnMatMulConfig_FusionKind_value = map[string]int32{
		"UNDEFINED": 0,
		"BIAS":      1,
		"RELU":      2,
		"TANH":      3,
		"GELU_ERF":  4,
		"GELU_TANH": 5,
	}
)

func (x OneDnnMatMulConfig_FusionKind) Enum() *OneDnnMatMulConfig_FusionKind {
	p := new(OneDnnMatMulConfig_FusionKind)
	*p = x
	return p
}

func (x OneDnnMatMulConfig_FusionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OneDnnMatMulConfig_FusionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_xla_service_cpu_backend_config_proto_enumTypes[0].Descriptor()
}

func (OneDnnMatMulConfig_FusionKind) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_xla_service_cpu_backend_config_proto_enumTypes[0]
}

func (x OneDnnMatMulConfig_FusionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OneDnnMatMulConfig_FusionKind.Descriptor instead.
func (OneDnnMatMulConfig_FusionKind) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescGZIP(), []int{1, 0}
}

// Backend config for XLA:CPU.
type BackendConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of partitions per outer dimension (in order, starting with
	// outer-most dimension first). Used by the parallel cpu backend to partition
	// HLOs into parallel tasks.
	OuterDimensionPartitions []int64 `protobuf:"varint,1,rep,packed,name=outer_dimension_partitions,json=outerDimensionPartitions,proto3" json:"outer_dimension_partitions,omitempty"`
	// Configuration to be used by oneDNN matmul
	OnednnMatmulConfig *OneDnnMatMulConfig `protobuf:"bytes,2,opt,name=onednn_matmul_config,json=onednnMatmulConfig,proto3" json:"onednn_matmul_config,omitempty"`
}

func (x *BackendConfig) Reset() {
	*x = BackendConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendConfig) ProtoMessage() {}

func (x *BackendConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendConfig.ProtoReflect.Descriptor instead.
func (*BackendConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescGZIP(), []int{0}
}

func (x *BackendConfig) GetOuterDimensionPartitions() []int64 {
	if x != nil {
		return x.OuterDimensionPartitions
	}
	return nil
}

func (x *BackendConfig) GetOnednnMatmulConfig() *OneDnnMatMulConfig {
	if x != nil {
		return x.OnednnMatmulConfig
	}
	return nil
}

type OneDnnMatMulConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FusedOps []OneDnnMatMulConfig_FusionKind `protobuf:"varint,3,rep,packed,name=fused_ops,json=fusedOps,proto3,enum=xla.cpu.OneDnnMatMulConfig_FusionKind" json:"fused_ops,omitempty"`
}

func (x *OneDnnMatMulConfig) Reset() {
	*x = OneDnnMatMulConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneDnnMatMulConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneDnnMatMulConfig) ProtoMessage() {}

func (x *OneDnnMatMulConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneDnnMatMulConfig.ProtoReflect.Descriptor instead.
func (*OneDnnMatMulConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescGZIP(), []int{1}
}

func (x *OneDnnMatMulConfig) GetFusedOps() []OneDnnMatMulConfig_FusionKind {
	if x != nil {
		return x.FusedOps
	}
	return nil
}

var File_tensorflow_compiler_xla_service_cpu_backend_config_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x63, 0x70, 0x75, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x6c, 0x61, 0x2e,
	0x63, 0x70, 0x75, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x1a, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x18, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x4d, 0x0a, 0x14, 0x6f, 0x6e, 0x65, 0x64, 0x6e, 0x6e, 0x5f, 0x6d, 0x61,
	0x74, 0x6d, 0x75, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x63, 0x70, 0x75, 0x2e, 0x4f, 0x6e, 0x65, 0x44,
	0x6e, 0x6e, 0x4d, 0x61, 0x74, 0x4d, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12,
	0x6f, 0x6e, 0x65, 0x64, 0x6e, 0x6e, 0x4d, 0x61, 0x74, 0x6d, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x4f, 0x6e, 0x65, 0x44, 0x6e, 0x6e, 0x4d, 0x61, 0x74,
	0x4d, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x09, 0x66, 0x75, 0x73,
	0x65, 0x64, 0x5f, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x78,
	0x6c, 0x61, 0x2e, 0x63, 0x70, 0x75, 0x2e, 0x4f, 0x6e, 0x65, 0x44, 0x6e, 0x6e, 0x4d, 0x61, 0x74,
	0x4d, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x08, 0x66, 0x75, 0x73, 0x65, 0x64, 0x4f, 0x70, 0x73, 0x22, 0x56,
	0x0a, 0x0a, 0x46, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x49, 0x41, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x4c, 0x55, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x41, 0x4e, 0x48, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x4c,
	0x55, 0x5f, 0x45, 0x52, 0x46, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x4c, 0x55, 0x5f,
	0x54, 0x41, 0x4e, 0x48, 0x10, 0x05, 0x42, 0xae, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x78,
	0x6c, 0x61, 0x2e, 0x63, 0x70, 0x75, 0x42, 0x12, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x70, 0x75, 0xa2, 0x02, 0x03, 0x58,
	0x43, 0x58, 0xaa, 0x02, 0x07, 0x58, 0x6c, 0x61, 0x2e, 0x43, 0x70, 0x75, 0xca, 0x02, 0x07, 0x58,
	0x6c, 0x61, 0x5c, 0x43, 0x70, 0x75, 0xe2, 0x02, 0x13, 0x58, 0x6c, 0x61, 0x5c, 0x43, 0x70, 0x75,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x58,
	0x6c, 0x61, 0x3a, 0x3a, 0x43, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescData = file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDesc
)

func file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDescData
}

var file_tensorflow_compiler_xla_service_cpu_backend_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_compiler_xla_service_cpu_backend_config_proto_goTypes = []interface{}{
	(OneDnnMatMulConfig_FusionKind)(0), // 0: xla.cpu.OneDnnMatMulConfig.FusionKind
	(*BackendConfig)(nil),              // 1: xla.cpu.BackendConfig
	(*OneDnnMatMulConfig)(nil),         // 2: xla.cpu.OneDnnMatMulConfig
}
var file_tensorflow_compiler_xla_service_cpu_backend_config_proto_depIdxs = []int32{
	2, // 0: xla.cpu.BackendConfig.onednn_matmul_config:type_name -> xla.cpu.OneDnnMatMulConfig
	0, // 1: xla.cpu.OneDnnMatMulConfig.fused_ops:type_name -> xla.cpu.OneDnnMatMulConfig.FusionKind
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_service_cpu_backend_config_proto_init() }
func file_tensorflow_compiler_xla_service_cpu_backend_config_proto_init() {
	if File_tensorflow_compiler_xla_service_cpu_backend_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendConfig); i {
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
		file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneDnnMatMulConfig); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_service_cpu_backend_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_service_cpu_backend_config_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_xla_service_cpu_backend_config_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_xla_service_cpu_backend_config_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_service_cpu_backend_config_proto = out.File
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_rawDesc = nil
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_goTypes = nil
	file_tensorflow_compiler_xla_service_cpu_backend_config_proto_depIdxs = nil
}