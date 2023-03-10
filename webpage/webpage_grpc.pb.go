// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: webpage.proto

package webpage

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

// WebPageClient is the client API for WebPage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebPageClient interface {
	// 解析url
	ParseUrl(ctx context.Context, in *ParseUrlRequest, opts ...grpc.CallOption) (*ParseUrlReply, error)
	// 解析url
	ParseUrlForEncoded(ctx context.Context, in *ParseUrlForEncodedRequest, opts ...grpc.CallOption) (*ParseUrlForEncodedReply, error)
	ParseUrlAndStore(ctx context.Context, in *ParseUrlAndStoreRequest, opts ...grpc.CallOption) (*ParseUrlAndStoreReply, error)
}

type webPageClient struct {
	cc grpc.ClientConnInterface
}

func NewWebPageClient(cc grpc.ClientConnInterface) WebPageClient {
	return &webPageClient{cc}
}

func (c *webPageClient) ParseUrl(ctx context.Context, in *ParseUrlRequest, opts ...grpc.CallOption) (*ParseUrlReply, error) {
	out := new(ParseUrlReply)
	err := c.cc.Invoke(ctx, "/webpage.WebPage/ParseUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageClient) ParseUrlForEncoded(ctx context.Context, in *ParseUrlForEncodedRequest, opts ...grpc.CallOption) (*ParseUrlForEncodedReply, error) {
	out := new(ParseUrlForEncodedReply)
	err := c.cc.Invoke(ctx, "/webpage.WebPage/ParseUrlForEncoded", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageClient) ParseUrlAndStore(ctx context.Context, in *ParseUrlAndStoreRequest, opts ...grpc.CallOption) (*ParseUrlAndStoreReply, error) {
	out := new(ParseUrlAndStoreReply)
	err := c.cc.Invoke(ctx, "/webpage.WebPage/ParseUrlAndStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebPageServer is the server API for WebPage service.
// All implementations must embed UnimplementedWebPageServer
// for forward compatibility
type WebPageServer interface {
	// 解析url
	ParseUrl(context.Context, *ParseUrlRequest) (*ParseUrlReply, error)
	// 解析url
	ParseUrlForEncoded(context.Context, *ParseUrlForEncodedRequest) (*ParseUrlForEncodedReply, error)
	ParseUrlAndStore(context.Context, *ParseUrlAndStoreRequest) (*ParseUrlAndStoreReply, error)
	mustEmbedUnimplementedWebPageServer()
}

// UnimplementedWebPageServer must be embedded to have forward compatible implementations.
type UnimplementedWebPageServer struct {
}

func (UnimplementedWebPageServer) ParseUrl(context.Context, *ParseUrlRequest) (*ParseUrlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseUrl not implemented")
}
func (UnimplementedWebPageServer) ParseUrlForEncoded(context.Context, *ParseUrlForEncodedRequest) (*ParseUrlForEncodedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseUrlForEncoded not implemented")
}
func (UnimplementedWebPageServer) ParseUrlAndStore(context.Context, *ParseUrlAndStoreRequest) (*ParseUrlAndStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseUrlAndStore not implemented")
}
func (UnimplementedWebPageServer) mustEmbedUnimplementedWebPageServer() {}

// UnsafeWebPageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebPageServer will
// result in compilation errors.
type UnsafeWebPageServer interface {
	mustEmbedUnimplementedWebPageServer()
}

func RegisterWebPageServer(s grpc.ServiceRegistrar, srv WebPageServer) {
	s.RegisterService(&WebPage_ServiceDesc, srv)
}

func _WebPage_ParseUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServer).ParseUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpage.WebPage/ParseUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServer).ParseUrl(ctx, req.(*ParseUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPage_ParseUrlForEncoded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseUrlForEncodedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServer).ParseUrlForEncoded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpage.WebPage/ParseUrlForEncoded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServer).ParseUrlForEncoded(ctx, req.(*ParseUrlForEncodedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPage_ParseUrlAndStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseUrlAndStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServer).ParseUrlAndStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpage.WebPage/ParseUrlAndStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServer).ParseUrlAndStore(ctx, req.(*ParseUrlAndStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebPage_ServiceDesc is the grpc.ServiceDesc for WebPage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebPage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webpage.WebPage",
	HandlerType: (*WebPageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseUrl",
			Handler:    _WebPage_ParseUrl_Handler,
		},
		{
			MethodName: "ParseUrlForEncoded",
			Handler:    _WebPage_ParseUrlForEncoded_Handler,
		},
		{
			MethodName: "ParseUrlAndStore",
			Handler:    _WebPage_ParseUrlAndStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webpage.proto",
}
