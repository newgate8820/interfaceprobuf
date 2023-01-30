// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitlab.chatserver.im/im/iminterfaceprotobufs/botserver/updateservice/updateservice.proto

/*
Package updateservice is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/im/iminterfaceprotobufs/botserver/updateservice/updateservice.proto

It has these top-level messages:

	GetOldBotsRequest
	GetOldBotsReply
	PushMessageListRequest
	PushMessageRequest
	PushMessageReply
	CallbackQueryRequest
	CallbackQueryReply
	GetInlineQueryRequest
	GetInlineQueryReply
	SendInlineQueryRequest
	SendInlineQueryReply
	PushInlineQueryRequest
	PushInlineQueryRequestV2
*/
package updateservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pbcomm "gitlab.chatserver.im/interfaceprobuf/pbcomm"

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

// 获取旧机器人请求
type GetOldBotsRequest struct {
}

func (m *GetOldBotsRequest) Reset()                    { *m = GetOldBotsRequest{} }
func (m *GetOldBotsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOldBotsRequest) ProtoMessage()               {}
func (*GetOldBotsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 获取旧机器人响应
type GetOldBotsReply struct {
	Bots []int32 `protobuf:"varint,1,rep,packed,name=bots" json:"bots,omitempty"`
}

func (m *GetOldBotsReply) Reset()                    { *m = GetOldBotsReply{} }
func (m *GetOldBotsReply) String() string            { return proto.CompactTextString(m) }
func (*GetOldBotsReply) ProtoMessage()               {}
func (*GetOldBotsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetOldBotsReply) GetBots() []int32 {
	if m != nil {
		return m.Bots
	}
	return nil
}

// 推送消息请求
type PushMessageListRequest struct {
	BotId              []uint32 `protobuf:"varint,1,rep,packed,name=bot_id,json=botId" json:"bot_id,omitempty"`
	UserData           []byte   `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	BrforeMirageChatid uint32   `protobuf:"varint,3,opt,name=brforeMirageChatid" json:"brforeMirageChatid,omitempty"`
}

func (m *PushMessageListRequest) Reset()                    { *m = PushMessageListRequest{} }
func (m *PushMessageListRequest) String() string            { return proto.CompactTextString(m) }
func (*PushMessageListRequest) ProtoMessage()               {}
func (*PushMessageListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushMessageListRequest) GetBotId() []uint32 {
	if m != nil {
		return m.BotId
	}
	return nil
}

func (m *PushMessageListRequest) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

func (m *PushMessageListRequest) GetBrforeMirageChatid() uint32 {
	if m != nil {
		return m.BrforeMirageChatid
	}
	return 0
}

// 推送消息请求
type PushMessageRequest struct {
	BotId              uint32 `protobuf:"varint,1,opt,name=bot_id,json=botId" json:"bot_id,omitempty"`
	UserData           []byte `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	BrforeMirageChatid uint32 `protobuf:"varint,3,opt,name=brforeMirageChatid" json:"brforeMirageChatid,omitempty"`
	Lang               string `protobuf:"bytes,4,opt,name=lang" json:"lang,omitempty"`
}

func (m *PushMessageRequest) Reset()                    { *m = PushMessageRequest{} }
func (m *PushMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*PushMessageRequest) ProtoMessage()               {}
func (*PushMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PushMessageRequest) GetBotId() uint32 {
	if m != nil {
		return m.BotId
	}
	return 0
}

func (m *PushMessageRequest) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

func (m *PushMessageRequest) GetBrforeMirageChatid() uint32 {
	if m != nil {
		return m.BrforeMirageChatid
	}
	return 0
}

func (m *PushMessageRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

// 推送消息响应
type PushMessageReply struct {
}

func (m *PushMessageReply) Reset()                    { *m = PushMessageReply{} }
func (m *PushMessageReply) String() string            { return proto.CompactTextString(m) }
func (*PushMessageReply) ProtoMessage()               {}
func (*PushMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 推送回调查询请求
type CallbackQueryRequest struct {
	FromId          uint32 `protobuf:"varint,1,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	InlineMessageId uint64 `protobuf:"varint,2,opt,name=inline_message_id,json=inlineMessageId" json:"inline_message_id,omitempty"`
	UserData        []byte `protobuf:"bytes,3,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	Date            uint32 `protobuf:"varint,4,opt,name=date" json:"date,omitempty"`
	Lang            string `protobuf:"bytes,5,opt,name=lang" json:"lang,omitempty"`
}

func (m *CallbackQueryRequest) Reset()                    { *m = CallbackQueryRequest{} }
func (m *CallbackQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*CallbackQueryRequest) ProtoMessage()               {}
func (*CallbackQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CallbackQueryRequest) GetFromId() uint32 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *CallbackQueryRequest) GetInlineMessageId() uint64 {
	if m != nil {
		return m.InlineMessageId
	}
	return 0
}

func (m *CallbackQueryRequest) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

func (m *CallbackQueryRequest) GetDate() uint32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *CallbackQueryRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

// 推送回调查询响应
type CallbackQueryReply struct {
}

func (m *CallbackQueryReply) Reset()                    { *m = CallbackQueryReply{} }
func (m *CallbackQueryReply) String() string            { return proto.CompactTextString(m) }
func (*CallbackQueryReply) ProtoMessage()               {}
func (*CallbackQueryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 推送内联查询请求
type GetInlineQueryRequest struct {
	FromId          int32  `protobuf:"varint,1,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	InlineQueryData []byte `protobuf:"bytes,2,opt,name=inline_query_data,json=inlineQueryData,proto3" json:"inline_query_data,omitempty"`
}

func (m *GetInlineQueryRequest) Reset()                    { *m = GetInlineQueryRequest{} }
func (m *GetInlineQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetInlineQueryRequest) ProtoMessage()               {}
func (*GetInlineQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetInlineQueryRequest) GetFromId() int32 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *GetInlineQueryRequest) GetInlineQueryData() []byte {
	if m != nil {
		return m.InlineQueryData
	}
	return nil
}

// 推送内联查询响应
type GetInlineQueryReply struct {
	Ok               bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Msg              string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	InlineResultData []byte `protobuf:"bytes,3,opt,name=inline_result_data,json=inlineResultData,proto3" json:"inline_result_data,omitempty"`
}

func (m *GetInlineQueryReply) Reset()                    { *m = GetInlineQueryReply{} }
func (m *GetInlineQueryReply) String() string            { return proto.CompactTextString(m) }
func (*GetInlineQueryReply) ProtoMessage()               {}
func (*GetInlineQueryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetInlineQueryReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetInlineQueryReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetInlineQueryReply) GetInlineResultData() []byte {
	if m != nil {
		return m.InlineResultData
	}
	return nil
}

type SendInlineQueryRequest struct {
	FromId        int32  `protobuf:"varint,1,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	SendQueryData []byte `protobuf:"bytes,2,opt,name=send_query_data,json=sendQueryData,proto3" json:"send_query_data,omitempty"`
}

func (m *SendInlineQueryRequest) Reset()                    { *m = SendInlineQueryRequest{} }
func (m *SendInlineQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*SendInlineQueryRequest) ProtoMessage()               {}
func (*SendInlineQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SendInlineQueryRequest) GetFromId() int32 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *SendInlineQueryRequest) GetSendQueryData() []byte {
	if m != nil {
		return m.SendQueryData
	}
	return nil
}

type SendInlineQueryReply struct {
	Ok             bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Msg            string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	SendResultData []byte `protobuf:"bytes,3,opt,name=send_result_data,json=sendResultData,proto3" json:"send_result_data,omitempty"`
}

func (m *SendInlineQueryReply) Reset()                    { *m = SendInlineQueryReply{} }
func (m *SendInlineQueryReply) String() string            { return proto.CompactTextString(m) }
func (*SendInlineQueryReply) ProtoMessage()               {}
func (*SendInlineQueryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SendInlineQueryReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SendInlineQueryReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SendInlineQueryReply) GetSendResultData() []byte {
	if m != nil {
		return m.SendResultData
	}
	return nil
}

type PushInlineQueryRequest struct {
	UserId    int32         `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	MessageId uint64        `protobuf:"varint,2,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	KeyId     uint64        `protobuf:"varint,3,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	SessionId uint64        `protobuf:"varint,4,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Datetime  int64         `protobuf:"varint,5,opt,name=datetime" json:"datetime,omitempty"`
	Data      []byte        `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Dbg       *pbcomm.Debug `protobuf:"bytes,7,opt,name=dbg" json:"dbg,omitempty"`
}

func (m *PushInlineQueryRequest) Reset()                    { *m = PushInlineQueryRequest{} }
func (m *PushInlineQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*PushInlineQueryRequest) ProtoMessage()               {}
func (*PushInlineQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PushInlineQueryRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PushInlineQueryRequest) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PushInlineQueryRequest) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *PushInlineQueryRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushInlineQueryRequest) GetDatetime() int64 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

