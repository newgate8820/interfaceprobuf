// protoc --gogofast_out=plugins=grpc:. *.proto
// sudo protoc --gogofast_out=plugins=grpc:. gitlab.chatserver.im/interfaceprobuf/imverifyservice/imverifyservice.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imverifyservice.proto

package imverifyservice

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

type VerifyServiceResultCode int32

const (
	VerifyServiceResultCode_ResultCode_Ok                             VerifyServiceResultCode = 0    // ok
	VerifyServiceResultCode_ResultCode_Token_Invalid                  VerifyServiceResultCode = 1001 // token无效
	VerifyServiceResultCode_ResultCode_Parameter_Exception            VerifyServiceResultCode = 1002 // 参数异常
	VerifyServiceResultCode_ResultCode_Cailiao_Side_Service_Exception VerifyServiceResultCode = 1003 // 彩聊侧服务异常
	VerifyServiceResultCode_ResultCode_Game_Side_Service_Exception    VerifyServiceResultCode = 1004 // 游戏侧服务异常
)

// Enum value maps for VerifyServiceResultCode.
var (
	VerifyServiceResultCode_name = map[int32]string{
		0:    "ResultCode_Ok",
		1001: "ResultCode_Token_Invalid",
		1002: "ResultCode_Parameter_Exception",
		1003: "ResultCode_Cailiao_Side_Service_Exception",
		1004: "ResultCode_Game_Side_Service_Exception",
	}
	VerifyServiceResultCode_value = map[string]int32{
		"ResultCode_Ok":                             0,
		"ResultCode_Token_Invalid":                  1001,
		"ResultCode_Parameter_Exception":            1002,
		"ResultCode_Cailiao_Side_Service_Exception": 1003,
		"ResultCode_Game_Side_Service_Exception":    1004,
	}
)

func (x VerifyServiceResultCode) Enum() *VerifyServiceResultCode {
	p := new(VerifyServiceResultCode)
	*p = x
	return p
}

func (x VerifyServiceResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerifyServiceResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_imverifyservice_proto_enumTypes[0].Descriptor()
}

func (VerifyServiceResultCode) Type() protoreflect.EnumType {
	return &file_imverifyservice_proto_enumTypes[0]
}

func (x VerifyServiceResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerifyServiceResultCode.Descriptor instead.
func (VerifyServiceResultCode) EnumDescriptor() ([]byte, []int) {
	return file_imverifyservice_proto_rawDescGZIP(), []int{0}
}

// 获取帐户token消息请求参数
type GetAccountTokenMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetAccountTokenMsg) Reset() {
	*x = GetAccountTokenMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imverifyservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountTokenMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountTokenMsg) ProtoMessage() {}

func (x *GetAccountTokenMsg) ProtoReflect() protoreflect.Message {
	mi := &file_imverifyservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountTokenMsg.ProtoReflect.Descriptor instead.
func (*GetAccountTokenMsg) Descriptor() ([]byte, []int) {
	return file_imverifyservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountTokenMsg) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 获取帐户token消息响应参数
type GetAccountTokenMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode VerifyServiceResultCode `protobuf:"varint,1,opt,name=ResultCode,proto3,enum=imverifyservice.VerifyServiceResultCode" json:"ResultCode,omitempty"`
	Token      string                  `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	URL        string                  `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *GetAccountTokenMsgReply) Reset() {
	*x = GetAccountTokenMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imverifyservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountTokenMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountTokenMsgReply) ProtoMessage() {}

func (x *GetAccountTokenMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_imverifyservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountTokenMsgReply.ProtoReflect.Descriptor instead.
func (*GetAccountTokenMsgReply) Descriptor() ([]byte, []int) {
	return file_imverifyservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountTokenMsgReply) GetResultCode() VerifyServiceResultCode {
	if x != nil {
		return x.ResultCode
	}
	return VerifyServiceResultCode_ResultCode_Ok
}

func (x *GetAccountTokenMsgReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAccountTokenMsgReply) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

// 验证帐户token消息请求参数
type VerifyAccountTokenMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *VerifyAccountTokenMsg) Reset() {
	*x = VerifyAccountTokenMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imverifyservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccountTokenMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccountTokenMsg) ProtoMessage() {}

