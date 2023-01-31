// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: advertisemonitor.proto

package advertisemonitor

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

// AdvertiseMonitorClient is the client API for AdvertiseMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertiseMonitorClient interface {
	//獲取所有禁言用戶
	GetBlockList(ctx context.Context, in *ReqGetBlockList, opts ...grpc.CallOption) (*ReplyGetBlockList, error)
	// 解除單個用户限制
	UnBlockSingleUser(ctx context.Context, in *ReqUnBlockBannedUser, opts ...grpc.CallOption) (*ReplyData, error)
	// 解除所有用户限制
	UnBlockAllUser(ctx context.Context, in *ReqUnBlockAllUser, opts ...grpc.CallOption) (*ReplyData, error)
	// 獲取群成員消息發送條件限制
	GetSendMessageRuleSettings(ctx context.Context, in *ReqGetSendMessageRuleSettings, opts ...grpc.CallOption) (*ReplyGetSendMessageRuleSettings, error)
	// 設定某個群的發消息規則
	UpdateMessageSendingLimit(ctx context.Context, in *ReqUpdateMessageSendingLimit, opts ...grpc.CallOption) (*ReplyData, error)
	// 設定某個群的發消息規則
	UpdateMessageMaskWord(ctx context.Context, in *ReqUpdateMessageMaskWord, opts ...grpc.CallOption) (*ReplyData, error)
	// 新增手動禁言名單
	SetManualBlockList(ctx context.Context, in *ReqSetManualBlockList, opts ...grpc.CallOption) (*ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(ctx context.Context, in *ReqDelManualBlockList, opts ...grpc.CallOption) (*ReplyData, error)
}

type advertiseMonitorClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertiseMonitorClient(cc grpc.ClientConnInterface) AdvertiseMonitorClient {
	return &advertiseMonitorClient{cc}
}

func (c *advertiseMonitorClient) GetBlockList(ctx context.Context, in *ReqGetBlockList, opts ...grpc.CallOption) (*ReplyGetBlockList, error) {
	out := new(ReplyGetBlockList)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/GetBlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) UnBlockSingleUser(ctx context.Context, in *ReqUnBlockBannedUser, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/UnBlockSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) UnBlockAllUser(ctx context.Context, in *ReqUnBlockAllUser, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/UnBlockAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) GetSendMessageRuleSettings(ctx context.Context, in *ReqGetSendMessageRuleSettings, opts ...grpc.CallOption) (*ReplyGetSendMessageRuleSettings, error) {
	out := new(ReplyGetSendMessageRuleSettings)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/GetSendMessageRuleSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) UpdateMessageSendingLimit(ctx context.Context, in *ReqUpdateMessageSendingLimit, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/UpdateMessageSendingLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) UpdateMessageMaskWord(ctx context.Context, in *ReqUpdateMessageMaskWord, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/UpdateMessageMaskWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) SetManualBlockList(ctx context.Context, in *ReqSetManualBlockList, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/SetManualBlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiseMonitorClient) DelManualBlockList(ctx context.Context, in *ReqDelManualBlockList, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := c.cc.Invoke(ctx, "/advertisemonitor.AdvertiseMonitor/DelManualBlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertiseMonitorServer is the server API for AdvertiseMonitor service.
// All implementations must embed UnimplementedAdvertiseMonitorServer
// for forward compatibility
type AdvertiseMonitorServer interface {
	//獲取所有禁言用戶
	GetBlockList(context.Context, *ReqGetBlockList) (*ReplyGetBlockList, error)
	// 解除單個用户限制
	UnBlockSingleUser(context.Context, *ReqUnBlockBannedUser) (*ReplyData, error)
	// 解除所有用户限制
	UnBlockAllUser(context.Context, *ReqUnBlockAllUser) (*ReplyData, error)
	// 獲取群成員消息發送條件限制
	GetSendMessageRuleSettings(context.Context, *ReqGetSendMessageRuleSettings) (*ReplyGetSendMessageRuleSettings, error)
	// 設定某個群的發消息規則
	UpdateMessageSendingLimit(context.Context, *ReqUpdateMessageSendingLimit) (*ReplyData, error)
	// 設定某個群的發消息規則
	UpdateMessageMaskWord(context.Context, *ReqUpdateMessageMaskWord) (*ReplyData, error)
	// 新增手動禁言名單
	SetManualBlockList(context.Context, *ReqSetManualBlockList) (*ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(context.Context, *ReqDelManualBlockList) (*ReplyData, error)
	mustEmbedUnimplementedAdvertiseMonitorServer()
}

// UnimplementedAdvertiseMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedAdvertiseMonitorServer struct {
}

func (UnimplementedAdvertiseMonitorServer) GetBlockList(context.Context, *ReqGetBlockList) (*ReplyGetBlockList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockList not implemented")
}
func (UnimplementedAdvertiseMonitorServer) UnBlockSingleUser(context.Context, *ReqUnBlockBannedUser) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockSingleUser not implemented")
}
func (UnimplementedAdvertiseMonitorServer) UnBlockAllUser(context.Context, *ReqUnBlockAllUser) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockAllUser not implemented")
}
func (UnimplementedAdvertiseMonitorServer) GetSendMessageRuleSettings(context.Context, *ReqGetSendMessageRuleSettings) (*ReplyGetSendMessageRuleSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendMessageRuleSettings not implemented")
}
func (UnimplementedAdvertiseMonitorServer) UpdateMessageSendingLimit(context.Context, *ReqUpdateMessageSendingLimit) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageSendingLimit not implemented")
}
func (UnimplementedAdvertiseMonitorServer) UpdateMessageMaskWord(context.Context, *ReqUpdateMessageMaskWord) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageMaskWord not implemented")
}
func (UnimplementedAdvertiseMonitorServer) SetManualBlockList(context.Context, *ReqSetManualBlockList) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetManualBlockList not implemented")
}
func (UnimplementedAdvertiseMonitorServer) DelManualBlockList(context.Context, *ReqDelManualBlockList) (*ReplyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelManualBlockList not implemented")
}
func (UnimplementedAdvertiseMonitorServer) mustEmbedUnimplementedAdvertiseMonitorServer() {}

// UnsafeAdvertiseMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertiseMonitorServer will
// result in compilation errors.
type UnsafeAdvertiseMonitorServer interface {
	mustEmbedUnimplementedAdvertiseMonitorServer()
}

func RegisterAdvertiseMonitorServer(s grpc.ServiceRegistrar, srv AdvertiseMonitorServer) {
	s.RegisterService(&AdvertiseMonitor_ServiceDesc, srv)
}

func _AdvertiseMonitor_GetBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).GetBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/GetBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).GetBlockList(ctx, req.(*ReqGetBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_UnBlockSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUnBlockBannedUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).UnBlockSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/UnBlockSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).UnBlockSingleUser(ctx, req.(*ReqUnBlockBannedUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_UnBlockAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUnBlockAllUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).UnBlockAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/UnBlockAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).UnBlockAllUser(ctx, req.(*ReqUnBlockAllUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_GetSendMessageRuleSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetSendMessageRuleSettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).GetSendMessageRuleSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/GetSendMessageRuleSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).GetSendMessageRuleSettings(ctx, req.(*ReqGetSendMessageRuleSettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_UpdateMessageSendingLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateMessageSendingLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).UpdateMessageSendingLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/UpdateMessageSendingLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).UpdateMessageSendingLimit(ctx, req.(*ReqUpdateMessageSendingLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_UpdateMessageMaskWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateMessageMaskWord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).UpdateMessageMaskWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/UpdateMessageMaskWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).UpdateMessageMaskWord(ctx, req.(*ReqUpdateMessageMaskWord))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_SetManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSetManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).SetManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/SetManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).SetManualBlockList(ctx, req.(*ReqSetManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertiseMonitor_DelManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDelManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiseMonitorServer).DelManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisemonitor.AdvertiseMonitor/DelManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiseMonitorServer).DelManualBlockList(ctx, req.(*ReqDelManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

// AdvertiseMonitor_ServiceDesc is the grpc.ServiceDesc for AdvertiseMonitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdvertiseMonitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "advertisemonitor.AdvertiseMonitor",
	HandlerType: (*AdvertiseMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockList",
			Handler:    _AdvertiseMonitor_GetBlockList_Handler,
		},
		{
			MethodName: "UnBlockSingleUser",
			Handler:    _AdvertiseMonitor_UnBlockSingleUser_Handler,
		},
		{
			MethodName: "UnBlockAllUser",
			Handler:    _AdvertiseMonitor_UnBlockAllUser_Handler,
		},
		{
			MethodName: "GetSendMessageRuleSettings",
			Handler:    _AdvertiseMonitor_GetSendMessageRuleSettings_Handler,
		},
		{
			MethodName: "UpdateMessageSendingLimit",
			Handler:    _AdvertiseMonitor_UpdateMessageSendingLimit_Handler,
		},
		{
			MethodName: "UpdateMessageMaskWord",
			Handler:    _AdvertiseMonitor_UpdateMessageMaskWord_Handler,
		},
		{
			MethodName: "SetManualBlockList",
			Handler:    _AdvertiseMonitor_SetManualBlockList_Handler,
		},
		{
			MethodName: "DelManualBlockList",
			Handler:    _AdvertiseMonitor_DelManualBlockList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertisemonitor.proto",
}
