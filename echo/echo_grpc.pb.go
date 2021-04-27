// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.8
// source: echo/echo.proto

package echo

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

// EchoesClient is the client API for Echoes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoesClient interface {
	// Sends a greeting
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type echoesClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoesClient(cc grpc.ClientConnInterface) EchoesClient {
	return &echoesClient{cc}
}

func (c *echoesClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/echo.Echoes/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoesServer is the server API for Echoes service.
// All implementations must embed UnimplementedEchoesServer
// for forward compatibility
type EchoesServer interface {
	// Sends a greeting
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
	mustEmbedUnimplementedEchoesServer()
}

// UnimplementedEchoesServer must be embedded to have forward compatible implementations.
type UnimplementedEchoesServer struct {
}

func (UnimplementedEchoesServer) Echo(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoesServer) mustEmbedUnimplementedEchoesServer() {}

// UnsafeEchoesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoesServer will
// result in compilation errors.
type UnsafeEchoesServer interface {
	mustEmbedUnimplementedEchoesServer()
}

func RegisterEchoesServer(s grpc.ServiceRegistrar, srv EchoesServer) {
	s.RegisterService(&Echoes_ServiceDesc, srv)
}

func _Echoes_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoesServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echoes/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoesServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Echoes_ServiceDesc is the grpc.ServiceDesc for Echoes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Echoes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echoes",
	HandlerType: (*EchoesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echoes_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo/echo.proto",
}