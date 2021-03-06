// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: grpc-service/candidate-service.proto

package grpc_service

import (
	proto "github.com/golang/protobuf/proto"
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

type OrderDirection int32

const (
	OrderDirection_ASC  OrderDirection = 0
	OrderDirection_DESC OrderDirection = 1
)

// Enum value maps for OrderDirection.
var (
	OrderDirection_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderDirection_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderDirection) Enum() *OrderDirection {
	p := new(OrderDirection)
	*p = x
	return p
}

func (x OrderDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_candidate_service_proto_enumTypes[0].Descriptor()
}

func (OrderDirection) Type() protoreflect.EnumType {
	return &file_grpc_service_candidate_service_proto_enumTypes[0]
}

func (x OrderDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderDirection.Descriptor instead.
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{0}
}

type Contact_Type int32

const (
	Contact_PHONE Contact_Type = 0
	Contact_EMAIL Contact_Type = 1
)

// Enum value maps for Contact_Type.
var (
	Contact_Type_name = map[int32]string{
		0: "PHONE",
		1: "EMAIL",
	}
	Contact_Type_value = map[string]int32{
		"PHONE": 0,
		"EMAIL": 1,
	}
)

func (x Contact_Type) Enum() *Contact_Type {
	p := new(Contact_Type)
	*p = x
	return p
}

func (x Contact_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contact_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_candidate_service_proto_enumTypes[1].Descriptor()
}

func (Contact_Type) Type() protoreflect.EnumType {
	return &file_grpc_service_candidate_service_proto_enumTypes[1]
}

func (x Contact_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contact_Type.Descriptor instead.
func (Contact_Type) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{6, 0}
}

type Candidate_Order int32

const (
	Candidate_FULL_NAME Candidate_Order = 0
)

// Enum value maps for Candidate_Order.
var (
	Candidate_Order_name = map[int32]string{
		0: "FULL_NAME",
	}
	Candidate_Order_value = map[string]int32{
		"FULL_NAME": 0,
	}
)

func (x Candidate_Order) Enum() *Candidate_Order {
	p := new(Candidate_Order)
	*p = x
	return p
}

func (x Candidate_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Candidate_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_candidate_service_proto_enumTypes[2].Descriptor()
}

func (Candidate_Order) Type() protoreflect.EnumType {
	return &file_grpc_service_candidate_service_proto_enumTypes[2]
}

func (x Candidate_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Candidate_Order.Descriptor instead.
func (Candidate_Order) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{7, 0}
}

type CreateCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname  string     `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Contacts []*Contact `protobuf:"bytes,3,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *CreateCandidateRequest) Reset() {
	*x = CreateCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandidateRequest) ProtoMessage() {}

func (x *CreateCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandidateRequest.ProtoReflect.Descriptor instead.
func (*CreateCandidateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCandidateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCandidateRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *CreateCandidateRequest) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type CreateCandidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateCandidateResponse) Reset() {
	*x = CreateCandidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCandidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandidateResponse) ProtoMessage() {}

func (x *CreateCandidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandidateResponse.ProtoReflect.Descriptor instead.
func (*CreateCandidateResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCandidateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCandidateRequest) Reset() {
	*x = GetCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidateRequest) ProtoMessage() {}

func (x *GetCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidateRequest.ProtoReflect.Descriptor instead.
func (*GetCandidateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCandidateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCandidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate *Candidate `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *GetCandidateResponse) Reset() {
	*x = GetCandidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidateResponse) ProtoMessage() {}

func (x *GetCandidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidateResponse.ProtoReflect.Descriptor instead.
func (*GetCandidateResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCandidateResponse) GetCandidate() *Candidate {
	if x != nil {
		return x.Candidate
	}
	return nil
}

type SearchCandidatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           int32           `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size           int32           `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	OrderBy        Candidate_Order `protobuf:"varint,3,opt,name=orderBy,proto3,enum=Candidate_Order" json:"orderBy,omitempty"`
	OrderDirection OrderDirection  `protobuf:"varint,4,opt,name=orderDirection,proto3,enum=OrderDirection" json:"orderDirection,omitempty"`
}

func (x *SearchCandidatesRequest) Reset() {
	*x = SearchCandidatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCandidatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCandidatesRequest) ProtoMessage() {}

func (x *SearchCandidatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCandidatesRequest.ProtoReflect.Descriptor instead.
func (*SearchCandidatesRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchCandidatesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchCandidatesRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchCandidatesRequest) GetOrderBy() Candidate_Order {
	if x != nil {
		return x.OrderBy
	}
	return Candidate_FULL_NAME
}

func (x *SearchCandidatesRequest) GetOrderDirection() OrderDirection {
	if x != nil {
		return x.OrderDirection
	}
	return OrderDirection_ASC
}

type SearchCandidatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Candidate `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SearchCandidatesResponse) Reset() {
	*x = SearchCandidatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCandidatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCandidatesResponse) ProtoMessage() {}

func (x *SearchCandidatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCandidatesResponse.ProtoReflect.Descriptor instead.
func (*SearchCandidatesResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{5}
}

func (x *SearchCandidatesResponse) GetList() []*Candidate {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SearchCandidatesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Contact_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Contact_Type" json:"type,omitempty"`
	Value string       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{6}
}

func (x *Contact) GetType() Contact_Type {
	if x != nil {
		return x.Type
	}
	return Contact_PHONE
}

func (x *Contact) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname  string     `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Contacts []*Contact `protobuf:"bytes,4,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_candidate_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_candidate_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_grpc_service_candidate_service_proto_rawDescGZIP(), []int{7}
}

func (x *Candidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Candidate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Candidate) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Candidate) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

var File_grpc_service_candidate_service_proto protoreflect.FileDescriptor

var file_grpc_service_candidate_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x50, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x16, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x2a,
	0x23, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x32, 0xe5, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x79, 0x65, 0x72,
	0x6b, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_service_candidate_service_proto_rawDescOnce sync.Once
	file_grpc_service_candidate_service_proto_rawDescData = file_grpc_service_candidate_service_proto_rawDesc
)

