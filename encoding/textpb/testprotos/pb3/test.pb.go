// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoding/textpb/testprotos/pb3/test.proto

package pb3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
	Enum_TEN  Enum = 10
)

var Enum_name = map[int32]string{
	0:  "ZERO",
	1:  "ONE",
	2:  "TWO",
	10: "TEN",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
	"TEN":  10,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{0}
}

type Enums_NestedEnum int32

const (
	Enums_CERO Enums_NestedEnum = 0
	Enums_UNO  Enums_NestedEnum = 1
	Enums_DOS  Enums_NestedEnum = 2
	Enums_DIEZ Enums_NestedEnum = 10
)

var Enums_NestedEnum_name = map[int32]string{
	0:  "CERO",
	1:  "UNO",
	2:  "DOS",
	10: "DIEZ",
}

var Enums_NestedEnum_value = map[string]int32{
	"CERO": 0,
	"UNO":  1,
	"DOS":  2,
	"DIEZ": 10,
}

func (x Enums_NestedEnum) String() string {
	return proto.EnumName(Enums_NestedEnum_name, int32(x))
}

func (Enums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{1, 0}
}

// Scalars contains scalar field types.
type Scalars struct {
	SBool                bool     `protobuf:"varint,1,opt,name=s_bool,json=sBool,proto3" json:"s_bool,omitempty"`
	SInt32               int32    `protobuf:"varint,2,opt,name=s_int32,json=sInt32,proto3" json:"s_int32,omitempty"`
	SInt64               int64    `protobuf:"varint,3,opt,name=s_int64,json=sInt64,proto3" json:"s_int64,omitempty"`
	SUint32              uint32   `protobuf:"varint,4,opt,name=s_uint32,json=sUint32,proto3" json:"s_uint32,omitempty"`
	SUint64              uint64   `protobuf:"varint,5,opt,name=s_uint64,json=sUint64,proto3" json:"s_uint64,omitempty"`
	SSint32              int32    `protobuf:"zigzag32,6,opt,name=s_sint32,json=sSint32,proto3" json:"s_sint32,omitempty"`
	SSint64              int64    `protobuf:"zigzag64,7,opt,name=s_sint64,json=sSint64,proto3" json:"s_sint64,omitempty"`
	SFixed32             uint32   `protobuf:"fixed32,8,opt,name=s_fixed32,json=sFixed32,proto3" json:"s_fixed32,omitempty"`
	SFixed64             uint64   `protobuf:"fixed64,9,opt,name=s_fixed64,json=sFixed64,proto3" json:"s_fixed64,omitempty"`
	SSfixed32            int32    `protobuf:"fixed32,10,opt,name=s_sfixed32,json=sSfixed32,proto3" json:"s_sfixed32,omitempty"`
	SSfixed64            int64    `protobuf:"fixed64,11,opt,name=s_sfixed64,json=sSfixed64,proto3" json:"s_sfixed64,omitempty"`
	SFloat               float32  `protobuf:"fixed32,20,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	SDouble              float64  `protobuf:"fixed64,21,opt,name=s_double,json=sDouble,proto3" json:"s_double,omitempty"`
	SBytes               []byte   `protobuf:"bytes,14,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	SString              string   `protobuf:"bytes,13,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalars) Reset()         { *m = Scalars{} }
func (m *Scalars) String() string { return proto.CompactTextString(m) }
func (*Scalars) ProtoMessage()    {}
func (*Scalars) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{0}
}

func (m *Scalars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalars.Unmarshal(m, b)
}
func (m *Scalars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalars.Marshal(b, m, deterministic)
}
func (m *Scalars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalars.Merge(m, src)
}
func (m *Scalars) XXX_Size() int {
	return xxx_messageInfo_Scalars.Size(m)
}
func (m *Scalars) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalars.DiscardUnknown(m)
}

var xxx_messageInfo_Scalars proto.InternalMessageInfo

func (m *Scalars) GetSBool() bool {
	if m != nil {
		return m.SBool
	}
	return false
}

func (m *Scalars) GetSInt32() int32 {
	if m != nil {
		return m.SInt32
	}
	return 0
}

func (m *Scalars) GetSInt64() int64 {
	if m != nil {
		return m.SInt64
	}
	return 0
}

func (m *Scalars) GetSUint32() uint32 {
	if m != nil {
		return m.SUint32
	}
	return 0
}

