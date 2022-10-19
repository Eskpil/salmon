// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: m2m.proto

package definitions

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

// M2MClient is the client API for M2M service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type M2MClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	FinishTask(ctx context.Context, in *FinishTaskRequest, opts ...grpc.CallOption) (*FinishTaskResponse, error)
}

type m2MClient struct {
	cc grpc.ClientConnInterface
}

func NewM2MClient(cc grpc.ClientConnInterface) M2MClient {
	return &m2MClient{cc}
}

func (c *m2MClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/definitions.m2m/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *m2MClient) FinishTask(ctx context.Context, in *FinishTaskRequest, opts ...grpc.CallOption) (*FinishTaskResponse, error) {
	out := new(FinishTaskResponse)
	err := c.cc.Invoke(ctx, "/definitions.m2m/FinishTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// M2MServer is the server API for M2M service.
// All implementations must embed UnimplementedM2MServer
// for forward compatibility
type M2MServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	FinishTask(context.Context, *FinishTaskRequest) (*FinishTaskResponse, error)
	mustEmbedUnimplementedM2MServer()
}

// UnimplementedM2MServer must be embedded to have forward compatible implementations.
type UnimplementedM2MServer struct {
}

func (UnimplementedM2MServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedM2MServer) FinishTask(context.Context, *FinishTaskRequest) (*FinishTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTask not implemented")
}
func (UnimplementedM2MServer) mustEmbedUnimplementedM2MServer() {}

// UnsafeM2MServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to M2MServer will
// result in compilation errors.
type UnsafeM2MServer interface {
	mustEmbedUnimplementedM2MServer()
}

func RegisterM2MServer(s grpc.ServiceRegistrar, srv M2MServer) {
	s.RegisterService(&M2M_ServiceDesc, srv)
}

func _M2M_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M2MServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.m2m/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M2MServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _M2M_FinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M2MServer).FinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.m2m/FinishTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M2MServer).FinishTask(ctx, req.(*FinishTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// M2M_ServiceDesc is the grpc.ServiceDesc for M2M service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var M2M_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "definitions.m2m",
	HandlerType: (*M2MServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _M2M_Heartbeat_Handler,
		},
		{
			MethodName: "FinishTask",
			Handler:    _M2M_FinishTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "m2m.proto",
}