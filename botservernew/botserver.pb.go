// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitlab.chatserver.im/interfaceprobuf/botserver/botserver.proto

/*
Package pbbotserver is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/interfaceprobuf/botserver/botserver.proto

It has these top-level messages:

	ReqGetAdminGroups
	ReqGetGroups
	Peer
	ReplyGetGroups
	ReqDeleteMessage
	ReqSendChatAction
	ReqData
	ReplyData
	ReqGetChatMembersCount
	ReplyGetChatMembersCount
*/
package pbbotserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pbcomm "gitlab.chatserver.im/interfaceprobuf/pbcomm"
import pbadvertisemonitor "gitlab.chatserver.im/interfaceprobuf/advertisemonitornew"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChatType 聊天类型
type ChatType int32

const (
	ChatType_ChatTypeZero ChatType = 0
	ChatType_PeerUser     ChatType = 1
	ChatType_PeerChat     ChatType = 2
	ChatType_ChannelChat  ChatType = 3
)

var ChatType_name = map[int32]string{
	0: "ChatTypeZero",
	1: "PeerUser",
	2: "PeerChat",
	3: "ChannelChat",
}
var ChatType_value = map[string]int32{
	"ChatTypeZero": 0,
	"PeerUser":     1,
	"PeerChat":     2,
	"ChannelChat":  3,
}

func (x ChatType) String() string {
	return proto.EnumName(ChatType_name, int32(x))
}
func (ChatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReqGetAdminGroups struct {
	UserID int32 `protobuf:"varint,1,opt,name=UserID,json=userID" json:"UserID,omitempty"`
	// int32 PeerFilter = 2; // 2：只要普通群 4：只要超级群 8：只要频道 （可相互组合 ex: 2|4|8 暂时不支持单聊)
	Debug *pbcomm.Debug `protobuf:"bytes,2,opt,name=debug" json:"debug,omitempty"`
}

func (m *ReqGetAdminGroups) Reset()                    { *m = ReqGetAdminGroups{} }
func (m *ReqGetAdminGroups) String() string            { return proto.CompactTextString(m) }
func (*ReqGetAdminGroups) ProtoMessage()               {}
func (*ReqGetAdminGroups) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqGetAdminGroups) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ReqGetAdminGroups) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type ReqGetGroups struct {
	UserID     int32 `protobuf:"varint,1,opt,name=UserID,json=userID" json:"UserID,omitempty"`
	PeerFilter int32 `protobuf:"varint,2,opt,name=PeerFilter,json=peerFilter" json:"PeerFilter,omitempty"`
}

func (m *ReqGetGroups) Reset()                    { *m = ReqGetGroups{} }
func (m *ReqGetGroups) String() string            { return proto.CompactTextString(m) }
func (*ReqGetGroups) ProtoMessage()               {}
func (*ReqGetGroups) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReqGetGroups) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ReqGetGroups) GetPeerFilter() int32 {
	if m != nil {
		return m.PeerFilter
	}
	return 0
}

