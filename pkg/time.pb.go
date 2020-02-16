// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/time.proto

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

// Track current date and time in the system.
type Time struct {
	// Current weekday in the system
	Weekday string `protobuf:"bytes,1,opt,name=weekday,proto3" json:"weekday"`
	// Current year in the system
	Year int32 `protobuf:"varint,2,opt,name=year,proto3" json:"year"`
	// Current month in the system
	Month int32 `protobuf:"varint,3,opt,name=month,proto3" json:"month"`
	// Current day in the system
	Day int32 `protobuf:"varint,4,opt,name=day,proto3" json:"day"`
	// Current hour in the system
	Hour int32 `protobuf:"varint,5,opt,name=hour,proto3" json:"hour"`
	// Current minutes in the system
	Minutes int32 `protobuf:"varint,6,opt,name=minutes,proto3" json:"minutes"`
	// Current seconds in the system
	Seconds int32 `protobuf:"varint,7,opt,name=seconds,proto3" json:"seconds"`
	// Current timezone in the system
	Timezone string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone"`
	// Current local timezone in the system
	LocalTimezone string `protobuf:"bytes,9,opt,name=local_timezone,json=localTimezone,proto3" json:"local_timezone"`
	// Current timestamp (log format) in the system
	Timestamp string `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp"`
	// Current date and time (ISO format
	Datetime string `protobuf:"bytes,11,opt,name=datetime,proto3" json:"datetime"`
	// Current time (ISO format) in the system
	Iso_8601 string `protobuf:"bytes,12,opt,name=iso_8601,json=iso8601,proto3" json:"iso_8601"`
	// Timestamp value in 100 nanosecond units.
	WinTimestamp         int64    `protobuf:"varint,13,opt,name=win_timestamp,json=winTimestamp,proto3" json:"win_timestamp"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1ae1173ee2b89a0, []int{0}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetWeekday() string {
	if m != nil {
		return m.Weekday
	}
	return ""
}

func (m *Time) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Time) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Time) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Time) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *Time) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *Time) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Time) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Time) GetLocalTimezone() string {
	if m != nil {
		return m.LocalTimezone
	}
	return ""
}

func (m *Time) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Time) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *Time) GetIso_8601() string {
	if m != nil {
		return m.Iso_8601
	}
	return ""
}

func (m *Time) GetWinTimestamp() int64 {
	if m != nil {
		return m.WinTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Time)(nil), "schemas.Time")
}

func init() { proto.RegisterFile("pb/time.proto", fileDescriptor_a1ae1173ee2b89a0) }

var fileDescriptor_a1ae1173ee2b89a0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0x87, 0x43, 0xe9, 0xdf, 0xb1, 0x18, 0xb3, 0xf1, 0xb0, 0x1a, 0x0f, 0x8d, 0xc6, 0xa4, 0x17,
	0x4b, 0xd5, 0xa4, 0x9a, 0x78, 0xf3, 0x11, 0x1a, 0x4e, 0x5e, 0xc8, 0x02, 0x63, 0xd9, 0xd0, 0xdd,
	0x45, 0x76, 0x09, 0xc1, 0xb7, 0xf5, 0x4d, 0xcc, 0x2e, 0xd0, 0xde, 0xf6, 0xfb, 0xe6, 0x37, 0xcc,
	0x84, 0x81, 0xa0, 0x4c, 0x42, 0xc3, 0x05, 0x6e, 0xca, 0x4a, 0x19, 0x45, 0x66, 0x3a, 0xcd, 0x51,
	0x30, 0x7d, 0xff, 0x37, 0x82, 0x71, 0xc4, 0x05, 0x12, 0x0a, 0xb3, 0x06, 0xb1, 0xc8, 0x58, 0x4b,
	0xbd, 0x95, 0xb7, 0x5e, 0xec, 0x07, 0x24, 0x04, 0xc6, 0x2d, 0xb2, 0x8a, 0x8e, 0x56, 0xde, 0x7a,
	0xb2, 0x77, 0x6f, 0x72, 0x0d, 0x13, 0xa1, 0xa4, 0xc9, 0xa9, 0xef, 0x64, 0x07, 0xe4, 0x0a, 0x7c,
	0xdb, 0x3f, 0x76, 0xce, 0xef, 0x7b, 0x73, 0x55, 0x57, 0x74, 0xd2, 0xf5, 0xda, 0xb7, 0x9d, 0x24,
	0xb8, 0xac, 0x0d, 0x6a, 0x3a, 0x75, 0x7a, 0x40, 0x5b, 0xd1, 0x98, 0x2a, 0x99, 0x69, 0x3a, 0xeb,
	0x2a, 0x3d, 0x92, 0x5b, 0x98, 0xdb, 0xed, 0x7f, 0x95, 0x44, 0x3a, 0x77, 0xeb, 0x9d, 0x98, 0x3c,
	0xc2, 0xe5, 0x51, 0xa5, 0xec, 0x18, 0x9f, 0x12, 0x0b, 0x97, 0x08, 0x9c, 0x8d, 0x86, 0xd8, 0x1d,
	0x2c, 0x6c, 0x40, 0x1b, 0x26, 0x4a, 0x0a, 0x2e, 0x71, 0x16, 0x76, 0x40, 0xc6, 0x0c, 0x5a, 0x41,
	0x2f, 0xba, 0x01, 0x03, 0x93, 0x1b, 0x98, 0x73, 0xad, 0xe2, 0xf7, 0xdd, 0xf6, 0x99, 0x2e, 0xbb,
	0x7f, 0xc3, 0xb5, 0xb2, 0x48, 0x1e, 0x20, 0x68, 0xb8, 0x8c, 0xcf, 0x1f, 0x0e, 0x56, 0xde, 0xda,
	0xdf, 0x2f, 0x1b, 0x2e, 0xa3, 0xc1, 0x7d, 0xbe, 0x7c, 0x6d, 0x0f, 0xdc, 0xe4, 0x75, 0xb2, 0x49,
	0x95, 0x08, 0x05, 0x4f, 0x0b, 0x2c, 0xdf, 0x76, 0xa1, 0xd2, 0x3f, 0x35, 0x56, 0xed, 0x93, 0xbb,
	0x48, 0x52, 0x7f, 0x87, 0x65, 0x71, 0xf8, 0xe8, 0xef, 0x92, 0x4c, 0x9d, 0x7d, 0xfd, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x11, 0x3c, 0x40, 0xb8, 0x01, 0x00, 0x00,
}