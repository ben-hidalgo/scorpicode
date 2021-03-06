// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hats.proto

package hatspb

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

// A Hat is a piece of headwear
type Hat struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Size                 string   `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Style                string   `protobuf:"bytes,7,opt,name=style,proto3" json:"style,omitempty"`
	Batch                string   `protobuf:"bytes,8,opt,name=batch,proto3" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hat) Reset()         { *m = Hat{} }
func (m *Hat) String() string { return proto.CompactTextString(m) }
func (*Hat) ProtoMessage()    {}
func (*Hat) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{0}
}

func (m *Hat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hat.Unmarshal(m, b)
}
func (m *Hat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hat.Marshal(b, m, deterministic)
}
func (m *Hat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hat.Merge(m, src)
}
func (m *Hat) XXX_Size() int {
	return xxx_messageInfo_Hat.Size(m)
}
func (m *Hat) XXX_DiscardUnknown() {
	xxx_messageInfo_Hat.DiscardUnknown(m)
}

var xxx_messageInfo_Hat proto.InternalMessageInfo

func (m *Hat) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hat) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Hat) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Hat) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Hat) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *Hat) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Hat) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *Hat) GetBatch() string {
	if m != nil {
		return m.Batch
	}
	return ""
}

type Order struct {
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

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Order) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Order) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *Order) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Order) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *Order) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Order) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Order) GetBatch() string {
	if m != nil {
		return m.Batch
	}
	return ""
}

type MakeHatsRequest struct {
	Size                 string   `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Color                string   `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Style                string   `protobuf:"bytes,3,opt,name=style,proto3" json:"style,omitempty"`
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Notes                string   `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeHatsRequest) Reset()         { *m = MakeHatsRequest{} }
func (m *MakeHatsRequest) String() string { return proto.CompactTextString(m) }
func (*MakeHatsRequest) ProtoMessage()    {}
func (*MakeHatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{2}
}

func (m *MakeHatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeHatsRequest.Unmarshal(m, b)
}
func (m *MakeHatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeHatsRequest.Marshal(b, m, deterministic)
}
func (m *MakeHatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeHatsRequest.Merge(m, src)
}
func (m *MakeHatsRequest) XXX_Size() int {
	return xxx_messageInfo_MakeHatsRequest.Size(m)
}
func (m *MakeHatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeHatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeHatsRequest proto.InternalMessageInfo

func (m *MakeHatsRequest) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *MakeHatsRequest) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *MakeHatsRequest) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *MakeHatsRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *MakeHatsRequest) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

type MakeHatsResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeHatsResponse) Reset()         { *m = MakeHatsResponse{} }
func (m *MakeHatsResponse) String() string { return proto.CompactTextString(m) }
func (*MakeHatsResponse) ProtoMessage()    {}
func (*MakeHatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{3}
}

func (m *MakeHatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeHatsResponse.Unmarshal(m, b)
}
func (m *MakeHatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeHatsResponse.Marshal(b, m, deterministic)
}
func (m *MakeHatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeHatsResponse.Merge(m, src)
}
func (m *MakeHatsResponse) XXX_Size() int {
	return xxx_messageInfo_MakeHatsResponse.Size(m)
}
func (m *MakeHatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeHatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MakeHatsResponse proto.InternalMessageInfo

func (m *MakeHatsResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type ListHatsRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHatsRequest) Reset()         { *m = ListHatsRequest{} }
func (m *ListHatsRequest) String() string { return proto.CompactTextString(m) }
func (*ListHatsRequest) ProtoMessage()    {}
func (*ListHatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{4}
}

func (m *ListHatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHatsRequest.Unmarshal(m, b)
}
func (m *ListHatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHatsRequest.Marshal(b, m, deterministic)
}
func (m *ListHatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHatsRequest.Merge(m, src)
}
func (m *ListHatsRequest) XXX_Size() int {
	return xxx_messageInfo_ListHatsRequest.Size(m)
}
func (m *ListHatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHatsRequest proto.InternalMessageInfo

func (m *ListHatsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListHatsRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListHatsResponse struct {
	Hats                 []*Hat   `protobuf:"bytes,1,rep,name=hats,proto3" json:"hats,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHatsResponse) Reset()         { *m = ListHatsResponse{} }
func (m *ListHatsResponse) String() string { return proto.CompactTextString(m) }
func (*ListHatsResponse) ProtoMessage()    {}
func (*ListHatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{5}
}

func (m *ListHatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHatsResponse.Unmarshal(m, b)
}
func (m *ListHatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHatsResponse.Marshal(b, m, deterministic)
}
func (m *ListHatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHatsResponse.Merge(m, src)
}
func (m *ListHatsResponse) XXX_Size() int {
	return xxx_messageInfo_ListHatsResponse.Size(m)
}
func (m *ListHatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHatsResponse proto.InternalMessageInfo

func (m *ListHatsResponse) GetHats() []*Hat {
	if m != nil {
		return m.Hats
	}
	return nil
}

func (m *ListHatsResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListHatsResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Size struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}
func (*Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{6}
}

func (m *Size) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Size.Unmarshal(m, b)
}
func (m *Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Size.Marshal(b, m, deterministic)
}
func (m *Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Size.Merge(m, src)
}
func (m *Size) XXX_Size() int {
	return xxx_messageInfo_Size.Size(m)
}
func (m *Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Size proto.InternalMessageInfo

func (m *Size) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Size) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListSizesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSizesRequest) Reset()         { *m = ListSizesRequest{} }
func (m *ListSizesRequest) String() string { return proto.CompactTextString(m) }
func (*ListSizesRequest) ProtoMessage()    {}
func (*ListSizesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{7}
}

