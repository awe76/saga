// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: sagatracerapis/v1/tracerapis.proto

package v1

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

// SagaTracerServiceClient is the client API for SagaTracerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SagaTracerServiceClient interface {
	TraceWorkflow(ctx context.Context, in *TraceWorkflowRequest, opts ...grpc.CallOption) (SagaTracerService_TraceWorkflowClient, error)
}

type sagaTracerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSagaTracerServiceClient(cc grpc.ClientConnInterface) SagaTracerServiceClient {
	return &sagaTracerServiceClient{cc}
}

func (c *sagaTracerServiceClient) TraceWorkflow(ctx context.Context, in *TraceWorkflowRequest, opts ...grpc.CallOption) (SagaTracerService_TraceWorkflowClient, error) {
	stream, err := c.cc.NewStream(ctx, &SagaTracerService_ServiceDesc.Streams[0], "/sagatracerapis.v1.SagaTracerService/TraceWorkflow", opts...)
	if err != nil {
		return nil, err
	}
	x := &sagaTracerServiceTraceWorkflowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SagaTracerService_TraceWorkflowClient interface {
	Recv() (*TraceWorkflowResponse, error)
	grpc.ClientStream
}

type sagaTracerServiceTraceWorkflowClient struct {
	grpc.ClientStream
}

func (x *sagaTracerServiceTraceWorkflowClient) Recv() (*TraceWorkflowResponse, error) {
	m := new(TraceWorkflowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SagaTracerServiceServer is the server API for SagaTracerService service.
// All implementations must embed UnimplementedSagaTracerServiceServer
// for forward compatibility
type SagaTracerServiceServer interface {
	TraceWorkflow(*TraceWorkflowRequest, SagaTracerService_TraceWorkflowServer) error
	mustEmbedUnimplementedSagaTracerServiceServer()
}

// UnimplementedSagaTracerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSagaTracerServiceServer struct {
}

func (UnimplementedSagaTracerServiceServer) TraceWorkflow(*TraceWorkflowRequest, SagaTracerService_TraceWorkflowServer) error {
	return status.Errorf(codes.Unimplemented, "method TraceWorkflow not implemented")
}
func (UnimplementedSagaTracerServiceServer) mustEmbedUnimplementedSagaTracerServiceServer() {}

// UnsafeSagaTracerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SagaTracerServiceServer will
// result in compilation errors.
type UnsafeSagaTracerServiceServer interface {
	mustEmbedUnimplementedSagaTracerServiceServer()
}

func RegisterSagaTracerServiceServer(s grpc.ServiceRegistrar, srv SagaTracerServiceServer) {
	s.RegisterService(&SagaTracerService_ServiceDesc, srv)
}

func _SagaTracerService_TraceWorkflow_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceWorkflowRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SagaTracerServiceServer).TraceWorkflow(m, &sagaTracerServiceTraceWorkflowServer{stream})
}

type SagaTracerService_TraceWorkflowServer interface {
	Send(*TraceWorkflowResponse) error
	grpc.ServerStream
}

type sagaTracerServiceTraceWorkflowServer struct {
	grpc.ServerStream
}

func (x *sagaTracerServiceTraceWorkflowServer) Send(m *TraceWorkflowResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SagaTracerService_ServiceDesc is the grpc.ServiceDesc for SagaTracerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SagaTracerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sagatracerapis.v1.SagaTracerService",
	HandlerType: (*SagaTracerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TraceWorkflow",
			Handler:       _SagaTracerService_TraceWorkflow_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sagatracerapis/v1/tracerapis.proto",
}