func (m *Scalars) GetSUint64() uint64 {
	if m != nil {
		return m.SUint64
	}
	return 0
}

func (m *Scalars) GetSSint32() int32 {
	if m != nil {
		return m.SSint32
	}
	return 0
}

func (m *Scalars) GetSSint64() int64 {
	if m != nil {
		return m.SSint64
	}
	return 0
}

func (m *Scalars) GetSFixed32() uint32 {
	if m != nil {
		return m.SFixed32
	}
	return 0
}

func (m *Scalars) GetSFixed64() uint64 {
	if m != nil {
		return m.SFixed64
	}
	return 0
}

func (m *Scalars) GetSSfixed32() int32 {
	if m != nil {
		return m.SSfixed32
	}
	return 0
}

func (m *Scalars) GetSSfixed64() int64 {
	if m != nil {
		return m.SSfixed64
	}
	return 0
}

func (m *Scalars) GetSFloat() float32 {
	if m != nil {
		return m.SFloat
	}
	return 0
}

func (m *Scalars) GetSDouble() float64 {
	if m != nil {
		return m.SDouble
	}
	return 0
}

func (m *Scalars) GetSBytes() []byte {
	if m != nil {
		return m.SBytes
	}
	return nil
}

func (m *Scalars) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

// Message contains enum fields.
type Enums struct {
	SEnum                Enum               `protobuf:"varint,1,opt,name=s_enum,json=sEnum,proto3,enum=pb3.Enum" json:"s_enum,omitempty"`
	RptEnum              []Enum             `protobuf:"varint,2,rep,packed,name=rpt_enum,json=rptEnum,proto3,enum=pb3.Enum" json:"rpt_enum,omitempty"`
	SNestedEnum          Enums_NestedEnum   `protobuf:"varint,3,opt,name=s_nested_enum,json=sNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"s_nested_enum,omitempty"`
	RptNestedEnum        []Enums_NestedEnum `protobuf:"varint,4,rep,packed,name=rpt_nested_enum,json=rptNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"rpt_nested_enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Enums) Reset()         { *m = Enums{} }
func (m *Enums) String() string { return proto.CompactTextString(m) }
func (*Enums) ProtoMessage()    {}
func (*Enums) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{1}
}

func (m *Enums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enums.Unmarshal(m, b)
}
func (m *Enums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enums.Marshal(b, m, deterministic)
}
func (m *Enums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enums.Merge(m, src)
}
func (m *Enums) XXX_Size() int {
	return xxx_messageInfo_Enums.Size(m)
}
func (m *Enums) XXX_DiscardUnknown() {
	xxx_messageInfo_Enums.DiscardUnknown(m)
}

var xxx_messageInfo_Enums proto.InternalMessageInfo

func (m *Enums) GetSEnum() Enum {
	if m != nil {
		return m.SEnum
	}
	return Enum_ZERO
}

func (m *Enums) GetRptEnum() []Enum {
	if m != nil {
		return m.RptEnum
	}
	return nil
}

func (m *Enums) GetSNestedEnum() Enums_NestedEnum {
	if m != nil {
		return m.SNestedEnum
	}
	return Enums_CERO
}

func (m *Enums) GetRptNestedEnum() []Enums_NestedEnum {
	if m != nil {
		return m.RptNestedEnum
	}
	return nil
}

// Message contains message and group fields.
type Nests struct {
	SNested              *Nested   `protobuf:"bytes,1,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	RptNested            []*Nested `protobuf:"bytes,2,rep,name=rpt_nested,json=rptNested,proto3" json:"rpt_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Nests) Reset()         { *m = Nests{} }
func (m *Nests) String() string { return proto.CompactTextString(m) }
func (*Nests) ProtoMessage()    {}
func (*Nests) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{2}
}

func (m *Nests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nests.Unmarshal(m, b)
}
func (m *Nests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nests.Marshal(b, m, deterministic)
}
func (m *Nests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nests.Merge(m, src)
}
func (m *Nests) XXX_Size() int {
	return xxx_messageInfo_Nests.Size(m)
}
func (m *Nests) XXX_DiscardUnknown() {
	xxx_messageInfo_Nests.DiscardUnknown(m)
}

var xxx_messageInfo_Nests proto.InternalMessageInfo

func (m *Nests) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

