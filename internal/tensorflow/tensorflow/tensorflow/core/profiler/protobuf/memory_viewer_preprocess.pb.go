// Protobuf definitions for communicating the results of the memory
// visualization analysis subprocess (written in C++) to the outer script which
// renders HTML from Python.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/profiler/protobuf/memory_viewer_preprocess.proto

package protobuf

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

// Describes a heap object that is displayed in a plot in the memory
// visualization HTML.
type HeapObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Color:
	//
	//	*HeapObject_Numbered
	//	*HeapObject_Named
	Color                isHeapObject_Color `protobuf_oneof:"color"`
	Label                string             `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	LogicalBufferId      int32              `protobuf:"varint,4,opt,name=logical_buffer_id,json=logicalBufferId,proto3" json:"logical_buffer_id,omitempty"`
	LogicalBufferSizeMib float64            `protobuf:"fixed64,5,opt,name=logical_buffer_size_mib,json=logicalBufferSizeMib,proto3" json:"logical_buffer_size_mib,omitempty"`
	UnpaddedShapeMib     float64            `protobuf:"fixed64,6,opt,name=unpadded_shape_mib,json=unpaddedShapeMib,proto3" json:"unpadded_shape_mib,omitempty"`
	InstructionName      string             `protobuf:"bytes,7,opt,name=instruction_name,json=instructionName,proto3" json:"instruction_name,omitempty"`
	ShapeString          string             `protobuf:"bytes,8,opt,name=shape_string,json=shapeString,proto3" json:"shape_string,omitempty"`
	TfOpName             string             `protobuf:"bytes,9,opt,name=tf_op_name,json=tfOpName,proto3" json:"tf_op_name,omitempty"`
	GroupName            string             `protobuf:"bytes,10,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	OpCode               string             `protobuf:"bytes,11,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
}

func (x *HeapObject) Reset() {
	*x = HeapObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeapObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeapObject) ProtoMessage() {}

func (x *HeapObject) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeapObject.ProtoReflect.Descriptor instead.
func (*HeapObject) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP(), []int{0}
}

func (m *HeapObject) GetColor() isHeapObject_Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func (x *HeapObject) GetNumbered() int32 {
	if x, ok := x.GetColor().(*HeapObject_Numbered); ok {
		return x.Numbered
	}
	return 0
}

func (x *HeapObject) GetNamed() string {
	if x, ok := x.GetColor().(*HeapObject_Named); ok {
		return x.Named
	}
	return ""
}

func (x *HeapObject) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *HeapObject) GetLogicalBufferId() int32 {
	if x != nil {
		return x.LogicalBufferId
	}
	return 0
}

func (x *HeapObject) GetLogicalBufferSizeMib() float64 {
	if x != nil {
		return x.LogicalBufferSizeMib
	}
	return 0
}

func (x *HeapObject) GetUnpaddedShapeMib() float64 {
	if x != nil {
		return x.UnpaddedShapeMib
	}
	return 0
}

func (x *HeapObject) GetInstructionName() string {
	if x != nil {
		return x.InstructionName
	}
	return ""
}

func (x *HeapObject) GetShapeString() string {
	if x != nil {
		return x.ShapeString
	}
	return ""
}

func (x *HeapObject) GetTfOpName() string {
	if x != nil {
		return x.TfOpName
	}
	return ""
}

func (x *HeapObject) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *HeapObject) GetOpCode() string {
	if x != nil {
		return x.OpCode
	}
	return ""
}

type isHeapObject_Color interface {
	isHeapObject_Color()
}

type HeapObject_Numbered struct {
	Numbered int32 `protobuf:"varint,1,opt,name=numbered,proto3,oneof"`
}

type HeapObject_Named struct {
	Named string `protobuf:"bytes,2,opt,name=named,proto3,oneof"`
}

func (*HeapObject_Numbered) isHeapObject_Color() {}

func (*HeapObject_Named) isHeapObject_Color() {}

