//*
// 定义提供给机器人逻辑服务使用的API
// 只能由机器人逻辑服务来调用

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: middleware.proto

package middleware

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 路由请求数据
type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetDc uint32 `protobuf:"varint,1,opt,name=target_dc,json=targetDc,proto3" json:"target_dc,omitempty"` // 目标DC
	Tldata   []byte `protobuf:"bytes,2,opt,name=tldata,proto3" json:"tldata,omitempty"`                      // 协议数据
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *RouteRequest) GetTargetDc() uint32 {
	if x != nil {
		return x.TargetDc
	}
	return 0
}

func (x *RouteRequest) GetTldata() []byte {
	if x != nil {
		return x.Tldata
	}
	return nil
}

// 路由响应数据
type RouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tldata []byte `protobuf:"bytes,1,opt,name=tldata,proto3" json:"tldata,omitempty"` // 协议数据
}

func (x *RouteResponse) Reset() {
	*x = RouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse) ProtoMessage() {}

func (x *RouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse.ProtoReflect.Descriptor instead.
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return file_middleware_proto_rawDescGZIP(), []int{1}
}

func (x *RouteResponse) GetTldata() []byte {
	if x != nil {
		return x.Tldata
	}
	return nil
}

var File_middleware_proto protoreflect.FileDescriptor

var file_middleware_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x22, 0x43,
	0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6c, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6c, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb4, 0x01, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x52, 0x0a, 0x19, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x19, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_middleware_proto_rawDescOnce sync.Once
	file_middleware_proto_rawDescData = file_middleware_proto_rawDesc
)

func file_middleware_proto_rawDescGZIP() []byte {
	file_middleware_proto_rawDescOnce.Do(func() {
		file_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_middleware_proto_rawDescData)
	})
	return file_middleware_proto_rawDescData
}

var file_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_middleware_proto_goTypes = []interface{}{
	(*RouteRequest)(nil),  // 0: middleware.RouteRequest
	(*RouteResponse)(nil), // 1: middleware.RouteResponse
}
var file_middleware_proto_depIdxs = []int32{
	0, // 0: middleware.Middleware.RouteGetBotCallbackAnswer:input_type -> middleware.RouteRequest
	0, // 1: middleware.Middleware.RouteSetBotCallbackAnswer:input_type -> middleware.RouteRequest
	1, // 2: middleware.Middleware.RouteGetBotCallbackAnswer:output_type -> middleware.RouteResponse
	1, // 3: middleware.Middleware.RouteSetBotCallbackAnswer:output_type -> middleware.RouteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_middleware_proto_init() }
func file_middleware_proto_init() {
	if File_middleware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_middleware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_middleware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_middleware_proto_goTypes,
		DependencyIndexes: file_middleware_proto_depIdxs,
		MessageInfos:      file_middleware_proto_msgTypes,
	}.Build()
	File_middleware_proto = out.File
	file_middleware_proto_rawDesc = nil
	file_middleware_proto_goTypes = nil
	file_middleware_proto_depIdxs = nil
}