func file_grpc_service_candidate_service_proto_rawDescGZIP() []byte {
	file_grpc_service_candidate_service_proto_rawDescOnce.Do(func() {
		file_grpc_service_candidate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_service_candidate_service_proto_rawDescData)
	})
	return file_grpc_service_candidate_service_proto_rawDescData
}

var file_grpc_service_candidate_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_grpc_service_candidate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_service_candidate_service_proto_goTypes = []interface{}{
	(OrderDirection)(0),              // 0: OrderDirection
	(Contact_Type)(0),                // 1: Contact.Type
	(Candidate_Order)(0),             // 2: Candidate.Order
	(*CreateCandidateRequest)(nil),   // 3: CreateCandidateRequest
	(*CreateCandidateResponse)(nil),  // 4: CreateCandidateResponse
	(*GetCandidateRequest)(nil),      // 5: GetCandidateRequest
	(*GetCandidateResponse)(nil),     // 6: GetCandidateResponse
	(*SearchCandidatesRequest)(nil),  // 7: SearchCandidatesRequest
	(*SearchCandidatesResponse)(nil), // 8: SearchCandidatesResponse
	(*Contact)(nil),                  // 9: Contact
	(*Candidate)(nil),                // 10: Candidate
}
var file_grpc_service_candidate_service_proto_depIdxs = []int32{
	9,  // 0: CreateCandidateRequest.contacts:type_name -> Contact
	10, // 1: GetCandidateResponse.candidate:type_name -> Candidate
	2,  // 2: SearchCandidatesRequest.orderBy:type_name -> Candidate.Order
	0,  // 3: SearchCandidatesRequest.orderDirection:type_name -> OrderDirection
	10, // 4: SearchCandidatesResponse.list:type_name -> Candidate
	1,  // 5: Contact.type:type_name -> Contact.Type
	9,  // 6: Candidate.contacts:type_name -> Contact
	3,  // 7: CandidatesService.CreateCandidate:input_type -> CreateCandidateRequest
	5,  // 8: CandidatesService.GetCandidate:input_type -> GetCandidateRequest
	7,  // 9: CandidatesService.SearchCandidates:input_type -> SearchCandidatesRequest
	4,  // 10: CandidatesService.CreateCandidate:output_type -> CreateCandidateResponse
	6,  // 11: CandidatesService.GetCandidate:output_type -> GetCandidateResponse
	8,  // 12: CandidatesService.SearchCandidates:output_type -> SearchCandidatesResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_grpc_service_candidate_service_proto_init() }
func file_grpc_service_candidate_service_proto_init() {
	if File_grpc_service_candidate_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_service_candidate_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCandidateRequest); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCandidateResponse); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidateRequest); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidateResponse); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCandidatesRequest); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCandidatesResponse); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_grpc_service_candidate_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
			RawDescriptor: file_grpc_service_candidate_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_service_candidate_service_proto_goTypes,
		DependencyIndexes: file_grpc_service_candidate_service_proto_depIdxs,
		EnumInfos:         file_grpc_service_candidate_service_proto_enumTypes,
		MessageInfos:      file_grpc_service_candidate_service_proto_msgTypes,
	}.Build()
	File_grpc_service_candidate_service_proto = out.File
	file_grpc_service_candidate_service_proto_rawDesc = nil
	file_grpc_service_candidate_service_proto_goTypes = nil
	file_grpc_service_candidate_service_proto_depIdxs = nil
}
