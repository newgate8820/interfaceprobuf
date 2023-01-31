//protoc --gogofast_out=plugins=grpc:. ./gitlab.chatserver.im/interfaceprobuf/walletidserver/userid.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: userid.proto

package walletidserver

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

type ReErrCode int32

const (
	ReErrCode_Success       ReErrCode = 0 //成功
	ReErrCode_InternalError ReErrCode = 1 //内部错误
)

// Enum value maps for ReErrCode.
var (
	ReErrCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
	}
	ReErrCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
	}
)

func (x ReErrCode) Enum() *ReErrCode {
	p := new(ReErrCode)
	*p = x
	return p
}

func (x ReErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_userid_proto_enumTypes[0].Descriptor()
}

func (ReErrCode) Type() protoreflect.EnumType {
	return &file_userid_proto_enumTypes[0]
}

func (x ReErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReErrCode.Descriptor instead.
func (ReErrCode) EnumDescriptor() ([]byte, []int) {
	return file_userid_proto_rawDescGZIP(), []int{0}
}

type GetUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserIdReq) Reset() {
	*x = GetUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIdReq) ProtoMessage() {}

func (x *GetUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_userid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIdReq.ProtoReflect.Descriptor instead.
func (*GetUserIdReq) Descriptor() ([]byte, []int) {
	return file_userid_proto_rawDescGZIP(), []int{0}
}

type GetUserIdRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                        //获取的用户ＩＤ
	ErrorCode ReErrCode `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=walletidserver.ReErrCode" json:"error_code,omitempty"` //错误码
}

func (x *GetUserIdRep) Reset() {
	*x = GetUserIdRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIdRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIdRep) ProtoMessage() {}

func (x *GetUserIdRep) ProtoReflect() protoreflect.Message {
	mi := &file_userid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIdRep.ProtoReflect.Descriptor instead.
func (*GetUserIdRep) Descriptor() ([]byte, []int) {
	return file_userid_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserIdRep) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserIdRep) GetErrorCode() ReErrCode {
	if x != nil {
		return x.ErrorCode
	}
	return ReErrCode_Success
}

var File_userid_proto protoreflect.FileDescriptor

var file_userid_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x69, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x0e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x61,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x70, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x69, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x45,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x2a, 0x2b, 0x0a, 0x09, 0x52, 0x65, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x32, 0x56,
	0x0a, 0x09, 0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x69, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x69,
	0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x69, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userid_proto_rawDescOnce sync.Once
	file_userid_proto_rawDescData = file_userid_proto_rawDesc
)

func file_userid_proto_rawDescGZIP() []byte {
	file_userid_proto_rawDescOnce.Do(func() {
		file_userid_proto_rawDescData = protoimpl.X.CompressGZIP(file_userid_proto_rawDescData)
	})
	return file_userid_proto_rawDescData
}

var file_userid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userid_proto_goTypes = []interface{}{
	(ReErrCode)(0),       // 0: walletidserver.ReErrCode
	(*GetUserIdReq)(nil), // 1: walletidserver.GetUserIdReq
	(*GetUserIdRep)(nil), // 2: walletidserver.GetUserIdRep
}
var file_userid_proto_depIdxs = []int32{
	0, // 0: walletidserver.GetUserIdRep.error_code:type_name -> walletidserver.ReErrCode
	1, // 1: walletidserver.IdService.GetUserId:input_type -> walletidserver.GetUserIdReq
	2, // 2: walletidserver.IdService.GetUserId:output_type -> walletidserver.GetUserIdRep
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userid_proto_init() }
func file_userid_proto_init() {
	if File_userid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIdReq); i {
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
		file_userid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIdRep); i {
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
			RawDescriptor: file_userid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userid_proto_goTypes,
		DependencyIndexes: file_userid_proto_depIdxs,
		EnumInfos:         file_userid_proto_enumTypes,
		MessageInfos:      file_userid_proto_msgTypes,
	}.Build()
	File_userid_proto = out.File
	file_userid_proto_rawDesc = nil
	file_userid_proto_goTypes = nil
	file_userid_proto_depIdxs = nil
}
