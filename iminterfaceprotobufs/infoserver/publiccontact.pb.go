// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: publiccontact.proto

package infoserver

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

type ReqUpdateMoments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId []int32 `protobuf:"varint,1,rep,packed,name=userId,proto3" json:"userId,omitempty"` //被更新用户ID列表
	SelfId int32   `protobuf:"varint,2,opt,name=selfId,proto3" json:"selfId,omitempty"`        //发送用户id
}

func (x *ReqUpdateMoments) Reset() {
	*x = ReqUpdateMoments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publiccontact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateMoments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateMoments) ProtoMessage() {}

func (x *ReqUpdateMoments) ProtoReflect() protoreflect.Message {
	mi := &file_publiccontact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateMoments.ProtoReflect.Descriptor instead.
func (*ReqUpdateMoments) Descriptor() ([]byte, []int) {
	return file_publiccontact_proto_rawDescGZIP(), []int{0}
}

func (x *ReqUpdateMoments) GetUserId() []int32 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ReqUpdateMoments) GetSelfId() int32 {
	if x != nil {
		return x.SelfId
	}
	return 0
}

type ResUpdateMoments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` //1为成功,-1为失败
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    //返回描述
}

func (x *ResUpdateMoments) Reset() {
	*x = ResUpdateMoments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publiccontact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResUpdateMoments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResUpdateMoments) ProtoMessage() {}

func (x *ResUpdateMoments) ProtoReflect() protoreflect.Message {
	mi := &file_publiccontact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResUpdateMoments.ProtoReflect.Descriptor instead.
func (*ResUpdateMoments) Descriptor() ([]byte, []int) {
	return file_publiccontact_proto_rawDescGZIP(), []int{1}
}

func (x *ResUpdateMoments) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResUpdateMoments) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_publiccontact_proto protoreflect.FileDescriptor

var file_publiccontact_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x42, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x65, 0x6c, 0x66, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0x63, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publiccontact_proto_rawDescOnce sync.Once
	file_publiccontact_proto_rawDescData = file_publiccontact_proto_rawDesc
)

func file_publiccontact_proto_rawDescGZIP() []byte {
	file_publiccontact_proto_rawDescOnce.Do(func() {
		file_publiccontact_proto_rawDescData = protoimpl.X.CompressGZIP(file_publiccontact_proto_rawDescData)
	})
	return file_publiccontact_proto_rawDescData
}

var file_publiccontact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_publiccontact_proto_goTypes = []interface{}{
	(*ReqUpdateMoments)(nil), // 0: infoserver.ReqUpdateMoments
	(*ResUpdateMoments)(nil), // 1: infoserver.ResUpdateMoments
}
var file_publiccontact_proto_depIdxs = []int32{
	0, // 0: infoserver.ContactPushService.UpdateMoments:input_type -> infoserver.ReqUpdateMoments
	1, // 1: infoserver.ContactPushService.UpdateMoments:output_type -> infoserver.ResUpdateMoments
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publiccontact_proto_init() }
func file_publiccontact_proto_init() {
	if File_publiccontact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publiccontact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdateMoments); i {
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
		file_publiccontact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResUpdateMoments); i {
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
			RawDescriptor: file_publiccontact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publiccontact_proto_goTypes,
		DependencyIndexes: file_publiccontact_proto_depIdxs,
		MessageInfos:      file_publiccontact_proto_msgTypes,
	}.Build()
	File_publiccontact_proto = out.File
	file_publiccontact_proto_rawDesc = nil
	file_publiccontact_proto_goTypes = nil
	file_publiccontact_proto_depIdxs = nil
}