func (m *ListSizesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSizesRequest.Unmarshal(m, b)
}
func (m *ListSizesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSizesRequest.Marshal(b, m, deterministic)
}
func (m *ListSizesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSizesRequest.Merge(m, src)
}
func (m *ListSizesRequest) XXX_Size() int {
	return xxx_messageInfo_ListSizesRequest.Size(m)
}
func (m *ListSizesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSizesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSizesRequest proto.InternalMessageInfo

type ListSizesResponse struct {
	Sizes                []*Size  `protobuf:"bytes,1,rep,name=sizes,proto3" json:"sizes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSizesResponse) Reset()         { *m = ListSizesResponse{} }
func (m *ListSizesResponse) String() string { return proto.CompactTextString(m) }
func (*ListSizesResponse) ProtoMessage()    {}
func (*ListSizesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{8}
}

func (m *ListSizesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSizesResponse.Unmarshal(m, b)
}
func (m *ListSizesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSizesResponse.Marshal(b, m, deterministic)
}
func (m *ListSizesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSizesResponse.Merge(m, src)
}
func (m *ListSizesResponse) XXX_Size() int {
	return xxx_messageInfo_ListSizesResponse.Size(m)
}
func (m *ListSizesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSizesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSizesResponse proto.InternalMessageInfo

func (m *ListSizesResponse) GetSizes() []*Size {
	if m != nil {
		return m.Sizes
	}
	return nil
}

type DeleteHatRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHatRequest) Reset()         { *m = DeleteHatRequest{} }
func (m *DeleteHatRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHatRequest) ProtoMessage()    {}
func (*DeleteHatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{9}
}

func (m *DeleteHatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHatRequest.Unmarshal(m, b)
}
func (m *DeleteHatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHatRequest.Marshal(b, m, deterministic)
}
func (m *DeleteHatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHatRequest.Merge(m, src)
}
func (m *DeleteHatRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteHatRequest.Size(m)
}
func (m *DeleteHatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHatRequest proto.InternalMessageInfo

func (m *DeleteHatRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteHatRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type DeleteHatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHatResponse) Reset()         { *m = DeleteHatResponse{} }
func (m *DeleteHatResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteHatResponse) ProtoMessage()    {}
func (*DeleteHatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{10}
}

func (m *DeleteHatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHatResponse.Unmarshal(m, b)
}
func (m *DeleteHatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHatResponse.Marshal(b, m, deterministic)
}
func (m *DeleteHatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHatResponse.Merge(m, src)
}
func (m *DeleteHatResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteHatResponse.Size(m)
}
func (m *DeleteHatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHatResponse proto.InternalMessageInfo

type FetchHatRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchHatRequest) Reset()         { *m = FetchHatRequest{} }
func (m *FetchHatRequest) String() string { return proto.CompactTextString(m) }
func (*FetchHatRequest) ProtoMessage()    {}
func (*FetchHatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{11}
}

func (m *FetchHatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchHatRequest.Unmarshal(m, b)
}
func (m *FetchHatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchHatRequest.Marshal(b, m, deterministic)
}
func (m *FetchHatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchHatRequest.Merge(m, src)
}
func (m *FetchHatRequest) XXX_Size() int {
	return xxx_messageInfo_FetchHatRequest.Size(m)
}
func (m *FetchHatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchHatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchHatRequest proto.InternalMessageInfo

func (m *FetchHatRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FetchHatResponse struct {
	Hat                  *Hat     `protobuf:"bytes,1,opt,name=hat,proto3" json:"hat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchHatResponse) Reset()         { *m = FetchHatResponse{} }
func (m *FetchHatResponse) String() string { return proto.CompactTextString(m) }
func (*FetchHatResponse) ProtoMessage()    {}
func (*FetchHatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{12}
}

func (m *FetchHatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchHatResponse.Unmarshal(m, b)
}
func (m *FetchHatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchHatResponse.Marshal(b, m, deterministic)
}
func (m *FetchHatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchHatResponse.Merge(m, src)
}
func (m *FetchHatResponse) XXX_Size() int {
	return xxx_messageInfo_FetchHatResponse.Size(m)
}
func (m *FetchHatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchHatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchHatResponse proto.InternalMessageInfo

func (m *FetchHatResponse) GetHat() *Hat {
	if m != nil {
		return m.Hat
	}
	return nil
}

type FetchOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchOrderRequest) Reset()         { *m = FetchOrderRequest{} }
func (m *FetchOrderRequest) String() string { return proto.CompactTextString(m) }
func (*FetchOrderRequest) ProtoMessage()    {}
func (*FetchOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{13}
}

func (m *FetchOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOrderRequest.Unmarshal(m, b)
}
func (m *FetchOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOrderRequest.Marshal(b, m, deterministic)
}
func (m *FetchOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOrderRequest.Merge(m, src)
}
func (m *FetchOrderRequest) XXX_Size() int {
	return xxx_messageInfo_FetchOrderRequest.Size(m)
}
func (m *FetchOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOrderRequest proto.InternalMessageInfo

func (m *FetchOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FetchOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchOrderResponse) Reset()         { *m = FetchOrderResponse{} }
func (m *FetchOrderResponse) String() string { return proto.CompactTextString(m) }
func (*FetchOrderResponse) ProtoMessage()    {}
func (*FetchOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b65fa2650eb6b0bd, []int{14}
}

func (m *FetchOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOrderResponse.Unmarshal(m, b)
}
func (m *FetchOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOrderResponse.Marshal(b, m, deterministic)
}
func (m *FetchOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOrderResponse.Merge(m, src)
}
func (m *FetchOrderResponse) XXX_Size() int {
	return xxx_messageInfo_FetchOrderResponse.Size(m)
}
func (m *FetchOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOrderResponse proto.InternalMessageInfo

func (m *FetchOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterType((*Hat)(nil), "hats.Hat")
	proto.RegisterType((*Order)(nil), "hats.Order")
	proto.RegisterType((*MakeHatsRequest)(nil), "hats.MakeHatsRequest")
	proto.RegisterType((*MakeHatsResponse)(nil), "hats.MakeHatsResponse")
	proto.RegisterType((*ListHatsRequest)(nil), "hats.ListHatsRequest")
	proto.RegisterType((*ListHatsResponse)(nil), "hats.ListHatsResponse")
	proto.RegisterType((*Size)(nil), "hats.Size")
	proto.RegisterType((*ListSizesRequest)(nil), "hats.ListSizesRequest")
	proto.RegisterType((*ListSizesResponse)(nil), "hats.ListSizesResponse")
	proto.RegisterType((*DeleteHatRequest)(nil), "hats.DeleteHatRequest")
	proto.RegisterType((*DeleteHatResponse)(nil), "hats.DeleteHatResponse")
	proto.RegisterType((*FetchHatRequest)(nil), "hats.FetchHatRequest")
	proto.RegisterType((*FetchHatResponse)(nil), "hats.FetchHatResponse")
	proto.RegisterType((*FetchOrderRequest)(nil), "hats.FetchOrderRequest")
	proto.RegisterType((*FetchOrderResponse)(nil), "hats.FetchOrderResponse")
}

func init() { proto.RegisterFile("hats.proto", fileDescriptor_b65fa2650eb6b0bd) }

var fileDescriptor_b65fa2650eb6b0bd = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0x4d, 0x6b, 0x1b, 0x31,
	0x10, 0x65, 0x3f, 0xe4, 0xd8, 0x13, 0xa8, 0x6d, 0x35, 0x4d, 0x84, 0x4b, 0xc0, 0x51, 0x2f, 0x3e,
	0xb9, 0xe0, 0x12, 0x4a, 0x20, 0x10, 0x52, 0x4a, 0xf1, 0xa1, 0xa5, 0xb0, 0xbd, 0xf5, 0x12, 0x64,
	0x5b, 0x89, 0x97, 0x6e, 0xbc, 0xce, 0x6a, 0x5c, 0x48, 0xaf, 0xfd, 0x69, 0xfd, 0x4b, 0x85, 0x5e,
	0x8b, 0xa4, 0x95, 0x57, 0xd9, 0xb5, 0xa1, 0xd7, 0xde, 0x76, 0xde, 0xe8, 0x69, 0x9e, 0xe6, 0x3d,
	0x58, 0x80, 0xa5, 0x40, 0x35, 0x5e, 0x17, 0x39, 0xe6, 0x34, 0xd6, 0xdf, 0xfc, 0x57, 0x00, 0xd1,
	0x54, 0x20, 0x7d, 0x06, 0x61, 0xba, 0x60, 0xc1, 0x30, 0x18, 0x75, 0x92, 0x30, 0x5d, 0xd0, 0x53,
	0x80, 0x79, 0x21, 0x05, 0xca, 0xc5, 0x8d, 0x40, 0x16, 0x1a, 0xbc, 0x53, 0x22, 0xd7, 0xa8, 0xdb,
	0x9b, 0xf5, 0xc2, 0xb5, 0x23, 0xdb, 0x2e, 0x91, 0x6b, 0xa4, 0x0c, 0x0e, 0xbe, 0xcb, 0x42, 0xa5,
	0xf9, 0x8a, 0xc5, 0xc3, 0x60, 0x44, 0x12, 0x57, 0x52, 0x0a, 0xb1, 0x4a, 0x7f, 0x48, 0x46, 0x0c,
	0xc5, 0x7c, 0xd3, 0x23, 0x20, 0xf3, 0x3c, 0xcb, 0x0b, 0xd6, 0x32, 0xa0, 0x2d, 0x34, 0xaa, 0xf0,
	0x31, 0x93, 0xec, 0xc0, 0xa2, 0xa6, 0xd0, 0xe8, 0x4c, 0xe0, 0x7c, 0xc9, 0xda, 0x16, 0x35, 0x05,
	0xff, 0x13, 0x00, 0xf9, 0x5c, 0x2c, 0x64, 0xf1, 0xdf, 0xbc, 0x63, 0x00, 0xed, 0x87, 0x8d, 0x58,
	0x61, 0x8a, 0x8f, 0xe6, 0x29, 0x24, 0xd9, 0xd6, 0x9a, 0xb1, 0xca, 0x51, 0x2a, 0xd6, 0xb1, 0x0c,
	0x53, 0x54, 0x2f, 0x07, 0xff, 0xe5, 0x3f, 0x03, 0xe8, 0x7e, 0x12, 0xdf, 0xe4, 0x54, 0xa0, 0x4a,
	0xe4, 0xc3, 0x46, 0x2a, 0xdc, 0x6a, 0x0b, 0x76, 0x69, 0x0b, 0x77, 0x6a, 0x8b, 0xf6, 0x69, 0x8b,
	0xf7, 0x69, 0x23, 0x9e, 0x36, 0x7e, 0x0e, 0xbd, 0x4a, 0x84, 0x5a, 0xe7, 0x2b, 0x25, 0xe9, 0x19,
	0x90, 0x5c, 0x5b, 0x62, 0x64, 0x1c, 0x4e, 0x0e, 0xc7, 0x26, 0x7b, 0xc6, 0xa5, 0xc4, 0x76, 0xf8,
	0x15, 0x74, 0x3f, 0xa6, 0x0a, 0x7d, 0xed, 0x47, 0x40, 0xb2, 0xf4, 0x3e, 0x45, 0xc3, 0x22, 0x89,
	0x2d, 0xe8, 0x31, 0xb4, 0xf2, 0xdb, 0x5b, 0x25, 0xad, 0x83, 0x24, 0x29, 0x2b, 0x7e, 0x03, 0xbd,
	0xea, 0x82, 0x72, 0xee, 0x29, 0x98, 0x64, 0xb3, 0x60, 0x18, 0x8d, 0x0e, 0x27, 0x1d, 0x3b, 0x76,
	0x2a, 0x30, 0x31, 0x70, 0x35, 0x20, 0xdc, 0x3d, 0x20, 0x7a, 0x32, 0x60, 0x0c, 0xf1, 0x17, 0xbd,
	0x3e, 0xbd, 0xd2, 0x6c, 0x73, 0xb7, 0x5d, 0x69, 0xb6, 0xb9, 0xd3, 0xd8, 0x4a, 0xdc, 0xcb, 0x72,
	0xa3, 0xe6, 0x9b, 0x53, 0x2b, 0x48, 0x73, 0xdc, 0x93, 0xf8, 0x39, 0xf4, 0x3d, 0xac, 0x54, 0x39,
	0x04, 0xa2, 0x7d, 0x71, 0x32, 0xc1, 0xca, 0xd4, 0x67, 0x12, 0xdb, 0xe0, 0x97, 0xd0, 0x7b, 0x2f,
	0x33, 0x89, 0x7a, 0xab, 0x6e, 0x3b, 0xf5, 0x74, 0x7b, 0xf9, 0x0c, 0x9f, 0xe4, 0x93, 0x3f, 0x87,
	0xbe, 0xc7, 0xb6, 0x43, 0xf9, 0x19, 0x74, 0x3f, 0x48, 0x9c, 0x2f, 0xf7, 0xdf, 0xc8, 0x5f, 0x43,
	0xaf, 0x3a, 0x52, 0x6a, 0x7d, 0x09, 0xd1, 0x52, 0x60, 0xe9, 0xa3, 0xb7, 0x50, 0x8d, 0xf2, 0x57,
	0xd0, 0x37, 0x04, 0x6b, 0xec, 0x9e, 0x5b, 0xdf, 0x02, 0xf5, 0x0f, 0xfd, 0x73, 0x42, 0x26, 0xbf,
	0x43, 0x88, 0xb5, 0xbb, 0xf4, 0x02, 0xda, 0x2e, 0x61, 0xf4, 0x85, 0x3d, 0x58, 0x8b, 0xfd, 0xe0,
	0xb8, 0x0e, 0x97, 0x63, 0x2e, 0xa0, 0xed, 0x42, 0xe2, 0xa8, 0xb5, 0xd4, 0x39, 0x6a, 0x23, 0x4b,
	0x97, 0xd0, 0xd9, 0x5a, 0x47, 0xbd, 0x43, 0xbe, 0xbf, 0x83, 0x93, 0x06, 0x5e, 0xb1, 0xb7, 0x1e,
	0x38, 0x76, 0xdd, 0x52, 0xc7, 0x6e, 0x98, 0xa5, 0x65, 0x3b, 0x27, 0x9c, 0xec, 0x9a, 0x79, 0x4e,
	0x76, 0xc3, 0xb0, 0x2b, 0x80, 0x6a, 0xdd, 0xf4, 0xc4, 0x3b, 0xe5, 0xbb, 0x34, 0x60, 0xcd, 0x86,
	0xbd, 0xe0, 0x5d, 0xfb, 0x6b, 0x4b, 0xb7, 0xd6, 0xb3, 0x59, 0xcb, 0xfc, 0x2c, 0xde, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x80, 0x2d, 0xb9, 0x09, 0x3a, 0x06, 0x00, 0x00,
}
