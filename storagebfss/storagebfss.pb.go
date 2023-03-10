// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: storagebfss.proto

package storagebfss

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

// 块信息
type ObjectBlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockSize uint32 `protobuf:"varint,1,opt,name=BlockSize,proto3" json:"BlockSize,omitempty"` // 当前块在这个对象中所占的数据大小（由于块是加密的，所以这里的大小是按16字节块对齐的）
	DataOff   uint32 `protobuf:"varint,2,opt,name=DataOff,proto3" json:"DataOff,omitempty"`     // 当前块在这个对象中的数据偏移（由于块是加密的，所以这里的偏移是按16字节块对齐的）
	BlockKey  []byte `protobuf:"bytes,3,opt,name=BlockKey,proto3" json:"BlockKey,omitempty"`    // 块对应的密钥
}

func (x *ObjectBlockInfo) Reset() {
	*x = ObjectBlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectBlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectBlockInfo) ProtoMessage() {}

func (x *ObjectBlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectBlockInfo.ProtoReflect.Descriptor instead.
func (*ObjectBlockInfo) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectBlockInfo) GetBlockSize() uint32 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *ObjectBlockInfo) GetDataOff() uint32 {
	if x != nil {
		return x.DataOff
	}
	return 0
}

func (x *ObjectBlockInfo) GetBlockKey() []byte {
	if x != nil {
		return x.BlockKey
	}
	return nil
}

// 请求文件下载信息
type BFSSGetFilePtReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"` // 文件ID
}

func (x *BFSSGetFilePtReq) Reset() {
	*x = BFSSGetFilePtReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSGetFilePtReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSGetFilePtReq) ProtoMessage() {}

func (x *BFSSGetFilePtReq) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSGetFilePtReq.ProtoReflect.Descriptor instead.
func (*BFSSGetFilePtReq) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{1}
}

func (x *BFSSGetFilePtReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

// 返回文件下载信息
type BFSSGetFilePtResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectSize       uint32             `protobuf:"varint,1,opt,name=ObjectSize,proto3" json:"ObjectSize,omitempty"`            // 对象大小
	AccessHash       string             `protobuf:"bytes,2,opt,name=AccessHash,proto3" json:"AccessHash,omitempty"`             // hash
	ObjectBlockInfos []*ObjectBlockInfo `protobuf:"bytes,3,rep,name=ObjectBlockInfos,proto3" json:"ObjectBlockInfos,omitempty"` // 块信息，一个对象需要很多块存储，跨块
}

func (x *BFSSGetFilePtResp) Reset() {
	*x = BFSSGetFilePtResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSGetFilePtResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSGetFilePtResp) ProtoMessage() {}

func (x *BFSSGetFilePtResp) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSGetFilePtResp.ProtoReflect.Descriptor instead.
func (*BFSSGetFilePtResp) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{2}
}

func (x *BFSSGetFilePtResp) GetObjectSize() uint32 {
	if x != nil {
		return x.ObjectSize
	}
	return 0
}

func (x *BFSSGetFilePtResp) GetAccessHash() string {
	if x != nil {
		return x.AccessHash
	}
	return ""
}

func (x *BFSSGetFilePtResp) GetObjectBlockInfos() []*ObjectBlockInfo {
	if x != nil {
		return x.ObjectBlockInfos
	}
	return nil
}

// 请求上传小文件
type BFSSSaveFilePartPtReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId   uint64 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"`     // 文件ID
	FileSize uint32 `protobuf:"varint,2,opt,name=FileSize,proto3" json:"FileSize,omitempty"` // 文件大小
	FilePart uint32 `protobuf:"varint,3,opt,name=FilePart,proto3" json:"FilePart,omitempty"` // 当前块
	Offset   uint32 `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`     // 平移
	Data     []byte `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`          // buff
}

func (x *BFSSSaveFilePartPtReq) Reset() {
	*x = BFSSSaveFilePartPtReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveFilePartPtReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveFilePartPtReq) ProtoMessage() {}

func (x *BFSSSaveFilePartPtReq) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveFilePartPtReq.ProtoReflect.Descriptor instead.
func (*BFSSSaveFilePartPtReq) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{3}
}

func (x *BFSSSaveFilePartPtReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *BFSSSaveFilePartPtReq) GetFileSize() uint32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *BFSSSaveFilePartPtReq) GetFilePart() uint32 {
	if x != nil {
		return x.FilePart
	}
	return 0
}

func (x *BFSSSaveFilePartPtReq) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *BFSSSaveFilePartPtReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 上传小文件返回
type BFSSSaveFilePartPtResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"` // 返回是否上传成功
}

func (x *BFSSSaveFilePartPtResp) Reset() {
	*x = BFSSSaveFilePartPtResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveFilePartPtResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveFilePartPtResp) ProtoMessage() {}

func (x *BFSSSaveFilePartPtResp) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveFilePartPtResp.ProtoReflect.Descriptor instead.
func (*BFSSSaveFilePartPtResp) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{4}
}

func (x *BFSSSaveFilePartPtResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 请求上传大文件
type BFSSSaveBigFilePartPtReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId         uint64 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"`                 // 文件ID
	FileSize       uint32 `protobuf:"varint,2,opt,name=FileSize,proto3" json:"FileSize,omitempty"`             // 文件大小
	FileTotalParts uint32 `protobuf:"varint,3,opt,name=FileTotalParts,proto3" json:"FileTotalParts,omitempty"` // 总块数
	FilePart       uint32 `protobuf:"varint,4,opt,name=FilePart,proto3" json:"FilePart,omitempty"`             // 当前块
	Offset         uint32 `protobuf:"varint,5,opt,name=Offset,proto3" json:"Offset,omitempty"`                 // 平移
	Data           []byte `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`                      // buff
}

func (x *BFSSSaveBigFilePartPtReq) Reset() {
	*x = BFSSSaveBigFilePartPtReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveBigFilePartPtReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveBigFilePartPtReq) ProtoMessage() {}

func (x *BFSSSaveBigFilePartPtReq) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveBigFilePartPtReq.ProtoReflect.Descriptor instead.
func (*BFSSSaveBigFilePartPtReq) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{5}
}

func (x *BFSSSaveBigFilePartPtReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *BFSSSaveBigFilePartPtReq) GetFileSize() uint32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *BFSSSaveBigFilePartPtReq) GetFileTotalParts() uint32 {
	if x != nil {
		return x.FileTotalParts
	}
	return 0
}

func (x *BFSSSaveBigFilePartPtReq) GetFilePart() uint32 {
	if x != nil {
		return x.FilePart
	}
	return 0
}

func (x *BFSSSaveBigFilePartPtReq) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *BFSSSaveBigFilePartPtReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 上传大文件返回
type BFSSSaveBigFilePartPtResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"` // 返回是否上传成功
}

func (x *BFSSSaveBigFilePartPtResp) Reset() {
	*x = BFSSSaveBigFilePartPtResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveBigFilePartPtResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveBigFilePartPtResp) ProtoMessage() {}

func (x *BFSSSaveBigFilePartPtResp) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveBigFilePartPtResp.ProtoReflect.Descriptor instead.
func (*BFSSSaveBigFilePartPtResp) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{6}
}

func (x *BFSSSaveBigFilePartPtResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 上传文件完成通知
type BFSSSaveFileCompleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId   uint64 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"`     // 文件ID
	FileSize uint32 `protobuf:"varint,2,opt,name=FileSize,proto3" json:"FileSize,omitempty"` // 文件大小
	Flag     uint32 `protobuf:"varint,3,opt,name=Flag,proto3" json:"Flag,omitempty"`         // Flag=1小文件，Flag=2大文件
}

func (x *BFSSSaveFileCompleteReq) Reset() {
	*x = BFSSSaveFileCompleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveFileCompleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveFileCompleteReq) ProtoMessage() {}

func (x *BFSSSaveFileCompleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveFileCompleteReq.ProtoReflect.Descriptor instead.
func (*BFSSSaveFileCompleteReq) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{7}
}

func (x *BFSSSaveFileCompleteReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *BFSSSaveFileCompleteReq) GetFileSize() uint32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *BFSSSaveFileCompleteReq) GetFlag() uint32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

// 上传文件完成返回
type BFSSSaveFileCompleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"` // 返回是否上传成功
}