// Describes the start / exclusive limit HLO program points for a given buffer
// lifetime, used for rendering a box on the plot.
type BufferSpan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *BufferSpan) Reset() {
	*x = BufferSpan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufferSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufferSpan) ProtoMessage() {}

func (x *BufferSpan) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufferSpan.ProtoReflect.Descriptor instead.
func (*BufferSpan) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP(), []int{1}
}

func (x *BufferSpan) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *BufferSpan) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type LogicalBuffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Shape      string  `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	SizeMib    float64 `protobuf:"fixed64,3,opt,name=size_mib,json=sizeMib,proto3" json:"size_mib,omitempty"`
	HloName    string  `protobuf:"bytes,4,opt,name=hlo_name,json=hloName,proto3" json:"hlo_name,omitempty"`
	ShapeIndex []int64 `protobuf:"varint,5,rep,packed,name=shape_index,json=shapeIndex,proto3" json:"shape_index,omitempty"`
}

func (x *LogicalBuffer) Reset() {
	*x = LogicalBuffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicalBuffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicalBuffer) ProtoMessage() {}

func (x *LogicalBuffer) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicalBuffer.ProtoReflect.Descriptor instead.
func (*LogicalBuffer) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP(), []int{2}
}

func (x *LogicalBuffer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogicalBuffer) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

func (x *LogicalBuffer) GetSizeMib() float64 {
	if x != nil {
		return x.SizeMib
	}
	return 0
}

func (x *LogicalBuffer) GetHloName() string {
	if x != nil {
		return x.HloName
	}
	return ""
}

func (x *LogicalBuffer) GetShapeIndex() []int64 {
	if x != nil {
		return x.ShapeIndex
	}
	return nil
}

type BufferAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeMib        float64          `protobuf:"fixed64,2,opt,name=size_mib,json=sizeMib,proto3" json:"size_mib,omitempty"`
	Attributes     []string         `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	LogicalBuffers []*LogicalBuffer `protobuf:"bytes,4,rep,name=logical_buffers,json=logicalBuffers,proto3" json:"logical_buffers,omitempty"`
	CommonShape    string           `protobuf:"bytes,5,opt,name=common_shape,json=commonShape,proto3" json:"common_shape,omitempty"`
}

func (x *BufferAllocation) Reset() {
	*x = BufferAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufferAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufferAllocation) ProtoMessage() {}

func (x *BufferAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufferAllocation.ProtoReflect.Descriptor instead.
func (*BufferAllocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP(), []int{3}
}

func (x *BufferAllocation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BufferAllocation) GetSizeMib() float64 {
	if x != nil {
		return x.SizeMib
	}
	return 0
}

func (x *BufferAllocation) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *BufferAllocation) GetLogicalBuffers() []*LogicalBuffer {
	if x != nil {
		return x.LogicalBuffers
	}
	return nil
}

func (x *BufferAllocation) GetCommonShape() string {
	if x != nil {
		return x.CommonShape
	}
	return ""
}

