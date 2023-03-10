// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: updateservice.proto

package updateservice

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
	// GetOldBots 获取旧机器人
	GetOldBots(ctx context.Context, in *GetOldBotsRequest, opts ...grpc.CallOption) (*GetOldBotsReply, error)
	// PushMessage 推送消息
	PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error)
	// PushMessageList 推送消息
	PushMessageList(ctx context.Context, in *PushMessageListRequest, opts ...grpc.CallOption) (*PushMessageReply, error)
	// PushCallbackQuery 推送回调查询
	PushCallbackQuery(ctx context.Context, in *CallbackQueryRequest, opts ...grpc.CallOption) (*CallbackQueryReply, error)
	// GetInlineQuery 获取内联查询
	GetInlineQuery(ctx context.Context, in *GetInlineQueryRequest, opts ...grpc.CallOption) (*GetInlineQueryReply, error)
	// SendInlineQuery 发送内联查询
	SendInlineQuery(ctx context.Context, in *SendInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQuery(ctx context.Context, in *PushInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQueryV2(ctx context.Context, in *PushInlineQueryRequestV2, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
}

type botClient struct {
	cc grpc.ClientConnInterface
}

func NewBotClient(cc grpc.ClientConnInterface) BotClient {
	return &botClient{cc}
}

func (c *botClient) GetOldBots(ctx context.Context, in *GetOldBotsRequest, opts ...grpc.CallOption) (*GetOldBotsReply, error) {
	out := new(GetOldBotsReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/GetOldBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error) {
	out := new(PushMessageReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushMessageList(ctx context.Context, in *PushMessageListRequest, opts ...grpc.CallOption) (*PushMessageReply, error) {
	out := new(PushMessageReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/PushMessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushCallbackQuery(ctx context.Context, in *CallbackQueryRequest, opts ...grpc.CallOption) (*CallbackQueryReply, error) {
	out := new(CallbackQueryReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/PushCallbackQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) GetInlineQuery(ctx context.Context, in *GetInlineQueryRequest, opts ...grpc.CallOption) (*GetInlineQueryReply, error) {
	out := new(GetInlineQueryReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/GetInlineQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) SendInlineQuery(ctx context.Context, in *SendInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/SendInlineQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushInlineQuery(ctx context.Context, in *PushInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/PushInlineQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushInlineQueryV2(ctx context.Context, in *PushInlineQueryRequestV2, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := c.cc.Invoke(ctx, "/updateservice.Bot/PushInlineQueryV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServer is the server API for Bot service.
// All implementations must embed UnimplementedBotServer
// for forward compatibility
type BotServer interface {
	// GetOldBots 获取旧机器人
	GetOldBots(context.Context, *GetOldBotsRequest) (*GetOldBotsReply, error)
	// PushMessage 推送消息
	PushMessage(context.Context, *PushMessageRequest) (*PushMessageReply, error)
	// PushMessageList 推送消息
	PushMessageList(context.Context, *PushMessageListRequest) (*PushMessageReply, error)
	// PushCallbackQuery 推送回调查询
	PushCallbackQuery(context.Context, *CallbackQueryRequest) (*CallbackQueryReply, error)
	// GetInlineQuery 获取内联查询
	GetInlineQuery(context.Context, *GetInlineQueryRequest) (*GetInlineQueryReply, error)
	// SendInlineQuery 发送内联查询
	SendInlineQuery(context.Context, *SendInlineQueryRequest) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQuery(context.Context, *PushInlineQueryRequest) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQueryV2(context.Context, *PushInlineQueryRequestV2) (*SendInlineQueryReply, error)
	mustEmbedUnimplementedBotServer()
}

// UnimplementedBotServer must be embedded to have forward compatible implementations.
type UnimplementedBotServer struct {
}

func (UnimplementedBotServer) GetOldBots(context.Context, *GetOldBotsRequest) (*GetOldBotsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOldBots not implemented")
}
func (UnimplementedBotServer) PushMessage(context.Context, *PushMessageRequest) (*PushMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}
func (UnimplementedBotServer) PushMessageList(context.Context, *PushMessageListRequest) (*PushMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessageList not implemented")
}
func (UnimplementedBotServer) PushCallbackQuery(context.Context, *CallbackQueryRequest) (*CallbackQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCallbackQuery not implemented")
}
func (UnimplementedBotServer) GetInlineQuery(context.Context, *GetInlineQueryRequest) (*GetInlineQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInlineQuery not implemented")
}
func (UnimplementedBotServer) SendInlineQuery(context.Context, *SendInlineQueryRequest) (*SendInlineQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendInlineQuery not implemented")
}
func (UnimplementedBotServer) PushInlineQuery(context.Context, *PushInlineQueryRequest) (*SendInlineQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushInlineQuery not implemented")
}
func (UnimplementedBotServer) PushInlineQueryV2(context.Context, *PushInlineQueryRequestV2) (*SendInlineQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushInlineQueryV2 not implemented")
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

func _Bot_GetOldBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOldBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetOldBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/GetOldBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetOldBots(ctx, req.(*GetOldBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushMessage(ctx, req.(*PushMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushMessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushMessageList(ctx, req.(*PushMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushCallbackQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushCallbackQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushCallbackQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushCallbackQuery(ctx, req.(*CallbackQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_GetInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/GetInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetInlineQuery(ctx, req.(*GetInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_SendInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).SendInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/SendInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).SendInlineQuery(ctx, req.(*SendInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushInlineQuery(ctx, req.(*PushInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushInlineQueryV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushInlineQueryRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushInlineQueryV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushInlineQueryV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushInlineQueryV2(ctx, req.(*PushInlineQueryRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

// Bot_ServiceDesc is the grpc.ServiceDesc for Bot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "updateservice.Bot",
	HandlerType: (*BotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOldBots",
			Handler:    _Bot_GetOldBots_Handler,
		},
		{
			MethodName: "PushMessage",
			Handler:    _Bot_PushMessage_Handler,
		},
		{
			MethodName: "PushMessageList",
			Handler:    _Bot_PushMessageList_Handler,
		},
		{
			MethodName: "PushCallbackQuery",
			Handler:    _Bot_PushCallbackQuery_Handler,
		},
		{
			MethodName: "GetInlineQuery",
			Handler:    _Bot_GetInlineQuery_Handler,
		},
		{
			MethodName: "SendInlineQuery",
			Handler:    _Bot_SendInlineQuery_Handler,
		},
		{
			MethodName: "PushInlineQuery",
			Handler:    _Bot_PushInlineQuery_Handler,
		},
		{
			MethodName: "PushInlineQueryV2",
			Handler:    _Bot_PushInlineQueryV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updateservice.proto",
}
