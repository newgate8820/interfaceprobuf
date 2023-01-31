// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: imvedioappserver.proto

package imvedioappserver

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

// VideoServerClient is the client API for VideoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServerClient interface {
	// 创建房间
	CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomReply, error)
	// 获取进入房间的Token
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenReply, error)
	// 批量获取Token
	BatchGetToken(ctx context.Context, in *BatchGetTokenReq, opts ...grpc.CallOption) (*BatchGetTokenReply, error)
	// 校验Token
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenReply, error)
	// 通过room_id获取房间里的用户列表
	GetUsersByRoomId(ctx context.Context, in *GetUsersByRoomIdReq, opts ...grpc.CallOption) (*GetUsersByRoomIdReply, error)
	// 获取房间创建者信息
	GetRoomAdminUser(ctx context.Context, in *GetRoomAdminUserReq, opts ...grpc.CallOption) (*GetRoomAdminUserReply, error)
	// 保存参与者id
	SaveParticipantKeyId(ctx context.Context, in *SaveParticipantKeyIdReq, opts ...grpc.CallOption) (*SaveParticipantKeyIdReply, error)
	GetParticipantKeyId(ctx context.Context, in *GetParticipantKeyIdReq, opts ...grpc.CallOption) (*GetParticipantKeyIdReply, error)
	GetUserMeeting(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*BoolResult, error)
	LetUserOffLine(ctx context.Context, in *UsersIdReq, opts ...grpc.CallOption) (*BoolResult, error)
	KeepLive(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BoolResult, error)
	ChangeModel(ctx context.Context, in *ChgModelReq, opts ...grpc.CallOption) (*BoolResult, error)
	GetModel(ctx context.Context, in *ModelReq, opts ...grpc.CallOption) (*ModelRsp, error)
}

type videoServerClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServerClient(cc grpc.ClientConnInterface) VideoServerClient {
	return &videoServerClient{cc}
}

