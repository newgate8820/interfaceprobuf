// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pbcomm.proto

/*
Package pbcomm is a generated protocol buffer package.

It is generated from these files:

	pbcomm.proto

It has these top-level messages:

	Debug
*/
package pbcomm

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

// 日志格式结构
type Debug struct {
	Userid     int32  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Ip         string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	KeyId      uint64 `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	SessionId  uint64 `protobuf:"varint,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Layer      int32  `protobuf:"varint,6,opt,name=layer,proto3" json:"layer,omitempty"`
	Ostype     uint32 `protobuf:"varint,7,opt,name=ostype,proto3" json:"ostype,omitempty"`
	CurrentCrc uint32 `protobuf:"varint,8,opt,name=current_crc,json=currentCrc,proto3" json:"current_crc,omitempty"`
	AppVersion string `protobuf:"bytes,9,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
}

func (m *Debug) Reset()                    { *m = Debug{} }
func (m *Debug) String() string            { return proto.CompactTextString(m) }
func (*Debug) ProtoMessage()               {}
func (*Debug) Descriptor() ([]byte, []int) { return fileDescriptorPbcomm, []int{0} }

func (m *Debug) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *Debug) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Debug) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Debug) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Debug) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Debug) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *Debug) GetOstype() uint32 {
	if m != nil {
		return m.Ostype
	}
	return 0
}

func (m *Debug) GetCurrentCrc() uint32 {
	if m != nil {
		return m.CurrentCrc
	}
	return 0
}

func (m *Debug) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*Debug)(nil), "pbcomm.Debug")
}
func (m *Debug) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Debug) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Userid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.Userid))
	}
	if len(m.Phone) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(len(m.Phone)))
		i += copy(dAtA[i:], m.Phone)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.KeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.KeyId))
	}
	if m.SessionId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.SessionId))
	}
	if m.Layer != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.Layer))
	}
	if m.Ostype != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.Ostype))
	}
	if m.CurrentCrc != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(m.CurrentCrc))
	}
	if len(m.AppVersion) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintPbcomm(dAtA, i, uint64(len(m.AppVersion)))
		i += copy(dAtA[i:], m.AppVersion)
	}
	return i, nil
}

func encodeVarintPbcomm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Debug) Size() (n int) {
	var l int
	_ = l
	if m.Userid != 0 {
		n += 1 + sovPbcomm(uint64(m.Userid))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovPbcomm(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovPbcomm(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovPbcomm(uint64(m.KeyId))
	}
	if m.SessionId != 0 {
		n += 1 + sovPbcomm(uint64(m.SessionId))
	}
	if m.Layer != 0 {
		n += 1 + sovPbcomm(uint64(m.Layer))
	}
	if m.Ostype != 0 {
		n += 1 + sovPbcomm(uint64(m.Ostype))
	}
	if m.CurrentCrc != 0 {
		n += 1 + sovPbcomm(uint64(m.CurrentCrc))
	}
	l = len(m.AppVersion)
	if l > 0 {
		n += 1 + l + sovPbcomm(uint64(l))
	}
	return n
}

func sovPbcomm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPbcomm(x uint64) (n int) {
	return sovPbcomm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Debug) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPbcomm
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
			return fmt.Errorf("proto: Debug: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Debug: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			m.Userid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcomm
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
					return ErrIntOverflowPbcomm
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
				return ErrInvalidLengthPbcomm
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
					return ErrIntOverflowPbcomm
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
				return ErrInvalidLengthPbcomm
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
					return ErrIntOverflowPbcomm
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
					return ErrIntOverflowPbcomm
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
					return ErrIntOverflowPbcomm
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
					return ErrIntOverflowPbcomm
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentCrc", wireType)
			}
			m.CurrentCrc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcomm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentCrc |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcomm
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
				return ErrInvalidLengthPbcomm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPbcomm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPbcomm
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
func skipPbcomm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPbcomm
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
					return 0, ErrIntOverflowPbcomm
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
					return 0, ErrIntOverflowPbcomm
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
				return 0, ErrInvalidLengthPbcomm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPbcomm
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
				next, err := skipPbcomm(dAtA[start:])
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
	ErrInvalidLengthPbcomm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPbcomm   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pbcomm.proto", fileDescriptorPbcomm) }

var fileDescriptorPbcomm = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x90, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0x86, 0x73, 0xf0, 0xb6, 0xca, 0xf1, 0xea, 0xd0, 0xa8, 0xe9, 0x22, 0x12, 0x27, 0x26, 0x17,
	0xdf, 0x40, 0x5d, 0xee, 0x66, 0x18, 0x5c, 0x09, 0x94, 0x13, 0x6d, 0xae, 0xd0, 0x93, 0x16, 0x4c,
	0x78, 0x43, 0x47, 0x1f, 0xc1, 0xb0, 0xfa, 0x12, 0x86, 0xc2, 0xf8, 0x7d, 0x7f, 0xf3, 0xf7, 0xcf,
	0xc1, 0x3d, 0x37, 0xc6, 0x75, 0xdd, 0x03, 0x7b, 0x37, 0x38, 0x25, 0x57, 0xba, 0xff, 0x03, 0x14,
	0x2f, 0xd4, 0x8c, 0xef, 0xea, 0x06, 0xe5, 0x18, 0xc8, 0xdb, 0x56, 0x43, 0x0e, 0x85, 0x28, 0x37,
	0x52, 0x57, 0x28, 0x5e, 0x3f, 0x5c, 0x4f, 0x3a, 0xc9, 0xa1, 0x48, 0xcb, 0x15, 0xd4, 0x25, 0x26,
	0x96, 0xf5, 0x49, 0x54, 0x89, 0x65, 0x75, 0x8d, 0xf2, 0x48, 0x53, 0x65, 0x5b, 0xbd, 0xcb, 0xa1,
	0xd8, 0x95, 0xe2, 0x48, 0xd3, 0xa1, 0x55, 0xb7, 0x88, 0x81, 0x42, 0xb0, 0xae, 0x5f, 0x22, 0x11,
	0xa3, 0x74, 0x33, 0x87, 0xd8, 0xfd, 0x59, 0x4f, 0xe4, 0xb5, 0x8c, 0x5f, 0xae, 0xb0, 0x2c, 0x71,
	0x61, 0x98, 0x98, 0xf4, 0x69, 0x0e, 0xc5, 0x45, 0xb9, 0x91, 0xba, 0xc3, 0x73, 0x33, 0x7a, 0x4f,
	0xfd, 0x50, 0x19, 0x6f, 0xf4, 0x59, 0x0c, 0x71, 0x53, 0xcf, 0xde, 0x2c, 0x0f, 0x6a, 0xe6, 0xea,
	0x8b, 0xfc, 0xd2, 0xaf, 0xd3, 0xb8, 0x0e, 0x6b, 0xe6, 0xb7, 0xd5, 0x3c, 0xed, 0xbf, 0xe7, 0x0c,
	0x7e, 0xe6, 0x0c, 0x7e, 0xe7, 0x0c, 0x1a, 0x19, 0x4f, 0xf1, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x00, 0xec, 0xde, 0x1f, 0x1a, 0x01, 0x00, 0x00,
}
