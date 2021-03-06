// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/windows_security_product.proto

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

// Enumeration of registered Windows security products.
type WindowsSecurityProduct struct {
	// Type of security product
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// Name of product
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// State of protection
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state"`
	// Timestamp for the product state
	StateTimestamp string `protobuf:"bytes,4,opt,name=state_timestamp,json=stateTimestamp,proto3" json:"state_timestamp"`
	// Remediation path
	RemediationPath string `protobuf:"bytes,5,opt,name=remediation_path,json=remediationPath,proto3" json:"remediation_path"`
	// 1 if product signatures are up to date
	SignaturesUpToDate   int32    `protobuf:"varint,6,opt,name=signatures_up_to_date,json=signaturesUpToDate,proto3" json:"signatures_up_to_date"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WindowsSecurityProduct) Reset()         { *m = WindowsSecurityProduct{} }
func (m *WindowsSecurityProduct) String() string { return proto.CompactTextString(m) }
func (*WindowsSecurityProduct) ProtoMessage()    {}
func (*WindowsSecurityProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_5099a9284bda5cd7, []int{0}
}

func (m *WindowsSecurityProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowsSecurityProduct.Unmarshal(m, b)
}
func (m *WindowsSecurityProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowsSecurityProduct.Marshal(b, m, deterministic)
}
func (m *WindowsSecurityProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowsSecurityProduct.Merge(m, src)
}
func (m *WindowsSecurityProduct) XXX_Size() int {
	return xxx_messageInfo_WindowsSecurityProduct.Size(m)
}
func (m *WindowsSecurityProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowsSecurityProduct.DiscardUnknown(m)
}

var xxx_messageInfo_WindowsSecurityProduct proto.InternalMessageInfo

func (m *WindowsSecurityProduct) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *WindowsSecurityProduct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WindowsSecurityProduct) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *WindowsSecurityProduct) GetStateTimestamp() string {
	if m != nil {
		return m.StateTimestamp
	}
	return ""
}

func (m *WindowsSecurityProduct) GetRemediationPath() string {
	if m != nil {
		return m.RemediationPath
	}
	return ""
}

func (m *WindowsSecurityProduct) GetSignaturesUpToDate() int32 {
	if m != nil {
		return m.SignaturesUpToDate
	}
	return 0
}

func init() {
	proto.RegisterType((*WindowsSecurityProduct)(nil), "schemas.WindowsSecurityProduct")
}

func init() { proto.RegisterFile("pb/windows_security_product.proto", fileDescriptor_5099a9284bda5cd7) }

var fileDescriptor_5099a9284bda5cd7 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xa9, 0x4e, 0x47, 0xcc, 0xc2, 0x91, 0xa0, 0x92, 0xe5, 0xe8, 0xc6, 0x71, 0xe1, 0xd4,
	0x1f, 0xd0, 0x85, 0x3b, 0xf1, 0x01, 0x86, 0x71, 0x44, 0x70, 0x13, 0xd2, 0xf4, 0xda, 0x86, 0x21,
	0xcd, 0x35, 0xb9, 0x61, 0xe8, 0xc3, 0xfa, 0x2e, 0xd2, 0xb4, 0xa2, 0xbb, 0x73, 0xbf, 0xf3, 0xc1,
	0x85, 0xc3, 0xce, 0xb1, 0x2c, 0x76, 0xa6, 0xad, 0xdc, 0x2e, 0xc8, 0x00, 0x3a, 0x7a, 0x43, 0x9d,
	0x44, 0xef, 0xaa, 0xa8, 0x69, 0x89, 0xde, 0x91, 0xe3, 0x07, 0x41, 0x37, 0x60, 0x55, 0xb8, 0xf8,
	0xce, 0xd8, 0xd9, 0xfb, 0xe0, 0xbe, 0x8e, 0xea, 0x6a, 0x30, 0x39, 0x67, 0x13, 0xea, 0x10, 0x44,
	0x36, 0xcf, 0x16, 0x87, 0xeb, 0x94, 0x7b, 0xd6, 0x2a, 0x0b, 0x62, 0x6f, 0x60, 0x7d, 0xe6, 0x27,
	0x2c, 0x0f, 0xa4, 0x08, 0xc4, 0x7e, 0x82, 0xc3, 0xc1, 0x2f, 0xd9, 0x2c, 0x05, 0x49, 0xc6, 0x42,
	0x20, 0x65, 0x51, 0x4c, 0x52, 0x7f, 0x94, 0xf0, 0xe6, 0x97, 0xf2, 0x2b, 0x76, 0xec, 0xc1, 0x42,
	0x65, 0x14, 0x19, 0xd7, 0x4a, 0x54, 0xd4, 0x88, 0x3c, 0x99, 0xb3, 0x7f, 0x7c, 0xa5, 0xa8, 0xe1,
	0xb7, 0xec, 0x34, 0x98, 0xba, 0x55, 0x14, 0x3d, 0x04, 0x19, 0x51, 0x92, 0x93, 0x55, 0xff, 0x79,
	0x3a, 0xcf, 0x16, 0xf9, 0x9a, 0xff, 0x95, 0x6f, 0xb8, 0x71, 0x2f, 0x8a, 0xe0, 0xf9, 0xee, 0xe3,
	0xa6, 0x36, 0xd4, 0xc4, 0x72, 0xa9, 0x9d, 0x2d, 0xac, 0xd1, 0x5b, 0xc0, 0xc7, 0x87, 0xc2, 0x85,
	0xaf, 0x08, 0xbe, 0xbb, 0x4e, 0x6b, 0x94, 0xf1, 0xb3, 0xc0, 0x6d, 0xfd, 0x34, 0x6e, 0x52, 0x4e,
	0x13, 0xbd, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x78, 0x58, 0x58, 0x17, 0x48, 0x01, 0x00, 0x00,
}
