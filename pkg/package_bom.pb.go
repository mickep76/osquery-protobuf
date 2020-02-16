// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/package_bom.proto

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

// OS X package bill of materials (BOM) file list.
type PackageBom struct {
	// Package file or directory
	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath"`
	// Expected user of file or directory
	Uid int32 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid"`
	// Expected group of file or directory
	Gid int32 `protobuf:"varint,3,opt,name=gid,proto3" json:"gid"`
	// Expected permissions
	Mode int32 `protobuf:"varint,4,opt,name=mode,proto3" json:"mode"`
	// Expected file size
	Size int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size"`
	// Timestamp the file was installed
	ModifiedTime int32 `protobuf:"varint,6,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time"`
	// Path of package bom
	Path                 string   `protobuf:"bytes,7,opt,name=path,proto3" json:"path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageBom) Reset()         { *m = PackageBom{} }
func (m *PackageBom) String() string { return proto.CompactTextString(m) }
func (*PackageBom) ProtoMessage()    {}
func (*PackageBom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e2ba1e64643d682, []int{0}
}

func (m *PackageBom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageBom.Unmarshal(m, b)
}
func (m *PackageBom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageBom.Marshal(b, m, deterministic)
}
func (m *PackageBom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageBom.Merge(m, src)
}
func (m *PackageBom) XXX_Size() int {
	return xxx_messageInfo_PackageBom.Size(m)
}
func (m *PackageBom) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageBom.DiscardUnknown(m)
}

var xxx_messageInfo_PackageBom proto.InternalMessageInfo

func (m *PackageBom) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

func (m *PackageBom) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PackageBom) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *PackageBom) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *PackageBom) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PackageBom) GetModifiedTime() int32 {
	if m != nil {
		return m.ModifiedTime
	}
	return 0
}

func (m *PackageBom) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*PackageBom)(nil), "schemas.PackageBom")
}

func init() { proto.RegisterFile("pb/package_bom.proto", fileDescriptor_1e2ba1e64643d682) }

var fileDescriptor_1e2ba1e64643d682 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0x59, 0x73, 0x7f, 0x74, 0x51, 0x90, 0xc5, 0x62, 0xb1, 0x0a, 0xda, 0xa4, 0xf1, 0x22,
	0x0a, 0x5a, 0xd8, 0xdd, 0x13, 0x48, 0xb0, 0xb2, 0x39, 0xb2, 0xd9, 0xc9, 0x66, 0x88, 0xc3, 0xac,
	0x97, 0xdd, 0x42, 0xdf, 0xc9, 0x77, 0x94, 0xcc, 0xe9, 0x75, 0xbf, 0xf9, 0x0d, 0x1f, 0x7c, 0x9f,
	0xbe, 0x8a, 0xae, 0x8e, 0x6d, 0x37, 0xb6, 0x01, 0x76, 0x8e, 0x69, 0x13, 0xf7, 0x9c, 0xd8, 0xac,
	0xa7, 0x6e, 0x00, 0x6a, 0xa7, 0x9b, 0x1f, 0xa5, 0xf5, 0xeb, 0xe1, 0xbd, 0x65, 0x32, 0xd7, 0xfa,
	0xb4, 0xc7, 0x0f, 0x88, 0x6d, 0x1a, 0xac, 0x2a, 0x55, 0x75, 0xd6, 0x1c, 0x6f, 0x73, 0xa9, 0x8b,
	0x8c, 0xde, 0x9e, 0x94, 0xaa, 0x5a, 0x36, 0x33, 0xce, 0x26, 0xa0, 0xb7, 0xc5, 0xc1, 0x04, 0xf4,
	0xc6, 0xe8, 0x05, 0xb1, 0x07, 0xbb, 0x10, 0x25, 0x3c, 0xbb, 0x09, 0xbf, 0xc1, 0x2e, 0x4b, 0x55,
	0x15, 0x8d, 0xb0, 0xb9, 0xd5, 0x17, 0xc4, 0x1e, 0x7b, 0x04, 0xbf, 0x4b, 0x48, 0x60, 0x57, 0x12,
	0x38, 0xff, 0x97, 0x6f, 0x48, 0x12, 0x94, 0x22, 0x6b, 0x29, 0x22, 0xbc, 0x7d, 0x78, 0xbf, 0x0f,
	0x98, 0x86, 0xec, 0x36, 0x1d, 0x53, 0x4d, 0xd8, 0x8d, 0x10, 0x9f, 0x9f, 0x6a, 0x9e, 0x3e, 0x33,
	0xec, 0xbf, 0xee, 0x64, 0x9d, 0xcb, 0x7d, 0x1d, 0xc7, 0xf0, 0xf2, 0xb7, 0xd1, 0xad, 0xc4, 0x3e,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x09, 0xf8, 0xc2, 0x0b, 0x01, 0x00, 0x00,
}