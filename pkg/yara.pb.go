// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/yara.proto

package schemas

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

// Track YARA matches for files or PIDs.
type Yara struct {
	// List of YARA matches
	Matches string `protobuf:"bytes,1,opt,name=matches,proto3" json:"matches"`
	// Number of YARA matches
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
	// Matching strings
	Strings string `protobuf:"bytes,3,opt,name=strings,proto3" json:"strings"`
	// Matching tags
	Tags                 string   `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Yara) Reset()         { *m = Yara{} }
func (m *Yara) String() string { return proto.CompactTextString(m) }
func (*Yara) ProtoMessage()    {}
func (*Yara) Descriptor() ([]byte, []int) {
	return fileDescriptor_525dc6cd5d3d5c82, []int{0}
}

func (m *Yara) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Yara.Unmarshal(m, b)
}
func (m *Yara) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Yara.Marshal(b, m, deterministic)
}
func (m *Yara) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Yara.Merge(m, src)
}
func (m *Yara) XXX_Size() int {
	return xxx_messageInfo_Yara.Size(m)
}
func (m *Yara) XXX_DiscardUnknown() {
	xxx_messageInfo_Yara.DiscardUnknown(m)
}

var xxx_messageInfo_Yara proto.InternalMessageInfo

func (m *Yara) GetMatches() string {
	if m != nil {
		return m.Matches
	}
	return ""
}

func (m *Yara) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Yara) GetStrings() string {
	if m != nil {
		return m.Strings
	}
	return ""
}

func (m *Yara) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func init() {
	proto.RegisterType((*Yara)(nil), "schemas.Yara")
}

func init() { proto.RegisterFile("pb/yara.proto", fileDescriptor_525dc6cd5d3d5c82) }

var fileDescriptor_525dc6cd5d3d5c82 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0xae, 0x82, 0x30,
	0x18, 0x46, 0xc3, 0xbd, 0x20, 0xb1, 0x89, 0x4b, 0xe3, 0xd0, 0x91, 0x38, 0xb1, 0x48, 0x8d, 0x26,
	0x3a, 0xb8, 0xf9, 0x08, 0x6c, 0xba, 0xfd, 0xad, 0xb5, 0x10, 0x52, 0x5a, 0xfb, 0xb7, 0x03, 0x6f,
	0x6f, 0x2c, 0xb8, 0x7d, 0xe7, 0xcb, 0x19, 0x0e, 0xd9, 0x38, 0xc1, 0x27, 0xf0, 0xd0, 0x38, 0x6f,
	0x83, 0xa5, 0x25, 0xca, 0x4e, 0x19, 0xc0, 0xdd, 0x93, 0xe4, 0x77, 0xf0, 0x40, 0x19, 0x29, 0x0d,
	0x04, 0xd9, 0x29, 0x64, 0x59, 0x95, 0xd5, 0xeb, 0xf6, 0x87, 0x74, 0x4b, 0x0a, 0x69, 0xe3, 0x18,
	0xd8, 0x5f, 0x95, 0xd5, 0x45, 0x3b, 0xc3, 0xd7, 0xc7, 0xe0, 0xfb, 0x51, 0x23, 0xfb, 0x9f, 0xfd,
	0x05, 0x29, 0x25, 0x79, 0x00, 0x8d, 0x2c, 0x4f, 0x77, 0xda, 0xb7, 0xe3, 0xe3, 0xa0, 0xfb, 0xd0,
	0x45, 0xd1, 0x48, 0x6b, 0xb8, 0xe9, 0xe5, 0xa0, 0xdc, 0xe5, 0xcc, 0x2d, 0xbe, 0xa3, 0xf2, 0xd3,
	0x3e, 0x35, 0x89, 0xf8, 0xe2, 0x6e, 0xd0, 0xd7, 0xa5, 0x4c, 0xac, 0xd2, 0x7b, 0xfa, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x12, 0x1d, 0x6c, 0xba, 0x00, 0x00, 0x00,
}
