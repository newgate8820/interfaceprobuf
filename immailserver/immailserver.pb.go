// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: immailserver.proto

package immailserver

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

type SendMailResult_ResultCode int32

const (
	SendMailResult_Success       SendMailResult_ResultCode = 0
	SendMailResult_InternalError SendMailResult_ResultCode = 1
	SendMailResult_SendTooFast   SendMailResult_ResultCode = 2
	SendMailResult_NotFountEmail SendMailResult_ResultCode = 3
)

// Enum value maps for SendMailResult_ResultCode.
var (
	SendMailResult_ResultCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
		2: "SendTooFast",
		3: "NotFountEmail",
	}
	SendMailResult_ResultCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
		"SendTooFast":   2,
		"NotFountEmail": 3,
	}
)

func (x SendMailResult_ResultCode) Enum() *SendMailResult_ResultCode {
	p := new(SendMailResult_ResultCode)
	*p = x
	return p
}

func (x SendMailResult_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMailResult_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_immailserver_proto_enumTypes[0].Descriptor()
}

func (SendMailResult_ResultCode) Type() protoreflect.EnumType {
	return &file_immailserver_proto_enumTypes[0]
}

func (x SendMailResult_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMailResult_ResultCode.Descriptor instead.
func (SendMailResult_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{1, 0}
}

type SendMailSSLResult_ResultCode int32

const (
	SendMailSSLResult_Success       SendMailSSLResult_ResultCode = 0
	SendMailSSLResult_InternalError SendMailSSLResult_ResultCode = 1
)

// Enum value maps for SendMailSSLResult_ResultCode.
var (
	SendMailSSLResult_ResultCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
	}
	SendMailSSLResult_ResultCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
	}
)

func (x SendMailSSLResult_ResultCode) Enum() *SendMailSSLResult_ResultCode {
	p := new(SendMailSSLResult_ResultCode)
	*p = x
	return p
}

func (x SendMailSSLResult_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMailSSLResult_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_immailserver_proto_enumTypes[1].Descriptor()
}

func (SendMailSSLResult_ResultCode) Type() protoreflect.EnumType {
	return &file_immailserver_proto_enumTypes[1]
}

func (x SendMailSSLResult_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMailSSLResult_ResultCode.Descriptor instead.
func (SendMailSSLResult_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{3, 0}
}

type SendMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`  //用户id
	ToAddrs []string `protobuf:"bytes,2,rep,name=toAddrs,proto3" json:"toAddrs,omitempty"` // 发送地址
	Subject string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"` //标题
	Content []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"` //内容
}

func (x *SendMailReq) Reset() {
	*x = SendMailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_immailserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailReq) ProtoMessage() {}

func (x *SendMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_immailserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailReq.ProtoReflect.Descriptor instead.
func (*SendMailReq) Descriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendMailReq) GetToAddrs() []string {
	if x != nil {
		return x.ToAddrs
	}
	return nil
}

func (x *SendMailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailReq) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type SendMailResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code SendMailResult_ResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=immailserver.SendMailResult_ResultCode" json:"code,omitempty"`
	Err  string                    `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SendMailResult) Reset() {
	*x = SendMailResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_immailserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResult) ProtoMessage() {}

func (x *SendMailResult) ProtoReflect() protoreflect.Message {
	mi := &file_immailserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResult.ProtoReflect.Descriptor instead.
func (*SendMailResult) Descriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailResult) GetCode() SendMailResult_ResultCode {
	if x != nil {
		return x.Code
	}
	return SendMailResult_Success
}

func (x *SendMailResult) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type SendMailSSLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToAddrs []string `protobuf:"bytes,1,rep,name=toAddrs,proto3" json:"toAddrs,omitempty"` // 发送地址
	Subject string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"` //标题
	Content []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"` //内容
}

func (x *SendMailSSLReq) Reset() {
	*x = SendMailSSLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_immailserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailSSLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailSSLReq) ProtoMessage() {}

func (x *SendMailSSLReq) ProtoReflect() protoreflect.Message {
	mi := &file_immailserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailSSLReq.ProtoReflect.Descriptor instead.
func (*SendMailSSLReq) Descriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{2}
}

