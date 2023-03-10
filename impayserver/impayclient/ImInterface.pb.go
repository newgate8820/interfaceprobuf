// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ImInterface.proto

package impayclient

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

// 我的发布 订单处理方式
type HandleType int32

const (
	HandleType_UP   HandleType = 0 //上架
	HandleType_DOWN HandleType = 1 //下架
	HandleType_DEL  HandleType = 2 //删除
)

// Enum value maps for HandleType.
var (
	HandleType_name = map[int32]string{
		0: "UP",
		1: "DOWN",
		2: "DEL",
	}
	HandleType_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
		"DEL":  2,
	}
)

func (x HandleType) Enum() *HandleType {
	p := new(HandleType)
	*p = x
	return p
}

func (x HandleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HandleType) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[0].Descriptor()
}

func (HandleType) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[0]
}

func (x HandleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HandleType.Descriptor instead.
func (HandleType) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{0}
}

// 交易类型
type TransType int32

const (
	TransType_BUY  TransType = 0 //购买
	TransType_SELL TransType = 1 //出售
)

// Enum value maps for TransType.
var (
	TransType_name = map[int32]string{
		0: "BUY",
		1: "SELL",
	}
	TransType_value = map[string]int32{
		"BUY":  0,
		"SELL": 1,
	}
)

func (x TransType) Enum() *TransType {
	p := new(TransType)
	*p = x
	return p
}

func (x TransType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransType) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[1].Descriptor()
}

func (TransType) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[1]
}

func (x TransType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransType.Descriptor instead.
func (TransType) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{1}
}

// 发布类型
type ReleaseType int32

const (
	ReleaseType_AGENT ReleaseType = 0 //代理
	ReleaseType_USER  ReleaseType = 1 //用户
)

// Enum value maps for ReleaseType.
var (
	ReleaseType_name = map[int32]string{
		0: "AGENT",
		1: "USER",
	}
	ReleaseType_value = map[string]int32{
		"AGENT": 0,
		"USER":  1,
	}
)

func (x ReleaseType) Enum() *ReleaseType {
	p := new(ReleaseType)
	*p = x
	return p
}

func (x ReleaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReleaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[2].Descriptor()
}

func (ReleaseType) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[2]
}

func (x ReleaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReleaseType.Descriptor instead.
func (ReleaseType) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{2}
}

// 支付方式
type PayStyle int32

const (
	PayStyle_UNKNOWN PayStyle = 0 //未知
	PayStyle_ALIPAY  PayStyle = 1 //支付宝
	PayStyle_WECHAT  PayStyle = 2 //微信
	PayStyle_UNION   PayStyle = 3 //银行卡
)

// Enum value maps for PayStyle.
var (
	PayStyle_name = map[int32]string{
		0: "UNKNOWN",
		1: "ALIPAY",
		2: "WECHAT",
		3: "UNION",
	}
	PayStyle_value = map[string]int32{
		"UNKNOWN": 0,
		"ALIPAY":  1,
		"WECHAT":  2,
		"UNION":   3,
	}
)

func (x PayStyle) Enum() *PayStyle {
	p := new(PayStyle)
	*p = x
	return p
}

func (x PayStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[3].Descriptor()
}

func (PayStyle) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[3]
}

func (x PayStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayStyle.Descriptor instead.
func (PayStyle) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{3}
}

// 排序方式
type OrderStyle int32

const (
	OrderStyle_ASC  OrderStyle = 0
	OrderStyle_DESC OrderStyle = 1
)

// Enum value maps for OrderStyle.
var (
	OrderStyle_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderStyle_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderStyle) Enum() *OrderStyle {
	p := new(OrderStyle)
	*p = x
	return p
}

func (x OrderStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[4].Descriptor()
}

func (OrderStyle) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[4]
}

func (x OrderStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStyle.Descriptor instead.
func (OrderStyle) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{4}
}

// 交易订单状态
type TransStatus int32

const (
	TransStatus_COMPLETE TransStatus = 0 //完成
	TransStatus_CANCEL   TransStatus = 1 //取消
	TransStatus_PAID     TransStatus = 2 // 已付款
	TransStatus_UNPAID   TransStatus = 3 // 未付款
)

// Enum value maps for TransStatus.
var (
	TransStatus_name = map[int32]string{
		0: "COMPLETE",
		1: "CANCEL",
		2: "PAID",
		3: "UNPAID",
	}
	TransStatus_value = map[string]int32{
		"COMPLETE": 0,
		"CANCEL":   1,
		"PAID":     2,
		"UNPAID":   3,
	}
)

