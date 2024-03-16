// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ips/user/v1/user.proto

package userv1

import (
	v1 "github.com/RyuChk/ips-map-service/internal/gen/proto/ips/shared/rssi/v1"
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

type GetCoordinateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signals    []*v1.RSSI     `protobuf:"bytes,1,rep,name=signals,proto3" json:"signals,omitempty"`
	DeviceInfo *v1.DeviceInfo `protobuf:"bytes,2,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
	User       string         `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	IsAdmin    bool           `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	Building   string         `protobuf:"bytes,5,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *GetCoordinateRequest) Reset() {
	*x = GetCoordinateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordinateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordinateRequest) ProtoMessage() {}

func (x *GetCoordinateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordinateRequest.ProtoReflect.Descriptor instead.
func (*GetCoordinateRequest) Descriptor() ([]byte, []int) {
	return file_ips_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetCoordinateRequest) GetSignals() []*v1.RSSI {
	if x != nil {
		return x.Signals
	}
	return nil
}

func (x *GetCoordinateRequest) GetDeviceInfo() *v1.DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *GetCoordinateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetCoordinateRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *GetCoordinateRequest) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

type GetCoordinateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *v1.Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Label    string       `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Building string       `protobuf:"bytes,3,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *GetCoordinateResponse) Reset() {
	*x = GetCoordinateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordinateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordinateResponse) ProtoMessage() {}

func (x *GetCoordinateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordinateResponse.ProtoReflect.Descriptor instead.
func (*GetCoordinateResponse) Descriptor() ([]byte, []int) {
	return file_ips_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetCoordinateResponse) GetPosition() *v1.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *GetCoordinateResponse) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *GetCoordinateResponse) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

type RegisterApRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssid       string `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterApRequest) Reset() {
	*x = RegisterApRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterApRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterApRequest) ProtoMessage() {}

func (x *RegisterApRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterApRequest.ProtoReflect.Descriptor instead.
func (*RegisterApRequest) Descriptor() ([]byte, []int) {
	return file_ips_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterApRequest) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *RegisterApRequest) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *RegisterApRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterApResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterApResponse) Reset() {
	*x = RegisterApResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterApResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterApResponse) ProtoMessage() {}

func (x *RegisterApResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterApResponse.ProtoReflect.Descriptor instead.
func (*RegisterApResponse) Descriptor() ([]byte, []int) {
	return file_ips_user_v1_user_proto_rawDescGZIP(), []int{3}
}

var File_ips_user_v1_user_proto protoreflect.FileDescriptor

var file_ips_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x70, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x72, 0x73, 0x73, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x07, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x69, 0x70, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x53, 0x53, 0x49, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x83, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x70, 0x73, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x22, 0x5c, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbb, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x70, 0x12, 0x1e, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb3, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x70,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x79, 0x75, 0x43, 0x68, 0x6b, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x6d, 0x61,
	0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x49, 0x55, 0x58, 0xaa, 0x02, 0x0b, 0x49, 0x70, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x49, 0x70, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x17, 0x49, 0x70, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x49, 0x70,
	0x73, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ips_user_v1_user_proto_rawDescOnce sync.Once
	file_ips_user_v1_user_proto_rawDescData = file_ips_user_v1_user_proto_rawDesc
)

func file_ips_user_v1_user_proto_rawDescGZIP() []byte {
	file_ips_user_v1_user_proto_rawDescOnce.Do(func() {
		file_ips_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_user_v1_user_proto_rawDescData)
	})
	return file_ips_user_v1_user_proto_rawDescData
}

var file_ips_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ips_user_v1_user_proto_goTypes = []interface{}{
	(*GetCoordinateRequest)(nil),  // 0: ips.user.v1.GetCoordinateRequest
	(*GetCoordinateResponse)(nil), // 1: ips.user.v1.GetCoordinateResponse
	(*RegisterApRequest)(nil),     // 2: ips.user.v1.RegisterApRequest
	(*RegisterApResponse)(nil),    // 3: ips.user.v1.RegisterApResponse
	(*v1.RSSI)(nil),               // 4: ips.shared.rssi.v1.RSSI
	(*v1.DeviceInfo)(nil),         // 5: ips.shared.rssi.v1.DeviceInfo
	(*v1.Position)(nil),           // 6: ips.shared.rssi.v1.Position
}
var file_ips_user_v1_user_proto_depIdxs = []int32{
	4, // 0: ips.user.v1.GetCoordinateRequest.signals:type_name -> ips.shared.rssi.v1.RSSI
	5, // 1: ips.user.v1.GetCoordinateRequest.device_info:type_name -> ips.shared.rssi.v1.DeviceInfo
	6, // 2: ips.user.v1.GetCoordinateResponse.position:type_name -> ips.shared.rssi.v1.Position
	0, // 3: ips.user.v1.UserManagerService.GetCoordinate:input_type -> ips.user.v1.GetCoordinateRequest
	2, // 4: ips.user.v1.UserManagerService.RegisterAp:input_type -> ips.user.v1.RegisterApRequest
	1, // 5: ips.user.v1.UserManagerService.GetCoordinate:output_type -> ips.user.v1.GetCoordinateResponse
	3, // 6: ips.user.v1.UserManagerService.RegisterAp:output_type -> ips.user.v1.RegisterApResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ips_user_v1_user_proto_init() }
func file_ips_user_v1_user_proto_init() {
	if File_ips_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordinateRequest); i {
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
		file_ips_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordinateResponse); i {
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
		file_ips_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterApRequest); i {
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
		file_ips_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterApResponse); i {
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
			RawDescriptor: file_ips_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ips_user_v1_user_proto_goTypes,
		DependencyIndexes: file_ips_user_v1_user_proto_depIdxs,
		MessageInfos:      file_ips_user_v1_user_proto_msgTypes,
	}.Build()
	File_ips_user_v1_user_proto = out.File
	file_ips_user_v1_user_proto_rawDesc = nil
	file_ips_user_v1_user_proto_goTypes = nil
	file_ips_user_v1_user_proto_depIdxs = nil
}
