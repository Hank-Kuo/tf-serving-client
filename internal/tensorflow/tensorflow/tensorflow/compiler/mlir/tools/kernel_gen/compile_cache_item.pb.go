// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/mlir/tools/kernel_gen/compile_cache_item.proto

package kernel_gen

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

// Protocolbuffer representing a compilation input and output. This will be used
// for caching JIT compiles of kernels.
type CompilationCacheItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalModule string `protobuf:"bytes,1,opt,name=original_module,json=originalModule,proto3" json:"original_module,omitempty"`
	ResultModule   string `protobuf:"bytes,2,opt,name=result_module,json=resultModule,proto3" json:"result_module,omitempty"`
}

func (x *CompilationCacheItem) Reset() {
	*x = CompilationCacheItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationCacheItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationCacheItem) ProtoMessage() {}

func (x *CompilationCacheItem) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationCacheItem.ProtoReflect.Descriptor instead.
func (*CompilationCacheItem) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescGZIP(), []int{0}
}

func (x *CompilationCacheItem) GetOriginalModule() string {
	if x != nil {
		return x.OriginalModule
	}
	return ""
}

func (x *CompilationCacheItem) GetResultModule() string {
	if x != nil {
		return x.ResultModule
	}
	return ""
}

var File_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDesc = []byte{
	0x0a, 0x42, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x22, 0x64, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x42, 0xde, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f,
	0x67, 0x65, 0x6e, 0x42, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69,
	0x72, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x67,
	0x65, 0x6e, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x4d, 0x4b, 0x58, 0xaa, 0x02, 0x0e, 0x4d, 0x6c,
	0x69, 0x72, 0x2e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x6e, 0xca, 0x02, 0x0e, 0x4d,
	0x6c, 0x69, 0x72, 0x5c, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x6e, 0xe2, 0x02, 0x1a,
	0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x6e, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4d, 0x6c, 0x69,
	0x72, 0x3a, 0x3a, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescData = file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDesc
)

func file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescData)
	})
	return file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDescData
}

var file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_goTypes = []interface{}{
	(*CompilationCacheItem)(nil), // 0: mlir.kernel_gen.CompilationCacheItem
}
var file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_init() }
func file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_init() {
	if File_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationCacheItem); i {
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
			RawDescriptor: file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto = out.File
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_rawDesc = nil
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_goTypes = nil
	file_tensorflow_compiler_mlir_tools_kernel_gen_compile_cache_item_proto_depIdxs = nil
}