func (m *Nests) GetRptNested() []*Nested {
	if m != nil {
		return m.RptNested
	}
	return nil
}

// Message type used as submessage.
type Nested struct {
	SString              string   `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	SNested              *Nested  `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422, []int{3}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

func (m *Nested) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb3.Enum", Enum_name, Enum_value)
	proto.RegisterEnum("pb3.Enums_NestedEnum", Enums_NestedEnum_name, Enums_NestedEnum_value)
	proto.RegisterType((*Scalars)(nil), "pb3.Scalars")
	proto.RegisterType((*Enums)(nil), "pb3.Enums")
	proto.RegisterType((*Nests)(nil), "pb3.Nests")
	proto.RegisterType((*Nested)(nil), "pb3.Nested")
}

func init() {
	proto.RegisterFile("encoding/textpb/testprotos/pb3/test.proto", fileDescriptor_0854715c5b41c422)
}

var fileDescriptor_0854715c5b41c422 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0xfd, 0x29, 0x4e, 0xec, 0xe4, 0xe6, 0x97, 0xd6, 0x13, 0x2b, 0xf3, 0xfe, 0x81, 0x08, 0x63,
	0x68, 0x1d, 0xc4, 0x90, 0x18, 0xc3, 0x60, 0xdb, 0x43, 0xd7, 0x14, 0xca, 0x20, 0x01, 0x67, 0x65,
	0xd0, 0x3d, 0x84, 0x38, 0x71, 0xbd, 0x80, 0x6b, 0x19, 0x5f, 0x79, 0x74, 0x5f, 0x66, 0xdf, 0x74,
	0x30, 0x24, 0x39, 0x89, 0x3b, 0xe8, 0x9e, 0xac, 0xe3, 0x73, 0x8e, 0xee, 0x3d, 0xf7, 0x22, 0x78,
	0x93, 0xe4, 0x6b, 0xb1, 0xd9, 0xe6, 0xa9, 0x2f, 0x93, 0x3b, 0x59, 0xc4, 0xbe, 0x4c, 0x50, 0x16,
	0xa5, 0x90, 0x02, 0xfd, 0x22, 0x9e, 0x68, 0x38, 0xd2, 0x98, 0x5a, 0x45, 0x3c, 0x79, 0xf6, 0x22,
	0x15, 0x22, 0xcd, 0x12, 0x5f, 0xff, 0x8a, 0xab, 0x1b, 0x1f, 0x65, 0x59, 0xad, 0x6b, 0xc9, 0xf0,
	0x97, 0x05, 0xce, 0x62, 0xbd, 0xca, 0x56, 0x25, 0xd2, 0x13, 0xb0, 0x71, 0x19, 0x0b, 0x91, 0x79,
	0x84, 0x11, 0xde, 0x8d, 0x3a, 0x78, 0x26, 0x44, 0x46, 0x9f, 0x80, 0x83, 0xcb, 0x6d, 0x2e, 0x27,
	0x63, 0xaf, 0xc5, 0x08, 0xef, 0x44, 0x36, 0x5e, 0x2a, 0xb4, 0x27, 0xc2, 0xc0, 0xb3, 0x18, 0xe1,
	0x96, 0x21, 0xc2, 0x80, 0x3e, 0x85, 0x2e, 0x2e, 0x2b, 0x63, 0x69, 0x33, 0xc2, 0x07, 0x91, 0x83,
	0x57, 0x1a, 0x1e, 0xa8, 0x30, 0xf0, 0x3a, 0x8c, 0xf0, 0x76, 0x4d, 0xed, 0x5c, 0x68, 0x5c, 0x36,
	0x23, 0xfc, 0x51, 0xe4, 0xe0, 0xa2, 0xe1, 0x42, 0xe3, 0x72, 0x18, 0xe1, 0xb4, 0xa6, 0xc2, 0x80,
	0x3e, 0x87, 0x1e, 0x2e, 0x6f, 0xb6, 0x77, 0xc9, 0x66, 0x32, 0xf6, 0xba, 0x8c, 0x70, 0x27, 0xea,
	0xe2, 0x85, 0xc1, 0x0d, 0x32, 0x0c, 0xbc, 0x1e, 0x23, 0xdc, 0xde, 0x91, 0x61, 0x40, 0x5f, 0x02,
	0xe0, 0x12, 0x77, 0x56, 0x60, 0x84, 0x1f, 0x47, 0x3d, 0x5c, 0xd4, 0x3f, 0x9a, 0x74, 0x18, 0x78,
	0x7d, 0x46, 0xb8, 0xbb, 0xa7, 0xc3, 0xc0, 0x84, 0xbf, 0xc9, 0xc4, 0x4a, 0x7a, 0x8f, 0x19, 0xe1,
	0xad, 0xc8, 0xc6, 0x0b, 0x85, 0x4c, 0xaf, 0x1b, 0x51, 0xc5, 0x59, 0xe2, 0x9d, 0x30, 0xc2, 0x49,
	0xe4, 0xe0, 0xb9, 0x86, 0xc6, 0x13, 0xff, 0x94, 0x09, 0x7a, 0x47, 0x8c, 0xf0, 0xff, 0x23, 0x1b,
	0xcf, 0x14, 0xaa, 0xf3, 0xc9, 0x72, 0x9b, 0xa7, 0xde, 0x80, 0x11, 0xde, 0x53, 0xf9, 0x34, 0x1c,
	0xfe, 0x26, 0xd0, 0x99, 0xe6, 0xd5, 0x2d, 0x52, 0xa6, 0xd6, 0x93, 0xe4, 0xd5, 0xad, 0x5e, 0xcf,
	0xd1, 0xb8, 0x37, 0x2a, 0xe2, 0xc9, 0x48, 0x71, 0x51, 0x07, 0xd5, 0x87, 0xbe, 0x82, 0x6e, 0x59,
	0x48, 0xa3, 0x69, 0x31, 0xeb, 0xbe, 0xc6, 0x29, 0x0b, 0xa9, 0x55, 0xef, 0x60, 0x80, 0xcb, 0x3c,
	0x41, 0x99, 0x6c, 0x8c, 0xd4, 0xd2, 0xd7, 0x9d, 0xec, 0xa5, 0x38, 0x9a, 0x69, 0x56, 0xdb, 0xfa,
	0x78, 0x00, 0xf4, 0x03, 0x1c, 0xab, 0x02, 0x4d, 0x73, 0x5b, 0xd7, 0x79, 0xc0, 0x3c, 0x28, 0x0b,
	0x79, 0x80, 0xc3, 0x31, 0x40, 0xe3, 0xb2, 0x2e, 0xb4, 0x3f, 0x4d, 0xa3, 0xb9, 0xfb, 0x1f, 0x75,
	0xc0, 0xba, 0x9a, 0xcd, 0x5d, 0xa2, 0x0e, 0xe7, 0xf3, 0x85, 0xdb, 0x52, 0xdc, 0xf9, 0xe5, 0xf4,
	0xda, 0x85, 0xe1, 0x37, 0xe8, 0x28, 0x0f, 0xd2, 0xd7, 0x6a, 0x46, 0xa6, 0xb2, 0x1e, 0x40, 0x7f,
	0xdc, 0xd7, 0x45, 0xcd, 0x8d, 0x91, 0x53, 0xf7, 0x49, 0x4f, 0x01, 0x0e, 0x3d, 0xea, 0x31, 0xfc,
	0xa5, 0xec, 0xed, 0x9b, 0x1a, 0x7e, 0x06, 0xbb, 0x76, 0x35, 0x37, 0x40, 0xee, 0x6d, 0xe0, 0x5e,
	0xe1, 0xd6, 0xc3, 0x85, 0x4f, 0xdf, 0x42, 0x7b, 0x97, 0xeb, 0x7a, 0x9f, 0x6b, 0x3e, 0x9b, 0x9a,
	0x5c, 0x5f, 0xbe, 0xce, 0xdd, 0x96, 0x3e, 0x4c, 0x67, 0x2e, 0x9c, 0x7d, 0xbc, 0x7e, 0x9f, 0x6e,
	0xe5, 0xf7, 0x2a, 0x1e, 0xad, 0xc5, 0xad, 0x9f, 0x8a, 0x6c, 0x95, 0xa7, 0x87, 0x27, 0xfa, 0x63,
	0xec, 0xff, 0xfb, 0x95, 0xc7, 0xb6, 0x3e, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xef, 0xcf,
	0x95, 0x7a, 0x0e, 0x04, 0x00, 0x00,
}
