// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: ghost.proto

package ghost

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{0}
}

type APIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responsemessage string `protobuf:"bytes,1,opt,name=responsemessage,proto3" json:"responsemessage,omitempty"`
	Responsecode    int32  `protobuf:"varint,2,opt,name=responsecode,proto3" json:"responsecode,omitempty"`
}

func (x *APIResponse) Reset() {
	*x = APIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIResponse) ProtoMessage() {}

func (x *APIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIResponse.ProtoReflect.Descriptor instead.
func (*APIResponse) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{1}
}

func (x *APIResponse) GetResponsemessage() string {
	if x != nil {
		return x.Responsemessage
	}
	return ""
}

func (x *APIResponse) GetResponsecode() int32 {
	if x != nil {
		return x.Responsecode
	}
	return 0
}

type APIMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responsemessage []string `protobuf:"bytes,1,rep,name=responsemessage,proto3" json:"responsemessage,omitempty"`
}

func (x *APIMessage) Reset() {
	*x = APIMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIMessage) ProtoMessage() {}

func (x *APIMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIMessage.ProtoReflect.Descriptor instead.
func (*APIMessage) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{2}
}

func (x *APIMessage) GetResponsemessage() []string {
	if x != nil {
		return x.Responsemessage
	}
	return nil
}

type DiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dir string `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
}

func (x *DiskRequest) Reset() {
	*x = DiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskRequest) ProtoMessage() {}

func (x *DiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskRequest.ProtoReflect.Descriptor instead.
func (*DiskRequest) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{3}
}

func (x *DiskRequest) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

type DefinitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemaname string `protobuf:"bytes,1,opt,name=schemaname,proto3" json:"schemaname,omitempty"`
	Tablename  string `protobuf:"bytes,2,opt,name=tablename,proto3" json:"tablename,omitempty"`
}

func (x *DefinitionRequest) Reset() {
	*x = DefinitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefinitionRequest) ProtoMessage() {}

func (x *DefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefinitionRequest.ProtoReflect.Descriptor instead.
func (*DefinitionRequest) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{4}
}

func (x *DefinitionRequest) GetSchemaname() string {
	if x != nil {
		return x.Schemaname
	}
	return ""
}

func (x *DefinitionRequest) GetTablename() string {
	if x != nil {
		return x.Tablename
	}
	return ""
}

type IbdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dir        string `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	Schemaname string `protobuf:"bytes,2,opt,name=schemaname,proto3" json:"schemaname,omitempty"`
	Tablename  string `protobuf:"bytes,3,opt,name=tablename,proto3" json:"tablename,omitempty"`
}

func (x *IbdRequest) Reset() {
	*x = IbdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IbdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IbdRequest) ProtoMessage() {}

func (x *IbdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IbdRequest.ProtoReflect.Descriptor instead.
func (*IbdRequest) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{5}
}

func (x *IbdRequest) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *IbdRequest) GetSchemaname() string {
	if x != nil {
		return x.Schemaname
	}
	return ""
}

func (x *IbdRequest) GetTablename() string {
	if x != nil {
		return x.Tablename
	}
	return ""
}

type GhostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemaname string `protobuf:"bytes,1,opt,name=schemaname,proto3" json:"schemaname,omitempty"`
	Tablename  string `protobuf:"bytes,2,opt,name=tablename,proto3" json:"tablename,omitempty"`
	Statement  string `protobuf:"bytes,3,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (x *GhostRequest) Reset() {
	*x = GhostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GhostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GhostRequest) ProtoMessage() {}

func (x *GhostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GhostRequest.ProtoReflect.Descriptor instead.
func (*GhostRequest) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{6}
}

func (x *GhostRequest) GetSchemaname() string {
	if x != nil {
		return x.Schemaname
	}
	return ""
}

func (x *GhostRequest) GetTablename() string {
	if x != nil {
		return x.Tablename
	}
	return ""
}

func (x *GhostRequest) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

type InteractiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemaname   string `protobuf:"bytes,1,opt,name=schemaname,proto3" json:"schemaname,omitempty"`
	Tablename    string `protobuf:"bytes,2,opt,name=tablename,proto3" json:"tablename,omitempty"`
	Ghostcommand string `protobuf:"bytes,3,opt,name=ghostcommand,proto3" json:"ghostcommand,omitempty"`
}

func (x *InteractiveRequest) Reset() {
	*x = InteractiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractiveRequest) ProtoMessage() {}

func (x *InteractiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractiveRequest.ProtoReflect.Descriptor instead.
func (*InteractiveRequest) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{7}
}

func (x *InteractiveRequest) GetSchemaname() string {
	if x != nil {
		return x.Schemaname
	}
	return ""
}

func (x *InteractiveRequest) GetTablename() string {
	if x != nil {
		return x.Tablename
	}
	return ""
}

func (x *InteractiveRequest) GetGhostcommand() string {
	if x != nil {
		return x.Ghostcommand
	}
	return ""
}

var File_ghost_proto protoreflect.FileDescriptor

var file_ghost_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x0b, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0b, 0x64,
	0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x22, 0x51, 0x0a, 0x11,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x5c, 0x0a, 0x0a, 0x69, 0x62, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a,
	0x0c, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x12, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x32, 0xd0, 0x03, 0x0a, 0x05, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x64,
	0x69, 0x73, 0x6b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0c, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x07, 0x63, 0x75, 0x74,
	0x6f, 0x76, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0c, 0x70, 0x75,
	0x74, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x07, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x64, 0x72, 0x79, 0x72, 0x75, 0x6e, 0x12, 0x0d, 0x2e, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30,
	0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x68, 0x75,
	0x70, 0x12, 0x0d, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x13, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x72, 0x6f, 0x77, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x07, 0x69, 0x62, 0x64, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0b, 0x2e, 0x69, 0x62, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ghost_proto_rawDescOnce sync.Once
	file_ghost_proto_rawDescData = file_ghost_proto_rawDesc
)

func file_ghost_proto_rawDescGZIP() []byte {
	file_ghost_proto_rawDescOnce.Do(func() {
		file_ghost_proto_rawDescData = protoimpl.X.CompressGZIP(file_ghost_proto_rawDescData)
	})
	return file_ghost_proto_rawDescData
}

var file_ghost_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ghost_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: Empty
	(*APIResponse)(nil),        // 1: APIResponse
	(*APIMessage)(nil),         // 2: APIMessage
	(*DiskRequest)(nil),        // 3: diskRequest
	(*DefinitionRequest)(nil),  // 4: definitionRequest
	(*IbdRequest)(nil),         // 5: ibdRequest
	(*GhostRequest)(nil),       // 6: ghostRequest
	(*InteractiveRequest)(nil), // 7: interactiveRequest
}
var file_ghost_proto_depIdxs = []int32{
	3,  // 0: ghost.diskcheck:input_type -> diskRequest
	4,  // 1: ghost.checkdefinition:input_type -> definitionRequest
	0,  // 2: ghost.cutover:input_type -> Empty
	0,  // 3: ghost.putpanicflag:input_type -> Empty
	0,  // 4: ghost.cleanup:input_type -> Empty
	6,  // 5: ghost.dryrun:input_type -> ghostRequest
	6,  // 6: ghost.execute:input_type -> ghostRequest
	6,  // 7: ghost.executeNohup:input_type -> ghostRequest
	7,  // 8: ghost.interactive:input_type -> interactiveRequest
	4,  // 9: ghost.rowcount:input_type -> definitionRequest
	5,  // 10: ghost.ibdsize:input_type -> ibdRequest
	1,  // 11: ghost.diskcheck:output_type -> APIResponse
	1,  // 12: ghost.checkdefinition:output_type -> APIResponse
	1,  // 13: ghost.cutover:output_type -> APIResponse
	1,  // 14: ghost.putpanicflag:output_type -> APIResponse
	1,  // 15: ghost.cleanup:output_type -> APIResponse
	1,  // 16: ghost.dryrun:output_type -> APIResponse
	2,  // 17: ghost.execute:output_type -> APIMessage
	1,  // 18: ghost.executeNohup:output_type -> APIResponse
	1,  // 19: ghost.interactive:output_type -> APIResponse
	1,  // 20: ghost.rowcount:output_type -> APIResponse
	1,  // 21: ghost.ibdsize:output_type -> APIResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ghost_proto_init() }
func file_ghost_proto_init() {
	if File_ghost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ghost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_ghost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIResponse); i {
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
		file_ghost_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIMessage); i {
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
		file_ghost_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskRequest); i {
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
		file_ghost_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefinitionRequest); i {
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
		file_ghost_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IbdRequest); i {
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
		file_ghost_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GhostRequest); i {
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
		file_ghost_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractiveRequest); i {
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
			RawDescriptor: file_ghost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ghost_proto_goTypes,
		DependencyIndexes: file_ghost_proto_depIdxs,
		MessageInfos:      file_ghost_proto_msgTypes,
	}.Build()
	File_ghost_proto = out.File
	file_ghost_proto_rawDesc = nil
	file_ghost_proto_goTypes = nil
	file_ghost_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GhostClient is the client API for Ghost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GhostClient interface {
	Diskcheck(ctx context.Context, in *DiskRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Checkdefinition(ctx context.Context, in *DefinitionRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Cutover(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error)
	Putpanicflag(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error)
	Cleanup(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error)
	Dryrun(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Execute(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (Ghost_ExecuteClient, error)
	ExecuteNohup(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Interactive(ctx context.Context, in *InteractiveRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Rowcount(ctx context.Context, in *DefinitionRequest, opts ...grpc.CallOption) (*APIResponse, error)
	Ibdsize(ctx context.Context, in *IbdRequest, opts ...grpc.CallOption) (*APIResponse, error)
}

type ghostClient struct {
	cc grpc.ClientConnInterface
}

func NewGhostClient(cc grpc.ClientConnInterface) GhostClient {
	return &ghostClient{cc}
}

func (c *ghostClient) Diskcheck(ctx context.Context, in *DiskRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/diskcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Checkdefinition(ctx context.Context, in *DefinitionRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/checkdefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Cutover(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/cutover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Putpanicflag(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/putpanicflag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Cleanup(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/cleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Dryrun(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/dryrun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Execute(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (Ghost_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ghost_serviceDesc.Streams[0], "/ghost/execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &ghostExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ghost_ExecuteClient interface {
	Recv() (*APIMessage, error)
	grpc.ClientStream
}

type ghostExecuteClient struct {
	grpc.ClientStream
}

func (x *ghostExecuteClient) Recv() (*APIMessage, error) {
	m := new(APIMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ghostClient) ExecuteNohup(ctx context.Context, in *GhostRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/executeNohup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Interactive(ctx context.Context, in *InteractiveRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/interactive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Rowcount(ctx context.Context, in *DefinitionRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/rowcount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghostClient) Ibdsize(ctx context.Context, in *IbdRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/ghost/ibdsize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GhostServer is the server API for Ghost service.
type GhostServer interface {
	Diskcheck(context.Context, *DiskRequest) (*APIResponse, error)
	Checkdefinition(context.Context, *DefinitionRequest) (*APIResponse, error)
	Cutover(context.Context, *Empty) (*APIResponse, error)
	Putpanicflag(context.Context, *Empty) (*APIResponse, error)
	Cleanup(context.Context, *Empty) (*APIResponse, error)
	Dryrun(context.Context, *GhostRequest) (*APIResponse, error)
	Execute(*GhostRequest, Ghost_ExecuteServer) error
	ExecuteNohup(context.Context, *GhostRequest) (*APIResponse, error)
	Interactive(context.Context, *InteractiveRequest) (*APIResponse, error)
	Rowcount(context.Context, *DefinitionRequest) (*APIResponse, error)
	Ibdsize(context.Context, *IbdRequest) (*APIResponse, error)
}

// UnimplementedGhostServer can be embedded to have forward compatible implementations.
type UnimplementedGhostServer struct {
}

func (*UnimplementedGhostServer) Diskcheck(context.Context, *DiskRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diskcheck not implemented")
}
func (*UnimplementedGhostServer) Checkdefinition(context.Context, *DefinitionRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkdefinition not implemented")
}
func (*UnimplementedGhostServer) Cutover(context.Context, *Empty) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cutover not implemented")
}
func (*UnimplementedGhostServer) Putpanicflag(context.Context, *Empty) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Putpanicflag not implemented")
}
func (*UnimplementedGhostServer) Cleanup(context.Context, *Empty) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (*UnimplementedGhostServer) Dryrun(context.Context, *GhostRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dryrun not implemented")
}
func (*UnimplementedGhostServer) Execute(*GhostRequest, Ghost_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedGhostServer) ExecuteNohup(context.Context, *GhostRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteNohup not implemented")
}
func (*UnimplementedGhostServer) Interactive(context.Context, *InteractiveRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interactive not implemented")
}
func (*UnimplementedGhostServer) Rowcount(context.Context, *DefinitionRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rowcount not implemented")
}
func (*UnimplementedGhostServer) Ibdsize(context.Context, *IbdRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ibdsize not implemented")
}

func RegisterGhostServer(s *grpc.Server, srv GhostServer) {
	s.RegisterService(&_Ghost_serviceDesc, srv)
}

func _Ghost_Diskcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Diskcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Diskcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Diskcheck(ctx, req.(*DiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Checkdefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Checkdefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Checkdefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Checkdefinition(ctx, req.(*DefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Cutover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Cutover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Cutover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Cutover(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Putpanicflag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Putpanicflag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Putpanicflag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Putpanicflag(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Cleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Cleanup(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Dryrun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GhostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Dryrun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Dryrun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Dryrun(ctx, req.(*GhostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GhostRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GhostServer).Execute(m, &ghostExecuteServer{stream})
}

type Ghost_ExecuteServer interface {
	Send(*APIMessage) error
	grpc.ServerStream
}

type ghostExecuteServer struct {
	grpc.ServerStream
}

func (x *ghostExecuteServer) Send(m *APIMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Ghost_ExecuteNohup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GhostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).ExecuteNohup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/ExecuteNohup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).ExecuteNohup(ctx, req.(*GhostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Interactive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Interactive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Interactive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Interactive(ctx, req.(*InteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Rowcount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Rowcount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Rowcount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Rowcount(ctx, req.(*DefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ghost_Ibdsize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhostServer).Ibdsize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghost/Ibdsize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhostServer).Ibdsize(ctx, req.(*IbdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ghost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ghost",
	HandlerType: (*GhostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "diskcheck",
			Handler:    _Ghost_Diskcheck_Handler,
		},
		{
			MethodName: "checkdefinition",
			Handler:    _Ghost_Checkdefinition_Handler,
		},
		{
			MethodName: "cutover",
			Handler:    _Ghost_Cutover_Handler,
		},
		{
			MethodName: "putpanicflag",
			Handler:    _Ghost_Putpanicflag_Handler,
		},
		{
			MethodName: "cleanup",
			Handler:    _Ghost_Cleanup_Handler,
		},
		{
			MethodName: "dryrun",
			Handler:    _Ghost_Dryrun_Handler,
		},
		{
			MethodName: "executeNohup",
			Handler:    _Ghost_ExecuteNohup_Handler,
		},
		{
			MethodName: "interactive",
			Handler:    _Ghost_Interactive_Handler,
		},
		{
			MethodName: "rowcount",
			Handler:    _Ghost_Rowcount_Handler,
		},
		{
			MethodName: "ibdsize",
			Handler:    _Ghost_Ibdsize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _Ghost_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ghost.proto",
}