func (x TransStatus) Enum() *TransStatus {
	p := new(TransStatus)
	*p = x
	return p
}

func (x TransStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ImInterface_proto_enumTypes[5].Descriptor()
}

func (TransStatus) Type() protoreflect.EnumType {
	return &file_ImInterface_proto_enumTypes[5]
}

func (x TransStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransStatus.Descriptor instead.
func (TransStatus) EnumDescriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{5}
}

type SaveIcoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileId   uint64 `protobuf:"varint,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileHash uint64 `protobuf:"varint,3,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	FileSize uint64 `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Data     string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveIcoReq) Reset() {
	*x = SaveIcoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveIcoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveIcoReq) ProtoMessage() {}

func (x *SaveIcoReq) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveIcoReq.ProtoReflect.Descriptor instead.
func (*SaveIcoReq) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{0}
}

func (x *SaveIcoReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SaveIcoReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *SaveIcoReq) GetFileHash() uint64 {
	if x != nil {
		return x.FileHash
	}
	return 0
}

func (x *SaveIcoReq) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *SaveIcoReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 创建账户
type CreateAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	RegTime int32  `protobuf:"varint,3,opt,name=reg_time,json=regTime,proto3" json:"reg_time,omitempty"`
	Data    string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateAccount) Reset() {
	*x = CreateAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccount) ProtoMessage() {}

func (x *CreateAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccount.ProtoReflect.Descriptor instead.
func (*CreateAccount) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccount) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateAccount) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateAccount) GetRegTime() int32 {
	if x != nil {
		return x.RegTime
	}
	return 0
}

func (x *CreateAccount) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 获取账户信息
type GetAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAccount) Reset() {
	*x = GetAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccount) ProtoMessage() {}

func (x *GetAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccount.ProtoReflect.Descriptor instead.
func (*GetAccount) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccount) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 公共请求（im负责传发）
type ImCommonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          //由im添加， 用于验证
	Info      string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`                             //前端请求参数
	MessageId int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` //由im添加， 用于验证
	ChannelId int32  `protobuf:"varint,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` //由im添加， 用于验证
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`                           //电话
}

func (x *ImCommonReq) Reset() {
	*x = ImCommonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImCommonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImCommonReq) ProtoMessage() {}

func (x *ImCommonReq) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImCommonReq.ProtoReflect.Descriptor instead.
func (*ImCommonReq) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{3}
}

func (x *ImCommonReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ImCommonReq) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *ImCommonReq) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ImCommonReq) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ImCommonReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type TransListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //由im添加， 用于验证
	Info   string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`                    //前端请求参数
}

func (x *TransListReq) Reset() {
	*x = TransListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransListReq) ProtoMessage() {}

func (x *TransListReq) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransListReq.ProtoReflect.Descriptor instead.
func (*TransListReq) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{4}
}

func (x *TransListReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TransListReq) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// 公共响应
type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message     string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data        string  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	UserId      []int32 `protobuf:"varint,4,rep,packed,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BusinessKey string  `protobuf:"bytes,5,opt,name=business_key,json=businessKey,proto3" json:"business_key,omitempty"` //返回前端业务key
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{5}
}

func (x *CommonResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CommonResp) GetUserId() []int32 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CommonResp) GetBusinessKey() string {
	if x != nil {
		return x.BusinessKey
	}
	return ""
}

type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Phone  string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"` //电话
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{6}
}

