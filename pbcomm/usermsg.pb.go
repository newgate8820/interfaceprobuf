// protoc --gogofast_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: usermsg.proto

package pbcomm

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

// 用户消息
type UserMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64        `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                      // 唯一键
	UserId       int32        `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`              // 用户ID
	WithId       int32        `protobuf:"varint,3,opt,name=WithId,proto3" json:"WithId,omitempty"`              // 消息交互方
	WithIdType   int32        `protobuf:"varint,4,opt,name=WithIdType,proto3" json:"WithIdType,omitempty"`      // 消息交互方类型 1.user 2:chat 3:channel
	MessageType  int32        `protobuf:"varint,5,opt,name=MessageType,proto3" json:"MessageType,omitempty"`    // 消息类型
	Flags        int32        `protobuf:"varint,6,opt,name=Flags,proto3" json:"Flags,omitempty"`                // 消息标识
	Out          bool         `protobuf:"varint,7,opt,name=Out,proto3" json:"Out,omitempty"`                    // 是否为自己发送
	Mentioned    bool         `protobuf:"varint,8,opt,name=Mentioned,proto3" json:"Mentioned,omitempty"`        // 是否提及他人
	MediaUnread  bool         `protobuf:"varint,9,opt,name=MediaUnread,proto3" json:"MediaUnread,omitempty"`    // 媒体是否已读
	Silent       bool         `protobuf:"varint,10,opt,name=Silent,proto3" json:"Silent,omitempty"`             // 是否静音
	Post         bool         `protobuf:"varint,11,opt,name=Post,proto3" json:"Post,omitempty"`                 // 是否邮递
	GroupId      int64        `protobuf:"varint,12,opt,name=GroupId,proto3" json:"GroupId,omitempty"`           // 图片组ID
	MsgId        int32        `protobuf:"varint,13,opt,name=MsgId,proto3" json:"MsgId,omitempty"`               // 消息id
	ToId         int32        `protobuf:"varint,14,opt,name=ToId,proto3" json:"ToId,omitempty"`                 // 接收方id
	Pts          int32        `protobuf:"varint,15,opt,name=Pts,proto3" json:"Pts,omitempty"`                   // 消息的pts
	FwdFrom      []byte       `protobuf:"bytes,16,opt,name=FwdFrom,proto3" json:"FwdFrom,omitempty"`            // 引用
	ViaBotId     int32        `protobuf:"varint,17,opt,name=ViaBotId,proto3" json:"ViaBotId,omitempty"`         // 机器人id
	ReplyToMsgId int32        `protobuf:"varint,18,opt,name=ReplyToMsgId,proto3" json:"ReplyToMsgId,omitempty"` // 回复id
	Date         int32        `protobuf:"varint,19,opt,name=Date,proto3" json:"Date,omitempty"`                 // 时间
	Message      string       `protobuf:"bytes,20,opt,name=Message,proto3" json:"Message,omitempty"`            // 消息内容
	Media        int64        `protobuf:"varint,21,opt,name=Media,proto3" json:"Media,omitempty"`               // 媒体id
	MediaData    []byte       `protobuf:"bytes,22,opt,name=MediaData,proto3" json:"MediaData,omitempty"`        // 媒体二进制
	ReplyMarkup  []byte       `protobuf:"bytes,23,opt,name=ReplyMarkup,proto3" json:"ReplyMarkup,omitempty"`    // 小键盘
	Entities     [][]byte     `protobuf:"bytes,24,rep,name=Entities,proto3" json:"Entities,omitempty"`          // 高亮
	FromId       int32        `protobuf:"varint,25,opt,name=FromId,proto3" json:"FromId,omitempty"`             // 发送方
	Views        int32        `protobuf:"varint,26,opt,name=Views,proto3" json:"Views,omitempty"`               // 查看人数
	EditDate     int32        `protobuf:"varint,27,opt,name=EditDate,proto3" json:"EditDate,omitempty"`         // 编辑时间
	Action       []byte       `protobuf:"bytes,28,opt,name=Action,proto3" json:"Action,omitempty"`              // 行为
	RandomId     int64        `protobuf:"varint,29,opt,name=RandomId,proto3" json:"RandomId,omitempty"`         // 随机id
	Uuid         int64        `protobuf:"varint,30,opt,name=Uuid,proto3" json:"Uuid,omitempty"`                 // 消息全局唯一id
	InsertDate   int32        `protobuf:"varint,31,opt,name=InsertDate,proto3" json:"InsertDate,omitempty"`     // 插入时间
	UpdateDate   int32        `protobuf:"varint,32,opt,name=UpdateDate,proto3" json:"UpdateDate,omitempty"`     // 更新时间
	Encry        int32        `protobuf:"varint,33,opt,name=Encry,proto3" json:"Encry,omitempty"`               // 是否加密信息 0: 不加密 1：加密
	Combination  *Combination `protobuf:"bytes,34,opt,name=Combination,proto3" json:"Combination,omitempty"`    // 预留字段,生成的时候需要在结果地方追加 xorm:"json"
}

func (x *UserMsg) Reset() {
	*x = UserMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMsg) ProtoMessage() {}

func (x *UserMsg) ProtoReflect() protoreflect.Message {
	mi := &file_usermsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMsg.ProtoReflect.Descriptor instead.
func (*UserMsg) Descriptor() ([]byte, []int) {
	return file_usermsg_proto_rawDescGZIP(), []int{0}
}

func (x *UserMsg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserMsg) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserMsg) GetWithId() int32 {
	if x != nil {
		return x.WithId
	}
	return 0
}

func (x *UserMsg) GetWithIdType() int32 {
	if x != nil {
		return x.WithIdType
	}
	return 0
}

func (x *UserMsg) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *UserMsg) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *UserMsg) GetOut() bool {
	if x != nil {
		return x.Out
	}
	return false
}

func (x *UserMsg) GetMentioned() bool {
	if x != nil {
		return x.Mentioned
	}
	return false
}

func (x *UserMsg) GetMediaUnread() bool {
	if x != nil {
		return x.MediaUnread
	}
	return false
}

func (x *UserMsg) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

func (x *UserMsg) GetPost() bool {
	if x != nil {
		return x.Post
	}
	return false
}

func (x *UserMsg) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserMsg) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *UserMsg) GetToId() int32 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *UserMsg) GetPts() int32 {
	if x != nil {
		return x.Pts
	}
	return 0
}

func (x *UserMsg) GetFwdFrom() []byte {
	if x != nil {
		return x.FwdFrom
	}
	return nil
}

func (x *UserMsg) GetViaBotId() int32 {
	if x != nil {
		return x.ViaBotId
	}
	return 0
}

func (x *UserMsg) GetReplyToMsgId() int32 {
	if x != nil {
		return x.ReplyToMsgId
	}
	return 0
}

func (x *UserMsg) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *UserMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserMsg) GetMedia() int64 {
	if x != nil {
		return x.Media
	}
	return 0
}

func (x *UserMsg) GetMediaData() []byte {
	if x != nil {
		return x.MediaData
	}
	return nil
}

func (x *UserMsg) GetReplyMarkup() []byte {
	if x != nil {
		return x.ReplyMarkup
	}
	return nil
}

func (x *UserMsg) GetEntities() [][]byte {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *UserMsg) GetFromId() int32 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *UserMsg) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *UserMsg) GetEditDate() int32 {
	if x != nil {
		return x.EditDate
	}
	return 0
}

func (x *UserMsg) GetAction() []byte {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *UserMsg) GetRandomId() int64 {
	if x != nil {
		return x.RandomId
	}
	return 0
}

func (x *UserMsg) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *UserMsg) GetInsertDate() int32 {
	if x != nil {
		return x.InsertDate
	}
	return 0
}

func (x *UserMsg) GetUpdateDate() int32 {
	if x != nil {
		return x.UpdateDate
	}
	return 0
}

func (x *UserMsg) GetEncry() int32 {
	if x != nil {
		return x.Encry
	}
	return 0
}

func (x *UserMsg) GetCombination() *Combination {
	if x != nil {
		return x.Combination
	}
	return nil
}

type MentionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *MentionStatus) Reset() {
	*x = MentionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MentionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MentionStatus) ProtoMessage() {}

func (x *MentionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_usermsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MentionStatus.ProtoReflect.Descriptor instead.
func (*MentionStatus) Descriptor() ([]byte, []int) {
	return file_usermsg_proto_rawDescGZIP(), []int{1}
}

func (x *MentionStatus) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MentionStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Combination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MentionedSlice []*MentionStatus `protobuf:"bytes,1,rep,name=MentionedSlice,proto3" json:"MentionedSlice,omitempty"` // @人 二进制
	KeyId          []int64          `protobuf:"varint,2,rep,packed,name=KeyId,proto3" json:"KeyId,omitempty"`           // keyId 判断
}

func (x *Combination) Reset() {
	*x = Combination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Combination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Combination) ProtoMessage() {}

func (x *Combination) ProtoReflect() protoreflect.Message {
	mi := &file_usermsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Combination.ProtoReflect.Descriptor instead.
func (*Combination) Descriptor() ([]byte, []int) {
	return file_usermsg_proto_rawDescGZIP(), []int{2}
}

func (x *Combination) GetMentionedSlice() []*MentionStatus {
	if x != nil {
		return x.MentionedSlice
	}
	return nil
}

func (x *Combination) GetKeyId() []int64 {
	if x != nil {
		return x.KeyId
	}
	return nil
}

var File_usermsg_proto protoreflect.FileDescriptor

var file_usermsg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x22, 0x8e, 0x07, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x57, 0x69, 0x74,
	0x68, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4f,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53,
	0x69, 0x6c, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x49,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x50, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x74, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x46, 0x77, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x46, 0x77, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x61,
	0x42, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x56, 0x69, 0x61,
	0x42, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x72, 0x6f,
	0x6d, 0x49, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x72, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x69, 0x65, 0x77, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x45, 0x64, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x42, 0x18, 0x5a,
	0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usermsg_proto_rawDescOnce sync.Once
	file_usermsg_proto_rawDescData = file_usermsg_proto_rawDesc
)

func file_usermsg_proto_rawDescGZIP() []byte {
	file_usermsg_proto_rawDescOnce.Do(func() {
		file_usermsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_usermsg_proto_rawDescData)
	})
	return file_usermsg_proto_rawDescData
}

var file_usermsg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_usermsg_proto_goTypes = []interface{}{
	(*UserMsg)(nil),       // 0: pbcomm.UserMsg
	(*MentionStatus)(nil), // 1: pbcomm.MentionStatus
	(*Combination)(nil),   // 2: pbcomm.Combination
}
var file_usermsg_proto_depIdxs = []int32{
	2, // 0: pbcomm.UserMsg.Combination:type_name -> pbcomm.Combination
	1, // 1: pbcomm.Combination.MentionedSlice:type_name -> pbcomm.MentionStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_usermsg_proto_init() }
func file_usermsg_proto_init() {
	if File_usermsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usermsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMsg); i {
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
		file_usermsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MentionStatus); i {
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
		file_usermsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Combination); i {
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
			RawDescriptor: file_usermsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_usermsg_proto_goTypes,
		DependencyIndexes: file_usermsg_proto_depIdxs,
		MessageInfos:      file_usermsg_proto_msgTypes,
	}.Build()
	File_usermsg_proto = out.File
	file_usermsg_proto_rawDesc = nil
	file_usermsg_proto_goTypes = nil
	file_usermsg_proto_depIdxs = nil
}
