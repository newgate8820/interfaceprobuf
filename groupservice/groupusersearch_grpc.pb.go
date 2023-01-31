// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: groupusersearch.proto

package groupservice

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

// GroupUserSearchClient is the client API for GroupUserSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupUserSearchClient interface {
	// 根据传入ID搜索用户
	SearchUsersByName(ctx context.Context, in *SearchUsersByNameReq, opts ...grpc.CallOption) (*SearchUsersByNameReply, error)
}

type groupUserSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupUserSearchClient(cc grpc.ClientConnInterface) GroupUserSearchClient {
	return &groupUserSearchClient{cc}
}

func (c *groupUserSearchClient) SearchUsersByName(ctx context.Context, in *SearchUsersByNameReq, opts ...grpc.CallOption) (*SearchUsersByNameReply, error) {
	out := new(SearchUsersByNameReply)
	err := c.cc.Invoke(ctx, "/groupservice.GroupUserSearch/SearchUsersByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupUserSearchServer is the server API for GroupUserSearch service.
// All implementations must embed UnimplementedGroupUserSearchServer
// for forward compatibility
type GroupUserSearchServer interface {
	// 根据传入ID搜索用户
	SearchUsersByName(context.Context, *SearchUsersByNameReq) (*SearchUsersByNameReply, error)
	mustEmbedUnimplementedGroupUserSearchServer()
}

// UnimplementedGroupUserSearchServer must be embedded to have forward compatible implementations.
type UnimplementedGroupUserSearchServer struct {
}

func (UnimplementedGroupUserSearchServer) SearchUsersByName(context.Context, *SearchUsersByNameReq) (*SearchUsersByNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsersByName not implemented")
}
func (UnimplementedGroupUserSearchServer) mustEmbedUnimplementedGroupUserSearchServer() {}

// UnsafeGroupUserSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupUserSearchServer will
// result in compilation errors.
type UnsafeGroupUserSearchServer interface {
	mustEmbedUnimplementedGroupUserSearchServer()
}

func RegisterGroupUserSearchServer(s grpc.ServiceRegistrar, srv GroupUserSearchServer) {
	s.RegisterService(&GroupUserSearch_ServiceDesc, srv)
}

func _GroupUserSearch_SearchUsersByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupUserSearchServer).SearchUsersByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/groupservice.GroupUserSearch/SearchUsersByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupUserSearchServer).SearchUsersByName(ctx, req.(*SearchUsersByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupUserSearch_ServiceDesc is the grpc.ServiceDesc for GroupUserSearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupUserSearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "groupservice.GroupUserSearch",
	HandlerType: (*GroupUserSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUsersByName",
			Handler:    _GroupUserSearch_SearchUsersByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "groupusersearch.proto",
}
