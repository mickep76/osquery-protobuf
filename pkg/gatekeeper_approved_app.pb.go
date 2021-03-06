// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/gatekeeper_approved_app.proto

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

// Gatekeeper apps a user has allowed to run.
type GatekeeperApprovedApp struct {
	// Path of executable allowed to run
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// Code signing requirement language
	Requirement string `protobuf:"bytes,2,opt,name=requirement,proto3" json:"requirement"`
	// Last change time
	Ctime float64 `protobuf:"fixed64,3,opt,name=ctime,proto3" json:"ctime"`
	// Last modification time
	Mtime                float64  `protobuf:"fixed64,4,opt,name=mtime,proto3" json:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatekeeperApprovedApp) Reset()         { *m = GatekeeperApprovedApp{} }
func (m *GatekeeperApprovedApp) String() string { return proto.CompactTextString(m) }
func (*GatekeeperApprovedApp) ProtoMessage()    {}
func (*GatekeeperApprovedApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9f819fda1d40fcb, []int{0}
}

func (m *GatekeeperApprovedApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatekeeperApprovedApp.Unmarshal(m, b)
}
func (m *GatekeeperApprovedApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatekeeperApprovedApp.Marshal(b, m, deterministic)
}
func (m *GatekeeperApprovedApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatekeeperApprovedApp.Merge(m, src)
}
func (m *GatekeeperApprovedApp) XXX_Size() int {
	return xxx_messageInfo_GatekeeperApprovedApp.Size(m)
}
func (m *GatekeeperApprovedApp) XXX_DiscardUnknown() {
	xxx_messageInfo_GatekeeperApprovedApp.DiscardUnknown(m)
}

var xxx_messageInfo_GatekeeperApprovedApp proto.InternalMessageInfo

func (m *GatekeeperApprovedApp) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GatekeeperApprovedApp) GetRequirement() string {
	if m != nil {
		return m.Requirement
	}
	return ""
}

func (m *GatekeeperApprovedApp) GetCtime() float64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *GatekeeperApprovedApp) GetMtime() float64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func init() {
	proto.RegisterType((*GatekeeperApprovedApp)(nil), "schemas.GatekeeperApprovedApp")
}

func init() { proto.RegisterFile("pb/gatekeeper_approved_app.proto", fileDescriptor_b9f819fda1d40fcb) }

var fileDescriptor_b9f819fda1d40fcb = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x05, 0x60, 0xd4, 0xba, 0x2d, 0x55, 0x37, 0xd1, 0x82, 0x47, 0xd1, 0xc9, 0x4b, 0xad, 0x92,
	0x40, 0x32, 0x64, 0x72, 0x96, 0xec, 0x1e, 0xb3, 0x04, 0x49, 0xb9, 0xd8, 0xc2, 0x28, 0x3a, 0xcb,
	0x52, 0xc0, 0xff, 0x3e, 0x58, 0x36, 0x21, 0xdb, 0xbd, 0xef, 0xde, 0xf0, 0x28, 0x47, 0x25, 0x1a,
	0x19, 0xa0, 0x03, 0x40, 0xf0, 0x27, 0x89, 0xe8, 0xdd, 0x0d, 0xce, 0xd3, 0x51, 0xa2, 0x77, 0xc1,
	0xb1, 0x8f, 0x41, 0xb7, 0x60, 0xe5, 0xf0, 0x3b, 0xd2, 0x9f, 0xc3, 0xa3, 0x59, 0x2d, 0xc5, 0x0a,
	0x91, 0x31, 0x9a, 0xa1, 0x0c, 0x6d, 0x4e, 0x38, 0x29, 0x3e, 0xeb, 0x74, 0x33, 0x4e, 0xbf, 0x3c,
	0xf4, 0xd1, 0x78, 0xb0, 0x70, 0x0d, 0xf9, 0x4b, 0x7a, 0x3d, 0x13, 0xfb, 0xa6, 0x6f, 0x3a, 0x18,
	0x0b, 0xf9, 0x2b, 0x27, 0x05, 0xa9, 0xe7, 0x30, 0xa9, 0x4d, 0x9a, 0xcd, 0x9a, 0xc2, 0x7e, 0x75,
	0xfc, 0x6f, 0x4c, 0x68, 0xa3, 0x2a, 0xb5, 0xb3, 0xc2, 0x1a, 0xdd, 0x01, 0x6e, 0x37, 0xc2, 0x0d,
	0x7d, 0x04, 0x3f, 0xfe, 0xa5, 0xa1, 0x2a, 0x5e, 0x04, 0x76, 0xcd, 0x6e, 0x99, 0xab, 0xde, 0x93,
	0xae, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xcb, 0x18, 0x6f, 0xe2, 0x00, 0x00, 0x00,
}
