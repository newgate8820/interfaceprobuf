// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imlogindayhisserver.proto

package imlogindayhisserver

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

type SetLoginHisDayResult_ResultCode int32

const (
	SetLoginHisDayResult_Success       SetLoginHisDayResult_ResultCode = 0
	SetLoginHisDayResult_InternalError SetLoginHisDayResult_ResultCode = 1
)

// Enum value maps for SetLoginHisDayResult_ResultCode.
var (
	SetLoginHisDayResult_ResultCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
	}
	SetLoginHisDayResult_ResultCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
	}
)

func (x SetLoginHisDayResult_ResultCode) Enum() *SetLoginHisDayResult_ResultCode {
	p := new(SetLoginHisDayResult_ResultCode)
	*p = x
	return p
}

func (x SetLoginHisDayResult_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetLoginHisDayResult_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_imlogindayhisserver_proto_enumTypes[0].Descriptor()
}

func (SetLoginHisDayResult_ResultCode) Type() protoreflect.EnumType {
	return &file_imlogindayhisserver_proto_enumTypes[0]
}

func (x SetLoginHisDayResult_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetLoginHisDayResult_ResultCode.Descriptor instead.
func (SetLoginHisDayResult_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{1, 0}
}

type GetNowDayHisDayResult_ResultCode int32

const (
	GetNowDayHisDayResult_Success       GetNowDayHisDayResult_ResultCode = 0
	GetNowDayHisDayResult_InternalError GetNowDayHisDayResult_ResultCode = 1
)

// Enum value maps for GetNowDayHisDayResult_ResultCode.
var (
	GetNowDayHisDayResult_ResultCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
	}
	GetNowDayHisDayResult_ResultCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
	}
)

func (x GetNowDayHisDayResult_ResultCode) Enum() *GetNowDayHisDayResult_ResultCode {
	p := new(GetNowDayHisDayResult_ResultCode)
	*p = x
	return p
}

func (x GetNowDayHisDayResult_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetNowDayHisDayResult_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_imlogindayhisserver_proto_enumTypes[1].Descriptor()
}

func (GetNowDayHisDayResult_ResultCode) Type() protoreflect.EnumType {
	return &file_imlogindayhisserver_proto_enumTypes[1]
}

func (x GetNowDayHisDayResult_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetNowDayHisDayResult_ResultCode.Descriptor instead.
func (GetNowDayHisDayResult_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{3, 0}
}

type GetHisCountResult_ResultCode int32

const (
	GetHisCountResult_Success       GetHisCountResult_ResultCode = 0
	GetHisCountResult_InternalError GetHisCountResult_ResultCode = 1
)

// Enum value maps for GetHisCountResult_ResultCode.
var (
	GetHisCountResult_ResultCode_name = map[int32]string{
		0: "Success",
		1: "InternalError",
	}
	GetHisCountResult_ResultCode_value = map[string]int32{
		"Success":       0,
		"InternalError": 1,
	}
)

func (x GetHisCountResult_ResultCode) Enum() *GetHisCountResult_ResultCode {
	p := new(GetHisCountResult_ResultCode)
	*p = x
	return p
}

func (x GetHisCountResult_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetHisCountResult_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_imlogindayhisserver_proto_enumTypes[2].Descriptor()
}

func (GetHisCountResult_ResultCode) Type() protoreflect.EnumType {
	return &file_imlogindayhisserver_proto_enumTypes[2]
}

func (x GetHisCountResult_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetHisCountResult_ResultCode.Descriptor instead.
func (GetHisCountResult_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{5, 0}
}

type SetLoginHisDayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ?????????id
}

func (x *SetLoginHisDayReq) Reset() {
	*x = SetLoginHisDayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLoginHisDayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLoginHisDayReq) ProtoMessage() {}

func (x *SetLoginHisDayReq) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLoginHisDayReq.ProtoReflect.Descriptor instead.
func (*SetLoginHisDayReq) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{0}
}

func (x *SetLoginHisDayReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SetLoginHisDayResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool                            `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`                                                      //
	Count  int32                           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                                                        //????????????
	Code   SetLoginHisDayResult_ResultCode `protobuf:"varint,3,opt,name=code,proto3,enum=imlogindayhisserver.SetLoginHisDayResult_ResultCode" json:"code,omitempty"` //????????????
}

