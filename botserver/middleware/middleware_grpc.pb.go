// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: middleware.proto

package middleware

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	// RouteGetBotCallbackAnswer 路由获取回调应答
	RouteGetBotCallbackAnswer(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	// RouteSetBotCallbackAnswer 路由设置回调应答
	RouteSetBotCallbackAnswer(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) RouteGetBotCallbackAnswer(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/middleware.Middleware/RouteGetBotCallbackAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) RouteSetBotCallbackAnswer(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/middleware.Middleware/RouteSetBotCallbackAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	// RouteGetBotCallbackAnswer 路由获取回调应答
	RouteGetBotCallbackAnswer(context.Context, *RouteRequest) (*RouteResponse, error)
	// RouteSetBotCallbackAnswer 路由设置回调应答
	RouteSetBotCallbackAnswer(context.Context, *RouteRequest) (*RouteResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) RouteGetBotCallbackAnswer(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteGetBotCallbackAnswer not implemented")
}
func (UnimplementedMiddlewareServer) RouteSetBotCallbackAnswer(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteSetBotCallbackAnswer not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_RouteGetBotCallbackAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).RouteGetBotCallbackAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/middleware.Middleware/RouteGetBotCallbackAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).RouteGetBotCallbackAnswer(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_RouteSetBotCallbackAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).RouteSetBotCallbackAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/middleware.Middleware/RouteSetBotCallbackAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).RouteSetBotCallbackAnswer(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "middleware.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RouteGetBotCallbackAnswer",
			Handler:    _Middleware_RouteGetBotCallbackAnswer_Handler,
		},
		{
			MethodName: "RouteSetBotCallbackAnswer",
			Handler:    _Middleware_RouteSetBotCallbackAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middleware.proto",
}