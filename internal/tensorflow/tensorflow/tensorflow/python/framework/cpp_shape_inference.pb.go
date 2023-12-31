// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/python/framework/cpp_shape_inference.proto

package framework

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

// LINT.IfChange
type CppShapeInferenceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape      *framework.TensorShapeProto         `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	HandleData *CppShapeInferenceResult_HandleData `protobuf:"bytes,4,opt,name=handle_data,json=handleData,proto3" json:"handle_data,omitempty"`
}

func (x *CppShapeInferenceResult) Reset() {
	*x = CppShapeInferenceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CppShapeInferenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CppShapeInferenceResult) ProtoMessage() {}

func (x *CppShapeInferenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CppShapeInferenceResult.ProtoReflect.Descriptor instead.
func (*CppShapeInferenceResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescGZIP(), []int{0}
}

func (x *CppShapeInferenceResult) GetShape() *framework.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *CppShapeInferenceResult) GetHandleData() *CppShapeInferenceResult_HandleData {
	if x != nil {
		return x.HandleData
	}
	return nil
}

type CppShapeInferenceInputsNeeded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputTensorsNeeded         []int32 `protobuf:"varint,1,rep,packed,name=input_tensors_needed,json=inputTensorsNeeded,proto3" json:"input_tensors_needed,omitempty"`
	InputTensorsAsShapesNeeded []int32 `protobuf:"varint,2,rep,packed,name=input_tensors_as_shapes_needed,json=inputTensorsAsShapesNeeded,proto3" json:"input_tensors_as_shapes_needed,omitempty"`
}

func (x *CppShapeInferenceInputsNeeded) Reset() {
	*x = CppShapeInferenceInputsNeeded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CppShapeInferenceInputsNeeded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CppShapeInferenceInputsNeeded) ProtoMessage() {}

func (x *CppShapeInferenceInputsNeeded) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CppShapeInferenceInputsNeeded.ProtoReflect.Descriptor instead.
func (*CppShapeInferenceInputsNeeded) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescGZIP(), []int{1}
}

func (x *CppShapeInferenceInputsNeeded) GetInputTensorsNeeded() []int32 {
	if x != nil {
		return x.InputTensorsNeeded
	}
	return nil
}

func (x *CppShapeInferenceInputsNeeded) GetInputTensorsAsShapesNeeded() []int32 {
	if x != nil {
		return x.InputTensorsAsShapesNeeded
	}
	return nil
}

type CppShapeInferenceResult_HandleShapeAndType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape *framework.TensorShapeProto `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype framework.DataType          `protobuf:"varint,2,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Type  *framework.FullTypeDef      `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CppShapeInferenceResult_HandleShapeAndType) Reset() {
	*x = CppShapeInferenceResult_HandleShapeAndType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CppShapeInferenceResult_HandleShapeAndType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CppShapeInferenceResult_HandleShapeAndType) ProtoMessage() {}

func (x *CppShapeInferenceResult_HandleShapeAndType) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CppShapeInferenceResult_HandleShapeAndType.ProtoReflect.Descriptor instead.
func (*CppShapeInferenceResult_HandleShapeAndType) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CppShapeInferenceResult_HandleShapeAndType) GetShape() *framework.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *CppShapeInferenceResult_HandleShapeAndType) GetDtype() framework.DataType {
	if x != nil {
		return x.Dtype
	}
	return framework.DataType(0)
}

func (x *CppShapeInferenceResult_HandleShapeAndType) GetType() *framework.FullTypeDef {
	if x != nil {
		return x.Type
	}
	return nil
}

type CppShapeInferenceResult_HandleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSet bool `protobuf:"varint,1,opt,name=is_set,json=isSet,proto3" json:"is_set,omitempty"`
	// Only valid if <is_set>.
	ShapeAndType []*CppShapeInferenceResult_HandleShapeAndType `protobuf:"bytes,2,rep,name=shape_and_type,json=shapeAndType,proto3" json:"shape_and_type,omitempty"`
}

func (x *CppShapeInferenceResult_HandleData) Reset() {
	*x = CppShapeInferenceResult_HandleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CppShapeInferenceResult_HandleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CppShapeInferenceResult_HandleData) ProtoMessage() {}

func (x *CppShapeInferenceResult_HandleData) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CppShapeInferenceResult_HandleData.ProtoReflect.Descriptor instead.
func (*CppShapeInferenceResult_HandleData) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CppShapeInferenceResult_HandleData) GetIsSet() bool {
	if x != nil {
		return x.IsSet
	}
	return false
}

