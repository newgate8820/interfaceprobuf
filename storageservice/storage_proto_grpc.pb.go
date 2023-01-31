// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: storage_proto.proto

package storageservice

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

// FileStorageServiceClient is the client API for FileStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStorageServiceClient interface {
	// ServiceLimit 服务上传下载限制
	ServiceLimit(ctx context.Context, in *MsgServiceLimitReq, opts ...grpc.CallOption) (*MsgServiceLimitResp, error)
	// UploadFile 上传文件或文件片
	UploadFile(ctx context.Context, in *MsgUploadFileReq, opts ...grpc.CallOption) (*MsgUploadFileResp, error)
	// DownloadFile 下载文件或文件片
	DownloadFile(ctx context.Context, in *MsgDownloadFileReq, opts ...grpc.CallOption) (*MsgDownloadFileResp, error)
	// DownloadFile 下载大块文件到缓存
	DownloadFileCache(ctx context.Context, in *MsgDownloadFileReq, opts ...grpc.CallOption) (*MsgDownloadFileResp, error)
	// QueryObjInfo 查询对象信息
	QueryObjInfo(ctx context.Context, in *MsgQueryObjInfoReq, opts ...grpc.CallOption) (*MsgQueryObjInfoResp, error)
	// 网关上传文件
	GatewayUploadFile(ctx context.Context, in *MsgFilePartsReq, opts ...grpc.CallOption) (*MsgFilePartsResp, error)
	// 网关取上传文件的Fid, 与DC约定：1.收到此消息认为文件片已经传完；2.keyid+fileid与存储fid映射关系丢弃
	GatewayGetUploadFileFid(ctx context.Context, in *MsgGetFidReq, opts ...grpc.CallOption) (*MsgGetFidResp, error)
	// 网关上传文件描述信息
	GatewayUploadFileDesc(ctx context.Context, in *MsgFileDescReq, opts ...grpc.CallOption) (*MsgFileDescResp, error)
	// 网关获取文件描述信息
	GatewayGetFileDesc(ctx context.Context, in *MsgGetFileDescReq, opts ...grpc.CallOption) (*MsgGetFileDescResp, error)
}

type fileStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStorageServiceClient(cc grpc.ClientConnInterface) FileStorageServiceClient {
	return &fileStorageServiceClient{cc}
}

