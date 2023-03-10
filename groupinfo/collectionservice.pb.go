// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: collectionservice.proto

package groupinfo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GroupCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId     int32 `protobuf:"varint,1,opt,name=Chat_id,json=ChatId,proto3" json:"Chat_id,omitempty"`
	UserId     int32 `protobuf:"varint,2,opt,name=User_id,json=UserId,proto3" json:"User_id,omitempty"`
	IsSuper    bool  `protobuf:"varint,3,opt,name=Is_super,json=IsSuper,proto3" json:"Is_super,omitempty"`
	IsChannel  bool  `protobuf:"varint,4,opt,name=Is_channel,json=IsChannel,proto3" json:"Is_channel,omitempty"`
	AccessHash int64 `protobuf:"varint,5,opt,name=Access_hash,json=AccessHash,proto3" json:"Access_hash,omitempty"`
}

func (x *GroupCollection) Reset() {
	*x = GroupCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCollection) ProtoMessage() {}

func (x *GroupCollection) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCollection.ProtoReflect.Descriptor instead.
func (*GroupCollection) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{0}
}

func (x *GroupCollection) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GroupCollection) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GroupCollection) GetIsSuper() bool {
	if x != nil {
		return x.IsSuper
	}
	return false
}

func (x *GroupCollection) GetIsChannel() bool {
	if x != nil {
		return x.IsChannel
	}
	return false
}

func (x *GroupCollection) GetAccessHash() int64 {
	if x != nil {
		return x.AccessHash
	}
	return 0
}

type SetCollectionDialogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection   *GroupCollection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	IsCollection bool             `protobuf:"varint,2,opt,name=isCollection,proto3" json:"isCollection,omitempty"`
	Dlog         *pbcomm.Debug    `protobuf:"bytes,3,opt,name=Dlog,proto3" json:"Dlog,omitempty"`
}

func (x *SetCollectionDialogReq) Reset() {
	*x = SetCollectionDialogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCollectionDialogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCollectionDialogReq) ProtoMessage() {}

func (x *SetCollectionDialogReq) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCollectionDialogReq.ProtoReflect.Descriptor instead.
func (*SetCollectionDialogReq) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{1}
}

func (x *SetCollectionDialogReq) GetCollection() *GroupCollection {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *SetCollectionDialogReq) GetIsCollection() bool {
	if x != nil {
		return x.IsCollection
	}
	return false
}

func (x *SetCollectionDialogReq) GetDlog() *pbcomm.Debug {
	if x != nil {
		return x.Dlog
	}
	return nil
}

type SetCollectionDialogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Flags      int32  `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	Resultcode int32  `protobuf:"varint,3,opt,name=resultcode,proto3" json:"resultcode,omitempty"`
}

func (x *SetCollectionDialogResp) Reset() {
	*x = SetCollectionDialogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCollectionDialogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCollectionDialogResp) ProtoMessage() {}

func (x *SetCollectionDialogResp) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCollectionDialogResp.ProtoReflect.Descriptor instead.
func (*SetCollectionDialogResp) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{2}
}

func (x *SetCollectionDialogResp) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *SetCollectionDialogResp) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *SetCollectionDialogResp) GetResultcode() int32 {
	if x != nil {
		return x.Resultcode
	}
	return 0
}

type GetCollectionDialogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsChannel bool          `protobuf:"varint,1,opt,name=isChannel,proto3" json:"isChannel,omitempty"`
	UserId    int32         `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Dlog      *pbcomm.Debug `protobuf:"bytes,3,opt,name=Dlog,proto3" json:"Dlog,omitempty"`
}

func (x *GetCollectionDialogsReq) Reset() {
	*x = GetCollectionDialogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDialogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDialogsReq) ProtoMessage() {}

func (x *GetCollectionDialogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionDialogsReq.ProtoReflect.Descriptor instead.
func (*GetCollectionDialogsReq) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetCollectionDialogsReq) GetIsChannel() bool {
	if x != nil {
		return x.IsChannel
	}
	return false
}

