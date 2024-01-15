// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: api/binanceexchange_user/v1/binanceexchange_user.proto

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

type SetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendBody *SetUserRequest_SendBody `protobuf:"bytes,1,opt,name=send_body,json=sendBody,proto3" json:"send_body,omitempty"`
}

func (x *SetUserRequest) Reset() {
	*x = SetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRequest) ProtoMessage() {}

func (x *SetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserRequest.ProtoReflect.Descriptor instead.
func (*SetUserRequest) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{0}
}

func (x *SetUserRequest) GetSendBody() *SetUserRequest_SendBody {
	if x != nil {
		return x.SendBody
	}
	return nil
}

type SetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUserReply) Reset() {
	*x = SetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserReply) ProtoMessage() {}

func (x *SetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserReply.ProtoReflect.Descriptor instead.
func (*SetUserReply) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{1}
}

type PullUserStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserStatusRequest) Reset() {
	*x = PullUserStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserStatusRequest) ProtoMessage() {}

func (x *PullUserStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserStatusRequest.ProtoReflect.Descriptor instead.
func (*PullUserStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{2}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PullUserStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserStatusReply) Reset() {
	*x = PullUserStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserStatusReply) ProtoMessage() {}

func (x *PullUserStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserStatusReply.ProtoReflect.Descriptor instead.
func (*PullUserStatusReply) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{5}
}

type PullUserCredentialsBscRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserCredentialsBscRequest) Reset() {
	*x = PullUserCredentialsBscRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserCredentialsBscRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserCredentialsBscRequest) ProtoMessage() {}

func (x *PullUserCredentialsBscRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserCredentialsBscRequest.ProtoReflect.Descriptor instead.
func (*PullUserCredentialsBscRequest) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{6}
}

type PullUserCredentialsBscReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserCredentialsBscReply) Reset() {
	*x = PullUserCredentialsBscReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserCredentialsBscReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserCredentialsBscReply) ProtoMessage() {}

func (x *PullUserCredentialsBscReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserCredentialsBscReply.ProtoReflect.Descriptor instead.
func (*PullUserCredentialsBscReply) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{7}
}

type SetUserRequest_SendBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Apikey    string `protobuf:"bytes,2,opt,name=apikey,proto3" json:"apikey,omitempty"`
	Apisecret string `protobuf:"bytes,3,opt,name=apisecret,proto3" json:"apisecret,omitempty"`
}

func (x *SetUserRequest_SendBody) Reset() {
	*x = SetUserRequest_SendBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRequest_SendBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRequest_SendBody) ProtoMessage() {}

func (x *SetUserRequest_SendBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserRequest_SendBody.ProtoReflect.Descriptor instead.
func (*SetUserRequest_SendBody) Descriptor() ([]byte, []int) {
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SetUserRequest_SendBody) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SetUserRequest_SendBody) GetApikey() string {
	if x != nil {
		return x.Apikey
	}
	return ""
}

func (x *SetUserRequest_SendBody) GetApisecret() string {
	if x != nil {
		return x.Apisecret
	}
	return ""
}

var File_api_binanceexchange_user_v1_binanceexchange_user_proto protoreflect.FileDescriptor

var file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDesc = []byte{
	0x0a, 0x36, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x6f, 0x64, 0x79,
	0x1a, 0x5a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15,
	0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x75, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1f, 0x0a, 0x1d, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x73, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x73, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x32, 0xd0, 0x03, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x60, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x3a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x22,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x55, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x72, 0x0a, 0x0e, 0x50, 0x75, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2c, 0x12, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x6c,
	0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x93, 0x01,
	0x0a, 0x16, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x73, 0x63, 0x12, 0x1e, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x73,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x73,
	0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f,
	0x62, 0x73, 0x63, 0x42, 0x4b, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x33, 0x62, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescOnce sync.Once
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescData = file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDesc
)

func file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescGZIP() []byte {
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescOnce.Do(func() {
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescData)
	})
	return file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDescData
}

var file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_binanceexchange_user_v1_binanceexchange_user_proto_goTypes = []interface{}{
	(*SetUserRequest)(nil),                // 0: SetUserRequest
	(*SetUserReply)(nil),                  // 1: SetUserReply
	(*PullUserStatusRequest)(nil),         // 2: PullUserStatusRequest
	(*GetUserRequest)(nil),                // 3: GetUserRequest
	(*GetUserReply)(nil),                  // 4: GetUserReply
	(*PullUserStatusReply)(nil),           // 5: PullUserStatusReply
	(*PullUserCredentialsBscRequest)(nil), // 6: PullUserCredentialsBscRequest
	(*PullUserCredentialsBscReply)(nil),   // 7: PullUserCredentialsBscReply
	(*SetUserRequest_SendBody)(nil),       // 8: SetUserRequest.SendBody
}
var file_api_binanceexchange_user_v1_binanceexchange_user_proto_depIdxs = []int32{
	8, // 0: SetUserRequest.send_body:type_name -> SetUserRequest.SendBody
	0, // 1: BinanceUser.SetUser:input_type -> SetUserRequest
	3, // 2: BinanceUser.GetUser:input_type -> GetUserRequest
	2, // 3: BinanceUser.PullUserStatus:input_type -> PullUserStatusRequest
	6, // 4: BinanceUser.PullUserCredentialsBsc:input_type -> PullUserCredentialsBscRequest
	1, // 5: BinanceUser.SetUser:output_type -> SetUserReply
	4, // 6: BinanceUser.GetUser:output_type -> GetUserReply
	5, // 7: BinanceUser.PullUserStatus:output_type -> PullUserStatusReply
	7, // 8: BinanceUser.PullUserCredentialsBsc:output_type -> PullUserCredentialsBscReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_binanceexchange_user_v1_binanceexchange_user_proto_init() }
func file_api_binanceexchange_user_v1_binanceexchange_user_proto_init() {
	if File_api_binanceexchange_user_v1_binanceexchange_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRequest); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserReply); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserStatusRequest); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserStatusReply); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserCredentialsBscRequest); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserCredentialsBscReply); i {
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
		file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRequest_SendBody); i {
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
			RawDescriptor: file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_binanceexchange_user_v1_binanceexchange_user_proto_goTypes,
		DependencyIndexes: file_api_binanceexchange_user_v1_binanceexchange_user_proto_depIdxs,
		MessageInfos:      file_api_binanceexchange_user_v1_binanceexchange_user_proto_msgTypes,
	}.Build()
	File_api_binanceexchange_user_v1_binanceexchange_user_proto = out.File
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_rawDesc = nil
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_goTypes = nil
	file_api_binanceexchange_user_v1_binanceexchange_user_proto_depIdxs = nil
}