// Groups together all results from the preprocessing C++ step.
type PreprocessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Heap sizes at each HLO program point (the HLO sequential order).
	HeapSizes []float64 `protobuf:"fixed64,1,rep,packed,name=heap_sizes,json=heapSizes,proto3" json:"heap_sizes,omitempty"`
	// Unpadded heap sizes (calculated as the minimal sizes based on the data type
	// and dimensionality) at each HLO program point (the HLO sequential order).
	UnpaddedHeapSizes []float64 `protobuf:"fixed64,2,rep,packed,name=unpadded_heap_sizes,json=unpaddedHeapSizes,proto3" json:"unpadded_heap_sizes,omitempty"`
	// Heap objects at the peak memory usage point ordered by HLO program "birth"
	// time.
	MaxHeap []*HeapObject `protobuf:"bytes,3,rep,name=max_heap,json=maxHeap,proto3" json:"max_heap,omitempty"`
	// Heap objects at the peak memory usage point ordered by size, descending.
	MaxHeapBySize []*HeapObject `protobuf:"bytes,4,rep,name=max_heap_by_size,json=maxHeapBySize,proto3" json:"max_heap_by_size,omitempty"`
	// Mapping from logical buffer ID to the HLO sequential order span in which it
	// is alive.
	LogicalBufferSpans map[int32]*BufferSpan `protobuf:"bytes,5,rep,name=logical_buffer_spans,json=logicalBufferSpans,proto3" json:"logical_buffer_spans,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Indexes to get back and forth from the by-size and by-program-order
	// sequences.
	MaxHeapToBySize      []int32 `protobuf:"varint,6,rep,packed,name=max_heap_to_by_size,json=maxHeapToBySize,proto3" json:"max_heap_to_by_size,omitempty"`
	BySizeToMaxHeap      []int32 `protobuf:"varint,7,rep,packed,name=by_size_to_max_heap,json=bySizeToMaxHeap,proto3" json:"by_size_to_max_heap,omitempty"`
	ModuleName           string  `protobuf:"bytes,8,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	EntryComputationName string  `protobuf:"bytes,9,opt,name=entry_computation_name,json=entryComputationName,proto3" json:"entry_computation_name,omitempty"`
	// Peak heap size for the HLO program.
	PeakHeapMib float64 `protobuf:"fixed64,10,opt,name=peak_heap_mib,json=peakHeapMib,proto3" json:"peak_heap_mib,omitempty"`
	// Peak unpadded heap size for the HLO program.
	PeakUnpaddedHeapMib float64 `protobuf:"fixed64,11,opt,name=peak_unpadded_heap_mib,json=peakUnpaddedHeapMib,proto3" json:"peak_unpadded_heap_mib,omitempty"`
	// HLO program point number at which the peak heap size occurs.
	PeakHeapSizePosition int32 `protobuf:"varint,12,opt,name=peak_heap_size_position,json=peakHeapSizePosition,proto3" json:"peak_heap_size_position,omitempty"`
	// Size of the entry computation parameters in MiB.
	//
	// This does not reflect whether those MiB are reusable during the computation
	// or not, it is simply a size value.
	EntryComputationParametersMib float64 `protobuf:"fixed64,13,opt,name=entry_computation_parameters_mib,json=entryComputationParametersMib,proto3" json:"entry_computation_parameters_mib,omitempty"`
	NonReusableMib                float64 `protobuf:"fixed64,14,opt,name=non_reusable_mib,json=nonReusableMib,proto3" json:"non_reusable_mib,omitempty"`
	MaybeLiveOutMib               float64 `protobuf:"fixed64,15,opt,name=maybe_live_out_mib,json=maybeLiveOutMib,proto3" json:"maybe_live_out_mib,omitempty"`
	// total size of indefinite/global and temporary buffer allocations.
	TotalBufferAllocationMib float64 `protobuf:"fixed64,18,opt,name=total_buffer_allocation_mib,json=totalBufferAllocationMib,proto3" json:"total_buffer_allocation_mib,omitempty"`
	// total size of indefinite/global buffer allocations.
	IndefiniteBufferAllocationMib float64             `protobuf:"fixed64,19,opt,name=indefinite_buffer_allocation_mib,json=indefiniteBufferAllocationMib,proto3" json:"indefinite_buffer_allocation_mib,omitempty"`
	IndefiniteLifetimes           []*BufferAllocation `protobuf:"bytes,16,rep,name=indefinite_lifetimes,json=indefiniteLifetimes,proto3" json:"indefinite_lifetimes,omitempty"`
	AllocationTimeline            string              `protobuf:"bytes,17,opt,name=allocation_timeline,json=allocationTimeline,proto3" json:"allocation_timeline,omitempty"`
}

func (x *PreprocessResult) Reset() {
	*x = PreprocessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreprocessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreprocessResult) ProtoMessage() {}

