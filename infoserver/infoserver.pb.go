// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitlab.chatserver.im/interfaceprobuf/infoserver/infoserver.proto

/*
Package infoserver is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/interfaceprobuf/infoserver/infoserver.proto

It has these top-level messages:

	MomentState
	ReqUpdateMoments
	ResUpdateMoments
*/
package infoserver

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

type MomentState struct {
	UserId []int32 `protobuf:"varint,1,rep,packed,name=userId" json:"userId,omitempty"`
	Type   int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *MomentState) Reset()                    { *m = MomentState{} }
func (m *MomentState) String() string            { return proto.CompactTextString(m) }
func (*MomentState) ProtoMessage()               {}
func (*MomentState) Descriptor() ([]byte, []int) { return fileDescriptorInfoserver, []int{0} }

func (m *MomentState) GetUserId() []int32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *MomentState) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqUpdateMoments struct {
	MomentStates []*MomentState `protobuf:"bytes,1,rep,name=momentStates" json:"momentStates,omitempty"`
	SelfId       int32          `protobuf:"varint,2,opt,name=selfId,proto3" json:"selfId,omitempty"`
}

func (m *ReqUpdateMoments) Reset()                    { *m = ReqUpdateMoments{} }
func (m *ReqUpdateMoments) String() string            { return proto.CompactTextString(m) }
func (*ReqUpdateMoments) ProtoMessage()               {}
func (*ReqUpdateMoments) Descriptor() ([]byte, []int) { return fileDescriptorInfoserver, []int{1} }

func (m *ReqUpdateMoments) GetMomentStates() []*MomentState {
	if m != nil {
		return m.MomentStates
	}
	return nil
}

func (m *ReqUpdateMoments) GetSelfId() int32 {
	if m != nil {
		return m.SelfId
	}
	return 0
}

type ResUpdateMoments struct {
	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ResUpdateMoments) Reset()                    { *m = ResUpdateMoments{} }
func (m *ResUpdateMoments) String() string            { return proto.CompactTextString(m) }
func (*ResUpdateMoments) ProtoMessage()               {}
func (*ResUpdateMoments) Descriptor() ([]byte, []int) { return fileDescriptorInfoserver, []int{2} }

func (m *ResUpdateMoments) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResUpdateMoments) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*MomentState)(nil), "infoserver.MomentState")
	proto.RegisterType((*ReqUpdateMoments)(nil), "infoserver.ReqUpdateMoments")
	proto.RegisterType((*ResUpdateMoments)(nil), "infoserver.ResUpdateMoments")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContactPushService service

type ContactPushServiceClient interface {
	// 更新自己朋友的未读状态
	UpdateMoments(ctx context.Context, in *ReqUpdateMoments, opts ...grpc.CallOption) (*ResUpdateMoments, error)
}

type contactPushServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactPushServiceClient(cc *grpc.ClientConn) ContactPushServiceClient {
	return &contactPushServiceClient{cc}
}

func (c *contactPushServiceClient) UpdateMoments(ctx context.Context, in *ReqUpdateMoments, opts ...grpc.CallOption) (*ResUpdateMoments, error) {
	out := new(ResUpdateMoments)
	err := grpc.Invoke(ctx, "/infoserver.ContactPushService/UpdateMoments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContactPushService service

type ContactPushServiceServer interface {
	// 更新自己朋友的未读状态
	UpdateMoments(context.Context, *ReqUpdateMoments) (*ResUpdateMoments, error)
}

func RegisterContactPushServiceServer(s *grpc.Server, srv ContactPushServiceServer) {
	s.RegisterService(&_ContactPushService_serviceDesc, srv)
}

func _ContactPushService_UpdateMoments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateMoments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactPushServiceServer).UpdateMoments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infoserver.ContactPushService/UpdateMoments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactPushServiceServer).UpdateMoments(ctx, req.(*ReqUpdateMoments))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactPushService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infoserver.ContactPushService",
	HandlerType: (*ContactPushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMoments",
			Handler:    _ContactPushService_UpdateMoments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/interfaceprobuf/infoserver/infoserver.proto",
}

