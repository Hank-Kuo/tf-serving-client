// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/protobuf/autotuning.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	xla "grpc-client/internal/tensorflow/tensorflow/tensorflow/compiler/xla"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/compiler/xla/autotuning.proto.

type AutotuneResult_FailureKind = xla.AutotuneResult_FailureKind

const AutotuneResult_UNKNOWN = xla.AutotuneResult_UNKNOWN
const AutotuneResult_REDZONE_MODIFIED = xla.AutotuneResult_REDZONE_MODIFIED
const AutotuneResult_WRONG_RESULT = xla.AutotuneResult_WRONG_RESULT
const AutotuneResult_DISQUALIFIED = xla.AutotuneResult_DISQUALIFIED

var AutotuneResult_FailureKind_name = xla.AutotuneResult_FailureKind_name
var AutotuneResult_FailureKind_value = xla.AutotuneResult_FailureKind_value

type CudnnVersion = xla.CudnnVersion
type ComputeCapability = xla.ComputeCapability
type AutotuneResult = xla.AutotuneResult
type AutotuneResult_Conv = xla.AutotuneResult_Conv
type AutotuneResult_Gemm = xla.AutotuneResult_Gemm
type AutotuneResult_Triton = xla.AutotuneResult_Triton
type AutotuneResult_CudaConvPlan = xla.AutotuneResult_CudaConvPlan
type AutotuneResult_Algorithm = xla.AutotuneResult_Algorithm
type AutotuningLog = xla.AutotuningLog
type AutotuneResult_FailureResult = xla.AutotuneResult_FailureResult
type AutotuneResult_FailureResult_ReferenceConv = xla.AutotuneResult_FailureResult_ReferenceConv
type AutotuneResult_FailureResult_ReferenceGemm = xla.AutotuneResult_FailureResult_ReferenceGemm
type AutotuneResult_FailureResult_ReferenceCudaConvPlan = xla.AutotuneResult_FailureResult_ReferenceCudaConvPlan
type AutotuneResult_FailureResult_ReferenceAlgorithm = xla.AutotuneResult_FailureResult_ReferenceAlgorithm
type AutotuneResult_ConvKey = xla.AutotuneResult_ConvKey
type AutotuneResult_GemmKey = xla.AutotuneResult_GemmKey
type AutotuneResult_CudaConvPlanKey = xla.AutotuneResult_CudaConvPlanKey
type AutotuneResult_TritonGemmKey = xla.AutotuneResult_TritonGemmKey

var File_tensorflow_core_protobuf_autotuning_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_autotuning_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x1a, 0x28, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xcd, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x42, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x44, 0x58, 0xaa, 0x02,
	0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x3a, 0x3a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_tensorflow_core_protobuf_autotuning_proto_goTypes = []interface{}{}
var file_tensorflow_core_protobuf_autotuning_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_autotuning_proto_init() }
func file_tensorflow_core_protobuf_autotuning_proto_init() {
	if File_tensorflow_core_protobuf_autotuning_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_autotuning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_autotuning_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_autotuning_proto_depIdxs,
	}.Build()
	File_tensorflow_core_protobuf_autotuning_proto = out.File
	file_tensorflow_core_protobuf_autotuning_proto_rawDesc = nil
	file_tensorflow_core_protobuf_autotuning_proto_goTypes = nil
	file_tensorflow_core_protobuf_autotuning_proto_depIdxs = nil
}