func (m *PushInlineQueryRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PushInlineQueryRequest) GetDbg() *pbcomm.Debug {
	if m != nil {
		return m.Dbg
	}
	return nil
}

type PushInlineQueryRequestV2 struct {
	UserId    int32         `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	MessageId uint64        `protobuf:"varint,2,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	KeyId     uint64        `protobuf:"varint,3,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	SessionId uint64        `protobuf:"varint,4,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Datetime  int64         `protobuf:"varint,5,opt,name=datetime" json:"datetime,omitempty"`
	Data      []byte        `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Dbg       *pbcomm.Debug `protobuf:"bytes,7,opt,name=dbg" json:"dbg,omitempty"`
	Lang      string        `protobuf:"bytes,8,opt,name=lang" json:"lang,omitempty"`
}

func (m *PushInlineQueryRequestV2) Reset()                    { *m = PushInlineQueryRequestV2{} }
func (m *PushInlineQueryRequestV2) String() string            { return proto.CompactTextString(m) }
func (*PushInlineQueryRequestV2) ProtoMessage()               {}
func (*PushInlineQueryRequestV2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PushInlineQueryRequestV2) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PushInlineQueryRequestV2) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PushInlineQueryRequestV2) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *PushInlineQueryRequestV2) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushInlineQueryRequestV2) GetDatetime() int64 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

