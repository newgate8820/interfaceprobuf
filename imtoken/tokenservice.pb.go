// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitlab.chatserver.im/im/utillib/imtoken/tokenservice/tokenservice.proto

/*
Package tokenservice is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/im/utillib/imtoken/tokenservice/tokenservice.proto

It has these top-level messages:

	ReqGetTokenMsg
	GetTokenReply
	ReqValidateToken
	ValidateTokenReply
*/
package tokenservice

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TokenErrorCode int32

const (
	TokenErrorCode_Token_OK          TokenErrorCode = 0
	TokenErrorCode_Token_EXPIRED     TokenErrorCode = 1004
	TokenErrorCode_Token_INVALID     TokenErrorCode = 1002
	TokenErrorCode_Token_KEY_CHANGED TokenErrorCode = 1003
)

var TokenErrorCode_name = map[int32]string{
	0:    "Token_OK",
	1004: "Token_EXPIRED",
	1002: "Token_INVALID",
	1003: "Token_KEY_CHANGED",
}
var TokenErrorCode_value = map[string]int32{
	"Token_OK":          0,
	"Token_EXPIRED":     1004,
	"Token_INVALID":     1002,
	"Token_KEY_CHANGED": 1003,
}

func (x TokenErrorCode) String() string {
	return proto.EnumName(TokenErrorCode_name, int32(x))
}
func (TokenErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorTokenservice, []int{0} }

type ReqGetTokenMsg struct {
	KeyId  uint64 `protobuf:"varint,1,opt,name=KeyId,proto3" json:"KeyId,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Flag   uint32 `protobuf:"varint,3,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (m *ReqGetTokenMsg) Reset()                    { *m = ReqGetTokenMsg{} }
func (m *ReqGetTokenMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTokenMsg) ProtoMessage()               {}
func (*ReqGetTokenMsg) Descriptor() ([]byte, []int) { return fileDescriptorTokenservice, []int{0} }

func (m *ReqGetTokenMsg) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *ReqGetTokenMsg) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ReqGetTokenMsg) GetFlag() uint32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

type GetTokenReply struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token []byte `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *GetTokenReply) Reset()                    { *m = GetTokenReply{} }
func (m *GetTokenReply) String() string            { return proto.CompactTextString(m) }
func (*GetTokenReply) ProtoMessage()               {}
func (*GetTokenReply) Descriptor() ([]byte, []int) { return fileDescriptorTokenservice, []int{1} }

func (m *GetTokenReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetTokenReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetTokenReply) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type ReqValidateToken struct {
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *ReqValidateToken) Reset()                    { *m = ReqValidateToken{} }
func (m *ReqValidateToken) String() string            { return proto.CompactTextString(m) }
func (*ReqValidateToken) ProtoMessage()               {}
func (*ReqValidateToken) Descriptor() ([]byte, []int) { return fileDescriptorTokenservice, []int{2} }

func (m *ReqValidateToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenReply struct {
	Code  TokenErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=tokenservice.TokenErrorCode" json:"code,omitempty"`
	Error string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *ValidateTokenReply) Reset()                    { *m = ValidateTokenReply{} }
func (m *ValidateTokenReply) String() string            { return proto.CompactTextString(m) }
func (*ValidateTokenReply) ProtoMessage()               {}
func (*ValidateTokenReply) Descriptor() ([]byte, []int) { return fileDescriptorTokenservice, []int{3} }

func (m *ValidateTokenReply) GetCode() TokenErrorCode {
	if m != nil {
		return m.Code
	}
	return TokenErrorCode_Token_OK
}

