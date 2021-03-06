// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rmq.proto

package rmqpb

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

// OrderHats represents a request to make hats
type OrderHats struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Size                 string   `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Style                string   `protobuf:"bytes,7,opt,name=style,proto3" json:"style,omitempty"`
	Quantity             int32    `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Notes                string   `protobuf:"bytes,9,opt,name=notes,proto3" json:"notes,omitempty"`
	Batch                string   `protobuf:"bytes,10,opt,name=batch,proto3" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderHats) Reset()         { *m = OrderHats{} }
func (m *OrderHats) String() string { return proto.CompactTextString(m) }
func (*OrderHats) ProtoMessage()    {}
func (*OrderHats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf5c3027a94e836, []int{0}
}

func (m *OrderHats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderHats.Unmarshal(m, b)
}
func (m *OrderHats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderHats.Marshal(b, m, deterministic)
}
func (m *OrderHats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderHats.Merge(m, src)
}
func (m *OrderHats) XXX_Size() int {
	return xxx_messageInfo_OrderHats.Size(m)
}
func (m *OrderHats) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderHats.DiscardUnknown(m)
}

var xxx_messageInfo_OrderHats proto.InternalMessageInfo

func (m *OrderHats) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderHats) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderHats) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *OrderHats) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *OrderHats) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *OrderHats) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *OrderHats) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *OrderHats) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderHats) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *OrderHats) GetBatch() string {
	if m != nil {
		return m.Batch
	}
	return ""
}

// HatCreated represents a hat that has been created
type HatCreated struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Size                 string   `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Style                string   `protobuf:"bytes,7,opt,name=style,proto3" json:"style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HatCreated) Reset()         { *m = HatCreated{} }
func (m *HatCreated) String() string { return proto.CompactTextString(m) }
func (*HatCreated) ProtoMessage()    {}
func (*HatCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf5c3027a94e836, []int{1}
}

func (m *HatCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HatCreated.Unmarshal(m, b)
}
func (m *HatCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HatCreated.Marshal(b, m, deterministic)
}
func (m *HatCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HatCreated.Merge(m, src)
}
func (m *HatCreated) XXX_Size() int {
	return xxx_messageInfo_HatCreated.Size(m)
}
func (m *HatCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_HatCreated.DiscardUnknown(m)
}

var xxx_messageInfo_HatCreated proto.InternalMessageInfo

func (m *HatCreated) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HatCreated) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *HatCreated) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *HatCreated) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HatCreated) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *HatCreated) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *HatCreated) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderHats)(nil), "rmq.OrderHats")
	proto.RegisterType((*HatCreated)(nil), "rmq.HatCreated")
}

func init() { proto.RegisterFile("rmq.proto", fileDescriptor_ddf5c3027a94e836) }

var fileDescriptor_ddf5c3027a94e836 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x91, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x91, 0x13, 0xc7, 0xd1, 0x0d, 0x1d, 0x44, 0x87, 0xa3, 0x50, 0x08, 0x99, 0x32, 0x75,
	0xe9, 0x13, 0xa4, 0x5d, 0xb2, 0x15, 0x32, 0x76, 0x29, 0xb2, 0x25, 0xa8, 0x20, 0xb6, 0xec, 0xd3,
	0xa5, 0x90, 0x3e, 0x57, 0xdf, 0xae, 0x4b, 0xd1, 0x29, 0xee, 0x3b, 0x64, 0xd3, 0xf7, 0xfd, 0xfc,
	0x82, 0x9f, 0x03, 0x4d, 0xfd, 0xf4, 0x34, 0x52, 0xe4, 0x68, 0x16, 0xd4, 0x4f, 0xdb, 0x5f, 0x05,
	0xfa, 0x8d, 0x9c, 0xa7, 0x83, 0xe5, 0x64, 0xee, 0xa0, 0x0a, 0x0e, 0xd5, 0x46, 0xed, 0xf4, 0xb1,
	0x0a, 0xce, 0x3c, 0x02, 0x74, 0xe4, 0x2d, 0x7b, 0xf7, 0x61, 0x19, 0x2b, 0xf1, 0xfa, 0x6a, 0xf6,
	0x9c, 0xe3, 0xf3, 0xe8, 0xe6, 0x78, 0x51, 0xe2, 0xab, 0xd9, 0xb3, 0x41, 0x68, 0xbe, 0x3c, 0xa5,
	0x10, 0x07, 0x5c, 0x6e, 0xd4, 0xae, 0x3e, 0xce, 0x68, 0x0c, 0x2c, 0x53, 0xf8, 0xf6, 0x58, 0x4b,
	0x45, 0xde, 0xe6, 0x1e, 0xea, 0x2e, 0x9e, 0x22, 0xe1, 0x4a, 0x64, 0x81, 0x6c, 0x13, 0x5f, 0x4e,
	0x1e, 0x9b, 0x62, 0x05, 0xcc, 0x03, 0xac, 0xa7, 0xb3, 0x1d, 0x38, 0xf0, 0x05, 0xd7, 0xf2, 0xf5,
	0x3f, 0xe7, 0xc6, 0x10, 0xd9, 0x27, 0xd4, 0xa5, 0x21, 0x90, 0x6d, 0x6b, 0xb9, 0xfb, 0x44, 0x28,
	0x56, 0x60, 0xfb, 0xa3, 0x00, 0x0e, 0x96, 0x5f, 0xcb, 0xa2, 0x5b, 0x99, 0xff, 0xd2, 0xbc, 0xd7,
	0xd4, 0x4f, 0x63, 0xdb, 0xae, 0xe4, 0x92, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xfd,
	0x2b, 0x72, 0xd6, 0x01, 0x00, 0x00,
}
