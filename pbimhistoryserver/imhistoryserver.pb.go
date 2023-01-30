// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitlab.chatserver.im/interfaceprobuf/pbimhistoryserver/imhistoryserver.proto

/*
Package pbimhistoryserver is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/interfaceprobuf/pbimhistoryserver/imhistoryserver.proto

It has these top-level messages:

	GetHistoryReq
	GetHistoryResult
	GetMessagesReq
	GetMessagesResult
*/
package pbimhistoryserver

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pbcomm "gitlab.chatserver.im/interfaceprobuf/pbcomm"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type GetHistoryReq struct {
	ObjBin    []byte        `protobuf:"bytes,1,opt,name=obj_bin,json=objBin,proto3" json:"obj_bin,omitempty"`
	CrcId     int64         `protobuf:"varint,2,opt,name=crc_id,json=crcId,proto3" json:"crc_id,omitempty"`
	MessageId int64         `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	KeyId     int64         `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	UserId    int32         `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Debug     *pbcomm.Debug `protobuf:"bytes,6,opt,name=debug" json:"debug,omitempty"`
}

func (m *GetHistoryReq) Reset()                    { *m = GetHistoryReq{} }
func (m *GetHistoryReq) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryReq) ProtoMessage()               {}
func (*GetHistoryReq) Descriptor() ([]byte, []int) { return fileDescriptorImhistoryserver, []int{0} }

func (m *GetHistoryReq) GetObjBin() []byte {
	if m != nil {
		return m.ObjBin
	}
	return nil
}

func (m *GetHistoryReq) GetCrcId() int64 {
	if m != nil {
		return m.CrcId
	}
	return 0
}

func (m *GetHistoryReq) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *GetHistoryReq) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *GetHistoryReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetHistoryReq) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type GetHistoryResult struct {
	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (m *GetHistoryResult) Reset()                    { *m = GetHistoryResult{} }
func (m *GetHistoryResult) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryResult) ProtoMessage()               {}
func (*GetHistoryResult) Descriptor() ([]byte, []int) { return fileDescriptorImhistoryserver, []int{1} }

func (m *GetHistoryResult) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetHistoryResult) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type GetMessagesReq struct {
	ObjBin    []byte        `protobuf:"bytes,1,opt,name=obj_bin,json=objBin,proto3" json:"obj_bin,omitempty"`
	CrcId     int64         `protobuf:"varint,2,opt,name=crc_id,json=crcId,proto3" json:"crc_id,omitempty"`
	MessageId int64         `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	KeyId     int64         `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	UserId    int32         `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Debug     *pbcomm.Debug `protobuf:"bytes,6,opt,name=debug" json:"debug,omitempty"`
}

func (m *GetMessagesReq) Reset()                    { *m = GetMessagesReq{} }
func (m *GetMessagesReq) String() string            { return proto.CompactTextString(m) }
func (*GetMessagesReq) ProtoMessage()               {}
func (*GetMessagesReq) Descriptor() ([]byte, []int) { return fileDescriptorImhistoryserver, []int{2} }

func (m *GetMessagesReq) GetObjBin() []byte {
	if m != nil {
		return m.ObjBin
	}
	return nil
}

func (m *GetMessagesReq) GetCrcId() int64 {
	if m != nil {
		return m.CrcId
	}
	return 0
}

func (m *GetMessagesReq) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *GetMessagesReq) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *GetMessagesReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetMessagesReq) GetDebug() *pbcomm.Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

