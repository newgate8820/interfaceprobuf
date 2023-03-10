//protoc --gogofast_out=plugins=grpc:. gitlab.chatserver.im/interfaceprobuf/activeuserserver/activeuserserver.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: activeuserinfoserver.proto

package groupservice

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

// 激活用户
type ActivateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyId  int64 `protobuf:"varint,1,opt,name=KeyId,proto3" json:"KeyId,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *ActivateUserReq) Reset() {
	*x = ActivateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserReq) ProtoMessage() {}

func (x *ActivateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserReq.ProtoReflect.Descriptor instead.
func (*ActivateUserReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateUserReq) GetKeyId() int64 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *ActivateUserReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ActivateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ActiveErrorCode `protobuf:"varint,1,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *ActivateUserReply) Reset() {
	*x = ActivateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserReply) ProtoMessage() {}

func (x *ActivateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserReply.ProtoReflect.Descriptor instead.
func (*ActivateUserReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{1}
}

func (x *ActivateUserReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

// 根据群名称搜索自己所在的群
type SearchSelfGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *SearchSelfGroupReq) Reset() {
	*x = SearchSelfGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSelfGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSelfGroupReq) ProtoMessage() {}

func (x *SearchSelfGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSelfGroupReq.ProtoReflect.Descriptor instead.
func (*SearchSelfGroupReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{2}
}

func (x *SearchSelfGroupReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchSelfGroupReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SearchSelfGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatIds    []int32         `protobuf:"varint,1,rep,packed,name=ChatIds,proto3" json:"ChatIds,omitempty"`       //普通群Id列表
	ChannelIds []int32         `protobuf:"varint,2,rep,packed,name=ChannelIds,proto3" json:"ChannelIds,omitempty"` //超级群、频道Id列表
	ErrorCode  ActiveErrorCode `protobuf:"varint,3,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *SearchSelfGroupReply) Reset() {
	*x = SearchSelfGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSelfGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSelfGroupReply) ProtoMessage() {}

func (x *SearchSelfGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSelfGroupReply.ProtoReflect.Descriptor instead.
func (*SearchSelfGroupReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{3}
}

func (x *SearchSelfGroupReply) GetChatIds() []int32 {
	if x != nil {
		return x.ChatIds
	}
	return nil
}

func (x *SearchSelfGroupReply) GetChannelIds() []int32 {
	if x != nil {
		return x.ChannelIds
	}
	return nil
}

func (x *SearchSelfGroupReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

// 获取两个用户相同的群数量
type GetTwoUserChatsCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserOne int32 `protobuf:"varint,1,opt,name=UserOne,proto3" json:"UserOne,omitempty"`
	UserTwo int32 `protobuf:"varint,2,opt,name=UserTwo,proto3" json:"UserTwo,omitempty"`
}

func (x *GetTwoUserChatsCountReq) Reset() {
	*x = GetTwoUserChatsCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTwoUserChatsCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTwoUserChatsCountReq) ProtoMessage() {}

func (x *GetTwoUserChatsCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTwoUserChatsCountReq.ProtoReflect.Descriptor instead.
func (*GetTwoUserChatsCountReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{4}
}

func (x *GetTwoUserChatsCountReq) GetUserOne() int32 {
	if x != nil {
		return x.UserOne
	}
	return 0
}

func (x *GetTwoUserChatsCountReq) GetUserTwo() int32 {
	if x != nil {
		return x.UserTwo
	}
	return 0
}

type GetTwoUserChatsCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count     int32           `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	ErrorCode ActiveErrorCode `protobuf:"varint,2,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *GetTwoUserChatsCountReply) Reset() {
	*x = GetTwoUserChatsCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTwoUserChatsCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTwoUserChatsCountReply) ProtoMessage() {}

func (x *GetTwoUserChatsCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTwoUserChatsCountReply.ProtoReflect.Descriptor instead.
func (*GetTwoUserChatsCountReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{5}
}

func (x *GetTwoUserChatsCountReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetTwoUserChatsCountReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

// 获取两个用户相同的群
type GetTwoUserChatsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserOne int32 `protobuf:"varint,1,opt,name=UserOne,proto3" json:"UserOne,omitempty"`
	UserTwo int32 `protobuf:"varint,2,opt,name=UserTwo,proto3" json:"UserTwo,omitempty"`
}

func (x *GetTwoUserChatsReq) Reset() {
	*x = GetTwoUserChatsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTwoUserChatsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTwoUserChatsReq) ProtoMessage() {}

func (x *GetTwoUserChatsReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTwoUserChatsReq.ProtoReflect.Descriptor instead.
func (*GetTwoUserChatsReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{6}
}

func (x *GetTwoUserChatsReq) GetUserOne() int32 {
	if x != nil {
		return x.UserOne
	}
	return 0
}

func (x *GetTwoUserChatsReq) GetUserTwo() int32 {
	if x != nil {
		return x.UserTwo
	}
	return 0
}

type GetTwoUserChatsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatIds    []int32         `protobuf:"varint,1,rep,packed,name=ChatIds,proto3" json:"ChatIds,omitempty"`       //普通群Id列表
	ChannelIds []int32         `protobuf:"varint,2,rep,packed,name=ChannelIds,proto3" json:"ChannelIds,omitempty"` //超级群、频道Id列表
	ErrorCode  ActiveErrorCode `protobuf:"varint,3,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *GetTwoUserChatsReply) Reset() {
	*x = GetTwoUserChatsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTwoUserChatsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTwoUserChatsReply) ProtoMessage() {}

func (x *GetTwoUserChatsReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTwoUserChatsReply.ProtoReflect.Descriptor instead.
func (*GetTwoUserChatsReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{7}
}

func (x *GetTwoUserChatsReply) GetChatIds() []int32 {
	if x != nil {
		return x.ChatIds
	}
	return nil
}

func (x *GetTwoUserChatsReply) GetChannelIds() []int32 {
	if x != nil {
		return x.ChannelIds
	}
	return nil
}

func (x *GetTwoUserChatsReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

// 判断哪些用户在线
type CheckOnlineUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int32 `protobuf:"varint,1,rep,packed,name=UserIds,proto3" json:"UserIds,omitempty"`
}

func (x *CheckOnlineUsersReq) Reset() {
	*x = CheckOnlineUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOnlineUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOnlineUsersReq) ProtoMessage() {}

func (x *CheckOnlineUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOnlineUsersReq.ProtoReflect.Descriptor instead.
func (*CheckOnlineUsersReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{8}
}

func (x *CheckOnlineUsersReq) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type CheckOnlineUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds   []int32         `protobuf:"varint,1,rep,packed,name=UserIds,proto3" json:"UserIds,omitempty"` //在线用户
	ErrorCode ActiveErrorCode `protobuf:"varint,2,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *CheckOnlineUsersReply) Reset() {
	*x = CheckOnlineUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOnlineUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOnlineUsersReply) ProtoMessage() {}

func (x *CheckOnlineUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOnlineUsersReply.ProtoReflect.Descriptor instead.
func (*CheckOnlineUsersReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{9}
}

func (x *CheckOnlineUsersReply) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *CheckOnlineUsersReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

// 根据传入ID搜索用户
type SearchUsersByStrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int32 `protobuf:"varint,1,rep,packed,name=UserIds,proto3" json:"UserIds,omitempty"` //传入的群成员ID
	Str     string  `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`                 //需要匹配的字符串
}

func (x *SearchUsersByStrReq) Reset() {
	*x = SearchUsersByStrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersByStrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersByStrReq) ProtoMessage() {}

func (x *SearchUsersByStrReq) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersByStrReq.ProtoReflect.Descriptor instead.
func (*SearchUsersByStrReq) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{10}
}

func (x *SearchUsersByStrReq) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *SearchUsersByStrReq) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type SearchUsersByStrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds   []int32         `protobuf:"varint,1,rep,packed,name=UserIds,proto3" json:"UserIds,omitempty"` //匹配到的用户ID
	ErrorCode ActiveErrorCode `protobuf:"varint,2,opt,name=ErrorCode,proto3,enum=groupservice.ActiveErrorCode" json:"ErrorCode,omitempty"`
}

func (x *SearchUsersByStrReply) Reset() {
	*x = SearchUsersByStrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activeuserinfoserver_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersByStrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersByStrReply) ProtoMessage() {}

func (x *SearchUsersByStrReply) ProtoReflect() protoreflect.Message {
	mi := &file_activeuserinfoserver_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersByStrReply.ProtoReflect.Descriptor instead.
func (*SearchUsersByStrReply) Descriptor() ([]byte, []int) {
	return file_activeuserinfoserver_proto_rawDescGZIP(), []int{11}
}

func (x *SearchUsersByStrReply) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *SearchUsersByStrReply) GetErrorCode() ActiveErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ActiveErrorCode_Active_OK
}

var File_activeuserinfoserver_proto protoreflect.FileDescriptor

var file_activeuserinfoserver_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x12, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f,
	0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x50, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x6c, 0x66, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x65, 0x6c, 0x66, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x07, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x77, 0x6f, 0x22, 0x6e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x22, 0x8d,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2f,
	0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22,
	0x6e, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x41, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x53, 0x74, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x22, 0x6e, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x53, 0x74, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x77, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x77, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x77, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x6c,
	0x66, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x6c, 0x66,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65,
	0x6c, 0x66, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x10,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x74, 0x72,
	0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x74, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x53, 0x74, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_activeuserinfoserver_proto_rawDescOnce sync.Once
	file_activeuserinfoserver_proto_rawDescData = file_activeuserinfoserver_proto_rawDesc
)

func file_activeuserinfoserver_proto_rawDescGZIP() []byte {
	file_activeuserinfoserver_proto_rawDescOnce.Do(func() {
		file_activeuserinfoserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_activeuserinfoserver_proto_rawDescData)
	})
	return file_activeuserinfoserver_proto_rawDescData
}

var file_activeuserinfoserver_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_activeuserinfoserver_proto_goTypes = []interface{}{
	(*ActivateUserReq)(nil),           // 0: groupservice.ActivateUserReq
	(*ActivateUserReply)(nil),         // 1: groupservice.ActivateUserReply
	(*SearchSelfGroupReq)(nil),        // 2: groupservice.SearchSelfGroupReq
	(*SearchSelfGroupReply)(nil),      // 3: groupservice.SearchSelfGroupReply
	(*GetTwoUserChatsCountReq)(nil),   // 4: groupservice.GetTwoUserChatsCountReq
	(*GetTwoUserChatsCountReply)(nil), // 5: groupservice.GetTwoUserChatsCountReply
	(*GetTwoUserChatsReq)(nil),        // 6: groupservice.GetTwoUserChatsReq
	(*GetTwoUserChatsReply)(nil),      // 7: groupservice.GetTwoUserChatsReply
	(*CheckOnlineUsersReq)(nil),       // 8: groupservice.CheckOnlineUsersReq
	(*CheckOnlineUsersReply)(nil),     // 9: groupservice.CheckOnlineUsersReply
	(*SearchUsersByStrReq)(nil),       // 10: groupservice.SearchUsersByStrReq
	(*SearchUsersByStrReply)(nil),     // 11: groupservice.SearchUsersByStrReply
	(ActiveErrorCode)(0),              // 12: groupservice.ActiveErrorCode
}
var file_activeuserinfoserver_proto_depIdxs = []int32{
	12, // 0: groupservice.ActivateUserReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	12, // 1: groupservice.SearchSelfGroupReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	12, // 2: groupservice.GetTwoUserChatsCountReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	12, // 3: groupservice.GetTwoUserChatsReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	12, // 4: groupservice.CheckOnlineUsersReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	12, // 5: groupservice.SearchUsersByStrReply.ErrorCode:type_name -> groupservice.ActiveErrorCode
	0,  // 6: groupservice.ActiveUserInfoService.ActivateUser:input_type -> groupservice.ActivateUserReq
	4,  // 7: groupservice.ActiveUserInfoService.GetTwoUserChatsCount:input_type -> groupservice.GetTwoUserChatsCountReq
	6,  // 8: groupservice.ActiveUserInfoService.GetTwoUserChats:input_type -> groupservice.GetTwoUserChatsReq
	2,  // 9: groupservice.ActiveUserInfoService.SearchSelfGroup:input_type -> groupservice.SearchSelfGroupReq
	8,  // 10: groupservice.ActiveUserInfoService.CheckOnlineUsers:input_type -> groupservice.CheckOnlineUsersReq
	10, // 11: groupservice.ActiveUserInfoService.SearchUsersByStr:input_type -> groupservice.SearchUsersByStrReq
	1,  // 12: groupservice.ActiveUserInfoService.ActivateUser:output_type -> groupservice.ActivateUserReply
	5,  // 13: groupservice.ActiveUserInfoService.GetTwoUserChatsCount:output_type -> groupservice.GetTwoUserChatsCountReply
	7,  // 14: groupservice.ActiveUserInfoService.GetTwoUserChats:output_type -> groupservice.GetTwoUserChatsReply
	3,  // 15: groupservice.ActiveUserInfoService.SearchSelfGroup:output_type -> groupservice.SearchSelfGroupReply
	9,  // 16: groupservice.ActiveUserInfoService.CheckOnlineUsers:output_type -> groupservice.CheckOnlineUsersReply
	11, // 17: groupservice.ActiveUserInfoService.SearchUsersByStr:output_type -> groupservice.SearchUsersByStrReply
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_activeuserinfoserver_proto_init() }
func file_activeuserinfoserver_proto_init() {
	if File_activeuserinfoserver_proto != nil {
		return
	}
	file_groupservice_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_activeuserinfoserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserReply); i {
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
		file_activeuserinfoserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSelfGroupReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSelfGroupReply); i {
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
		file_activeuserinfoserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTwoUserChatsCountReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTwoUserChatsCountReply); i {
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
		file_activeuserinfoserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTwoUserChatsReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTwoUserChatsReply); i {
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
		file_activeuserinfoserver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOnlineUsersReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOnlineUsersReply); i {
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
		file_activeuserinfoserver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUsersByStrReq); i {
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
		file_activeuserinfoserver_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUsersByStrReply); i {
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
			RawDescriptor: file_activeuserinfoserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activeuserinfoserver_proto_goTypes,
		DependencyIndexes: file_activeuserinfoserver_proto_depIdxs,
		MessageInfos:      file_activeuserinfoserver_proto_msgTypes,
	}.Build()
	File_activeuserinfoserver_proto = out.File
	file_activeuserinfoserver_proto_rawDesc = nil
	file_activeuserinfoserver_proto_goTypes = nil
	file_activeuserinfoserver_proto_depIdxs = nil
}
