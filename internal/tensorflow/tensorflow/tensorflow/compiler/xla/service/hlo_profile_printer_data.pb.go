// Copyright 2018 The TensorFlow Authors. All Rights Reserved.
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
// source: tensorflow/compiler/xla/service/hlo_profile_printer_data.proto

package service

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

// Describes how to pretty-print a profile counter array gathered for a specific
// HloModule.
type HloProfilePrinterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HloComputationInfos for every HloComputation in the HloModule.
	ComputationInfos []*HloProfilePrinterData_HloComputationInfo `protobuf:"bytes,1,rep,name=computation_infos,json=computationInfos,proto3" json:"computation_infos,omitempty"`
	// The size of the profile counters array we will pretty-print.
	ProfileCountersSize int64 `protobuf:"varint,2,opt,name=profile_counters_size,json=profileCountersSize,proto3" json:"profile_counters_size,omitempty"`
	// Maps extra metric name to the index into the profile counters array.
	ExtraMetrics map[string]int64 `protobuf:"bytes,3,rep,name=extra_metrics,json=extraMetrics,proto3" json:"extra_metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Name of the entry computation.
	EntryComputation string `protobuf:"bytes,4,opt,name=entry_computation,json=entryComputation,proto3" json:"entry_computation,omitempty"`
}

func (x *HloProfilePrinterData) Reset() {
	*x = HloProfilePrinterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HloProfilePrinterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HloProfilePrinterData) ProtoMessage() {}

func (x *HloProfilePrinterData) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HloProfilePrinterData.ProtoReflect.Descriptor instead.
func (*HloProfilePrinterData) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescGZIP(), []int{0}
}

func (x *HloProfilePrinterData) GetComputationInfos() []*HloProfilePrinterData_HloComputationInfo {
	if x != nil {
		return x.ComputationInfos
	}
	return nil
}

func (x *HloProfilePrinterData) GetProfileCountersSize() int64 {
	if x != nil {
		return x.ProfileCountersSize
	}
	return 0
}

func (x *HloProfilePrinterData) GetExtraMetrics() map[string]int64 {
	if x != nil {
		return x.ExtraMetrics
	}
	return nil
}

func (x *HloProfilePrinterData) GetEntryComputation() string {
	if x != nil {
		return x.EntryComputation
	}
	return ""
}

// Pretty-printer information about an HloInstruction.
type HloProfilePrinterData_HloInstructionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongName  string `protobuf:"bytes,1,opt,name=long_name,json=longName,proto3" json:"long_name,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	Category  string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// Metrics computed by HloCostAnalysis.
	FlopCount           float32 `protobuf:"fixed32,4,opt,name=flop_count,json=flopCount,proto3" json:"flop_count,omitempty"`
	TranscendentalCount float32 `protobuf:"fixed32,5,opt,name=transcendental_count,json=transcendentalCount,proto3" json:"transcendental_count,omitempty"`
	BytesAccessed       int64   `protobuf:"varint,9,opt,name=bytes_accessed,json=bytesAccessed,proto3" json:"bytes_accessed,omitempty"`
	OptimalSeconds      float32 `protobuf:"fixed32,7,opt,name=optimal_seconds,json=optimalSeconds,proto3" json:"optimal_seconds,omitempty"`
	// The index into the profile counters array for the HloInstruction
	// corresponding to this HloInstructionInfo.
	ProfileIndex int64 `protobuf:"varint,8,opt,name=profile_index,json=profileIndex,proto3" json:"profile_index,omitempty"`
}

func (x *HloProfilePrinterData_HloInstructionInfo) Reset() {
	*x = HloProfilePrinterData_HloInstructionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HloProfilePrinterData_HloInstructionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HloProfilePrinterData_HloInstructionInfo) ProtoMessage() {}