func (m *ValidateTokenReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqGetTokenMsg)(nil), "tokenservice.ReqGetTokenMsg")
	proto.RegisterType((*GetTokenReply)(nil), "tokenservice.GetTokenReply")
	proto.RegisterType((*ReqValidateToken)(nil), "tokenservice.ReqValidateToken")
	proto.RegisterType((*ValidateTokenReply)(nil), "tokenservice.ValidateTokenReply")
	proto.RegisterEnum("tokenservice.TokenErrorCode", TokenErrorCode_name, TokenErrorCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tokenservice service

type TokenserviceClient interface {
	// 获取token
	GetToken(ctx context.Context, in *ReqGetTokenMsg, opts ...grpc.CallOption) (*GetTokenReply, error)
	// 验证tonken
	ValidateToken(ctx context.Context, in *ReqValidateToken, opts ...grpc.CallOption) (*ValidateTokenReply, error)
}

type tokenserviceClient struct {
	cc *grpc.ClientConn
}

func NewTokenserviceClient(cc *grpc.ClientConn) TokenserviceClient {
	return &tokenserviceClient{cc}
}

func (c *tokenserviceClient) GetToken(ctx context.Context, in *ReqGetTokenMsg, opts ...grpc.CallOption) (*GetTokenReply, error) {
	out := new(GetTokenReply)
	err := grpc.Invoke(ctx, "/tokenservice.Tokenservice/GetToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenserviceClient) ValidateToken(ctx context.Context, in *ReqValidateToken, opts ...grpc.CallOption) (*ValidateTokenReply, error) {
	out := new(ValidateTokenReply)
	err := grpc.Invoke(ctx, "/tokenservice.Tokenservice/ValidateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tokenservice service

type TokenserviceServer interface {
	// 获取token
	GetToken(context.Context, *ReqGetTokenMsg) (*GetTokenReply, error)
	// 验证tonken
	ValidateToken(context.Context, *ReqValidateToken) (*ValidateTokenReply, error)
}

func RegisterTokenserviceServer(s *grpc.Server, srv TokenserviceServer) {
	s.RegisterService(&_Tokenservice_serviceDesc, srv)
}

func _Tokenservice_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetTokenMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenserviceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenservice.Tokenservice/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenserviceServer).GetToken(ctx, req.(*ReqGetTokenMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokenservice_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqValidateToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenserviceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenservice.Tokenservice/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenserviceServer).ValidateToken(ctx, req.(*ReqValidateToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokenservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenservice.Tokenservice",
	HandlerType: (*TokenserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Tokenservice_GetToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Tokenservice_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/im/utillib/imtoken/tokenservice/tokenservice.proto",
}

func (m *ReqGetTokenMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqGetTokenMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.KeyId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(m.KeyId))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(m.UserId))
	}
	if m.Flag != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(m.Flag))
	}
	return i, nil
}

func (m *GetTokenReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTokenReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ok {
		dAtA[i] = 0x8
		i++
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *ReqValidateToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqValidateToken) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *ValidateTokenReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidateTokenReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(m.Code))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTokenservice(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func encodeVarintTokenservice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReqGetTokenMsg) Size() (n int) {
	var l int
	_ = l
	if m.KeyId != 0 {
		n += 1 + sovTokenservice(uint64(m.KeyId))
	}
	if m.UserId != 0 {
		n += 1 + sovTokenservice(uint64(m.UserId))
	}
	if m.Flag != 0 {
		n += 1 + sovTokenservice(uint64(m.Flag))
	}
	return n
}

func (m *GetTokenReply) Size() (n int) {
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTokenservice(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTokenservice(uint64(l))
	}
	return n
}

func (m *ReqValidateToken) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTokenservice(uint64(l))
	}
	return n
}

func (m *ValidateTokenReply) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTokenservice(uint64(m.Code))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTokenservice(uint64(l))
	}
	return n
}

