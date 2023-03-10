// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitlab.chatserver.im/interfaceprobuf/pbmoments/media/media.proto

/*
Package __moment is a generated protocol buffer package.

It is generated from these files:

	gitlab.chatserver.im/interfaceprobuf/pbmoments/media/media.proto

It has these top-level messages:

	HighLight
	MediaGetRep
	Media
*/
package __moment

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

type HighLightType int32

const (
	HighLightType_HighLightType_UNKNOWN  HighLightType = 0
	HighLightType_HighLightType_USERNAME HighLightType = 1
	HighLightType_HighLightType_URL      HighLightType = 2
)

var HighLightType_name = map[int32]string{
	0: "HighLightType_UNKNOWN",
	1: "HighLightType_USERNAME",
	2: "HighLightType_URL",
}
var HighLightType_value = map[string]int32{
	"HighLightType_UNKNOWN":  0,
	"HighLightType_USERNAME": 1,
	"HighLightType_URL":      2,
}

func (x HighLightType) String() string {
	return proto.EnumName(HighLightType_name, int32(x))
}
func (HighLightType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMedia, []int{0} }

type HighLight struct {
	Type       HighLightType `protobuf:"varint,1,opt,name=Type,proto3,enum=moment.HighLightType" json:"Type,omitempty"`
	Offset     int32         `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	UserID     int32         `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName   string        `protobuf:"bytes,4,opt,name=UserName,proto3" json:"UserName,omitempty"`
	AccessHash uint64        `protobuf:"varint,5,opt,name=AccessHash,proto3" json:"AccessHash,omitempty"`
	Limit      int32         `protobuf:"varint,6,opt,name=Limit,proto3" json:"Limit,omitempty"`
	UOffset    int32         `protobuf:"varint,7,opt,name=UOffset,proto3" json:"UOffset,omitempty"`
	ULimit     int32         `protobuf:"varint,8,opt,name=ULimit,proto3" json:"ULimit,omitempty"`
}

func (m *HighLight) Reset()                    { *m = HighLight{} }
func (m *HighLight) String() string            { return proto.CompactTextString(m) }
func (*HighLight) ProtoMessage()               {}
func (*HighLight) Descriptor() ([]byte, []int) { return fileDescriptorMedia, []int{0} }

func (m *HighLight) GetType() HighLightType {
	if m != nil {
		return m.Type
	}
	return HighLightType_HighLightType_UNKNOWN
}

func (m *HighLight) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *HighLight) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *HighLight) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *HighLight) GetAccessHash() uint64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func (m *HighLight) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *HighLight) GetUOffset() int32 {
	if m != nil {
		return m.UOffset
	}
	return 0
}

func (m *HighLight) GetULimit() int32 {
	if m != nil {
		return m.ULimit
	}
	return 0
}

type MediaGetRep struct {
	Medias []*Media `protobuf:"bytes,1,rep,name=Medias" json:"Medias,omitempty"`
}

func (m *MediaGetRep) Reset()                    { *m = MediaGetRep{} }
func (m *MediaGetRep) String() string            { return proto.CompactTextString(m) }
func (*MediaGetRep) ProtoMessage()               {}
func (*MediaGetRep) Descriptor() ([]byte, []int) { return fileDescriptorMedia, []int{1} }

func (m *MediaGetRep) GetMedias() []*Media {
	if m != nil {
		return m.Medias
	}
	return nil
}

