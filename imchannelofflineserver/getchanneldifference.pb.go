// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: getchanneldifference.proto

package imchannelofflineserver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	dialogserver "interfaceprobuf/dialogserver"
	pbcomm "interfaceprobuf/pbcomm"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// #####################超级群get diff###################
type ReqGetChannelDifference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"` //用户ID
	KeyId       int64  `protobuf:"varint,2,opt,name=KeyId,proto3" json:"KeyId,omitempty"`   //用户KeyId
	ReMessageId uint64 `protobuf:"varint,18,opt,name=reMessageId,proto3" json:"reMessageId,omitempty"`
	CrcId       uint32 `protobuf:"varint,12,opt,name=CrcId,proto3" json:"CrcId,omitempty"`
	Flags       int32  `protobuf:"varint,8,opt,name=Flags,proto3" json:"Flags,omitempty"`
	Force       bool   `protobuf:"varint,9,opt,name=Force,proto3" json:"Force,omitempty"`
	Device      int32  `protobuf:"varint,6,opt,name=device,proto3" json:"device,omitempty"` //1, pc 2, android 3, ios
	ReqBin      []byte `protobuf:"bytes,7,opt,name=reqBin,proto3" json:"reqBin,omitempty"`
	RangeMinId  int32  `protobuf:"varint,10,opt,name=rangeMinId,proto3" json:"rangeMinId,omitempty"`
	RangeMaxId  int32  `protobuf:"varint,11,opt,name=rangeMaxId,proto3" json:"rangeMaxId,omitempty"`
	Layer       int32  `protobuf:"varint,3,opt,name=layer,proto3" json:"layer,omitempty"` // 层
	AppVision   string `protobuf:"bytes,4,opt,name=appVision,proto3" json:"appVision,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Ip          string `protobuf:"bytes,13,opt,name=ip,proto3" json:"ip,omitempty"`
	SessionId   uint64 `protobuf:"varint,14,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Ostype      uint32 `protobuf:"varint,16,opt,name=ostype,proto3" json:"ostype,omitempty"`
	CurrentCrc  uint32 `protobuf:"varint,17,opt,name=currentCrc,proto3" json:"currentCrc,omitempty"`
	ChannelId   int32  `protobuf:"varint,19,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *ReqGetChannelDifference) Reset() {
	*x = ReqGetChannelDifference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getchanneldifference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetChannelDifference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetChannelDifference) ProtoMessage() {}

func (x *ReqGetChannelDifference) ProtoReflect() protoreflect.Message {
	mi := &file_getchanneldifference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetChannelDifference.ProtoReflect.Descriptor instead.
func (*ReqGetChannelDifference) Descriptor() ([]byte, []int) {
	return file_getchanneldifference_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGetChannelDifference) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetKeyId() int64 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetReMessageId() uint64 {
	if x != nil {
		return x.ReMessageId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetCrcId() uint32 {
	if x != nil {
		return x.CrcId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *ReqGetChannelDifference) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *ReqGetChannelDifference) GetDevice() int32 {
	if x != nil {
		return x.Device
	}
	return 0
}

func (x *ReqGetChannelDifference) GetReqBin() []byte {
	if x != nil {
		return x.ReqBin
	}
	return nil
}

func (x *ReqGetChannelDifference) GetRangeMinId() int32 {
	if x != nil {
		return x.RangeMinId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetRangeMaxId() int32 {
	if x != nil {
		return x.RangeMaxId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetLayer() int32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *ReqGetChannelDifference) GetAppVision() string {
	if x != nil {
		return x.AppVision
	}
	return ""
}

func (x *ReqGetChannelDifference) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ReqGetChannelDifference) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ReqGetChannelDifference) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *ReqGetChannelDifference) GetOstype() uint32 {
	if x != nil {
		return x.Ostype
	}
	return 0
}

func (x *ReqGetChannelDifference) GetCurrentCrc() uint32 {
	if x != nil {
		return x.CurrentCrc
	}
	return 0
}

func (x *ReqGetChannelDifference) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

// 入参结构体
type ChanDiffStr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId      int32                    `protobuf:"varint,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	FromKeyId   int64                    `protobuf:"varint,2,opt,name=fromKeyId,proto3" json:"fromKeyId,omitempty"`
	Flags       int32                    `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	Force       bool                     `protobuf:"varint,4,opt,name=force,proto3" json:"force,omitempty"`
	ChannelId   int32                    `protobuf:"varint,5,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Pts         int32                    `protobuf:"varint,6,opt,name=pts,proto3" json:"pts,omitempty"`
	Limit       int32                    `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Range_MinId int32                    `protobuf:"varint,8,opt,name=range_MinId,json=rangeMinId,proto3" json:"range_MinId,omitempty"`
	Range_MaxId int32                    `protobuf:"varint,9,opt,name=range_MaxId,json=rangeMaxId,proto3" json:"range_MaxId,omitempty"`
	Device      int32                    `protobuf:"varint,10,opt,name=device,proto3" json:"device,omitempty"` //1, pc 2, android 3, ios
	ReMessageId uint64                   `protobuf:"varint,13,opt,name=reMessageId,proto3" json:"reMessageId,omitempty"`
	UserDialog  *dialogserver.UserDialog `protobuf:"bytes,11,opt,name=userDialog,proto3" json:"userDialog,omitempty"`
	Debug       *pbcomm.Debug            `protobuf:"bytes,12,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *ChanDiffStr) Reset() {
	*x = ChanDiffStr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getchanneldifference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChanDiffStr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChanDiffStr) ProtoMessage() {}

