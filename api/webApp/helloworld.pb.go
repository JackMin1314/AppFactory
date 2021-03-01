// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.2
// source: helloworld.proto

package helloworld

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

type CreateHelloworldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHelloworldRequest) Reset() {
	*x = CreateHelloworldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHelloworldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHelloworldRequest) ProtoMessage() {}

func (x *CreateHelloworldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHelloworldRequest.ProtoReflect.Descriptor instead.
func (*CreateHelloworldRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

type CreateHelloworldReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHelloworldReply) Reset() {
	*x = CreateHelloworldReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHelloworldReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHelloworldReply) ProtoMessage() {}

func (x *CreateHelloworldReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHelloworldReply.ProtoReflect.Descriptor instead.
func (*CreateHelloworldReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{1}
}

type UpdateHelloworldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHelloworldRequest) Reset() {
	*x = UpdateHelloworldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHelloworldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHelloworldRequest) ProtoMessage() {}

func (x *UpdateHelloworldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHelloworldRequest.ProtoReflect.Descriptor instead.
func (*UpdateHelloworldRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{2}
}

type UpdateHelloworldReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHelloworldReply) Reset() {
	*x = UpdateHelloworldReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHelloworldReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHelloworldReply) ProtoMessage() {}

func (x *UpdateHelloworldReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHelloworldReply.ProtoReflect.Descriptor instead.
func (*UpdateHelloworldReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{3}
}

type DeleteHelloworldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHelloworldRequest) Reset() {
	*x = DeleteHelloworldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHelloworldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHelloworldRequest) ProtoMessage() {}

func (x *DeleteHelloworldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHelloworldRequest.ProtoReflect.Descriptor instead.
func (*DeleteHelloworldRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{4}
}

type DeleteHelloworldReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHelloworldReply) Reset() {
	*x = DeleteHelloworldReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHelloworldReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHelloworldReply) ProtoMessage() {}

func (x *DeleteHelloworldReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHelloworldReply.ProtoReflect.Descriptor instead.
func (*DeleteHelloworldReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{5}
}

type GetHelloworldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHelloworldRequest) Reset() {
	*x = GetHelloworldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHelloworldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloworldRequest) ProtoMessage() {}

func (x *GetHelloworldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloworldRequest.ProtoReflect.Descriptor instead.
func (*GetHelloworldRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{6}
}

type GetHelloworldReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHelloworldReply) Reset() {
	*x = GetHelloworldReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHelloworldReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloworldReply) ProtoMessage() {}

func (x *GetHelloworldReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloworldReply.ProtoReflect.Descriptor instead.
func (*GetHelloworldReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{7}
}

type ListHelloworldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHelloworldRequest) Reset() {
	*x = ListHelloworldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHelloworldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHelloworldRequest) ProtoMessage() {}

func (x *ListHelloworldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHelloworldRequest.ProtoReflect.Descriptor instead.
func (*ListHelloworldRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{8}
}

type ListHelloworldReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHelloworldReply) Reset() {
	*x = ListHelloworldReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHelloworldReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHelloworldReply) ProtoMessage() {}

func (x *ListHelloworldReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHelloworldReply.ProtoReflect.Descriptor instead.
func (*ListHelloworldReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{9}
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf1, 0x03, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x62, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x62, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x62, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x5c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x38, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x50, 0x01, 0x5a, 0x24, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x3b,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_helloworld_proto_goTypes = []interface{}{
	(*CreateHelloworldRequest)(nil), // 0: api.helloworld.CreateHelloworldRequest
	(*CreateHelloworldReply)(nil),   // 1: api.helloworld.CreateHelloworldReply
	(*UpdateHelloworldRequest)(nil), // 2: api.helloworld.UpdateHelloworldRequest
	(*UpdateHelloworldReply)(nil),   // 3: api.helloworld.UpdateHelloworldReply
	(*DeleteHelloworldRequest)(nil), // 4: api.helloworld.DeleteHelloworldRequest
	(*DeleteHelloworldReply)(nil),   // 5: api.helloworld.DeleteHelloworldReply
	(*GetHelloworldRequest)(nil),    // 6: api.helloworld.GetHelloworldRequest
	(*GetHelloworldReply)(nil),      // 7: api.helloworld.GetHelloworldReply
	(*ListHelloworldRequest)(nil),   // 8: api.helloworld.ListHelloworldRequest
	(*ListHelloworldReply)(nil),     // 9: api.helloworld.ListHelloworldReply
}
var file_helloworld_proto_depIdxs = []int32{
	0, // 0: api.helloworld.Helloworld.CreateHelloworld:input_type -> api.helloworld.CreateHelloworldRequest
	2, // 1: api.helloworld.Helloworld.UpdateHelloworld:input_type -> api.helloworld.UpdateHelloworldRequest
	4, // 2: api.helloworld.Helloworld.DeleteHelloworld:input_type -> api.helloworld.DeleteHelloworldRequest
	6, // 3: api.helloworld.Helloworld.GetHelloworld:input_type -> api.helloworld.GetHelloworldRequest
	8, // 4: api.helloworld.Helloworld.ListHelloworld:input_type -> api.helloworld.ListHelloworldRequest
	1, // 5: api.helloworld.Helloworld.CreateHelloworld:output_type -> api.helloworld.CreateHelloworldReply
	3, // 6: api.helloworld.Helloworld.UpdateHelloworld:output_type -> api.helloworld.UpdateHelloworldReply
	5, // 7: api.helloworld.Helloworld.DeleteHelloworld:output_type -> api.helloworld.DeleteHelloworldReply
	7, // 8: api.helloworld.Helloworld.GetHelloworld:output_type -> api.helloworld.GetHelloworldReply
	9, // 9: api.helloworld.Helloworld.ListHelloworld:output_type -> api.helloworld.ListHelloworldReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHelloworldRequest); i {
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
		file_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHelloworldReply); i {
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
		file_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHelloworldRequest); i {
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
		file_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHelloworldReply); i {
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
		file_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHelloworldRequest); i {
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
		file_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHelloworldReply); i {
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
		file_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHelloworldRequest); i {
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
		file_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHelloworldReply); i {
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
		file_helloworld_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHelloworldRequest); i {
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
		file_helloworld_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHelloworldReply); i {
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
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloworldClient interface {
	CreateHelloworld(ctx context.Context, in *CreateHelloworldRequest, opts ...grpc.CallOption) (*CreateHelloworldReply, error)
	UpdateHelloworld(ctx context.Context, in *UpdateHelloworldRequest, opts ...grpc.CallOption) (*UpdateHelloworldReply, error)
	DeleteHelloworld(ctx context.Context, in *DeleteHelloworldRequest, opts ...grpc.CallOption) (*DeleteHelloworldReply, error)
	GetHelloworld(ctx context.Context, in *GetHelloworldRequest, opts ...grpc.CallOption) (*GetHelloworldReply, error)
	ListHelloworld(ctx context.Context, in *ListHelloworldRequest, opts ...grpc.CallOption) (*ListHelloworldReply, error)
}

type helloworldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldClient(cc grpc.ClientConnInterface) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) CreateHelloworld(ctx context.Context, in *CreateHelloworldRequest, opts ...grpc.CallOption) (*CreateHelloworldReply, error) {
	out := new(CreateHelloworldReply)
	err := c.cc.Invoke(ctx, "/api.helloworld.Helloworld/CreateHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) UpdateHelloworld(ctx context.Context, in *UpdateHelloworldRequest, opts ...grpc.CallOption) (*UpdateHelloworldReply, error) {
	out := new(UpdateHelloworldReply)
	err := c.cc.Invoke(ctx, "/api.helloworld.Helloworld/UpdateHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) DeleteHelloworld(ctx context.Context, in *DeleteHelloworldRequest, opts ...grpc.CallOption) (*DeleteHelloworldReply, error) {
	out := new(DeleteHelloworldReply)
	err := c.cc.Invoke(ctx, "/api.helloworld.Helloworld/DeleteHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) GetHelloworld(ctx context.Context, in *GetHelloworldRequest, opts ...grpc.CallOption) (*GetHelloworldReply, error) {
	out := new(GetHelloworldReply)
	err := c.cc.Invoke(ctx, "/api.helloworld.Helloworld/GetHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) ListHelloworld(ctx context.Context, in *ListHelloworldRequest, opts ...grpc.CallOption) (*ListHelloworldReply, error) {
	out := new(ListHelloworldReply)
	err := c.cc.Invoke(ctx, "/api.helloworld.Helloworld/ListHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServer is the server API for Helloworld service.
type HelloworldServer interface {
	CreateHelloworld(context.Context, *CreateHelloworldRequest) (*CreateHelloworldReply, error)
	UpdateHelloworld(context.Context, *UpdateHelloworldRequest) (*UpdateHelloworldReply, error)
	DeleteHelloworld(context.Context, *DeleteHelloworldRequest) (*DeleteHelloworldReply, error)
	GetHelloworld(context.Context, *GetHelloworldRequest) (*GetHelloworldReply, error)
	ListHelloworld(context.Context, *ListHelloworldRequest) (*ListHelloworldReply, error)
}

// UnimplementedHelloworldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (*UnimplementedHelloworldServer) CreateHelloworld(context.Context, *CreateHelloworldRequest) (*CreateHelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHelloworld not implemented")
}
func (*UnimplementedHelloworldServer) UpdateHelloworld(context.Context, *UpdateHelloworldRequest) (*UpdateHelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHelloworld not implemented")
}
func (*UnimplementedHelloworldServer) DeleteHelloworld(context.Context, *DeleteHelloworldRequest) (*DeleteHelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHelloworld not implemented")
}
func (*UnimplementedHelloworldServer) GetHelloworld(context.Context, *GetHelloworldRequest) (*GetHelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelloworld not implemented")
}
func (*UnimplementedHelloworldServer) ListHelloworld(context.Context, *ListHelloworldRequest) (*ListHelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHelloworld not implemented")
}

func RegisterHelloworldServer(s *grpc.Server, srv HelloworldServer) {
	s.RegisterService(&_Helloworld_serviceDesc, srv)
}

func _Helloworld_CreateHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).CreateHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.helloworld.Helloworld/CreateHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).CreateHelloworld(ctx, req.(*CreateHelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_UpdateHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).UpdateHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.helloworld.Helloworld/UpdateHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).UpdateHelloworld(ctx, req.(*UpdateHelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_DeleteHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).DeleteHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.helloworld.Helloworld/DeleteHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).DeleteHelloworld(ctx, req.(*DeleteHelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_GetHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).GetHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.helloworld.Helloworld/GetHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).GetHelloworld(ctx, req.(*GetHelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_ListHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).ListHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.helloworld.Helloworld/ListHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).ListHelloworld(ctx, req.(*ListHelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.helloworld.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHelloworld",
			Handler:    _Helloworld_CreateHelloworld_Handler,
		},
		{
			MethodName: "UpdateHelloworld",
			Handler:    _Helloworld_UpdateHelloworld_Handler,
		},
		{
			MethodName: "DeleteHelloworld",
			Handler:    _Helloworld_DeleteHelloworld_Handler,
		},
		{
			MethodName: "GetHelloworld",
			Handler:    _Helloworld_GetHelloworld_Handler,
		},
		{
			MethodName: "ListHelloworld",
			Handler:    _Helloworld_ListHelloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
