// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitlab.chatserver.im/interfaceprobuf/imslowmodeserver/imslowmodeserver.proto

/*
Package imslowmodeserver is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/interfaceprobuf/imslowmodeserver/imslowmodeserver.proto

It has these top-level messages:

	ToggleSlowModeReq
	ToggleSlowModeRsp
	CheckSlowLimitReq
	CheckSlowLimitRsp
*/
package imslowmodeserver

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

type ToggleSlowModeReq struct {
	ChnlId  int32 `protobuf:"varint,1,opt,name=Chnl_id,json=ChnlId,proto3" json:"Chnl_id,omitempty"`
	Seconds int32 `protobuf:"varint,2,opt,name=Seconds,proto3" json:"Seconds,omitempty"`
}

func (m *ToggleSlowModeReq) Reset()         { *m = ToggleSlowModeReq{} }
func (m *ToggleSlowModeReq) String() string { return proto.CompactTextString(m) }
func (*ToggleSlowModeReq) ProtoMessage()    {}
func (*ToggleSlowModeReq) Descriptor() ([]byte, []int) {
	return fileDescriptorImslowmodeserver, []int{0}
}

func (m *ToggleSlowModeReq) GetChnlId() int32 {
	if m != nil {
		return m.ChnlId
	}
	return 0
}

func (m *ToggleSlowModeReq) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

type ToggleSlowModeRsp struct {
	ResCode int32  `protobuf:"varint,1,opt,name=Res_code,json=ResCode,proto3" json:"Res_code,omitempty"`
	ResInfo string `protobuf:"bytes,2,opt,name=Res_info,json=ResInfo,proto3" json:"Res_info,omitempty"`
}

func (m *ToggleSlowModeRsp) Reset()         { *m = ToggleSlowModeRsp{} }
func (m *ToggleSlowModeRsp) String() string { return proto.CompactTextString(m) }
func (*ToggleSlowModeRsp) ProtoMessage()    {}
func (*ToggleSlowModeRsp) Descriptor() ([]byte, []int) {
	return fileDescriptorImslowmodeserver, []int{1}
}

func (m *ToggleSlowModeRsp) GetResCode() int32 {
	if m != nil {
		return m.ResCode
	}
	return 0
}

func (m *ToggleSlowModeRsp) GetResInfo() string {
	if m != nil {
		return m.ResInfo
	}
	return ""
}

type CheckSlowLimitReq struct {
	ChnlId  int32 `protobuf:"varint,1,opt,name=Chnl_id,json=ChnlId,proto3" json:"Chnl_id,omitempty"`
	MembId  int32 `protobuf:"varint,2,opt,name=Memb_id,json=MembId,proto3" json:"Memb_id,omitempty"`
	MsgTime int64 `protobuf:"varint,3,opt,name=MsgTime,proto3" json:"MsgTime,omitempty"`
	Seconds int32 `protobuf:"varint,4,opt,name=Seconds,proto3" json:"Seconds,omitempty"`
}

func (m *CheckSlowLimitReq) Reset()         { *m = CheckSlowLimitReq{} }
func (m *CheckSlowLimitReq) String() string { return proto.CompactTextString(m) }
func (*CheckSlowLimitReq) ProtoMessage()    {}
func (*CheckSlowLimitReq) Descriptor() ([]byte, []int) {
	return fileDescriptorImslowmodeserver, []int{2}
}

func (m *CheckSlowLimitReq) GetChnlId() int32 {
	if m != nil {
		return m.ChnlId
	}
	return 0
}

func (m *CheckSlowLimitReq) GetMembId() int32 {
	if m != nil {
		return m.MembId
	}
	return 0
}

func (m *CheckSlowLimitReq) GetMsgTime() int64 {
	if m != nil {
		return m.MsgTime
	}
	return 0
}

func (m *CheckSlowLimitReq) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

