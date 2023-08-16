// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/tf2xla/kernels/callback.proto

package kernels

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	framework "grpc-client/internal/tensorflow/tensorflow/tensorflow/core/framework"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TfCallbackData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op      *framework.NodeDef                        `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Inputs  []*TfCallbackData_InputBufferDescription  `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	Outputs []*TfCallbackData_OutputBufferDescription `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty"`
}

func (x *TfCallbackData) Reset() {
	*x = TfCallbackData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfCallbackData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfCallbackData) ProtoMessage() {}

func (x *TfCallbackData) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfCallbackData.ProtoReflect.Descriptor instead.
func (*TfCallbackData) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescGZIP(), []int{0}
}

func (x *TfCallbackData) GetOp() *framework.NodeDef {
	if x != nil {
		return x.Op
	}
	return nil
}

func (x *TfCallbackData) GetInputs() []*TfCallbackData_InputBufferDescription {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *TfCallbackData) GetOutputs() []*TfCallbackData_OutputBufferDescription {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type TfCallbackData_BufferDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape *framework.TensorShapeProto `protobuf:"bytes,1,opt,name=shape" json:"shape,omitempty"`
	Type  *framework.DataType         `protobuf:"varint,2,opt,name=type,enum=tensorflow.DataType" json:"type,omitempty"`
}

func (x *TfCallbackData_BufferDescription) Reset() {
	*x = TfCallbackData_BufferDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfCallbackData_BufferDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfCallbackData_BufferDescription) ProtoMessage() {}

func (x *TfCallbackData_BufferDescription) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfCallbackData_BufferDescription.ProtoReflect.Descriptor instead.
func (*TfCallbackData_BufferDescription) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TfCallbackData_BufferDescription) GetShape() *framework.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *TfCallbackData_BufferDescription) GetType() framework.DataType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return framework.DataType(0)
}