func (x *SetLoginHisDayResult) Reset() {
	*x = SetLoginHisDayResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLoginHisDayResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLoginHisDayResult) ProtoMessage() {}

func (x *SetLoginHisDayResult) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLoginHisDayResult.ProtoReflect.Descriptor instead.
func (*SetLoginHisDayResult) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{1}
}

func (x *SetLoginHisDayResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *SetLoginHisDayResult) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SetLoginHisDayResult) GetCode() SetLoginHisDayResult_ResultCode {
	if x != nil {
		return x.Code
	}
	return SetLoginHisDayResult_Success
}

type GetNowDayHisDayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ?????????id
}

func (x *GetNowDayHisDayReq) Reset() {
	*x = GetNowDayHisDayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNowDayHisDayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNowDayHisDayReq) ProtoMessage() {}

func (x *GetNowDayHisDayReq) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNowDayHisDayReq.ProtoReflect.Descriptor instead.
func (*GetNowDayHisDayReq) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{2}
}

func (x *GetNowDayHisDayReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetNowDayHisDayResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool                             `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`                                                       //
	Code   GetNowDayHisDayResult_ResultCode `protobuf:"varint,2,opt,name=code,proto3,enum=imlogindayhisserver.GetNowDayHisDayResult_ResultCode" json:"code,omitempty"` //????????????
}

func (x *GetNowDayHisDayResult) Reset() {
	*x = GetNowDayHisDayResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNowDayHisDayResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNowDayHisDayResult) ProtoMessage() {}

func (x *GetNowDayHisDayResult) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNowDayHisDayResult.ProtoReflect.Descriptor instead.
func (*GetNowDayHisDayResult) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{3}
}

func (x *GetNowDayHisDayResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *GetNowDayHisDayResult) GetCode() GetNowDayHisDayResult_ResultCode {
	if x != nil {
		return x.Code
	}
	return GetNowDayHisDayResult_Success
}

type GetHisCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ?????????id
}

func (x *GetHisCountReq) Reset() {
	*x = GetHisCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHisCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHisCountReq) ProtoMessage() {}

func (x *GetHisCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHisCountReq.ProtoReflect.Descriptor instead.
func (*GetHisCountReq) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{4}
}

func (x *GetHisCountReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetHisCountResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32                        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Code  GetHisCountResult_ResultCode `protobuf:"varint,2,opt,name=code,proto3,enum=imlogindayhisserver.GetHisCountResult_ResultCode" json:"code,omitempty"`
}

func (x *GetHisCountResult) Reset() {
	*x = GetHisCountResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imlogindayhisserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHisCountResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHisCountResult) ProtoMessage() {}

func (x *GetHisCountResult) ProtoReflect() protoreflect.Message {
	mi := &file_imlogindayhisserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHisCountResult.ProtoReflect.Descriptor instead.
func (*GetHisCountResult) Descriptor() ([]byte, []int) {
	return file_imlogindayhisserver_proto_rawDescGZIP(), []int{5}
}

func (x *GetHisCountResult) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetHisCountResult) GetCode() GetHisCountResult_ResultCode {
	if x != nil {
		return x.Code
	}
	return GetHisCountResult_Success
}

var File_imlogindayhisserver_proto protoreflect.FileDescriptor

var file_imlogindayhisserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6d, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x2c, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x44,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xbc,
	0x01, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x44, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79,
	0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x2c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x22, 0x2d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x44, 0x61, 0x79, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x44, 0x61, 0x79, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x49,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x69,
	0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x44, 0x61, 0x79, 0x48, 0x69, 0x73, 0x44,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x69,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x69,
	0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x01, 0x32, 0xc4, 0x02, 0x0a, 0x13, 0x49, 0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44,
	0x61, 0x79, 0x48, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79, 0x12, 0x26, 0x2e,
	0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x44,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64,
	0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x68, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x44, 0x61, 0x79, 0x48,
	0x69, 0x73, 0x44, 0x61, 0x79, 0x12, 0x27, 0x2e, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64,
	0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x77, 0x44, 0x61, 0x79, 0x48, 0x69, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x2a,
	0x2e, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x44, 0x61, 0x79, 0x48, 0x69,
	0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x69, 0x6d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x26, 0x2e, 0x69, 0x6d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x79, 0x68, 0x69, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imlogindayhisserver_proto_rawDescOnce sync.Once
	file_imlogindayhisserver_proto_rawDescData = file_imlogindayhisserver_proto_rawDesc
)

