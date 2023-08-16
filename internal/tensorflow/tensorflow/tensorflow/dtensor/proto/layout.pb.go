// Copyright 2022 The TensorFlow Authors. All Rights Reserved.
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
// source: tensorflow/dtensor/proto/layout.proto

package proto

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

type LayoutProto_LayoutType int32

const (
	LayoutProto_UNKNOWN       LayoutProto_LayoutType = 0
	LayoutProto_STATIC        LayoutProto_LayoutType = 1 // A tiled layout for static and evenly shaped local shards.
	LayoutProto_PARTED        LayoutProto_LayoutType = 2 // A parted layout contains axes that are treated as
	LayoutProto_SINGLE_DEVICE LayoutProto_LayoutType = 3 // A single device layout that represents DTensor on a
)

// Enum value maps for LayoutProto_LayoutType.
var (
	LayoutProto_LayoutType_name = map[int32]string{
		0: "UNKNOWN",
		1: "STATIC",
		2: "PARTED",
		3: "SINGLE_DEVICE",
	}
	LayoutProto_LayoutType_value = map[string]int32{
		"UNKNOWN":       0,
		"STATIC":        1,
		"PARTED":        2,
		"SINGLE_DEVICE": 3,
	}
)

func (x LayoutProto_LayoutType) Enum() *LayoutProto_LayoutType {
	p := new(LayoutProto_LayoutType)
	*p = x
	return p
}

func (x LayoutProto_LayoutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LayoutProto_LayoutType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_dtensor_proto_layout_proto_enumTypes[0].Descriptor()
}

func (LayoutProto_LayoutType) Type() protoreflect.EnumType {
	return &file_tensorflow_dtensor_proto_layout_proto_enumTypes[0]
}

func (x LayoutProto_LayoutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LayoutProto_LayoutType.Descriptor instead.
func (LayoutProto_LayoutType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP(), []int{2, 0}
}

// Defines a sharding spec over a tensor.
// A sharding spec can either be a mesh dimension of the associated mesh or the
// special values scalar, or not_sharded.
type ShardingSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the mesh dimension this tensor dimension is sharded over.
	ShardingSpec string `protobuf:"bytes,2,opt,name=sharding_spec,json=shardingSpec,proto3" json:"sharding_spec,omitempty"`
}

func (x *ShardingSpec) Reset() {
	*x = ShardingSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardingSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardingSpec) ProtoMessage() {}

func (x *ShardingSpec) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardingSpec.ProtoReflect.Descriptor instead.
func (*ShardingSpec) Descriptor() ([]byte, []int) {
	return file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP(), []int{0}
}

func (x *ShardingSpec) GetShardingSpec() string {
	if x != nil {
		return x.ShardingSpec
	}
	return ""
}

type MeshDimensionProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of mesh dimension, required to create Mesh.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Length of mesh dimension.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *MeshDimensionProto) Reset() {
	*x = MeshDimensionProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshDimensionProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshDimensionProto) ProtoMessage() {}

