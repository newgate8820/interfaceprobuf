// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: chatservice.proto

package chatserver

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

// GroupServerServiceClient is the client API for GroupServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServerServiceClient interface {
	// ***********************业务接口 **************************
	// 群创建的接口
	CreateChat(ctx context.Context, in *CreateRquest, opts ...grpc.CallOption) (*ChatInfo, error)
	// 根据chat_id获取群的基本信息
	GetChatInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新群名称
	UpdateChatTitle(ctx context.Context, in *ChatTitleReq, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新群名称头像
	UpdateChatPhoto(ctx context.Context, in *ChatPhotoReq, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新群migrate_to
	UpdateChatMigrate(ctx context.Context, in *ChatMigrateReq, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新deactive
	// rpc UpdateChatDeactice (chatdeactiveReq) returns (chatInfo) {
	// }
	// 更新admins_enabled
	UpdateChatAdminsEn(ctx context.Context, in *ChatadminEnableReq, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新invite link
	UpdateChatLink(ctx context.Context, in *ChatExportLink, opts ...grpc.CallOption) (*ChatInfo, error)
	// 更新invite link
	// rpc UpdateChatVersion (chatVersion) returns (chatInfo) {
	// }
	// 更新participants_count
	// rpc UpdateChatParticipantsCount (chatPartCount) returns (chatInfo) {
	// }
	// 删除群信息
	DelChatInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*BoolResult, error)
	// ############### 有关普通群成员的操作#####################
	AddChatUser(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error)
	// 删除群成员
	DelChatUser(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error)
	// 设置管理员
	SetChatUserAdmin(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error)
	// 取消管理员
	DelChatUserAdmin(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error)
	// 机器人获取基本信息
	GetChatLittleInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatLittleInfo, error)
	// 获取群成员列表
	GetChatParticipants(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatParticipants, error)
	// 获取指定成员的信息
	GetChatUserInfo(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*ChatMemberInfo, error)
	// 获取群成员用户id
	GetChatUsersId(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatUsersId, error)
	GetChatInfos(ctx context.Context, in *ChatIdsReq, opts ...grpc.CallOption) (*ChatInfos, error)
	GetParticipants(ctx context.Context, in *ChatIdsReq, opts ...grpc.CallOption) (*ChatParticipants, error)
	GetCommonChats(ctx context.Context, in *CommonChatReq, opts ...grpc.CallOption) (*CommonChatsId, error)
}

type groupServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServerServiceClient(cc grpc.ClientConnInterface) GroupServerServiceClient {
	return &groupServerServiceClient{cc}
}

func (c *groupServerServiceClient) CreateChat(ctx context.Context, in *CreateRquest, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) UpdateChatTitle(ctx context.Context, in *ChatTitleReq, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/UpdateChatTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) UpdateChatPhoto(ctx context.Context, in *ChatPhotoReq, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/UpdateChatPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) UpdateChatMigrate(ctx context.Context, in *ChatMigrateReq, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/UpdateChatMigrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) UpdateChatAdminsEn(ctx context.Context, in *ChatadminEnableReq, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/UpdateChatAdminsEn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) UpdateChatLink(ctx context.Context, in *ChatExportLink, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/UpdateChatLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) DelChatInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/DelChatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) AddChatUser(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/AddChatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) DelChatUser(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/DelChatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) SetChatUserAdmin(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/SetChatUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) DelChatUserAdmin(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/DelChatUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatLittleInfo(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatLittleInfo, error) {
	out := new(ChatLittleInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatLittleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatParticipants(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatParticipants, error) {
	out := new(ChatParticipants)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatUserInfo(ctx context.Context, in *ChatMemberBase, opts ...grpc.CallOption) (*ChatMemberInfo, error) {
	out := new(ChatMemberInfo)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatUsersId(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatUsersId, error) {
	out := new(ChatUsersId)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatUsersId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetChatInfos(ctx context.Context, in *ChatIdsReq, opts ...grpc.CallOption) (*ChatInfos, error) {
	out := new(ChatInfos)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetChatInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetParticipants(ctx context.Context, in *ChatIdsReq, opts ...grpc.CallOption) (*ChatParticipants, error) {
	out := new(ChatParticipants)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServerServiceClient) GetCommonChats(ctx context.Context, in *CommonChatReq, opts ...grpc.CallOption) (*CommonChatsId, error) {
	out := new(CommonChatsId)
	err := c.cc.Invoke(ctx, "/chatserver.GroupServerService/GetCommonChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServerServiceServer is the server API for GroupServerService service.
// All implementations must embed UnimplementedGroupServerServiceServer
// for forward compatibility
type GroupServerServiceServer interface {
	// ***********************业务接口 **************************
	// 群创建的接口
	CreateChat(context.Context, *CreateRquest) (*ChatInfo, error)
	// 根据chat_id获取群的基本信息
	GetChatInfo(context.Context, *ChatId) (*ChatInfo, error)
	// 更新群名称
	UpdateChatTitle(context.Context, *ChatTitleReq) (*ChatInfo, error)
	// 更新群名称头像
	UpdateChatPhoto(context.Context, *ChatPhotoReq) (*ChatInfo, error)
	// 更新群migrate_to
	UpdateChatMigrate(context.Context, *ChatMigrateReq) (*ChatInfo, error)
	// 更新deactive
	// rpc UpdateChatDeactice (chatdeactiveReq) returns (chatInfo) {
	// }
	// 更新admins_enabled
	UpdateChatAdminsEn(context.Context, *ChatadminEnableReq) (*ChatInfo, error)
	// 更新invite link
	UpdateChatLink(context.Context, *ChatExportLink) (*ChatInfo, error)
	// 更新invite link
	// rpc UpdateChatVersion (chatVersion) returns (chatInfo) {
	// }
	// 更新participants_count
	// rpc UpdateChatParticipantsCount (chatPartCount) returns (chatInfo) {
	// }
	// 删除群信息
	DelChatInfo(context.Context, *ChatId) (*BoolResult, error)
	// ############### 有关普通群成员的操作#####################
	AddChatUser(context.Context, *ChatMemberBase) (*BoolResult, error)
	// 删除群成员
	DelChatUser(context.Context, *ChatMemberBase) (*BoolResult, error)
	// 设置管理员
	SetChatUserAdmin(context.Context, *ChatMemberBase) (*BoolResult, error)
	// 取消管理员
	DelChatUserAdmin(context.Context, *ChatMemberBase) (*BoolResult, error)
	// 机器人获取基本信息
	GetChatLittleInfo(context.Context, *ChatId) (*ChatLittleInfo, error)
	// 获取群成员列表
	GetChatParticipants(context.Context, *ChatId) (*ChatParticipants, error)
	// 获取指定成员的信息
	GetChatUserInfo(context.Context, *ChatMemberBase) (*ChatMemberInfo, error)
	// 获取群成员用户id
	GetChatUsersId(context.Context, *ChatId) (*ChatUsersId, error)
	GetChatInfos(context.Context, *ChatIdsReq) (*ChatInfos, error)
	GetParticipants(context.Context, *ChatIdsReq) (*ChatParticipants, error)
	GetCommonChats(context.Context, *CommonChatReq) (*CommonChatsId, error)
	mustEmbedUnimplementedGroupServerServiceServer()
}

// UnimplementedGroupServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServerServiceServer struct {
}

func (UnimplementedGroupServerServiceServer) CreateChat(context.Context, *CreateRquest) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatInfo(context.Context, *ChatId) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatInfo not implemented")
}
func (UnimplementedGroupServerServiceServer) UpdateChatTitle(context.Context, *ChatTitleReq) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatTitle not implemented")
}
func (UnimplementedGroupServerServiceServer) UpdateChatPhoto(context.Context, *ChatPhotoReq) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatPhoto not implemented")
}
func (UnimplementedGroupServerServiceServer) UpdateChatMigrate(context.Context, *ChatMigrateReq) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatMigrate not implemented")
}
func (UnimplementedGroupServerServiceServer) UpdateChatAdminsEn(context.Context, *ChatadminEnableReq) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatAdminsEn not implemented")
}
func (UnimplementedGroupServerServiceServer) UpdateChatLink(context.Context, *ChatExportLink) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatLink not implemented")
}
func (UnimplementedGroupServerServiceServer) DelChatInfo(context.Context, *ChatId) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelChatInfo not implemented")
}
func (UnimplementedGroupServerServiceServer) AddChatUser(context.Context, *ChatMemberBase) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChatUser not implemented")
}
func (UnimplementedGroupServerServiceServer) DelChatUser(context.Context, *ChatMemberBase) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelChatUser not implemented")
}
func (UnimplementedGroupServerServiceServer) SetChatUserAdmin(context.Context, *ChatMemberBase) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatUserAdmin not implemented")
}
func (UnimplementedGroupServerServiceServer) DelChatUserAdmin(context.Context, *ChatMemberBase) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelChatUserAdmin not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatLittleInfo(context.Context, *ChatId) (*ChatLittleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatLittleInfo not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatParticipants(context.Context, *ChatId) (*ChatParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatParticipants not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatUserInfo(context.Context, *ChatMemberBase) (*ChatMemberInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatUserInfo not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatUsersId(context.Context, *ChatId) (*ChatUsersId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatUsersId not implemented")
}
func (UnimplementedGroupServerServiceServer) GetChatInfos(context.Context, *ChatIdsReq) (*ChatInfos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatInfos not implemented")
}
func (UnimplementedGroupServerServiceServer) GetParticipants(context.Context, *ChatIdsReq) (*ChatParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipants not implemented")
}
func (UnimplementedGroupServerServiceServer) GetCommonChats(context.Context, *CommonChatReq) (*CommonChatsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommonChats not implemented")
}
func (UnimplementedGroupServerServiceServer) mustEmbedUnimplementedGroupServerServiceServer() {}

// UnsafeGroupServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServerServiceServer will
// result in compilation errors.
type UnsafeGroupServerServiceServer interface {
	mustEmbedUnimplementedGroupServerServiceServer()
}

func RegisterGroupServerServiceServer(s grpc.ServiceRegistrar, srv GroupServerServiceServer) {
	s.RegisterService(&GroupServerService_ServiceDesc, srv)
}

func _GroupServerService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).CreateChat(ctx, req.(*CreateRquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatInfo(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_UpdateChatTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatTitleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).UpdateChatTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/UpdateChatTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).UpdateChatTitle(ctx, req.(*ChatTitleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_UpdateChatPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatPhotoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).UpdateChatPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/UpdateChatPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).UpdateChatPhoto(ctx, req.(*ChatPhotoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_UpdateChatMigrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMigrateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).UpdateChatMigrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/UpdateChatMigrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).UpdateChatMigrate(ctx, req.(*ChatMigrateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_UpdateChatAdminsEn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatadminEnableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).UpdateChatAdminsEn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/UpdateChatAdminsEn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).UpdateChatAdminsEn(ctx, req.(*ChatadminEnableReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_UpdateChatLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatExportLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).UpdateChatLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/UpdateChatLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).UpdateChatLink(ctx, req.(*ChatExportLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_DelChatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).DelChatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/DelChatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).DelChatInfo(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_AddChatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).AddChatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/AddChatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).AddChatUser(ctx, req.(*ChatMemberBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_DelChatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).DelChatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/DelChatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).DelChatUser(ctx, req.(*ChatMemberBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_SetChatUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).SetChatUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/SetChatUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).SetChatUserAdmin(ctx, req.(*ChatMemberBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_DelChatUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).DelChatUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/DelChatUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).DelChatUserAdmin(ctx, req.(*ChatMemberBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatLittleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatLittleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatLittleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatLittleInfo(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatParticipants(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatUserInfo(ctx, req.(*ChatMemberBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatUsersId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatUsersId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatUsersId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatUsersId(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetChatInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetChatInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetChatInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetChatInfos(ctx, req.(*ChatIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetParticipants(ctx, req.(*ChatIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupServerService_GetCommonChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServerServiceServer).GetCommonChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.GroupServerService/GetCommonChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServerServiceServer).GetCommonChats(ctx, req.(*CommonChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupServerService_ServiceDesc is the grpc.ServiceDesc for GroupServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatserver.GroupServerService",
	HandlerType: (*GroupServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _GroupServerService_CreateChat_Handler,
		},
		{
			MethodName: "GetChatInfo",
			Handler:    _GroupServerService_GetChatInfo_Handler,
		},
		{
			MethodName: "UpdateChatTitle",
			Handler:    _GroupServerService_UpdateChatTitle_Handler,
		},
		{
			MethodName: "UpdateChatPhoto",
			Handler:    _GroupServerService_UpdateChatPhoto_Handler,
		},
		{
			MethodName: "UpdateChatMigrate",
			Handler:    _GroupServerService_UpdateChatMigrate_Handler,
		},
		{
			MethodName: "UpdateChatAdminsEn",
			Handler:    _GroupServerService_UpdateChatAdminsEn_Handler,
		},
		{
			MethodName: "UpdateChatLink",
			Handler:    _GroupServerService_UpdateChatLink_Handler,
		},
		{
			MethodName: "DelChatInfo",
			Handler:    _GroupServerService_DelChatInfo_Handler,
		},
		{
			MethodName: "AddChatUser",
			Handler:    _GroupServerService_AddChatUser_Handler,
		},
		{
			MethodName: "DelChatUser",
			Handler:    _GroupServerService_DelChatUser_Handler,
		},
		{
			MethodName: "SetChatUserAdmin",
			Handler:    _GroupServerService_SetChatUserAdmin_Handler,
		},
		{
			MethodName: "DelChatUserAdmin",
			Handler:    _GroupServerService_DelChatUserAdmin_Handler,
		},
		{
			MethodName: "GetChatLittleInfo",
			Handler:    _GroupServerService_GetChatLittleInfo_Handler,
		},
		{
			MethodName: "GetChatParticipants",
			Handler:    _GroupServerService_GetChatParticipants_Handler,
		},
		{
			MethodName: "GetChatUserInfo",
			Handler:    _GroupServerService_GetChatUserInfo_Handler,
		},
		{
			MethodName: "GetChatUsersId",
			Handler:    _GroupServerService_GetChatUsersId_Handler,
		},
		{
			MethodName: "GetChatInfos",
			Handler:    _GroupServerService_GetChatInfos_Handler,
		},
		{
			MethodName: "GetParticipants",
			Handler:    _GroupServerService_GetParticipants_Handler,
		},
		{
			MethodName: "GetCommonChats",
			Handler:    _GroupServerService_GetCommonChats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatservice.proto",
}