func (x *GetCollectionDialogsReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCollectionDialogsReq) GetDlog() *pbcomm.Debug {
	if x != nil {
		return x.Dlog
	}
	return nil
}

type GetCollectionDialogsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     []*GroupCollection `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Resultcode int32              `protobuf:"varint,2,opt,name=resultcode,proto3" json:"resultcode,omitempty"`
}

func (x *GetCollectionDialogsResp) Reset() {
	*x = GetCollectionDialogsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDialogsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDialogsResp) ProtoMessage() {}

func (x *GetCollectionDialogsResp) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionDialogsResp.ProtoReflect.Descriptor instead.
func (*GetCollectionDialogsResp) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionDialogsResp) GetResult() []*GroupCollection {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetCollectionDialogsResp) GetResultcode() int32 {
	if x != nil {
		return x.Resultcode
	}
	return 0
}

type GetGroupCollectionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsChannel bool          `protobuf:"varint,1,opt,name=isChannel,proto3" json:"isChannel,omitempty"`
	UserId    int32         `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Hash      string        `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Dlog      *pbcomm.Debug `protobuf:"bytes,4,opt,name=Dlog,proto3" json:"Dlog,omitempty"`
}

func (x *GetGroupCollectionsReq) Reset() {
	*x = GetGroupCollectionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupCollectionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupCollectionsReq) ProtoMessage() {}

func (x *GetGroupCollectionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupCollectionsReq.ProtoReflect.Descriptor instead.
func (*GetGroupCollectionsReq) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{5}
}

func (x *GetGroupCollectionsReq) GetIsChannel() bool {
	if x != nil {
		return x.IsChannel
	}
	return false
}

func (x *GetGroupCollectionsReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetGroupCollectionsReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetGroupCollectionsReq) GetDlog() *pbcomm.Debug {
	if x != nil {
		return x.Dlog
	}
	return nil
}

type GetGroupCollectionsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     []*GroupCollection `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	IsModified bool               `protobuf:"varint,2,opt,name=isModified,proto3" json:"isModified,omitempty"`
	Hash       string             `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Resultcode int32              `protobuf:"varint,4,opt,name=resultcode,proto3" json:"resultcode,omitempty"`
}

func (x *GetGroupCollectionsResp) Reset() {
	*x = GetGroupCollectionsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupCollectionsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupCollectionsResp) ProtoMessage() {}

func (x *GetGroupCollectionsResp) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupCollectionsResp.ProtoReflect.Descriptor instead.
func (*GetGroupCollectionsResp) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupCollectionsResp) GetResult() []*GroupCollection {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetGroupCollectionsResp) GetIsModified() bool {
	if x != nil {
		return x.IsModified
	}
	return false
}

func (x *GetGroupCollectionsResp) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetGroupCollectionsResp) GetResultcode() int32 {
	if x != nil {
		return x.Resultcode
	}
	return 0
}

type GetDialogCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection *GroupCollection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Dlog       *pbcomm.Debug    `protobuf:"bytes,2,opt,name=Dlog,proto3" json:"Dlog,omitempty"`
}

func (x *GetDialogCollectionReq) Reset() {
	*x = GetDialogCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogCollectionReq) ProtoMessage() {}

func (x *GetDialogCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogCollectionReq.ProtoReflect.Descriptor instead.
func (*GetDialogCollectionReq) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{7}
}

func (x *GetDialogCollectionReq) GetCollection() *GroupCollection {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *GetDialogCollectionReq) GetDlog() *pbcomm.Debug {
	if x != nil {
		return x.Dlog
	}
	return nil
}

type GetDialogCollectionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     bool  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Resultcode int32 `protobuf:"varint,2,opt,name=resultcode,proto3" json:"resultcode,omitempty"`
}

func (x *GetDialogCollectionResp) Reset() {
	*x = GetDialogCollectionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogCollectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogCollectionResp) ProtoMessage() {}

func (x *GetDialogCollectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogCollectionResp.ProtoReflect.Descriptor instead.
func (*GetDialogCollectionResp) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{8}
}

func (x *GetDialogCollectionResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *GetDialogCollectionResp) GetResultcode() int32 {
	if x != nil {
		return x.Resultcode
	}
	return 0
}

type UpdateCollectionDialogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32         `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChatId    int32         `protobuf:"varint,2,opt,name=chatId,proto3" json:"chatId,omitempty"`
	ChannelId int32         `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Dlog      *pbcomm.Debug `protobuf:"bytes,4,opt,name=Dlog,proto3" json:"Dlog,omitempty"`
}

func (x *UpdateCollectionDialogReq) Reset() {
	*x = UpdateCollectionDialogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionDialogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionDialogReq) ProtoMessage() {}

func (x *UpdateCollectionDialogReq) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionDialogReq.ProtoReflect.Descriptor instead.
func (*UpdateCollectionDialogReq) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCollectionDialogReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateCollectionDialogReq) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *UpdateCollectionDialogReq) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateCollectionDialogReq) GetDlog() *pbcomm.Debug {
	if x != nil {
		return x.Dlog
	}
	return nil
}

type UpdateCollectionDialogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Resultcode int32 `protobuf:"varint,2,opt,name=resultcode,proto3" json:"resultcode,omitempty"`
}

func (x *UpdateCollectionDialogResp) Reset() {
	*x = UpdateCollectionDialogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionservice_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionDialogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionDialogResp) ProtoMessage() {}

func (x *UpdateCollectionDialogResp) ProtoReflect() protoreflect.Message {
	mi := &file_collectionservice_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionDialogResp.ProtoReflect.Descriptor instead.
func (*UpdateCollectionDialogResp) Descriptor() ([]byte, []int) {
	return file_collectionservice_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateCollectionDialogResp) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *UpdateCollectionDialogResp) GetResultcode() int32 {
	if x != nil {
		return x.Resultcode
	}
	return 0
}

var File_collectionservice_proto protoreflect.FileDescriptor

var file_collectionservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x70, 0x62, 0x63,
	0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x49, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x49, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x49, 0x73,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x73, 0x68, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x53,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x52, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x22, 0x67, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52,
	0x04, 0x44, 0x6c, 0x6f, 0x67, 0x22, 0x6e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x6c,
	0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x22, 0xa1, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x77, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x52, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x22, 0x51, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x6c, 0x6f,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x04, 0x44, 0x6c, 0x6f, 0x67, 0x22, 0x54, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x84, 0x04, 0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a,
	0x13, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x61, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x67, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x1a, 0x25, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collectionservice_proto_rawDescOnce sync.Once
	file_collectionservice_proto_rawDescData = file_collectionservice_proto_rawDesc
)

func file_collectionservice_proto_rawDescGZIP() []byte {
	file_collectionservice_proto_rawDescOnce.Do(func() {
		file_collectionservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_collectionservice_proto_rawDescData)
	})
	return file_collectionservice_proto_rawDescData
}

var file_collectionservice_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_collectionservice_proto_goTypes = []interface{}{
	(*GroupCollection)(nil),            // 0: groupinfo.GroupCollection
	(*SetCollectionDialogReq)(nil),     // 1: groupinfo.SetCollectionDialogReq
	(*SetCollectionDialogResp)(nil),    // 2: groupinfo.SetCollectionDialogResp
	(*GetCollectionDialogsReq)(nil),    // 3: groupinfo.GetCollectionDialogsReq
	(*GetCollectionDialogsResp)(nil),   // 4: groupinfo.GetCollectionDialogsResp
	(*GetGroupCollectionsReq)(nil),     // 5: groupinfo.GetGroupCollectionsReq
	(*GetGroupCollectionsResp)(nil),    // 6: groupinfo.GetGroupCollectionsResp
	(*GetDialogCollectionReq)(nil),     // 7: groupinfo.GetDialogCollectionReq
	(*GetDialogCollectionResp)(nil),    // 8: groupinfo.GetDialogCollectionResp
	(*UpdateCollectionDialogReq)(nil),  // 9: groupinfo.UpdateCollectionDialogReq
	(*UpdateCollectionDialogResp)(nil), // 10: groupinfo.UpdateCollectionDialogResp
	(*pbcomm.Debug)(nil),               // 11: pbcomm.Debug
}
var file_collectionservice_proto_depIdxs = []int32{
	0,  // 0: groupinfo.SetCollectionDialogReq.collection:type_name -> groupinfo.GroupCollection
	11, // 1: groupinfo.SetCollectionDialogReq.Dlog:type_name -> pbcomm.Debug
	11, // 2: groupinfo.GetCollectionDialogsReq.Dlog:type_name -> pbcomm.Debug
	0,  // 3: groupinfo.GetCollectionDialogsResp.result:type_name -> groupinfo.GroupCollection
	11, // 4: groupinfo.GetGroupCollectionsReq.Dlog:type_name -> pbcomm.Debug
	0,  // 5: groupinfo.GetGroupCollectionsResp.result:type_name -> groupinfo.GroupCollection
	0,  // 6: groupinfo.GetDialogCollectionReq.collection:type_name -> groupinfo.GroupCollection
	11, // 7: groupinfo.GetDialogCollectionReq.Dlog:type_name -> pbcomm.Debug
	11, // 8: groupinfo.UpdateCollectionDialogReq.Dlog:type_name -> pbcomm.Debug
	1,  // 9: groupinfo.GroupCollectionService.SetCollectionDialog:input_type -> groupinfo.SetCollectionDialogReq
	3,  // 10: groupinfo.GroupCollectionService.GetCollectionDialogs:input_type -> groupinfo.GetCollectionDialogsReq
	5,  // 11: groupinfo.GroupCollectionService.GetGroupCollections:input_type -> groupinfo.GetGroupCollectionsReq
	7,  // 12: groupinfo.GroupCollectionService.GetDialogCollection:input_type -> groupinfo.GetDialogCollectionReq
	9,  // 13: groupinfo.GroupCollectionService.UpdateCollectionDialog:input_type -> groupinfo.UpdateCollectionDialogReq
	2,  // 14: groupinfo.GroupCollectionService.SetCollectionDialog:output_type -> groupinfo.SetCollectionDialogResp
	4,  // 15: groupinfo.GroupCollectionService.GetCollectionDialogs:output_type -> groupinfo.GetCollectionDialogsResp
	6,  // 16: groupinfo.GroupCollectionService.GetGroupCollections:output_type -> groupinfo.GetGroupCollectionsResp
	8,  // 17: groupinfo.GroupCollectionService.GetDialogCollection:output_type -> groupinfo.GetDialogCollectionResp
	10, // 18: groupinfo.GroupCollectionService.UpdateCollectionDialog:output_type -> groupinfo.UpdateCollectionDialogResp
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_collectionservice_proto_init() }
func file_collectionservice_proto_init() {
	if File_collectionservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collectionservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCollection); i {
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
		file_collectionservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCollectionDialogReq); i {
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
		file_collectionservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCollectionDialogResp); i {
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
		file_collectionservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionDialogsReq); i {
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
		file_collectionservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionDialogsResp); i {
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
		file_collectionservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupCollectionsReq); i {
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
		file_collectionservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupCollectionsResp); i {
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
		file_collectionservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogCollectionReq); i {
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
		file_collectionservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogCollectionResp); i {
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
		file_collectionservice_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionDialogReq); i {
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
		file_collectionservice_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionDialogResp); i {
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
			RawDescriptor: file_collectionservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collectionservice_proto_goTypes,
		DependencyIndexes: file_collectionservice_proto_depIdxs,
		MessageInfos:      file_collectionservice_proto_msgTypes,
	}.Build()
	File_collectionservice_proto = out.File
	file_collectionservice_proto_rawDesc = nil
	file_collectionservice_proto_goTypes = nil
	file_collectionservice_proto_depIdxs = nil
}
