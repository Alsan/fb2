// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: filebrowser.proto

package proto

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

// reusable messages
type ReplyStatus int32

const (
	ReplyStatus_Unknown ReplyStatus = 0
	ReplyStatus_Ok      ReplyStatus = 1
	ReplyStatus_Failed  ReplyStatus = 2
)

// Enum value maps for ReplyStatus.
var (
	ReplyStatus_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	ReplyStatus_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x ReplyStatus) Enum() *ReplyStatus {
	p := new(ReplyStatus)
	*p = x
	return p
}

func (x ReplyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReplyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_filebrowser_proto_enumTypes[0].Descriptor()
}

func (ReplyStatus) Type() protoreflect.EnumType {
	return &file_filebrowser_proto_enumTypes[0]
}

func (x ReplyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReplyStatus.Descriptor instead.
func (ReplyStatus) EnumDescriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{0}
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []string `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{0}
}

func (x *List) GetItem() []string {
	if x != nil {
		return x.Item
	}
	return nil
}

// login request and response
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReplyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ReplyStatus" json:"status,omitempty"`
	// Types that are assignable to Data:
	//	*LoginReply_Message
	//	*LoginReply_Token
	Data isLoginReply_Data `protobuf_oneof:"data"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReply) GetStatus() ReplyStatus {
	if x != nil {
		return x.Status
	}
	return ReplyStatus_Unknown
}

func (m *LoginReply) GetData() isLoginReply_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *LoginReply) GetMessage() string {
	if x, ok := x.GetData().(*LoginReply_Message); ok {
		return x.Message
	}
	return ""
}

func (x *LoginReply) GetToken() string {
	if x, ok := x.GetData().(*LoginReply_Token); ok {
		return x.Token
	}
	return ""
}

type isLoginReply_Data interface {
	isLoginReply_Data()
}

type LoginReply_Message struct {
	Message string `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

type LoginReply_Token struct {
	Token string `protobuf:"bytes,3,opt,name=token,proto3,oneof"`
}

func (*LoginReply_Message) isLoginReply_Data() {}

func (*LoginReply_Token) isLoginReply_Data() {}

// file list request and response
type FileListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Path   string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Filter *string `protobuf:"bytes,3,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *FileListRequest) Reset() {
	*x = FileListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileListRequest) ProtoMessage() {}

func (x *FileListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileListRequest.ProtoReflect.Descriptor instead.
func (*FileListRequest) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{3}
}

func (x *FileListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FileListRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileListRequest) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

type FileListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReplyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ReplyStatus" json:"status,omitempty"`
	// Types that are assignable to Data:
	//	*FileListReply_Message
	//	*FileListReply_List
	Data isFileListReply_Data `protobuf_oneof:"data"`
}

func (x *FileListReply) Reset() {
	*x = FileListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileListReply) ProtoMessage() {}

func (x *FileListReply) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileListReply.ProtoReflect.Descriptor instead.
func (*FileListReply) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{4}
}

func (x *FileListReply) GetStatus() ReplyStatus {
	if x != nil {
		return x.Status
	}
	return ReplyStatus_Unknown
}

func (m *FileListReply) GetData() isFileListReply_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FileListReply) GetMessage() string {
	if x, ok := x.GetData().(*FileListReply_Message); ok {
		return x.Message
	}
	return ""
}

func (x *FileListReply) GetList() *List {
	if x, ok := x.GetData().(*FileListReply_List); ok {
		return x.List
	}
	return nil
}

type isFileListReply_Data interface {
	isFileListReply_Data()
}

type FileListReply_Message struct {
	Message string `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

type FileListReply_List struct {
	List *List `protobuf:"bytes,3,opt,name=list,proto3,oneof"`
}

func (*FileListReply_Message) isFileListReply_Data() {}

func (*FileListReply_List) isFileListReply_Data() {}

// upload file request and response
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Size     string `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
	Checksum string `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *FileInfo) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Types that are assignable to Request:
	//	*UploadFileRequest_Metadata
	//	*UploadFileRequest_Content
	Request isUploadFileRequest_Request `protobuf_oneof:"request"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{6}
}

func (x *UploadFileRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (m *UploadFileRequest) GetRequest() isUploadFileRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *UploadFileRequest) GetMetadata() *FileInfo {
	if x, ok := x.GetRequest().(*UploadFileRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *UploadFileRequest) GetContent() []byte {
	if x, ok := x.GetRequest().(*UploadFileRequest_Content); ok {
		return x.Content
	}
	return nil
}

type isUploadFileRequest_Request interface {
	isUploadFileRequest_Request()
}

type UploadFileRequest_Metadata struct {
	Metadata *FileInfo `protobuf:"bytes,2,opt,name=metadata,proto3,oneof"`
}

type UploadFileRequest_Content struct {
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3,oneof"`
}

func (*UploadFileRequest_Metadata) isUploadFileRequest_Request() {}

func (*UploadFileRequest_Content) isUploadFileRequest_Request() {}

type UploadFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ReplyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ReplyStatus" json:"status,omitempty"`
	Message *string     `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *UploadFileReply) Reset() {
	*x = UploadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileReply) ProtoMessage() {}

func (x *UploadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileReply.ProtoReflect.Descriptor instead.
func (*UploadFileReply) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{7}
}

func (x *UploadFileReply) GetStatus() ReplyStatus {
	if x != nil {
		return x.Status
	}
	return ReplyStatus_Unknown
}

func (x *UploadFileReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// download file request and response
type DownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadFileRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DownloadFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DownloadFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReplyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ReplyStatus" json:"status,omitempty"`
	// Types that are assignable to Data:
	//	*DownloadFileReply_Message
	//	*DownloadFileReply_Content
	Data isDownloadFileReply_Data `protobuf_oneof:"data"`
}

