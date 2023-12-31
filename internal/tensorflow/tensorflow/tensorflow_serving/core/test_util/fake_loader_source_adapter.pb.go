// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow_serving/core/test_util/fake_loader_source_adapter.proto

package test_util

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

// Config proto for FakeLoaderSourceAdapter.
type FakeLoaderSourceAdapterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FakeLoaderSourceAdapter's 'suffix' ctor parameter.
	Suffix string `protobuf:"bytes,1,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *FakeLoaderSourceAdapterConfig) Reset() {
	*x = FakeLoaderSourceAdapterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeLoaderSourceAdapterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeLoaderSourceAdapterConfig) ProtoMessage() {}

func (x *FakeLoaderSourceAdapterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeLoaderSourceAdapterConfig.ProtoReflect.Descriptor instead.
func (*FakeLoaderSourceAdapterConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *FakeLoaderSourceAdapterConfig) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

var File_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto protoreflect.FileDescriptor

var file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDesc = []byte{
	0x0a, 0x42, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x66, 0x61, 0x6b, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x74,
	0x69, 0x6c, 0x22, 0x37, 0x0a, 0x1d, 0x46, 0x61, 0x6b, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x42, 0x9c, 0x02, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x74, 0x69, 0x6c,
	0x42, 0x1c, 0x46, 0x61, 0x6b, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0xa2, 0x02,
	0x03, 0x54, 0x53, 0x54, 0xaa, 0x02, 0x1b, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x74,
	0x69, 0x6c, 0xca, 0x02, 0x1b, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x55, 0x74, 0x69, 0x6c,
	0xe2, 0x02, 0x27, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescOnce sync.Once
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescData = file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDesc
)

func file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescData)
	})
	return file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDescData
}

var file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_goTypes = []interface{}{
	(*FakeLoaderSourceAdapterConfig)(nil), // 0: tensorflow.serving.test_util.FakeLoaderSourceAdapterConfig
}
var file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_init() }
func file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_init() {
	if File_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeLoaderSourceAdapterConfig); i {
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
			RawDescriptor: file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto = out.File
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_rawDesc = nil
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_goTypes = nil
	file_tensorflow_serving_core_test_util_fake_loader_source_adapter_proto_depIdxs = nil
}