func (x *VerifyAccountTokenMsg) ProtoReflect() protoreflect.Message {
	mi := &file_imverifyservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccountTokenMsg.ProtoReflect.Descriptor instead.
func (*VerifyAccountTokenMsg) Descriptor() ([]byte, []int) {
	return file_imverifyservice_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyAccountTokenMsg) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VerifyAccountTokenMsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 验证帐户token消息响应参数
type VerifyAccountTokenMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode VerifyServiceResultCode `protobuf:"varint,1,opt,name=ResultCode,proto3,enum=imverifyservice.VerifyServiceResultCode" json:"ResultCode,omitempty"`
	UserId     int32                   `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *VerifyAccountTokenMsgReply) Reset() {
	*x = VerifyAccountTokenMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imverifyservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccountTokenMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccountTokenMsgReply) ProtoMessage() {}

func (x *VerifyAccountTokenMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_imverifyservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccountTokenMsgReply.ProtoReflect.Descriptor instead.
func (*VerifyAccountTokenMsgReply) Descriptor() ([]byte, []int) {
	return file_imverifyservice_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyAccountTokenMsgReply) GetResultCode() VerifyServiceResultCode {
	if x != nil {
		return x.ResultCode
	}
	return VerifyServiceResultCode_ResultCode_Ok
}

func (x *VerifyAccountTokenMsgReply) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_imverifyservice_proto protoreflect.FileDescriptor

var file_imverifyservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x48, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x22, 0x45, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7e, 0x0a, 0x1a, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0xcd, 0x01, 0x0a, 0x17,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x18, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xe9, 0x07, 0x12, 0x23, 0x0a, 0x1e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xea, 0x07, 0x12, 0x2e,
	0x0a, 0x29, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x43, 0x61, 0x69,
	0x6c, 0x69, 0x61, 0x6f, 0x5f, 0x53, 0x69, 0x64, 0x65, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xeb, 0x07, 0x12, 0x2b,
	0x0a, 0x26, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x47, 0x61, 0x6d,
	0x65, 0x5f, 0x53, 0x69, 0x64, 0x65, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x45,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xec, 0x07, 0x32, 0xe2, 0x01, 0x0a, 0x0f,
	0x49, 0x6d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x2e, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x1a, 0x28, 0x2e, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e, 0x69, 0x6d, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73,
	0x67, 0x1a, 0x2b, 0x2e, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x21, 0x5a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imverifyservice_proto_rawDescOnce sync.Once
	file_imverifyservice_proto_rawDescData = file_imverifyservice_proto_rawDesc
)

func file_imverifyservice_proto_rawDescGZIP() []byte {
	file_imverifyservice_proto_rawDescOnce.Do(func() {
		file_imverifyservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_imverifyservice_proto_rawDescData)
	})
	return file_imverifyservice_proto_rawDescData
}

var file_imverifyservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_imverifyservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imverifyservice_proto_goTypes = []interface{}{
	(VerifyServiceResultCode)(0),       // 0: imverifyservice.VerifyServiceResultCode
	(*GetAccountTokenMsg)(nil),         // 1: imverifyservice.GetAccountTokenMsg
	(*GetAccountTokenMsgReply)(nil),    // 2: imverifyservice.GetAccountTokenMsgReply
	(*VerifyAccountTokenMsg)(nil),      // 3: imverifyservice.VerifyAccountTokenMsg
	(*VerifyAccountTokenMsgReply)(nil), // 4: imverifyservice.VerifyAccountTokenMsgReply
}
var file_imverifyservice_proto_depIdxs = []int32{
	0, // 0: imverifyservice.GetAccountTokenMsgReply.ResultCode:type_name -> imverifyservice.VerifyServiceResultCode
	0, // 1: imverifyservice.VerifyAccountTokenMsgReply.ResultCode:type_name -> imverifyservice.VerifyServiceResultCode
	1, // 2: imverifyservice.ImVerifyService.GetAccountToken:input_type -> imverifyservice.GetAccountTokenMsg
	3, // 3: imverifyservice.ImVerifyService.VerifyAccountToken:input_type -> imverifyservice.VerifyAccountTokenMsg
	2, // 4: imverifyservice.ImVerifyService.GetAccountToken:output_type -> imverifyservice.GetAccountTokenMsgReply
	4, // 5: imverifyservice.ImVerifyService.VerifyAccountToken:output_type -> imverifyservice.VerifyAccountTokenMsgReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_imverifyservice_proto_init() }
func file_imverifyservice_proto_init() {
	if File_imverifyservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imverifyservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountTokenMsg); i {
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
		file_imverifyservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountTokenMsgReply); i {
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
		file_imverifyservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccountTokenMsg); i {
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
		file_imverifyservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccountTokenMsgReply); i {
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
			RawDescriptor: file_imverifyservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imverifyservice_proto_goTypes,
		DependencyIndexes: file_imverifyservice_proto_depIdxs,
		EnumInfos:         file_imverifyservice_proto_enumTypes,
		MessageInfos:      file_imverifyservice_proto_msgTypes,
	}.Build()
	File_imverifyservice_proto = out.File
	file_imverifyservice_proto_rawDesc = nil
	file_imverifyservice_proto_goTypes = nil
	file_imverifyservice_proto_depIdxs = nil
}
