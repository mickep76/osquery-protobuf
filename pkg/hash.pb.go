// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/hash.proto

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

// Filesystem hash data.
type Hash struct {
	// Must provide a path or directory
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// Must provide a path or directory
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory"`
	// MD5 hash of provided filesystem data
	Md5 string `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5"`
	// SHA1 hash of provided filesystem data
	Sha1 string `protobuf:"bytes,4,opt,name=sha1,proto3" json:"sha1"`
	// SHA256 hash of provided filesystem data
	Sha256 string `protobuf:"bytes,5,opt,name=sha256,proto3" json:"sha256"`
	// ssdeep hash of provided filesystem data
	Ssdeep               string   `protobuf:"bytes,6,opt,name=ssdeep,proto3" json:"ssdeep"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ecf4a1d21b99ba7, []int{0}
}

func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Hash) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *Hash) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *Hash) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *Hash) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Hash) GetSsdeep() string {
	if m != nil {
		return m.Ssdeep
	}
	return ""
}

func init() {
	proto.RegisterType((*Hash)(nil), "schemas.Hash")
}

func init() { proto.RegisterFile("pb/hash.proto", fileDescriptor_3ecf4a1d21b99ba7) }

var fileDescriptor_3ecf4a1d21b99ba7 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0xcf,
	0x48, 0x2c, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0x2c, 0x56, 0xea, 0x63, 0xe4, 0x62, 0xf1, 0x48, 0x2c, 0xce, 0x10, 0x12, 0xe2, 0x62, 0x29,
	0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x64, 0xb8, 0x38,
	0x53, 0x32, 0x8b, 0x52, 0x93, 0x4b, 0xf2, 0x8b, 0x2a, 0x25, 0x98, 0xc0, 0x12, 0x08, 0x01, 0x21,
	0x01, 0x2e, 0xe6, 0xdc, 0x14, 0x53, 0x09, 0x66, 0xb0, 0x38, 0x88, 0x09, 0x32, 0xa3, 0x38, 0x23,
	0xd1, 0x50, 0x82, 0x05, 0x62, 0x06, 0x88, 0x2d, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x91, 0x68, 0x64,
	0x6a, 0x26, 0xc1, 0x0a, 0x16, 0x85, 0xf2, 0xc0, 0xe2, 0xc5, 0x29, 0xa9, 0xa9, 0x05, 0x12, 0x6c,
	0x50, 0x71, 0x30, 0xcf, 0xc9, 0x28, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x3f, 0x37, 0x33, 0x39, 0x3b, 0xb5, 0xc0, 0xdc, 0x4c, 0x3f, 0xbf, 0xb8, 0xb0, 0x34,
	0xb5, 0xa8, 0x52, 0x17, 0xec, 0xfc, 0xa4, 0xd2, 0x34, 0xfd, 0x82, 0xec, 0x74, 0x6b, 0xa8, 0x27,
	0x92, 0xd8, 0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xd7, 0xa2, 0xfd, 0xe5,
	0x00, 0x00, 0x00,
}
