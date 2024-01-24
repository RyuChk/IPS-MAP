// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/map/v1/map.proto

package mapv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddFloorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Number      int32  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Building    string `protobuf:"bytes,4,opt,name=building,proto3" json:"building,omitempty"`
	Symbol      string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *AddFloorRequest) Reset() {
	*x = AddFloorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFloorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFloorRequest) ProtoMessage() {}

func (x *AddFloorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFloorRequest.ProtoReflect.Descriptor instead.
func (*AddFloorRequest) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{0}
}

func (x *AddFloorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddFloorRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddFloorRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *AddFloorRequest) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *AddFloorRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type AddFloorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *FloorDetail           `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AddFloorResponse) Reset() {
	*x = AddFloorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFloorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFloorResponse) ProtoMessage() {}

func (x *AddFloorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFloorResponse.ProtoReflect.Descriptor instead.
func (*AddFloorResponse) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{1}
}

func (x *AddFloorResponse) GetInfo() *FloorDetail {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *AddFloorResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AddFloorResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FloorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Number      int32  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Building    string `protobuf:"bytes,4,opt,name=building,proto3" json:"building,omitempty"`
	Symbol      string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *FloorDetail) Reset() {
	*x = FloorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloorDetail) ProtoMessage() {}

func (x *FloorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloorDetail.ProtoReflect.Descriptor instead.
func (*FloorDetail) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{2}
}

func (x *FloorDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FloorDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FloorDetail) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *FloorDetail) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *FloorDetail) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type GetFloorListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Building string `protobuf:"bytes,1,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *GetFloorListRequest) Reset() {
	*x = GetFloorListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFloorListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFloorListRequest) ProtoMessage() {}

func (x *GetFloorListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFloorListRequest.ProtoReflect.Descriptor instead.
func (*GetFloorListRequest) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{3}
}

func (x *GetFloorListRequest) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

type GetFloorListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Floors []*FloorDetail `protobuf:"bytes,1,rep,name=floors,proto3" json:"floors,omitempty"`
}

func (x *GetFloorListResponse) Reset() {
	*x = GetFloorListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFloorListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFloorListResponse) ProtoMessage() {}

func (x *GetFloorListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFloorListResponse.ProtoReflect.Descriptor instead.
func (*GetFloorListResponse) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{4}
}

func (x *GetFloorListResponse) GetFloors() []*FloorDetail {
	if x != nil {
		return x.Floors
	}
	return nil
}

type GetMapURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FloorNumber int32  `protobuf:"varint,1,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	Building    string `protobuf:"bytes,2,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *GetMapURLRequest) Reset() {
	*x = GetMapURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapURLRequest) ProtoMessage() {}

func (x *GetMapURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapURLRequest.ProtoReflect.Descriptor instead.
func (*GetMapURLRequest) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{5}
}

func (x *GetMapURLRequest) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *GetMapURLRequest) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

type GetMapURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail    *FloorDetail           `protobuf:"bytes,1,opt,name=Detail,proto3" json:"Detail,omitempty"`
	Url       string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetMapURLResponse) Reset() {
	*x = GetMapURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapURLResponse) ProtoMessage() {}

func (x *GetMapURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapURLResponse.ProtoReflect.Descriptor instead.
func (*GetMapURLResponse) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{6}
}

func (x *GetMapURLResponse) GetDetail() *FloorDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *GetMapURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetMapURLResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AddMapURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FloorNumber int32  `protobuf:"varint,1,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	Building    string `protobuf:"bytes,2,opt,name=building,proto3" json:"building,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *AddMapURLRequest) Reset() {
	*x = AddMapURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMapURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMapURLRequest) ProtoMessage() {}

func (x *AddMapURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMapURLRequest.ProtoReflect.Descriptor instead.
func (*AddMapURLRequest) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{7}
}

func (x *AddMapURLRequest) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *AddMapURLRequest) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *AddMapURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type AddMapURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail    *FloorDetail           `protobuf:"bytes,1,opt,name=Detail,proto3" json:"Detail,omitempty"`
	Url       string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *AddMapURLResponse) Reset() {
	*x = AddMapURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_map_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMapURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMapURLResponse) ProtoMessage() {}

func (x *AddMapURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_map_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMapURLResponse.ProtoReflect.Descriptor instead.
func (*AddMapURLResponse) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_map_proto_rawDescGZIP(), []int{8}
}

func (x *AddMapURLResponse) GetDetail() *FloorDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *AddMapURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddMapURLResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AddMapURLResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_ips_map_v1_map_proto protoreflect.FileDescriptor

