// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: imslowmodeserver.proto

package imslowmodeserver

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

type ToggleSlowModeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChnlId  int32 `protobuf:"varint,1,opt,name=Chnl_id,json=ChnlId,proto3" json:"Chnl_id,omitempty"`
	Seconds int32 `protobuf:"varint,2,opt,name=Seconds,proto3" json:"Seconds,omitempty"` // 间隔时间
}

func (x *ToggleSlowModeReq) Reset() {
	*x = ToggleSlowModeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imslowmodeserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleSlowModeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleSlowModeReq) ProtoMessage() {}

func (x *ToggleSlowModeReq) ProtoReflect() protoreflect.Message {
	mi := &file_imslowmodeserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleSlowModeReq.ProtoReflect.Descriptor instead.
func (*ToggleSlowModeReq) Descriptor() ([]byte, []int) {
	return file_imslowmodeserver_proto_rawDescGZIP(), []int{0}
}

func (x *ToggleSlowModeReq) GetChnlId() int32 {
	if x != nil {
		return x.ChnlId
	}
	return 0
}

func (x *ToggleSlowModeReq) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type ToggleSlowModeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResCode int32  `protobuf:"varint,1,opt,name=Res_code,json=ResCode,proto3" json:"Res_code,omitempty"` // 返回状态
	ResInfo string `protobuf:"bytes,2,opt,name=Res_info,json=ResInfo,proto3" json:"Res_info,omitempty"`
}

func (x *ToggleSlowModeRsp) Reset() {
	*x = ToggleSlowModeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imslowmodeserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleSlowModeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleSlowModeRsp) ProtoMessage() {}

func (x *ToggleSlowModeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_imslowmodeserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleSlowModeRsp.ProtoReflect.Descriptor instead.
func (*ToggleSlowModeRsp) Descriptor() ([]byte, []int) {
	return file_imslowmodeserver_proto_rawDescGZIP(), []int{1}
}

func (x *ToggleSlowModeRsp) GetResCode() int32 {
	if x != nil {
		return x.ResCode
	}
	return 0
}

func (x *ToggleSlowModeRsp) GetResInfo() string {
	if x != nil {
		return x.ResInfo
	}
	return ""
}

type CheckSlowLimitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChnlId  int32 `protobuf:"varint,1,opt,name=Chnl_id,json=ChnlId,proto3" json:"Chnl_id,omitempty"`
	MembId  int32 `protobuf:"varint,2,opt,name=Memb_id,json=MembId,proto3" json:"Memb_id,omitempty"`
	MsgTime int64 `protobuf:"varint,3,opt,name=MsgTime,proto3" json:"MsgTime,omitempty"` // 消息时间（s)
	Seconds int32 `protobuf:"varint,4,opt,name=Seconds,proto3" json:"Seconds,omitempty"` // 消息间隔时间
}

func (x *CheckSlowLimitReq) Reset() {
	*x = CheckSlowLimitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imslowmodeserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSlowLimitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSlowLimitReq) ProtoMessage() {}

