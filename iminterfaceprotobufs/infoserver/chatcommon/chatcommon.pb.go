// Code generated by protoc-gen-gogo.
// source: chatcommon.proto
// DO NOT EDIT!

/*
Package chatcommon is a generated protocol buffer package.

It is generated from these files:

	chatcommon.proto

It has these top-level messages:

	ChatLogPrefix
*/
package chatcommon

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ChatResultCode int32

const (
	ChatResultCode_Chat_RPC_OK              ChatResultCode = 0
	ChatResultCode_Chat_PARAMETER_EXCEPTION ChatResultCode = 101
	ChatResultCode_Chat_LOGIC_EXCEPTION     ChatResultCode = 102
	ChatResultCode_Chat_REDIS_EXCEPTION     ChatResultCode = 103
	ChatResultCode_Chat_DB_EXCEPTION        ChatResultCode = 104
	ChatResultCode_Chat_UNKNOWN             ChatResultCode = 65535
)

var ChatResultCode_name = map[int32]string{
	0:     "Chat_RPC_OK",
	101:   "Chat_PARAMETER_EXCEPTION",
	102:   "Chat_LOGIC_EXCEPTION",
	103:   "Chat_REDIS_EXCEPTION",
	104:   "Chat_DB_EXCEPTION",
	65535: "Chat_UNKNOWN",
}
var ChatResultCode_value = map[string]int32{
	"Chat_RPC_OK":              0,
	"Chat_PARAMETER_EXCEPTION": 101,
	"Chat_LOGIC_EXCEPTION":     102,
	"Chat_REDIS_EXCEPTION":     103,
	"Chat_DB_EXCEPTION":        104,
	"Chat_UNKNOWN":             65535,
}

func (x ChatResultCode) String() string {
	return proto.EnumName(ChatResultCode_name, int32(x))
}
func (ChatResultCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorChatcommon, []int{0} }

