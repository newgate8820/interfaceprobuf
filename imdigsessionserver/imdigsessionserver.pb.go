// sudo protoc --gogofast_out=plugins=grpc:. gitlab.chatserver.im/interfaceprobuf/imdigsessionserver/imdigsessionserver.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imdigsessionserver.proto

package imdigsessionserver

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

type TokenErrorCode int32

const (
	TokenErrorCode_Token_OK                TokenErrorCode = 0    //ok
	TokenErrorCode_Token_REQ_PARAME_ERR    TokenErrorCode = 1001 //请求参数错误
	TokenErrorCode_Token_CTEATE_TOKEN_FAIL TokenErrorCode = 1002 //获取token失败
	TokenErrorCode_Token_ERR               TokenErrorCode = 1003 //token错误
	TokenErrorCode_Token_EXPIRED           TokenErrorCode = 1004 //token超时失效
	TokenErrorCode_Token_NOT_EXIST         TokenErrorCode = 1005 //token不存在
	TokenErrorCode_Token_INTERNAL_ERR      TokenErrorCode = 1006 //内部错误
)

// Enum value maps for TokenErrorCode.
var (
	TokenErrorCode_name = map[int32]string{
		0:    "Token_OK",
		1001: "Token_REQ_PARAME_ERR",
		1002: "Token_CTEATE_TOKEN_FAIL",
		1003: "Token_ERR",
		1004: "Token_EXPIRED",
		1005: "Token_NOT_EXIST",
		1006: "Token_INTERNAL_ERR",
	}
	TokenErrorCode_value = map[string]int32{
		"Token_OK":                0,
		"Token_REQ_PARAME_ERR":    1001,
		"Token_CTEATE_TOKEN_FAIL": 1002,
		"Token_ERR":               1003,
		"Token_EXPIRED":           1004,
		"Token_NOT_EXIST":         1005,
		"Token_INTERNAL_ERR":      1006,
	}
)

func (x TokenErrorCode) Enum() *TokenErrorCode {
	p := new(TokenErrorCode)
	*p = x
	return p
}

func (x TokenErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_imdigsessionserver_proto_enumTypes[0].Descriptor()
}

func (TokenErrorCode) Type() protoreflect.EnumType {
	return &file_imdigsessionserver_proto_enumTypes[0]
}

func (x TokenErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenErrorCode.Descriptor instead.
func (TokenErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_imdigsessionserver_proto_rawDescGZIP(), []int{0}
}

type ReqGetTokenMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyId  string `protobuf:"bytes,1,opt,name=keyId,proto3" json:"keyId,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"` //
	Nonce  string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`    //上传的随机字符串长度为16字节
}

func (x *ReqGetTokenMsg) Reset() {
	*x = ReqGetTokenMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imdigsessionserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetTokenMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetTokenMsg) ProtoMessage() {}

func (x *ReqGetTokenMsg) ProtoReflect() protoreflect.Message {
	mi := &file_imdigsessionserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetTokenMsg.ProtoReflect.Descriptor instead.
func (*ReqGetTokenMsg) Descriptor() ([]byte, []int) {
	return file_imdigsessionserver_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGetTokenMsg) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *ReqGetTokenMsg) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqGetTokenMsg) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type GetTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           TokenErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=imdigsessionserver.TokenErrorCode" json:"code,omitempty"`
	Error          string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token          string         `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Nonce          string         `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`                    //随机字符串长度为16字节
	ExpirationTime uint64         `protobuf:"varint,5,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"` //过期时间戳（秒级）
}

func (x *GetTokenReply) Reset() {
	*x = GetTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imdigsessionserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenReply) ProtoMessage() {}

func (x *GetTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_imdigsessionserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenReply.ProtoReflect.Descriptor instead.
func (*GetTokenReply) Descriptor() ([]byte, []int) {
	return file_imdigsessionserver_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenReply) GetCode() TokenErrorCode {
	if x != nil {
		return x.Code
	}
	return TokenErrorCode_Token_OK
}

func (x *GetTokenReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTokenReply) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *GetTokenReply) GetExpirationTime() uint64 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

type ReqGetUserToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReqGetUserToken) Reset() {
	*x = ReqGetUserToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imdigsessionserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetUserToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetUserToken) ProtoMessage() {}

func (x *ReqGetUserToken) ProtoReflect() protoreflect.Message {
	mi := &file_imdigsessionserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetUserToken.ProtoReflect.Descriptor instead.
func (*ReqGetUserToken) Descriptor() ([]byte, []int) {
	return file_imdigsessionserver_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGetUserToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetUserTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           TokenErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=imdigsessionserver.TokenErrorCode" json:"code,omitempty"`
	UserId         int64          `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Key            string         `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	ExpirationTime uint64         `protobuf:"varint,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"` //过期时间戳（秒级）
}

