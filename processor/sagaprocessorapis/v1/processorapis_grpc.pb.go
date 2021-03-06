// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: sagaprocessorapis/v1/processorapis.proto

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

// SagaProcessorServiceClient is the client API for SagaProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SagaProcessorServiceClient interface {
	ExecuteWorkflow(ctx context.Context, in *ExecuteWorkflowRequest, opts ...grpc.CallOption) (*ExecuteWorkflowResponse, error)
}

type sagaProcessorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSagaProcessorServiceClient(cc grpc.ClientConnInterface) SagaProcessorServiceClient {
	return &sagaProcessorServiceClient{cc}
}

func (c *sagaProcessorServiceClient) ExecuteWorkflow(ctx context.Context, in *ExecuteWorkflowRequest, opts ...grpc.CallOption) (*ExecuteWorkflowResponse, error) {
	out := new(ExecuteWorkflowResponse)
	err := c.cc.Invoke(ctx, "/sagaprocessorapis.v1.SagaProcessorService/ExecuteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SagaProcessorServiceServer is the server API for SagaProcessorService service.
// All implementations must embed UnimplementedSagaProcessorServiceServer
// for forward compatibility
type SagaProcessorServiceServer interface {
	ExecuteWorkflow(context.Context, *ExecuteWorkflowRequest) (*ExecuteWorkflowResponse, error)
	mustEmbedUnimplementedSagaProcessorServiceServer()
}

// UnimplementedSagaProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSagaProcessorServiceServer struct {
}

func (UnimplementedSagaProcessorServiceServer) ExecuteWorkflow(context.Context, *ExecuteWorkflowRequest) (*ExecuteWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteWorkflow not implemented")
}
func (UnimplementedSagaProcessorServiceServer) mustEmbedUnimplementedSagaProcessorServiceServer() {}

// UnsafeSagaProcessorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SagaProcessorServiceServer will
// result in compilation errors.
type UnsafeSagaProcessorServiceServer interface {
	mustEmbedUnimplementedSagaProcessorServiceServer()
}

func RegisterSagaProcessorServiceServer(s grpc.ServiceRegistrar, srv SagaProcessorServiceServer) {
	s.RegisterService(&SagaProcessorService_ServiceDesc, srv)
}

func _SagaProcessorService_ExecuteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SagaProcessorServiceServer).ExecuteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sagaprocessorapis.v1.SagaProcessorService/ExecuteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SagaProcessorServiceServer).ExecuteWorkflow(ctx, req.(*ExecuteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SagaProcessorService_ServiceDesc is the grpc.ServiceDesc for SagaProcessorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SagaProcessorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sagaprocessorapis.v1.SagaProcessorService",
	HandlerType: (*SagaProcessorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteWorkflow",
			Handler:    _SagaProcessorService_ExecuteWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sagaprocessorapis/v1/processorapis.proto",
}