type Peer struct {
	PeerID   int32  `protobuf:"varint,1,opt,name=PeerID,json=peerID" json:"PeerID,omitempty"`
	PeerName string `protobuf:"bytes,2,opt,name=PeerName,json=peerName" json:"PeerName,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Peer) GetPeerID() int32 {
	if m != nil {
		return m.PeerID
	}
	return 0
}

func (m *Peer) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type ReplyGetGroups struct {
	Groups      []*Peer `protobuf:"bytes,1,rep,name=Groups,json=groups" json:"Groups,omitempty"`
	SuperGroups []*Peer `protobuf:"bytes,2,rep,name=SuperGroups,json=superGroups" json:"SuperGroups,omitempty"`
	Channels    []*Peer `protobuf:"bytes,3,rep,name=Channels,json=channels" json:"Channels,omitempty"`
}

func (m *ReplyGetGroups) Reset()                    { *m = ReplyGetGroups{} }
func (m *ReplyGetGroups) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetGroups) ProtoMessage()               {}
func (*ReplyGetGroups) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReplyGetGroups) GetGroups() []*Peer {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ReplyGetGroups) GetSuperGroups() []*Peer {
	if m != nil {
		return m.SuperGroups
	}
	return nil
}

func (m *ReplyGetGroups) GetChannels() []*Peer {
	if m != nil {
		return m.Channels
	}
	return nil
}

type ReqDeleteMessage struct {
	ChatType ChatType      `protobuf:"varint,1,opt,name=ChatType,json=chatType,enum=pbbotserver.ChatType" json:"ChatType,omitempty"`
	UserId   int32         `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ChatId   int32         `protobuf:"varint,3,opt,name=chat_id,json=chatId" json:"chat_id,omitempty"`
	MsgId    []int32       `protobuf:"varint,4,rep,packed,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	Debug    *pbcomm.Debug `protobuf:"bytes,5,opt,name=debug" json:"debug,omitempty"`
}

func (m *ReqDeleteMessage) Reset()                    { *m = ReqDeleteMessage{} }
func (m *ReqDeleteMessage) String() string            { return proto.CompactTextString(m) }
func (*ReqDeleteMessage) ProtoMessage()               {}
func (*ReqDeleteMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqDeleteMessage) GetChatType() ChatType {
	if m != nil {
		return m.ChatType
	}
	return ChatType_ChatTypeZero
}

func (m *ReqDeleteMessage) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ReqDeleteMessage) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *ReqDeleteMessage) GetMsgId() []int32 {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *ReqDeleteMessage) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type ReqSendChatAction struct {
	ChatType   ChatType      `protobuf:"varint,1,opt,name=ChatType,json=chatType,enum=pbbotserver.ChatType" json:"ChatType,omitempty"`
	UserId     int32         `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ChatId     int32         `protobuf:"varint,3,opt,name=chat_id,json=chatId" json:"chat_id,omitempty"`
	ActionData []byte        `protobuf:"bytes,4,opt,name=actionData,proto3" json:"actionData,omitempty"`
	Debug      *pbcomm.Debug `protobuf:"bytes,5,opt,name=debug" json:"debug,omitempty"`
}

func (m *ReqSendChatAction) Reset()                    { *m = ReqSendChatAction{} }
func (m *ReqSendChatAction) String() string            { return proto.CompactTextString(m) }
func (*ReqSendChatAction) ProtoMessage()               {}
func (*ReqSendChatAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqSendChatAction) GetChatType() ChatType {
	if m != nil {
		return m.ChatType
	}
	return ChatType_ChatTypeZero
}

func (m *ReqSendChatAction) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ReqSendChatAction) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *ReqSendChatAction) GetActionData() []byte {
	if m != nil {
		return m.ActionData
	}
	return nil
}

func (m *ReqSendChatAction) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type ReqData struct {
	ChatType ChatType `protobuf:"varint,1,opt,name=ChatType,json=chatType,enum=pbbotserver.ChatType" json:"ChatType,omitempty"`
	ReqData  []byte   `protobuf:"bytes,2,opt,name=reqData,proto3" json:"reqData,omitempty"`
	BotId    int32    `protobuf:"varint,3,opt,name=bot_id,json=botId" json:"bot_id,omitempty"`
}

func (m *ReqData) Reset()                    { *m = ReqData{} }
func (m *ReqData) String() string            { return proto.CompactTextString(m) }
func (*ReqData) ProtoMessage()               {}
func (*ReqData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqData) GetChatType() ChatType {
	if m != nil {
		return m.ChatType
	}
	return ChatType_ChatTypeZero
}

func (m *ReqData) GetReqData() []byte {
	if m != nil {
		return m.ReqData
	}
	return nil
}

func (m *ReqData) GetBotId() int32 {
	if m != nil {
		return m.BotId
	}
	return 0
}