func (x *PreprocessResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreprocessResult.ProtoReflect.Descriptor instead.
func (*PreprocessResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP(), []int{4}
}

func (x *PreprocessResult) GetHeapSizes() []float64 {
	if x != nil {
		return x.HeapSizes
	}
	return nil
}

func (x *PreprocessResult) GetUnpaddedHeapSizes() []float64 {
	if x != nil {
		return x.UnpaddedHeapSizes
	}
	return nil
}

func (x *PreprocessResult) GetMaxHeap() []*HeapObject {
	if x != nil {
		return x.MaxHeap
	}
	return nil
}

func (x *PreprocessResult) GetMaxHeapBySize() []*HeapObject {
	if x != nil {
		return x.MaxHeapBySize
	}
	return nil
}

func (x *PreprocessResult) GetLogicalBufferSpans() map[int32]*BufferSpan {
	if x != nil {
		return x.LogicalBufferSpans
	}
	return nil
}

func (x *PreprocessResult) GetMaxHeapToBySize() []int32 {
	if x != nil {
		return x.MaxHeapToBySize
	}
	return nil
}

func (x *PreprocessResult) GetBySizeToMaxHeap() []int32 {
	if x != nil {
		return x.BySizeToMaxHeap
	}
	return nil
}

func (x *PreprocessResult) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *PreprocessResult) GetEntryComputationName() string {
	if x != nil {
		return x.EntryComputationName
	}
	return ""
}

func (x *PreprocessResult) GetPeakHeapMib() float64 {
	if x != nil {
		return x.PeakHeapMib
	}
	return 0
}

func (x *PreprocessResult) GetPeakUnpaddedHeapMib() float64 {
	if x != nil {
		return x.PeakUnpaddedHeapMib
	}
	return 0
}

func (x *PreprocessResult) GetPeakHeapSizePosition() int32 {
	if x != nil {
		return x.PeakHeapSizePosition
	}
	return 0
}

func (x *PreprocessResult) GetEntryComputationParametersMib() float64 {
	if x != nil {
		return x.EntryComputationParametersMib
	}
	return 0
}

func (x *PreprocessResult) GetNonReusableMib() float64 {
	if x != nil {
		return x.NonReusableMib
	}
	return 0
}

func (x *PreprocessResult) GetMaybeLiveOutMib() float64 {
	if x != nil {
		return x.MaybeLiveOutMib
	}
	return 0
}

func (x *PreprocessResult) GetTotalBufferAllocationMib() float64 {
	if x != nil {
		return x.TotalBufferAllocationMib
	}
	return 0
}

func (x *PreprocessResult) GetIndefiniteBufferAllocationMib() float64 {
	if x != nil {
		return x.IndefiniteBufferAllocationMib
	}
	return 0
}

func (x *PreprocessResult) GetIndefiniteLifetimes() []*BufferAllocation {
	if x != nil {
		return x.IndefiniteLifetimes
	}
	return nil
}

func (x *PreprocessResult) GetAllocationTimeline() string {
	if x != nil {
		return x.AllocationTimeline
	}
	return ""
}