func (m *PushInlineQueryRequestV2) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PushInlineQueryRequestV2) GetDbg() *pbcomm.Debug {
	if m != nil {
		return m.Dbg
	}
	return nil
}

func (m *PushInlineQueryRequestV2) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOldBotsRequest)(nil), "updateservice.GetOldBotsRequest")
	proto.RegisterType((*GetOldBotsReply)(nil), "updateservice.GetOldBotsReply")
	proto.RegisterType((*PushMessageListRequest)(nil), "updateservice.PushMessageListRequest")
	proto.RegisterType((*PushMessageRequest)(nil), "updateservice.PushMessageRequest")
	proto.RegisterType((*PushMessageReply)(nil), "updateservice.PushMessageReply")
	proto.RegisterType((*CallbackQueryRequest)(nil), "updateservice.CallbackQueryRequest")
	proto.RegisterType((*CallbackQueryReply)(nil), "updateservice.CallbackQueryReply")
	proto.RegisterType((*GetInlineQueryRequest)(nil), "updateservice.GetInlineQueryRequest")
	proto.RegisterType((*GetInlineQueryReply)(nil), "updateservice.GetInlineQueryReply")
	proto.RegisterType((*SendInlineQueryRequest)(nil), "updateservice.SendInlineQueryRequest")
	proto.RegisterType((*SendInlineQueryReply)(nil), "updateservice.SendInlineQueryReply")
	proto.RegisterType((*PushInlineQueryRequest)(nil), "updateservice.PushInlineQueryRequest")
	proto.RegisterType((*PushInlineQueryRequestV2)(nil), "updateservice.PushInlineQueryRequestV2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Bot service

type BotClient interface {
	// GetOldBots 获取旧机器人
	GetOldBots(ctx context.Context, in *GetOldBotsRequest, opts ...grpc.CallOption) (*GetOldBotsReply, error)
	// PushMessage 推送消息
	PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error)
	// PushMessageList 推送消息
	PushMessageList(ctx context.Context, in *PushMessageListRequest, opts ...grpc.CallOption) (*PushMessageReply, error)
	// PushCallbackQuery 推送回调查询
	PushCallbackQuery(ctx context.Context, in *CallbackQueryRequest, opts ...grpc.CallOption) (*CallbackQueryReply, error)
	// GetInlineQuery 获取内联查询
	GetInlineQuery(ctx context.Context, in *GetInlineQueryRequest, opts ...grpc.CallOption) (*GetInlineQueryReply, error)
	// SendInlineQuery 发送内联查询
	SendInlineQuery(ctx context.Context, in *SendInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQuery(ctx context.Context, in *PushInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQueryV2(ctx context.Context, in *PushInlineQueryRequestV2, opts ...grpc.CallOption) (*SendInlineQueryReply, error)
}

type botClient struct {
	cc *grpc.ClientConn
}

func NewBotClient(cc *grpc.ClientConn) BotClient {
	return &botClient{cc}
}

func (c *botClient) GetOldBots(ctx context.Context, in *GetOldBotsRequest, opts ...grpc.CallOption) (*GetOldBotsReply, error) {
	out := new(GetOldBotsReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/GetOldBots", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error) {
	out := new(PushMessageReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/PushMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushMessageList(ctx context.Context, in *PushMessageListRequest, opts ...grpc.CallOption) (*PushMessageReply, error) {
	out := new(PushMessageReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/PushMessageList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushCallbackQuery(ctx context.Context, in *CallbackQueryRequest, opts ...grpc.CallOption) (*CallbackQueryReply, error) {
	out := new(CallbackQueryReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/PushCallbackQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) GetInlineQuery(ctx context.Context, in *GetInlineQueryRequest, opts ...grpc.CallOption) (*GetInlineQueryReply, error) {
	out := new(GetInlineQueryReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/GetInlineQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) SendInlineQuery(ctx context.Context, in *SendInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/SendInlineQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushInlineQuery(ctx context.Context, in *PushInlineQueryRequest, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/PushInlineQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) PushInlineQueryV2(ctx context.Context, in *PushInlineQueryRequestV2, opts ...grpc.CallOption) (*SendInlineQueryReply, error) {
	out := new(SendInlineQueryReply)
	err := grpc.Invoke(ctx, "/updateservice.Bot/PushInlineQueryV2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bot service

type BotServer interface {
	// GetOldBots 获取旧机器人
	GetOldBots(context.Context, *GetOldBotsRequest) (*GetOldBotsReply, error)
	// PushMessage 推送消息
	PushMessage(context.Context, *PushMessageRequest) (*PushMessageReply, error)
	// PushMessageList 推送消息
	PushMessageList(context.Context, *PushMessageListRequest) (*PushMessageReply, error)
	// PushCallbackQuery 推送回调查询
	PushCallbackQuery(context.Context, *CallbackQueryRequest) (*CallbackQueryReply, error)
	// GetInlineQuery 获取内联查询
	GetInlineQuery(context.Context, *GetInlineQueryRequest) (*GetInlineQueryReply, error)
	// SendInlineQuery 发送内联查询
	SendInlineQuery(context.Context, *SendInlineQueryRequest) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQuery(context.Context, *PushInlineQueryRequest) (*SendInlineQueryReply, error)
	// PushInlineQuery 内联查询
	PushInlineQueryV2(context.Context, *PushInlineQueryRequestV2) (*SendInlineQueryReply, error)
}

func RegisterBotServer(s *grpc.Server, srv BotServer) {
	s.RegisterService(&_Bot_serviceDesc, srv)
}

func _Bot_GetOldBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOldBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetOldBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/GetOldBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetOldBots(ctx, req.(*GetOldBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushMessage(ctx, req.(*PushMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushMessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushMessageList(ctx, req.(*PushMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushCallbackQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushCallbackQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushCallbackQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushCallbackQuery(ctx, req.(*CallbackQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_GetInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/GetInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetInlineQuery(ctx, req.(*GetInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_SendInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).SendInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/SendInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).SendInlineQuery(ctx, req.(*SendInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushInlineQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushInlineQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushInlineQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushInlineQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushInlineQuery(ctx, req.(*PushInlineQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_PushInlineQueryV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushInlineQueryRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).PushInlineQueryV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateservice.Bot/PushInlineQueryV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).PushInlineQueryV2(ctx, req.(*PushInlineQueryRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "updateservice.Bot",
	HandlerType: (*BotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOldBots",
			Handler:    _Bot_GetOldBots_Handler,
		},
		{
			MethodName: "PushMessage",
			Handler:    _Bot_PushMessage_Handler,
		},
		{
			MethodName: "PushMessageList",
			Handler:    _Bot_PushMessageList_Handler,
		},
		{
			MethodName: "PushCallbackQuery",
			Handler:    _Bot_PushCallbackQuery_Handler,
		},
		{
			MethodName: "GetInlineQuery",
			Handler:    _Bot_GetInlineQuery_Handler,
		},
		{
			MethodName: "SendInlineQuery",
			Handler:    _Bot_SendInlineQuery_Handler,
		},
		{
			MethodName: "PushInlineQuery",
			Handler:    _Bot_PushInlineQuery_Handler,
		},
		{
			MethodName: "PushInlineQueryV2",
			Handler:    _Bot_PushInlineQueryV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/im/iminterfaceprotobufs/botserver/updateservice/updateservice.proto",
}

func init() {
	proto.RegisterFile("gitlab.chatserver.im/im/iminterfaceprotobufs/botserver/updateservice/updateservice.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xfe, 0x65, 0x69, 0xbb, 0xf6, 0xec, 0xd7, 0xb5, 0xf3, 0xfe, 0x10, 0x05, 0xc1, 0xba, 0x8c,
	0x41, 0x85, 0x50, 0x2b, 0x95, 0x4b, 0xee, 0xb6, 0x49, 0x53, 0x25, 0x26, 0x86, 0x27, 0x4d, 0x42,
	0x0c, 0x55, 0x4e, 0xe3, 0x65, 0x51, 0x93, 0xba, 0x8b, 0x1d, 0xa4, 0x4a, 0x3c, 0x02, 0x0f, 0xc1,
	0x5b, 0xf1, 0x14, 0x7b, 0x07, 0x64, 0x27, 0xb4, 0x69, 0x12, 0xba, 0x82, 0xc4, 0x05, 0x57, 0xb5,
	0xcf, 0x39, 0xfe, 0xbe, 0xef, 0xb8, 0xc7, 0x9f, 0x02, 0x97, 0xae, 0x27, 0x7c, 0x62, 0x77, 0x86,
	0xb7, 0x44, 0x70, 0x1a, 0x7e, 0xa6, 0x61, 0xc7, 0x0b, 0xba, 0x5e, 0xd0, 0x9d, 0x30, 0x41, 0x04,
	0x9b, 0x84, 0x4c, 0x30, 0x3b, 0xba, 0xe1, 0x5d, 0x9b, 0x25, 0xf9, 0x6e, 0x34, 0x71, 0x88, 0xa0,
	0x72, 0xe3, 0x0d, 0xe9, 0xe2, 0xae, 0xa3, 0xea, 0x51, 0x7d, 0x21, 0x68, 0xbe, 0x59, 0x95, 0x63,
	0x62, 0x0f, 0x59, 0x10, 0x24, 0x3f, 0x31, 0x96, 0xb5, 0x0d, 0x5b, 0x67, 0x54, 0xbc, 0xf3, 0x9d,
	0x63, 0x26, 0x38, 0xa6, 0x77, 0x11, 0xe5, 0xc2, 0x3a, 0x82, 0x46, 0x3a, 0x38, 0xf1, 0xa7, 0x08,
	0x41, 0x49, 0x4a, 0x34, 0xb4, 0x96, 0xde, 0x2e, 0x63, 0xb5, 0xb6, 0xbe, 0xc0, 0xde, 0x45, 0xc4,
	0x6f, 0xcf, 0x29, 0xe7, 0xc4, 0xa5, 0x6f, 0x3d, 0x2e, 0x12, 0x00, 0xb4, 0x0b, 0x15, 0x9b, 0x89,
	0x81, 0xe7, 0xa8, 0xfa, 0x3a, 0x2e, 0xdb, 0x4c, 0xf4, 0x1d, 0xf4, 0x18, 0x6a, 0x11, 0xa7, 0xe1,
	0xc0, 0x21, 0x82, 0x18, 0x6b, 0x2d, 0xad, 0xfd, 0x3f, 0xae, 0xca, 0xc0, 0x29, 0x11, 0x04, 0x75,
	0x00, 0xd9, 0xe1, 0x0d, 0x0b, 0xe9, 0xb9, 0x17, 0x12, 0x97, 0x9e, 0xdc, 0x12, 0xe1, 0x39, 0x86,
	0xde, 0xd2, 0xda, 0x75, 0x5c, 0x90, 0xb1, 0xbe, 0x6a, 0x80, 0x52, 0xf4, 0x45, 0xd4, 0xda, 0xdf,
	0xa1, 0x96, 0x97, 0xe1, 0x93, 0xb1, 0x6b, 0x94, 0x5a, 0x5a, 0xbb, 0x86, 0xd5, 0xda, 0x42, 0xd0,
	0x5c, 0x50, 0x33, 0xf1, 0xa7, 0xd6, 0x37, 0x0d, 0x76, 0x4e, 0x88, 0xef, 0xdb, 0x64, 0x38, 0x7a,
	0x1f, 0xd1, 0x70, 0xfa, 0x53, 0xe4, 0x23, 0x58, 0xbf, 0x09, 0x59, 0x30, 0x57, 0x59, 0x91, 0xdb,
	0xbe, 0x83, 0x5e, 0xc2, 0x96, 0x37, 0xf6, 0xbd, 0x31, 0x1d, 0x04, 0x31, 0x90, 0x2c, 0x91, 0x72,
	0x4b, 0xb8, 0x11, 0x27, 0x12, 0x82, 0x6c, 0x4b, 0x7a, 0xa6, 0x25, 0x04, 0x25, 0x39, 0x23, 0x4a,
	0x62, 0x1d, 0xab, 0xf5, 0x4c, 0x76, 0x39, 0x25, 0x7b, 0x07, 0x50, 0x46, 0xa1, 0x14, 0x7e, 0x0d,
	0xbb, 0x67, 0x54, 0xf4, 0x15, 0xe1, 0x32, 0xe1, 0xe5, 0x02, 0xe1, 0x77, 0xb2, 0x3e, 0x7d, 0xcf,
	0x89, 0x70, 0x85, 0x23, 0xb5, 0x59, 0x14, 0xb6, 0xb3, 0xe8, 0x72, 0xc4, 0x36, 0x61, 0x8d, 0x8d,
	0x14, 0x6c, 0x15, 0xaf, 0xb1, 0x11, 0x6a, 0x82, 0x1e, 0x70, 0x57, 0x81, 0xd4, 0xb0, 0x5c, 0xa2,
	0x57, 0x80, 0x12, 0x92, 0x90, 0xf2, 0xc8, 0x17, 0xe9, 0xd6, 0x9b, 0x71, 0x06, 0xab, 0x84, 0xa2,
	0xf9, 0x00, 0x7b, 0x97, 0x74, 0xec, 0xfc, 0x4e, 0x17, 0xcf, 0xa1, 0xc1, 0xe9, 0xd8, 0xc9, 0xf7,
	0x50, 0x97, 0xe1, 0x79, 0x07, 0x36, 0xec, 0xe4, 0xa0, 0x57, 0x6b, 0xa1, 0x0d, 0x4d, 0xc5, 0x90,
	0x6f, 0x60, 0x53, 0xc6, 0x53, 0xf2, 0xbf, 0x6b, 0xf1, 0xf3, 0x2a, 0xd6, 0xaf, 0xfe, 0xf9, 0xb9,
	0x7e, 0xb9, 0xed, 0x3b, 0xe8, 0x09, 0x40, 0x6e, 0x6e, 0x6a, 0xc1, 0x6c, 0x62, 0x76, 0xa1, 0x32,
	0xa2, 0xd3, 0x41, 0x32, 0xdb, 0x25, 0x5c, 0x1e, 0xd1, 0x69, 0x7c, 0x8a, 0x53, 0xce, 0x3d, 0x36,
	0x96, 0xa9, 0x52, 0x7c, 0x2a, 0x89, 0xf4, 0x1d, 0x64, 0x42, 0x55, 0x8e, 0x8f, 0xf0, 0x02, 0xaa,
	0x46, 0x47, 0xc7, 0xb3, 0x7d, 0x32, 0x66, 0xc4, 0xa8, 0xa8, 0x16, 0xd4, 0x1a, 0xed, 0x83, 0xee,
	0xd8, 0xae, 0xb1, 0xde, 0xd2, 0xda, 0x1b, 0xbd, 0x7a, 0x27, 0xb1, 0x9b, 0x53, 0x6a, 0x47, 0x2e,
	0x96, 0x19, 0xeb, 0x5e, 0x03, 0xa3, 0xb8, 0xb3, 0xab, 0xde, 0x3f, 0xde, 0xdb, 0xec, 0x8d, 0x55,
	0xe7, 0x6f, 0xac, 0x77, 0x5f, 0x06, 0xfd, 0x98, 0x09, 0x74, 0x01, 0x30, 0xb7, 0x55, 0xd4, 0xea,
	0x2c, 0x7a, 0x7b, 0xce, 0x86, 0xcd, 0xa7, 0x4b, 0x2a, 0xe4, 0x2b, 0xfd, 0x0f, 0x5d, 0xc2, 0x46,
	0xca, 0x74, 0xd0, 0x41, 0xe6, 0x40, 0xde, 0x1e, 0xcd, 0xfd, 0x65, 0x25, 0x31, 0xe8, 0x47, 0x68,
	0x64, 0x6c, 0x1d, 0x1d, 0xfd, 0xfa, 0x54, 0xca, 0xf6, 0x57, 0x01, 0xff, 0x04, 0x5b, 0x32, 0xba,
	0xe0, 0x39, 0xe8, 0x30, 0x73, 0xae, 0xc8, 0x33, 0xcd, 0x83, 0xe5, 0x45, 0x31, 0xfc, 0x35, 0x6c,
	0x2e, 0x5a, 0x0b, 0x7a, 0x96, 0xbf, 0xc4, 0xfc, 0xdc, 0x99, 0xd6, 0x03, 0x55, 0x31, 0xfa, 0x00,
	0x1a, 0x99, 0x67, 0x9f, 0xbb, 0x99, 0x62, 0xc7, 0x31, 0x0f, 0x1f, 0x2a, 0x9b, 0x11, 0x64, 0x1e,
	0x46, 0xe1, 0xd5, 0xff, 0x39, 0xc1, 0x30, 0xbe, 0xfe, 0x54, 0xe6, 0xaa, 0x87, 0x5e, 0xac, 0x44,
	0x71, 0xd5, 0x5b, 0x91, 0xc4, 0xae, 0xa8, 0x4f, 0x8b, 0xd7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x76, 0x56, 0x8a, 0xfd, 0x08, 0x00, 0x00,
}
