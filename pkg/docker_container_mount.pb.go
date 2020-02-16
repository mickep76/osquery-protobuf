// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_container_mount.proto

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

// Docker container mounts.
type DockerContainerMount struct {
	// Container ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// Type of mount (bind
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// Optional mount name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// Source path on host
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
	// Destination path inside container
	Destination string `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination"`
	// Driver providing the mount
	Driver string `protobuf:"bytes,6,opt,name=driver,proto3" json:"driver"`
	// Mount options (rw
	Mode string `protobuf:"bytes,7,opt,name=mode,proto3" json:"mode"`
	// 1 if read/write. 0 otherwise
	Rw int32 `protobuf:"varint,8,opt,name=rw,proto3" json:"rw"`
	// Mount propagation
	Propagation          string   `protobuf:"bytes,9,opt,name=propagation,proto3" json:"propagation"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerContainerMount) Reset()         { *m = DockerContainerMount{} }
func (m *DockerContainerMount) String() string { return proto.CompactTextString(m) }
func (*DockerContainerMount) ProtoMessage()    {}
func (*DockerContainerMount) Descriptor() ([]byte, []int) {
	return fileDescriptor_de4d22d687cb206a, []int{0}
}

func (m *DockerContainerMount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerContainerMount.Unmarshal(m, b)
}
func (m *DockerContainerMount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerContainerMount.Marshal(b, m, deterministic)
}
func (m *DockerContainerMount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerContainerMount.Merge(m, src)
}
func (m *DockerContainerMount) XXX_Size() int {
	return xxx_messageInfo_DockerContainerMount.Size(m)
}
func (m *DockerContainerMount) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerContainerMount.DiscardUnknown(m)
}

var xxx_messageInfo_DockerContainerMount proto.InternalMessageInfo

func (m *DockerContainerMount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DockerContainerMount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DockerContainerMount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerContainerMount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *DockerContainerMount) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *DockerContainerMount) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *DockerContainerMount) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *DockerContainerMount) GetRw() int32 {
	if m != nil {
		return m.Rw
	}
	return 0
}

func (m *DockerContainerMount) GetPropagation() string {
	if m != nil {
		return m.Propagation
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerContainerMount)(nil), "schemas.DockerContainerMount")
}

func init() { proto.RegisterFile("pb/docker_container_mount.proto", fileDescriptor_de4d22d687cb206a) }

var fileDescriptor_de4d22d687cb206a = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xd0, 0xa6, 0xd4, 0x48, 0x0c, 0x16, 0x42, 0xde, 0x88, 0x98, 0xba, 0xd0, 0x20,
	0x90, 0x60, 0x60, 0x03, 0x56, 0x96, 0x8e, 0x2c, 0x55, 0x62, 0x1f, 0xa9, 0x15, 0xd9, 0x67, 0xce,
	0x36, 0x55, 0xff, 0x30, 0xbf, 0x03, 0xd9, 0x2e, 0x52, 0xb7, 0x77, 0x9f, 0xdf, 0xbd, 0xb3, 0x1e,
	0xbb, 0x71, 0x43, 0xa7, 0x50, 0x4e, 0x40, 0x5b, 0x89, 0x36, 0xf4, 0xda, 0x02, 0x6d, 0x0d, 0x46,
	0x1b, 0xd6, 0x8e, 0x30, 0x20, 0x5f, 0x78, 0xb9, 0x03, 0xd3, 0xfb, 0xdb, 0xdf, 0x8a, 0x5d, 0xbd,
	0x67, 0xe7, 0xdb, 0xbf, 0xf1, 0x23, 0xf9, 0xf8, 0x25, 0xab, 0xb5, 0x12, 0x55, 0x5b, 0xad, 0x96,
	0x9b, 0x5a, 0x2b, 0xce, 0xd9, 0x2c, 0x1c, 0x1c, 0x88, 0x3a, 0x93, 0xac, 0x13, 0xb3, 0xbd, 0x01,
	0x71, 0x56, 0x58, 0xd2, 0xfc, 0x9a, 0x35, 0x1e, 0x23, 0x49, 0x10, 0xb3, 0x4c, 0x8f, 0x13, 0x6f,
	0xd9, 0x85, 0x02, 0x1f, 0xb4, 0xed, 0x83, 0x46, 0x2b, 0xe6, 0xf9, 0xf1, 0x14, 0xa5, 0x4d, 0x45,
	0xfa, 0x07, 0x48, 0x34, 0x65, 0xb3, 0x4c, 0xe9, 0x8a, 0x41, 0x05, 0x62, 0x51, 0xae, 0x24, 0x9d,
	0x7e, 0x47, 0x7b, 0x71, 0xde, 0x56, 0xab, 0xf9, 0xa6, 0xa6, 0x7d, 0x4a, 0x77, 0x84, 0xae, 0x1f,
	0x4b, 0xfa, 0xb2, 0xa4, 0x9f, 0xa0, 0xd7, 0x87, 0xcf, 0xfb, 0x51, 0x87, 0x5d, 0x1c, 0xd6, 0x12,
	0x4d, 0x67, 0xb4, 0x9c, 0xc0, 0x3d, 0x3f, 0x75, 0xe8, 0xbf, 0x23, 0xd0, 0xe1, 0x2e, 0xd7, 0x32,
	0xc4, 0xaf, 0xce, 0x4d, 0xe3, 0xcb, 0xb1, 0x9c, 0xa1, 0xc9, 0xf4, 0xf1, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x7e, 0x97, 0x7e, 0x4f, 0x01, 0x00, 0x00,
}