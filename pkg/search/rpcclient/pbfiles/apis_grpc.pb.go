// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.1
// source: apis.proto

package apis

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

// MataServiceClient is the client API for MataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MataServiceClient interface {
	CreateServerQueue(ctx context.Context, in *CreateServerQueueReq, opts ...grpc.CallOption) (*CreateServerQueueRes, error)
	DeleteServerQueue(ctx context.Context, in *DeleteServerQueueReq, opts ...grpc.CallOption) (*DeleteServerQueueRes, error)
	CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerRes, error)
	DeleteServer(ctx context.Context, in *DeleteServerReq, opts ...grpc.CallOption) (*DeleteServerRes, error)
}

type mataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMataServiceClient(cc grpc.ClientConnInterface) MataServiceClient {
	return &mataServiceClient{cc}
}

func (c *mataServiceClient) CreateServerQueue(ctx context.Context, in *CreateServerQueueReq, opts ...grpc.CallOption) (*CreateServerQueueRes, error) {
	out := new(CreateServerQueueRes)
	err := c.cc.Invoke(ctx, "/MataService/CreateServerQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mataServiceClient) DeleteServerQueue(ctx context.Context, in *DeleteServerQueueReq, opts ...grpc.CallOption) (*DeleteServerQueueRes, error) {
	out := new(DeleteServerQueueRes)
	err := c.cc.Invoke(ctx, "/MataService/DeleteServerQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mataServiceClient) CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerRes, error) {
	out := new(CreateServerRes)
	err := c.cc.Invoke(ctx, "/MataService/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mataServiceClient) DeleteServer(ctx context.Context, in *DeleteServerReq, opts ...grpc.CallOption) (*DeleteServerRes, error) {
	out := new(DeleteServerRes)
	err := c.cc.Invoke(ctx, "/MataService/DeleteServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MataServiceServer is the server API for MataService service.
// All implementations must embed UnimplementedMataServiceServer
// for forward compatibility
type MataServiceServer interface {
	CreateServerQueue(context.Context, *CreateServerQueueReq) (*CreateServerQueueRes, error)
	DeleteServerQueue(context.Context, *DeleteServerQueueReq) (*DeleteServerQueueRes, error)
	CreateServer(context.Context, *CreateServerReq) (*CreateServerRes, error)
	DeleteServer(context.Context, *DeleteServerReq) (*DeleteServerRes, error)
	mustEmbedUnimplementedMataServiceServer()
}

// UnimplementedMataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMataServiceServer struct {
}

func (UnimplementedMataServiceServer) CreateServerQueue(context.Context, *CreateServerQueueReq) (*CreateServerQueueRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServerQueue not implemented")
}
func (UnimplementedMataServiceServer) DeleteServerQueue(context.Context, *DeleteServerQueueReq) (*DeleteServerQueueRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerQueue not implemented")
}
func (UnimplementedMataServiceServer) CreateServer(context.Context, *CreateServerReq) (*CreateServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedMataServiceServer) DeleteServer(context.Context, *DeleteServerReq) (*DeleteServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedMataServiceServer) mustEmbedUnimplementedMataServiceServer() {}

// UnsafeMataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MataServiceServer will
// result in compilation errors.
type UnsafeMataServiceServer interface {
	mustEmbedUnimplementedMataServiceServer()
}

func RegisterMataServiceServer(s grpc.ServiceRegistrar, srv MataServiceServer) {
	s.RegisterService(&MataService_ServiceDesc, srv)
}

func _MataService_CreateServerQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerQueueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MataServiceServer).CreateServerQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MataService/CreateServerQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MataServiceServer).CreateServerQueue(ctx, req.(*CreateServerQueueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MataService_DeleteServerQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerQueueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MataServiceServer).DeleteServerQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MataService/DeleteServerQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MataServiceServer).DeleteServerQueue(ctx, req.(*DeleteServerQueueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MataService_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MataServiceServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MataService/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MataServiceServer).CreateServer(ctx, req.(*CreateServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MataService_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MataServiceServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MataService/DeleteServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MataServiceServer).DeleteServer(ctx, req.(*DeleteServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MataService_ServiceDesc is the grpc.ServiceDesc for MataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MataService",
	HandlerType: (*MataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServerQueue",
			Handler:    _MataService_CreateServerQueue_Handler,
		},
		{
			MethodName: "DeleteServerQueue",
			Handler:    _MataService_DeleteServerQueue_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _MataService_CreateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _MataService_DeleteServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis.proto",
}