type ReplyData struct {
	Reply []byte `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Error int32  `protobuf:"varint,2,opt,name=error" json:"error,omitempty"`
}

func (m *ReplyData) Reset()                    { *m = ReplyData{} }
func (m *ReplyData) String() string            { return proto.CompactTextString(m) }
func (*ReplyData) ProtoMessage()               {}
func (*ReplyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReplyData) GetReply() []byte {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *ReplyData) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

type ReqGetChatMembersCount struct {
	ChatType ChatType      `protobuf:"varint,1,opt,name=ChatType,json=chatType,enum=pbbotserver.ChatType" json:"ChatType,omitempty"`
	UserId   int32         `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ChatId   int32         `protobuf:"varint,3,opt,name=chat_id,json=chatId" json:"chat_id,omitempty"`
	Debug    *pbcomm.Debug `protobuf:"bytes,4,opt,name=debug" json:"debug,omitempty"`
}

func (m *ReqGetChatMembersCount) Reset()                    { *m = ReqGetChatMembersCount{} }
func (m *ReqGetChatMembersCount) String() string            { return proto.CompactTextString(m) }
func (*ReqGetChatMembersCount) ProtoMessage()               {}
func (*ReqGetChatMembersCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReqGetChatMembersCount) GetChatType() ChatType {
	if m != nil {
		return m.ChatType
	}
	return ChatType_ChatTypeZero
}

func (m *ReqGetChatMembersCount) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ReqGetChatMembersCount) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *ReqGetChatMembersCount) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type ReplyGetChatMembersCount struct {
	ChatMembersCount int32 `protobuf:"varint,1,opt,name=ChatMembersCount,json=chatMembersCount" json:"ChatMembersCount,omitempty"`
	ErrorCode        int32 `protobuf:"varint,2,opt,name=ErrorCode,json=errorCode" json:"ErrorCode,omitempty"`
}

func (m *ReplyGetChatMembersCount) Reset()                    { *m = ReplyGetChatMembersCount{} }
func (m *ReplyGetChatMembersCount) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetChatMembersCount) ProtoMessage()               {}
func (*ReplyGetChatMembersCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReplyGetChatMembersCount) GetChatMembersCount() int32 {
	if m != nil {
		return m.ChatMembersCount
	}
	return 0
}

