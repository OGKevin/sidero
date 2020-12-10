// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SystemInformation struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Manufacturer         string   `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName          string   `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	SerialNumber         string   `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	SkuNumber            string   `protobuf:"bytes,6,opt,name=sku_number,json=skuNumber,proto3" json:"sku_number,omitempty"`
	Family               string   `protobuf:"bytes,7,opt,name=family,proto3" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemInformation) Reset()         { *m = SystemInformation{} }
func (m *SystemInformation) String() string { return proto.CompactTextString(m) }
func (*SystemInformation) ProtoMessage()    {}
func (*SystemInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *SystemInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemInformation.Unmarshal(m, b)
}

func (m *SystemInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemInformation.Marshal(b, m, deterministic)
}

func (m *SystemInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInformation.Merge(m, src)
}

func (m *SystemInformation) XXX_Size() int {
	return xxx_messageInfo_SystemInformation.Size(m)
}

func (m *SystemInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInformation.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInformation proto.InternalMessageInfo

func (m *SystemInformation) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SystemInformation) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *SystemInformation) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *SystemInformation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SystemInformation) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *SystemInformation) GetSkuNumber() string {
	if m != nil {
		return m.SkuNumber
	}
	return ""
}

func (m *SystemInformation) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

type CPU struct {
	Manufacturer         string   `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}

func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}

func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}

func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}

func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *CPU) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type CreateServerRequest struct {
	SystemInformation    *SystemInformation `protobuf:"bytes,1,opt,name=system_information,json=systemInformation,proto3" json:"system_information,omitempty"`
	Cpu                  *CPU               `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Hostname             string             `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateServerRequest) Reset()         { *m = CreateServerRequest{} }
func (m *CreateServerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServerRequest) ProtoMessage()    {}
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *CreateServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServerRequest.Unmarshal(m, b)
}

func (m *CreateServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServerRequest.Marshal(b, m, deterministic)
}

func (m *CreateServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServerRequest.Merge(m, src)
}

func (m *CreateServerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServerRequest.Size(m)
}

func (m *CreateServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServerRequest proto.InternalMessageInfo

func (m *CreateServerRequest) GetSystemInformation() *SystemInformation {
	if m != nil {
		return m.SystemInformation
	}
	return nil
}

func (m *CreateServerRequest) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CreateServerRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Address struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}

func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}

func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}

func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}