var File_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDesc = []byte{
	0x0a, 0x40, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x96, 0x03, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x70,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x17, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x14, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69,
	0x7a, 0x65, 0x4d, 0x69, 0x62, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x6e, 0x70, 0x61, 0x64, 0x64, 0x65,
	0x64, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x10, 0x75, 0x6e, 0x70, 0x61, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65,
	0x4d, 0x69, 0x62, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x70, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x66, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x66, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x22, 0x38, 0x0a, 0x0a, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x69, 0x62, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x6c, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x6c, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x70,
	0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xcd, 0x01, 0x0a, 0x10, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x69, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x42,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x22, 0xb6, 0x09, 0x0a, 0x10, 0x50, 0x72,
	0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x68, 0x65, 0x61, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x09, 0x68, 0x65, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x13, 0x75, 0x6e, 0x70, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x11, 0x75, 0x6e, 0x70, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x48, 0x65, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x3a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70, 0x12, 0x48, 0x0a, 0x10, 0x6d, 0x61, 0x78,
	0x5f, 0x68, 0x65, 0x61, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x70, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70, 0x42, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x6f, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x70, 0x61, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x61, 0x70,
	0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70, 0x54, 0x6f, 0x42, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x13, 0x62, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0f, 0x62, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x4d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x65, 0x61, 0x6b, 0x5f,
	0x68, 0x65, 0x61, 0x70, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x70, 0x65, 0x61, 0x6b, 0x48, 0x65, 0x61, 0x70, 0x4d, 0x69, 0x62, 0x12, 0x33, 0x0a, 0x16, 0x70,
	0x65, 0x61, 0x6b, 0x5f, 0x75, 0x6e, 0x70, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61,
	0x70, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x70, 0x65, 0x61,
	0x6b, 0x55, 0x6e, 0x70, 0x61, 0x64, 0x64, 0x65, 0x64, 0x48, 0x65, 0x61, 0x70, 0x4d, 0x69, 0x62,
	0x12, 0x35, 0x0a, 0x17, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x14, 0x70, 0x65, 0x61, 0x6b, 0x48, 0x65, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x1d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x4d, 0x69, 0x62,
	0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6d, 0x69, 0x62, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6e, 0x6f, 0x6e, 0x52,
	0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x69, 0x62, 0x12, 0x2b, 0x0a, 0x12, 0x6d, 0x61,
	0x79, 0x62, 0x65, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x62,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x4c, 0x69, 0x76,
	0x65, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x62, 0x12, 0x3d, 0x0a, 0x1b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x62, 0x12, 0x47, 0x0a, 0x20, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x1d, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x62, 0x12,
	0x58, 0x0a, 0x14, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65,
	0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x66, 0x0a, 0x17, 0x4c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0xf1, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x1b,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x50, 0x72, 0x65, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x50,
	0x58, 0xaa, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xca, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xe2, 0x02, 0x1f,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescOnce sync.Once
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescData = file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDesc
)

func file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescGZIP() []byte {
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescData)
	})
	return file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDescData
}

var file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_goTypes = []interface{}{
	(*HeapObject)(nil),       // 0: tensorflow.profiler.HeapObject
	(*BufferSpan)(nil),       // 1: tensorflow.profiler.BufferSpan
	(*LogicalBuffer)(nil),    // 2: tensorflow.profiler.LogicalBuffer
	(*BufferAllocation)(nil), // 3: tensorflow.profiler.BufferAllocation
	(*PreprocessResult)(nil), // 4: tensorflow.profiler.PreprocessResult
	nil,                      // 5: tensorflow.profiler.PreprocessResult.LogicalBufferSpansEntry
}
var file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_depIdxs = []int32{
	2, // 0: tensorflow.profiler.BufferAllocation.logical_buffers:type_name -> tensorflow.profiler.LogicalBuffer
	0, // 1: tensorflow.profiler.PreprocessResult.max_heap:type_name -> tensorflow.profiler.HeapObject
	0, // 2: tensorflow.profiler.PreprocessResult.max_heap_by_size:type_name -> tensorflow.profiler.HeapObject
	5, // 3: tensorflow.profiler.PreprocessResult.logical_buffer_spans:type_name -> tensorflow.profiler.PreprocessResult.LogicalBufferSpansEntry
	3, // 4: tensorflow.profiler.PreprocessResult.indefinite_lifetimes:type_name -> tensorflow.profiler.BufferAllocation
	1, // 5: tensorflow.profiler.PreprocessResult.LogicalBufferSpansEntry.value:type_name -> tensorflow.profiler.BufferSpan
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_init() }
func file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_init() {
	if File_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeapObject); i {
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
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufferSpan); i {
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
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicalBuffer); i {
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
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufferAllocation); i {
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
		file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreprocessResult); i {
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
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HeapObject_Numbered)(nil),
		(*HeapObject_Named)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_msgTypes,
	}.Build()
	File_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto = out.File
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_memory_viewer_preprocess_proto_depIdxs = nil
}
