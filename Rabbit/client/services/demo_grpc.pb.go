// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

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

// JuegosClient is the client API for Juegos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JuegosClient interface {
	Jugar(ctx context.Context, in *JuegosRequest, opts ...grpc.CallOption) (*JuegosReply, error)
}

type juegosClient struct {
	cc grpc.ClientConnInterface
}

func NewJuegosClient(cc grpc.ClientConnInterface) JuegosClient {
	return &juegosClient{cc}
}

func (c *juegosClient) Jugar(ctx context.Context, in *JuegosRequest, opts ...grpc.CallOption) (*JuegosReply, error) {
	out := new(JuegosReply)
	err := c.cc.Invoke(ctx, "/proto.Juegos/Jugar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JuegosServer is the server API for Juegos service.
// All implementations must embed UnimplementedJuegosServer
// for forward compatibility
type JuegosServer interface {
	Jugar(context.Context, *JuegosRequest) (*JuegosReply, error)
	mustEmbedUnimplementedJuegosServer()
}

// UnimplementedJuegosServer must be embedded to have forward compatible implementations.
type UnimplementedJuegosServer struct {
}

func (UnimplementedJuegosServer) Jugar(context.Context, *JuegosRequest) (*JuegosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jugar not implemented")
}
func (UnimplementedJuegosServer) mustEmbedUnimplementedJuegosServer() {}

// UnsafeJuegosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JuegosServer will
// result in compilation errors.
type UnsafeJuegosServer interface {
	mustEmbedUnimplementedJuegosServer()
}

func RegisterJuegosServer(s grpc.ServiceRegistrar, srv JuegosServer) {
	s.RegisterService(&Juegos_ServiceDesc, srv)
}

func _Juegos_Jugar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JuegosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JuegosServer).Jugar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juegos/Jugar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JuegosServer).Jugar(ctx, req.(*JuegosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Juegos_ServiceDesc is the grpc.ServiceDesc for Juegos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Juegos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Juegos",
	HandlerType: (*JuegosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Jugar",
			Handler:    _Juegos_Jugar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/demo.proto",
}