// 日志需要打印
type ChatLogPrefix struct {
	Userid    int32  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Phone     string `protobuf:"bytes,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Ip        string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	KeyId     uint64 `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	SessionId uint64 `protobuf:"varint,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Layer     int32  `protobuf:"varint,6,opt,name=layer,proto3" json:"layer,omitempty"`
	Ostype    uint32 `protobuf:"varint,7,opt,name=ostype,proto3" json:"ostype,omitempty"`
}

func (m *ChatLogPrefix) Reset()                    { *m = ChatLogPrefix{} }
func (m *ChatLogPrefix) String() string            { return proto.CompactTextString(m) }
func (*ChatLogPrefix) ProtoMessage()               {}
func (*ChatLogPrefix) Descriptor() ([]byte, []int) { return fileDescriptorChatcommon, []int{0} }

func (m *ChatLogPrefix) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *ChatLogPrefix) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChatLogPrefix) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ChatLogPrefix) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *ChatLogPrefix) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *ChatLogPrefix) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *ChatLogPrefix) GetOstype() uint32 {
	if m != nil {
		return m.Ostype
	}
	return 0
}

func init() {
	proto.RegisterType((*ChatLogPrefix)(nil), "chatcommon.ChatLogPrefix")
	proto.RegisterEnum("chatcommon.ChatResultCode", ChatResultCode_name, ChatResultCode_value)
}
func (m *ChatLogPrefix) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatLogPrefix) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Userid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(m.Userid))
	}
	if len(m.Phone) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(len(m.Phone)))
		i += copy(dAtA[i:], m.Phone)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.KeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(m.KeyId))
	}
	if m.SessionId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(m.SessionId))
	}
	if m.Layer != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(m.Layer))
	}
	if m.Ostype != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintChatcommon(dAtA, i, uint64(m.Ostype))
	}
	return i, nil
}

func encodeFixed64Chatcommon(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Chatcommon(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintChatcommon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChatLogPrefix) Size() (n int) {
	var l int
	_ = l
	if m.Userid != 0 {
		n += 1 + sovChatcommon(uint64(m.Userid))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovChatcommon(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovChatcommon(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovChatcommon(uint64(m.KeyId))
	}
	if m.SessionId != 0 {
		n += 1 + sovChatcommon(uint64(m.SessionId))
	}
	if m.Layer != 0 {
		n += 1 + sovChatcommon(uint64(m.Layer))
	}
	if m.Ostype != 0 {
		n += 1 + sovChatcommon(uint64(m.Ostype))
	}
	return n
}

func sovChatcommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChatcommon(x uint64) (n int) {
	return sovChatcommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChatLogPrefix) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatcommon
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
			return fmt.Errorf("proto: ChatLogPrefix: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatLogPrefix: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			m.Userid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Userid |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
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
				return ErrInvalidLengthChatcommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
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
				return ErrInvalidLengthChatcommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layer", wireType)
			}
			m.Layer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Layer |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ostype", wireType)
			}
			m.Ostype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatcommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ostype |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChatcommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatcommon
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
func skipChatcommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChatcommon
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
					return 0, ErrIntOverflowChatcommon
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
					return 0, ErrIntOverflowChatcommon
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
				return 0, ErrInvalidLengthChatcommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChatcommon
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
				next, err := skipChatcommon(dAtA[start:])
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
	ErrInvalidLengthChatcommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChatcommon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("chatcommon.proto", fileDescriptorChatcommon) }

var fileDescriptorChatcommon = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0x87, 0x1d, 0xa0, 0x35, 0x3c, 0x01, 0xeb, 0x0b, 0x98, 0x59, 0x28, 0x69, 0x5c, 0x35, 0x2e,
	0xdc, 0x78, 0x02, 0x28, 0x8d, 0x69, 0xc0, 0xb6, 0x19, 0x31, 0xba, 0x6b, 0x90, 0x0e, 0xb4, 0xe1,
	0xcf, 0x34, 0x9d, 0x92, 0xd8, 0xcb, 0x78, 0x07, 0x6f, 0xe1, 0xd2, 0x23, 0x18, 0x2e, 0x82, 0xe9,
	0x40, 0x94, 0xe5, 0xf7, 0x7d, 0xc9, 0xcc, 0x2f, 0x0f, 0x8c, 0x69, 0x3c, 0xc9, 0xa7, 0x62, 0xb5,
	0x12, 0xeb, 0xbb, 0x34, 0x13, 0xb9, 0x40, 0xf8, 0x37, 0x37, 0x9f, 0x04, 0x9a, 0x76, 0x3c, 0xc9,
	0x47, 0x62, 0x1e, 0x64, 0x7c, 0x96, 0xbc, 0xe3, 0x25, 0xe8, 0x1b, 0xc9, 0xb3, 0x24, 0xa2, 0xc4,
	0x24, 0x96, 0xc6, 0x0e, 0x84, 0x6d, 0xd0, 0x82, 0x58, 0xac, 0x39, 0xad, 0x98, 0xc4, 0xaa, 0xb3,
	0x3d, 0x60, 0x0b, 0x2a, 0x49, 0x4a, 0xab, 0x4a, 0x55, 0x92, 0x14, 0x3b, 0xa0, 0x2f, 0x78, 0x11,
	0x26, 0x11, 0xad, 0x99, 0xc4, 0xaa, 0x31, 0x6d, 0xc1, 0x0b, 0x37, 0xc2, 0x6b, 0x00, 0xc9, 0xa5,
	0x4c, 0xc4, 0xba, 0x4c, 0x9a, 0x4a, 0xf5, 0x83, 0x71, 0xd5, 0xdb, 0xcb, 0x49, 0xc1, 0x33, 0xaa,
	0xab, 0x2f, 0xf7, 0x50, 0x2e, 0x11, 0x32, 0x2f, 0x52, 0x4e, 0x4f, 0x4d, 0x62, 0x35, 0xd9, 0x81,
	0x6e, 0x3f, 0x08, 0xb4, 0xca, 0xcd, 0x8c, 0xcb, 0xcd, 0x32, 0xb7, 0x45, 0xc4, 0xf1, 0x1c, 0xce,
	0x4a, 0x13, 0xb2, 0xc0, 0x0e, 0xfd, 0xa1, 0x71, 0x82, 0x57, 0x40, 0x95, 0x08, 0x7a, 0xac, 0xf7,
	0xe8, 0x8c, 0x1d, 0x16, 0x3a, 0xaf, 0xb6, 0x13, 0x8c, 0x5d, 0xdf, 0x33, 0x38, 0x52, 0x68, 0xab,
	0x3a, 0xf2, 0x1f, 0x5c, 0xfb, 0xa8, 0xcc, 0xfe, 0x0a, 0x73, 0x06, 0xee, 0xd3, 0x51, 0x99, 0x63,
	0x07, 0x2e, 0x54, 0x19, 0xf4, 0x8f, 0x74, 0x8c, 0x08, 0x0d, 0xa5, 0x9f, 0xbd, 0xa1, 0xe7, 0xbf,
	0x78, 0xc6, 0x6e, 0x57, 0xed, 0x37, 0xbe, 0xb6, 0x5d, 0xf2, 0xbd, 0xed, 0x92, 0x9f, 0x6d, 0x97,
	0xbc, 0xe9, 0xea, 0xea, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x69, 0x9e, 0x93, 0x89,
	0x01, 0x00, 0x00,
}