func (c *videoServerClient) CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomReply, error) {
	out := new(CreateRoomReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenReply, error) {
	out := new(GetTokenReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) BatchGetToken(ctx context.Context, in *BatchGetTokenReq, opts ...grpc.CallOption) (*BatchGetTokenReply, error) {
	out := new(BatchGetTokenReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/BatchGetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenReply, error) {
	out := new(VerifyTokenReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetUsersByRoomId(ctx context.Context, in *GetUsersByRoomIdReq, opts ...grpc.CallOption) (*GetUsersByRoomIdReply, error) {
	out := new(GetUsersByRoomIdReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetUsersByRoomId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetRoomAdminUser(ctx context.Context, in *GetRoomAdminUserReq, opts ...grpc.CallOption) (*GetRoomAdminUserReply, error) {
	out := new(GetRoomAdminUserReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetRoomAdminUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) SaveParticipantKeyId(ctx context.Context, in *SaveParticipantKeyIdReq, opts ...grpc.CallOption) (*SaveParticipantKeyIdReply, error) {
	out := new(SaveParticipantKeyIdReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/SaveParticipantKeyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetParticipantKeyId(ctx context.Context, in *GetParticipantKeyIdReq, opts ...grpc.CallOption) (*GetParticipantKeyIdReply, error) {
	out := new(GetParticipantKeyIdReply)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetParticipantKeyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetUserMeeting(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetUserMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) LetUserOffLine(ctx context.Context, in *UsersIdReq, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/LetUserOffLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) KeepLive(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/KeepLive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) ChangeModel(ctx context.Context, in *ChgModelReq, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/ChangeModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServerClient) GetModel(ctx context.Context, in *ModelReq, opts ...grpc.CallOption) (*ModelRsp, error) {
	out := new(ModelRsp)
	err := c.cc.Invoke(ctx, "/imvedioappserver.VideoServer/GetModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServerServer is the server API for VideoServer service.
// All implementations must embed UnimplementedVideoServerServer
// for forward compatibility
type VideoServerServer interface {
	// 创建房间
	CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomReply, error)
	// 获取进入房间的Token
	GetToken(context.Context, *GetTokenReq) (*GetTokenReply, error)
	// 批量获取Token
	BatchGetToken(context.Context, *BatchGetTokenReq) (*BatchGetTokenReply, error)
	// 校验Token
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenReply, error)
	// 通过room_id获取房间里的用户列表
	GetUsersByRoomId(context.Context, *GetUsersByRoomIdReq) (*GetUsersByRoomIdReply, error)
	// 获取房间创建者信息
	GetRoomAdminUser(context.Context, *GetRoomAdminUserReq) (*GetRoomAdminUserReply, error)
	// 保存参与者id
	SaveParticipantKeyId(context.Context, *SaveParticipantKeyIdReq) (*SaveParticipantKeyIdReply, error)
	GetParticipantKeyId(context.Context, *GetParticipantKeyIdReq) (*GetParticipantKeyIdReply, error)
	GetUserMeeting(context.Context, *UserIdReq) (*BoolResult, error)
	LetUserOffLine(context.Context, *UsersIdReq) (*BoolResult, error)
	KeepLive(context.Context, *IdReq) (*BoolResult, error)
	ChangeModel(context.Context, *ChgModelReq) (*BoolResult, error)
	GetModel(context.Context, *ModelReq) (*ModelRsp, error)
	mustEmbedUnimplementedVideoServerServer()
}

// UnimplementedVideoServerServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServerServer struct {
}

func (UnimplementedVideoServerServer) CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedVideoServerServer) GetToken(context.Context, *GetTokenReq) (*GetTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedVideoServerServer) BatchGetToken(context.Context, *BatchGetTokenReq) (*BatchGetTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetToken not implemented")
}
func (UnimplementedVideoServerServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedVideoServerServer) GetUsersByRoomId(context.Context, *GetUsersByRoomIdReq) (*GetUsersByRoomIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByRoomId not implemented")
}
func (UnimplementedVideoServerServer) GetRoomAdminUser(context.Context, *GetRoomAdminUserReq) (*GetRoomAdminUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomAdminUser not implemented")
}
func (UnimplementedVideoServerServer) SaveParticipantKeyId(context.Context, *SaveParticipantKeyIdReq) (*SaveParticipantKeyIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveParticipantKeyId not implemented")
}
func (UnimplementedVideoServerServer) GetParticipantKeyId(context.Context, *GetParticipantKeyIdReq) (*GetParticipantKeyIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipantKeyId not implemented")
}
func (UnimplementedVideoServerServer) GetUserMeeting(context.Context, *UserIdReq) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMeeting not implemented")
}
func (UnimplementedVideoServerServer) LetUserOffLine(context.Context, *UsersIdReq) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LetUserOffLine not implemented")
}
func (UnimplementedVideoServerServer) KeepLive(context.Context, *IdReq) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepLive not implemented")
}
func (UnimplementedVideoServerServer) ChangeModel(context.Context, *ChgModelReq) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeModel not implemented")
}
func (UnimplementedVideoServerServer) GetModel(context.Context, *ModelReq) (*ModelRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModel not implemented")
}
func (UnimplementedVideoServerServer) mustEmbedUnimplementedVideoServerServer() {}

// UnsafeVideoServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServerServer will
// result in compilation errors.
type UnsafeVideoServerServer interface {
	mustEmbedUnimplementedVideoServerServer()
}

func RegisterVideoServerServer(s grpc.ServiceRegistrar, srv VideoServerServer) {
	s.RegisterService(&VideoServer_ServiceDesc, srv)
}

func _VideoServer_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).CreateRoom(ctx, req.(*CreateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_BatchGetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).BatchGetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/BatchGetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).BatchGetToken(ctx, req.(*BatchGetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).VerifyToken(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetUsersByRoomId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByRoomIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetUsersByRoomId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetUsersByRoomId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetUsersByRoomId(ctx, req.(*GetUsersByRoomIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetRoomAdminUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomAdminUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetRoomAdminUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetRoomAdminUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetRoomAdminUser(ctx, req.(*GetRoomAdminUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_SaveParticipantKeyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveParticipantKeyIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).SaveParticipantKeyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/SaveParticipantKeyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).SaveParticipantKeyId(ctx, req.(*SaveParticipantKeyIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetParticipantKeyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParticipantKeyIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetParticipantKeyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetParticipantKeyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetParticipantKeyId(ctx, req.(*GetParticipantKeyIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetUserMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetUserMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetUserMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetUserMeeting(ctx, req.(*UserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_LetUserOffLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).LetUserOffLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/LetUserOffLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).LetUserOffLine(ctx, req.(*UsersIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_KeepLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).KeepLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/KeepLive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).KeepLive(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_ChangeModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChgModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).ChangeModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/ChangeModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).ChangeModel(ctx, req.(*ChgModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoServer_GetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServerServer).GetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imvedioappserver.VideoServer/GetModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServerServer).GetModel(ctx, req.(*ModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoServer_ServiceDesc is the grpc.ServiceDesc for VideoServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imvedioappserver.VideoServer",
	HandlerType: (*VideoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _VideoServer_CreateRoom_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _VideoServer_GetToken_Handler,
		},
		{
			MethodName: "BatchGetToken",
			Handler:    _VideoServer_BatchGetToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _VideoServer_VerifyToken_Handler,
		},
		{
			MethodName: "GetUsersByRoomId",
			Handler:    _VideoServer_GetUsersByRoomId_Handler,
		},
		{
			MethodName: "GetRoomAdminUser",
			Handler:    _VideoServer_GetRoomAdminUser_Handler,
		},
		{
			MethodName: "SaveParticipantKeyId",
			Handler:    _VideoServer_SaveParticipantKeyId_Handler,
		},
		{
			MethodName: "GetParticipantKeyId",
			Handler:    _VideoServer_GetParticipantKeyId_Handler,
		},
		{
			MethodName: "GetUserMeeting",
			Handler:    _VideoServer_GetUserMeeting_Handler,
		},
		{
			MethodName: "LetUserOffLine",
			Handler:    _VideoServer_LetUserOffLine_Handler,
		},
		{
			MethodName: "KeepLive",
			Handler:    _VideoServer_KeepLive_Handler,
		},
		{
			MethodName: "ChangeModel",
			Handler:    _VideoServer_ChangeModel_Handler,
		},
		{
			MethodName: "GetModel",
			Handler:    _VideoServer_GetModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imvedioappserver.proto",
}