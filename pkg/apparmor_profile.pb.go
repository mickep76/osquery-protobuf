// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/apparmor_profile.proto

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

// Track active AppArmor profiles.
type ApparmorProfile struct {
	// Unique
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// Policy name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// Which executable(s) a profile will attach to.
	Attach string `protobuf:"bytes,3,opt,name=attach,proto3" json:"attach"`
	// How the policy is applied.
	Mode string `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode"`
	// A unique hash that identifies this policy.
	Sha1                 string   `protobuf:"bytes,5,opt,name=sha1,proto3" json:"sha1"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApparmorProfile) Reset()         { *m = ApparmorProfile{} }
func (m *ApparmorProfile) String() string { return proto.CompactTextString(m) }
func (*ApparmorProfile) ProtoMessage()    {}
func (*ApparmorProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_90fbff8a0d3470b0, []int{0}
}

func (m *ApparmorProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApparmorProfile.Unmarshal(m, b)
}
func (m *ApparmorProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApparmorProfile.Marshal(b, m, deterministic)
}
func (m *ApparmorProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApparmorProfile.Merge(m, src)
}
func (m *ApparmorProfile) XXX_Size() int {
	return xxx_messageInfo_ApparmorProfile.Size(m)
}
func (m *ApparmorProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ApparmorProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ApparmorProfile proto.InternalMessageInfo

func (m *ApparmorProfile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ApparmorProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApparmorProfile) GetAttach() string {
	if m != nil {
		return m.Attach
	}
	return ""
}

func (m *ApparmorProfile) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ApparmorProfile) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func init() {
	proto.RegisterType((*ApparmorProfile)(nil), "schemas.ApparmorProfile")
}

func init() { proto.RegisterFile("pb/apparmor_profile.proto", fileDescriptor_90fbff8a0d3470b0) }

var fileDescriptor_90fbff8a0d3470b0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xbd, 0x6e, 0xc2, 0x30,
	0x14, 0x46, 0x95, 0x36, 0x4d, 0x55, 0x2f, 0x95, 0x3c, 0x54, 0xee, 0x86, 0x98, 0x58, 0x88, 0xf9,
	0x91, 0x60, 0x60, 0x82, 0x27, 0x40, 0x8c, 0x2c, 0xe8, 0xda, 0x38, 0x71, 0x14, 0x8c, 0x2f, 0xb6,
	0x33, 0xe4, 0xed, 0x91, 0xed, 0x6c, 0xe7, 0x3b, 0x3a, 0xc3, 0x47, 0xfe, 0x51, 0x70, 0x40, 0x04,
	0x67, 0xac, 0xbb, 0xa1, 0xb3, 0x4d, 0xf7, 0x50, 0x35, 0x3a, 0x1b, 0x2c, 0xfd, 0xf6, 0x52, 0x2b,
	0x03, 0x7e, 0x3e, 0x92, 0xdf, 0xe3, 0x94, 0x9c, 0x73, 0x41, 0x29, 0x29, 0x11, 0x82, 0x66, 0xc5,
	0xac, 0x58, 0xfc, 0x5c, 0x12, 0x47, 0xf7, 0x04, 0xa3, 0xd8, 0x47, 0x76, 0x91, 0xe9, 0x1f, 0xa9,
	0x20, 0x04, 0x90, 0x9a, 0x7d, 0x26, 0x3b, 0xad, 0xd8, 0x1a, 0x7b, 0x57, 0xac, 0xcc, 0x6d, 0xe4,
	0xe8, 0xbc, 0x86, 0x35, 0xfb, 0xca, 0x2e, 0xf2, 0x69, 0x73, 0x5d, 0xb5, 0x5d, 0xd0, 0x83, 0xa8,
	0xa5, 0x35, 0xdc, 0x74, 0xb2, 0x57, 0xb8, 0xdf, 0x71, 0xeb, 0x5f, 0x83, 0x72, 0xe3, 0x32, 0x1d,
	0x15, 0x43, 0xc3, 0xb1, 0x6f, 0x0f, 0xd3, 0x5d, 0x51, 0x25, 0xbb, 0x7d, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0x26, 0x53, 0xb1, 0xdb, 0x00, 0x00, 0x00,
}