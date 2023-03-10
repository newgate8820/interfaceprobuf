// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: photoresize.proto

package photocropping

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

// PhotoResizeClient is the client API for PhotoResize service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoResizeClient interface {
	// Resize
	Resize(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeReply, error)
	// ResizeMedia
	ResizeMedia(ctx context.Context, in *ResizeMediaRequest, opts ...grpc.CallOption) (*ResizeMediaReply, error)
}

type photoResizeClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoResizeClient(cc grpc.ClientConnInterface) PhotoResizeClient {
	return &photoResizeClient{cc}
}

func (c *photoResizeClient) Resize(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeReply, error) {
	out := new(ResizeReply)
	err := c.cc.Invoke(ctx, "/photocropping.PhotoResize/Resize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoResizeClient) ResizeMedia(ctx context.Context, in *ResizeMediaRequest, opts ...grpc.CallOption) (*ResizeMediaReply, error) {
	out := new(ResizeMediaReply)
	err := c.cc.Invoke(ctx, "/photocropping.PhotoResize/ResizeMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoResizeServer is the server API for PhotoResize service.
// All implementations must embed UnimplementedPhotoResizeServer
// for forward compatibility
type PhotoResizeServer interface {
	// Resize
	Resize(context.Context, *ResizeRequest) (*ResizeReply, error)
	// ResizeMedia
	ResizeMedia(context.Context, *ResizeMediaRequest) (*ResizeMediaReply, error)
	mustEmbedUnimplementedPhotoResizeServer()
}

// UnimplementedPhotoResizeServer must be embedded to have forward compatible implementations.
type UnimplementedPhotoResizeServer struct {
}

func (UnimplementedPhotoResizeServer) Resize(context.Context, *ResizeRequest) (*ResizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resize not implemented")
}
func (UnimplementedPhotoResizeServer) ResizeMedia(context.Context, *ResizeMediaRequest) (*ResizeMediaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeMedia not implemented")
}
func (UnimplementedPhotoResizeServer) mustEmbedUnimplementedPhotoResizeServer() {}

// UnsafePhotoResizeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoResizeServer will
// result in compilation errors.
type UnsafePhotoResizeServer interface {
	mustEmbedUnimplementedPhotoResizeServer()
}

func RegisterPhotoResizeServer(s grpc.ServiceRegistrar, srv PhotoResizeServer) {
	s.RegisterService(&PhotoResize_ServiceDesc, srv)
}

func _PhotoResize_Resize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoResizeServer).Resize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photocropping.PhotoResize/Resize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoResizeServer).Resize(ctx, req.(*ResizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoResize_ResizeMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoResizeServer).ResizeMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photocropping.PhotoResize/ResizeMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoResizeServer).ResizeMedia(ctx, req.(*ResizeMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhotoResize_ServiceDesc is the grpc.ServiceDesc for PhotoResize service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoResize_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "photocropping.PhotoResize",
	HandlerType: (*PhotoResizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resize",
			Handler:    _PhotoResize_Resize_Handler,
		},
		{
			MethodName: "ResizeMedia",
			Handler:    _PhotoResize_ResizeMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photoresize.proto",
}