func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateServerResponse struct {
	Wipe                 bool     `protobuf:"varint,1,opt,name=wipe,proto3" json:"wipe,omitempty"`
	InsecureWipe         bool     `protobuf:"varint,2,opt,name=insecure_wipe,json=insecureWipe,proto3" json:"insecure_wipe,omitempty"`
	RebootTimeout        float64  `protobuf:"fixed64,3,opt,name=reboot_timeout,json=rebootTimeout,proto3" json:"reboot_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServerResponse) Reset()         { *m = CreateServerResponse{} }
func (m *CreateServerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServerResponse) ProtoMessage()    {}
func (*CreateServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *CreateServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServerResponse.Unmarshal(m, b)
}

func (m *CreateServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServerResponse.Marshal(b, m, deterministic)
}

func (m *CreateServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServerResponse.Merge(m, src)
}

func (m *CreateServerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServerResponse.Size(m)
}

func (m *CreateServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServerResponse proto.InternalMessageInfo

func (m *CreateServerResponse) GetWipe() bool {
	if m != nil {
		return m.Wipe
	}
	return false
}

func (m *CreateServerResponse) GetInsecureWipe() bool {
	if m != nil {
		return m.InsecureWipe
	}
	return false
}

func (m *CreateServerResponse) GetRebootTimeout() float64 {
	if m != nil {
		return m.RebootTimeout
	}
	return 0
}

type MarkServerAsWipedRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkServerAsWipedRequest) Reset()         { *m = MarkServerAsWipedRequest{} }
func (m *MarkServerAsWipedRequest) String() string { return proto.CompactTextString(m) }
func (*MarkServerAsWipedRequest) ProtoMessage()    {}
func (*MarkServerAsWipedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *MarkServerAsWipedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkServerAsWipedRequest.Unmarshal(m, b)
}

func (m *MarkServerAsWipedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkServerAsWipedRequest.Marshal(b, m, deterministic)
}

func (m *MarkServerAsWipedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkServerAsWipedRequest.Merge(m, src)
}

func (m *MarkServerAsWipedRequest) XXX_Size() int {
	return xxx_messageInfo_MarkServerAsWipedRequest.Size(m)
}

func (m *MarkServerAsWipedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkServerAsWipedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkServerAsWipedRequest proto.InternalMessageInfo

func (m *MarkServerAsWipedRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type HeartbeatRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}

func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}

func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}

func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}

func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type MarkServerAsWipedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkServerAsWipedResponse) Reset()         { *m = MarkServerAsWipedResponse{} }
func (m *MarkServerAsWipedResponse) String() string { return proto.CompactTextString(m) }
func (*MarkServerAsWipedResponse) ProtoMessage()    {}
func (*MarkServerAsWipedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *MarkServerAsWipedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkServerAsWipedResponse.Unmarshal(m, b)
}

func (m *MarkServerAsWipedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkServerAsWipedResponse.Marshal(b, m, deterministic)
}

func (m *MarkServerAsWipedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkServerAsWipedResponse.Merge(m, src)
}

func (m *MarkServerAsWipedResponse) XXX_Size() int {
	return xxx_messageInfo_MarkServerAsWipedResponse.Size(m)
}

func (m *MarkServerAsWipedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkServerAsWipedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarkServerAsWipedResponse proto.InternalMessageInfo

type HeartbeatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}

func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}

func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}

func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}

func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

type ReconcileServerAddressesRequest struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Address              []*Address `protobuf:"bytes,2,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReconcileServerAddressesRequest) Reset()         { *m = ReconcileServerAddressesRequest{} }
func (m *ReconcileServerAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*ReconcileServerAddressesRequest) ProtoMessage()    {}
func (*ReconcileServerAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *ReconcileServerAddressesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReconcileServerAddressesRequest.Unmarshal(m, b)
}

func (m *ReconcileServerAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReconcileServerAddressesRequest.Marshal(b, m, deterministic)
}

func (m *ReconcileServerAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReconcileServerAddressesRequest.Merge(m, src)
}

func (m *ReconcileServerAddressesRequest) XXX_Size() int {
	return xxx_messageInfo_ReconcileServerAddressesRequest.Size(m)
}

func (m *ReconcileServerAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReconcileServerAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReconcileServerAddressesRequest proto.InternalMessageInfo

func (m *ReconcileServerAddressesRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ReconcileServerAddressesRequest) GetAddress() []*Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type ReconcileServerAddressesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReconcileServerAddressesResponse) Reset()         { *m = ReconcileServerAddressesResponse{} }
func (m *ReconcileServerAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*ReconcileServerAddressesResponse) ProtoMessage()    {}
func (*ReconcileServerAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *ReconcileServerAddressesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReconcileServerAddressesResponse.Unmarshal(m, b)
}

func (m *ReconcileServerAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReconcileServerAddressesResponse.Marshal(b, m, deterministic)
}

func (m *ReconcileServerAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReconcileServerAddressesResponse.Merge(m, src)
}

func (m *ReconcileServerAddressesResponse) XXX_Size() int {
	return xxx_messageInfo_ReconcileServerAddressesResponse.Size(m)
}

func (m *ReconcileServerAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReconcileServerAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReconcileServerAddressesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SystemInformation)(nil), "api.SystemInformation")
	proto.RegisterType((*CPU)(nil), "api.CPU")
	proto.RegisterType((*CreateServerRequest)(nil), "api.CreateServerRequest")
	proto.RegisterType((*Address)(nil), "api.Address")
	proto.RegisterType((*CreateServerResponse)(nil), "api.CreateServerResponse")
	proto.RegisterType((*MarkServerAsWipedRequest)(nil), "api.MarkServerAsWipedRequest")
	proto.RegisterType((*HeartbeatRequest)(nil), "api.HeartbeatRequest")
	proto.RegisterType((*MarkServerAsWipedResponse)(nil), "api.MarkServerAsWipedResponse")
	proto.RegisterType((*HeartbeatResponse)(nil), "api.HeartbeatResponse")
	proto.RegisterType((*ReconcileServerAddressesRequest)(nil), "api.ReconcileServerAddressesRequest")
	proto.RegisterType((*ReconcileServerAddressesResponse)(nil), "api.ReconcileServerAddressesResponse")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xda, 0xfd, 0xf5, 0x6b, 0x87, 0xa8, 0x07, 0x53, 0x16, 0x34, 0x18, 0x81, 0x4d, 0xbb,
	0x59, 0x23, 0x95, 0x0b, 0x24, 0xee, 0x46, 0x85, 0xc4, 0x84, 0x98, 0xa6, 0x8c, 0x09, 0x09, 0x09,
	0x55, 0x6e, 0xf2, 0xad, 0xb3, 0x96, 0xd8, 0xc1, 0x3f, 0x43, 0x7d, 0x0f, 0x5e, 0x89, 0x37, 0xe1,
	0x41, 0x50, 0x1c, 0x67, 0x6b, 0xd7, 0x76, 0xbb, 0xb3, 0xcf, 0xf9, 0x7e, 0xcf, 0x71, 0x02, 0x2d,
	0x5a, 0xb0, 0x5e, 0x21, 0x85, 0x16, 0xa4, 0x49, 0x0b, 0x16, 0xfe, 0xf3, 0xa0, 0x7b, 0x3e, 0x51,
	0x1a, 0xf3, 0x13, 0x7e, 0x29, 0x64, 0x4e, 0x35, 0x13, 0x9c, 0x10, 0x58, 0x31, 0x86, 0xa5, 0xbe,
	0xb7, 0xe7, 0x1d, 0xb6, 0x62, 0x7b, 0x26, 0x21, 0x74, 0x72, 0xca, 0xcd, 0x25, 0x4d, 0xb4, 0x91,
	0x28, 0xfd, 0x86, 0xe5, 0x66, 0x30, 0xf2, 0x1a, 0x3a, 0x85, 0x14, 0xa9, 0x49, 0xf4, 0x90, 0xd3,
	0x1c, 0xfd, 0xa6, 0x8d, 0x69, 0x3b, 0xec, 0x94, 0xe6, 0x48, 0x7c, 0x58, 0xbf, 0x41, 0xa9, 0x98,
	0xe0, 0xfe, 0x8a, 0x65, 0xeb, 0x2b, 0x79, 0x03, 0x9b, 0x0a, 0x25, 0xa3, 0xd9, 0x90, 0x9b, 0x7c,
	0x84, 0xd2, 0x5f, 0xad, 0x3a, 0x54, 0xe0, 0xa9, 0xc5, 0xc8, 0x2e, 0x80, 0xba, 0x36, 0x75, 0xc4,
	0x9a, 0x8d, 0x68, 0xa9, 0x6b, 0xe3, 0xe8, 0x6d, 0x58, 0xbb, 0xa4, 0x39, 0xcb, 0x26, 0xfe, 0xba,
	0xa5, 0xdc, 0x2d, 0x1c, 0x40, 0x73, 0x70, 0x76, 0x31, 0xb7, 0x83, 0xb7, 0x60, 0x87, 0xa9, 0x01,
	0x1b, 0x33, 0x03, 0x86, 0x7f, 0x3c, 0xd8, 0x1a, 0x48, 0xa4, 0x1a, 0xcf, 0x51, 0xde, 0xa0, 0x8c,
	0xf1, 0x97, 0x41, 0xa5, 0xc9, 0x27, 0x20, 0xca, 0x4a, 0x38, 0x64, 0x77, 0x1a, 0xda, 0xda, 0xed,
	0xfe, 0x76, 0xaf, 0x14, 0x7c, 0x4e, 0xe1, 0xb8, 0xab, 0xe6, 0x44, 0x0f, 0xa0, 0x99, 0x14, 0xc6,
	0x36, 0x6d, 0xf7, 0x37, 0x6c, 0xde, 0xe0, 0xec, 0x22, 0x2e, 0x41, 0x12, 0xc0, 0xc6, 0x95, 0x50,
	0x7a, 0x4a, 0xd4, 0xdb, 0x7b, 0xf8, 0x1e, 0xd6, 0x8f, 0xd3, 0x54, 0xa2, 0x52, 0xa5, 0x6f, 0x7a,
	0x52, 0x60, 0xed, 0x5b, 0x79, 0x2e, 0xf7, 0xa1, 0x15, 0x5d, 0xef, 0xe3, 0xae, 0xe1, 0x0d, 0x3c,
	0x9b, 0x5d, 0x47, 0x15, 0x82, 0x2b, 0x2c, 0xab, 0xfc, 0x66, 0xae, 0xca, 0x46, 0x6c, 0xcf, 0xa5,
	0x39, 0x8c, 0x2b, 0x4c, 0x8c, 0xc4, 0xa1, 0x25, 0x1b, 0x96, 0xec, 0xd4, 0xe0, 0xf7, 0x32, 0x68,
	0x1f, 0x9e, 0x48, 0x1c, 0x09, 0xa1, 0x87, 0x9a, 0xe5, 0x28, 0x8c, 0xb6, 0xb3, 0x7a, 0xf1, 0x66,
	0x85, 0x7e, 0xab, 0xc0, 0xb0, 0x07, 0xfe, 0x57, 0x2a, 0xaf, 0xab, 0xae, 0xc7, 0xaa, 0x4c, 0x4d,
	0x6b, 0x2d, 0x17, 0xbc, 0xbc, 0xf0, 0x00, 0x9e, 0x7e, 0x46, 0x2a, 0xf5, 0x08, 0xa9, 0x7e, 0x28,
	0xee, 0x05, 0xec, 0x2c, 0xa8, 0x5b, 0x2d, 0x15, 0x6e, 0x41, 0x77, 0xaa, 0x88, 0x03, 0x7f, 0xc2,
	0xab, 0x18, 0x13, 0xc1, 0x13, 0x96, 0x39, 0x11, 0x9c, 0x92, 0xa8, 0x1e, 0x68, 0x44, 0x0e, 0xa6,
	0x25, 0x6d, 0x1e, 0xb6, 0xfb, 0x1d, 0xeb, 0x96, 0xcb, 0xbd, 0x13, 0x38, 0x84, 0xbd, 0xe5, 0xe5,
	0xab, 0x11, 0xfa, 0x7f, 0x1b, 0xb0, 0x7a, 0x3c, 0x46, 0xae, 0xc9, 0x00, 0x3a, 0xd3, 0x76, 0x10,
	0xbf, 0x7a, 0x02, 0xf3, 0x0f, 0x2e, 0xd8, 0x59, 0xc0, 0x38, 0xef, 0x62, 0xe8, 0xce, 0x69, 0x40,
	0x76, 0x6d, 0xfc, 0x32, 0xcd, 0x83, 0x97, 0xcb, 0x68, 0x57, 0x73, 0x0c, 0xfe, 0xb2, 0x35, 0xc8,
	0x5b, 0x9b, 0xfb, 0x88, 0x88, 0xc1, 0xfe, 0x23, 0x51, 0xae, 0xd1, 0x07, 0x68, 0xdd, 0x7a, 0x44,
	0x9e, 0xdb, 0x9c, 0xfb, 0xc6, 0x07, 0xdb, 0xf7, 0xe1, 0x2a, 0xf7, 0xe3, 0x97, 0x1f, 0x27, 0x63,
	0xa6, 0xaf, 0xcc, 0xa8, 0x97, 0x88, 0x3c, 0xd2, 0x34, 0x13, 0xea, 0xa8, 0xfa, 0xc6, 0x54, 0xa4,
	0x58, 0x8a, 0x52, 0x44, 0xb4, 0x28, 0xa2, 0x1c, 0x35, 0xcd, 0x8e, 0x12, 0xc1, 0xb5, 0x14, 0x59,
	0x86, 0xf2, 0x28, 0xa7, 0x9c, 0x8e, 0x51, 0x46, 0x8c, 0x6b, 0x94, 0x9c, 0x66, 0x11, 0x2d, 0xd8,
	0x68, 0xcd, 0xfe, 0x21, 0xdf, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x22, 0xab, 0xa4, 0x2e,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error)
	MarkServerAsWiped(ctx context.Context, in *MarkServerAsWipedRequest, opts ...grpc.CallOption) (*MarkServerAsWipedResponse, error)
	ReconcileServerAddresses(ctx context.Context, in *ReconcileServerAddressesRequest, opts ...grpc.CallOption) (*ReconcileServerAddressesResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error) {
	out := new(CreateServerResponse)
	err := c.cc.Invoke(ctx, "/api.Agent/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) MarkServerAsWiped(ctx context.Context, in *MarkServerAsWipedRequest, opts ...grpc.CallOption) (*MarkServerAsWipedResponse, error) {
	out := new(MarkServerAsWipedResponse)
	err := c.cc.Invoke(ctx, "/api.Agent/MarkServerAsWiped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ReconcileServerAddresses(ctx context.Context, in *ReconcileServerAddressesRequest, opts ...grpc.CallOption) (*ReconcileServerAddressesResponse, error) {
	out := new(ReconcileServerAddressesResponse)
	err := c.cc.Invoke(ctx, "/api.Agent/ReconcileServerAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/api.Agent/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error)
	MarkServerAsWiped(context.Context, *MarkServerAsWipedRequest) (*MarkServerAsWipedResponse, error)
	ReconcileServerAddresses(context.Context, *ReconcileServerAddressesRequest) (*ReconcileServerAddressesResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) CreateServer(ctx context.Context, req *CreateServerRequest) (*CreateServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}

func (*UnimplementedAgentServer) MarkServerAsWiped(ctx context.Context, req *MarkServerAsWipedRequest) (*MarkServerAsWipedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkServerAsWiped not implemented")
}

func (*UnimplementedAgentServer) ReconcileServerAddresses(ctx context.Context, req *ReconcileServerAddressesRequest) (*ReconcileServerAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReconcileServerAddresses not implemented")
}

func (*UnimplementedAgentServer) Heartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Agent/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_MarkServerAsWiped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkServerAsWipedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).MarkServerAsWiped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Agent/MarkServerAsWiped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).MarkServerAsWiped(ctx, req.(*MarkServerAsWipedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ReconcileServerAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconcileServerAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ReconcileServerAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Agent/ReconcileServerAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ReconcileServerAddresses(ctx, req.(*ReconcileServerAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Agent/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServer",
			Handler:    _Agent_CreateServer_Handler,
		},
		{
			MethodName: "MarkServerAsWiped",
			Handler:    _Agent_MarkServerAsWiped_Handler,
		},
		{
			MethodName: "ReconcileServerAddresses",
			Handler:    _Agent_ReconcileServerAddresses_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Agent_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
