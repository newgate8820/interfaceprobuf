// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: userservice.proto

package userservice

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

// BotClient is the client API for Bot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotClient interface {
	// GetBotDialog 获取机器人对话列表的用户ID数组
	GetBotDialogUserIds(ctx context.Context, in *GetBotDialogUserIdsRequest, opts ...grpc.CallOption) (*GetBotDialogUserIdsReply, error)
}

type botClient struct {
	cc grpc.ClientConnInterface
}

func NewBotClient(cc grpc.ClientConnInterface) BotClient {
	return &botClient{cc}
}

func (c *botClient) GetBotDialogUserIds(ctx context.Context, in *GetBotDialogUserIdsRequest, opts ...grpc.CallOption) (*GetBotDialogUserIdsReply, error) {
	out := new(GetBotDialogUserIdsReply)
	err := c.cc.Invoke(ctx, "/userservice.Bot/GetBotDialogUserIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServer is the server API for Bot service.
// All implementations must embed UnimplementedBotServer
// for forward compatibility
type BotServer interface {
	// GetBotDialog 获取机器人对话列表的用户ID数组
	GetBotDialogUserIds(context.Context, *GetBotDialogUserIdsRequest) (*GetBotDialogUserIdsReply, error)
	mustEmbedUnimplementedBotServer()
}

// UnimplementedBotServer must be embedded to have forward compatible implementations.
type UnimplementedBotServer struct {
}

func (UnimplementedBotServer) GetBotDialogUserIds(context.Context, *GetBotDialogUserIdsRequest) (*GetBotDialogUserIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotDialogUserIds not implemented")
}
func (UnimplementedBotServer) mustEmbedUnimplementedBotServer() {}

// UnsafeBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServer will
// result in compilation errors.
type UnsafeBotServer interface {
	mustEmbedUnimplementedBotServer()
}

func RegisterBotServer(s grpc.ServiceRegistrar, srv BotServer) {
	s.RegisterService(&Bot_ServiceDesc, srv)
}

func _Bot_GetBotDialogUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotDialogUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetBotDialogUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.Bot/GetBotDialogUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetBotDialogUserIds(ctx, req.(*GetBotDialogUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bot_ServiceDesc is the grpc.ServiceDesc for Bot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userservice.Bot",
	HandlerType: (*BotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBotDialogUserIds",
			Handler:    _Bot_GetBotDialogUserIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}
