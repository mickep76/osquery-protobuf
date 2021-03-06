// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/disk_encryption.proto

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

// Disk encryption status and information.
type DiskEncryption struct {
	// Disk name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Disk Universally Unique Identifier
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid"`
	// 1 If encrypted: true (disk is encrypted)
	Encrypted int32 `protobuf:"varint,3,opt,name=encrypted,proto3" json:"encrypted"`
	// Description of cipher type and mode if available
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	// Currently authenticated user if available (Apple)
	Uid string `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid"`
	// UUID of authenticated user if available (Apple)
	UserUuid string `protobuf:"bytes,6,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid"`
	// Disk encryption status with one of following values: encrypted | not encrypted | undefined
	EncryptionStatus     string   `protobuf:"bytes,7,opt,name=encryption_status,json=encryptionStatus,proto3" json:"encryption_status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiskEncryption) Reset()         { *m = DiskEncryption{} }
func (m *DiskEncryption) String() string { return proto.CompactTextString(m) }
func (*DiskEncryption) ProtoMessage()    {}
func (*DiskEncryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_8513803356d784c8, []int{0}
}

func (m *DiskEncryption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskEncryption.Unmarshal(m, b)
}
func (m *DiskEncryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskEncryption.Marshal(b, m, deterministic)
}
func (m *DiskEncryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskEncryption.Merge(m, src)
}
func (m *DiskEncryption) XXX_Size() int {
	return xxx_messageInfo_DiskEncryption.Size(m)
}
func (m *DiskEncryption) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskEncryption.DiscardUnknown(m)
}

var xxx_messageInfo_DiskEncryption proto.InternalMessageInfo

func (m *DiskEncryption) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DiskEncryption) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DiskEncryption) GetEncrypted() int32 {
	if m != nil {
		return m.Encrypted
	}
	return 0
}

func (m *DiskEncryption) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DiskEncryption) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DiskEncryption) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *DiskEncryption) GetEncryptionStatus() string {
	if m != nil {
		return m.EncryptionStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*DiskEncryption)(nil), "schemas.DiskEncryption")
}

func init() { proto.RegisterFile("pb/disk_encryption.proto", fileDescriptor_8513803356d784c8) }

var fileDescriptor_8513803356d784c8 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0xfb, 0xcf, 0xe6, 0x20, 0x6b, 0x4e, 0x01, 0x3d, 0x2c, 0x9e, 0x16, 0xc4, 0xad,
	0x28, 0xe8, 0xc1, 0x9b, 0xe8, 0x0b, 0xac, 0x78, 0xf1, 0x52, 0xda, 0x74, 0xdc, 0x0d, 0x21, 0x4d,
	0xcc, 0x24, 0x87, 0xbe, 0xa0, 0xcf, 0x25, 0x33, 0x5b, 0xec, 0xed, 0xc7, 0x6f, 0x66, 0x3e, 0xe6,
	0x13, 0x2a, 0xb4, 0x55, 0x67, 0xd0, 0xd6, 0xd0, 0xeb, 0x38, 0x84, 0x64, 0x7c, 0xbf, 0x0b, 0xd1,
	0x27, 0x2f, 0x57, 0xa8, 0x8f, 0xe0, 0x1a, 0xbc, 0xf9, 0x2d, 0xc4, 0xc5, 0x9b, 0x41, 0xfb, 0xfe,
	0xbf, 0x21, 0xa5, 0x98, 0xf7, 0x8d, 0x03, 0x55, 0x6c, 0x8a, 0x6d, 0xb9, 0x67, 0x26, 0x97, 0xb3,
	0xe9, 0xd4, 0xd9, 0xc9, 0x11, 0xcb, 0x6b, 0x51, 0x8e, 0xb9, 0xd0, 0xa9, 0xd9, 0xa6, 0xd8, 0x2e,
	0xf6, 0x93, 0xa0, 0x8b, 0x34, 0x04, 0x50, 0xf3, 0xd3, 0x05, 0xb1, 0x5c, 0x8b, 0x19, 0x85, 0x2c,
	0x58, 0x11, 0xca, 0x2b, 0x51, 0x66, 0x84, 0x58, 0x73, 0xf8, 0x92, 0xfd, 0x39, 0x89, 0x4f, 0x1a,
	0xde, 0x8a, 0xcb, 0xe9, 0xf1, 0x1a, 0x53, 0x93, 0x32, 0xaa, 0x15, 0x2f, 0xad, 0xa7, 0xc1, 0x07,
	0xfb, 0xd7, 0x87, 0xaf, 0xfb, 0x83, 0x49, 0xc7, 0xdc, 0xee, 0xb4, 0x77, 0x95, 0x33, 0xda, 0x42,
	0x78, 0x7e, 0xaa, 0x3c, 0xfe, 0x64, 0x88, 0xc3, 0x1d, 0xd7, 0x6e, 0xf3, 0x77, 0x15, 0xec, 0xe1,
	0x65, 0x2c, 0xdf, 0x2e, 0xd9, 0x3e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x16, 0x16, 0x25,
	0x28, 0x01, 0x00, 0x00,
}
