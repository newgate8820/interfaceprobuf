// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: immediamonitorserver.proto

package immediamonitorserver

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

// MediaMonitorServiceClient is the client API for MediaMonitorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaMonitorServiceClient interface {
	// MediaMonitorInfoSave 媒体监督所需相关信息存储
	MediaMonitorInfoSave(ctx context.Context, in *MediaMonitorInfoSaveReq, opts ...grpc.CallOption) (*MediaMonitorInfoSaveResp, error)
}

type mediaMonitorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaMonitorServiceClient(cc grpc.ClientConnInterface) MediaMonitorServiceClient {
	return &mediaMonitorServiceClient{cc}
}

func (c *mediaMonitorServiceClient) MediaMonitorInfoSave(ctx context.Context, in *MediaMonitorInfoSaveReq, opts ...grpc.CallOption) (*MediaMonitorInfoSaveResp, error) {
	out := new(MediaMonitorInfoSaveResp)
	err := c.cc.Invoke(ctx, "/immediamonitorserver.MediaMonitorService/MediaMonitorInfoSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaMonitorServiceServer is the server API for MediaMonitorService service.
// All implementations must embed UnimplementedMediaMonitorServiceServer
// for forward compatibility
type MediaMonitorServiceServer interface {
	// MediaMonitorInfoSave 媒体监督所需相关信息存储
	MediaMonitorInfoSave(context.Context, *MediaMonitorInfoSaveReq) (*MediaMonitorInfoSaveResp, error)
	mustEmbedUnimplementedMediaMonitorServiceServer()
}

// UnimplementedMediaMonitorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaMonitorServiceServer struct {
}

func (UnimplementedMediaMonitorServiceServer) MediaMonitorInfoSave(context.Context, *MediaMonitorInfoSaveReq) (*MediaMonitorInfoSaveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaMonitorInfoSave not implemented")
}
func (UnimplementedMediaMonitorServiceServer) mustEmbedUnimplementedMediaMonitorServiceServer() {}

// UnsafeMediaMonitorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaMonitorServiceServer will
// result in compilation errors.
type UnsafeMediaMonitorServiceServer interface {
	mustEmbedUnimplementedMediaMonitorServiceServer()
}

func RegisterMediaMonitorServiceServer(s grpc.ServiceRegistrar, srv MediaMonitorServiceServer) {
	s.RegisterService(&MediaMonitorService_ServiceDesc, srv)
}

func _MediaMonitorService_MediaMonitorInfoSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaMonitorInfoSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaMonitorServiceServer).MediaMonitorInfoSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immediamonitorserver.MediaMonitorService/MediaMonitorInfoSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaMonitorServiceServer).MediaMonitorInfoSave(ctx, req.(*MediaMonitorInfoSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaMonitorService_ServiceDesc is the grpc.ServiceDesc for MediaMonitorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaMonitorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immediamonitorserver.MediaMonitorService",
	HandlerType: (*MediaMonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MediaMonitorInfoSave",
			Handler:    _MediaMonitorService_MediaMonitorInfoSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "immediamonitorserver.proto",
}