func (x *HloProfilePrinterData_HloInstructionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HloProfilePrinterData_HloInstructionInfo.ProtoReflect.Descriptor instead.
func (*HloProfilePrinterData_HloInstructionInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetLongName() string {
	if x != nil {
		return x.LongName
	}
	return ""
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetFlopCount() float32 {
	if x != nil {
		return x.FlopCount
	}
	return 0
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetTranscendentalCount() float32 {
	if x != nil {
		return x.TranscendentalCount
	}
	return 0
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetBytesAccessed() int64 {
	if x != nil {
		return x.BytesAccessed
	}
	return 0
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetOptimalSeconds() float32 {
	if x != nil {
		return x.OptimalSeconds
	}
	return 0
}

func (x *HloProfilePrinterData_HloInstructionInfo) GetProfileIndex() int64 {
	if x != nil {
		return x.ProfileIndex
	}
	return 0
}

// Pretty-printer information about an HloComputation.
type HloProfilePrinterData_HloComputationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index into the profile counters array for the HloComputation
	// corresponding to this HloComputationInfo.
	ProfileIndex int64 `protobuf:"varint,2,opt,name=profile_index,json=profileIndex,proto3" json:"profile_index,omitempty"`
	// HloInstructionInfos for every HloInstruction in the HloComputation for
	// corresponding to this HloComputattionInfo.
	InstructionInfos []*HloProfilePrinterData_HloInstructionInfo `protobuf:"bytes,3,rep,name=instruction_infos,json=instructionInfos,proto3" json:"instruction_infos,omitempty"`
}

func (x *HloProfilePrinterData_HloComputationInfo) Reset() {
	*x = HloProfilePrinterData_HloComputationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HloProfilePrinterData_HloComputationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HloProfilePrinterData_HloComputationInfo) ProtoMessage() {}

func (x *HloProfilePrinterData_HloComputationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HloProfilePrinterData_HloComputationInfo.ProtoReflect.Descriptor instead.
func (*HloProfilePrinterData_HloComputationInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HloProfilePrinterData_HloComputationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HloProfilePrinterData_HloComputationInfo) GetProfileIndex() int64 {
	if x != nil {
		return x.ProfileIndex
	}
	return 0
}

func (x *HloProfilePrinterData_HloComputationInfo) GetInstructionInfos() []*HloProfilePrinterData_HloInstructionInfo {
	if x != nil {
		return x.InstructionInfos
	}
	return nil
}

var File_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x68, 0x6c, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x78, 0x6c, 0x61, 0x22, 0xd0, 0x06, 0x0a, 0x15, 0x48, 0x6c, 0x6f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x5a, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x78, 0x6c, 0x61,
	0x2e, 0x48, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x6c, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x51, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x48, 0x6c, 0x6f,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0xb9, 0x02, 0x0a, 0x12, 0x48, 0x6c, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a,
	0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x1a, 0xa9, 0x01, 0x0a, 0x12,
	0x48, 0x6c, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x5a, 0x0a, 0x11, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x48, 0x6c, 0x6f,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x48, 0x6c, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xa0, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x2e, 0x78, 0x6c, 0x61, 0x42, 0x1a, 0x48, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4a, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x03, 0x58, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xca, 0x02, 0x03,
	0x58, 0x6c, 0x61, 0xe2, 0x02, 0x0f, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x58, 0x6c, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescData = file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDesc
)

func file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDescData
}

var file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_goTypes = []interface{}{
	(*HloProfilePrinterData)(nil),                    // 0: xla.HloProfilePrinterData
	(*HloProfilePrinterData_HloInstructionInfo)(nil), // 1: xla.HloProfilePrinterData.HloInstructionInfo
	(*HloProfilePrinterData_HloComputationInfo)(nil), // 2: xla.HloProfilePrinterData.HloComputationInfo
	nil, // 3: xla.HloProfilePrinterData.ExtraMetricsEntry
}
var file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_depIdxs = []int32{
	2, // 0: xla.HloProfilePrinterData.computation_infos:type_name -> xla.HloProfilePrinterData.HloComputationInfo
	3, // 1: xla.HloProfilePrinterData.extra_metrics:type_name -> xla.HloProfilePrinterData.ExtraMetricsEntry
	1, // 2: xla.HloProfilePrinterData.HloComputationInfo.instruction_infos:type_name -> xla.HloProfilePrinterData.HloInstructionInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_init() }
func file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_init() {
	if File_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HloProfilePrinterData); i {
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
		file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HloProfilePrinterData_HloInstructionInfo); i {
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
		file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HloProfilePrinterData_HloComputationInfo); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto = out.File
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_rawDesc = nil
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_goTypes = nil
	file_tensorflow_compiler_xla_service_hlo_profile_printer_data_proto_depIdxs = nil
}