func (x *BFSSSaveFileCompleteResp) Reset() {
	*x = BFSSSaveFileCompleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSSaveFileCompleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSSaveFileCompleteResp) ProtoMessage() {}

func (x *BFSSSaveFileCompleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSSaveFileCompleteResp.ProtoReflect.Descriptor instead.
func (*BFSSSaveFileCompleteResp) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{8}
}

func (x *BFSSSaveFileCompleteResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 上传文件完成通知
type BFSSResetObjectIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"` // 文件ID
	NewId  uint64 `protobuf:"varint,2,opt,name=NewId,proto3" json:"NewId,omitempty"`   // 新ID
	NewTag string `protobuf:"bytes,3,opt,name=NewTag,proto3" json:"NewTag,omitempty"`  // Flag=1小文件，Flag=2大文件
}

func (x *BFSSResetObjectIdReq) Reset() {
	*x = BFSSResetObjectIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSResetObjectIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSResetObjectIdReq) ProtoMessage() {}

func (x *BFSSResetObjectIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSResetObjectIdReq.ProtoReflect.Descriptor instead.
func (*BFSSResetObjectIdReq) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{9}
}

func (x *BFSSResetObjectIdReq) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *BFSSResetObjectIdReq) GetNewId() uint64 {
	if x != nil {
		return x.NewId
	}
	return 0
}

func (x *BFSSResetObjectIdReq) GetNewTag() string {
	if x != nil {
		return x.NewTag
	}
	return ""
}

// 上传文件完成返回
type BFSSResetObjectIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"` // 返回是否上传成功
}

func (x *BFSSResetObjectIdResp) Reset() {
	*x = BFSSResetObjectIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storagebfss_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFSSResetObjectIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFSSResetObjectIdResp) ProtoMessage() {}

