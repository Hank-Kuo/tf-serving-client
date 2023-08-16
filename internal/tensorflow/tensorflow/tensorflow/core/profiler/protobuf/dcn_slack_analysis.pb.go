// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/profiler/protobuf/dcn_slack_analysis.proto

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

type DcnSlack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rendezvous string `protobuf:"bytes,1,opt,name=rendezvous,proto3" json:"rendezvous,omitempty"`
	// Xprof observed send start time.
	SendStartTimeUs uint64 `protobuf:"varint,2,opt,name=send_start_time_us,json=sendStartTimeUs,proto3" json:"send_start_time_us,omitempty"`
	// Xprof observed recv_done end time.
	RecvDoneEndTimeUs uint64 `protobuf:"varint,3,opt,name=recv_done_end_time_us,json=recvDoneEndTimeUs,proto3" json:"recv_done_end_time_us,omitempty"`
	// Slack is defined as the time the collective has to send and recv data
	// without stalling the tpu. The effect of the network and other overlapping
	// collectives are removed from the collective of interest.
	//
	// HOST 1 :
	// |--------|SEND1|-------|SEND1.DONE|-------|RECV1|------|RECV1.DONE|-------
	// HOST 2:
	// |------|SEND2|-------|SEND2.DONE|-------|RECV2|------|RECV2.DONE    |-----
	//
	// Slack is computed as
	// RECV2.DONE.StartTime - SEND2.StartTime - (Overlapping Communication)
	// In this case, Overlapping communication is the duration of SEND2,
	// SEND2.DONE and RECV2. In cases where other collectives are interspaced
	// between this collective, Overlapping duration would include their durations
	// as well. Host 1 is ignored while computing the slack, as we assume that the
	// similar ops are executing each core. This also prevents clock drifts to
	// effect the analysis.
	SlackUs                     uint64 `protobuf:"varint,4,opt,name=slack_us,json=slackUs,proto3" json:"slack_us,omitempty"`
	BytesTransmittedOverNetwork uint64 `protobuf:"varint,5,opt,name=bytes_transmitted_over_network,json=bytesTransmittedOverNetwork,proto3" json:"bytes_transmitted_over_network,omitempty"`
	// Duration the collective stalled the TPU.
	StallDurationUs uint64 `protobuf:"varint,6,opt,name=stall_duration_us,json=stallDurationUs,proto3" json:"stall_duration_us,omitempty"`
	// Recv op name
	RecvOpName string `protobuf:"bytes,7,opt,name=recv_op_name,json=recvOpName,proto3" json:"recv_op_name,omitempty"`
	// Send op name
	SendOpName string `protobuf:"bytes,8,opt,name=send_op_name,json=sendOpName,proto3" json:"send_op_name,omitempty"`
}

func (x *DcnSlack) Reset() {
	*x = DcnSlack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcnSlack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcnSlack) ProtoMessage() {}

func (x *DcnSlack) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcnSlack.ProtoReflect.Descriptor instead.
func (*DcnSlack) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *DcnSlack) GetRendezvous() string {
	if x != nil {
		return x.Rendezvous
	}
	return ""
}

func (x *DcnSlack) GetSendStartTimeUs() uint64 {
	if x != nil {
		return x.SendStartTimeUs
	}
	return 0
}

func (x *DcnSlack) GetRecvDoneEndTimeUs() uint64 {
	if x != nil {
		return x.RecvDoneEndTimeUs
	}
	return 0
}

func (x *DcnSlack) GetSlackUs() uint64 {
	if x != nil {
		return x.SlackUs
	}
	return 0
}

func (x *DcnSlack) GetBytesTransmittedOverNetwork() uint64 {
	if x != nil {
		return x.BytesTransmittedOverNetwork
	}
	return 0
}

func (x *DcnSlack) GetStallDurationUs() uint64 {
	if x != nil {
		return x.StallDurationUs
	}
	return 0
}

func (x *DcnSlack) GetRecvOpName() string {
	if x != nil {
		return x.RecvOpName
	}
	return ""
}

func (x *DcnSlack) GetSendOpName() string {
	if x != nil {
		return x.SendOpName
	}
	return ""
}

type DcnSlackSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rendezvous name for the collective.
	Rendezvous string `protobuf:"bytes,1,opt,name=rendezvous,proto3" json:"rendezvous,omitempty"`
	// Slack Time in Microseconds,
	SlackUs uint64 `protobuf:"varint,2,opt,name=slack_us,json=slackUs,proto3" json:"slack_us,omitempty"`
	// Number of occurrences in the sampled duration.
	Occurrences uint64 `protobuf:"varint,3,opt,name=occurrences,proto3" json:"occurrences,omitempty"`
	// Bytes transmitted over the network.
	BytesTransmittedOverNetwork uint64 `protobuf:"varint,4,opt,name=bytes_transmitted_over_network,json=bytesTransmittedOverNetwork,proto3" json:"bytes_transmitted_over_network,omitempty"`
	// Duration the collective stalled the TPU.
	StallDurationUs uint64 `protobuf:"varint,5,opt,name=stall_duration_us,json=stallDurationUs,proto3" json:"stall_duration_us,omitempty"`
	// Observed duration.
	ObservedDurationUs uint64 `protobuf:"varint,6,opt,name=observed_duration_us,json=observedDurationUs,proto3" json:"observed_duration_us,omitempty"`
	// Recv op name.
	RecvOpName string `protobuf:"bytes,7,opt,name=recv_op_name,json=recvOpName,proto3" json:"recv_op_name,omitempty"`
	// Send op name.
	SendOpName string `protobuf:"bytes,8,opt,name=send_op_name,json=sendOpName,proto3" json:"send_op_name,omitempty"`
}

