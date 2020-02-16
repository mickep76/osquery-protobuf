// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/osquery_registry.proto

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

// List the osquery registry plugins.
type OsqueryRegistry struct {
	// Name of the osquery registry
	Registry string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry"`
	// Name of the plugin item
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// Extension route UUID (0 for core)
	OwnerUuid int32 `protobuf:"varint,3,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid"`
	// 1 If the plugin is internal else 0
	Internal int32 `protobuf:"varint,4,opt,name=internal,proto3" json:"internal"`
	// 1 If this plugin is active else 0
	Active               int32    `protobuf:"varint,5,opt,name=active,proto3" json:"active"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OsqueryRegistry) Reset()         { *m = OsqueryRegistry{} }
func (m *OsqueryRegistry) String() string { return proto.CompactTextString(m) }
func (*OsqueryRegistry) ProtoMessage()    {}
func (*OsqueryRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_19c06398ec19cc33, []int{0}
}

func (m *OsqueryRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OsqueryRegistry.Unmarshal(m, b)
}
func (m *OsqueryRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OsqueryRegistry.Marshal(b, m, deterministic)
}
func (m *OsqueryRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsqueryRegistry.Merge(m, src)
}
func (m *OsqueryRegistry) XXX_Size() int {
	return xxx_messageInfo_OsqueryRegistry.Size(m)
}
func (m *OsqueryRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_OsqueryRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_OsqueryRegistry proto.InternalMessageInfo

func (m *OsqueryRegistry) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *OsqueryRegistry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OsqueryRegistry) GetOwnerUuid() int32 {
	if m != nil {
		return m.OwnerUuid
	}
	return 0
}

func (m *OsqueryRegistry) GetInternal() int32 {
	if m != nil {
		return m.Internal
	}
	return 0
}

func (m *OsqueryRegistry) GetActive() int32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func init() {
	proto.RegisterType((*OsqueryRegistry)(nil), "schemas.OsqueryRegistry")
}

func init() { proto.RegisterFile("pb/osquery_registry.proto", fileDescriptor_19c06398ec19cc33) }

var fileDescriptor_19c06398ec19cc33 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0xd2, 0xcf,
	0x2f, 0x2e, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0x9a,
	0xc2, 0xc8, 0xc5, 0xef, 0x0f, 0x51, 0x13, 0x04, 0x55, 0x22, 0x24, 0xc5, 0xc5, 0x01, 0x53, 0x2e,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x81, 0xc5, 0xc1, 0x6c, 0x21, 0x59, 0x2e, 0xae, 0xfc, 0xf2, 0xbc, 0xd4, 0xa2, 0xf8,
	0xd2, 0xd2, 0xcc, 0x14, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x4e, 0xb0, 0x48, 0x68, 0x69,
	0x66, 0x0a, 0xc8, 0xb8, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xbc, 0xc4, 0x1c, 0x09, 0x16, 0xb0, 0x24,
	0x9c, 0x2f, 0x24, 0xc6, 0xc5, 0x96, 0x98, 0x5c, 0x92, 0x59, 0x96, 0x2a, 0xc1, 0x0a, 0x96, 0x81,
	0xf2, 0x9c, 0x8c, 0xa2, 0x0c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x73, 0x33, 0x93, 0xb3, 0x53, 0x0b, 0xcc, 0xcd, 0x60, 0xbe, 0xd1, 0x05, 0x7b, 0x22, 0xa9, 0x34,
	0x4d, 0xbf, 0x20, 0x3b, 0xdd, 0x1a, 0xea, 0x95, 0x24, 0x36, 0xb0, 0xa8, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0xcd, 0x5d, 0xd7, 0xf7, 0x00, 0x00, 0x00,
}