func (x *ChanDiffStr) ProtoReflect() protoreflect.Message {
	mi := &file_getchanneldifference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChanDiffStr.ProtoReflect.Descriptor instead.
func (*ChanDiffStr) Descriptor() ([]byte, []int) {
	return file_getchanneldifference_proto_rawDescGZIP(), []int{1}
}

func (x *ChanDiffStr) GetFromId() int32 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *ChanDiffStr) GetFromKeyId() int64 {
	if x != nil {
		return x.FromKeyId
	}
	return 0
}

func (x *ChanDiffStr) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *ChanDiffStr) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *ChanDiffStr) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ChanDiffStr) GetPts() int32 {
	if x != nil {
		return x.Pts
	}
	return 0
}

func (x *ChanDiffStr) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ChanDiffStr) GetRange_MinId() int32 {
	if x != nil {
		return x.Range_MinId
	}
	return 0
}

func (x *ChanDiffStr) GetRange_MaxId() int32 {
	if x != nil {
		return x.Range_MaxId
	}
	return 0
}

func (x *ChanDiffStr) GetDevice() int32 {
	if x != nil {
		return x.Device
	}
	return 0
}

func (x *ChanDiffStr) GetReMessageId() uint64 {
	if x != nil {
		return x.ReMessageId
	}
	return 0
}

func (x *ChanDiffStr) GetUserDialog() *dialogserver.UserDialog {
	if x != nil {
		return x.UserDialog
	}
	return nil
}

func (x *ChanDiffStr) GetDebug() *pbcomm.Debug {
	if x != nil {
		return x.Debug
	}
	return nil
}

// 返回
type RspGetChannelDifference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
}

func (x *RspGetChannelDifference) Reset() {
	*x = RspGetChannelDifference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getchanneldifference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspGetChannelDifference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspGetChannelDifference) ProtoMessage() {}

func (x *RspGetChannelDifference) ProtoReflect() protoreflect.Message {
	mi := &file_getchanneldifference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspGetChannelDifference.ProtoReflect.Descriptor instead.
func (*RspGetChannelDifference) Descriptor() ([]byte, []int) {
	return file_getchanneldifference_proto_rawDescGZIP(), []int{2}
}

func (x *RspGetChannelDifference) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *RspGetChannelDifference) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

var File_getchanneldifference_proto protoreflect.FileDescriptor

var file_getchanneldifference_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x65, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x1a, 0x1f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x70, 0x62,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x17, 0x52,
	0x65, 0x71, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4b,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x72, 0x63, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x72, 0x63, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x42, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x72, 0x65, 0x71, 0x42, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x4d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x61, 0x78, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x73, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x43, 0x72, 0x63, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x90, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x44,
	0x69, 0x66, 0x66, 0x53, 0x74, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x4d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x4d, 0x61, 0x78, 0x49, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x4f, 0x0a, 0x17, 0x52, 0x73, 0x70,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x98, 0x02, 0x0a, 0x1b, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x1a, 0x2f, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x73, 0x70,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6f,
	0x73, 0x12, 0x2f, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x1a, 0x2f, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x73, 0x70, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_getchanneldifference_proto_rawDescOnce sync.Once
	file_getchanneldifference_proto_rawDescData = file_getchanneldifference_proto_rawDesc
)

func file_getchanneldifference_proto_rawDescGZIP() []byte {
	file_getchanneldifference_proto_rawDescOnce.Do(func() {
		file_getchanneldifference_proto_rawDescData = protoimpl.X.CompressGZIP(file_getchanneldifference_proto_rawDescData)
	})
	return file_getchanneldifference_proto_rawDescData
}

var file_getchanneldifference_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_getchanneldifference_proto_goTypes = []interface{}{
	(*ReqGetChannelDifference)(nil), // 0: imchannelofflineserver.ReqGetChannelDifference
	(*ChanDiffStr)(nil),             // 1: imchannelofflineserver.ChanDiffStr
	(*RspGetChannelDifference)(nil), // 2: imchannelofflineserver.RspGetChannelDifference
	(*dialogserver.UserDialog)(nil), // 3: dialogserver.UserDialog
	(*pbcomm.Debug)(nil),            // 4: pbcomm.Debug
}
var file_getchanneldifference_proto_depIdxs = []int32{
	3, // 0: imchannelofflineserver.ChanDiffStr.userDialog:type_name -> dialogserver.UserDialog
	4, // 1: imchannelofflineserver.ChanDiffStr.debug:type_name -> pbcomm.Debug
	0, // 2: imchannelofflineserver.ChannelOfflineServerService.GetChannelDifference:input_type -> imchannelofflineserver.ReqGetChannelDifference
	0, // 3: imchannelofflineserver.ChannelOfflineServerService.GetChannelDifferenceIos:input_type -> imchannelofflineserver.ReqGetChannelDifference
	2, // 4: imchannelofflineserver.ChannelOfflineServerService.GetChannelDifference:output_type -> imchannelofflineserver.RspGetChannelDifference
	2, // 5: imchannelofflineserver.ChannelOfflineServerService.GetChannelDifferenceIos:output_type -> imchannelofflineserver.RspGetChannelDifference
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_getchanneldifference_proto_init() }
func file_getchanneldifference_proto_init() {
	if File_getchanneldifference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_getchanneldifference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetChannelDifference); i {
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
		file_getchanneldifference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChanDiffStr); i {
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
		file_getchanneldifference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspGetChannelDifference); i {
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
			RawDescriptor: file_getchanneldifference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_getchanneldifference_proto_goTypes,
		DependencyIndexes: file_getchanneldifference_proto_depIdxs,
		MessageInfos:      file_getchanneldifference_proto_msgTypes,
	}.Build()
	File_getchanneldifference_proto = out.File
	file_getchanneldifference_proto_rawDesc = nil
	file_getchanneldifference_proto_goTypes = nil
	file_getchanneldifference_proto_depIdxs = nil
}
