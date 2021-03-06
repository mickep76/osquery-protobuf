// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

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

// Lists all installed Windows services and their relevant data.
type Service struct {
	// Service name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Service Type: OWN_PROCESS
	ServiceType string `protobuf:"bytes,2,opt,name=service_type,json=serviceType,proto3" json:"service_type"`
	// Service Display name
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	// Service Current status: STOPPED
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`
	// the Process ID of the service
	Pid int32 `protobuf:"varint,5,opt,name=pid,proto3" json:"pid"`
	// Service start type: BOOT_START
	StartType string `protobuf:"bytes,6,opt,name=start_type,json=startType,proto3" json:"start_type"`
	// The error code that the service uses to report an error that occurs when it is starting or stopping
	Win32ExitCode int32 `protobuf:"varint,7,opt,name=win32_exit_code,json=win32ExitCode,proto3" json:"win32_exit_code"`
	// The service-specific error code that the service returns when an error occurs while the service is starting or stopping
	ServiceExitCode int32 `protobuf:"varint,8,opt,name=service_exit_code,json=serviceExitCode,proto3" json:"service_exit_code"`
	// Path to Service Executable
	Path string `protobuf:"bytes,9,opt,name=path,proto3" json:"path"`
	// Path to ServiceDll
	ModulePath string `protobuf:"bytes,10,opt,name=module_path,json=modulePath,proto3" json:"module_path"`
	// Service Description
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description"`
	// The name of the account that the service process will be logged on as when it runs. This name can be of the form Domain\\UserName. If the account belongs to the built-in domain
	UserAccount          string   `protobuf:"bytes,12,opt,name=user_account,json=userAccount,proto3" json:"user_account"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *Service) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Service) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Service) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Service) GetStartType() string {
	if m != nil {
		return m.StartType
	}
	return ""
}

func (m *Service) GetWin32ExitCode() int32 {
	if m != nil {
		return m.Win32ExitCode
	}
	return 0
}

func (m *Service) GetServiceExitCode() int32 {
	if m != nil {
		return m.ServiceExitCode
	}
	return 0
}

func (m *Service) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Service) GetModulePath() string {
	if m != nil {
		return m.ModulePath
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetUserAccount() string {
	if m != nil {
		return m.UserAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*Service)(nil), "schemas.Service")
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptor_6ff5ab49d8a5fcc4) }

var fileDescriptor_6ff5ab49d8a5fcc4 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0x47, 0xd5, 0xff, 0xf4, 0x52, 0xd4, 0xe2, 0x01, 0x79, 0x41, 0x14, 0x06, 0x54, 0x21, 0xd1,
	0xa0, 0x56, 0x82, 0x81, 0x09, 0x10, 0x2b, 0x42, 0x85, 0x89, 0x25, 0x72, 0x1c, 0xd3, 0x58, 0x6d,
	0x62, 0x13, 0xdb, 0xd0, 0x7c, 0x75, 0x26, 0x94, 0xb3, 0xa1, 0xdb, 0xe5, 0xfd, 0xde, 0xe5, 0x72,
	0x39, 0x98, 0xe8, 0x34, 0x36, 0xa2, 0xfa, 0x92, 0x5c, 0xcc, 0x75, 0xa5, 0xac, 0x22, 0x03, 0xc3,
	0x73, 0x51, 0x30, 0x73, 0xfe, 0xd3, 0x86, 0xc1, 0xab, 0x8f, 0x08, 0x81, 0x6e, 0xc9, 0x0a, 0x41,
	0x5b, 0xd3, 0xd6, 0x6c, 0xb8, 0xc2, 0x9a, 0x9c, 0xc1, 0x28, 0x74, 0x26, 0xb6, 0xd6, 0x82, 0xb6,
	0x31, 0x8b, 0x02, 0x7b, 0xab, 0x35, 0x2a, 0x99, 0x34, 0x7a, 0xcb, 0xea, 0x04, 0xdb, 0x3b, 0x5e,
	0x09, 0xec, 0xb9, 0x79, 0xcb, 0x31, 0xf4, 0x8d, 0x65, 0xd6, 0x19, 0xda, 0xc5, 0x30, 0x3c, 0x91,
	0x09, 0x74, 0xb4, 0xcc, 0x68, 0x6f, 0xda, 0x9a, 0xf5, 0x56, 0x4d, 0x49, 0x4e, 0x00, 0x8c, 0x65,
	0x95, 0xf5, 0xd3, 0xfa, 0x68, 0x0f, 0x91, 0xe0, 0xac, 0x0b, 0x18, 0x7f, 0xcb, 0x72, 0xb9, 0x48,
	0xc4, 0x4e, 0xda, 0x84, 0xab, 0x4c, 0xd0, 0x01, 0x36, 0x1f, 0x22, 0x7e, 0xda, 0x49, 0xfb, 0xa8,
	0x32, 0x41, 0x2e, 0xe1, 0xe8, 0xef, 0xb3, 0xf7, 0xe6, 0x01, 0x9a, 0xe3, 0x10, 0xfc, 0xbb, 0x04,
	0xba, 0x9a, 0xd9, 0x9c, 0x0e, 0xfd, 0xda, 0x4d, 0x4d, 0x4e, 0x21, 0x2a, 0x54, 0xe6, 0xb6, 0x22,
	0xc1, 0x08, 0x30, 0x02, 0x8f, 0x5e, 0x1a, 0x61, 0x0a, 0x51, 0x26, 0x0c, 0xaf, 0xa4, 0xb6, 0x52,
	0x95, 0x34, 0x0a, 0x3b, 0xef, 0x51, 0xf3, 0x5b, 0x9c, 0x11, 0x55, 0xc2, 0x38, 0x57, 0xae, 0xb4,
	0x74, 0xe4, 0x95, 0x86, 0xdd, 0x7b, 0xf4, 0xb0, 0x78, 0xbf, 0x5e, 0x4b, 0x9b, 0xbb, 0x74, 0xce,
	0x55, 0x11, 0x17, 0x92, 0x6f, 0x84, 0xbe, 0xbd, 0x89, 0x95, 0xf9, 0x74, 0xa2, 0xaa, 0xaf, 0xf0,
	0x54, 0xa9, 0xfb, 0x88, 0xf5, 0x66, 0x7d, 0x17, 0x0e, 0x96, 0xf6, 0x91, 0x2e, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x28, 0x55, 0x85, 0x54, 0xd4, 0x01, 0x00, 0x00,
}