var file_ips_map_v1_map_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0xb5, 0x01, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69,
	0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f,
	0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x6f,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x73, 0x22,
	0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x70,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x6f, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xba, 0x02, 0x0a, 0x0a, 0x4d,
	0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x69, 0x70, 0x73, 0x2e,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x70, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c,
	0x12, 0x1c, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x70, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x2e, 0x69, 0x70, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x70, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x70, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xab, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x4d, 0x61, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x79, 0x75, 0x43, 0x68, 0x6b, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x6d, 0x61,
	0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73,
	0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x49, 0x4d, 0x58, 0xaa, 0x02, 0x0a, 0x49, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x49, 0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x49, 0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x49, 0x70, 0x73, 0x3a, 0x3a, 0x4d, 0x61,
	0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ips_map_v1_map_proto_rawDescOnce sync.Once
	file_ips_map_v1_map_proto_rawDescData = file_ips_map_v1_map_proto_rawDesc
)

func file_ips_map_v1_map_proto_rawDescGZIP() []byte {
	file_ips_map_v1_map_proto_rawDescOnce.Do(func() {
		file_ips_map_v1_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_map_v1_map_proto_rawDescData)
	})
	return file_ips_map_v1_map_proto_rawDescData
}

var file_ips_map_v1_map_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ips_map_v1_map_proto_goTypes = []interface{}{
	(*AddFloorRequest)(nil),       // 0: ips.map.v1.AddFloorRequest
	(*AddFloorResponse)(nil),      // 1: ips.map.v1.AddFloorResponse
	(*FloorDetail)(nil),           // 2: ips.map.v1.FloorDetail
	(*GetFloorListRequest)(nil),   // 3: ips.map.v1.GetFloorListRequest
	(*GetFloorListResponse)(nil),  // 4: ips.map.v1.GetFloorListResponse
	(*GetMapURLRequest)(nil),      // 5: ips.map.v1.GetMapURLRequest
	(*GetMapURLResponse)(nil),     // 6: ips.map.v1.GetMapURLResponse
	(*AddMapURLRequest)(nil),      // 7: ips.map.v1.AddMapURLRequest
	(*AddMapURLResponse)(nil),     // 8: ips.map.v1.AddMapURLResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_ips_map_v1_map_proto_depIdxs = []int32{
	2,  // 0: ips.map.v1.AddFloorResponse.info:type_name -> ips.map.v1.FloorDetail
	9,  // 1: ips.map.v1.AddFloorResponse.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: ips.map.v1.AddFloorResponse.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 3: ips.map.v1.GetFloorListResponse.floors:type_name -> ips.map.v1.FloorDetail
	2,  // 4: ips.map.v1.GetMapURLResponse.Detail:type_name -> ips.map.v1.FloorDetail
	9,  // 5: ips.map.v1.GetMapURLResponse.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 6: ips.map.v1.AddMapURLResponse.Detail:type_name -> ips.map.v1.FloorDetail
	9,  // 7: ips.map.v1.AddMapURLResponse.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 8: ips.map.v1.AddMapURLResponse.created_at:type_name -> google.protobuf.Timestamp
	3,  // 9: ips.map.v1.MapService.GetFloorList:input_type -> ips.map.v1.GetFloorListRequest
	0,  // 10: ips.map.v1.MapService.AddFloor:input_type -> ips.map.v1.AddFloorRequest
	5,  // 11: ips.map.v1.MapService.GetMapURL:input_type -> ips.map.v1.GetMapURLRequest
	7,  // 12: ips.map.v1.MapService.AddMapURL:input_type -> ips.map.v1.AddMapURLRequest
	4,  // 13: ips.map.v1.MapService.GetFloorList:output_type -> ips.map.v1.GetFloorListResponse
	1,  // 14: ips.map.v1.MapService.AddFloor:output_type -> ips.map.v1.AddFloorResponse
	6,  // 15: ips.map.v1.MapService.GetMapURL:output_type -> ips.map.v1.GetMapURLResponse
	8,  // 16: ips.map.v1.MapService.AddMapURL:output_type -> ips.map.v1.AddMapURLResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_ips_map_v1_map_proto_init() }
func file_ips_map_v1_map_proto_init() {
	if File_ips_map_v1_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_map_v1_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFloorRequest); i {
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
		file_ips_map_v1_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFloorResponse); i {
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
		file_ips_map_v1_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloorDetail); i {
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
		file_ips_map_v1_map_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFloorListRequest); i {
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
		file_ips_map_v1_map_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFloorListResponse); i {
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
		file_ips_map_v1_map_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapURLRequest); i {
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
		file_ips_map_v1_map_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapURLResponse); i {
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
		file_ips_map_v1_map_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMapURLRequest); i {
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
		file_ips_map_v1_map_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMapURLResponse); i {
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
			RawDescriptor: file_ips_map_v1_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ips_map_v1_map_proto_goTypes,
		DependencyIndexes: file_ips_map_v1_map_proto_depIdxs,
		MessageInfos:      file_ips_map_v1_map_proto_msgTypes,
	}.Build()
	File_ips_map_v1_map_proto = out.File
	file_ips_map_v1_map_proto_rawDesc = nil
	file_ips_map_v1_map_proto_goTypes = nil
	file_ips_map_v1_map_proto_depIdxs = nil
}
