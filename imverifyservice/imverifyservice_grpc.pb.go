// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: imverifyservice.proto

package imverifyservice

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

// ImVerifyServiceClient is the client API for ImVerifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImVerifyServiceClient interface {
	// 获取帐户token
	GetAccountToken(ctx context.Context, in *GetAccountTokenMsg, opts ...grpc.CallOption) (*GetAccountTokenMsgReply, error)
	// 验证帐户token
	VerifyAccountToken(ctx context.Context, in *VerifyAccountTokenMsg, opts ...grpc.CallOption) (*VerifyAccountTokenMsgReply, error)
}

type imVerifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImVerifyServiceClient(cc grpc.ClientConnInterface) ImVerifyServiceClient {
	return &imVerifyServiceClient{cc}
}

func (c *imVerifyServiceClient) GetAccountToken(ctx context.Context, in *GetAccountTokenMsg, opts ...grpc.CallOption) (*GetAccountTokenMsgReply, error) {
	out := new(GetAccountTokenMsgReply)
	err := c.cc.Invoke(ctx, "/imverifyservice.ImVerifyService/GetAccountToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imVerifyServiceClient) VerifyAccountToken(ctx context.Context, in *VerifyAccountTokenMsg, opts ...grpc.CallOption) (*VerifyAccountTokenMsgReply, error) {
	out := new(VerifyAccountTokenMsgReply)
	err := c.cc.Invoke(ctx, "/imverifyservice.ImVerifyService/VerifyAccountToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImVerifyServiceServer is the server API for ImVerifyService service.
// All implementations must embed UnimplementedImVerifyServiceServer
// for forward compatibility
type ImVerifyServiceServer interface {
	// 获取帐户token
	GetAccountToken(context.Context, *GetAccountTokenMsg) (*GetAccountTokenMsgReply, error)
	// 验证帐户token
	VerifyAccountToken(context.Context, *VerifyAccountTokenMsg) (*VerifyAccountTokenMsgReply, error)
	mustEmbedUnimplementedImVerifyServiceServer()
}

// UnimplementedImVerifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImVerifyServiceServer struct {
}

func (UnimplementedImVerifyServiceServer) GetAccountToken(context.Context, *GetAccountTokenMsg) (*GetAccountTokenMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountToken not implemented")
}
func (UnimplementedImVerifyServiceServer) VerifyAccountToken(context.Context, *VerifyAccountTokenMsg) (*VerifyAccountTokenMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccountToken not implemented")
}
func (UnimplementedImVerifyServiceServer) mustEmbedUnimplementedImVerifyServiceServer() {}

// UnsafeImVerifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImVerifyServiceServer will
// result in compilation errors.
type UnsafeImVerifyServiceServer interface {
	mustEmbedUnimplementedImVerifyServiceServer()
}

func RegisterImVerifyServiceServer(s grpc.ServiceRegistrar, srv ImVerifyServiceServer) {
	s.RegisterService(&ImVerifyService_ServiceDesc, srv)
}

func _ImVerifyService_GetAccountToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImVerifyServiceServer).GetAccountToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imverifyservice.ImVerifyService/GetAccountToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImVerifyServiceServer).GetAccountToken(ctx, req.(*GetAccountTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImVerifyService_VerifyAccountToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccountTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImVerifyServiceServer).VerifyAccountToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imverifyservice.ImVerifyService/VerifyAccountToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImVerifyServiceServer).VerifyAccountToken(ctx, req.(*VerifyAccountTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// ImVerifyService_ServiceDesc is the grpc.ServiceDesc for ImVerifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImVerifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imverifyservice.ImVerifyService",
	HandlerType: (*ImVerifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountToken",
			Handler:    _ImVerifyService_GetAccountToken_Handler,
		},
		{
			MethodName: "VerifyAccountToken",
			Handler:    _ImVerifyService_VerifyAccountToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imverifyservice.proto",
}
