// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imshortvideoserver.proto

package imshortvideoserver

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

type ShortVideoBussinessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                //用户ID
	KeyId       int64  `protobuf:"varint,2,opt,name=Key_id,json=KeyId,proto3" json:"Key_id,omitempty"`                   //
	ClientIp    string `protobuf:"bytes,3,opt,name=Client_ip,json=ClientIp,proto3" json:"Client_ip,omitempty"`           //客户端IP
	BussinessId int32  `protobuf:"varint,4,opt,name=Bussiness_id,json=BussinessId,proto3" json:"Bussiness_id,omitempty"` // 業務Id
	ReqData     string `protobuf:"bytes,5,opt,name=ReqData,proto3" json:"ReqData,omitempty"`                             // 請求 的json
}

func (x *ShortVideoBussinessReq) Reset() {
	*x = ShortVideoBussinessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imshortvideoserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortVideoBussinessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortVideoBussinessReq) ProtoMessage() {}

func (x *ShortVideoBussinessReq) ProtoReflect() protoreflect.Message {
	mi := &file_imshortvideoserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortVideoBussinessReq.ProtoReflect.Descriptor instead.
func (*ShortVideoBussinessReq) Descriptor() ([]byte, []int) {
	return file_imshortvideoserver_proto_rawDescGZIP(), []int{0}
}

func (x *ShortVideoBussinessReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShortVideoBussinessReq) GetKeyId() int64 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *ShortVideoBussinessReq) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *ShortVideoBussinessReq) GetBussinessId() int32 {
	if x != nil {
		return x.BussinessId
	}
	return 0
}

func (x *ShortVideoBussinessReq) GetReqData() string {
	if x != nil {
		return x.ReqData
	}
	return ""
}

type ShortVideoBussinessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResData string  `protobuf:"bytes,1,opt,name=ResData,proto3" json:"ResData,omitempty"`         // 返回的 的json
	UserIds []int32 `protobuf:"varint,2,rep,packed,name=userIds,proto3" json:"userIds,omitempty"` // 需要Im 组装的用户信息的ID 数组
}

func (x *ShortVideoBussinessResp) Reset() {
	*x = ShortVideoBussinessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imshortvideoserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortVideoBussinessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortVideoBussinessResp) ProtoMessage() {}

func (x *ShortVideoBussinessResp) ProtoReflect() protoreflect.Message {
	mi := &file_imshortvideoserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortVideoBussinessResp.ProtoReflect.Descriptor instead.
func (*ShortVideoBussinessResp) Descriptor() ([]byte, []int) {
	return file_imshortvideoserver_proto_rawDescGZIP(), []int{1}
}

func (x *ShortVideoBussinessResp) GetResData() string {
	if x != nil {
		return x.ResData
	}
	return ""
}

func (x *ShortVideoBussinessResp) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_imshortvideoserver_proto protoreflect.FileDescriptor

var file_imshortvideoserver_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69, 0x6d, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xa2,
	0x01, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x75, 0x73,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x42, 0x75,
	0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x71, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x32, 0x81, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x2e, 0x69, 0x6d,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x75, 0x73, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x69, 0x6d, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imshortvideoserver_proto_rawDescOnce sync.Once
	file_imshortvideoserver_proto_rawDescData = file_imshortvideoserver_proto_rawDesc
)

func file_imshortvideoserver_proto_rawDescGZIP() []byte {
	file_imshortvideoserver_proto_rawDescOnce.Do(func() {
		file_imshortvideoserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_imshortvideoserver_proto_rawDescData)
	})
	return file_imshortvideoserver_proto_rawDescData
}

var file_imshortvideoserver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_imshortvideoserver_proto_goTypes = []interface{}{
	(*ShortVideoBussinessReq)(nil),  // 0: imshortvideoserver.ShortVideoBussinessReq
	(*ShortVideoBussinessResp)(nil), // 1: imshortvideoserver.ShortVideoBussinessResp
}
var file_imshortvideoserver_proto_depIdxs = []int32{
	0, // 0: imshortvideoserver.ServerService.ShortVideoBussiness:input_type -> imshortvideoserver.ShortVideoBussinessReq
	1, // 1: imshortvideoserver.ServerService.ShortVideoBussiness:output_type -> imshortvideoserver.ShortVideoBussinessResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imshortvideoserver_proto_init() }
func file_imshortvideoserver_proto_init() {
	if File_imshortvideoserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imshortvideoserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortVideoBussinessReq); i {
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
		file_imshortvideoserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortVideoBussinessResp); i {
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
			RawDescriptor: file_imshortvideoserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imshortvideoserver_proto_goTypes,
		DependencyIndexes: file_imshortvideoserver_proto_depIdxs,
		MessageInfos:      file_imshortvideoserver_proto_msgTypes,
	}.Build()
	File_imshortvideoserver_proto = out.File
	file_imshortvideoserver_proto_rawDesc = nil
	file_imshortvideoserver_proto_goTypes = nil
	file_imshortvideoserver_proto_depIdxs = nil
}
