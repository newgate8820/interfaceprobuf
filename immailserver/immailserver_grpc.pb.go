// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: immailserver.proto

package immailserver

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

// ImMailServerClient is the client API for ImMailServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImMailServerClient interface {
	SendMail(ctx context.Context, in *SendMailReq, opts ...grpc.CallOption) (*SendMailResult, error)
	// 发送邮件
	SendMailSSL(ctx context.Context, in *SendMailSSLReq, opts ...grpc.CallOption) (*SendMailSSLResult, error)
}

type imMailServerClient struct {
	cc grpc.ClientConnInterface
}

func NewImMailServerClient(cc grpc.ClientConnInterface) ImMailServerClient {
	return &imMailServerClient{cc}
}

func (c *imMailServerClient) SendMail(ctx context.Context, in *SendMailReq, opts ...grpc.CallOption) (*SendMailResult, error) {
	out := new(SendMailResult)
	err := c.cc.Invoke(ctx, "/immailserver.ImMailServer/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imMailServerClient) SendMailSSL(ctx context.Context, in *SendMailSSLReq, opts ...grpc.CallOption) (*SendMailSSLResult, error) {
	out := new(SendMailSSLResult)
	err := c.cc.Invoke(ctx, "/immailserver.ImMailServer/SendMailSSL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImMailServerServer is the server API for ImMailServer service.
// All implementations must embed UnimplementedImMailServerServer
// for forward compatibility
type ImMailServerServer interface {
	SendMail(context.Context, *SendMailReq) (*SendMailResult, error)
	// 发送邮件
	SendMailSSL(context.Context, *SendMailSSLReq) (*SendMailSSLResult, error)
	mustEmbedUnimplementedImMailServerServer()
}

// UnimplementedImMailServerServer must be embedded to have forward compatible implementations.
type UnimplementedImMailServerServer struct {
}

func (UnimplementedImMailServerServer) SendMail(context.Context, *SendMailReq) (*SendMailResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedImMailServerServer) SendMailSSL(context.Context, *SendMailSSLReq) (*SendMailSSLResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMailSSL not implemented")
}
func (UnimplementedImMailServerServer) mustEmbedUnimplementedImMailServerServer() {}

// UnsafeImMailServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImMailServerServer will
// result in compilation errors.
type UnsafeImMailServerServer interface {
	mustEmbedUnimplementedImMailServerServer()
}

func RegisterImMailServerServer(s grpc.ServiceRegistrar, srv ImMailServerServer) {
	s.RegisterService(&ImMailServer_ServiceDesc, srv)
}

func _ImMailServer_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImMailServerServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immailserver.ImMailServer/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImMailServerServer).SendMail(ctx, req.(*SendMailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImMailServer_SendMailSSL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailSSLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImMailServerServer).SendMailSSL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immailserver.ImMailServer/SendMailSSL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImMailServerServer).SendMailSSL(ctx, req.(*SendMailSSLReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ImMailServer_ServiceDesc is the grpc.ServiceDesc for ImMailServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImMailServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immailserver.ImMailServer",
	HandlerType: (*ImMailServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _ImMailServer_SendMail_Handler,
		},
		{
			MethodName: "SendMailSSL",
			Handler:    _ImMailServer_SendMailSSL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "immailserver.proto",
}
