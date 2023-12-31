// Copyright 2020 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/tpu/kernels/tpu_compilation_cache.proto

package kernels

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	tf2xla "grpc-client/internal/tensorflow/tensorflow/tensorflow/compiler/tf2xla"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Response for GetTpuProgram RPC.
type GetTpuProgramResponseExternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto               *GetTpuProgramResponseExternal_Blob `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	HostComputeMetadata *tf2xla.HostComputeMetadata         `protobuf:"bytes,2,opt,name=host_compute_metadata,json=hostComputeMetadata,proto3" json:"host_compute_metadata,omitempty"`
	MayModifyVariables  bool                                `protobuf:"varint,3,opt,name=may_modify_variables,json=mayModifyVariables,proto3" json:"may_modify_variables,omitempty"`
	CompilerMetadata    *GetTpuProgramResponseExternal_Blob `protobuf:"bytes,4,opt,name=compiler_metadata,json=compilerMetadata,proto3" json:"compiler_metadata,omitempty"`
	// Whether the program is empty, which could be true for sharding/unsharding
	// entries.
	IsEmpty bool `protobuf:"varint,5,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
}

func (x *GetTpuProgramResponseExternal) Reset() {
	*x = GetTpuProgramResponseExternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTpuProgramResponseExternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTpuProgramResponseExternal) ProtoMessage() {}

func (x *GetTpuProgramResponseExternal) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTpuProgramResponseExternal.ProtoReflect.Descriptor instead.
func (*GetTpuProgramResponseExternal) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescGZIP(), []int{0}
}

func (x *GetTpuProgramResponseExternal) GetProto() *GetTpuProgramResponseExternal_Blob {
	if x != nil {
		return x.Proto
	}
	return nil
}

func (x *GetTpuProgramResponseExternal) GetHostComputeMetadata() *tf2xla.HostComputeMetadata {
	if x != nil {
		return x.HostComputeMetadata
	}
	return nil
}

func (x *GetTpuProgramResponseExternal) GetMayModifyVariables() bool {
	if x != nil {
		return x.MayModifyVariables
	}
	return false
}

func (x *GetTpuProgramResponseExternal) GetCompilerMetadata() *GetTpuProgramResponseExternal_Blob {
	if x != nil {
		return x.CompilerMetadata
	}
	return nil
}

func (x *GetTpuProgramResponseExternal) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

type GetTpuProgramResponseExternal_Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTpuProgramResponseExternal_Blob) Reset() {
	*x = GetTpuProgramResponseExternal_Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTpuProgramResponseExternal_Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTpuProgramResponseExternal_Blob) ProtoMessage() {}

func (x *GetTpuProgramResponseExternal_Blob) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTpuProgramResponseExternal_Blob.ProtoReflect.Descriptor instead.
func (*GetTpuProgramResponseExternal_Blob) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetTpuProgramResponseExternal_Blob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto protoreflect.FileDescriptor

var file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDesc = []byte{
	0x0a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x70, 0x75, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x70,
	0x75, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x1a, 0x36, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x74,
	0x66, 0x32, 0x78, 0x6c, 0x61, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x70, 0x75, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x74,
	0x70, 0x75, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x48, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x74, 0x70, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x5a, 0x0a,
	0x15, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x13, 0x68, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x79,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6d, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x8c, 0x01, 0x0a, 0x22, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x54, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x24, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74,
	0x70, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x22, 0x00, 0x42, 0xcf, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x42, 0x18, 0x54, 0x70, 0x75, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x70, 0x75, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x73, 0xa2, 0x02, 0x03,
	0x54, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x54, 0x70, 0x75, 0xca, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x54, 0x70, 0x75, 0xe2, 0x02, 0x1a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5c, 0x54, 0x70, 0x75, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a,
	0x3a, 0x54, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescOnce sync.Once
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescData = file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDesc
)

func file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescGZIP() []byte {
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescData)
	})
	return file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDescData
}

var file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_goTypes = []interface{}{
	(*GetTpuProgramResponseExternal)(nil),      // 0: tensorflow.tpu.GetTpuProgramResponseExternal
	(*GetTpuProgramResponseExternal_Blob)(nil), // 1: tensorflow.tpu.GetTpuProgramResponseExternal.Blob
	(*tf2xla.HostComputeMetadata)(nil),         // 2: tensorflow.tf2xla.HostComputeMetadata
	(*GetTpuProgramRequest)(nil),               // 3: tensorflow.tpu.GetTpuProgramRequest
}
var file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_depIdxs = []int32{
	1, // 0: tensorflow.tpu.GetTpuProgramResponseExternal.proto:type_name -> tensorflow.tpu.GetTpuProgramResponseExternal.Blob
	2, // 1: tensorflow.tpu.GetTpuProgramResponseExternal.host_compute_metadata:type_name -> tensorflow.tf2xla.HostComputeMetadata
	1, // 2: tensorflow.tpu.GetTpuProgramResponseExternal.compiler_metadata:type_name -> tensorflow.tpu.GetTpuProgramResponseExternal.Blob
	3, // 3: tensorflow.tpu.TpuCompilationCacheServiceExternal.GetTpuProgram:input_type -> tensorflow.tpu.GetTpuProgramRequest
	0, // 4: tensorflow.tpu.TpuCompilationCacheServiceExternal.GetTpuProgram:output_type -> tensorflow.tpu.GetTpuProgramResponseExternal
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_init() }
func file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_init() {
	if File_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto != nil {
		return
	}
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTpuProgramResponseExternal); i {
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
		file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTpuProgramResponseExternal_Blob); i {
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
			RawDescriptor: file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_msgTypes,
	}.Build()
	File_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto = out.File
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_rawDesc = nil
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_goTypes = nil
	file_tensorflow_core_tpu_kernels_tpu_compilation_cache_proto_depIdxs = nil
}
