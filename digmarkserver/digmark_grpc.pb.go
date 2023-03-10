// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: digmark.proto

package digmarkserver

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

// DigMarkServiceClient is the client API for DigMarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DigMarkServiceClient interface {
	// 公共接口
	DigMarkCommon(ctx context.Context, in *DigMarkCommonReq, opts ...grpc.CallOption) (*DigMarkCommonResp, error)
}

type digMarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDigMarkServiceClient(cc grpc.ClientConnInterface) DigMarkServiceClient {
	return &digMarkServiceClient{cc}
}

func (c *digMarkServiceClient) DigMarkCommon(ctx context.Context, in *DigMarkCommonReq, opts ...grpc.CallOption) (*DigMarkCommonResp, error) {
	out := new(DigMarkCommonResp)
	err := c.cc.Invoke(ctx, "/digmarkserver.DigMarkService/DigMarkCommon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigMarkServiceServer is the server API for DigMarkService service.
// All implementations must embed UnimplementedDigMarkServiceServer
// for forward compatibility
type DigMarkServiceServer interface {
	// 公共接口
	DigMarkCommon(context.Context, *DigMarkCommonReq) (*DigMarkCommonResp, error)
	mustEmbedUnimplementedDigMarkServiceServer()
}

// UnimplementedDigMarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDigMarkServiceServer struct {
}

func (UnimplementedDigMarkServiceServer) DigMarkCommon(context.Context, *DigMarkCommonReq) (*DigMarkCommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DigMarkCommon not implemented")
}
func (UnimplementedDigMarkServiceServer) mustEmbedUnimplementedDigMarkServiceServer() {}

// UnsafeDigMarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DigMarkServiceServer will
// result in compilation errors.
type UnsafeDigMarkServiceServer interface {
	mustEmbedUnimplementedDigMarkServiceServer()
}

func RegisterDigMarkServiceServer(s grpc.ServiceRegistrar, srv DigMarkServiceServer) {
	s.RegisterService(&DigMarkService_ServiceDesc, srv)
}

func _DigMarkService_DigMarkCommon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigMarkCommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigMarkServiceServer).DigMarkCommon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/digmarkserver.DigMarkService/DigMarkCommon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigMarkServiceServer).DigMarkCommon(ctx, req.(*DigMarkCommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DigMarkService_ServiceDesc is the grpc.ServiceDesc for DigMarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DigMarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digmarkserver.DigMarkService",
	HandlerType: (*DigMarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DigMarkCommon",
			Handler:    _DigMarkService_DigMarkCommon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "digmark.proto",
}