func (x *CppShapeInferenceResult_HandleData) GetShapeAndType() []*CppShapeInferenceResult_HandleShapeAndType {
	if x != nil {
		return x.ShapeAndType
	}
	return nil
}

var File_tensorflow_python_framework_cpp_shape_inference_proto protoreflect.FileDescriptor

var file_tensorflow_python_framework_cpp_shape_inference_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74,
	0x68, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x70,
	0x70, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x1a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a, 0x17, 0x43, 0x70, 0x70, 0x53, 0x68, 0x61, 0x70, 0x65,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x32, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x70, 0x70, 0x53, 0x68, 0x61, 0x70, 0x65, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0xa7, 0x01, 0x0a, 0x12, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x65, 0x66, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x1a, 0x81,
	0x01, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x53, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0e, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x61, 0x6e,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x70, 0x70, 0x53, 0x68, 0x61,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x70, 0x65, 0x41, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x95,
	0x01, 0x0a, 0x1d, 0x43, 0x70, 0x70, 0x53, 0x68, 0x61, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x73, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x4e, 0x65, 0x65, 0x64,
	0x65, 0x64, 0x12, 0x42, 0x0a, 0x1e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x73, 0x5f, 0x6e, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x1a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x41, 0x73, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73,
	0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x42, 0xbb, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x16, 0x43, 0x70, 0x70, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f,
	0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02,
	0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescOnce sync.Once
	file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescData = file_tensorflow_python_framework_cpp_shape_inference_proto_rawDesc
)

func file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescGZIP() []byte {
	file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescOnce.Do(func() {
		file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescData)
	})
	return file_tensorflow_python_framework_cpp_shape_inference_proto_rawDescData
}

var file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_python_framework_cpp_shape_inference_proto_goTypes = []interface{}{
	(*CppShapeInferenceResult)(nil),                    // 0: tensorflow.CppShapeInferenceResult
	(*CppShapeInferenceInputsNeeded)(nil),              // 1: tensorflow.CppShapeInferenceInputsNeeded
	(*CppShapeInferenceResult_HandleShapeAndType)(nil), // 2: tensorflow.CppShapeInferenceResult.HandleShapeAndType
	(*CppShapeInferenceResult_HandleData)(nil),         // 3: tensorflow.CppShapeInferenceResult.HandleData
	(*framework.TensorShapeProto)(nil),                 // 4: tensorflow.TensorShapeProto
	(framework.DataType)(0),                            // 5: tensorflow.DataType
	(*framework.FullTypeDef)(nil),                      // 6: tensorflow.FullTypeDef
}
var file_tensorflow_python_framework_cpp_shape_inference_proto_depIdxs = []int32{
	4, // 0: tensorflow.CppShapeInferenceResult.shape:type_name -> tensorflow.TensorShapeProto
	3, // 1: tensorflow.CppShapeInferenceResult.handle_data:type_name -> tensorflow.CppShapeInferenceResult.HandleData
	4, // 2: tensorflow.CppShapeInferenceResult.HandleShapeAndType.shape:type_name -> tensorflow.TensorShapeProto
	5, // 3: tensorflow.CppShapeInferenceResult.HandleShapeAndType.dtype:type_name -> tensorflow.DataType
	6, // 4: tensorflow.CppShapeInferenceResult.HandleShapeAndType.type:type_name -> tensorflow.FullTypeDef
	2, // 5: tensorflow.CppShapeInferenceResult.HandleData.shape_and_type:type_name -> tensorflow.CppShapeInferenceResult.HandleShapeAndType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_python_framework_cpp_shape_inference_proto_init() }
func file_tensorflow_python_framework_cpp_shape_inference_proto_init() {
	if File_tensorflow_python_framework_cpp_shape_inference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CppShapeInferenceResult); i {
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
		file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CppShapeInferenceInputsNeeded); i {
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
		file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CppShapeInferenceResult_HandleShapeAndType); i {
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
		file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CppShapeInferenceResult_HandleData); i {
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
			RawDescriptor: file_tensorflow_python_framework_cpp_shape_inference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_python_framework_cpp_shape_inference_proto_goTypes,
		DependencyIndexes: file_tensorflow_python_framework_cpp_shape_inference_proto_depIdxs,
		MessageInfos:      file_tensorflow_python_framework_cpp_shape_inference_proto_msgTypes,
	}.Build()
	File_tensorflow_python_framework_cpp_shape_inference_proto = out.File
	file_tensorflow_python_framework_cpp_shape_inference_proto_rawDesc = nil
	file_tensorflow_python_framework_cpp_shape_inference_proto_goTypes = nil
	file_tensorflow_python_framework_cpp_shape_inference_proto_depIdxs = nil
}