func (x *BFSSResetObjectIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_storagebfss_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFSSResetObjectIdResp.ProtoReflect.Descriptor instead.
func (*BFSSResetObjectIdResp) Descriptor() ([]byte, []int) {
	return file_storagebfss_proto_rawDescGZIP(), []int{10}
}

func (x *BFSSResetObjectIdResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_storagebfss_proto protoreflect.FileDescriptor

var file_storagebfss_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73,
	0x22, 0x65, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x66, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x22, 0x2a, 0x0a, 0x10, 0x42, 0x46, 0x53, 0x53, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x42, 0x46, 0x53, 0x53, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x48, 0x0a, 0x10, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73,
	0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x16, 0x42, 0x46, 0x53,
	0x53, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xbe, 0x01,
	0x0a, 0x18, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x42, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x61, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x35,
	0x0a, 0x19, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x42, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x61, 0x0a, 0x17, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x34, 0x0a, 0x18, 0x42, 0x46, 0x53, 0x53,
	0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5c,
	0x0a, 0x14, 0x42, 0x46, 0x53, 0x53, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4e,
	0x65, 0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x22, 0x31, 0x0a, 0x15,
	0x42, 0x46, 0x53, 0x53, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0xf3, 0x03, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x46, 0x53, 0x53, 0x12, 0x50, 0x0a, 0x0d, 0x42, 0x46, 0x53, 0x53, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66,
	0x73, 0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73,
	0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46,
	0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x15, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61,
	0x76, 0x65, 0x42, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x12,
	0x25, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46,
	0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x42, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x50, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x42, 0x69, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x50, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x65, 0x0a, 0x14, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x53, 0x61, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46, 0x53,
	0x53, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x42, 0x46, 0x53, 0x53, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46, 0x53, 0x53, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x66, 0x73, 0x73, 0x2e, 0x42, 0x46,
	0x53, 0x53, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x66, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storagebfss_proto_rawDescOnce sync.Once
	file_storagebfss_proto_rawDescData = file_storagebfss_proto_rawDesc
)

func file_storagebfss_proto_rawDescGZIP() []byte {
	file_storagebfss_proto_rawDescOnce.Do(func() {
		file_storagebfss_proto_rawDescData = protoimpl.X.CompressGZIP(file_storagebfss_proto_rawDescData)
	})
	return file_storagebfss_proto_rawDescData
}

var file_storagebfss_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_storagebfss_proto_goTypes = []interface{}{
	(*ObjectBlockInfo)(nil),           // 0: storagebfss.ObjectBlockInfo
	(*BFSSGetFilePtReq)(nil),          // 1: storagebfss.BFSSGetFilePtReq
	(*BFSSGetFilePtResp)(nil),         // 2: storagebfss.BFSSGetFilePtResp
	(*BFSSSaveFilePartPtReq)(nil),     // 3: storagebfss.BFSSSaveFilePartPtReq
	(*BFSSSaveFilePartPtResp)(nil),    // 4: storagebfss.BFSSSaveFilePartPtResp
	(*BFSSSaveBigFilePartPtReq)(nil),  // 5: storagebfss.BFSSSaveBigFilePartPtReq
	(*BFSSSaveBigFilePartPtResp)(nil), // 6: storagebfss.BFSSSaveBigFilePartPtResp
	(*BFSSSaveFileCompleteReq)(nil),   // 7: storagebfss.BFSSSaveFileCompleteReq
	(*BFSSSaveFileCompleteResp)(nil),  // 8: storagebfss.BFSSSaveFileCompleteResp
	(*BFSSResetObjectIdReq)(nil),      // 9: storagebfss.BFSSResetObjectIdReq
	(*BFSSResetObjectIdResp)(nil),     // 10: storagebfss.BFSSResetObjectIdResp
}
var file_storagebfss_proto_depIdxs = []int32{
	0,  // 0: storagebfss.BFSSGetFilePtResp.ObjectBlockInfos:type_name -> storagebfss.ObjectBlockInfo
	1,  // 1: storagebfss.FileStorageBFSS.BFSSGetFilePt:input_type -> storagebfss.BFSSGetFilePtReq
	3,  // 2: storagebfss.FileStorageBFSS.BFSSSaveFilePartPt:input_type -> storagebfss.BFSSSaveFilePartPtReq
	5,  // 3: storagebfss.FileStorageBFSS.BFSSSaveBigFilePartPt:input_type -> storagebfss.BFSSSaveBigFilePartPtReq
	7,  // 4: storagebfss.FileStorageBFSS.BFSSSaveFileComplete:input_type -> storagebfss.BFSSSaveFileCompleteReq
	9,  // 5: storagebfss.FileStorageBFSS.BFSSResetObjectId:input_type -> storagebfss.BFSSResetObjectIdReq
	2,  // 6: storagebfss.FileStorageBFSS.BFSSGetFilePt:output_type -> storagebfss.BFSSGetFilePtResp
	4,  // 7: storagebfss.FileStorageBFSS.BFSSSaveFilePartPt:output_type -> storagebfss.BFSSSaveFilePartPtResp
	6,  // 8: storagebfss.FileStorageBFSS.BFSSSaveBigFilePartPt:output_type -> storagebfss.BFSSSaveBigFilePartPtResp
	8,  // 9: storagebfss.FileStorageBFSS.BFSSSaveFileComplete:output_type -> storagebfss.BFSSSaveFileCompleteResp
	10, // 10: storagebfss.FileStorageBFSS.BFSSResetObjectId:output_type -> storagebfss.BFSSResetObjectIdResp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_storagebfss_proto_init() }
func file_storagebfss_proto_init() {
	if File_storagebfss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storagebfss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectBlockInfo); i {
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
		file_storagebfss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSGetFilePtReq); i {
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
		file_storagebfss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSGetFilePtResp); i {
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
		file_storagebfss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveFilePartPtReq); i {
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
		file_storagebfss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveFilePartPtResp); i {
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
		file_storagebfss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveBigFilePartPtReq); i {
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
		file_storagebfss_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveBigFilePartPtResp); i {
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
		file_storagebfss_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveFileCompleteReq); i {
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
		file_storagebfss_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSSaveFileCompleteResp); i {
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
		file_storagebfss_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSResetObjectIdReq); i {
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
		file_storagebfss_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFSSResetObjectIdResp); i {
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
			RawDescriptor: file_storagebfss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storagebfss_proto_goTypes,
		DependencyIndexes: file_storagebfss_proto_depIdxs,
		MessageInfos:      file_storagebfss_proto_msgTypes,
	}.Build()
	File_storagebfss_proto = out.File
	file_storagebfss_proto_rawDesc = nil
	file_storagebfss_proto_goTypes = nil
	file_storagebfss_proto_depIdxs = nil
}
