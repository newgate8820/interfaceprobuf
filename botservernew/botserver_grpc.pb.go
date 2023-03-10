// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: botserver.proto

package pbbotserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	advertisemonitornew "interfaceprobuf/advertisemonitornew"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BotServerClient is the client API for BotServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotServerClient interface {
	// 離開群組
	LeaveChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 獲取群訊息
	GetChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 獲取群成員數目
	GetChatMembersCount(ctx context.Context, in *ReqGetChatMembersCount, opts ...grpc.CallOption) (*ReplyGetChatMembersCount, error)
	// 獲取群管理員
	GetChatAdministrators(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 修改群名字
	SetChatTitle(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 修改群簡介
	SetChatDescription(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 置顶群消息
	UpdatePinnedChannelMessage(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	//    // 取消置顶群消息
	//    rpc UpinChatMessage (ReqData) returns (ReplyData) {
	//    }
	// 發送Action
	SendChatAction(ctx context.Context, in *ReqSendChatAction, opts ...grpc.CallOption) (*ReplyData, error)
	// 發送名片
	SendContact(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 刪除消息
	DeleteMessage(ctx context.Context, in *ReqDeleteMessage, opts ...grpc.CallOption) (*ReplyData, error)
	// GetGroups 获取群组消息
	GetGroups(ctx context.Context, in *ReqGetGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error)
	// GetGroups 获取使用者是管理員的群組消息
	GetAdminGroups(ctx context.Context, in *ReqGetAdminGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error)
	// 新增手動禁言名單
	SetManualBlockList(ctx context.Context, in *advertisemonitornew.ReqSetManualBlockList, opts ...grpc.CallOption) (*advertisemonitornew.ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(ctx context.Context, in *advertisemonitornew.ReqDelManualBlockList, opts ...grpc.CallOption) (*advertisemonitornew.ReplyData, error)
}

type botServerClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServerClient(cc grpc.ClientConnInterface) BotServerClient {
	return &botServerClient{cc}
}

func (c *botServerClient) LeaveChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/LeaveChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChatMembersCount(ctx context.Context, in *ReqGetChatMembersCount, opts ...grpc.CallOption) (*ReplyGetChatMembersCount, error) {
	out := new(ReplyGetChatMembersCount)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/GetChatMembersCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChatAdministrators(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/GetChatAdministrators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetChatTitle(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/SetChatTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetChatDescription(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/SetChatDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) UpdatePinnedChannelMessage(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/UpdatePinnedChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SendChatAction(ctx context.Context, in *ReqSendChatAction, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/SendChatAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SendContact(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/SendContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) DeleteMessage(ctx context.Context, in *ReqDeleteMessage, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/DeleteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetGroups(ctx context.Context, in *ReqGetGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error) {
	out := new(ReplyGetGroups)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetAdminGroups(ctx context.Context, in *ReqGetAdminGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error) {
	out := new(ReplyGetGroups)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/GetAdminGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetManualBlockList(ctx context.Context, in *advertisemonitornew.ReqSetManualBlockList, opts ...grpc.CallOption) (*advertisemonitornew.ReplyData, error) {
	out := new(advertisemonitornew.ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/SetManualBlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) DelManualBlockList(ctx context.Context, in *advertisemonitornew.ReqDelManualBlockList, opts ...grpc.CallOption) (*advertisemonitornew.ReplyData, error) {
	out := new(advertisemonitornew.ReplyData)
	err := c.cc.Invoke(ctx, "/pbbotserver.BotServer/DelManualBlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServerServer is the server API for BotServer service.
// All implementations must embed UnimplementedBotServerServer
// for forward compatibility
type BotServerServer interface {
	// 離開群組
	LeaveChat(context.Context, *ReqData) (*ReplyData, error)
	// 獲取群訊息
	GetChat(context.Context, *ReqData) (*ReplyData, error)
	// 獲取群成員數目
	GetChatMembersCount(context.Context, *ReqGetChatMembersCount) (*ReplyGetChatMembersCount, error)
	// 獲取群管理員
	GetChatAdministrators(context.Context, *ReqData) (*ReplyData, error)
	// 修改群名字
	SetChatTitle(context.Context, *ReqData) (*ReplyData, error)
	// 修改群簡介
	SetChatDescription(context.Context, *ReqData) (*ReplyData, error)
	// 置顶群消息
	UpdatePinnedChannelMessage(context.Context, *ReqData) (*ReplyData, error)
	//    // 取消置顶群消息
	//    rpc UpinChatMessage (ReqData) returns (ReplyData) {
	//    }
	// 發送Action
	SendChatAction(context.Context, *ReqSendChatAction) (*ReplyData, error)
	// 發送名片
	SendContact(context.Context, *ReqData) (*ReplyData, error)
	// 刪除消息
	DeleteMessage(context.Context, *ReqDeleteMessage) (*ReplyData, error)
	// GetGroups 获取群组消息
	GetGroups(context.Context, *ReqGetGroups) (*ReplyGetGroups, error)
	// GetGroups 获取使用者是管理員的群組消息
	GetAdminGroups(context.Context, *ReqGetAdminGroups) (*ReplyGetGroups, error)
	// 新增手動禁言名單
	SetManualBlockList(context.Context, *advertisemonitornew.ReqSetManualBlockList) (*advertisemonitornew.ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(context.Context, *advertisemonitornew.ReqDelManualBlockList) (*advertisemonitornew.ReplyData, error)
	mustEmbedUnimplementedBotServerServer()
}

// UnimplementedBotServerServer must be embedded to have forward compatible implementations.
type UnimplementedBotServerServer struct {
}

func (UnimplementedBotServerServer) LeaveChat(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChat not implemented")
}
func (UnimplementedBotServerServer) GetChat(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedBotServerServer) GetChatMembersCount(context.Context, *ReqGetChatMembersCount) (*ReplyGetChatMembersCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMembersCount not implemented")
}
func (UnimplementedBotServerServer) GetChatAdministrators(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatAdministrators not implemented")
}
func (UnimplementedBotServerServer) SetChatTitle(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatTitle not implemented")
}
func (UnimplementedBotServerServer) SetChatDescription(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatDescription not implemented")
}
func (UnimplementedBotServerServer) UpdatePinnedChannelMessage(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePinnedChannelMessage not implemented")
}
func (UnimplementedBotServerServer) SendChatAction(context.Context, *ReqSendChatAction) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatAction not implemented")
}
func (UnimplementedBotServerServer) SendContact(context.Context, *ReqData) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendContact not implemented")
}
func (UnimplementedBotServerServer) DeleteMessage(context.Context, *ReqDeleteMessage) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedBotServerServer) GetGroups(context.Context, *ReqGetGroups) (*ReplyGetGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedBotServerServer) GetAdminGroups(context.Context, *ReqGetAdminGroups) (*ReplyGetGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminGroups not implemented")
}
func (UnimplementedBotServerServer) SetManualBlockList(context.Context, *advertisemonitornew.ReqSetManualBlockList) (*advertisemonitornew.ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetManualBlockList not implemented")
}
func (UnimplementedBotServerServer) DelManualBlockList(context.Context, *advertisemonitornew.ReqDelManualBlockList) (*advertisemonitornew.ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelManualBlockList not implemented")
}
func (UnimplementedBotServerServer) mustEmbedUnimplementedBotServerServer() {}

// UnsafeBotServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServerServer will
// result in compilation errors.
type UnsafeBotServerServer interface {
	mustEmbedUnimplementedBotServerServer()
}

func RegisterBotServerServer(s grpc.ServiceRegistrar, srv BotServerServer) {
	s.RegisterService(&BotServer_ServiceDesc, srv)
}

func _BotServer_LeaveChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).LeaveChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/LeaveChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).LeaveChat(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChat(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChatMembersCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetChatMembersCount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChatMembersCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChatMembersCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChatMembersCount(ctx, req.(*ReqGetChatMembersCount))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChatAdministrators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChatAdministrators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChatAdministrators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChatAdministrators(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetChatTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetChatTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetChatTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetChatTitle(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetChatDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetChatDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetChatDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetChatDescription(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_UpdatePinnedChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).UpdatePinnedChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/UpdatePinnedChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).UpdatePinnedChannelMessage(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SendChatAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSendChatAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SendChatAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SendChatAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SendChatAction(ctx, req.(*ReqSendChatAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SendContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SendContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SendContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SendContact(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/DeleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).DeleteMessage(ctx, req.(*ReqDeleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetGroups(ctx, req.(*ReqGetGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetAdminGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetAdminGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetAdminGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetAdminGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetAdminGroups(ctx, req.(*ReqGetAdminGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(advertisemonitornew.ReqSetManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetManualBlockList(ctx, req.(*advertisemonitornew.ReqSetManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_DelManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(advertisemonitornew.ReqDelManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).DelManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/DelManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).DelManualBlockList(ctx, req.(*advertisemonitornew.ReqDelManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

// BotServer_ServiceDesc is the grpc.ServiceDesc for BotServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbbotserver.BotServer",
	HandlerType: (*BotServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaveChat",
			Handler:    _BotServer_LeaveChat_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _BotServer_GetChat_Handler,
		},
		{
			MethodName: "GetChatMembersCount",
			Handler:    _BotServer_GetChatMembersCount_Handler,
		},
		{
			MethodName: "GetChatAdministrators",
			Handler:    _BotServer_GetChatAdministrators_Handler,
		},
		{
			MethodName: "SetChatTitle",
			Handler:    _BotServer_SetChatTitle_Handler,
		},
		{
			MethodName: "SetChatDescription",
			Handler:    _BotServer_SetChatDescription_Handler,
		},
		{
			MethodName: "UpdatePinnedChannelMessage",
			Handler:    _BotServer_UpdatePinnedChannelMessage_Handler,
		},
		{
			MethodName: "SendChatAction",
			Handler:    _BotServer_SendChatAction_Handler,
		},
		{
			MethodName: "SendContact",
			Handler:    _BotServer_SendContact_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _BotServer_DeleteMessage_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _BotServer_GetGroups_Handler,
		},
		{
			MethodName: "GetAdminGroups",
			Handler:    _BotServer_GetAdminGroups_Handler,
		},
		{
			MethodName: "SetManualBlockList",
			Handler:    _BotServer_SetManualBlockList_Handler,
		},
		{
			MethodName: "DelManualBlockList",
			Handler:    _BotServer_DelManualBlockList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "botserver.proto",
}
