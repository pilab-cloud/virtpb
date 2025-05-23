// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pilab/cloud/frontend/v1/tasks_service.proto

package frontendv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TasksService_ListTasks_FullMethodName = "/pilab.cloud.frontend.v1.TasksService/ListTasks"
)

// TasksServiceClient is the client API for TasksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service Definition
type TasksServiceClient interface {
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
}

type tasksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksServiceClient(cc grpc.ClientConnInterface) TasksServiceClient {
	return &tasksServiceClient{cc}
}

func (c *tasksServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, TasksService_ListTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServiceServer is the server API for TasksService service.
// All implementations must embed UnimplementedTasksServiceServer
// for forward compatibility.
//
// Service Definition
type TasksServiceServer interface {
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	mustEmbedUnimplementedTasksServiceServer()
}

// UnimplementedTasksServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTasksServiceServer struct{}

func (UnimplementedTasksServiceServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTasksServiceServer) mustEmbedUnimplementedTasksServiceServer() {}
func (UnimplementedTasksServiceServer) testEmbeddedByValue()                      {}

// UnsafeTasksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServiceServer will
// result in compilation errors.
type UnsafeTasksServiceServer interface {
	mustEmbedUnimplementedTasksServiceServer()
}

func RegisterTasksServiceServer(s grpc.ServiceRegistrar, srv TasksServiceServer) {
	// If the following call pancis, it indicates UnimplementedTasksServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TasksService_ServiceDesc, srv)
}

func _TasksService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_ListTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TasksService_ServiceDesc is the grpc.ServiceDesc for TasksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TasksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pilab.cloud.frontend.v1.TasksService",
	HandlerType: (*TasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTasks",
			Handler:    _TasksService_ListTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pilab/cloud/frontend/v1/tasks_service.proto",
}