func (x *SendMailSSLReq) GetToAddrs() []string {
	if x != nil {
		return x.ToAddrs
	}
	return nil
}

func (x *SendMailSSLReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailSSLReq) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type SendMailSSLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code SendMailSSLResult_ResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=immailserver.SendMailSSLResult_ResultCode" json:"code,omitempty"` //接口标识
	Err  string                       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`                                                   //错误消息
}

func (x *SendMailSSLResult) Reset() {
	*x = SendMailSSLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_immailserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailSSLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailSSLResult) ProtoMessage() {}

func (x *SendMailSSLResult) ProtoReflect() protoreflect.Message {
	mi := &file_immailserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailSSLResult.ProtoReflect.Descriptor instead.
func (*SendMailSSLResult) Descriptor() ([]byte, []int) {
	return file_immailserver_proto_rawDescGZIP(), []int{3}
}

func (x *SendMailSSLResult) GetCode() SendMailSSLResult_ResultCode {
	if x != nil {
		return x.Code
	}
	return SendMailSSLResult_Success
}

func (x *SendMailSSLResult) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_immailserver_proto protoreflect.FileDescriptor

var file_immailserver_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x6d, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x6d, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0x73, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x41, 0x64,
	0x64, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x69, 0x6d, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x50, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x6f, 0x6f, 0x46, 0x61, 0x73, 0x74, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x03, 0x22, 0x5e, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x11,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x53, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x3e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x69, 0x6d, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x53, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x22, 0x2c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x01, 0x32, 0xa5, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x19,
	0x2e, 0x69, 0x6d, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x69, 0x6d, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x53, 0x4c, 0x12, 0x1c, 0x2e, 0x69, 0x6d, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x53, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x69, 0x6d, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x53,
	0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_immailserver_proto_rawDescOnce sync.Once
	file_immailserver_proto_rawDescData = file_immailserver_proto_rawDesc
)

func file_immailserver_proto_rawDescGZIP() []byte {
	file_immailserver_proto_rawDescOnce.Do(func() {
		file_immailserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_immailserver_proto_rawDescData)
	})
	return file_immailserver_proto_rawDescData
}

var file_immailserver_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_immailserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_immailserver_proto_goTypes = []interface{}{
	(SendMailResult_ResultCode)(0),    // 0: immailserver.SendMailResult.ResultCode
	(SendMailSSLResult_ResultCode)(0), // 1: immailserver.SendMailSSLResult.ResultCode
	(*SendMailReq)(nil),               // 2: immailserver.SendMailReq
	(*SendMailResult)(nil),            // 3: immailserver.SendMailResult
	(*SendMailSSLReq)(nil),            // 4: immailserver.SendMailSSLReq
	(*SendMailSSLResult)(nil),         // 5: immailserver.SendMailSSLResult
}
var file_immailserver_proto_depIdxs = []int32{
	0, // 0: immailserver.SendMailResult.code:type_name -> immailserver.SendMailResult.ResultCode
	1, // 1: immailserver.SendMailSSLResult.code:type_name -> immailserver.SendMailSSLResult.ResultCode
	2, // 2: immailserver.ImMailServer.SendMail:input_type -> immailserver.SendMailReq
	4, // 3: immailserver.ImMailServer.SendMailSSL:input_type -> immailserver.SendMailSSLReq
	3, // 4: immailserver.ImMailServer.SendMail:output_type -> immailserver.SendMailResult
	5, // 5: immailserver.ImMailServer.SendMailSSL:output_type -> immailserver.SendMailSSLResult
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_immailserver_proto_init() }
func file_immailserver_proto_init() {
	if File_immailserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_immailserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailReq); i {
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
		file_immailserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResult); i {
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
		file_immailserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailSSLReq); i {
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
		file_immailserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailSSLResult); i {
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
			RawDescriptor: file_immailserver_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_immailserver_proto_goTypes,
		DependencyIndexes: file_immailserver_proto_depIdxs,
		EnumInfos:         file_immailserver_proto_enumTypes,
		MessageInfos:      file_immailserver_proto_msgTypes,
	}.Build()
	File_immailserver_proto = out.File
	file_immailserver_proto_rawDesc = nil
	file_immailserver_proto_goTypes = nil
	file_immailserver_proto_depIdxs = nil
}