func (x *MeshDimensionProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshDimensionProto.ProtoReflect.Descriptor instead.
func (*MeshDimensionProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP(), []int{1}
}

func (x *MeshDimensionProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeshDimensionProto) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// Proto representation of a Layout.
type LayoutProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardingSpecs []*ShardingSpec        `protobuf:"bytes,1,rep,name=sharding_specs,json=shardingSpecs,proto3" json:"sharding_specs,omitempty"`
	MeshConfig    *MeshProto             `protobuf:"bytes,2,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	Type          LayoutProto_LayoutType `protobuf:"varint,3,opt,name=type,proto3,enum=tensorflow.dtensor.LayoutProto_LayoutType" json:"type,omitempty"`
}

func (x *LayoutProto) Reset() {
	*x = LayoutProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayoutProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayoutProto) ProtoMessage() {}

func (x *LayoutProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayoutProto.ProtoReflect.Descriptor instead.
func (*LayoutProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP(), []int{2}
}

func (x *LayoutProto) GetShardingSpecs() []*ShardingSpec {
	if x != nil {
		return x.ShardingSpecs
	}
	return nil
}

func (x *LayoutProto) GetMeshConfig() *MeshProto {
	if x != nil {
		return x.MeshConfig
	}
	return nil
}

func (x *LayoutProto) GetType() LayoutProto_LayoutType {
	if x != nil {
		return x.Type
	}
	return LayoutProto_UNKNOWN
}

// Proto representation of a Mesh.
// TODO(allenl): Add a unique mesh ID so we have an efficient way to identify
// meshes in mappings.
type MeshProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeshDimensions []*MeshDimensionProto `protobuf:"bytes,1,rep,name=mesh_dimensions,json=meshDimensions,proto3" json:"mesh_dimensions,omitempty"`
	// A list of global devices ids.
	GlobalDeviceIds []int64 `protobuf:"varint,2,rep,packed,name=global_device_ids,json=globalDeviceIds,proto3" json:"global_device_ids,omitempty"`
	// Devices ids handled by the local process.
	LocalDeviceIds []int64 `protobuf:"varint,4,rep,packed,name=local_device_ids,json=localDeviceIds,proto3" json:"local_device_ids,omitempty"`
	// Devices handled by the local process.
	LocalDevices []string `protobuf:"bytes,5,rep,name=local_devices,json=localDevices,proto3" json:"local_devices,omitempty"`
	// Global device names (Optional). Set when multiple device meshes are in use.
	GlobalDevices []string `protobuf:"bytes,6,rep,name=global_devices,json=globalDevices,proto3" json:"global_devices,omitempty"`
	// Required name which uniquely identifies this mesh in a program.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// If true, ops on this mesh will use XLA SPMD.
	UseXlaSpmd bool `protobuf:"varint,7,opt,name=use_xla_spmd,json=useXlaSpmd,proto3" json:"use_xla_spmd,omitempty"`
	// The device string when the mesh is a single device mesh. If it is not
	// empty, then global_device_ids, local_device_ids, local_device and global
	// devices are all empty.
	SingleDevice string `protobuf:"bytes,8,opt,name=single_device,json=singleDevice,proto3" json:"single_device,omitempty"`
}

func (x *MeshProto) Reset() {
	*x = MeshProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshProto) ProtoMessage() {}

func (x *MeshProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_dtensor_proto_layout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshProto.ProtoReflect.Descriptor instead.
func (*MeshProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP(), []int{3}
}

func (x *MeshProto) GetMeshDimensions() []*MeshDimensionProto {
	if x != nil {
		return x.MeshDimensions
	}
	return nil
}

func (x *MeshProto) GetGlobalDeviceIds() []int64 {
	if x != nil {
		return x.GlobalDeviceIds
	}
	return nil
}

func (x *MeshProto) GetLocalDeviceIds() []int64 {
	if x != nil {
		return x.LocalDeviceIds
	}
	return nil
}

func (x *MeshProto) GetLocalDevices() []string {
	if x != nil {
		return x.LocalDevices
	}
	return nil
}

func (x *MeshProto) GetGlobalDevices() []string {
	if x != nil {
		return x.GlobalDevices
	}
	return nil
}

func (x *MeshProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeshProto) GetUseXlaSpmd() bool {
	if x != nil {
		return x.UseXlaSpmd
	}
	return false
}

func (x *MeshProto) GetSingleDevice() string {
	if x != nil {
		return x.SingleDevice
	}
	return ""
}

var File_tensorflow_dtensor_proto_layout_proto protoreflect.FileDescriptor

var file_tensorflow_dtensor_proto_layout_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x64, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x39, 0x0a, 0x0c, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x3c, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x47, 0x0a, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0d,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x3e, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a,
	0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43,
	0x45, 0x10, 0x03, 0x22, 0xd9, 0x02, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x4f, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x73, 0x68, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x68, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f,
	0x78, 0x6c, 0x61, 0x5f, 0x73, 0x70, 0x6d, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x58, 0x6c, 0x61, 0x53, 0x70, 0x6d, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42,
	0xd3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x42, 0x0b, 0x4c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x64, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02,
	0x03, 0x54, 0x44, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x44, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xca, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xe2, 0x02,
	0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x44, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_dtensor_proto_layout_proto_rawDescOnce sync.Once
	file_tensorflow_dtensor_proto_layout_proto_rawDescData = file_tensorflow_dtensor_proto_layout_proto_rawDesc
)

func file_tensorflow_dtensor_proto_layout_proto_rawDescGZIP() []byte {
	file_tensorflow_dtensor_proto_layout_proto_rawDescOnce.Do(func() {
		file_tensorflow_dtensor_proto_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_dtensor_proto_layout_proto_rawDescData)
	})
	return file_tensorflow_dtensor_proto_layout_proto_rawDescData
}

var file_tensorflow_dtensor_proto_layout_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_dtensor_proto_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_dtensor_proto_layout_proto_goTypes = []interface{}{
	(LayoutProto_LayoutType)(0), // 0: tensorflow.dtensor.LayoutProto.LayoutType
	(*ShardingSpec)(nil),        // 1: tensorflow.dtensor.ShardingSpec
	(*MeshDimensionProto)(nil),  // 2: tensorflow.dtensor.MeshDimensionProto
	(*LayoutProto)(nil),         // 3: tensorflow.dtensor.LayoutProto
	(*MeshProto)(nil),           // 4: tensorflow.dtensor.MeshProto
}
var file_tensorflow_dtensor_proto_layout_proto_depIdxs = []int32{
	1, // 0: tensorflow.dtensor.LayoutProto.sharding_specs:type_name -> tensorflow.dtensor.ShardingSpec
	4, // 1: tensorflow.dtensor.LayoutProto.mesh_config:type_name -> tensorflow.dtensor.MeshProto
	0, // 2: tensorflow.dtensor.LayoutProto.type:type_name -> tensorflow.dtensor.LayoutProto.LayoutType
	2, // 3: tensorflow.dtensor.MeshProto.mesh_dimensions:type_name -> tensorflow.dtensor.MeshDimensionProto
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_dtensor_proto_layout_proto_init() }
func file_tensorflow_dtensor_proto_layout_proto_init() {
	if File_tensorflow_dtensor_proto_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_dtensor_proto_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardingSpec); i {
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
		file_tensorflow_dtensor_proto_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshDimensionProto); i {
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
		file_tensorflow_dtensor_proto_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayoutProto); i {
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
		file_tensorflow_dtensor_proto_layout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshProto); i {
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
			RawDescriptor: file_tensorflow_dtensor_proto_layout_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_dtensor_proto_layout_proto_goTypes,
		DependencyIndexes: file_tensorflow_dtensor_proto_layout_proto_depIdxs,
		EnumInfos:         file_tensorflow_dtensor_proto_layout_proto_enumTypes,
		MessageInfos:      file_tensorflow_dtensor_proto_layout_proto_msgTypes,
	}.Build()
	File_tensorflow_dtensor_proto_layout_proto = out.File
	file_tensorflow_dtensor_proto_layout_proto_rawDesc = nil
	file_tensorflow_dtensor_proto_layout_proto_goTypes = nil
	file_tensorflow_dtensor_proto_layout_proto_depIdxs = nil
}