type TfCallbackData_InputBufferDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BufferDescription *TfCallbackData_BufferDescription `protobuf:"bytes,1,opt,name=buffer_description,json=bufferDescription" json:"buffer_description,omitempty"`
	// The input value might be already fixed at the compilation time.
	// This value may or may not be present.
	Value *framework.TensorProto `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *TfCallbackData_InputBufferDescription) Reset() {
	*x = TfCallbackData_InputBufferDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfCallbackData_InputBufferDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfCallbackData_InputBufferDescription) ProtoMessage() {}

func (x *TfCallbackData_InputBufferDescription) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfCallbackData_InputBufferDescription.ProtoReflect.Descriptor instead.
func (*TfCallbackData_InputBufferDescription) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TfCallbackData_InputBufferDescription) GetBufferDescription() *TfCallbackData_BufferDescription {
	if x != nil {
		return x.BufferDescription
	}
	return nil
}

func (x *TfCallbackData_InputBufferDescription) GetValue() *framework.TensorProto {
	if x != nil {
		return x.Value
	}
	return nil
}

type TfCallbackData_OutputBufferDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BufferDescription *TfCallbackData_BufferDescription `protobuf:"bytes,1,opt,name=buffer_description,json=bufferDescription" json:"buffer_description,omitempty"`
	// Whether the buffer stores dynamically padded data: in that case, actual
	// concrete dimensions need to be stored after the buffer.
	IsDynamicallyPadded *bool `protobuf:"varint,2,opt,name=is_dynamically_padded,json=isDynamicallyPadded" json:"is_dynamically_padded,omitempty"`
}

func (x *TfCallbackData_OutputBufferDescription) Reset() {
	*x = TfCallbackData_OutputBufferDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfCallbackData_OutputBufferDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfCallbackData_OutputBufferDescription) ProtoMessage() {}

func (x *TfCallbackData_OutputBufferDescription) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfCallbackData_OutputBufferDescription.ProtoReflect.Descriptor instead.
func (*TfCallbackData_OutputBufferDescription) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescGZIP(), []int{0, 2}
}

func (x *TfCallbackData_OutputBufferDescription) GetBufferDescription() *TfCallbackData_BufferDescription {
	if x != nil {
		return x.BufferDescription
	}
	return nil
}

func (x *TfCallbackData_OutputBufferDescription) GetIsDynamicallyPadded() bool {
	if x != nil && x.IsDynamicallyPadded != nil {
		return *x.IsDynamicallyPadded
	}
	return false
}

var File_tensorflow_compiler_tf2xla_kernels_callback_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x2f, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a,
	0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x05, 0x0a, 0x0e, 0x54, 0x66, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x49,
	0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x66, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x66, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x71, 0x0a, 0x11, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xa4, 0x01, 0x0a, 0x16, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x12, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54,
	0x66, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x11, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0xaa, 0x01, 0x0a, 0x17, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a,
	0x12, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x66, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x73,
	0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x70, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x50, 0x61, 0x64, 0x64, 0x65, 0x64, 0x42, 0xb6,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x42, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x73, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
}

var (
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescData = file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDesc
)

func file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescData)
	})
	return file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDescData
}

var file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_compiler_tf2xla_kernels_callback_proto_goTypes = []interface{}{
	(*TfCallbackData)(nil),                         // 0: tensorflow.TfCallbackData
	(*TfCallbackData_BufferDescription)(nil),       // 1: tensorflow.TfCallbackData.BufferDescription
	(*TfCallbackData_InputBufferDescription)(nil),  // 2: tensorflow.TfCallbackData.InputBufferDescription
	(*TfCallbackData_OutputBufferDescription)(nil), // 3: tensorflow.TfCallbackData.OutputBufferDescription
	(*framework.NodeDef)(nil),                      // 4: tensorflow.NodeDef
	(*framework.TensorShapeProto)(nil),             // 5: tensorflow.TensorShapeProto
	(framework.DataType)(0),                        // 6: tensorflow.DataType
	(*framework.TensorProto)(nil),                  // 7: tensorflow.TensorProto
}
var file_tensorflow_compiler_tf2xla_kernels_callback_proto_depIdxs = []int32{
	4, // 0: tensorflow.TfCallbackData.op:type_name -> tensorflow.NodeDef
	2, // 1: tensorflow.TfCallbackData.inputs:type_name -> tensorflow.TfCallbackData.InputBufferDescription
	3, // 2: tensorflow.TfCallbackData.outputs:type_name -> tensorflow.TfCallbackData.OutputBufferDescription
	5, // 3: tensorflow.TfCallbackData.BufferDescription.shape:type_name -> tensorflow.TensorShapeProto
	6, // 4: tensorflow.TfCallbackData.BufferDescription.type:type_name -> tensorflow.DataType
	1, // 5: tensorflow.TfCallbackData.InputBufferDescription.buffer_description:type_name -> tensorflow.TfCallbackData.BufferDescription
	7, // 6: tensorflow.TfCallbackData.InputBufferDescription.value:type_name -> tensorflow.TensorProto
	1, // 7: tensorflow.TfCallbackData.OutputBufferDescription.buffer_description:type_name -> tensorflow.TfCallbackData.BufferDescription
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_tf2xla_kernels_callback_proto_init() }
func file_tensorflow_compiler_tf2xla_kernels_callback_proto_init() {
	if File_tensorflow_compiler_tf2xla_kernels_callback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfCallbackData); i {
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
		file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfCallbackData_BufferDescription); i {
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
		file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfCallbackData_InputBufferDescription); i {
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
		file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfCallbackData_OutputBufferDescription); i {
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
			RawDescriptor: file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_tf2xla_kernels_callback_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_tf2xla_kernels_callback_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_tf2xla_kernels_callback_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_tf2xla_kernels_callback_proto = out.File
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_rawDesc = nil
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_goTypes = nil
	file_tensorflow_compiler_tf2xla_kernels_callback_proto_depIdxs = nil
}