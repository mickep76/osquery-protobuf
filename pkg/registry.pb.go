// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/registry.proto

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

// All of the Windows registry hives.
type Registry struct {
	// Name of the key to search for
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// Full path to the value
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path"`
	// Name of the registry value entry
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// Type of the registry value
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	// Data content of registry value
	Data string `protobuf:"bytes,5,opt,name=data,proto3" json:"data"`
	// timestamp of the most recent registry write
	Mtime                int64    `protobuf:"varint,6,opt,name=mtime,proto3" json:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a42d6bc5de2ab13, []int{0}
}

func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Registry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Registry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registry) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Registry) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Registry) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func init() {
	proto.RegisterType((*Registry)(nil), "schemas.Registry")
}

func init() { proto.RegisterFile("pb/registry.proto", fileDescriptor_4a42d6bc5de2ab13) }

var fileDescriptor_4a42d6bc5de2ab13 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xce, 0x3f, 0x8f, 0xc2, 0x20,
	0x18, 0xc7, 0xf1, 0x70, 0xfd, 0x73, 0x77, 0x4c, 0x77, 0xc4, 0x81, 0xb1, 0x71, 0xea, 0x62, 0x31,
	0x9a, 0xe8, 0xe0, 0xe6, 0x4b, 0xe8, 0xe8, 0x06, 0xf5, 0xb1, 0x25, 0x0d, 0x82, 0x40, 0x07, 0x56,
	0x5f, 0xb9, 0x01, 0xba, 0x7d, 0x9f, 0xcf, 0xb3, 0xfc, 0xf0, 0xbf, 0x11, 0xcc, 0xc2, 0x28, 0x9d,
	0xb7, 0xa1, 0x33, 0x56, 0x7b, 0x4d, 0xbe, 0xdd, 0x30, 0x81, 0xe2, 0x6e, 0xfb, 0x46, 0xf8, 0xa7,
	0x5f, 0x7f, 0xe4, 0x0f, 0x17, 0x33, 0x04, 0x8a, 0x1a, 0xd4, 0xfe, 0xf6, 0x31, 0x09, 0xc1, 0xa5,
	0xe1, 0x7e, 0xa2, 0x5f, 0x89, 0x52, 0x47, 0x7b, 0x72, 0x05, 0xb4, 0xc8, 0x16, 0x3b, 0x9a, 0x0f,
	0x06, 0x68, 0x99, 0x2d, 0x76, 0xb4, 0x3b, 0xf7, 0x9c, 0x56, 0xd9, 0x62, 0x93, 0x0d, 0xae, 0x94,
	0x97, 0x0a, 0x68, 0xdd, 0xa0, 0xb6, 0xe8, 0xf3, 0x71, 0x3d, 0xdc, 0xf6, 0xa3, 0xf4, 0xd3, 0x22,
	0xba, 0x41, 0x2b, 0xa6, 0xe4, 0x30, 0x83, 0x39, 0x9f, 0x98, 0x76, 0xaf, 0x05, 0x6c, 0xd8, 0xa5,
	0xc9, 0x62, 0x79, 0x30, 0x33, 0x8f, 0x97, 0x75, 0xb8, 0xa8, 0x93, 0x1e, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x18, 0x78, 0x9d, 0x52, 0xdd, 0x00, 0x00, 0x00,
}