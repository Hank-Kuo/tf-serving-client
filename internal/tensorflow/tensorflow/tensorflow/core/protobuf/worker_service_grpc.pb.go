// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow/core/protobuf/worker_service.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	// See worker.proto for details.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// See worker.proto for details.
	CreateWorkerSession(ctx context.Context, in *CreateWorkerSessionRequest, opts ...grpc.CallOption) (*CreateWorkerSessionResponse, error)
	// See worker.proto for details.
	DeleteWorkerSession(ctx context.Context, in *DeleteWorkerSessionRequest, opts ...grpc.CallOption) (*DeleteWorkerSessionResponse, error)
	// See worker.proto for details.
	RegisterGraph(ctx context.Context, in *RegisterGraphRequest, opts ...grpc.CallOption) (*RegisterGraphResponse, error)
	// See worker.proto for details.
	DeregisterGraph(ctx context.Context, in *DeregisterGraphRequest, opts ...grpc.CallOption) (*DeregisterGraphResponse, error)
	// See worker.proto for details.
	RunGraph(ctx context.Context, in *RunGraphRequest, opts ...grpc.CallOption) (*RunGraphResponse, error)
	// See worker.proto for details.
	CleanupGraph(ctx context.Context, in *CleanupGraphRequest, opts ...grpc.CallOption) (*CleanupGraphResponse, error)
	// See worker.proto for details.
	CleanupAll(ctx context.Context, in *CleanupAllRequest, opts ...grpc.CallOption) (*CleanupAllResponse, error)
	// See worker.proto for details.
	RecvTensor(ctx context.Context, in *RecvTensorRequest, opts ...grpc.CallOption) (*RecvTensorResponse, error)
	// See worker.proto for details.
	MarkRecvFinished(ctx context.Context, in *MarkRecvFinishedRequest, opts ...grpc.CallOption) (*MarkRecvFinishedResponse, error)
	// See worker.proto for details.
	Logging(ctx context.Context, in *LoggingRequest, opts ...grpc.CallOption) (*LoggingResponse, error)
	// See worker.proto for details.
	Tracing(ctx context.Context, in *TracingRequest, opts ...grpc.CallOption) (*TracingResponse, error)
	// See worker.proto for details.
	RecvBuf(ctx context.Context, in *RecvBufRequest, opts ...grpc.CallOption) (*RecvBufResponse, error)
	// See worker.proto for details.
	GetStepSequence(ctx context.Context, in *GetStepSequenceRequest, opts ...grpc.CallOption) (*GetStepSequenceResponse, error)
	// See worker.proto for details.
	CompleteGroup(ctx context.Context, in *CompleteGroupRequest, opts ...grpc.CallOption) (*CompleteGroupResponse, error)
	// See worker.proto for details.
	CompleteInstance(ctx context.Context, in *CompleteInstanceRequest, opts ...grpc.CallOption) (*CompleteInstanceResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CreateWorkerSession(ctx context.Context, in *CreateWorkerSessionRequest, opts ...grpc.CallOption) (*CreateWorkerSessionResponse, error) {
	out := new(CreateWorkerSessionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/CreateWorkerSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) DeleteWorkerSession(ctx context.Context, in *DeleteWorkerSessionRequest, opts ...grpc.CallOption) (*DeleteWorkerSessionResponse, error) {
	out := new(DeleteWorkerSessionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/DeleteWorkerSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) RegisterGraph(ctx context.Context, in *RegisterGraphRequest, opts ...grpc.CallOption) (*RegisterGraphResponse, error) {
	out := new(RegisterGraphResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/RegisterGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) DeregisterGraph(ctx context.Context, in *DeregisterGraphRequest, opts ...grpc.CallOption) (*DeregisterGraphResponse, error) {
	out := new(DeregisterGraphResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/DeregisterGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) RunGraph(ctx context.Context, in *RunGraphRequest, opts ...grpc.CallOption) (*RunGraphResponse, error) {
	out := new(RunGraphResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/RunGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CleanupGraph(ctx context.Context, in *CleanupGraphRequest, opts ...grpc.CallOption) (*CleanupGraphResponse, error) {
	out := new(CleanupGraphResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/CleanupGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CleanupAll(ctx context.Context, in *CleanupAllRequest, opts ...grpc.CallOption) (*CleanupAllResponse, error) {
	out := new(CleanupAllResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/CleanupAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) RecvTensor(ctx context.Context, in *RecvTensorRequest, opts ...grpc.CallOption) (*RecvTensorResponse, error) {
	out := new(RecvTensorResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/RecvTensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) MarkRecvFinished(ctx context.Context, in *MarkRecvFinishedRequest, opts ...grpc.CallOption) (*MarkRecvFinishedResponse, error) {
	out := new(MarkRecvFinishedResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/MarkRecvFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Logging(ctx context.Context, in *LoggingRequest, opts ...grpc.CallOption) (*LoggingResponse, error) {
	out := new(LoggingResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/Logging", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Tracing(ctx context.Context, in *TracingRequest, opts ...grpc.CallOption) (*TracingResponse, error) {
	out := new(TracingResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/Tracing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) RecvBuf(ctx context.Context, in *RecvBufRequest, opts ...grpc.CallOption) (*RecvBufResponse, error) {
	out := new(RecvBufResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/RecvBuf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetStepSequence(ctx context.Context, in *GetStepSequenceRequest, opts ...grpc.CallOption) (*GetStepSequenceResponse, error) {
	out := new(GetStepSequenceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/GetStepSequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CompleteGroup(ctx context.Context, in *CompleteGroupRequest, opts ...grpc.CallOption) (*CompleteGroupResponse, error) {
	out := new(CompleteGroupResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/CompleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CompleteInstance(ctx context.Context, in *CompleteInstanceRequest, opts ...grpc.CallOption) (*CompleteInstanceResponse, error) {
	out := new(CompleteInstanceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.grpc.WorkerService/CompleteInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations should embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	// See worker.proto for details.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// See worker.proto for details.
	CreateWorkerSession(context.Context, *CreateWorkerSessionRequest) (*CreateWorkerSessionResponse, error)
	// See worker.proto for details.
	DeleteWorkerSession(context.Context, *DeleteWorkerSessionRequest) (*DeleteWorkerSessionResponse, error)
	// See worker.proto for details.
	RegisterGraph(context.Context, *RegisterGraphRequest) (*RegisterGraphResponse, error)
	// See worker.proto for details.
	DeregisterGraph(context.Context, *DeregisterGraphRequest) (*DeregisterGraphResponse, error)
	// See worker.proto for details.
	RunGraph(context.Context, *RunGraphRequest) (*RunGraphResponse, error)
	// See worker.proto for details.
	CleanupGraph(context.Context, *CleanupGraphRequest) (*CleanupGraphResponse, error)
	// See worker.proto for details.
	CleanupAll(context.Context, *CleanupAllRequest) (*CleanupAllResponse, error)
	// See worker.proto for details.
	RecvTensor(context.Context, *RecvTensorRequest) (*RecvTensorResponse, error)
	// See worker.proto for details.
	MarkRecvFinished(context.Context, *MarkRecvFinishedRequest) (*MarkRecvFinishedResponse, error)
	// See worker.proto for details.
	Logging(context.Context, *LoggingRequest) (*LoggingResponse, error)
	// See worker.proto for details.
	Tracing(context.Context, *TracingRequest) (*TracingResponse, error)
	// See worker.proto for details.
	RecvBuf(context.Context, *RecvBufRequest) (*RecvBufResponse, error)
	// See worker.proto for details.
	GetStepSequence(context.Context, *GetStepSequenceRequest) (*GetStepSequenceResponse, error)
	// See worker.proto for details.
	CompleteGroup(context.Context, *CompleteGroupRequest) (*CompleteGroupResponse, error)
	// See worker.proto for details.
	CompleteInstance(context.Context, *CompleteInstanceRequest) (*CompleteInstanceResponse, error)
}

// UnimplementedWorkerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedWorkerServiceServer) CreateWorkerSession(context.Context, *CreateWorkerSessionRequest) (*CreateWorkerSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkerSession not implemented")
}
func (UnimplementedWorkerServiceServer) DeleteWorkerSession(context.Context, *DeleteWorkerSessionRequest) (*DeleteWorkerSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkerSession not implemented")
}
func (UnimplementedWorkerServiceServer) RegisterGraph(context.Context, *RegisterGraphRequest) (*RegisterGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterGraph not implemented")
}
func (UnimplementedWorkerServiceServer) DeregisterGraph(context.Context, *DeregisterGraphRequest) (*DeregisterGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterGraph not implemented")
}
func (UnimplementedWorkerServiceServer) RunGraph(context.Context, *RunGraphRequest) (*RunGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunGraph not implemented")
}
func (UnimplementedWorkerServiceServer) CleanupGraph(context.Context, *CleanupGraphRequest) (*CleanupGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanupGraph not implemented")
}
func (UnimplementedWorkerServiceServer) CleanupAll(context.Context, *CleanupAllRequest) (*CleanupAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanupAll not implemented")
}
func (UnimplementedWorkerServiceServer) RecvTensor(context.Context, *RecvTensorRequest) (*RecvTensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecvTensor not implemented")
}
func (UnimplementedWorkerServiceServer) MarkRecvFinished(context.Context, *MarkRecvFinishedRequest) (*MarkRecvFinishedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRecvFinished not implemented")
}
func (UnimplementedWorkerServiceServer) Logging(context.Context, *LoggingRequest) (*LoggingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logging not implemented")
}
func (UnimplementedWorkerServiceServer) Tracing(context.Context, *TracingRequest) (*TracingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tracing not implemented")
}
func (UnimplementedWorkerServiceServer) RecvBuf(context.Context, *RecvBufRequest) (*RecvBufResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecvBuf not implemented")
}
func (UnimplementedWorkerServiceServer) GetStepSequence(context.Context, *GetStepSequenceRequest) (*GetStepSequenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStepSequence not implemented")
}
func (UnimplementedWorkerServiceServer) CompleteGroup(context.Context, *CompleteGroupRequest) (*CompleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteGroup not implemented")
}
func (UnimplementedWorkerServiceServer) CompleteInstance(context.Context, *CompleteInstanceRequest) (*CompleteInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteInstance not implemented")
}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CreateWorkerSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkerSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CreateWorkerSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/CreateWorkerSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CreateWorkerSession(ctx, req.(*CreateWorkerSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_DeleteWorkerSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkerSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).DeleteWorkerSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/DeleteWorkerSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).DeleteWorkerSession(ctx, req.(*DeleteWorkerSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_RegisterGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RegisterGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/RegisterGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RegisterGraph(ctx, req.(*RegisterGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_DeregisterGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).DeregisterGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/DeregisterGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).DeregisterGraph(ctx, req.(*DeregisterGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_RunGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RunGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/RunGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RunGraph(ctx, req.(*RunGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CleanupGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CleanupGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/CleanupGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CleanupGraph(ctx, req.(*CleanupGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CleanupAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CleanupAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/CleanupAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CleanupAll(ctx, req.(*CleanupAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_RecvTensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecvTensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RecvTensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/RecvTensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RecvTensor(ctx, req.(*RecvTensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_MarkRecvFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkRecvFinishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).MarkRecvFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/MarkRecvFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).MarkRecvFinished(ctx, req.(*MarkRecvFinishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Logging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoggingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Logging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/Logging",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Logging(ctx, req.(*LoggingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Tracing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Tracing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/Tracing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Tracing(ctx, req.(*TracingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_RecvBuf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecvBufRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RecvBuf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/RecvBuf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RecvBuf(ctx, req.(*RecvBufRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetStepSequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStepSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetStepSequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/GetStepSequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetStepSequence(ctx, req.(*GetStepSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CompleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CompleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/CompleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CompleteGroup(ctx, req.(*CompleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CompleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CompleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.grpc.WorkerService/CompleteInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CompleteInstance(ctx, req.(*CompleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.grpc.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _WorkerService_GetStatus_Handler,
		},
		{
			MethodName: "CreateWorkerSession",
			Handler:    _WorkerService_CreateWorkerSession_Handler,
		},
		{
			MethodName: "DeleteWorkerSession",
			Handler:    _WorkerService_DeleteWorkerSession_Handler,
		},
		{
			MethodName: "RegisterGraph",
			Handler:    _WorkerService_RegisterGraph_Handler,
		},
		{
			MethodName: "DeregisterGraph",
			Handler:    _WorkerService_DeregisterGraph_Handler,
		},
		{
			MethodName: "RunGraph",
			Handler:    _WorkerService_RunGraph_Handler,
		},
		{
			MethodName: "CleanupGraph",
			Handler:    _WorkerService_CleanupGraph_Handler,
		},
		{
			MethodName: "CleanupAll",
			Handler:    _WorkerService_CleanupAll_Handler,
		},
		{
			MethodName: "RecvTensor",
			Handler:    _WorkerService_RecvTensor_Handler,
		},
		{
			MethodName: "MarkRecvFinished",
			Handler:    _WorkerService_MarkRecvFinished_Handler,
		},
		{
			MethodName: "Logging",
			Handler:    _WorkerService_Logging_Handler,
		},
		{
			MethodName: "Tracing",
			Handler:    _WorkerService_Tracing_Handler,
		},
		{
			MethodName: "RecvBuf",
			Handler:    _WorkerService_RecvBuf_Handler,
		},
		{
			MethodName: "GetStepSequence",
			Handler:    _WorkerService_GetStepSequence_Handler,
		},
		{
			MethodName: "CompleteGroup",
			Handler:    _WorkerService_CompleteGroup_Handler,
		},
		{
			MethodName: "CompleteInstance",
			Handler:    _WorkerService_CompleteInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/core/protobuf/worker_service.proto",
}
