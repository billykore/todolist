// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: internal/proto/v1/todo.proto

package v1

import (
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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsDone      bool   `protobuf:"varint,4,opt,name=isDone,proto3" json:"isDone,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

type GetTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDone string `protobuf:"bytes,1,opt,name=isDone,proto3" json:"isDone,omitempty"`
}

func (x *GetTodosRequest) Reset() {
	*x = GetTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosRequest) ProtoMessage() {}

func (x *GetTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosRequest.ProtoReflect.Descriptor instead.
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *GetTodosRequest) GetIsDone() string {
	if x != nil {
		return x.IsDone
	}
	return ""
}

type GetTodosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *GetTodosReply) Reset() {
	*x = GetTodosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosReply) ProtoMessage() {}

func (x *GetTodosReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosReply.ProtoReflect.Descriptor instead.
func (*GetTodosReply) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{2}
}

func (x *GetTodosReply) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetTodoReply) Reset() {
	*x = GetTodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoReply) ProtoMessage() {}

func (x *GetTodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoReply.ProtoReflect.Descriptor instead.
func (*GetTodoReply) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodoReply) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type AddTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddTodoRequest) Reset() {
	*x = AddTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoRequest) ProtoMessage() {}

func (x *AddTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoRequest.ProtoReflect.Descriptor instead.
func (*AddTodoRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{5}
}

func (x *AddTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddTodoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DefaultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultReply) Reset() {
	*x = DefaultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_v1_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultReply) ProtoMessage() {}

func (x *DefaultReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_v1_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultReply.ProtoReflect.Descriptor instead.
func (*DefaultReply) Descriptor() ([]byte, []int) {
	return file_internal_proto_v1_todo_proto_rawDescGZIP(), []int{6}
}

func (x *DefaultReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_v1_todo_proto protoreflect.FileDescriptor

var file_internal_proto_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x44,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e,
	0x65, 0x22, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x22, 0x1d, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a,
	0x0c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf9, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12,
	0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x46, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x49,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x1a, 0x0b, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x23, 0x5a, 0x21, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2d, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_v1_todo_proto_rawDescOnce sync.Once
	file_internal_proto_v1_todo_proto_rawDescData = file_internal_proto_v1_todo_proto_rawDesc
)

func file_internal_proto_v1_todo_proto_rawDescGZIP() []byte {
	file_internal_proto_v1_todo_proto_rawDescOnce.Do(func() {
		file_internal_proto_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_v1_todo_proto_rawDescData)
	})
	return file_internal_proto_v1_todo_proto_rawDescData
}

var file_internal_proto_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_proto_v1_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),            // 0: todo.Todo
	(*GetTodosRequest)(nil), // 1: todo.GetTodosRequest
	(*GetTodosReply)(nil),   // 2: todo.GetTodosReply
	(*TodoRequest)(nil),     // 3: todo.TodoRequest
	(*GetTodoReply)(nil),    // 4: todo.GetTodoReply
	(*AddTodoRequest)(nil),  // 5: todo.AddTodoRequest
	(*DefaultReply)(nil),    // 6: todo.DefaultReply
}
var file_internal_proto_v1_todo_proto_depIdxs = []int32{
	0, // 0: todo.GetTodosReply.todos:type_name -> todo.Todo
	0, // 1: todo.GetTodoReply.todo:type_name -> todo.Todo
	1, // 2: todo.TodoService.GetTodos:input_type -> todo.GetTodosRequest
	3, // 3: todo.TodoService.GetTodo:input_type -> todo.TodoRequest
	5, // 4: todo.TodoService.AddTodo:input_type -> todo.AddTodoRequest
	3, // 5: todo.TodoService.SetDoneTodo:input_type -> todo.TodoRequest
	3, // 6: todo.TodoService.DeleteTodo:input_type -> todo.TodoRequest
	2, // 7: todo.TodoService.GetTodos:output_type -> todo.GetTodosReply
	4, // 8: todo.TodoService.GetTodo:output_type -> todo.GetTodoReply
	6, // 9: todo.TodoService.AddTodo:output_type -> todo.DefaultReply
	6, // 10: todo.TodoService.SetDoneTodo:output_type -> todo.DefaultReply
	6, // 11: todo.TodoService.DeleteTodo:output_type -> todo.DefaultReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_v1_todo_proto_init() }
func file_internal_proto_v1_todo_proto_init() {
	if File_internal_proto_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosRequest); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosReply); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoRequest); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoReply); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoRequest); i {
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
		file_internal_proto_v1_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultReply); i {
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
			RawDescriptor: file_internal_proto_v1_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_v1_todo_proto_goTypes,
		DependencyIndexes: file_internal_proto_v1_todo_proto_depIdxs,
		MessageInfos:      file_internal_proto_v1_todo_proto_msgTypes,
	}.Build()
	File_internal_proto_v1_todo_proto = out.File
	file_internal_proto_v1_todo_proto_rawDesc = nil
	file_internal_proto_v1_todo_proto_goTypes = nil
	file_internal_proto_v1_todo_proto_depIdxs = nil
}