type Media struct {
	// @inject_tag: db:"id"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// ForumID ????????????
	// @inject_tag: db:"main_id"
	MainID int64 `protobuf:"varint,2,opt,name=MainID,proto3" json:"MainID,omitempty"`
	// Seq ????????????
	// @inject_tag: db:"seq"
	Seq int64 `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
	// Name ?????????
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	// Ext ???????????? 1=????????????(bmp???jpg???jpeg???png) 2=??????(gif???mp4???) 3=?????? 4=DOC??????pdf,txt, doc,docx, xls, ppt ??? 5=????????????(zip,rar,tar) 6=????????????????????????URL 7=??????
	// @inject_tag: db:"ext"
	Ext int32 `protobuf:"varint,5,opt,name=Ext,proto3" json:"Ext,omitempty"`
	// Thum ???????????????
	// @inject_tag: db:"thum"
	Thum string `protobuf:"bytes,6,opt,name=Thum,proto3" json:"Thum,omitempty"`
	// Region s3 ?????? ?????????????????????????????????
	// @inject_tag: db:"region"
	Region string `protobuf:"bytes,7,opt,name=Region,proto3" json:"Region,omitempty"`
	// Size ????????????
	// @inject_tag: db:"size"
	Size_ int32 `protobuf:"varint,8,opt,name=Size,proto3" json:"Size,omitempty"`
	// ThumSize ???????????????
	// @inject_tag: db:"thum_size"
	ThumSize int32 `protobuf:"varint,9,opt,name=ThumSize,proto3" json:"ThumSize,omitempty"`
	// Hash ????????????????????????hash ???
	// @inject_tag: db:"hash"
	Hash string `protobuf:"bytes,10,opt,name=Hash,proto3" json:"Hash,omitempty"`
	// Duration ??????????????????
	// @inject_tag: db:"duration"
	Duration int32 `protobuf:"varint,11,opt,name=Duration,proto3" json:"Duration,omitempty"`
	// Height ????????????
	// @inject_tag: db:"height"
	Height int32 `protobuf:"varint,12,opt,name=Height,proto3" json:"Height,omitempty"`
	// Width ????????????
	// @inject_tag: db:"width"
	Width int32 `protobuf:"varint,13,opt,name=Width,proto3" json:"Width,omitempty"`
	// CreateTime ????????????
	// @inject_tag: db:"create_at"
	CreateAt int64 `protobuf:"varint,14,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	// @inject_tag: db:"status"
	Status int32 `protobuf:"varint,15,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: db:"delete_at"
	DeleteAt int64 `protobuf:"varint,16,opt,name=DeleteAt,proto3" json:"DeleteAt,omitempty"`
}

func (m *Media) Reset()                    { *m = Media{} }
func (m *Media) String() string            { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()               {}
func (*Media) Descriptor() ([]byte, []int) { return fileDescriptorMedia, []int{2} }

func (m *Media) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Media) GetMainID() int64 {
	if m != nil {
		return m.MainID
	}
	return 0
}

func (m *Media) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Media) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Media) GetExt() int32 {
	if m != nil {
		return m.Ext
	}
	return 0
}

func (m *Media) GetThum() string {
	if m != nil {
		return m.Thum
	}
	return ""
}

func (m *Media) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Media) GetSize_() int32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Media) GetThumSize() int32 {
	if m != nil {
		return m.ThumSize
	}
	return 0
}

func (m *Media) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Media) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Media) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Media) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Media) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *Media) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Media) GetDeleteAt() int64 {
	if m != nil {
		return m.DeleteAt
	}
	return 0
}

func init() {
	proto.RegisterType((*HighLight)(nil), "moment.HighLight")
	proto.RegisterType((*MediaGetRep)(nil), "moment.MediaGetRep")
	proto.RegisterType((*Media)(nil), "moment.Media")
	proto.RegisterEnum("moment.HighLightType", HighLightType_name, HighLightType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MediaDBService service

type MediaDBServiceClient interface {
	Get(ctx context.Context, in *Media, opts ...grpc.CallOption) (*MediaGetRep, error)
}

type mediaDBServiceClient struct {
	cc *grpc.ClientConn
}

func NewMediaDBServiceClient(cc *grpc.ClientConn) MediaDBServiceClient {
	return &mediaDBServiceClient{cc}
}

func (c *mediaDBServiceClient) Get(ctx context.Context, in *Media, opts ...grpc.CallOption) (*MediaGetRep, error) {
	out := new(MediaGetRep)
	err := grpc.Invoke(ctx, "/moment.MediaDBService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MediaDBService service

type MediaDBServiceServer interface {
	Get(context.Context, *Media) (*MediaGetRep, error)
}

func RegisterMediaDBServiceServer(s *grpc.Server, srv MediaDBServiceServer) {
	s.RegisterService(&_MediaDBService_serviceDesc, srv)
}

func _MediaDBService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Media)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaDBServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MediaDBService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaDBServiceServer).Get(ctx, req.(*Media))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaDBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moment.MediaDBService",
	HandlerType: (*MediaDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MediaDBService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.chatserver.im/interfaceprobuf/pbmoments/media/media.proto",
}

func (m *HighLight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HighLight) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Type))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Offset))
	}
	if m.UserID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.UserID))
	}
	if len(m.UserName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMedia(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if m.AccessHash != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.AccessHash))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Limit))
	}
	if m.UOffset != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.UOffset))
	}
	if m.ULimit != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.ULimit))
	}
	return i, nil
}

func (m *MediaGetRep) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MediaGetRep) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Medias) > 0 {
		for _, msg := range m.Medias {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMedia(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Media) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Media) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.ID))
	}
	if m.MainID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.MainID))
	}
	if m.Seq != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Seq))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Ext != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Ext))
	}
	if len(m.Thum) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Thum)))
		i += copy(dAtA[i:], m.Thum)
	}
	if len(m.Region) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Region)))
		i += copy(dAtA[i:], m.Region)
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Size_))
	}
	if m.ThumSize != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.ThumSize))
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.Duration != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Duration))
	}
	if m.Height != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Height))
	}
	if m.Width != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Width))
	}
	if m.CreateAt != 0 {
		dAtA[i] = 0x70
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.CreateAt))
	}
	if m.Status != 0 {
		dAtA[i] = 0x78
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.Status))
	}
	if m.DeleteAt != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintMedia(dAtA, i, uint64(m.DeleteAt))
	}
	return i, nil
}

func encodeVarintMedia(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HighLight) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMedia(uint64(m.Type))
	}
	if m.Offset != 0 {
		n += 1 + sovMedia(uint64(m.Offset))
	}
	if m.UserID != 0 {
		n += 1 + sovMedia(uint64(m.UserID))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.AccessHash != 0 {
		n += 1 + sovMedia(uint64(m.AccessHash))
	}
	if m.Limit != 0 {
		n += 1 + sovMedia(uint64(m.Limit))
	}
	if m.UOffset != 0 {
		n += 1 + sovMedia(uint64(m.UOffset))
	}
	if m.ULimit != 0 {
		n += 1 + sovMedia(uint64(m.ULimit))
	}
	return n
}

func (m *MediaGetRep) Size() (n int) {
	var l int
	_ = l
	if len(m.Medias) > 0 {
		for _, e := range m.Medias {
			l = e.Size()
			n += 1 + l + sovMedia(uint64(l))
		}
	}
	return n
}

func (m *Media) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovMedia(uint64(m.ID))
	}
	if m.MainID != 0 {
		n += 1 + sovMedia(uint64(m.MainID))
	}
	if m.Seq != 0 {
		n += 1 + sovMedia(uint64(m.Seq))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.Ext != 0 {
		n += 1 + sovMedia(uint64(m.Ext))
	}
	l = len(m.Thum)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovMedia(uint64(m.Size_))
	}
	if m.ThumSize != 0 {
		n += 1 + sovMedia(uint64(m.ThumSize))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovMedia(uint64(m.Duration))
	}
	if m.Height != 0 {
		n += 1 + sovMedia(uint64(m.Height))
	}
	if m.Width != 0 {
		n += 1 + sovMedia(uint64(m.Width))
	}
	if m.CreateAt != 0 {
		n += 1 + sovMedia(uint64(m.CreateAt))
	}
	if m.Status != 0 {
		n += 1 + sovMedia(uint64(m.Status))
	}
	if m.DeleteAt != 0 {
		n += 2 + sovMedia(uint64(m.DeleteAt))
	}
	return n
}

func sovMedia(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMedia(x uint64) (n int) {
	return sovMedia(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HighLight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMedia
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
			return fmt.Errorf("proto: HighLight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HighLight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (HighLightType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UOffset", wireType)
			}
			m.UOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UOffset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ULimit", wireType)
			}
			m.ULimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ULimit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMedia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMedia
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
func (m *MediaGetRep) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMedia
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
			return fmt.Errorf("proto: MediaGetRep: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MediaGetRep: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Medias", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Medias = append(m.Medias, &Media{})
			if err := m.Medias[len(m.Medias)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMedia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMedia
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
func (m *Media) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMedia
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
			return fmt.Errorf("proto: Media: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Media: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainID", wireType)
			}
			m.MainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			m.Ext = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ext |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Thum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Thum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThumSize", wireType)
			}
			m.ThumSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThumSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteAt", wireType)
			}
			m.DeleteAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeleteAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMedia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMedia
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
func skipMedia(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMedia
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
					return 0, ErrIntOverflowMedia
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
					return 0, ErrIntOverflowMedia
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
				return 0, ErrInvalidLengthMedia
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMedia
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
				next, err := skipMedia(dAtA[start:])
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
	ErrInvalidLengthMedia = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMedia   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.chatserver.im/interfaceprobuf/pbmoments/media/media.proto", fileDescriptorMedia)
}

var fileDescriptorMedia = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x51, 0x6b, 0x1a, 0x41,
	0x10, 0xc7, 0xb3, 0x9e, 0x1a, 0x9d, 0x54, 0x6b, 0xb7, 0x55, 0xb6, 0x3e, 0x88, 0x08, 0x05, 0xdb,
	0xc2, 0x09, 0xb6, 0x6f, 0xa5, 0x50, 0xd3, 0x93, 0x28, 0x55, 0x03, 0x6b, 0x24, 0x50, 0x28, 0xe5,
	0x34, 0xa3, 0xb7, 0x90, 0xd3, 0xeb, 0xdd, 0x1a, 0xda, 0x7e, 0xc2, 0x3e, 0xf6, 0x23, 0x14, 0x1f,
	0xfb, 0xd6, 0x6f, 0x50, 0x76, 0xf6, 0x94, 0x98, 0x97, 0x63, 0x7e, 0x33, 0xff, 0x9d, 0xdd, 0xf9,
	0xef, 0x1e, 0x7c, 0x58, 0x29, 0x7d, 0xeb, 0xcf, 0xdd, 0x45, 0xe0, 0xeb, 0x04, 0xe3, 0x3b, 0x8c,
	0x5d, 0x15, 0x76, 0xd4, 0x5a, 0x63, 0xbc, 0xf4, 0x17, 0x18, 0xc5, 0x9b, 0xf9, 0x76, 0xd9, 0x89,
	0xe6, 0xe1, 0x26, 0xc4, 0xb5, 0x4e, 0x3a, 0x21, 0xde, 0x28, 0xdf, 0x7e, 0xdd, 0x28, 0xde, 0xe8,
	0x0d, 0xcf, 0xdb, 0x52, 0xeb, 0x2f, 0x83, 0xe2, 0x40, 0xad, 0x82, 0x91, 0x5a, 0x05, 0x9a, 0xbf,
	0x84, 0xec, 0xd5, 0x8f, 0x08, 0x05, 0x6b, 0xb2, 0x76, 0xb9, 0x5b, 0x75, 0xad, 0xc8, 0x3d, 0x08,
	0x4c, 0x51, 0x92, 0x84, 0xd7, 0x20, 0x7f, 0xb9, 0x5c, 0x26, 0xa8, 0x45, 0xa6, 0xc9, 0xda, 0x39,
	0x99, 0x92, 0xc9, 0xcf, 0x12, 0x8c, 0x87, 0x9e, 0x70, 0x6c, 0xde, 0x12, 0xaf, 0x43, 0xc1, 0x44,
	0x13, 0x3f, 0x44, 0x91, 0x6d, 0xb2, 0x76, 0x51, 0x1e, 0x98, 0x37, 0x00, 0x7a, 0x8b, 0x05, 0x26,
	0xc9, 0xc0, 0x4f, 0x02, 0x91, 0x6b, 0xb2, 0x76, 0x56, 0xde, 0xcb, 0xf0, 0x67, 0x90, 0x1b, 0xa9,
	0x50, 0x69, 0x91, 0xa7, 0x96, 0x16, 0xb8, 0x80, 0xd3, 0x59, 0x7a, 0x84, 0x53, 0xca, 0xef, 0x91,
	0xce, 0x60, 0x17, 0x14, 0xd2, 0x33, 0x10, 0xb5, 0xde, 0xc2, 0xd9, 0xd8, 0x78, 0x70, 0x81, 0x5a,
	0x62, 0xc4, 0x5f, 0x40, 0x9e, 0x30, 0x11, 0xac, 0xe9, 0xb4, 0xcf, 0xba, 0xa5, 0xfd, 0xbc, 0x94,
	0x95, 0x69, 0xb1, 0xf5, 0x2f, 0x03, 0x39, 0x0a, 0x79, 0x19, 0x32, 0x43, 0x8f, 0xcc, 0x71, 0x64,
	0x66, 0xe8, 0x99, 0x7d, 0xc6, 0xbe, 0x5a, 0x0f, 0x3d, 0xf2, 0xc0, 0x91, 0x29, 0xf1, 0x0a, 0x38,
	0x53, 0xfc, 0x46, 0x06, 0x38, 0xd2, 0x84, 0x9c, 0x43, 0xf6, 0xde, 0xe4, 0x14, 0x1b, 0x55, 0xff,
	0xbb, 0xa6, 0x71, 0x73, 0xd2, 0x84, 0x46, 0x75, 0x15, 0x6c, 0x43, 0x1a, 0xb3, 0x28, 0x29, 0x36,
	0x7b, 0x48, 0x5c, 0xa9, 0xcd, 0x9a, 0x86, 0x2c, 0xca, 0x94, 0x8c, 0x76, 0xaa, 0x7e, 0x62, 0x3a,
	0x21, 0xc5, 0xc6, 0x63, 0xb3, 0x86, 0xf2, 0x45, 0xca, 0x1f, 0xd8, 0xe8, 0xc9, 0x5d, 0xb0, 0xbd,
	0xc9, 0xd7, 0x3a, 0x14, 0xbc, 0x6d, 0xec, 0x6b, 0xd3, 0xfd, 0xcc, 0xea, 0xf7, 0x6c, 0xf6, 0x1d,
	0xa0, 0xb9, 0x73, 0xf1, 0xc8, 0x7a, 0x68, 0xc9, 0xdc, 0xc5, 0xb5, 0xba, 0xd1, 0x81, 0x28, 0xd9,
	0xbb, 0x20, 0x30, 0x9d, 0x3e, 0xc6, 0xe8, 0x6b, 0xec, 0x69, 0x51, 0xa6, 0xb1, 0x0f, 0x6c, 0x3a,
	0x4d, 0xb5, 0xaf, 0xb7, 0x89, 0x78, 0x6c, 0x3b, 0x59, 0xa2, 0xdd, 0xf1, 0x16, 0x69, 0x4d, 0xc5,
	0xae, 0xd9, 0xf3, 0xab, 0x2f, 0x50, 0x3a, 0x7a, 0x74, 0xfc, 0x39, 0x54, 0x8f, 0x12, 0x5f, 0x67,
	0x93, 0x4f, 0x93, 0xcb, 0xeb, 0x49, 0xe5, 0x84, 0xd7, 0xa1, 0xf6, 0xa0, 0x34, 0xed, 0xcb, 0x49,
	0x6f, 0xdc, 0xaf, 0x30, 0x5e, 0x85, 0x27, 0x0f, 0x6a, 0x72, 0x54, 0xc9, 0x74, 0xdf, 0x43, 0x99,
	0x6e, 0xd4, 0x3b, 0x9f, 0x62, 0x7c, 0xa7, 0x16, 0xc8, 0x5f, 0x83, 0x73, 0x81, 0x9a, 0x1f, 0x3f,
	0x81, 0xfa, 0xd3, 0x23, 0xb4, 0xcf, 0xa6, 0x75, 0x72, 0x5e, 0xfb, 0xb5, 0x6b, 0xb0, 0xdf, 0xbb,
	0x06, 0xfb, 0xb3, 0x6b, 0xb0, 0xcf, 0x05, 0xf7, 0x9d, 0x55, 0xcd, 0xf3, 0xf4, 0x6f, 0xbd, 0xf9,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x63, 0xea, 0x76, 0x9f, 0x03, 0x00, 0x00,
}
