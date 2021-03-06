// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: demo.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamNum     string `protobuf:"bytes,1,opt,name=exam_num,json=examNum,proto3" json:"exam_num,omitempty"`
	StudentName string `protobuf:"bytes,2,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentRequest) GetExamNum() string {
	if x != nil {
		return x.ExamNum
	}
	return ""
}

func (x *GetStudentRequest) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

type GetStudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamNum      string `protobuf:"bytes,1,opt,name=exam_num,json=examNum,proto3" json:"exam_num,omitempty"`
	StudentName  string `protobuf:"bytes,2,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
	ClassName    string `protobuf:"bytes,3,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ChineseScore string `protobuf:"bytes,4,opt,name=chinese_score,json=chineseScore,proto3" json:"chinese_score,omitempty"`
	MathScore    string `protobuf:"bytes,5,opt,name=math_score,json=mathScore,proto3" json:"math_score,omitempty"`
	EnglishScore string `protobuf:"bytes,6,opt,name=english_score,json=englishScore,proto3" json:"english_score,omitempty"`
	TotalScore   string `protobuf:"bytes,7,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"`
	ClassRate    string `protobuf:"bytes,8,opt,name=class_rate,json=classRate,proto3" json:"class_rate,omitempty"`
	SchoolRate   string `protobuf:"bytes,9,opt,name=school_rate,json=schoolRate,proto3" json:"school_rate,omitempty"`
	StepRank     string `protobuf:"bytes,10,opt,name=step_rank,json=stepRank,proto3" json:"step_rank,omitempty"`
	UploadDate   string `protobuf:"bytes,11,opt,name=upload_date,json=uploadDate,proto3" json:"upload_date,omitempty"`
	IsDeleted    string `protobuf:"bytes,12,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	DeleteTime   string `protobuf:"bytes,13,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func (x *GetStudentReply) Reset() {
	*x = GetStudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentReply) ProtoMessage() {}

func (x *GetStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentReply.ProtoReflect.Descriptor instead.
func (*GetStudentReply) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *GetStudentReply) GetExamNum() string {
	if x != nil {
		return x.ExamNum
	}
	return ""
}

func (x *GetStudentReply) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *GetStudentReply) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *GetStudentReply) GetChineseScore() string {
	if x != nil {
		return x.ChineseScore
	}
	return ""
}

func (x *GetStudentReply) GetMathScore() string {
	if x != nil {
		return x.MathScore
	}
	return ""
}

func (x *GetStudentReply) GetEnglishScore() string {
	if x != nil {
		return x.EnglishScore
	}
	return ""
}

func (x *GetStudentReply) GetTotalScore() string {
	if x != nil {
		return x.TotalScore
	}
	return ""
}

func (x *GetStudentReply) GetClassRate() string {
	if x != nil {
		return x.ClassRate
	}
	return ""
}

func (x *GetStudentReply) GetSchoolRate() string {
	if x != nil {
		return x.SchoolRate
	}
	return ""
}

func (x *GetStudentReply) GetStepRank() string {
	if x != nil {
		return x.StepRank
	}
	return ""
}

func (x *GetStudentReply) GetUploadDate() string {
	if x != nil {
		return x.UploadDate
	}
	return ""
}

func (x *GetStudentReply) GetIsDeleted() string {
	if x != nil {
		return x.IsDeleted
	}
	return ""
}

func (x *GetStudentReply) GetDeleteTime() string {
	if x != nil {
		return x.DeleteTime
	}
	return ""
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x65, 0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78,
	0x61, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x03, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x78, 0x61, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x52, 0x61, 0x6e, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0xd2, 0x01, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x12, 0x59,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x65, 0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65,
	0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x6b, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65,
	0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x77, 0x65, 0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x22, 0x13, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x3d, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65,
	0x62, 0x41, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x41, 0x70, 0x70, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x41, 0x70, 0x70, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_demo_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),      // 0: api.webApp.v1.HelloRequest
	(*HelloReply)(nil),        // 1: api.webApp.v1.HelloReply
	(*GetStudentRequest)(nil), // 2: api.webApp.v1.GetStudentRequest
	(*GetStudentReply)(nil),   // 3: api.webApp.v1.GetStudentReply
}
var file_demo_proto_depIdxs = []int32{
	0, // 0: api.webApp.v1.AppExcel.SayHello:input_type -> api.webApp.v1.HelloRequest
	2, // 1: api.webApp.v1.AppExcel.GetStudent:input_type -> api.webApp.v1.GetStudentRequest
	1, // 2: api.webApp.v1.AppExcel.SayHello:output_type -> api.webApp.v1.HelloReply
	3, // 3: api.webApp.v1.AppExcel.GetStudent:output_type -> api.webApp.v1.GetStudentReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentRequest); i {
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
		file_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentReply); i {
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
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}