func (x *UserIdReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserIdReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 购买、出售 列表返回
type TransListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     []int32     `protobuf:"varint,1,rep,packed,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CommonResp *CommonResp `protobuf:"bytes,2,opt,name=common_resp,json=commonResp,proto3" json:"common_resp,omitempty"`
}

func (x *TransListResp) Reset() {
	*x = TransListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImInterface_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransListResp) ProtoMessage() {}

func (x *TransListResp) ProtoReflect() protoreflect.Message {
	mi := &file_ImInterface_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransListResp.ProtoReflect.Descriptor instead.
func (*TransListResp) Descriptor() ([]byte, []int) {
	return file_ImInterface_proto_rawDescGZIP(), []int{7}
}

func (x *TransListResp) GetUserId() []int32 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *TransListResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

var File_ImInterface_proto protoreflect.FileDescriptor

var file_ImInterface_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x49, 0x63, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x6d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x49, 0x6d, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x22, 0x39, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x62, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6d, 0x70,
	0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x2a,
	0x27, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a,
	0x02, 0x55, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x45, 0x4c, 0x10, 0x02, 0x2a, 0x1e, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x2a, 0x22, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x3a, 0x0a, 0x08,
	0x50, 0x61, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x49, 0x50, 0x41, 0x59, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x2a, 0x1f, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x2a, 0x3d, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x4e, 0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x32, 0xdd, 0x02, 0x0a, 0x0b, 0x49, 0x6d, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x70, 0x61,
	0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6d, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x69, 0x6d,
	0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x61, 0x76,
	0x65, 0x49, 0x63, 0x6f, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x63, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x69, 0x6d, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x51, 0x0a, 0x13, 0x69, 0x6d, 0x2e, 0x69,
	0x62, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x69, 0x6d, 0x42,
	0x09, 0x49, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x27, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d,
	0x70, 0x61, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6d, 0x70, 0x61, 0x79, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x69, 0x62, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ImInterface_proto_rawDescOnce sync.Once
	file_ImInterface_proto_rawDescData = file_ImInterface_proto_rawDesc
)

func file_ImInterface_proto_rawDescGZIP() []byte {
	file_ImInterface_proto_rawDescOnce.Do(func() {
		file_ImInterface_proto_rawDescData = protoimpl.X.CompressGZIP(file_ImInterface_proto_rawDescData)
	})
	return file_ImInterface_proto_rawDescData
}

var file_ImInterface_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_ImInterface_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ImInterface_proto_goTypes = []interface{}{
	(HandleType)(0),       // 0: impayclient.HandleType
	(TransType)(0),        // 1: impayclient.TransType
	(ReleaseType)(0),      // 2: impayclient.ReleaseType
	(PayStyle)(0),         // 3: impayclient.PayStyle
	(OrderStyle)(0),       // 4: impayclient.OrderStyle
	(TransStatus)(0),      // 5: impayclient.TransStatus
	(*SaveIcoReq)(nil),    // 6: impayclient.SaveIcoReq
	(*CreateAccount)(nil), // 7: impayclient.CreateAccount
	(*GetAccount)(nil),    // 8: impayclient.GetAccount
	(*ImCommonReq)(nil),   // 9: impayclient.ImCommonReq
	(*TransListReq)(nil),  // 10: impayclient.TransListReq
	(*CommonResp)(nil),    // 11: impayclient.CommonResp
	(*UserIdReq)(nil),     // 12: impayclient.UserIdReq
	(*TransListResp)(nil), // 13: impayclient.TransListResp
}
var file_ImInterface_proto_depIdxs = []int32{
	11, // 0: impayclient.TransListResp.common_resp:type_name -> impayclient.CommonResp
	7,  // 1: impayclient.ImInterface.createAccount:input_type -> impayclient.CreateAccount
	12, // 2: impayclient.ImInterface.getAccount:input_type -> impayclient.UserIdReq
	9,  // 3: impayclient.ImInterface.commonReq:input_type -> impayclient.ImCommonReq
	10, // 4: impayclient.ImInterface.transList:input_type -> impayclient.TransListReq
	6,  // 5: impayclient.ImInterface.saveIco:input_type -> impayclient.SaveIcoReq
	11, // 6: impayclient.ImInterface.createAccount:output_type -> impayclient.CommonResp
	11, // 7: impayclient.ImInterface.getAccount:output_type -> impayclient.CommonResp
	11, // 8: impayclient.ImInterface.commonReq:output_type -> impayclient.CommonResp
	13, // 9: impayclient.ImInterface.transList:output_type -> impayclient.TransListResp
	11, // 10: impayclient.ImInterface.saveIco:output_type -> impayclient.CommonResp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_ImInterface_proto_init() }
func file_ImInterface_proto_init() {
	if File_ImInterface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ImInterface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveIcoReq); i {
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
		file_ImInterface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccount); i {
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
		file_ImInterface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccount); i {
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
		file_ImInterface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImCommonReq); i {
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
		file_ImInterface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransListReq); i {
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
		file_ImInterface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_ImInterface_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdReq); i {
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
		file_ImInterface_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransListResp); i {
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
			RawDescriptor: file_ImInterface_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ImInterface_proto_goTypes,
		DependencyIndexes: file_ImInterface_proto_depIdxs,
		EnumInfos:         file_ImInterface_proto_enumTypes,
		MessageInfos:      file_ImInterface_proto_msgTypes,
	}.Build()
	File_ImInterface_proto = out.File
	file_ImInterface_proto_rawDesc = nil
	file_ImInterface_proto_goTypes = nil
	file_ImInterface_proto_depIdxs = nil
}