type CheckSlowLimitRsp struct {
	Is_Limit bool  `protobuf:"varint,1,opt,name=Is_Limit,json=IsLimit,proto3" json:"Is_Limit,omitempty"`
	NextTime int64 `protobuf:"varint,2,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *CheckSlowLimitRsp) Reset()         { *m = CheckSlowLimitRsp{} }
func (m *CheckSlowLimitRsp) String() string { return proto.CompactTextString(m) }
func (*CheckSlowLimitRsp) ProtoMessage()    {}
func (*CheckSlowLimitRsp) Descriptor() ([]byte, []int) {
	return fileDescriptorImslowmodeserver, []int{3}
}

func (m *CheckSlowLimitRsp) GetIs_Limit() bool {
	if m != nil {
		return m.Is_Limit
	}
	return false
}

func (m *CheckSlowLimitRsp) GetNextTime() int64 {
	if m != nil {
		return m.NextTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ToggleSlowModeReq)(nil), "imslowmodeserver.ToggleSlowModeReq")
	proto.RegisterType((*ToggleSlowModeRsp)(nil), "imslowmodeserver.ToggleSlowModeRsp")
	proto.RegisterType((*CheckSlowLimitReq)(nil), "imslowmodeserver.CheckSlowLimitReq")
	proto.RegisterType((*CheckSlowLimitRsp)(nil), "imslowmodeserver.CheckSlowLimitRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServerService service

type ServerServiceClient interface {
	// 切换慢速模式
	ToggleSlowMode(ctx context.Context, in *ToggleSlowModeReq, opts ...grpc.CallOption) (*ToggleSlowModeRsp, error)
	// 检查慢速限制
	CheckSlowLimit(ctx context.Context, in *CheckSlowLimitReq, opts ...grpc.CallOption) (*CheckSlowLimitRsp, error)
}

type serverServiceClient struct {
	cc *grpc.ClientConn
}

func NewServerServiceClient(cc *grpc.ClientConn) ServerServiceClient {
	return &serverServiceClient{cc}
}

func (c *serverServiceClient) ToggleSlowMode(ctx context.Context, in *ToggleSlowModeReq, opts ...grpc.CallOption) (*ToggleSlowModeRsp, error) {
	out := new(ToggleSlowModeRsp)
	err := grpc.Invoke(ctx, "/imslowmodeserver.ServerService/ToggleSlowMode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) CheckSlowLimit(ctx context.Context, in *CheckSlowLimitReq, opts ...grpc.CallOption) (*CheckSlowLimitRsp, error) {
	out := new(CheckSlowLimitRsp)
	err := grpc.Invoke(ctx, "/imslowmodeserver.ServerService/CheckSlowLimit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServerService service

type ServerServiceServer interface {
	// 切换慢速模式
	ToggleSlowMode(context.Context, *ToggleSlowModeReq) (*ToggleSlowModeRsp, error)
	// 检查慢速限制
	CheckSlowLimit(context.Context, *CheckSlowLimitReq) (*CheckSlowLimitRsp, error)
}

func RegisterServerServiceServer(s *grpc.Server, srv ServerServiceServer) {
	s.RegisterService(&_ServerService_serviceDesc, srv)
}

func _ServerService_ToggleSlowMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleSlowModeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ToggleSlowMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imslowmodeserver.ServerService/ToggleSlowMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ToggleSlowMode(ctx, req.(*ToggleSlowModeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_CheckSlowLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSlowLimitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).CheckSlowLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imslowmodeserver.ServerService/CheckSlowLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).CheckSlowLimit(ctx, req.(*CheckSlowLimitReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imslowmodeserver.ServerService",
	HandlerType: (*ServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToggleSlowMode",
			Handler:    _ServerService_ToggleSlowMode_Handler,
		},
		{
			MethodName: "CheckSlowLimit",
			Handler:    _ServerService_CheckSlowLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/interfaceprobuf/imslowmodeserver/imslowmodeserver.proto",
}

func (m *ToggleSlowModeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToggleSlowModeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ChnlId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.ChnlId))
	}
	if m.Seconds != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.Seconds))
	}
	return i, nil
}

func (m *ToggleSlowModeRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToggleSlowModeRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ResCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.ResCode))
	}
	if len(m.ResInfo) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(len(m.ResInfo)))
		i += copy(dAtA[i:], m.ResInfo)
	}
	return i, nil
}

func (m *CheckSlowLimitReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckSlowLimitReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ChnlId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.ChnlId))
	}
	if m.MembId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.MembId))
	}
	if m.MsgTime != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.MsgTime))
	}
	if m.Seconds != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.Seconds))
	}
	return i, nil
}

func (m *CheckSlowLimitRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckSlowLimitRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Is_Limit {
		dAtA[i] = 0x8
		i++
		if m.Is_Limit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NextTime != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintImslowmodeserver(dAtA, i, uint64(m.NextTime))
	}
	return i, nil
}

func encodeVarintImslowmodeserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ToggleSlowModeReq) Size() (n int) {
	var l int
	_ = l
	if m.ChnlId != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.ChnlId))
	}
	if m.Seconds != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.Seconds))
	}
	return n
}

func (m *ToggleSlowModeRsp) Size() (n int) {
	var l int
	_ = l
	if m.ResCode != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.ResCode))
	}
	l = len(m.ResInfo)
	if l > 0 {
		n += 1 + l + sovImslowmodeserver(uint64(l))
	}
	return n
}

func (m *CheckSlowLimitReq) Size() (n int) {
	var l int
	_ = l
	if m.ChnlId != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.ChnlId))
	}
	if m.MembId != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.MembId))
	}
	if m.MsgTime != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.MsgTime))
	}
	if m.Seconds != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.Seconds))
	}
	return n
}

func (m *CheckSlowLimitRsp) Size() (n int) {
	var l int
	_ = l
	if m.Is_Limit {
		n += 2
	}
	if m.NextTime != 0 {
		n += 1 + sovImslowmodeserver(uint64(m.NextTime))
	}
	return n
}

func sovImslowmodeserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImslowmodeserver(x uint64) (n int) {
	return sovImslowmodeserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ToggleSlowModeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImslowmodeserver
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
			return fmt.Errorf("proto: ToggleSlowModeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToggleSlowModeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChnlId", wireType)
			}
			m.ChnlId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChnlId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImslowmodeserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImslowmodeserver
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
func (m *ToggleSlowModeRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImslowmodeserver
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
			return fmt.Errorf("proto: ToggleSlowModeRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToggleSlowModeRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResCode", wireType)
			}
			m.ResCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
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
				return ErrInvalidLengthImslowmodeserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImslowmodeserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImslowmodeserver
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
func (m *CheckSlowLimitReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImslowmodeserver
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
			return fmt.Errorf("proto: CheckSlowLimitReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckSlowLimitReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChnlId", wireType)
			}
			m.ChnlId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChnlId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MembId", wireType)
			}
			m.MembId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MembId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTime", wireType)
			}
			m.MsgTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImslowmodeserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImslowmodeserver
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
func (m *CheckSlowLimitRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImslowmodeserver
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
			return fmt.Errorf("proto: CheckSlowLimitRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckSlowLimitRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Is_Limit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
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
			m.Is_Limit = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextTime", wireType)
			}
			m.NextTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImslowmodeserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImslowmodeserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImslowmodeserver
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
func skipImslowmodeserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImslowmodeserver
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
					return 0, ErrIntOverflowImslowmodeserver
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
					return 0, ErrIntOverflowImslowmodeserver
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
				return 0, ErrInvalidLengthImslowmodeserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImslowmodeserver
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
				next, err := skipImslowmodeserver(dAtA[start:])
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
	ErrInvalidLengthImslowmodeserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImslowmodeserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.chatserver.im/interfaceprobuf/imslowmodeserver/imslowmodeserver.proto", fileDescriptorImslowmodeserver)
}

var fileDescriptorImslowmodeserver = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0x5b, 0x6d, 0xea, 0xa2, 0x62, 0xf7, 0xd2, 0xd8, 0x43, 0x29, 0xf1, 0xd2, 0x53, 0x0a,
	0xfa, 0x07, 0x16, 0x84, 0x48, 0xeb, 0x21, 0xed, 0x51, 0x28, 0x4d, 0x76, 0x92, 0x2e, 0x26, 0x99,
	0x98, 0x8d, 0x56, 0xff, 0x50, 0xf0, 0xe2, 0x27, 0x48, 0xbf, 0x44, 0x76, 0x63, 0xc5, 0xb8, 0x62,
	0x2f, 0x0b, 0x6f, 0xde, 0xf0, 0xde, 0x9b, 0x99, 0xa5, 0x93, 0x58, 0x94, 0xc9, 0x32, 0x70, 0xc3,
	0xd5, 0xb2, 0x94, 0x50, 0x3c, 0x41, 0xe1, 0x8a, 0x74, 0x24, 0xb2, 0x12, 0x8a, 0x68, 0x19, 0x42,
	0x5e, 0x60, 0xf0, 0x18, 0x8d, 0x44, 0x2a, 0x13, 0x5c, 0xa7, 0xc8, 0xa1, 0xea, 0x31, 0x0a, 0x6e,
	0x5e, 0x60, 0x89, 0xec, 0xf4, 0x77, 0xdd, 0xb9, 0xa6, 0x9d, 0x39, 0xc6, 0x71, 0x02, 0xb3, 0x04,
	0xd7, 0x53, 0xe4, 0xe0, 0xc3, 0x03, 0xeb, 0x52, 0x6b, 0xbc, 0xca, 0x92, 0x85, 0xe0, 0x36, 0x19,
	0x90, 0xe1, 0x81, 0xdf, 0x52, 0xd0, 0xe3, 0xcc, 0xa6, 0xd6, 0x0c, 0x42, 0xcc, 0xb8, 0xb4, 0x1b,
	0x9a, 0xd8, 0x42, 0xc7, 0x33, 0x74, 0x64, 0xce, 0xce, 0x68, 0xdb, 0x07, 0xb9, 0x08, 0x91, 0xc3,
	0x97, 0x90, 0xe5, 0x83, 0x1c, 0x23, 0x87, 0x2d, 0x25, 0xb2, 0x08, 0xb5, 0xd4, 0xa1, 0xa6, 0xbc,
	0x2c, 0x42, 0xe7, 0x85, 0x76, 0xc6, 0x2b, 0x08, 0xef, 0x95, 0xd2, 0x44, 0xa4, 0xa2, 0xfc, 0x37,
	0x52, 0x97, 0x5a, 0x53, 0x48, 0x03, 0x45, 0x54, 0x91, 0x5a, 0x0a, 0x56, 0x59, 0xa7, 0x32, 0x9e,
	0x8b, 0x14, 0xec, 0xe6, 0x80, 0x0c, 0x9b, 0xfe, 0x16, 0xfe, 0x9c, 0x62, 0xbf, 0x3e, 0xc5, 0x8d,
	0x61, 0x5d, 0x4d, 0xe1, 0xc9, 0x85, 0x86, 0xda, 0xbb, 0xed, 0x5b, 0x9e, 0xd4, 0x90, 0xf5, 0x68,
	0xfb, 0x16, 0x9e, 0x4b, 0x6d, 0xd2, 0xd0, 0x26, 0xdf, 0xf8, 0xe2, 0x8d, 0xd0, 0xe3, 0x99, 0x5e,
	0xb2, 0x7a, 0x45, 0x08, 0xec, 0x8e, 0x9e, 0xd4, 0x77, 0xc4, 0xce, 0x5d, 0xe3, 0x50, 0xc6, 0x35,
	0x7a, 0xbb, 0x9b, 0x64, 0xee, 0xec, 0x29, 0xf5, 0x7a, 0xf6, 0xbf, 0xd4, 0x8d, 0xc5, 0xf6, 0x76,
	0x37, 0x29, 0xf5, 0xab, 0xa3, 0xd7, 0x4d, 0x9f, 0xbc, 0x6f, 0xfa, 0xe4, 0x63, 0xd3, 0x27, 0x41,
	0x4b, 0x7f, 0xa7, 0xcb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x3a, 0x43, 0xcc, 0x9e, 0x02,
	0x00, 0x00,
}