func (c *fileStorageServiceClient) ServiceLimit(ctx context.Context, in *MsgServiceLimitReq, opts ...grpc.CallOption) (*MsgServiceLimitResp, error) {
	out := new(MsgServiceLimitResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/ServiceLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) UploadFile(ctx context.Context, in *MsgUploadFileReq, opts ...grpc.CallOption) (*MsgUploadFileResp, error) {
	out := new(MsgUploadFileResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) DownloadFile(ctx context.Context, in *MsgDownloadFileReq, opts ...grpc.CallOption) (*MsgDownloadFileResp, error) {
	out := new(MsgDownloadFileResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) DownloadFileCache(ctx context.Context, in *MsgDownloadFileReq, opts ...grpc.CallOption) (*MsgDownloadFileResp, error) {
	out := new(MsgDownloadFileResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/DownloadFileCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) QueryObjInfo(ctx context.Context, in *MsgQueryObjInfoReq, opts ...grpc.CallOption) (*MsgQueryObjInfoResp, error) {
	out := new(MsgQueryObjInfoResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/QueryObjInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) GatewayUploadFile(ctx context.Context, in *MsgFilePartsReq, opts ...grpc.CallOption) (*MsgFilePartsResp, error) {
	out := new(MsgFilePartsResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/GatewayUploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) GatewayGetUploadFileFid(ctx context.Context, in *MsgGetFidReq, opts ...grpc.CallOption) (*MsgGetFidResp, error) {
	out := new(MsgGetFidResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/GatewayGetUploadFileFid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) GatewayUploadFileDesc(ctx context.Context, in *MsgFileDescReq, opts ...grpc.CallOption) (*MsgFileDescResp, error) {
	out := new(MsgFileDescResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/GatewayUploadFileDesc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) GatewayGetFileDesc(ctx context.Context, in *MsgGetFileDescReq, opts ...grpc.CallOption) (*MsgGetFileDescResp, error) {
	out := new(MsgGetFileDescResp)
	err := c.cc.Invoke(ctx, "/storageservice.FileStorageService/GatewayGetFileDesc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileStorageServiceServer is the server API for FileStorageService service.
// All implementations must embed UnimplementedFileStorageServiceServer
// for forward compatibility
type FileStorageServiceServer interface {
	// ServiceLimit 服务上传下载限制
	ServiceLimit(context.Context, *MsgServiceLimitReq) (*MsgServiceLimitResp, error)
	// UploadFile 上传文件或文件片
	UploadFile(context.Context, *MsgUploadFileReq) (*MsgUploadFileResp, error)
	// DownloadFile 下载文件或文件片
	DownloadFile(context.Context, *MsgDownloadFileReq) (*MsgDownloadFileResp, error)
	// DownloadFile 下载大块文件到缓存
	DownloadFileCache(context.Context, *MsgDownloadFileReq) (*MsgDownloadFileResp, error)
	// QueryObjInfo 查询对象信息
	QueryObjInfo(context.Context, *MsgQueryObjInfoReq) (*MsgQueryObjInfoResp, error)
	// 网关上传文件
	GatewayUploadFile(context.Context, *MsgFilePartsReq) (*MsgFilePartsResp, error)
	// 网关取上传文件的Fid, 与DC约定：1.收到此消息认为文件片已经传完；2.keyid+fileid与存储fid映射关系丢弃
	GatewayGetUploadFileFid(context.Context, *MsgGetFidReq) (*MsgGetFidResp, error)
	// 网关上传文件描述信息
	GatewayUploadFileDesc(context.Context, *MsgFileDescReq) (*MsgFileDescResp, error)
	// 网关获取文件描述信息
	GatewayGetFileDesc(context.Context, *MsgGetFileDescReq) (*MsgGetFileDescResp, error)
	mustEmbedUnimplementedFileStorageServiceServer()
}

// UnimplementedFileStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileStorageServiceServer struct {
}

func (UnimplementedFileStorageServiceServer) ServiceLimit(context.Context, *MsgServiceLimitReq) (*MsgServiceLimitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceLimit not implemented")
}
func (UnimplementedFileStorageServiceServer) UploadFile(context.Context, *MsgUploadFileReq) (*MsgUploadFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileStorageServiceServer) DownloadFile(context.Context, *MsgDownloadFileReq) (*MsgDownloadFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileStorageServiceServer) DownloadFileCache(context.Context, *MsgDownloadFileReq) (*MsgDownloadFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFileCache not implemented")
}
func (UnimplementedFileStorageServiceServer) QueryObjInfo(context.Context, *MsgQueryObjInfoReq) (*MsgQueryObjInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryObjInfo not implemented")
}
func (UnimplementedFileStorageServiceServer) GatewayUploadFile(context.Context, *MsgFilePartsReq) (*MsgFilePartsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayUploadFile not implemented")
}
func (UnimplementedFileStorageServiceServer) GatewayGetUploadFileFid(context.Context, *MsgGetFidReq) (*MsgGetFidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayGetUploadFileFid not implemented")
}
func (UnimplementedFileStorageServiceServer) GatewayUploadFileDesc(context.Context, *MsgFileDescReq) (*MsgFileDescResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayUploadFileDesc not implemented")
}
func (UnimplementedFileStorageServiceServer) GatewayGetFileDesc(context.Context, *MsgGetFileDescReq) (*MsgGetFileDescResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayGetFileDesc not implemented")
}
func (UnimplementedFileStorageServiceServer) mustEmbedUnimplementedFileStorageServiceServer() {}

// UnsafeFileStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStorageServiceServer will
// result in compilation errors.
type UnsafeFileStorageServiceServer interface {
	mustEmbedUnimplementedFileStorageServiceServer()
}

func RegisterFileStorageServiceServer(s grpc.ServiceRegistrar, srv FileStorageServiceServer) {
	s.RegisterService(&FileStorageService_ServiceDesc, srv)
}

func _FileStorageService_ServiceLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgServiceLimitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).ServiceLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/ServiceLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).ServiceLimit(ctx, req.(*MsgServiceLimitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUploadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).UploadFile(ctx, req.(*MsgUploadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDownloadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).DownloadFile(ctx, req.(*MsgDownloadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_DownloadFileCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDownloadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).DownloadFileCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/DownloadFileCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).DownloadFileCache(ctx, req.(*MsgDownloadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_QueryObjInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryObjInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).QueryObjInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/QueryObjInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).QueryObjInfo(ctx, req.(*MsgQueryObjInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_GatewayUploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFilePartsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).GatewayUploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/GatewayUploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).GatewayUploadFile(ctx, req.(*MsgFilePartsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_GatewayGetUploadFileFid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetFidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).GatewayGetUploadFileFid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/GatewayGetUploadFileFid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).GatewayGetUploadFileFid(ctx, req.(*MsgGetFidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_GatewayUploadFileDesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFileDescReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).GatewayUploadFileDesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/GatewayUploadFileDesc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).GatewayUploadFileDesc(ctx, req.(*MsgFileDescReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_GatewayGetFileDesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetFileDescReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).GatewayGetFileDesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageservice.FileStorageService/GatewayGetFileDesc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).GatewayGetFileDesc(ctx, req.(*MsgGetFileDescReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FileStorageService_ServiceDesc is the grpc.ServiceDesc for FileStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storageservice.FileStorageService",
	HandlerType: (*FileStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceLimit",
			Handler:    _FileStorageService_ServiceLimit_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _FileStorageService_UploadFile_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _FileStorageService_DownloadFile_Handler,
		},
		{
			MethodName: "DownloadFileCache",
			Handler:    _FileStorageService_DownloadFileCache_Handler,
		},
		{
			MethodName: "QueryObjInfo",
			Handler:    _FileStorageService_QueryObjInfo_Handler,
		},
		{
			MethodName: "GatewayUploadFile",
			Handler:    _FileStorageService_GatewayUploadFile_Handler,
		},
		{
			MethodName: "GatewayGetUploadFileFid",
			Handler:    _FileStorageService_GatewayGetUploadFileFid_Handler,
		},
		{
			MethodName: "GatewayUploadFileDesc",
			Handler:    _FileStorageService_GatewayUploadFileDesc_Handler,
		},
		{
			MethodName: "GatewayGetFileDesc",
			Handler:    _FileStorageService_GatewayGetFileDesc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage_proto.proto",
}
