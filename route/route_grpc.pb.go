// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: route.proto

package route

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

// RouteTableClient is the client API for RouteTable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteTableClient interface {
	Del(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*DelUserReply, error)
	Find(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserReply, error)
	FindEx(ctx context.Context, in *FindUserExRequest, opts ...grpc.CallOption) (*FindUserExReply, error)
	// rpc UpdateEx(UpdateUserExRequest) returns (UpdateUserExReply);
	FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersReply, error)
	UpdateExV2(ctx context.Context, in *UpdateUserExV2Request, opts ...grpc.CallOption) (*UpdateUserExV2Reply, error)
	UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error)
}

type routeTableClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteTableClient(cc grpc.ClientConnInterface) RouteTableClient {
	return &routeTableClient{cc}
}

func (c *routeTableClient) Del(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*DelUserReply, error) {
	out := new(DelUserReply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableClient) Find(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserReply, error) {
	out := new(FindUserReply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableClient) FindEx(ctx context.Context, in *FindUserExRequest, opts ...grpc.CallOption) (*FindUserExReply, error) {
	out := new(FindUserExReply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/FindEx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableClient) FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersReply, error) {
	out := new(FindUsersReply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/FindUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableClient) UpdateExV2(ctx context.Context, in *UpdateUserExV2Request, opts ...grpc.CallOption) (*UpdateUserExV2Reply, error) {
	out := new(UpdateUserExV2Reply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/UpdateExV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableClient) UpdateUsers(ctx context.Context, in *UpdateUsersRequest, opts ...grpc.CallOption) (*UpdateUsersReply, error) {
	out := new(UpdateUsersReply)
	err := c.cc.Invoke(ctx, "/route.RouteTable/UpdateUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteTableServer is the server API for RouteTable service.
// All implementations must embed UnimplementedRouteTableServer
// for forward compatibility
type RouteTableServer interface {
	Del(context.Context, *DelUserRequest) (*DelUserReply, error)
	Find(context.Context, *FindUserRequest) (*FindUserReply, error)
	FindEx(context.Context, *FindUserExRequest) (*FindUserExReply, error)
	// rpc UpdateEx(UpdateUserExRequest) returns (UpdateUserExReply);
	FindUsers(context.Context, *FindUsersRequest) (*FindUsersReply, error)
	UpdateExV2(context.Context, *UpdateUserExV2Request) (*UpdateUserExV2Reply, error)
	UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error)
	mustEmbedUnimplementedRouteTableServer()
}

// UnimplementedRouteTableServer must be embedded to have forward compatible implementations.
type UnimplementedRouteTableServer struct {
}

func (UnimplementedRouteTableServer) Del(context.Context, *DelUserRequest) (*DelUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedRouteTableServer) Find(context.Context, *FindUserRequest) (*FindUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedRouteTableServer) FindEx(context.Context, *FindUserExRequest) (*FindUserExReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEx not implemented")
}
func (UnimplementedRouteTableServer) FindUsers(context.Context, *FindUsersRequest) (*FindUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}
func (UnimplementedRouteTableServer) UpdateExV2(context.Context, *UpdateUserExV2Request) (*UpdateUserExV2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExV2 not implemented")
}
func (UnimplementedRouteTableServer) UpdateUsers(context.Context, *UpdateUsersRequest) (*UpdateUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsers not implemented")
}
func (UnimplementedRouteTableServer) mustEmbedUnimplementedRouteTableServer() {}

// UnsafeRouteTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteTableServer will
// result in compilation errors.
type UnsafeRouteTableServer interface {
	mustEmbedUnimplementedRouteTableServer()
}

func RegisterRouteTableServer(s grpc.ServiceRegistrar, srv RouteTableServer) {
	s.RegisterService(&RouteTable_ServiceDesc, srv)
}

func _RouteTable_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).Del(ctx, req.(*DelUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTable_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).Find(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTable_FindEx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserExRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).FindEx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/FindEx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).FindEx(ctx, req.(*FindUserExRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTable_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/FindUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).FindUsers(ctx, req.(*FindUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTable_UpdateExV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserExV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).UpdateExV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/UpdateExV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).UpdateExV2(ctx, req.(*UpdateUserExV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTable_UpdateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServer).UpdateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteTable/UpdateUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServer).UpdateUsers(ctx, req.(*UpdateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteTable_ServiceDesc is the grpc.ServiceDesc for RouteTable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteTable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.RouteTable",
	HandlerType: (*RouteTableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Del",
			Handler:    _RouteTable_Del_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _RouteTable_Find_Handler,
		},
		{
			MethodName: "FindEx",
			Handler:    _RouteTable_FindEx_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _RouteTable_FindUsers_Handler,
		},
		{
			MethodName: "UpdateExV2",
			Handler:    _RouteTable_UpdateExV2_Handler,
		},
		{
			MethodName: "UpdateUsers",
			Handler:    _RouteTable_UpdateUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}
