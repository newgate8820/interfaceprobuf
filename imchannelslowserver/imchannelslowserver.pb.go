// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imchannelslowserver.proto

//protoc --gogofast_out=plugins=grpc:. imchannelslowserver.proto

package imchannelslowserver

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

// 检查慢速模式
type CheckIsSlowModelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChannelId int32 `protobuf:"varint,2,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"` //超级群ID
	SendTime  int32 `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"`   //发送时间
}

func (x *CheckIsSlowModelReq) Reset() {
	*x = CheckIsSlowModelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imchannelslowserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsSlowModelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsSlowModelReq) ProtoMessage() {}

func (x *CheckIsSlowModelReq) ProtoReflect() protoreflect.Message {
	mi := &file_imchannelslowserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsSlowModelReq.ProtoReflect.Descriptor instead.
func (*CheckIsSlowModelReq) Descriptor() ([]byte, []int) {
	return file_imchannelslowserver_proto_rawDescGZIP(), []int{0}
}

func (x *CheckIsSlowModelReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckIsSlowModelReq) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *CheckIsSlowModelReq) GetSendTime() int32 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

// 慢速模式返回,能不能发送消息
type CheckIsSlowModelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanSendMessage bool  `protobuf:"varint,1,opt,name=canSendMessage,proto3" json:"canSendMessage,omitempty"` //是否能够发送消息
	NextTime       int32 `protobuf:"varint,2,opt,name=nextTime,proto3" json:"nextTime,omitempty"`             //下一次发送还剩下时长
}

func (x *CheckIsSlowModelResp) Reset() {
	*x = CheckIsSlowModelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imchannelslowserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsSlowModelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsSlowModelResp) ProtoMessage() {}

func (x *CheckIsSlowModelResp) ProtoReflect() protoreflect.Message {
	mi := &file_imchannelslowserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsSlowModelResp.ProtoReflect.Descriptor instead.
func (*CheckIsSlowModelResp) Descriptor() ([]byte, []int) {
	return file_imchannelslowserver_proto_rawDescGZIP(), []int{1}
}

func (x *CheckIsSlowModelResp) GetCanSendMessage() bool {
	if x != nil {
		return x.CanSendMessage
	}
	return false
}

func (x *CheckIsSlowModelResp) GetNextTime() int32 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

// 设置或者修改群是否为慢速模式
type UpdateChannelSlowModuleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpId           int32 `protobuf:"varint,1,opt,name=OpId,proto3" json:"OpId,omitempty"` //操作者的ID
	KeyId          int64 `protobuf:"varint,2,opt,name=keyId,proto3" json:"keyId,omitempty"`
	ChannelId      int32 `protobuf:"varint,4,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`           //超级群的ID
	IsSlowModule   bool  `protobuf:"varint,5,opt,name=IsSlowModule,proto3" json:"IsSlowModule,omitempty"`     // 是否慢速模式
	PeerSlowSecond int32 `protobuf:"varint,6,opt,name=PeerSlowSecond,proto3" json:"PeerSlowSecond,omitempty"` // 再慢速模式下
}

func (x *UpdateChannelSlowModuleReq) Reset() {
	*x = UpdateChannelSlowModuleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imchannelslowserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelSlowModuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelSlowModuleReq) ProtoMessage() {}

func (x *UpdateChannelSlowModuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_imchannelslowserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelSlowModuleReq.ProtoReflect.Descriptor instead.
func (*UpdateChannelSlowModuleReq) Descriptor() ([]byte, []int) {
	return file_imchannelslowserver_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateChannelSlowModuleReq) GetOpId() int32 {
	if x != nil {
		return x.OpId
	}
	return 0
}

func (x *UpdateChannelSlowModuleReq) GetKeyId() int64 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *UpdateChannelSlowModuleReq) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateChannelSlowModuleReq) GetIsSlowModule() bool {
	if x != nil {
		return x.IsSlowModule
	}
	return false
}

