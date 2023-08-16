// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/framework/device_attributes.proto

package framework

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

type InterconnectLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId int32  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Strength int32  `protobuf:"varint,3,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (x *InterconnectLink) Reset() {
	*x = InterconnectLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterconnectLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterconnectLink) ProtoMessage() {}

func (x *InterconnectLink) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterconnectLink.ProtoReflect.Descriptor instead.
func (*InterconnectLink) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_device_attributes_proto_rawDescGZIP(), []int{0}
}

func (x *InterconnectLink) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *InterconnectLink) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InterconnectLink) GetStrength() int32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

type LocalLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link []*InterconnectLink `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
}

func (x *LocalLinks) Reset() {
	*x = LocalLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalLinks) ProtoMessage() {}

func (x *LocalLinks) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalLinks.ProtoReflect.Descriptor instead.
func (*LocalLinks) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_device_attributes_proto_rawDescGZIP(), []int{1}
}

func (x *LocalLinks) GetLink() []*InterconnectLink {
	if x != nil {
		return x.Link
	}
	return nil
}

type DeviceLocality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId,proto3" json:"bus_id,omitempty"`
	// Optional NUMA locality of device.
	NumaNode int32 `protobuf:"varint,2,opt,name=numa_node,json=numaNode,proto3" json:"numa_node,omitempty"`
	// Optional local interconnect links to other devices.
	Links *LocalLinks `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *DeviceLocality) Reset() {
	*x = DeviceLocality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceLocality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLocality) ProtoMessage() {}

func (x *DeviceLocality) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLocality.ProtoReflect.Descriptor instead.
func (*DeviceLocality) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_device_attributes_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceLocality) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *DeviceLocality) GetNumaNode() int32 {
	if x != nil {
		return x.NumaNode
	}
	return 0
}

func (x *DeviceLocality) GetLinks() *LocalLinks {
	if x != nil {
		return x.Links
	}
	return nil
}

type DeviceAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc string `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc,proto3" json:"physical_device_desc,omitempty"`
	// A physical device ID for use in XLA DeviceAssignments, unique across
	// clients in a multi-client setup. Set to -1 if unavailable, non-negative
	// otherwise.
	XlaGlobalId int64 `protobuf:"varint,8,opt,name=xla_global_id,json=xlaGlobalId,proto3" json:"xla_global_id,omitempty"`
}

func (x *DeviceAttributes) Reset() {
	*x = DeviceAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAttributes) ProtoMessage() {}

func (x *DeviceAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_device_attributes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceAttributes.ProtoReflect.Descriptor instead.
func (*DeviceAttributes) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_device_attributes_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceAttributes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceAttributes) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *DeviceAttributes) GetMemoryLimit() int64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *DeviceAttributes) GetLocality() *DeviceLocality {
	if x != nil {
		return x.Locality
	}
	return nil
}

func (x *DeviceAttributes) GetIncarnation() uint64 {
	if x != nil {
		return x.Incarnation
	}
	return 0
}

func (x *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if x != nil {
		return x.PhysicalDeviceDesc
	}
	return ""
}

func (x *DeviceAttributes) GetXlaGlobalId() int64 {
	if x != nil {
		return x.XlaGlobalId
	}
	return 0
}

var File_tensorflow_core_framework_device_attributes_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_device_attributes_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22,
	0x5f, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x3e, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x30,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x22, 0x72, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d,
	0x61, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75,
	0x6d, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x22, 0x9a, 0x02, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x63,
	0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0b,
	0x69, 0x6e, 0x63, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x70,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x22, 0x0a,
	0x0d, 0x78, 0x6c, 0x61, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x78, 0x6c, 0x61, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49,
	0x64, 0x42, 0xb8, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x42, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_device_attributes_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_device_attributes_proto_rawDescData = file_tensorflow_core_framework_device_attributes_proto_rawDesc
)

func file_tensorflow_core_framework_device_attributes_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_device_attributes_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_device_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_device_attributes_proto_rawDescData)
	})
	return file_tensorflow_core_framework_device_attributes_proto_rawDescData
}

var file_tensorflow_core_framework_device_attributes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_framework_device_attributes_proto_goTypes = []interface{}{
	(*InterconnectLink)(nil), // 0: tensorflow.InterconnectLink
	(*LocalLinks)(nil),       // 1: tensorflow.LocalLinks
	(*DeviceLocality)(nil),   // 2: tensorflow.DeviceLocality
	(*DeviceAttributes)(nil), // 3: tensorflow.DeviceAttributes
}
var file_tensorflow_core_framework_device_attributes_proto_depIdxs = []int32{
	0, // 0: tensorflow.LocalLinks.link:type_name -> tensorflow.InterconnectLink
	1, // 1: tensorflow.DeviceLocality.links:type_name -> tensorflow.LocalLinks
	2, // 2: tensorflow.DeviceAttributes.locality:type_name -> tensorflow.DeviceLocality
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_device_attributes_proto_init() }
func file_tensorflow_core_framework_device_attributes_proto_init() {
	if File_tensorflow_core_framework_device_attributes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_device_attributes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterconnectLink); i {
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
		file_tensorflow_core_framework_device_attributes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalLinks); i {
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
		file_tensorflow_core_framework_device_attributes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceLocality); i {
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
		file_tensorflow_core_framework_device_attributes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceAttributes); i {
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
			RawDescriptor: file_tensorflow_core_framework_device_attributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_device_attributes_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_device_attributes_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_device_attributes_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_device_attributes_proto = out.File
	file_tensorflow_core_framework_device_attributes_proto_rawDesc = nil
	file_tensorflow_core_framework_device_attributes_proto_goTypes = nil
	file_tensorflow_core_framework_device_attributes_proto_depIdxs = nil
}