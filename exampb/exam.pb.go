// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: exampb/exam.proto

package exampb

import (
	studentpb "github.com/Edigiraldo/grpc/studentpb"
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

type Exam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Exam) Reset() {
	*x = Exam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exam) ProtoMessage() {}

func (x *Exam) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exam.ProtoReflect.Descriptor instead.
func (*Exam) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{0}
}

func (x *Exam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Exam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	ExamId   string `protobuf:"bytes,4,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

type GetExamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExamRequest) Reset() {
	*x = GetExamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamRequest) ProtoMessage() {}

func (x *GetExamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamRequest.ProtoReflect.Descriptor instead.
func (*GetExamRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{2}
}

func (x *GetExamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetExamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetExamResponse) Reset() {
	*x = SetExamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetExamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExamResponse) ProtoMessage() {}

func (x *SetExamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExamResponse.ProtoReflect.Descriptor instead.
func (*SetExamResponse) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{3}
}

func (x *SetExamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetExamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetQuestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SetQuestionsResponse) Reset() {
	*x = SetQuestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetQuestionsResponse) ProtoMessage() {}

func (x *SetQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetQuestionsResponse.ProtoReflect.Descriptor instead.
func (*SetQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{4}
}

func (x *SetQuestionsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EnrollStudentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ExamId    string `protobuf:"bytes,2,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *EnrollStudentsRequest) Reset() {
	*x = EnrollStudentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollStudentsRequest) ProtoMessage() {}

func (x *EnrollStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollStudentsRequest.ProtoReflect.Descriptor instead.
func (*EnrollStudentsRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{5}
}

func (x *EnrollStudentsRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *EnrollStudentsRequest) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

type EnrollStudentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *EnrollStudentsResponse) Reset() {
	*x = EnrollStudentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollStudentsResponse) ProtoMessage() {}

func (x *EnrollStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollStudentsResponse.ProtoReflect.Descriptor instead.
func (*EnrollStudentsResponse) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{6}
}

func (x *EnrollStudentsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetStudentsPerExamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamId string `protobuf:"bytes,1,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *GetStudentsPerExamRequest) Reset() {
	*x = GetStudentsPerExamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentsPerExamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentsPerExamRequest) ProtoMessage() {}

func (x *GetStudentsPerExamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentsPerExamRequest.ProtoReflect.Descriptor instead.
func (*GetStudentsPerExamRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{7}
}

func (x *GetStudentsPerExamRequest) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

var File_exampb_exam_proto protoreflect.FileDescriptor

var file_exampb_exam_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78, 0x61, 0x6d, 0x1a, 0x17, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67,
	0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x26, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4f, 0x0a, 0x15, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x32, 0xbe, 0x02, 0x0a, 0x0b, 0x45, 0x78,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x45, 0x78, 0x61,
	0x6d, 0x12, 0x0a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x1a, 0x15, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x12, 0x4b, 0x0a, 0x0e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x64, 0x69, 0x67, 0x69, 0x72, 0x61,
	0x6c, 0x64, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exampb_exam_proto_rawDescOnce sync.Once
	file_exampb_exam_proto_rawDescData = file_exampb_exam_proto_rawDesc
)

func file_exampb_exam_proto_rawDescGZIP() []byte {
	file_exampb_exam_proto_rawDescOnce.Do(func() {
		file_exampb_exam_proto_rawDescData = protoimpl.X.CompressGZIP(file_exampb_exam_proto_rawDescData)
	})
	return file_exampb_exam_proto_rawDescData
}

var file_exampb_exam_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_exampb_exam_proto_goTypes = []interface{}{
	(*Exam)(nil),                      // 0: exam.Exam
	(*Question)(nil),                  // 1: exam.Question
	(*GetExamRequest)(nil),            // 2: exam.GetExamRequest
	(*SetExamResponse)(nil),           // 3: exam.SetExamResponse
	(*SetQuestionsResponse)(nil),      // 4: exam.SetQuestionsResponse
	(*EnrollStudentsRequest)(nil),     // 5: exam.EnrollStudentsRequest
	(*EnrollStudentsResponse)(nil),    // 6: exam.EnrollStudentsResponse
	(*GetStudentsPerExamRequest)(nil), // 7: exam.GetStudentsPerExamRequest
	(*studentpb.Student)(nil),         // 8: student.Student
}
var file_exampb_exam_proto_depIdxs = []int32{
	2, // 0: exam.ExamService.GetExam:input_type -> exam.GetExamRequest
	0, // 1: exam.ExamService.SetExam:input_type -> exam.Exam
	1, // 2: exam.ExamService.SetQuestions:input_type -> exam.Question
	5, // 3: exam.ExamService.EnrollStudents:input_type -> exam.EnrollStudentsRequest
	7, // 4: exam.ExamService.GetStudentsPerExam:input_type -> exam.GetStudentsPerExamRequest
	0, // 5: exam.ExamService.GetExam:output_type -> exam.Exam
	3, // 6: exam.ExamService.SetExam:output_type -> exam.SetExamResponse
	4, // 7: exam.ExamService.SetQuestions:output_type -> exam.SetQuestionsResponse
	6, // 8: exam.ExamService.EnrollStudents:output_type -> exam.EnrollStudentsResponse
	8, // 9: exam.ExamService.GetStudentsPerExam:output_type -> student.Student
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exampb_exam_proto_init() }
func file_exampb_exam_proto_init() {
	if File_exampb_exam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exampb_exam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exam); i {
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
		file_exampb_exam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_exampb_exam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamRequest); i {
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
		file_exampb_exam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetExamResponse); i {
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
		file_exampb_exam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetQuestionsResponse); i {
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
		file_exampb_exam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollStudentsRequest); i {
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
		file_exampb_exam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollStudentsResponse); i {
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
		file_exampb_exam_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentsPerExamRequest); i {
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
			RawDescriptor: file_exampb_exam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exampb_exam_proto_goTypes,
		DependencyIndexes: file_exampb_exam_proto_depIdxs,
		MessageInfos:      file_exampb_exam_proto_msgTypes,
	}.Build()
	File_exampb_exam_proto = out.File
	file_exampb_exam_proto_rawDesc = nil
	file_exampb_exam_proto_goTypes = nil
	file_exampb_exam_proto_depIdxs = nil
}