func (x *UpdateChannelSlowModuleReq) GetPeerSlowSecond() int32 {
	if x != nil {
		return x.PeerSlowSecond
	}
	return 0
}

type UpdateChannelSlowModuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateChannelSlowModuleResp) Reset() {
	*x = UpdateChannelSlowModuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imchannelslowserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelSlowModuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelSlowModuleResp) ProtoMessage() {}

func (x *UpdateChannelSlowModuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_imchannelslowserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelSlowModuleResp.ProtoReflect.Descriptor instead.
func (*UpdateChannelSlowModuleResp) Descriptor() ([]byte, []int) {
	return file_imchannelslowserver_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateChannelSlowModuleResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_imchannelslowserver_proto protoreflect.FileDescriptor

var file_imchannelslowserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6d, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x67, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x1a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x4f, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x6c,
	0x6f, 0x77, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x35, 0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x86, 0x02, 0x0a, 0x19, 0x69, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x53, 0x6c, 0x6f, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a,
	0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6c,
	0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6c, 0x6f, 0x77,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x30, 0x2e, 0x69, 0x6d, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6c, 0x6f,
	0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x69, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x28, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x53,
	0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x69, 0x6d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x6c, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imchannelslowserver_proto_rawDescOnce sync.Once
	file_imchannelslowserver_proto_rawDescData = file_imchannelslowserver_proto_rawDesc
)

func file_imchannelslowserver_proto_rawDescGZIP() []byte {
	file_imchannelslowserver_proto_rawDescOnce.Do(func() {
		file_imchannelslowserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_imchannelslowserver_proto_rawDescData)
	})
	return file_imchannelslowserver_proto_rawDescData
}

var file_imchannelslowserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imchannelslowserver_proto_goTypes = []interface{}{
	(*CheckIsSlowModelReq)(nil),         // 0: imchannelslowserver.CheckIsSlowModelReq
	(*CheckIsSlowModelResp)(nil),        // 1: imchannelslowserver.CheckIsSlowModelResp
	(*UpdateChannelSlowModuleReq)(nil),  // 2: imchannelslowserver.updateChannelSlowModuleReq
	(*UpdateChannelSlowModuleResp)(nil), // 3: imchannelslowserver.updateChannelSlowModuleResp
}
var file_imchannelslowserver_proto_depIdxs = []int32{
	2, // 0: imchannelslowserver.imChanneSlowServerService.updateChannelSlowModule:input_type -> imchannelslowserver.updateChannelSlowModuleReq
	0, // 1: imchannelslowserver.imChanneSlowServerService.CheckIsSlowModel:input_type -> imchannelslowserver.CheckIsSlowModelReq
	3, // 2: imchannelslowserver.imChanneSlowServerService.updateChannelSlowModule:output_type -> imchannelslowserver.updateChannelSlowModuleResp
	1, // 3: imchannelslowserver.imChanneSlowServerService.CheckIsSlowModel:output_type -> imchannelslowserver.CheckIsSlowModelResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imchannelslowserver_proto_init() }
func file_imchannelslowserver_proto_init() {
	if File_imchannelslowserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imchannelslowserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsSlowModelReq); i {
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
		file_imchannelslowserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsSlowModelResp); i {
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
		file_imchannelslowserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelSlowModuleReq); i {
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
		file_imchannelslowserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelSlowModuleResp); i {
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
			RawDescriptor: file_imchannelslowserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imchannelslowserver_proto_goTypes,
		DependencyIndexes: file_imchannelslowserver_proto_depIdxs,
		MessageInfos:      file_imchannelslowserver_proto_msgTypes,
	}.Build()
	File_imchannelslowserver_proto = out.File
	file_imchannelslowserver_proto_rawDesc = nil
	file_imchannelslowserver_proto_goTypes = nil
	file_imchannelslowserver_proto_depIdxs = nil
}