func (m *ReplyGetChatMembersCount) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqGetAdminGroups)(nil), "pbbotserver.ReqGetAdminGroups")
	proto.RegisterType((*ReqGetGroups)(nil), "pbbotserver.ReqGetGroups")
	proto.RegisterType((*Peer)(nil), "pbbotserver.Peer")
	proto.RegisterType((*ReplyGetGroups)(nil), "pbbotserver.ReplyGetGroups")
	proto.RegisterType((*ReqDeleteMessage)(nil), "pbbotserver.ReqDeleteMessage")
	proto.RegisterType((*ReqSendChatAction)(nil), "pbbotserver.ReqSendChatAction")
	proto.RegisterType((*ReqData)(nil), "pbbotserver.ReqData")
	proto.RegisterType((*ReplyData)(nil), "pbbotserver.ReplyData")
	proto.RegisterType((*ReqGetChatMembersCount)(nil), "pbbotserver.ReqGetChatMembersCount")
	proto.RegisterType((*ReplyGetChatMembersCount)(nil), "pbbotserver.ReplyGetChatMembersCount")
	proto.RegisterEnum("pbbotserver.ChatType", ChatType_name, ChatType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BotServer service

type BotServerClient interface {
	// 離開群組
	LeaveChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 獲取群訊息
	GetChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 獲取群成員數目
	GetChatMembersCount(ctx context.Context, in *ReqGetChatMembersCount, opts ...grpc.CallOption) (*ReplyGetChatMembersCount, error)
	// 獲取群管理員
	GetChatAdministrators(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 修改群名字
	SetChatTitle(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 修改群簡介
	SetChatDescription(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 置顶群消息
	UpdatePinnedChannelMessage(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	//    // 取消置顶群消息
	//    rpc UpinChatMessage (ReqData) returns (ReplyData) {
	//    }
	// 發送Action
	SendChatAction(ctx context.Context, in *ReqSendChatAction, opts ...grpc.CallOption) (*ReplyData, error)
	// 發送名片
	SendContact(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error)
	// 刪除消息
	DeleteMessage(ctx context.Context, in *ReqDeleteMessage, opts ...grpc.CallOption) (*ReplyData, error)
	// GetGroups 获取群组消息
	GetGroups(ctx context.Context, in *ReqGetGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error)
	// GetGroups 获取使用者是管理員的群組消息
	GetAdminGroups(ctx context.Context, in *ReqGetAdminGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error)
	// 新增手動禁言名單
	SetManualBlockList(ctx context.Context, in *pbadvertisemonitor.ReqSetManualBlockList, opts ...grpc.CallOption) (*pbadvertisemonitor.ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(ctx context.Context, in *pbadvertisemonitor.ReqDelManualBlockList, opts ...grpc.CallOption) (*pbadvertisemonitor.ReplyData, error)
}

type botServerClient struct {
	cc *grpc.ClientConn
}

func NewBotServerClient(cc *grpc.ClientConn) BotServerClient {
	return &botServerClient{cc}
}

func (c *botServerClient) LeaveChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/LeaveChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChat(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/GetChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChatMembersCount(ctx context.Context, in *ReqGetChatMembersCount, opts ...grpc.CallOption) (*ReplyGetChatMembersCount, error) {
	out := new(ReplyGetChatMembersCount)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/GetChatMembersCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetChatAdministrators(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/GetChatAdministrators", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetChatTitle(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/SetChatTitle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetChatDescription(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/SetChatDescription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) UpdatePinnedChannelMessage(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/UpdatePinnedChannelMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SendChatAction(ctx context.Context, in *ReqSendChatAction, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/SendChatAction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SendContact(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/SendContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) DeleteMessage(ctx context.Context, in *ReqDeleteMessage, opts ...grpc.CallOption) (*ReplyData, error) {
	out := new(ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/DeleteMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetGroups(ctx context.Context, in *ReqGetGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error) {
	out := new(ReplyGetGroups)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/GetGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) GetAdminGroups(ctx context.Context, in *ReqGetAdminGroups, opts ...grpc.CallOption) (*ReplyGetGroups, error) {
	out := new(ReplyGetGroups)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/GetAdminGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) SetManualBlockList(ctx context.Context, in *pbadvertisemonitor.ReqSetManualBlockList, opts ...grpc.CallOption) (*pbadvertisemonitor.ReplyData, error) {
	out := new(pbadvertisemonitor.ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/SetManualBlockList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServerClient) DelManualBlockList(ctx context.Context, in *pbadvertisemonitor.ReqDelManualBlockList, opts ...grpc.CallOption) (*pbadvertisemonitor.ReplyData, error) {
	out := new(pbadvertisemonitor.ReplyData)
	err := grpc.Invoke(ctx, "/pbbotserver.BotServer/DelManualBlockList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BotServer service

type BotServerServer interface {
	// 離開群組
	LeaveChat(context.Context, *ReqData) (*ReplyData, error)
	// 獲取群訊息
	GetChat(context.Context, *ReqData) (*ReplyData, error)
	// 獲取群成員數目
	GetChatMembersCount(context.Context, *ReqGetChatMembersCount) (*ReplyGetChatMembersCount, error)
	// 獲取群管理員
	GetChatAdministrators(context.Context, *ReqData) (*ReplyData, error)
	// 修改群名字
	SetChatTitle(context.Context, *ReqData) (*ReplyData, error)
	// 修改群簡介
	SetChatDescription(context.Context, *ReqData) (*ReplyData, error)
	// 置顶群消息
	UpdatePinnedChannelMessage(context.Context, *ReqData) (*ReplyData, error)
	//    // 取消置顶群消息
	//    rpc UpinChatMessage (ReqData) returns (ReplyData) {
	//    }
	// 發送Action
	SendChatAction(context.Context, *ReqSendChatAction) (*ReplyData, error)
	// 發送名片
	SendContact(context.Context, *ReqData) (*ReplyData, error)
	// 刪除消息
	DeleteMessage(context.Context, *ReqDeleteMessage) (*ReplyData, error)
	// GetGroups 获取群组消息
	GetGroups(context.Context, *ReqGetGroups) (*ReplyGetGroups, error)
	// GetGroups 获取使用者是管理員的群組消息
	GetAdminGroups(context.Context, *ReqGetAdminGroups) (*ReplyGetGroups, error)
	// 新增手動禁言名單
	SetManualBlockList(context.Context, *pbadvertisemonitor.ReqSetManualBlockList) (*pbadvertisemonitor.ReplyData, error)
	// 刪除手動禁言名單
	DelManualBlockList(context.Context, *pbadvertisemonitor.ReqDelManualBlockList) (*pbadvertisemonitor.ReplyData, error)
}

func RegisterBotServerServer(s *grpc.Server, srv BotServerServer) {
	s.RegisterService(&_BotServer_serviceDesc, srv)
}

func _BotServer_LeaveChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).LeaveChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/LeaveChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).LeaveChat(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChat(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChatMembersCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetChatMembersCount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChatMembersCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChatMembersCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChatMembersCount(ctx, req.(*ReqGetChatMembersCount))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetChatAdministrators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetChatAdministrators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetChatAdministrators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetChatAdministrators(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetChatTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetChatTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetChatTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetChatTitle(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetChatDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetChatDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetChatDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetChatDescription(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_UpdatePinnedChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).UpdatePinnedChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/UpdatePinnedChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).UpdatePinnedChannelMessage(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SendChatAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSendChatAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SendChatAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SendChatAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SendChatAction(ctx, req.(*ReqSendChatAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SendContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SendContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SendContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SendContact(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/DeleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).DeleteMessage(ctx, req.(*ReqDeleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetGroups(ctx, req.(*ReqGetGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_GetAdminGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetAdminGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).GetAdminGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/GetAdminGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).GetAdminGroups(ctx, req.(*ReqGetAdminGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_SetManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbadvertisemonitor.ReqSetManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).SetManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/SetManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).SetManualBlockList(ctx, req.(*pbadvertisemonitor.ReqSetManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotServer_DelManualBlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbadvertisemonitor.ReqDelManualBlockList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).DelManualBlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbbotserver.BotServer/DelManualBlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).DelManualBlockList(ctx, req.(*pbadvertisemonitor.ReqDelManualBlockList))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbbotserver.BotServer",
	HandlerType: (*BotServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaveChat",
			Handler:    _BotServer_LeaveChat_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _BotServer_GetChat_Handler,
		},
		{
			MethodName: "GetChatMembersCount",
			Handler:    _BotServer_GetChatMembersCount_Handler,
		},
		{
			MethodName: "GetChatAdministrators",
			Handler:    _BotServer_GetChatAdministrators_Handler,
		},
		{
			MethodName: "SetChatTitle",
			Handler:    _BotServer_SetChatTitle_Handler,
		},
		{
			MethodName: "SetChatDescription",
			Handler:    _BotServer_SetChatDescription_Handler,
		},
		{
			MethodName: "UpdatePinnedChannelMessage",
			Handler:    _BotServer_UpdatePinnedChannelMessage_Handler,
		},
		{
			MethodName: "SendChatAction",
			Handler:    _BotServer_SendChatAction_Handler,
		},
		{
			MethodName: "SendContact",
			Handler:    _BotServer_SendContact_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _BotServer_DeleteMessage_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _BotServer_GetGroups_Handler,
		},
		{
			MethodName: "GetAdminGroups",
			Handler:    _BotServer_GetAdminGroups_Handler,
		},
		{
			MethodName: "SetManualBlockList",
			Handler:    _BotServer_SetManualBlockList_Handler,
		},
		{
			MethodName: "DelManualBlockList",
			Handler:    _BotServer_DelManualBlockList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/interfaceprobuf/botserver/botserver.proto",
}

func init() {
	proto.RegisterFile("gitlab.chatserver.im/interfaceprobuf/botserver/botserver.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x2d, 0x53, 0x96, 0x46, 0x8a, 0xab, 0x6c, 0x63, 0x57, 0x65, 0x1b, 0x43, 0x60, 0x50,
	0x40, 0x09, 0x50, 0x05, 0x55, 0x0e, 0xfd, 0x43, 0x03, 0x24, 0x56, 0xe2, 0xaa, 0xb0, 0x1b, 0x83,
	0x4a, 0x2e, 0xbd, 0x34, 0x4b, 0x72, 0xa2, 0x10, 0x25, 0xb9, 0xcc, 0x72, 0xe5, 0x22, 0x4f, 0x53,
	0xa0, 0xaf, 0x50, 0xa0, 0xb7, 0xbe, 0x5b, 0xb1, 0x3f, 0x34, 0x25, 0x4a, 0x91, 0xa5, 0x1c, 0x7c,
	0xe2, 0xce, 0x7c, 0xdf, 0x7c, 0x3b, 0x3b, 0x3b, 0xbb, 0x5c, 0x78, 0x3c, 0x8d, 0x44, 0x4c, 0xfd,
	0x41, 0xf0, 0x96, 0x8a, 0x1c, 0xf9, 0x25, 0xf2, 0x41, 0x94, 0x3c, 0x8c, 0x52, 0x81, 0xfc, 0x0d,
	0x0d, 0x30, 0xe3, 0xcc, 0x9f, 0xbd, 0x79, 0xe8, 0x33, 0x03, 0x96, 0xa3, 0x41, 0xc6, 0x99, 0x60,
	0xa4, 0x95, 0xf9, 0x57, 0x2e, 0xe7, 0xbb, 0x8d, 0xc4, 0x32, 0x3f, 0x60, 0x49, 0x62, 0x3e, 0x5a,
	0xc6, 0x79, 0xb1, 0x51, 0x24, 0x0d, 0x2f, 0x91, 0x8b, 0x28, 0xc7, 0x84, 0xa5, 0x91, 0x60, 0x3c,
	0xc5, 0x3f, 0x97, 0x7c, 0x5a, 0xd0, 0xbd, 0x80, 0xdb, 0x1e, 0xbe, 0x3b, 0x45, 0xf1, 0x24, 0x4c,
	0xa2, 0xf4, 0x94, 0xb3, 0x59, 0x96, 0x93, 0x23, 0xa8, 0xbf, 0xca, 0x91, 0x8f, 0x47, 0x5d, 0xab,
	0x67, 0xf5, 0x6d, 0xaf, 0x3e, 0x53, 0x16, 0xb9, 0x07, 0x76, 0x88, 0xfe, 0x6c, 0xda, 0xdd, 0xed,
	0x59, 0xfd, 0xd6, 0xf0, 0xd6, 0xc0, 0xe4, 0x36, 0x92, 0x4e, 0x4f, 0x63, 0xee, 0x73, 0x68, 0x6b,
	0xc5, 0x6b, 0xc4, 0x8e, 0x01, 0x2e, 0x10, 0xf9, 0xf3, 0x28, 0x16, 0xc8, 0x95, 0xa2, 0xed, 0x41,
	0x76, 0xe5, 0x71, 0x7f, 0x80, 0x3d, 0x89, 0xcb, 0x78, 0xf9, 0x2d, 0xe3, 0x33, 0x65, 0x11, 0x07,
	0x1a, 0xd2, 0xff, 0x2b, 0x4d, 0x50, 0x45, 0x37, 0xbd, 0x46, 0x66, 0x6c, 0xf7, 0x2f, 0x0b, 0x0e,
	0x3c, 0xcc, 0xe2, 0xf7, 0x65, 0x1a, 0xf7, 0xa1, 0xae, 0x47, 0x5d, 0xab, 0x57, 0xeb, 0xb7, 0x86,
	0xb7, 0x07, 0x73, 0x3b, 0x32, 0x90, 0x4a, 0x5e, 0x7d, 0xaa, 0xa9, 0x8f, 0xa0, 0x35, 0x99, 0x65,
	0xc8, 0x0d, 0x7f, 0xf7, 0x43, 0xfc, 0x56, 0x5e, 0xb2, 0xc8, 0xd7, 0xd0, 0x38, 0x79, 0x4b, 0xd3,
	0x14, 0xe3, 0xbc, 0x5b, 0xfb, 0x50, 0x44, 0x23, 0x30, 0x14, 0xf7, 0x1f, 0x0b, 0x3a, 0x1e, 0xbe,
	0x1b, 0x61, 0x8c, 0x02, 0xcf, 0x31, 0xcf, 0xe9, 0x14, 0xc9, 0x37, 0x4a, 0x43, 0xbc, 0x7c, 0x9f,
	0xa1, 0x5a, 0xec, 0xc1, 0xf0, 0x70, 0x41, 0xa3, 0x00, 0x95, 0x8e, 0x1a, 0x91, 0xcf, 0x60, 0x5f,
	0xd6, 0xf3, 0xf7, 0x28, 0x34, 0x25, 0xd4, 0xe5, 0x0d, 0x25, 0x20, 0x49, 0x12, 0xa8, 0x69, 0x40,
	0x9a, 0xe3, 0x90, 0x1c, 0x42, 0x3d, 0xc9, 0xa7, 0xd2, 0xbf, 0xd7, 0xab, 0xf5, 0x6d, 0xcf, 0x4e,
	0xf2, 0xe9, 0x38, 0x2c, 0xf7, 0xd6, 0x5e, 0xb3, 0xb7, 0xff, 0x59, 0xaa, 0x5d, 0x26, 0x98, 0x86,
	0x32, 0x97, 0x27, 0x81, 0x88, 0x58, 0x7a, 0x33, 0x69, 0x1f, 0x03, 0x50, 0x35, 0xdd, 0x88, 0x0a,
	0xda, 0xdd, 0xeb, 0x59, 0xfd, 0xb6, 0x37, 0xe7, 0xd9, 0x2c, 0xff, 0x04, 0xf6, 0x65, 0xd1, 0x25,
	0xff, 0x23, 0x92, 0xee, 0xc2, 0x3e, 0xd7, 0xd1, 0x2a, 0xe9, 0xb6, 0x57, 0x98, 0xb2, 0xa6, 0x3e,
	0x9b, 0x4b, 0xda, 0xf6, 0x99, 0x18, 0x87, 0xee, 0xb7, 0xd0, 0x54, 0x5d, 0xa8, 0x38, 0x77, 0xc0,
	0xe6, 0xd2, 0x50, 0xb3, 0xb5, 0x3d, 0x6d, 0x48, 0x2f, 0x72, 0xce, 0x8a, 0x03, 0xa0, 0x0d, 0xf7,
	0x6f, 0x0b, 0x8e, 0xf4, 0x21, 0x92, 0x69, 0x9c, 0x63, 0xe2, 0x23, 0xcf, 0x4f, 0xd8, 0x2c, 0x15,
	0x37, 0x53, 0xec, 0xab, 0x62, 0xee, 0xad, 0x29, 0x66, 0x08, 0xdd, 0xe2, 0x8c, 0x2d, 0x65, 0xf9,
	0x00, 0x3a, 0x55, 0x9f, 0x39, 0xbe, 0x9d, 0xa0, 0xca, 0xfd, 0x12, 0x9a, 0xcf, 0xe4, 0xaa, 0x4f,
	0x58, 0x88, 0x26, 0xc1, 0x26, 0x16, 0x8e, 0x07, 0xe3, 0x72, 0xbd, 0xa4, 0x03, 0xed, 0x62, 0xfc,
	0x1b, 0x72, 0xd6, 0xd9, 0x21, 0x6d, 0x7d, 0x09, 0xc8, 0x0b, 0xa6, 0x63, 0x15, 0x96, 0xe4, 0x74,
	0x76, 0xc9, 0x27, 0xd0, 0x32, 0x27, 0x52, 0x39, 0x6a, 0xc3, 0x7f, 0x1b, 0xd0, 0x7c, 0xca, 0xc4,
	0x44, 0x15, 0x8a, 0xfc, 0x08, 0xcd, 0x33, 0xa4, 0x97, 0x28, 0x41, 0x72, 0x67, 0xa1, 0x86, 0xa6,
	0x47, 0x9c, 0xa3, 0x8a, 0xd7, 0x6c, 0xa5, 0xbb, 0x43, 0xbe, 0x87, 0x7d, 0xb3, 0xec, 0xad, 0x43,
	0x03, 0xf8, 0x74, 0x55, 0xc5, 0xee, 0x55, 0x65, 0x56, 0x90, 0x9c, 0xaf, 0x96, 0x55, 0x57, 0xd0,
	0xdc, 0x1d, 0x72, 0x0a, 0x87, 0x06, 0x50, 0xf7, 0x7a, 0x94, 0x0b, 0x4e, 0x05, 0xe3, 0xf9, 0xd6,
	0xd9, 0x3e, 0x86, 0xf6, 0x44, 0x0b, 0xbd, 0x8c, 0x44, 0x8c, 0x5b, 0xc7, 0x8f, 0x80, 0x98, 0xf8,
	0x11, 0xe6, 0x01, 0x8f, 0x32, 0x75, 0x63, 0x6c, 0xab, 0x72, 0x06, 0xce, 0xab, 0x2c, 0xa4, 0x02,
	0x2f, 0xa2, 0x34, 0xc5, 0xd0, 0x6c, 0x6b, 0x71, 0x6d, 0x6e, 0xab, 0xf6, 0x0b, 0x1c, 0x54, 0x6e,
	0xb0, 0xe3, 0xaa, 0xc2, 0x22, 0xbe, 0x46, 0xeb, 0x27, 0x68, 0x29, 0x2e, 0x4b, 0x05, 0x0d, 0xb6,
	0x6f, 0x86, 0x9f, 0xe1, 0xd6, 0xe2, 0x2f, 0xe0, 0xee, 0x92, 0xc0, 0x3c, 0xbc, 0x46, 0xe9, 0x19,
	0x34, 0xcb, 0x9f, 0xdd, 0xe7, 0x2b, 0x9a, 0x49, 0x43, 0xce, 0x17, 0x2b, 0x5b, 0x48, 0x83, 0xee,
	0x0e, 0x79, 0x01, 0x07, 0x95, 0xc7, 0xc0, 0xf1, 0x0a, 0xad, 0x39, 0xfc, 0x3a, 0xc1, 0xd7, 0xaa,
	0x01, 0xce, 0x69, 0x3a, 0xa3, 0xf1, 0xd3, 0x98, 0x05, 0x7f, 0x9c, 0x45, 0xb9, 0x20, 0xf7, 0x07,
	0x99, 0xbf, 0xf4, 0x22, 0x51, 0x75, 0xaf, 0x52, 0x9d, 0xbb, 0xab, 0xa9, 0xe5, 0xca, 0x5f, 0x03,
	0x19, 0x61, 0xbc, 0xe9, 0x0c, 0xcb, 0xd4, 0x6b, 0x67, 0xf0, 0xeb, 0xea, 0xad, 0xf4, 0xe8, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x87, 0x4a, 0x86, 0x05, 0x0a, 0x00, 0x00,
}