func (x *DcnSlackSummary) Reset() {
	*x = DcnSlackSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcnSlackSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcnSlackSummary) ProtoMessage() {}

func (x *DcnSlackSummary) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcnSlackSummary.ProtoReflect.Descriptor instead.
func (*DcnSlackSummary) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *DcnSlackSummary) GetRendezvous() string {
	if x != nil {
		return x.Rendezvous
	}
	return ""
}

func (x *DcnSlackSummary) GetSlackUs() uint64 {
	if x != nil {
		return x.SlackUs
	}
	return 0
}

func (x *DcnSlackSummary) GetOccurrences() uint64 {
	if x != nil {
		return x.Occurrences
	}
	return 0
}

func (x *DcnSlackSummary) GetBytesTransmittedOverNetwork() uint64 {
	if x != nil {
		return x.BytesTransmittedOverNetwork
	}
	return 0
}

func (x *DcnSlackSummary) GetStallDurationUs() uint64 {
	if x != nil {
		return x.StallDurationUs
	}
	return 0
}

func (x *DcnSlackSummary) GetObservedDurationUs() uint64 {
	if x != nil {
		return x.ObservedDurationUs
	}
	return 0
}

func (x *DcnSlackSummary) GetRecvOpName() string {
	if x != nil {
		return x.RecvOpName
	}
	return ""
}

func (x *DcnSlackSummary) GetSendOpName() string {
	if x != nil {
		return x.SendOpName
	}
	return ""
}

type DcnSlackAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DcnSlack        []*DcnSlack        `protobuf:"bytes,1,rep,name=dcn_slack,json=dcnSlack,proto3" json:"dcn_slack,omitempty"`
	DcnSlackSummary []*DcnSlackSummary `protobuf:"bytes,2,rep,name=dcn_slack_summary,json=dcnSlackSummary,proto3" json:"dcn_slack_summary,omitempty"`
}

func (x *DcnSlackAnalysis) Reset() {
	*x = DcnSlackAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcnSlackAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcnSlackAnalysis) ProtoMessage() {}

func (x *DcnSlackAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcnSlackAnalysis.ProtoReflect.Descriptor instead.
func (*DcnSlackAnalysis) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *DcnSlackAnalysis) GetDcnSlack() []*DcnSlack {
	if x != nil {
		return x.DcnSlack
	}
	return nil
}

func (x *DcnSlackAnalysis) GetDcnSlackSummary() []*DcnSlackSummary {
	if x != nil {
		return x.DcnSlackSummary
	}
	return nil
}

var File_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x63, 0x6e, 0x5f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x22, 0xd9, 0x02, 0x0a, 0x08, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x12, 0x2b,
	0x0a, 0x12, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x30, 0x0a, 0x15, 0x72,
	0x65, 0x63, 0x76, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x72, 0x65, 0x63, 0x76,
	0x44, 0x6f, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x1b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2a, 0x0a,
	0x11, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x76, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x76, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd5, 0x02,
	0x0a, 0x0f, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x43,
	0x0a, 0x1e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x4f, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x4f,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61,
	0x63, 0x6b, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x63,
	0x6e, 0x5f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x64, 0x63,
	0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x50, 0x0a, 0x11, 0x64, 0x63, 0x6e, 0x5f, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0f, 0x64, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0xeb, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x42, 0x15, 0x44, 0x63, 0x6e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67,
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
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescOnce sync.Once
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescData = file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDesc
)

func file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescGZIP() []byte {
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescData)
	})
	return file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDescData
}

var file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_goTypes = []interface{}{
	(*DcnSlack)(nil),         // 0: tensorflow.profiler.DcnSlack
	(*DcnSlackSummary)(nil),  // 1: tensorflow.profiler.DcnSlackSummary
	(*DcnSlackAnalysis)(nil), // 2: tensorflow.profiler.DcnSlackAnalysis
}
var file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_depIdxs = []int32{
	0, // 0: tensorflow.profiler.DcnSlackAnalysis.dcn_slack:type_name -> tensorflow.profiler.DcnSlack
	1, // 1: tensorflow.profiler.DcnSlackAnalysis.dcn_slack_summary:type_name -> tensorflow.profiler.DcnSlackSummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_init() }
func file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_init() {
	if File_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcnSlack); i {
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
		file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcnSlackSummary); i {
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
		file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcnSlackAnalysis); i {
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
			RawDescriptor: file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_msgTypes,
	}.Build()
	File_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto = out.File
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_dcn_slack_analysis_proto_depIdxs = nil
}