// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/service/metrics.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines compilation stages for which metrics are collected.
type CompilationLogEntry_CompilationStage int32

const (
	CompilationLogEntry_UNSPECIFIED     CompilationLogEntry_CompilationStage = 0
	CompilationLogEntry_END_TO_END      CompilationLogEntry_CompilationStage = 1
	CompilationLogEntry_HLO_PASSES      CompilationLogEntry_CompilationStage = 2
	CompilationLogEntry_CODE_GENERATION CompilationLogEntry_CompilationStage = 3
	CompilationLogEntry_BACKEND_PASSES  CompilationLogEntry_CompilationStage = 4
)

// Enum value maps for CompilationLogEntry_CompilationStage.
var (
	CompilationLogEntry_CompilationStage_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "END_TO_END",
		2: "HLO_PASSES",
		3: "CODE_GENERATION",
		4: "BACKEND_PASSES",
	}
	CompilationLogEntry_CompilationStage_value = map[string]int32{
		"UNSPECIFIED":     0,
		"END_TO_END":      1,
		"HLO_PASSES":      2,
		"CODE_GENERATION": 3,
		"BACKEND_PASSES":  4,
	}
)

func (x CompilationLogEntry_CompilationStage) Enum() *CompilationLogEntry_CompilationStage {
	p := new(CompilationLogEntry_CompilationStage)
	*p = x
	return p
}

func (x CompilationLogEntry_CompilationStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompilationLogEntry_CompilationStage) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_xla_service_metrics_proto_enumTypes[0].Descriptor()
}

func (CompilationLogEntry_CompilationStage) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_xla_service_metrics_proto_enumTypes[0]
}

func (x CompilationLogEntry_CompilationStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompilationLogEntry_CompilationStage.Descriptor instead.
func (CompilationLogEntry_CompilationStage) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_metrics_proto_rawDescGZIP(), []int{1, 0}
}

// Defines pass specific metrics.
type PassMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the module on which the pass was run.
	ModuleId int64 `protobuf:"varint,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	// The name of the pass.
	PassName string `protobuf:"bytes,2,opt,name=pass_name,json=passName,proto3" json:"pass_name,omitempty"`
	// Duration of the pass.
	PassDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=pass_duration,json=passDuration,proto3" json:"pass_duration,omitempty"`
}

func (x *PassMetrics) Reset() {
	*x = PassMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassMetrics) ProtoMessage() {}

func (x *PassMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassMetrics.ProtoReflect.Descriptor instead.
func (*PassMetrics) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *PassMetrics) GetModuleId() int64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *PassMetrics) GetPassName() string {
	if x != nil {
		return x.PassName
	}
	return ""
}

func (x *PassMetrics) GetPassDuration() *durationpb.Duration {
	if x != nil {
		return x.PassDuration
	}
	return nil
}

// Defines XLA compilation metrics.
type CompilationLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time when the event captured by this log entry occurred.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Compilation stage recorded by this log entry.
	Stage CompilationLogEntry_CompilationStage `protobuf:"varint,2,opt,name=stage,proto3,enum=xla.CompilationLogEntry_CompilationStage" json:"stage,omitempty"`
	// Duration of the given compilation stage.
	Duration *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// Task index from which this log entry was recorded or
	// -1 if the task index could not be fetched.
	TaskIndex int32 `protobuf:"varint,4,opt,name=task_index,json=taskIndex,proto3" json:"task_index,omitempty"`
	// Pass specific metrics.
	PassMetrics []*PassMetrics `protobuf:"bytes,5,rep,name=pass_metrics,json=passMetrics,proto3" json:"pass_metrics,omitempty"`
}

func (x *CompilationLogEntry) Reset() {
	*x = CompilationLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationLogEntry) ProtoMessage() {}

func (x *CompilationLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationLogEntry.ProtoReflect.Descriptor instead.
func (*CompilationLogEntry) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *CompilationLogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CompilationLogEntry) GetStage() CompilationLogEntry_CompilationStage {
	if x != nil {
		return x.Stage
	}
	return CompilationLogEntry_UNSPECIFIED
}

func (x *CompilationLogEntry) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *CompilationLogEntry) GetTaskIndex() int32 {
	if x != nil {
		return x.TaskIndex
	}
	return 0
}

func (x *CompilationLogEntry) GetPassMetrics() []*PassMetrics {
	if x != nil {
		return x.PassMetrics
	}
	return nil
}

var File_tensorflow_compiler_xla_service_metrics_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_service_metrics_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x78, 0x6c, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x89, 0x03, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x6c, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x4e, 0x44,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x4c, 0x4f, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x53,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x41, 0x43, 0x4b, 0x45,
	0x4e, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x53, 0x10, 0x04, 0x42, 0x8f, 0x01, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x42, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x58, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xca,
	0x02, 0x03, 0x58, 0x6c, 0x61, 0xe2, 0x02, 0x0f, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x58, 0x6c, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_service_metrics_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_service_metrics_proto_rawDescData = file_tensorflow_compiler_xla_service_metrics_proto_rawDesc
)

func file_tensorflow_compiler_xla_service_metrics_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_service_metrics_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_service_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_service_metrics_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_service_metrics_proto_rawDescData
}

var file_tensorflow_compiler_xla_service_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_xla_service_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_compiler_xla_service_metrics_proto_goTypes = []interface{}{
	(CompilationLogEntry_CompilationStage)(0), // 0: xla.CompilationLogEntry.CompilationStage
	(*PassMetrics)(nil),                       // 1: xla.PassMetrics
	(*CompilationLogEntry)(nil),               // 2: xla.CompilationLogEntry
	(*durationpb.Duration)(nil),               // 3: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),             // 4: google.protobuf.Timestamp
}
var file_tensorflow_compiler_xla_service_metrics_proto_depIdxs = []int32{
	3, // 0: xla.PassMetrics.pass_duration:type_name -> google.protobuf.Duration
	4, // 1: xla.CompilationLogEntry.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: xla.CompilationLogEntry.stage:type_name -> xla.CompilationLogEntry.CompilationStage
	3, // 3: xla.CompilationLogEntry.duration:type_name -> google.protobuf.Duration
	1, // 4: xla.CompilationLogEntry.pass_metrics:type_name -> xla.PassMetrics
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_service_metrics_proto_init() }
func file_tensorflow_compiler_xla_service_metrics_proto_init() {
	if File_tensorflow_compiler_xla_service_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassMetrics); i {
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
		file_tensorflow_compiler_xla_service_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationLogEntry); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_service_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_service_metrics_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_service_metrics_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_xla_service_metrics_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_xla_service_metrics_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_service_metrics_proto = out.File
	file_tensorflow_compiler_xla_service_metrics_proto_rawDesc = nil
	file_tensorflow_compiler_xla_service_metrics_proto_goTypes = nil
	file_tensorflow_compiler_xla_service_metrics_proto_depIdxs = nil
}
