// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tokenserver.proto

package imtokenservice

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

// ImTokenServiceClient is the client API for ImTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImTokenServiceClient interface {
	// 获取token
	GetToken(ctx context.Context, in *ReqGetTokenMsg, opts ...grpc.CallOption) (*GetTokenReply, error)
	// 刷新token
	RefreshToken(ctx context.Context, in *ReqRefreshTokenMsg, opts ...grpc.CallOption) (*RefreshTokenReply, error)
	// 验证token
	ValidateToken(ctx context.Context, in *ReqValidateTokenMsg, opts ...grpc.CallOption) (*ValidateTokenReply, error)
	// 通过token 获取用户信息
	GetUInfoByToken(ctx context.Context, in *ReqGetUInfoByTokenMsg, opts ...grpc.CallOption) (*GetUInfoByTokenReply, error)
}

type imTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImTokenServiceClient(cc grpc.ClientConnInterface) ImTokenServiceClient {
	return &imTokenServiceClient{cc}
}

func (c *imTokenServiceClient) GetToken(ctx context.Context, in *ReqGetTokenMsg, opts ...grpc.CallOption) (*GetTokenReply, error) {
	out := new(GetTokenReply)
	err := c.cc.Invoke(ctx, "/imtokenservice.ImTokenService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imTokenServiceClient) RefreshToken(ctx context.Context, in *ReqRefreshTokenMsg, opts ...grpc.CallOption) (*RefreshTokenReply, error) {
	out := new(RefreshTokenReply)
	err := c.cc.Invoke(ctx, "/imtokenservice.ImTokenService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imTokenServiceClient) ValidateToken(ctx context.Context, in *ReqValidateTokenMsg, opts ...grpc.CallOption) (*ValidateTokenReply, error) {
	out := new(ValidateTokenReply)
	err := c.cc.Invoke(ctx, "/imtokenservice.ImTokenService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imTokenServiceClient) GetUInfoByToken(ctx context.Context, in *ReqGetUInfoByTokenMsg, opts ...grpc.CallOption) (*GetUInfoByTokenReply, error) {
	out := new(GetUInfoByTokenReply)
	err := c.cc.Invoke(ctx, "/imtokenservice.ImTokenService/GetUInfoByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImTokenServiceServer is the server API for ImTokenService service.
// All implementations must embed UnimplementedImTokenServiceServer
// for forward compatibility
type ImTokenServiceServer interface {
	// 获取token
	GetToken(context.Context, *ReqGetTokenMsg) (*GetTokenReply, error)
	// 刷新token
	RefreshToken(context.Context, *ReqRefreshTokenMsg) (*RefreshTokenReply, error)
	// 验证token
	ValidateToken(context.Context, *ReqValidateTokenMsg) (*ValidateTokenReply, error)
	// 通过token 获取用户信息
	GetUInfoByToken(context.Context, *ReqGetUInfoByTokenMsg) (*GetUInfoByTokenReply, error)
	mustEmbedUnimplementedImTokenServiceServer()
}

// UnimplementedImTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImTokenServiceServer struct {
}

func (UnimplementedImTokenServiceServer) GetToken(context.Context, *ReqGetTokenMsg) (*GetTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedImTokenServiceServer) RefreshToken(context.Context, *ReqRefreshTokenMsg) (*RefreshTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedImTokenServiceServer) ValidateToken(context.Context, *ReqValidateTokenMsg) (*ValidateTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedImTokenServiceServer) GetUInfoByToken(context.Context, *ReqGetUInfoByTokenMsg) (*GetUInfoByTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUInfoByToken not implemented")
}
func (UnimplementedImTokenServiceServer) mustEmbedUnimplementedImTokenServiceServer() {}

// UnsafeImTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImTokenServiceServer will
// result in compilation errors.
type UnsafeImTokenServiceServer interface {
	mustEmbedUnimplementedImTokenServiceServer()
}

func RegisterImTokenServiceServer(s grpc.ServiceRegistrar, srv ImTokenServiceServer) {
	s.RegisterService(&ImTokenService_ServiceDesc, srv)
}

func _ImTokenService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImTokenServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imtokenservice.ImTokenService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImTokenServiceServer).GetToken(ctx, req.(*ReqGetTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImTokenService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqRefreshTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImTokenServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imtokenservice.ImTokenService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImTokenServiceServer).RefreshToken(ctx, req.(*ReqRefreshTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImTokenService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqValidateTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImTokenServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imtokenservice.ImTokenService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImTokenServiceServer).ValidateToken(ctx, req.(*ReqValidateTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImTokenService_GetUInfoByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetUInfoByTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImTokenServiceServer).GetUInfoByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imtokenservice.ImTokenService/GetUInfoByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImTokenServiceServer).GetUInfoByToken(ctx, req.(*ReqGetUInfoByTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// ImTokenService_ServiceDesc is the grpc.ServiceDesc for ImTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imtokenservice.ImTokenService",
	HandlerType: (*ImTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _ImTokenService_GetToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _ImTokenService_RefreshToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _ImTokenService_ValidateToken_Handler,
		},
		{
			MethodName: "GetUInfoByToken",
			Handler:    _ImTokenService_GetUInfoByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenserver.proto",
}