func file_imlogindayhisserver_proto_rawDescGZIP() []byte {
	file_imlogindayhisserver_proto_rawDescOnce.Do(func() {
		file_imlogindayhisserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_imlogindayhisserver_proto_rawDescData)
	})
	return file_imlogindayhisserver_proto_rawDescData
}

var file_imlogindayhisserver_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_imlogindayhisserver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_imlogindayhisserver_proto_goTypes = []interface{}{
	(SetLoginHisDayResult_ResultCode)(0),  // 0: imlogindayhisserver.SetLoginHisDayResult.ResultCode
	(GetNowDayHisDayResult_ResultCode)(0), // 1: imlogindayhisserver.GetNowDayHisDayResult.ResultCode
	(GetHisCountResult_ResultCode)(0),     // 2: imlogindayhisserver.GetHisCountResult.ResultCode
	(*SetLoginHisDayReq)(nil),             // 3: imlogindayhisserver.SetLoginHisDayReq
	(*SetLoginHisDayResult)(nil),          // 4: imlogindayhisserver.SetLoginHisDayResult
	(*GetNowDayHisDayReq)(nil),            // 5: imlogindayhisserver.GetNowDayHisDayReq
	(*GetNowDayHisDayResult)(nil),         // 6: imlogindayhisserver.GetNowDayHisDayResult
	(*GetHisCountReq)(nil),                // 7: imlogindayhisserver.GetHisCountReq
	(*GetHisCountResult)(nil),             // 8: imlogindayhisserver.GetHisCountResult
}
var file_imlogindayhisserver_proto_depIdxs = []int32{
	0, // 0: imlogindayhisserver.SetLoginHisDayResult.code:type_name -> imlogindayhisserver.SetLoginHisDayResult.ResultCode
	1, // 1: imlogindayhisserver.GetNowDayHisDayResult.code:type_name -> imlogindayhisserver.GetNowDayHisDayResult.ResultCode
	2, // 2: imlogindayhisserver.GetHisCountResult.code:type_name -> imlogindayhisserver.GetHisCountResult.ResultCode
	3, // 3: imlogindayhisserver.ImLoginDayHisServer.SetLoginHisDay:input_type -> imlogindayhisserver.SetLoginHisDayReq
	5, // 4: imlogindayhisserver.ImLoginDayHisServer.GetNowDayHisDay:input_type -> imlogindayhisserver.GetNowDayHisDayReq
	7, // 5: imlogindayhisserver.ImLoginDayHisServer.GetHisCount:input_type -> imlogindayhisserver.GetHisCountReq
	4, // 6: imlogindayhisserver.ImLoginDayHisServer.SetLoginHisDay:output_type -> imlogindayhisserver.SetLoginHisDayResult
	6, // 7: imlogindayhisserver.ImLoginDayHisServer.GetNowDayHisDay:output_type -> imlogindayhisserver.GetNowDayHisDayResult
	8, // 8: imlogindayhisserver.ImLoginDayHisServer.GetHisCount:output_type -> imlogindayhisserver.GetHisCountResult
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_imlogindayhisserver_proto_init() }
func file_imlogindayhisserver_proto_init() {
	if File_imlogindayhisserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imlogindayhisserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLoginHisDayReq); i {
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
		file_imlogindayhisserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLoginHisDayResult); i {
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
		file_imlogindayhisserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNowDayHisDayReq); i {
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
		file_imlogindayhisserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNowDayHisDayResult); i {
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
		file_imlogindayhisserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHisCountReq); i {
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
		file_imlogindayhisserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHisCountResult); i {
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
			RawDescriptor: file_imlogindayhisserver_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imlogindayhisserver_proto_goTypes,
		DependencyIndexes: file_imlogindayhisserver_proto_depIdxs,
		EnumInfos:         file_imlogindayhisserver_proto_enumTypes,
		MessageInfos:      file_imlogindayhisserver_proto_msgTypes,
	}.Build()
	File_imlogindayhisserver_proto = out.File
	file_imlogindayhisserver_proto_rawDesc = nil
	file_imlogindayhisserver_proto_goTypes = nil
	file_imlogindayhisserver_proto_depIdxs = nil
}