func (x *GetUserTokenReply) Reset() {
	*x = GetUserTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imdigsessionserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReply) ProtoMessage() {}

func (x *GetUserTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_imdigsessionserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReply.ProtoReflect.Descriptor instead.
func (*GetUserTokenReply) Descriptor() ([]byte, []int) {
	return file_imdigsessionserver_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserTokenReply) GetCode() TokenErrorCode {
	if x != nil {
		return x.Code
	}
	return TokenErrorCode_Token_OK
}

func (x *GetUserTokenReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserTokenReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetUserTokenReply) GetExpirationTime() uint64 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

var File_imdigsessionserver_proto protoreflect.FileDescriptor

var file_imdigsessionserver_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69, 0x6d, 0x64, 0x69,
	0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x54,
	0x0a, 0x0e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x2a, 0xaa, 0x01, 0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x52, 0x45, 0x51, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xe9, 0x07, 0x12, 0x1c, 0x0a,
	0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x43, 0x54, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xea, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xeb, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0xec, 0x07, 0x12,
	0x14, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x10, 0xed, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xee, 0x07, 0x32, 0xc8,
	0x01, 0x0a, 0x13, 0x49, 0x6d, 0x44, 0x69, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x22, 0x2e, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x1a, 0x21, 0x2e, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x69, 0x6d,
	0x64, 0x69, 0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x25, 0x2e, 0x69, 0x6d, 0x64, 0x69, 0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x64,
	0x69, 0x67, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imdigsessionserver_proto_rawDescOnce sync.Once
	file_imdigsessionserver_proto_rawDescData = file_imdigsessionserver_proto_rawDesc
)

func file_imdigsessionserver_proto_rawDescGZIP() []byte {
	file_imdigsessionserver_proto_rawDescOnce.Do(func() {
		file_imdigsessionserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_imdigsessionserver_proto_rawDescData)
	})
	return file_imdigsessionserver_proto_rawDescData
}

var file_imdigsessionserver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_imdigsessionserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imdigsessionserver_proto_goTypes = []interface{}{
	(TokenErrorCode)(0),       // 0: imdigsessionserver.TokenErrorCode
	(*ReqGetTokenMsg)(nil),    // 1: imdigsessionserver.ReqGetTokenMsg
	(*GetTokenReply)(nil),     // 2: imdigsessionserver.GetTokenReply
	(*ReqGetUserToken)(nil),   // 3: imdigsessionserver.ReqGetUserToken
	(*GetUserTokenReply)(nil), // 4: imdigsessionserver.GetUserTokenReply
}
var file_imdigsessionserver_proto_depIdxs = []int32{
	0, // 0: imdigsessionserver.GetTokenReply.code:type_name -> imdigsessionserver.TokenErrorCode
	0, // 1: imdigsessionserver.GetUserTokenReply.code:type_name -> imdigsessionserver.TokenErrorCode
	1, // 2: imdigsessionserver.ImDigSessionService.GetToken:input_type -> imdigsessionserver.ReqGetTokenMsg
	3, // 3: imdigsessionserver.ImDigSessionService.GetUserToken:input_type -> imdigsessionserver.ReqGetUserToken
	2, // 4: imdigsessionserver.ImDigSessionService.GetToken:output_type -> imdigsessionserver.GetTokenReply
	4, // 5: imdigsessionserver.ImDigSessionService.GetUserToken:output_type -> imdigsessionserver.GetUserTokenReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_imdigsessionserver_proto_init() }
func file_imdigsessionserver_proto_init() {
	if File_imdigsessionserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imdigsessionserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetTokenMsg); i {
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
		file_imdigsessionserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenReply); i {
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
		file_imdigsessionserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetUserToken); i {
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
		file_imdigsessionserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReply); i {
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
			RawDescriptor: file_imdigsessionserver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imdigsessionserver_proto_goTypes,
		DependencyIndexes: file_imdigsessionserver_proto_depIdxs,
		EnumInfos:         file_imdigsessionserver_proto_enumTypes,
		MessageInfos:      file_imdigsessionserver_proto_msgTypes,
	}.Build()
	File_imdigsessionserver_proto = out.File
	file_imdigsessionserver_proto_rawDesc = nil
	file_imdigsessionserver_proto_goTypes = nil
	file_imdigsessionserver_proto_depIdxs = nil
}