type GetMessagesResult struct {
	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (m *GetMessagesResult) Reset()         { *m = GetMessagesResult{} }
func (m *GetMessagesResult) String() string { return proto.CompactTextString(m) }
func (*GetMessagesResult) ProtoMessage()    {}
func (*GetMessagesResult) Descriptor() ([]byte, []int) {
	return fileDescriptorImhistoryserver, []int{3}
}

func (m *GetMessagesResult) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetMessagesResult) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*GetHistoryReq)(nil), "pbimhistoryserver.GetHistoryReq")
	proto.RegisterType((*GetHistoryResult)(nil), "pbimhistoryserver.GetHistoryResult")
	proto.RegisterType((*GetMessagesReq)(nil), "pbimhistoryserver.GetMessagesReq")
	proto.RegisterType((*GetMessagesResult)(nil), "pbimhistoryserver.GetMessagesResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HistoryServer service

type HistoryServerClient interface {
	//    获取历史
	GetHistory(ctx context.Context, in *GetHistoryReq, opts ...grpc.CallOption) (*GetHistoryResult, error)
	GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesResult, error)
}

type historyServerClient struct {
	cc *grpc.ClientConn
}

func NewHistoryServerClient(cc *grpc.ClientConn) HistoryServerClient {
	return &historyServerClient{cc}
}

func (c *historyServerClient) GetHistory(ctx context.Context, in *GetHistoryReq, opts ...grpc.CallOption) (*GetHistoryResult, error) {
	out := new(GetHistoryResult)
	err := grpc.Invoke(ctx, "/pbimhistoryserver.HistoryServer/GetHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServerClient) GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesResult, error) {
	out := new(GetMessagesResult)
	err := grpc.Invoke(ctx, "/pbimhistoryserver.HistoryServer/GetMessages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HistoryServer service

type HistoryServerServer interface {
	//    获取历史
	GetHistory(context.Context, *GetHistoryReq) (*GetHistoryResult, error)
	GetMessages(context.Context, *GetMessagesReq) (*GetMessagesResult, error)
}

func RegisterHistoryServerServer(s *grpc.Server, srv HistoryServerServer) {
	s.RegisterService(&_HistoryServer_serviceDesc, srv)
}

func _HistoryServer_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServerServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbimhistoryserver.HistoryServer/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServerServer).GetHistory(ctx, req.(*GetHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryServer_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServerServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbimhistoryserver.HistoryServer/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServerServer).GetMessages(ctx, req.(*GetMessagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HistoryServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbimhistoryserver.HistoryServer",
	HandlerType: (*HistoryServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHistory",
			Handler:    _HistoryServer_GetHistory_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _HistoryServer_GetMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/interfaceprobuf/pbimhistoryserver/imhistoryserver.proto",
}

func (m *GetHistoryReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetHistoryReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ObjBin) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(len(m.ObjBin)))
		i += copy(dAtA[i:], m.ObjBin)
	}
	if m.CrcId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.CrcId))
	}
	if m.MessageId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.MessageId))
	}
	if m.KeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.KeyId))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.UserId))
	}
	if m.Debug != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.Debug.Size()))
		n1, err := m.Debug.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *GetHistoryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetHistoryResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	if m.ErrorCode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.ErrorCode))
	}
	return i, nil
}

func (m *GetMessagesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMessagesReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ObjBin) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(len(m.ObjBin)))
		i += copy(dAtA[i:], m.ObjBin)
	}
	if m.CrcId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.CrcId))
	}
	if m.MessageId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.MessageId))
	}
	if m.KeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.KeyId))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.UserId))
	}
	if m.Debug != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.Debug.Size()))
		n2, err := m.Debug.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *GetMessagesResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMessagesResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	if m.ErrorCode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImhistoryserver(dAtA, i, uint64(m.ErrorCode))
	}
	return i, nil
}

func encodeVarintImhistoryserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetHistoryReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.ObjBin)
	if l > 0 {
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	if m.CrcId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.CrcId))
	}
	if m.MessageId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.MessageId))
	}
	if m.KeyId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.KeyId))
	}
	if m.UserId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.UserId))
	}
	if m.Debug != nil {
		l = m.Debug.Size()
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	return n
}

func (m *GetHistoryResult) Size() (n int) {
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovImhistoryserver(uint64(m.ErrorCode))
	}
	return n
}

