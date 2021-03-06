// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/authorization_mechanism.proto

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

// OS X Authorization mechanisms database.
type AuthorizationMechanism struct {
	// Label of the authorization right
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label"`
	// Authorization plugin name
	Plugin string `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin"`
	// Name of the mechanism that will be called
	Mechanism string `protobuf:"bytes,3,opt,name=mechanism,proto3" json:"mechanism"`
	// If privileged it will run as root
	Privileged string `protobuf:"bytes,4,opt,name=privileged,proto3" json:"privileged"`
	// The whole string entry
	Entry                string   `protobuf:"bytes,5,opt,name=entry,proto3" json:"entry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationMechanism) Reset()         { *m = AuthorizationMechanism{} }
func (m *AuthorizationMechanism) String() string { return proto.CompactTextString(m) }
func (*AuthorizationMechanism) ProtoMessage()    {}
func (*AuthorizationMechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e2cfdd632cc1c6f, []int{0}
}

func (m *AuthorizationMechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationMechanism.Unmarshal(m, b)
}
func (m *AuthorizationMechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationMechanism.Marshal(b, m, deterministic)
}
func (m *AuthorizationMechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationMechanism.Merge(m, src)
}
func (m *AuthorizationMechanism) XXX_Size() int {
	return xxx_messageInfo_AuthorizationMechanism.Size(m)
}
func (m *AuthorizationMechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationMechanism.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationMechanism proto.InternalMessageInfo

func (m *AuthorizationMechanism) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AuthorizationMechanism) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *AuthorizationMechanism) GetMechanism() string {
	if m != nil {
		return m.Mechanism
	}
	return ""
}

func (m *AuthorizationMechanism) GetPrivileged() string {
	if m != nil {
		return m.Privileged
	}
	return ""
}

func (m *AuthorizationMechanism) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthorizationMechanism)(nil), "schemas.AuthorizationMechanism")
}

func init() { proto.RegisterFile("pb/authorization_mechanism.proto", fileDescriptor_2e2cfdd632cc1c6f) }

var fileDescriptor_2e2cfdd632cc1c6f = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x18, 0x84, 0x59, 0xf5, 0x4e, 0x2e, 0x65, 0x90, 0x23, 0x85, 0xc8, 0x61, 0x65, 0xe3, 0x46, 0x14,
	0xb4, 0xb0, 0xd2, 0xde, 0xc6, 0xd2, 0x46, 0x92, 0xf8, 0x9a, 0xbc, 0x6c, 0xbe, 0xcc, 0x87, 0xb0,
	0xfe, 0x14, 0x7f, 0xad, 0x98, 0x5d, 0xbd, 0x2d, 0xe7, 0x99, 0x81, 0x87, 0x21, 0xbb, 0x28, 0xb9,
	0xa8, 0xc5, 0x84, 0x84, 0x5f, 0xa2, 0x60, 0xf0, 0xaf, 0x0e, 0x94, 0x11, 0x1e, 0xb3, 0xeb, 0x63,
	0x0a, 0x25, 0xd0, 0xe3, 0xac, 0x0c, 0x38, 0x91, 0xcf, 0xbf, 0x3b, 0xb2, 0x7d, 0x58, 0x4e, 0x9f,
	0xfe, 0x96, 0xf4, 0x84, 0xac, 0xac, 0x90, 0x60, 0x59, 0xb7, 0xeb, 0x2e, 0x36, 0xcf, 0x53, 0xa0,
	0x5b, 0xb2, 0x8e, 0xb6, 0x6a, 0xf4, 0xec, 0xa0, 0xe1, 0x39, 0xd1, 0x53, 0xb2, 0xf9, 0x97, 0xb0,
	0xc3, 0x56, 0xed, 0x01, 0x3d, 0x23, 0x24, 0x26, 0xfc, 0x44, 0x0b, 0x1a, 0xde, 0xd8, 0x51, 0xab,
	0x17, 0xe4, 0xd7, 0x05, 0xbe, 0xa4, 0x91, 0xad, 0x26, 0x57, 0x0b, 0x8f, 0xd7, 0x2f, 0x57, 0x1a,
	0x8b, 0xa9, 0xb2, 0x57, 0xc1, 0x71, 0x87, 0x6a, 0x80, 0x78, 0x77, 0xcb, 0x43, 0xfe, 0xa8, 0x90,
	0xc6, 0xcb, 0x76, 0x45, 0xd6, 0x77, 0x1e, 0x07, 0x7d, 0x3f, 0x1f, 0x92, 0xeb, 0x46, 0x6f, 0x7e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xf0, 0xb7, 0x79, 0x04, 0x01, 0x00, 0x00,
}