func (x *DownloadFileReply) Reset() {
	*x = DownloadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filebrowser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileReply) ProtoMessage() {}

func (x *DownloadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_filebrowser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileReply.ProtoReflect.Descriptor instead.
func (*DownloadFileReply) Descriptor() ([]byte, []int) {
	return file_filebrowser_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadFileReply) GetStatus() ReplyStatus {
	if x != nil {
		return x.Status
	}
	return ReplyStatus_Unknown
}

func (m *DownloadFileReply) GetData() isDownloadFileReply_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DownloadFileReply) GetMessage() string {
	if x, ok := x.GetData().(*DownloadFileReply_Message); ok {
		return x.Message
	}
	return ""
}

func (x *DownloadFileReply) GetContent() []byte {
	if x, ok := x.GetData().(*DownloadFileReply_Content); ok {
		return x.Content
	}
	return nil
}

type isDownloadFileReply_Data interface {
	isDownloadFileReply_Data()
}

type DownloadFileReply_Message struct {
	Message string `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

type DownloadFileReply_Content struct {
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3,oneof"`
}

func (*DownloadFileReply_Message) isDownloadFileReply_Data() {}

func (*DownloadFileReply_Content) isDownloadFileReply_Data() {}

var File_filebrowser_proto protoreflect.FileDescriptor

var file_filebrowser_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22,
	0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x0d,
	0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x22, 0x79, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x0f, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3f, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x79, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2e, 0x0a, 0x0b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0xe4, 0x01, 0x0a, 0x15,
	0x46, 0x69, 0x6c, 0x65, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x52, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x3c, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x73, 0x61, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x72, 0x6f, 0x77, 0x73,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filebrowser_proto_rawDescOnce sync.Once
	file_filebrowser_proto_rawDescData = file_filebrowser_proto_rawDesc
)

func file_filebrowser_proto_rawDescGZIP() []byte {
	file_filebrowser_proto_rawDescOnce.Do(func() {
		file_filebrowser_proto_rawDescData = protoimpl.X.CompressGZIP(file_filebrowser_proto_rawDescData)
	})
	return file_filebrowser_proto_rawDescData
}

var file_filebrowser_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filebrowser_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_filebrowser_proto_goTypes = []interface{}{
	(ReplyStatus)(0),            // 0: ReplyStatus
	(*List)(nil),                // 1: List
	(*LoginRequest)(nil),        // 2: LoginRequest
	(*LoginReply)(nil),          // 3: LoginReply
	(*FileListRequest)(nil),     // 4: FileListRequest
	(*FileListReply)(nil),       // 5: FileListReply
	(*FileInfo)(nil),            // 6: FileInfo
	(*UploadFileRequest)(nil),   // 7: UploadFileRequest
	(*UploadFileReply)(nil),     // 8: UploadFileReply
	(*DownloadFileRequest)(nil), // 9: DownloadFileRequest
	(*DownloadFileReply)(nil),   // 10: DownloadFileReply
}
var file_filebrowser_proto_depIdxs = []int32{
	0,  // 0: LoginReply.status:type_name -> ReplyStatus
	0,  // 1: FileListReply.status:type_name -> ReplyStatus
	1,  // 2: FileListReply.list:type_name -> List
	6,  // 3: UploadFileRequest.metadata:type_name -> FileInfo
	0,  // 4: UploadFileReply.status:type_name -> ReplyStatus
	0,  // 5: DownloadFileReply.status:type_name -> ReplyStatus
	2,  // 6: FileBrowserRpcService.Login:input_type -> LoginRequest
	4,  // 7: FileBrowserRpcService.FileList:input_type -> FileListRequest
	7,  // 8: FileBrowserRpcService.UploadFile:input_type -> UploadFileRequest
	9,  // 9: FileBrowserRpcService.DownloadFile:input_type -> DownloadFileRequest
	3,  // 10: FileBrowserRpcService.Login:output_type -> LoginReply
	5,  // 11: FileBrowserRpcService.FileList:output_type -> FileListReply
	8,  // 12: FileBrowserRpcService.UploadFile:output_type -> UploadFileReply
	10, // 13: FileBrowserRpcService.DownloadFile:output_type -> DownloadFileReply
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_filebrowser_proto_init() }
func file_filebrowser_proto_init() {
	if File_filebrowser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filebrowser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
		file_filebrowser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_filebrowser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_filebrowser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileListRequest); i {
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
		file_filebrowser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileListReply); i {
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
		file_filebrowser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_filebrowser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_filebrowser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileReply); i {
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
		file_filebrowser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileRequest); i {
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
		file_filebrowser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileReply); i {
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
	file_filebrowser_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*LoginReply_Message)(nil),
		(*LoginReply_Token)(nil),
	}
	file_filebrowser_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_filebrowser_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FileListReply_Message)(nil),
		(*FileListReply_List)(nil),
	}
	file_filebrowser_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*UploadFileRequest_Metadata)(nil),
		(*UploadFileRequest_Content)(nil),
	}
	file_filebrowser_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_filebrowser_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*DownloadFileReply_Message)(nil),
		(*DownloadFileReply_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filebrowser_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filebrowser_proto_goTypes,
		DependencyIndexes: file_filebrowser_proto_depIdxs,
		EnumInfos:         file_filebrowser_proto_enumTypes,
		MessageInfos:      file_filebrowser_proto_msgTypes,
	}.Build()
	File_filebrowser_proto = out.File
	file_filebrowser_proto_rawDesc = nil
	file_filebrowser_proto_goTypes = nil
	file_filebrowser_proto_depIdxs = nil
}