func sovTokenservice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTokenservice(x uint64) (n int) {
	return sovTokenservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReqGetTokenMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReqGetTokenMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqGetTokenMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			m.Flag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flag |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTokenReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTokenReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTokenReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTokenservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTokenservice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token[:0], dAtA[iNdEx:postIndex]...)
			if m.Token == nil {
				m.Token = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReqValidateToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReqValidateToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqValidateToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTokenservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidateTokenReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidateTokenReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidateTokenReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (TokenErrorCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTokenservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTokenservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenservice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTokenservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTokenservice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTokenservice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTokenservice(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTokenservice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenservice   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.chatserver.im/im/utillib/imtoken/tokenservice/tokenservice.proto", fileDescriptorTokenservice)
}

var fileDescriptorTokenservice = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0xc5, 0x8c, 0x01, 0xbb, 0x4a, 0xa2, 0xcc, 0x9a, 0x10, 0x62, 0x53, 0x84, 0xf2, 0x84, 0xf6,
	0x90, 0x4c, 0xec, 0x0b, 0x18, 0x64, 0x59, 0x94, 0x8d, 0x55, 0x2e, 0x45, 0xad, 0x54, 0x15, 0x05,
	0x62, 0x51, 0x8b, 0x50, 0x43, 0x48, 0x2b, 0xf1, 0x4f, 0xfd, 0x90, 0x3e, 0xf6, 0x13, 0x2a, 0x1e,
	0xdb, 0x7e, 0x44, 0x15, 0xbb, 0x08, 0x5c, 0xfa, 0x62, 0xf9, 0x9c, 0x7b, 0xee, 0xb9, 0x57, 0xc7,
	0x06, 0x7f, 0xca, 0xb2, 0x24, 0x1a, 0x3b, 0x93, 0xcb, 0x28, 0x5b, 0xd1, 0xf4, 0x86, 0xa6, 0x0e,
	0x9b, 0xbb, 0x6c, 0xee, 0x5e, 0x67, 0x2c, 0x49, 0xd8, 0xd8, 0x65, 0xf3, 0x8c, 0xcf, 0xe8, 0x95,
	0x2b, 0xce, 0x5c, 0xc1, 0x26, 0x54, 0x01, 0xce, 0x22, 0xe5, 0x19, 0xc7, 0xda, 0x3e, 0x67, 0x13,
	0x30, 0x08, 0x5d, 0xfa, 0x34, 0x1b, 0xe4, 0xec, 0xbf, 0xd5, 0x14, 0x7f, 0x81, 0x8f, 0x21, 0x5d,
	0x07, 0x71, 0x1d, 0x35, 0x51, 0xab, 0x44, 0x24, 0xc0, 0x35, 0x28, 0x9f, 0xac, 0x68, 0x1a, 0xc4,
	0xf5, 0xa2, 0xa0, 0x5f, 0x11, 0xc6, 0x50, 0xfa, 0x9d, 0x44, 0xd3, 0xfa, 0x87, 0x26, 0x6a, 0xe9,
	0x44, 0xdc, 0xed, 0x10, 0xf4, 0xad, 0x21, 0xa1, 0x8b, 0x64, 0x8d, 0x0d, 0x28, 0xf2, 0x99, 0xf0,
	0xab, 0x92, 0x22, 0x9f, 0xe5, 0x23, 0x68, 0x9a, 0xf2, 0x54, 0x78, 0x7d, 0x22, 0x12, 0xe4, 0xac,
	0xe8, 0x11, 0x5e, 0x1a, 0x91, 0xc0, 0x6e, 0x81, 0x49, 0xe8, 0x72, 0x18, 0x25, 0x2c, 0x8e, 0x32,
	0x2a, 0xb8, 0x9d, 0x12, 0xc9, 0x7e, 0xa9, 0x3c, 0x07, 0xac, 0xc8, 0xe4, 0xec, 0x1f, 0x50, 0x9a,
	0xf0, 0x98, 0x0a, 0xa9, 0xd1, 0xfe, 0xe6, 0x28, 0x89, 0x08, 0x9d, 0x97, 0x4f, 0xef, 0xf2, 0x98,
	0x12, 0xa1, 0x7c, 0x7f, 0xbb, 0xef, 0x17, 0x60, 0xa8, 0x6a, 0xac, 0x41, 0x55, 0x30, 0xa3, 0xff,
	0xa1, 0x59, 0xc0, 0x18, 0x74, 0x89, 0xbc, 0xd3, 0xa3, 0x80, 0x78, 0x3d, 0xf3, 0xb9, 0xb2, 0xe3,
	0x82, 0xfe, 0xb0, 0xf3, 0x37, 0xe8, 0x99, 0x8f, 0x15, 0x5c, 0x83, 0xcf, 0x92, 0x0b, 0xbd, 0xb3,
	0x51, 0xf7, 0x4f, 0xa7, 0xef, 0x7b, 0x3d, 0xf3, 0xa9, 0xd2, 0xbe, 0x45, 0xa0, 0x0d, 0xf6, 0x76,
	0xc3, 0x3e, 0x54, 0xb7, 0x29, 0xe2, 0x37, 0x6b, 0xab, 0x2f, 0xd6, 0xf8, 0xaa, 0x56, 0x95, 0xec,
	0xed, 0x02, 0x3e, 0x06, 0x5d, 0x8d, 0xcf, 0x3a, 0x70, 0x53, 0xea, 0x8d, 0xa6, 0x5a, 0x3f, 0x0c,
	0xd5, 0x2e, 0xfc, 0xd2, 0xee, 0x36, 0x16, 0xba, 0xdf, 0x58, 0xe8, 0x61, 0x63, 0xa1, 0x71, 0x59,
	0x7c, 0xad, 0x9f, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0xa3, 0x19, 0x44, 0xa5, 0x02, 0x00,
	0x00,
}
