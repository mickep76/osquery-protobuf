// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/windows_event.proto

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

// Windows Event logs.
type WindowsEvent struct {
	// Timestamp the event was received
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	// System time at which the event occurred
	Datetime string `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime"`
	// Source or channel of the event
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
	// Provider name of the event
	ProviderName string `protobuf:"bytes,4,opt,name=provider_name,json=providerName,proto3" json:"provider_name"`
	// Provider guid of the event
	ProviderGuid string `protobuf:"bytes,5,opt,name=provider_guid,json=providerGuid,proto3" json:"provider_guid"`
	// Event ID of the event
	Eventid int32 `protobuf:"varint,6,opt,name=eventid,proto3" json:"eventid"`
	// Task value associated with the event
	Task int32 `protobuf:"varint,7,opt,name=task,proto3" json:"task"`
	// The severity level associated with the event
	Level int32 `protobuf:"varint,8,opt,name=level,proto3" json:"level"`
	// A bitmask of the keywords defined in the event
	Keywords int64 `protobuf:"varint,9,opt,name=keywords,proto3" json:"keywords"`
	// Data associated with the event
	Data string `protobuf:"bytes,10,opt,name=data,proto3" json:"data"`
	// Event ID
	Eid                  string   `protobuf:"bytes,11,opt,name=eid,proto3" json:"eid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WindowsEvent) Reset()         { *m = WindowsEvent{} }
func (m *WindowsEvent) String() string { return proto.CompactTextString(m) }
func (*WindowsEvent) ProtoMessage()    {}
func (*WindowsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8aedd829455ea58, []int{0}
}

func (m *WindowsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowsEvent.Unmarshal(m, b)
}
func (m *WindowsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowsEvent.Marshal(b, m, deterministic)
}
func (m *WindowsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowsEvent.Merge(m, src)
}
func (m *WindowsEvent) XXX_Size() int {
	return xxx_messageInfo_WindowsEvent.Size(m)
}
func (m *WindowsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WindowsEvent proto.InternalMessageInfo

func (m *WindowsEvent) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *WindowsEvent) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *WindowsEvent) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *WindowsEvent) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *WindowsEvent) GetProviderGuid() string {
	if m != nil {
		return m.ProviderGuid
	}
	return ""
}

func (m *WindowsEvent) GetEventid() int32 {
	if m != nil {
		return m.Eventid
	}
	return 0
}

func (m *WindowsEvent) GetTask() int32 {
	if m != nil {
		return m.Task
	}
	return 0
}

func (m *WindowsEvent) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *WindowsEvent) GetKeywords() int64 {
	if m != nil {
		return m.Keywords
	}
	return 0
}

func (m *WindowsEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *WindowsEvent) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func init() {
	proto.RegisterType((*WindowsEvent)(nil), "schemas.WindowsEvent")
}

func init() { proto.RegisterFile("pb/windows_event.proto", fileDescriptor_d8aedd829455ea58) }

var fileDescriptor_d8aedd829455ea58 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0x24, 0xed, 0x5a, 0x41, 0x16, 0x29, 0x8b, 0xa7, 0xa0, 0x97, 0x5c, 0x6c,
	0x44, 0x41, 0x0f, 0xde, 0x04, 0xf1, 0xe6, 0x21, 0x17, 0xc1, 0x4b, 0xd9, 0x64, 0xc7, 0x74, 0x49,
	0x37, 0x1b, 0xf7, 0x4f, 0x42, 0xbf, 0x90, 0x9f, 0x53, 0x32, 0xfd, 0x03, 0xbd, 0xbd, 0xf7, 0x7b,
	0x8f, 0xdd, 0x99, 0x21, 0xcb, 0xae, 0xcc, 0x07, 0xd9, 0x0a, 0x3d, 0xd8, 0x35, 0xf4, 0xd0, 0xba,
	0x55, 0x67, 0xb4, 0xd3, 0x34, 0xb1, 0xd5, 0x06, 0x14, 0xb7, 0xb7, 0x7f, 0x13, 0xb2, 0xf8, 0xda,
	0x17, 0xde, 0xc7, 0x9c, 0x52, 0x32, 0x75, 0x52, 0x01, 0x0b, 0xd2, 0x20, 0x0b, 0x0b, 0xd4, 0xf4,
	0x86, 0xcc, 0x04, 0x77, 0x80, 0x7c, 0x92, 0x06, 0xd9, 0xbc, 0x38, 0x79, 0xba, 0x24, 0xb1, 0xd5,
	0xde, 0x54, 0xc0, 0x42, 0x4c, 0x0e, 0x8e, 0xde, 0x91, 0xcb, 0xce, 0xe8, 0x5e, 0x0a, 0x30, 0xeb,
	0x96, 0x2b, 0x60, 0x53, 0x8c, 0x17, 0x47, 0xf8, 0xc9, 0xd5, 0x79, 0xa9, 0xf6, 0x52, 0xb0, 0xe8,
	0xbc, 0xf4, 0xe1, 0xa5, 0xa0, 0x8c, 0x24, 0x38, 0xba, 0x14, 0x2c, 0x4e, 0x83, 0x2c, 0x2a, 0x8e,
	0x16, 0x67, 0xe5, 0xb6, 0x61, 0x09, 0x62, 0xd4, 0xf4, 0x9a, 0x44, 0x5b, 0xe8, 0x61, 0xcb, 0x66,
	0x08, 0xf7, 0x66, 0xdc, 0xa0, 0x81, 0xdd, 0xa0, 0x8d, 0xb0, 0x6c, 0x8e, 0x9b, 0x9d, 0xfc, 0xf8,
	0x8a, 0xe0, 0x8e, 0x33, 0x82, 0x7f, 0xa3, 0xa6, 0x57, 0x24, 0x04, 0x29, 0xd8, 0x05, 0xa2, 0x51,
	0xbe, 0x3d, 0x7e, 0x3f, 0xd4, 0xd2, 0x6d, 0x7c, 0xb9, 0xaa, 0xb4, 0xca, 0x95, 0xac, 0x1a, 0xe8,
	0x5e, 0x9e, 0x73, 0x6d, 0x7f, 0x3d, 0x98, 0xdd, 0x3d, 0x9e, 0xb5, 0xf4, 0x3f, 0x79, 0xd7, 0xd4,
	0xaf, 0x87, 0xe3, 0x96, 0x31, 0xd2, 0xa7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x22, 0x1b,
	0x8d, 0x86, 0x01, 0x00, 0x00,
}