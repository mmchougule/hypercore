// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: pkg/api/services/vms.proto

package microvm

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	types "vistara-node/pkg/api/types"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMicroVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Microvm  *types.MicroVMSpec    `protobuf:"bytes,1,opt,name=microvm,proto3" json:"microvm,omitempty"`
	Metadata map[string]*anypb.Any `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateMicroVMRequest) Reset() {
	*x = CreateMicroVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMicroVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMicroVMRequest) ProtoMessage() {}

func (x *CreateMicroVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMicroVMRequest.ProtoReflect.Descriptor instead.
func (*CreateMicroVMRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMicroVMRequest) GetMicrovm() *types.MicroVMSpec {
	if x != nil {
		return x.Microvm
	}
	return nil
}

func (x *CreateMicroVMRequest) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateMicroVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Microvm *types.MicroVM `protobuf:"bytes,1,opt,name=microvm,proto3" json:"microvm,omitempty"`
}

func (x *CreateMicroVMResponse) Reset() {
	*x = CreateMicroVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMicroVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMicroVMResponse) ProtoMessage() {}

func (x *CreateMicroVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMicroVMResponse.ProtoReflect.Descriptor instead.
func (*CreateMicroVMResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMicroVMResponse) GetMicrovm() *types.MicroVM {
	if x != nil {
		return x.Microvm
	}
	return nil
}

type DeleteMicroVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMicroVMRequest) Reset() {
	*x = DeleteMicroVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMicroVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMicroVMRequest) ProtoMessage() {}

func (x *DeleteMicroVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMicroVMRequest.ProtoReflect.Descriptor instead.
func (*DeleteMicroVMRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMicroVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMicroVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMicroVMRequest) Reset() {
	*x = GetMicroVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMicroVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMicroVMRequest) ProtoMessage() {}

func (x *GetMicroVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMicroVMRequest.ProtoReflect.Descriptor instead.
func (*GetMicroVMRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{3}
}

func (x *GetMicroVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMicroVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Microvm *types.MicroVM `protobuf:"bytes,1,opt,name=microvm,proto3" json:"microvm,omitempty"`
}

func (x *GetMicroVMResponse) Reset() {
	*x = GetMicroVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMicroVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMicroVMResponse) ProtoMessage() {}

func (x *GetMicroVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMicroVMResponse.ProtoReflect.Descriptor instead.
func (*GetMicroVMResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{4}
}

func (x *GetMicroVMResponse) GetMicrovm() *types.MicroVM {
	if x != nil {
		return x.Microvm
	}
	return nil
}

type RuntimeMicroVM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Microvm     *types.MicroVM            `protobuf:"bytes,1,opt,name=microvm,proto3" json:"microvm,omitempty"`
	RuntimeData *types.MicroVMRuntimeData `protobuf:"bytes,2,opt,name=runtime_data,json=runtimeData,proto3" json:"runtime_data,omitempty"`
}

func (x *RuntimeMicroVM) Reset() {
	*x = RuntimeMicroVM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeMicroVM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeMicroVM) ProtoMessage() {}

func (x *RuntimeMicroVM) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeMicroVM.ProtoReflect.Descriptor instead.
func (*RuntimeMicroVM) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{5}
}

func (x *RuntimeMicroVM) GetMicrovm() *types.MicroVM {
	if x != nil {
		return x.Microvm
	}
	return nil
}

func (x *RuntimeMicroVM) GetRuntimeData() *types.MicroVMRuntimeData {
	if x != nil {
		return x.RuntimeData
	}
	return nil
}

type ListMicroVMsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Microvm []*RuntimeMicroVM `protobuf:"bytes,1,rep,name=microvm,proto3" json:"microvm,omitempty"`
}

func (x *ListMicroVMsResponse) Reset() {
	*x = ListMicroVMsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_services_vms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMicroVMsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMicroVMsResponse) ProtoMessage() {}

func (x *ListMicroVMsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_services_vms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMicroVMsResponse.ProtoReflect.Descriptor instead.
func (*ListMicroVMsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_services_vms_proto_rawDescGZIP(), []int{6}
}

func (x *ListMicroVMsResponse) GetMicrovm() []*RuntimeMicroVM {
	if x != nil {
		return x.Microvm
	}
	return nil
}

var File_pkg_api_services_vms_proto protoreflect.FileDescriptor

var file_pkg_api_services_vms_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x76, 0x6d, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x07, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x76, 0x6d, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x76, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x56, 0x4d, 0x52, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x22, 0x6c, 0x0a, 0x0e, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x12, 0x22, 0x0a,
	0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76,
	0x6d, 0x12, 0x36, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56,
	0x4d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x56, 0x4d, 0x52, 0x07, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x32, 0xcc, 0x02, 0x0a,
	0x09, 0x56, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x25, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x25, 0x2e, 0x76, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_services_vms_proto_rawDescOnce sync.Once
	file_pkg_api_services_vms_proto_rawDescData = file_pkg_api_services_vms_proto_rawDesc
)

func file_pkg_api_services_vms_proto_rawDescGZIP() []byte {
	file_pkg_api_services_vms_proto_rawDescOnce.Do(func() {
		file_pkg_api_services_vms_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_services_vms_proto_rawDescData)
	})
	return file_pkg_api_services_vms_proto_rawDescData
}

var file_pkg_api_services_vms_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_api_services_vms_proto_goTypes = []interface{}{
	(*CreateMicroVMRequest)(nil),     // 0: vm.services.api.CreateMicroVMRequest
	(*CreateMicroVMResponse)(nil),    // 1: vm.services.api.CreateMicroVMResponse
	(*DeleteMicroVMRequest)(nil),     // 2: vm.services.api.DeleteMicroVMRequest
	(*GetMicroVMRequest)(nil),        // 3: vm.services.api.GetMicroVMRequest
	(*GetMicroVMResponse)(nil),       // 4: vm.services.api.GetMicroVMResponse
	(*RuntimeMicroVM)(nil),           // 5: vm.services.api.RuntimeMicroVM
	(*ListMicroVMsResponse)(nil),     // 6: vm.services.api.ListMicroVMsResponse
	nil,                              // 7: vm.services.api.CreateMicroVMRequest.MetadataEntry
	(*types.MicroVMSpec)(nil),        // 8: MicroVMSpec
	(*types.MicroVM)(nil),            // 9: MicroVM
	(*types.MicroVMRuntimeData)(nil), // 10: MicroVMRuntimeData
	(*anypb.Any)(nil),                // 11: google.protobuf.Any
	(*emptypb.Empty)(nil),            // 12: google.protobuf.Empty
}
var file_pkg_api_services_vms_proto_depIdxs = []int32{
	8,  // 0: vm.services.api.CreateMicroVMRequest.microvm:type_name -> MicroVMSpec
	7,  // 1: vm.services.api.CreateMicroVMRequest.metadata:type_name -> vm.services.api.CreateMicroVMRequest.MetadataEntry
	9,  // 2: vm.services.api.CreateMicroVMResponse.microvm:type_name -> MicroVM
	9,  // 3: vm.services.api.GetMicroVMResponse.microvm:type_name -> MicroVM
	9,  // 4: vm.services.api.RuntimeMicroVM.microvm:type_name -> MicroVM
	10, // 5: vm.services.api.RuntimeMicroVM.runtime_data:type_name -> MicroVMRuntimeData
	5,  // 6: vm.services.api.ListMicroVMsResponse.microvm:type_name -> vm.services.api.RuntimeMicroVM
	11, // 7: vm.services.api.CreateMicroVMRequest.MetadataEntry.value:type_name -> google.protobuf.Any
	0,  // 8: vm.services.api.VMService.Create:input_type -> vm.services.api.CreateMicroVMRequest
	2,  // 9: vm.services.api.VMService.Delete:input_type -> vm.services.api.DeleteMicroVMRequest
	3,  // 10: vm.services.api.VMService.Get:input_type -> vm.services.api.GetMicroVMRequest
	12, // 11: vm.services.api.VMService.List:input_type -> google.protobuf.Empty
	1,  // 12: vm.services.api.VMService.Create:output_type -> vm.services.api.CreateMicroVMResponse
	12, // 13: vm.services.api.VMService.Delete:output_type -> google.protobuf.Empty
	4,  // 14: vm.services.api.VMService.Get:output_type -> vm.services.api.GetMicroVMResponse
	6,  // 15: vm.services.api.VMService.List:output_type -> vm.services.api.ListMicroVMsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_api_services_vms_proto_init() }
func file_pkg_api_services_vms_proto_init() {
	if File_pkg_api_services_vms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_services_vms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMicroVMRequest); i {
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
		file_pkg_api_services_vms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMicroVMResponse); i {
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
		file_pkg_api_services_vms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMicroVMRequest); i {
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
		file_pkg_api_services_vms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMicroVMRequest); i {
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
		file_pkg_api_services_vms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMicroVMResponse); i {
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
		file_pkg_api_services_vms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeMicroVM); i {
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
		file_pkg_api_services_vms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMicroVMsResponse); i {
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
			RawDescriptor: file_pkg_api_services_vms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_services_vms_proto_goTypes,
		DependencyIndexes: file_pkg_api_services_vms_proto_depIdxs,
		MessageInfos:      file_pkg_api_services_vms_proto_msgTypes,
	}.Build()
	File_pkg_api_services_vms_proto = out.File
	file_pkg_api_services_vms_proto_rawDesc = nil
	file_pkg_api_services_vms_proto_goTypes = nil
	file_pkg_api_services_vms_proto_depIdxs = nil
}
