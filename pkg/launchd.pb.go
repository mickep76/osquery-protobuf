// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/launchd.proto

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

// LaunchAgents and LaunchDaemons from default search paths.
type Launchd struct {
	// Path to daemon or agent plist
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// File name of plist (used by launchd)
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// Daemon or agent service name
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label"`
	// Path to target program
	Program string `protobuf:"bytes,4,opt,name=program,proto3" json:"program"`
	// Should the program run on launch load
	RunAtLoad string `protobuf:"bytes,5,opt,name=run_at_load,json=runAtLoad,proto3" json:"run_at_load"`
	// Should the process be restarted if killed
	KeepAlive string `protobuf:"bytes,6,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive"`
	// Deprecated key
	OnDemand string `protobuf:"bytes,7,opt,name=on_demand,json=onDemand,proto3" json:"on_demand"`
	// Skip loading this daemon or agent on boot
	Disabled string `protobuf:"bytes,8,opt,name=disabled,proto3" json:"disabled"`
	// Run this daemon or agent as this username
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username"`
	// Run this daemon or agent as this group
	Groupname string `protobuf:"bytes,10,opt,name=groupname,proto3" json:"groupname"`
	// Pipe stdout to a target path
	StdoutPath string `protobuf:"bytes,11,opt,name=stdout_path,json=stdoutPath,proto3" json:"stdout_path"`
	// Pipe stderr to a target path
	StderrPath string `protobuf:"bytes,12,opt,name=stderr_path,json=stderrPath,proto3" json:"stderr_path"`
	// Frequency to run in seconds
	StartInterval        string   `protobuf:"bytes,13,opt,name=start_interval,json=startInterval,proto3" json:"start_interval"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Launchd) Reset()         { *m = Launchd{} }
func (m *Launchd) String() string { return proto.CompactTextString(m) }
func (*Launchd) ProtoMessage()    {}
func (*Launchd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2a7c7291a7eb4eb, []int{0}
}

func (m *Launchd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Launchd.Unmarshal(m, b)
}
func (m *Launchd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Launchd.Marshal(b, m, deterministic)
}
func (m *Launchd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Launchd.Merge(m, src)
}
func (m *Launchd) XXX_Size() int {
	return xxx_messageInfo_Launchd.Size(m)
}
func (m *Launchd) XXX_DiscardUnknown() {
	xxx_messageInfo_Launchd.DiscardUnknown(m)
}

var xxx_messageInfo_Launchd proto.InternalMessageInfo

func (m *Launchd) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Launchd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Launchd) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Launchd) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *Launchd) GetRunAtLoad() string {
	if m != nil {
		return m.RunAtLoad
	}
	return ""
}

func (m *Launchd) GetKeepAlive() string {
	if m != nil {
		return m.KeepAlive
	}
	return ""
}

func (m *Launchd) GetOnDemand() string {
	if m != nil {
		return m.OnDemand
	}
	return ""
}

func (m *Launchd) GetDisabled() string {
	if m != nil {
		return m.Disabled
	}
	return ""
}

func (m *Launchd) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Launchd) GetGroupname() string {
	if m != nil {
		return m.Groupname
	}
	return ""
}

func (m *Launchd) GetStdoutPath() string {
	if m != nil {
		return m.StdoutPath
	}
	return ""
}

func (m *Launchd) GetStderrPath() string {
	if m != nil {
		return m.StderrPath
	}
	return ""
}

func (m *Launchd) GetStartInterval() string {
	if m != nil {
		return m.StartInterval
	}
	return ""
}

func init() {
	proto.RegisterType((*Launchd)(nil), "schemas.Launchd")
}

func init() { proto.RegisterFile("pb/launchd.proto", fileDescriptor_f2a7c7291a7eb4eb) }

var fileDescriptor_f2a7c7291a7eb4eb = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0xcf, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0x69, 0xbf, 0xb6, 0xdb, 0x9d, 0x7e, 0x15, 0x09, 0x1e, 0x82, 0xbf, 0x11, 0x04, 0x2f,
	0x76, 0x45, 0x41, 0x0f, 0x9e, 0x2a, 0x5e, 0x84, 0x1e, 0xc4, 0xa3, 0x97, 0x25, 0xbb, 0x89, 0xbb,
	0x4b, 0x77, 0x93, 0x38, 0x49, 0x0a, 0xfe, 0xef, 0x1e, 0x64, 0x27, 0x6d, 0x6f, 0xf3, 0x3e, 0xcf,
	0x0b, 0x09, 0x33, 0x70, 0x68, 0x8b, 0xac, 0x15, 0x41, 0x97, 0xb5, 0x5c, 0x58, 0x34, 0xde, 0xb0,
	0xc4, 0x95, 0xb5, 0xea, 0x84, 0xbb, 0xfa, 0x1d, 0x42, 0xb2, 0x8a, 0x8a, 0x31, 0x18, 0x59, 0xe1,
	0x6b, 0x3e, 0xb8, 0x1c, 0xdc, 0xa4, 0x1f, 0x34, 0xf7, 0x4c, 0x8b, 0x4e, 0xf1, 0x61, 0x64, 0xfd,
	0xcc, 0x8e, 0x60, 0xdc, 0x8a, 0x42, 0xb5, 0xfc, 0x1f, 0xc1, 0x18, 0x18, 0x87, 0xc4, 0xa2, 0xa9,
	0x50, 0x74, 0x7c, 0x44, 0x7c, 0x17, 0xd9, 0x39, 0xcc, 0x30, 0xe8, 0x5c, 0xf8, 0xbc, 0x35, 0x42,
	0xf2, 0x31, 0xd9, 0x14, 0x83, 0x5e, 0xfa, 0x95, 0x11, 0x92, 0x9d, 0x01, 0xac, 0x95, 0xb2, 0xb9,
	0x68, 0x9b, 0x8d, 0xe2, 0x93, 0xa8, 0x7b, 0xb2, 0xec, 0x01, 0x3b, 0x81, 0xd4, 0xe8, 0x5c, 0xaa,
	0x4e, 0x68, 0xc9, 0x13, 0xb2, 0x53, 0xa3, 0x5f, 0x29, 0xb3, 0x63, 0x98, 0xca, 0xc6, 0x89, 0xa2,
	0x55, 0x92, 0x4f, 0xa3, 0xdb, 0xe5, 0xde, 0x05, 0xa7, 0x90, 0xfe, 0x9f, 0x46, 0xb7, 0xcb, 0xec,
	0x14, 0xd2, 0x0a, 0x4d, 0xb0, 0x24, 0x21, 0x3e, 0xb9, 0x07, 0xec, 0x02, 0x66, 0xce, 0x4b, 0x13,
	0x7c, 0x4e, 0x0b, 0x99, 0x91, 0x87, 0x88, 0xde, 0xfb, 0xb5, 0xc4, 0x82, 0x42, 0x8c, 0x85, 0xff,
	0xfb, 0x82, 0x42, 0xa4, 0xc2, 0x35, 0x1c, 0x38, 0x2f, 0xd0, 0xe7, 0x8d, 0xf6, 0x0a, 0x37, 0xa2,
	0xe5, 0x73, 0xea, 0xcc, 0x89, 0xbe, 0x6d, 0xe1, 0xcb, 0xfd, 0xe7, 0x5d, 0xd5, 0xf8, 0x3a, 0x14,
	0x8b, 0xd2, 0x74, 0x59, 0xd7, 0x94, 0x6b, 0x65, 0x9f, 0x1e, 0x33, 0xe3, 0xbe, 0x83, 0xc2, 0x9f,
	0x5b, 0x3a, 0x56, 0x11, 0xbe, 0x32, 0xbb, 0xae, 0x9e, 0xb7, 0x27, 0x2b, 0x26, 0x44, 0x1f, 0xfe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x62, 0xca, 0xd2, 0xd6, 0x01, 0x00, 0x00,
}