func (m *GetMessagesReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.ObjBin)
	if l > 0 {
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	if m.CrcId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.CrcId))
	}
	if m.MessageId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.MessageId))
	}
	if m.KeyId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.KeyId))
	}
	if m.UserId != 0 {
		n += 1 + sovImhistoryserver(uint64(m.UserId))
	}
	if m.Debug != nil {
		l = m.Debug.Size()
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	return n
}

func (m *GetMessagesResult) Size() (n int) {
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovImhistoryserver(uint64(l))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovImhistoryserver(uint64(m.ErrorCode))
	}
	return n
}

func sovImhistoryserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImhistoryserver(x uint64) (n int) {
	return sovImhistoryserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetHistoryReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImhistoryserver
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
			return fmt.Errorf("proto: GetHistoryReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetHistoryReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
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
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjBin = append(m.ObjBin[:0], dAtA[iNdEx:postIndex]...)
			if m.ObjBin == nil {
				m.ObjBin = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrcId", wireType)
			}
			m.CrcId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CrcId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Debug == nil {
				m.Debug = &pbcomm.Debug{}
			}
			if err := m.Debug.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImhistoryserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImhistoryserver
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
func (m *GetHistoryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImhistoryserver
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
			return fmt.Errorf("proto: GetHistoryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetHistoryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
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
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImhistoryserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImhistoryserver
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
func (m *GetMessagesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImhistoryserver
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
			return fmt.Errorf("proto: GetMessagesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMessagesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
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
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjBin = append(m.ObjBin[:0], dAtA[iNdEx:postIndex]...)
			if m.ObjBin == nil {
				m.ObjBin = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrcId", wireType)
			}
			m.CrcId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CrcId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Debug == nil {
				m.Debug = &pbcomm.Debug{}
			}
			if err := m.Debug.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImhistoryserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImhistoryserver
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
func (m *GetMessagesResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImhistoryserver
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
			return fmt.Errorf("proto: GetMessagesResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMessagesResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
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
				return ErrInvalidLengthImhistoryserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImhistoryserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImhistoryserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImhistoryserver
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
func skipImhistoryserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImhistoryserver
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
					return 0, ErrIntOverflowImhistoryserver
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
					return 0, ErrIntOverflowImhistoryserver
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
				return 0, ErrInvalidLengthImhistoryserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImhistoryserver
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
				next, err := skipImhistoryserver(dAtA[start:])
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
	ErrInvalidLengthImhistoryserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImhistoryserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.chatserver.im/interfaceprobuf/pbimhistoryserver/imhistoryserver.proto", fileDescriptorImhistoryserver)
}

var fileDescriptorImhistoryserver = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0xcf, 0x6e, 0xda, 0x30,
	0x1c, 0xc7, 0xe7, 0xb1, 0x64, 0xc2, 0xc0, 0x34, 0x2c, 0x6d, 0x43, 0x48, 0x43, 0x59, 0xd8, 0x21,
	0xa7, 0x20, 0xb1, 0xcb, 0xce, 0x6c, 0x12, 0xcb, 0xd4, 0x5e, 0x52, 0x55, 0xea, 0x0d, 0xc5, 0xf6,
	0x0f, 0x30, 0x10, 0x4c, 0x1d, 0xa7, 0x12, 0xcf, 0xd5, 0x3e, 0x40, 0x8f, 0x3d, 0xf6, 0x11, 0x2a,
	0x9e, 0xa4, 0x72, 0x9c, 0xaa, 0x50, 0x2a, 0x51, 0xa9, 0xa7, 0x9e, 0x92, 0xdf, 0xff, 0x8f, 0xbe,
	0xfe, 0xe2, 0xa3, 0x89, 0xd0, 0x8b, 0x84, 0x86, 0x6c, 0x9a, 0xe8, 0x0c, 0xd4, 0x05, 0xa8, 0x50,
	0xa4, 0x3d, 0xb1, 0xd4, 0xa0, 0xc6, 0x09, 0x83, 0x95, 0x92, 0x34, 0x1f, 0xf7, 0x56, 0x54, 0xa4,
	0x53, 0x91, 0x69, 0xa9, 0xd6, 0xb6, 0xa9, 0xf7, 0x24, 0x0e, 0x57, 0x4a, 0x6a, 0x49, 0x9a, 0x7b,
	0x8d, 0xed, 0xdf, 0x2f, 0x3c, 0xc0, 0x64, 0x9a, 0x96, 0x1f, 0xbb, 0xcc, 0xbf, 0x44, 0xb8, 0x31,
	0x04, 0xfd, 0xcf, 0xae, 0x8b, 0xe1, 0x9c, 0x7c, 0xc3, 0x1f, 0x25, 0x9d, 0x8d, 0xa8, 0x58, 0xb6,
	0x90, 0x87, 0x82, 0x7a, 0xec, 0x4a, 0x3a, 0x1b, 0x88, 0x25, 0xf9, 0x82, 0x5d, 0xa6, 0xd8, 0x48,
	0xf0, 0xd6, 0x7b, 0x0f, 0x05, 0x95, 0xd8, 0x61, 0x8a, 0x45, 0x9c, 0x7c, 0xc7, 0x38, 0x85, 0x2c,
	0x4b, 0x26, 0x60, 0x4a, 0x95, 0xa2, 0x54, 0x2d, 0x33, 0x11, 0x37, 0x53, 0x73, 0x58, 0x9b, 0xd2,
	0x07, 0x3b, 0x35, 0x87, 0x75, 0xc4, 0xcd, 0x95, 0x3c, 0x03, 0x65, 0xf2, 0x8e, 0x87, 0x02, 0x27,
	0x76, 0x4d, 0x18, 0x71, 0xd2, 0xc5, 0x0e, 0x07, 0x9a, 0x4f, 0x5a, 0xae, 0x87, 0x82, 0x5a, 0xbf,
	0x11, 0x96, 0xb8, 0x7f, 0x4d, 0x32, 0xb6, 0x35, 0x3f, 0xc2, 0x9f, 0xb7, 0xa1, 0xb3, 0x7c, 0xa1,
	0xc9, 0x57, 0xec, 0xaa, 0xe2, 0xef, 0x01, 0xdb, 0x46, 0x86, 0x0f, 0x94, 0x92, 0x6a, 0xc4, 0x24,
	0x87, 0x02, 0xdd, 0x89, 0xab, 0x45, 0xe6, 0x8f, 0xe4, 0xe0, 0x5f, 0x21, 0xfc, 0x69, 0x08, 0xfa,
	0xd8, 0x02, 0x67, 0x6f, 0x45, 0x81, 0xff, 0xb8, 0xb9, 0x43, 0xfd, 0x0a, 0x09, 0xfa, 0xd7, 0x08,
	0x37, 0x4a, 0x2d, 0x4f, 0x0a, 0xf3, 0x90, 0x53, 0x8c, 0x1f, 0xf5, 0x25, 0x5e, 0xb8, 0xe7, 0xb8,
	0x70, 0xc7, 0x33, 0xed, 0xee, 0x81, 0x0e, 0x43, 0xe1, 0xbf, 0x23, 0x67, 0xb8, 0xb6, 0x05, 0x4d,
	0x7e, 0x3c, 0x3f, 0xb5, 0xf5, 0x14, 0xed, 0x9f, 0x87, 0x5a, 0xec, 0xe6, 0x41, 0xfd, 0x66, 0xd3,
	0x41, 0xb7, 0x9b, 0x0e, 0xba, 0xdb, 0x74, 0x10, 0x75, 0x0b, 0x6f, 0xff, 0xba, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x59, 0x93, 0x48, 0x78, 0x03, 0x00, 0x00,
}
