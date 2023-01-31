// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: gameapi.proto

package gameapi

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

// GameapiClient is the client API for Gameapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameapiClient interface {
	// 游戏登录
	Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ReplyLogin, error)
}

type gameapiClient struct {
	cc grpc.ClientConnInterface
}

func NewGameapiClient(cc grpc.ClientConnInterface) GameapiClient {
	return &gameapiClient{cc}
}

func (c *gameapiClient) Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ReplyLogin, error) {
	out := new(ReplyLogin)
	err := c.cc.Invoke(ctx, "/gameapi.gameapi/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameapiServer is the server API for Gameapi service.
// All implementations must embed UnimplementedGameapiServer
// for forward compatibility
type GameapiServer interface {
	// 游戏登录
	Login(context.Context, *ReqLogin) (*ReplyLogin, error)
	mustEmbedUnimplementedGameapiServer()
}

// UnimplementedGameapiServer must be embedded to have forward compatible implementations.
type UnimplementedGameapiServer struct {
}

func (UnimplementedGameapiServer) Login(context.Context, *ReqLogin) (*ReplyLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGameapiServer) mustEmbedUnimplementedGameapiServer() {}

// UnsafeGameapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameapiServer will
// result in compilation errors.
type UnsafeGameapiServer interface {
	mustEmbedUnimplementedGameapiServer()
}

func RegisterGameapiServer(s grpc.ServiceRegistrar, srv GameapiServer) {
	s.RegisterService(&Gameapi_ServiceDesc, srv)
}

func _Gameapi_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameapiServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameapi.gameapi/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameapiServer).Login(ctx, req.(*ReqLogin))
	}
	return interceptor(ctx, in, info, handler)
}

// Gameapi_ServiceDesc is the grpc.ServiceDesc for Gameapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gameapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gameapi.gameapi",
	HandlerType: (*GameapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Gameapi_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameapi.proto",
}