func (x *CheckSlowLimitReq) ProtoReflect() protoreflect.Message {
	mi := &file_imslowmodeserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSlowLimitReq.ProtoReflect.Descriptor instead.
func (*CheckSlowLimitReq) Descriptor() ([]byte, []int) {
	return file_imslowmodeserver_proto_rawDescGZIP(), []int{2}
}

func (x *CheckSlowLimitReq) GetChnlId() int32 {
	if x != nil {
		return x.ChnlId
	}
	return 0
}

func (x *CheckSlowLimitReq) GetMembId() int32 {
	if x != nil {
		return x.MembId
	}
	return 0
}

func (x *CheckSlowLimitReq) GetMsgTime() int64 {
	if x != nil {
		return x.MsgTime
	}
	return 0
}

func (x *CheckSlowLimitReq) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type CheckSlowLimitRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is_Limit bool  `protobuf:"varint,1,opt,name=Is_Limit,json=IsLimit,proto3" json:"Is_Limit,omitempty"`
	NextTime int64 `protobuf:"varint,2,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (x *CheckSlowLimitRsp) Reset() {
	*x = CheckSlowLimitRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imslowmodeserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSlowLimitRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSlowLimitRsp) ProtoMessage() {}

func (x *CheckSlowLimitRsp) ProtoReflect() protoreflect.Message {
	mi := &file_imslowmodeserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSlowLimitRsp.ProtoReflect.Descriptor instead.
func (*CheckSlowLimitRsp) Descriptor() ([]byte, []int) {
	return file_imslowmodeserver_proto_rawDescGZIP(), []int{3}
}

func (x *CheckSlowLimitRsp) GetIs_Limit() bool {
	if x != nil {
		return x.Is_Limit
	}
	return false
}

func (x *CheckSlowLimitRsp) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

var File_imslowmodeserver_proto protoreflect.FileDescriptor

var file_imslowmodeserver_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6d, 0x73, 0x6c, 0x6f, 0x77, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x6d, 0x73, 0x6c, 0x6f, 0x77,
	0x6d, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x11, 0x54, 0x6f,
	0x67, 0x67, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x43, 0x68, 0x6e, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x43, 0x68, 0x6e, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x22, 0x49, 0x0a, 0x11, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x77,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x79, 0x0a,
	0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x43, 0x68, 0x6e, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x43, 0x68, 0x6e, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x4d,
	0x65, 0x6d, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x65,
	0x6d, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x49, 0x73, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x49, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4e, 0x65, 0x78, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0xcb, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65,
	0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x69, 0x6d, 0x73, 0x6c, 0x6f,
	0x77, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x67, 0x67,
	0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e,
	0x69, 0x6d, 0x73, 0x6c, 0x6f, 0x77, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x23, 0x2e, 0x69, 0x6d, 0x73, 0x6c, 0x6f, 0x77, 0x6d,
	0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x69, 0x6d,
	0x73, 0x6c, 0x6f, 0x77, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6d, 0x73, 0x6c, 0x6f, 0x77, 0x6d, 0x6f, 0x64, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imslowmodeserver_proto_rawDescOnce sync.Once
	file_imslowmodeserver_proto_rawDescData = file_imslowmodeserver_proto_rawDesc
)

func file_imslowmodeserver_proto_rawDescGZIP() []byte {
	file_imslowmodeserver_proto_rawDescOnce.Do(func() {
		file_imslowmodeserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_imslowmodeserver_proto_rawDescData)
	})
	return file_imslowmodeserver_proto_rawDescData
}

var file_imslowmodeserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imslowmodeserver_proto_goTypes = []interface{}{
	(*ToggleSlowModeReq)(nil), // 0: imslowmodeserver.ToggleSlowModeReq
	(*ToggleSlowModeRsp)(nil), // 1: imslowmodeserver.ToggleSlowModeRsp
	(*CheckSlowLimitReq)(nil), // 2: imslowmodeserver.CheckSlowLimitReq
	(*CheckSlowLimitRsp)(nil), // 3: imslowmodeserver.CheckSlowLimitRsp
}
var file_imslowmodeserver_proto_depIdxs = []int32{
	0, // 0: imslowmodeserver.ServerService.ToggleSlowMode:input_type -> imslowmodeserver.ToggleSlowModeReq
	2, // 1: imslowmodeserver.ServerService.CheckSlowLimit:input_type -> imslowmodeserver.CheckSlowLimitReq
	1, // 2: imslowmodeserver.ServerService.ToggleSlowMode:output_type -> imslowmodeserver.ToggleSlowModeRsp
	3, // 3: imslowmodeserver.ServerService.CheckSlowLimit:output_type -> imslowmodeserver.CheckSlowLimitRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imslowmodeserver_proto_init() }
func file_imslowmodeserver_proto_init() {
	if File_imslowmodeserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imslowmodeserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleSlowModeReq); i {
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
		file_imslowmodeserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleSlowModeRsp); i {
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
		file_imslowmodeserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckSlowLimitReq); i {
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
		file_imslowmodeserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckSlowLimitRsp); i {
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
			RawDescriptor: file_imslowmodeserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imslowmodeserver_proto_goTypes,
		DependencyIndexes: file_imslowmodeserver_proto_depIdxs,
		MessageInfos:      file_imslowmodeserver_proto_msgTypes,
	}.Build()
	File_imslowmodeserver_proto = out.File
	file_imslowmodeserver_proto_rawDesc = nil
	file_imslowmodeserver_proto_goTypes = nil
	file_imslowmodeserver_proto_depIdxs = nil
}