func (m *MomentState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MomentState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA2 := make([]byte, len(m.UserId)*10)
		var j1 int
		for _, num1 := range m.UserId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintInfoserver(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintInfoserver(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func (m *ReqUpdateMoments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqUpdateMoments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MomentStates) > 0 {
		for _, msg := range m.MomentStates {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInfoserver(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.SelfId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintInfoserver(dAtA, i, uint64(m.SelfId))
	}
	return i, nil
}

func (m *ResUpdateMoments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResUpdateMoments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInfoserver(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInfoserver(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func encodeVarintInfoserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MomentState) Size() (n int) {
	var l int
	_ = l
	if len(m.UserId) > 0 {
		l = 0
		for _, e := range m.UserId {
			l += sovInfoserver(uint64(e))
		}
		n += 1 + sovInfoserver(uint64(l)) + l
	}
	if m.Type != 0 {
		n += 1 + sovInfoserver(uint64(m.Type))
	}
	return n
}

func (m *ReqUpdateMoments) Size() (n int) {
	var l int
	_ = l
	if len(m.MomentStates) > 0 {
		for _, e := range m.MomentStates {
			l = e.Size()
			n += 1 + l + sovInfoserver(uint64(l))
		}
	}
	if m.SelfId != 0 {
		n += 1 + sovInfoserver(uint64(m.SelfId))
	}
	return n
}

func (m *ResUpdateMoments) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovInfoserver(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovInfoserver(uint64(l))
	}
	return n
}

func sovInfoserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfoserver(x uint64) (n int) {
	return sovInfoserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MomentState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfoserver
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
			return fmt.Errorf("proto: MomentState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MomentState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInfoserver
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UserId = append(m.UserId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInfoserver
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthInfoserver
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowInfoserver
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.UserId = append(m.UserId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfoserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfoserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfoserver
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
func (m *ReqUpdateMoments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfoserver
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
			return fmt.Errorf("proto: ReqUpdateMoments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqUpdateMoments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MomentStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfoserver
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
				return ErrInvalidLengthInfoserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MomentStates = append(m.MomentStates, &MomentState{})
			if err := m.MomentStates[len(m.MomentStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfId", wireType)
			}
			m.SelfId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfoserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelfId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfoserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfoserver
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
func (m *ResUpdateMoments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfoserver
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
			return fmt.Errorf("proto: ResUpdateMoments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResUpdateMoments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfoserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfoserver
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
				return ErrInvalidLengthInfoserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfoserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfoserver
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
func skipInfoserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfoserver
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
					return 0, ErrIntOverflowInfoserver
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
					return 0, ErrIntOverflowInfoserver
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
				return 0, ErrInvalidLengthInfoserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfoserver
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
				next, err := skipInfoserver(dAtA[start:])
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
	ErrInvalidLengthInfoserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfoserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.chatserver.im/interfaceprobuf/infoserver/infoserver.proto", fileDescriptorInfoserver)
}

var fileDescriptorInfoserver = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0x7f, 0xfd, 0xb9, 0x40, 0x27, 0x29, 0x04, 0x2d, 0x5a, 0x53, 0x8a, 0x09, 0x5e, 0x65,
	0xe5, 0x40, 0xba, 0x69, 0xe9, 0xa6, 0xb4, 0xab, 0x2c, 0x02, 0x45, 0xa1, 0x0f, 0x20, 0xcb, 0x63,
	0x47, 0x10, 0x5b, 0xae, 0x34, 0x0e, 0xf4, 0x0d, 0xbb, 0xec, 0x23, 0x14, 0x3f, 0x49, 0xf1, 0x05,
	0x62, 0x67, 0x77, 0x46, 0x47, 0x3a, 0xdf, 0xd1, 0xc0, 0x4b, 0xaa, 0xe9, 0x28, 0xa3, 0x50, 0x1d,
	0x24, 0x39, 0xb4, 0x27, 0xb4, 0xa1, 0xce, 0xd6, 0x3a, 0x27, 0xb4, 0x89, 0x54, 0x58, 0x58, 0x13,
	0x95, 0xc9, 0x5a, 0xe7, 0x89, 0x69, 0xdd, 0x9e, 0x0c, 0x0b, 0x6b, 0xc8, 0x70, 0x38, 0x9f, 0x04,
	0x4f, 0x30, 0xdb, 0x99, 0x0c, 0x73, 0xda, 0x93, 0x24, 0xe4, 0x37, 0x30, 0x2d, 0x1d, 0xda, 0x6d,
	0xec, 0xb1, 0xe5, 0x68, 0x35, 0x11, 0xdd, 0xc4, 0x39, 0x8c, 0xe9, 0xab, 0x40, 0xef, 0xff, 0x92,
	0xad, 0x26, 0xa2, 0xd1, 0x41, 0x0a, 0x0b, 0x81, 0x9f, 0x1f, 0x45, 0x2c, 0x09, 0xdb, 0x0c, 0xc7,
	0x9f, 0x61, 0x9e, 0x9d, 0xe3, 0x5c, 0x93, 0x32, 0xdb, 0xdc, 0x86, 0xbd, 0x0e, 0x3d, 0x9c, 0x18,
	0x5c, 0xae, 0xe1, 0x0e, 0x8f, 0xc9, 0x36, 0xee, 0x30, 0xdd, 0x14, 0x3c, 0xd6, 0x20, 0x37, 0x04,
	0x71, 0x18, 0x2b, 0x13, 0xa3, 0xc7, 0xda, 0x42, 0xb5, 0xe6, 0x0b, 0x18, 0x65, 0x2e, 0x6d, 0x1e,
	0x5f, 0x89, 0x5a, 0x6e, 0x14, 0xf0, 0x37, 0x93, 0x93, 0x54, 0xf4, 0x5e, 0xba, 0xc3, 0x1e, 0xed,
	0x49, 0x2b, 0xe4, 0x3b, 0xb8, 0x1e, 0x86, 0xdd, 0xf7, 0xfb, 0x5d, 0xfe, 0xe9, 0xee, 0xc2, 0x1d,
	0x16, 0x09, 0xfe, 0xbd, 0xce, 0xbf, 0x2b, 0x9f, 0xfd, 0x54, 0x3e, 0xfb, 0xad, 0x7c, 0x16, 0x4d,
	0x9b, 0x1d, 0x3f, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x88, 0x15, 0x75, 0xa7, 0x01, 0x00,
	0x00,
}