// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BeanProtocol.proto

package BeanProtocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//设备就相当于一个盒子
type Device struct {
	//每个盒子都有一个唯一的ID
	DeviceID uint64 `protobuf:"fixed64,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	//设备连接下的ID地址，可能包含内网地址，外网地址
	Ip                   []*NetAddress `protobuf:"bytes,2,rep,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb77771335fe271b, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetDeviceID() uint64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *Device) GetIp() []*NetAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

//客户端相当于一个远程节点
type Client struct {
	//用户唯一ID
	Uuid uint64 `protobuf:"fixed64,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	//一个用户可能会有多个设备 repeated
	Device               []*Device `protobuf:"bytes,2,rep,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb77771335fe271b, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetUuid() uint64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *Client) GetDevice() []*Device {
	if m != nil {
		return m.Device
	}
	return nil
}

//repeated 数组类型
type NetAddress struct {
	//若这个IP为内网IP则为true，若位外网IP则为false
	Lan                  bool     `protobuf:"varint,1,opt,name=lan,proto3" json:"lan,omitempty"`
	Ip                   uint32   `protobuf:"fixed32,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetAddress) Reset()         { *m = NetAddress{} }
func (m *NetAddress) String() string { return proto.CompactTextString(m) }
func (*NetAddress) ProtoMessage()    {}
func (*NetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb77771335fe271b, []int{2}
}

func (m *NetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetAddress.Unmarshal(m, b)
}
func (m *NetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetAddress.Marshal(b, m, deterministic)
}
func (m *NetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAddress.Merge(m, src)
}
func (m *NetAddress) XXX_Size() int {
	return xxx_messageInfo_NetAddress.Size(m)
}
func (m *NetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NetAddress proto.InternalMessageInfo

func (m *NetAddress) GetLan() bool {
	if m != nil {
		return m.Lan
	}
	return false
}

func (m *NetAddress) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *NetAddress) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Device)(nil), "Device")
	proto.RegisterType((*Client)(nil), "Client")
	proto.RegisterType((*NetAddress)(nil), "NetAddress")
}

func init() { proto.RegisterFile("BeanProtocol.proto", fileDescriptor_bb77771335fe271b) }

var fileDescriptor_bb77771335fe271b = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0xca, 0x83, 0x30,
	0x14, 0x46, 0x89, 0xfe, 0x44, 0xb9, 0xf2, 0x97, 0x72, 0x27, 0x69, 0x87, 0x8a, 0x53, 0x26, 0x87,
	0x76, 0xee, 0xa0, 0x75, 0xe9, 0x52, 0x4a, 0xde, 0xc0, 0x9a, 0x0c, 0x01, 0x31, 0x21, 0xc6, 0x3e,
	0x7f, 0x49, 0x0c, 0xed, 0x76, 0x12, 0xf8, 0xce, 0xe1, 0x02, 0x76, 0x72, 0x98, 0x9f, 0x56, 0x3b,
	0x3d, 0xea, 0xa9, 0x31, 0x1e, 0xea, 0x16, 0x68, 0x2f, 0xdf, 0x6a, 0x94, 0x78, 0x80, 0x7c, 0xa3,
	0x7b, 0x5f, 0x92, 0x8a, 0x30, 0xca, 0xbf, 0x6f, 0x3c, 0x42, 0xa2, 0x4c, 0x99, 0x54, 0x29, 0x2b,
	0xce, 0x45, 0xf3, 0x90, 0xae, 0x15, 0xc2, 0xca, 0x65, 0xe1, 0x89, 0x32, 0xf5, 0x15, 0xe8, 0x6d,
	0x52, 0x72, 0x76, 0x88, 0xf0, 0xb7, 0xae, 0x4a, 0xc4, 0x79, 0x60, 0x3c, 0x01, 0x15, 0x41, 0x13,
	0xe7, 0x59, 0xb3, 0x59, 0x79, 0xfc, 0xae, 0x3b, 0x80, 0x9f, 0x10, 0xf7, 0x90, 0x4e, 0xc3, 0x1c,
	0x0c, 0x39, 0xf7, 0x88, 0xbb, 0xd8, 0x26, 0x2c, 0xf3, 0x39, 0x1f, 0x31, 0xda, 0xba, 0x32, 0xad,
	0x08, 0xfb, 0xe7, 0x81, 0x5f, 0x34, 0x1c, 0x73, 0xf9, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20, 0x7a,
	0x32, 0x27, 0xe2, 0x00, 0x00, 0x00,
}