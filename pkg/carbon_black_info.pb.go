// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/carbon_black_info.proto

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

// Returns info about a Carbon Black sensor install.
type CarbonBlackInfo struct {
	// Sensor ID of the Carbon Black sensor
	SensorId int32 `protobuf:"varint,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id"`
	// Sensor group
	ConfigName string `protobuf:"bytes,2,opt,name=config_name,json=configName,proto3" json:"config_name"`
	// If the sensor is configured to send back binaries to the Carbon Black server
	CollectStoreFiles int32 `protobuf:"varint,3,opt,name=collect_store_files,json=collectStoreFiles,proto3" json:"collect_store_files"`
	// If the sensor is configured to capture module loads
	CollectModuleLoads int32 `protobuf:"varint,4,opt,name=collect_module_loads,json=collectModuleLoads,proto3" json:"collect_module_loads"`
	// If the sensor is configured to collect metadata of binaries
	CollectModuleInfo int32 `protobuf:"varint,5,opt,name=collect_module_info,json=collectModuleInfo,proto3" json:"collect_module_info"`
	// If the sensor is configured to collect file modification events
	CollectFileMods int32 `protobuf:"varint,6,opt,name=collect_file_mods,json=collectFileMods,proto3" json:"collect_file_mods"`
	// If the sensor is configured to collect registry modification events
	CollectRegMods int32 `protobuf:"varint,7,opt,name=collect_reg_mods,json=collectRegMods,proto3" json:"collect_reg_mods"`
	// If the sensor is configured to collect network connections
	CollectNetConns int32 `protobuf:"varint,8,opt,name=collect_net_conns,json=collectNetConns,proto3" json:"collect_net_conns"`
	// If the sensor is configured to process events
	CollectProcesses int32 `protobuf:"varint,9,opt,name=collect_processes,json=collectProcesses,proto3" json:"collect_processes"`
	// If the sensor is configured to cross process events
	CollectCrossProcesses int32 `protobuf:"varint,10,opt,name=collect_cross_processes,json=collectCrossProcesses,proto3" json:"collect_cross_processes"`
	// If the sensor is configured to EMET events
	CollectEmetEvents int32 `protobuf:"varint,11,opt,name=collect_emet_events,json=collectEmetEvents,proto3" json:"collect_emet_events"`
	// If the sensor is configured to collect non binary file writes
	CollectDataFileWrites int32 `protobuf:"varint,12,opt,name=collect_data_file_writes,json=collectDataFileWrites,proto3" json:"collect_data_file_writes"`
	// If the sensor is configured to collect the user running a process
	CollectProcessUserContext int32 `protobuf:"varint,13,opt,name=collect_process_user_context,json=collectProcessUserContext,proto3" json:"collect_process_user_context"`
	// Unknown
	CollectSensorOperations int32 `protobuf:"varint,14,opt,name=collect_sensor_operations,json=collectSensorOperations,proto3" json:"collect_sensor_operations"`
	// Event file disk quota in MB
	LogFileDiskQuotaMb int32 `protobuf:"varint,15,opt,name=log_file_disk_quota_mb,json=logFileDiskQuotaMb,proto3" json:"log_file_disk_quota_mb"`
	// Event file disk quota in a percentage
	LogFileDiskQuotaPercentage int32 `protobuf:"varint,16,opt,name=log_file_disk_quota_percentage,json=logFileDiskQuotaPercentage,proto3" json:"log_file_disk_quota_percentage"`
	// If the sensor is configured to report tamper events
	ProtectionDisabled int32 `protobuf:"varint,17,opt,name=protection_disabled,json=protectionDisabled,proto3" json:"protection_disabled"`
	// IP address of the sensor
	SensorIpAddr string `protobuf:"bytes,18,opt,name=sensor_ip_addr,json=sensorIpAddr,proto3" json:"sensor_ip_addr"`
	// Carbon Black server
	SensorBackendServer string `protobuf:"bytes,19,opt,name=sensor_backend_server,json=sensorBackendServer,proto3" json:"sensor_backend_server"`
	// Size in bytes of Carbon Black event files on disk
	EventQueue int32 `protobuf:"varint,20,opt,name=event_queue,json=eventQueue,proto3" json:"event_queue"`
	// Size in bytes of binaries waiting to be sent to Carbon Black server
	BinaryQueue          int32    `protobuf:"varint,21,opt,name=binary_queue,json=binaryQueue,proto3" json:"binary_queue"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarbonBlackInfo) Reset()         { *m = CarbonBlackInfo{} }
func (m *CarbonBlackInfo) String() string { return proto.CompactTextString(m) }
func (*CarbonBlackInfo) ProtoMessage()    {}
func (*CarbonBlackInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ea9b56d58ba2c, []int{0}
}

func (m *CarbonBlackInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarbonBlackInfo.Unmarshal(m, b)
}
func (m *CarbonBlackInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarbonBlackInfo.Marshal(b, m, deterministic)
}
func (m *CarbonBlackInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarbonBlackInfo.Merge(m, src)
}
func (m *CarbonBlackInfo) XXX_Size() int {
	return xxx_messageInfo_CarbonBlackInfo.Size(m)
}
func (m *CarbonBlackInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CarbonBlackInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CarbonBlackInfo proto.InternalMessageInfo

func (m *CarbonBlackInfo) GetSensorId() int32 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *CarbonBlackInfo) GetConfigName() string {
	if m != nil {
		return m.ConfigName
	}
	return ""
}

func (m *CarbonBlackInfo) GetCollectStoreFiles() int32 {
	if m != nil {
		return m.CollectStoreFiles
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectModuleLoads() int32 {
	if m != nil {
		return m.CollectModuleLoads
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectModuleInfo() int32 {
	if m != nil {
		return m.CollectModuleInfo
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectFileMods() int32 {
	if m != nil {
		return m.CollectFileMods
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectRegMods() int32 {
	if m != nil {
		return m.CollectRegMods
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectNetConns() int32 {
	if m != nil {
		return m.CollectNetConns
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectProcesses() int32 {
	if m != nil {
		return m.CollectProcesses
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectCrossProcesses() int32 {
	if m != nil {
		return m.CollectCrossProcesses
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectEmetEvents() int32 {
	if m != nil {
		return m.CollectEmetEvents
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectDataFileWrites() int32 {
	if m != nil {
		return m.CollectDataFileWrites
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectProcessUserContext() int32 {
	if m != nil {
		return m.CollectProcessUserContext
	}
	return 0
}

func (m *CarbonBlackInfo) GetCollectSensorOperations() int32 {
	if m != nil {
		return m.CollectSensorOperations
	}
	return 0
}

func (m *CarbonBlackInfo) GetLogFileDiskQuotaMb() int32 {
	if m != nil {
		return m.LogFileDiskQuotaMb
	}
	return 0
}

func (m *CarbonBlackInfo) GetLogFileDiskQuotaPercentage() int32 {
	if m != nil {
		return m.LogFileDiskQuotaPercentage
	}
	return 0
}

func (m *CarbonBlackInfo) GetProtectionDisabled() int32 {
	if m != nil {
		return m.ProtectionDisabled
	}
	return 0
}

func (m *CarbonBlackInfo) GetSensorIpAddr() string {
	if m != nil {
		return m.SensorIpAddr
	}
	return ""
}

func (m *CarbonBlackInfo) GetSensorBackendServer() string {
	if m != nil {
		return m.SensorBackendServer
	}
	return ""
}

func (m *CarbonBlackInfo) GetEventQueue() int32 {
	if m != nil {
		return m.EventQueue
	}
	return 0
}

func (m *CarbonBlackInfo) GetBinaryQueue() int32 {
	if m != nil {
		return m.BinaryQueue
	}
	return 0
}

func init() {
	proto.RegisterType((*CarbonBlackInfo)(nil), "schemas.CarbonBlackInfo")
}

func init() { proto.RegisterFile("pb/carbon_black_info.proto", fileDescriptor_416ea9b56d58ba2c) }

var fileDescriptor_416ea9b56d58ba2c = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdb, 0x4f, 0x14, 0x31,
	0x14, 0xc6, 0x83, 0xca, 0xad, 0x20, 0x97, 0x2e, 0x68, 0x41, 0xa3, 0x68, 0x7c, 0x20, 0x1a, 0x59,
	0x82, 0x09, 0x24, 0xfa, 0x60, 0xdc, 0x05, 0x13, 0x12, 0x41, 0x2e, 0x31, 0x26, 0xbe, 0x34, 0x9d,
	0xf6, 0xec, 0x30, 0x99, 0x99, 0x76, 0xb6, 0xed, 0xa0, 0xfc, 0xe1, 0xbe, 0x9b, 0x9e, 0xce, 0xec,
	0x85, 0xf8, 0xfa, 0x7d, 0xbf, 0xaf, 0xe7, 0xf4, 0x4c, 0xcf, 0x90, 0xed, 0x2a, 0xe9, 0x4a, 0x61,
	0x13, 0xa3, 0x79, 0x52, 0x08, 0x99, 0xf3, 0x4c, 0x0f, 0xcc, 0x5e, 0x65, 0x8d, 0x37, 0x74, 0xde,
	0xc9, 0x1b, 0x28, 0x85, 0x7b, 0xfd, 0x77, 0x9e, 0xac, 0xf6, 0x11, 0xea, 0x05, 0xe6, 0x54, 0x0f,
	0x0c, 0x7d, 0x46, 0x16, 0x1d, 0x68, 0x67, 0x2c, 0xcf, 0x14, 0x9b, 0xd9, 0x99, 0xd9, 0x9d, 0xbd,
	0x5a, 0x88, 0xc2, 0xa9, 0xa2, 0x2f, 0xc9, 0x92, 0x34, 0x7a, 0x90, 0xa5, 0x5c, 0x8b, 0x12, 0xd8,
	0x83, 0x9d, 0x99, 0xdd, 0xc5, 0x2b, 0x12, 0xa5, 0x73, 0x51, 0x02, 0xdd, 0x23, 0x1d, 0x69, 0x8a,
	0x02, 0xa4, 0xe7, 0xce, 0x1b, 0x0b, 0x7c, 0x90, 0x15, 0xe0, 0xd8, 0x43, 0x3c, 0x67, 0xbd, 0xb1,
	0xae, 0x83, 0xf3, 0x35, 0x18, 0x74, 0x9f, 0x6c, 0xb4, 0x7c, 0x69, 0x54, 0x5d, 0x00, 0x2f, 0x8c,
	0x50, 0x8e, 0x3d, 0xc2, 0x00, 0x6d, 0xbc, 0x33, 0xb4, 0xbe, 0x05, 0x67, 0xb2, 0x42, 0x93, 0x08,
	0x37, 0x63, 0xb3, 0x53, 0x15, 0x62, 0x00, 0xef, 0xf3, 0x96, 0xb4, 0x22, 0xf6, 0x12, 0x42, 0x8e,
	0xcd, 0x21, 0xbd, 0xda, 0x18, 0xa1, 0x95, 0x33, 0xa3, 0x1c, 0xdd, 0x25, 0x6b, 0x2d, 0x6b, 0x21,
	0x8d, 0xe8, 0x3c, 0xa2, 0x2b, 0x8d, 0x7e, 0x05, 0x29, 0x92, 0x13, 0xa7, 0x6a, 0xf0, 0x5c, 0x1a,
	0xad, 0x1d, 0x5b, 0x98, 0x3a, 0xf5, 0x1c, 0x7c, 0x3f, 0xc8, 0xf4, 0xdd, 0x98, 0xad, 0xac, 0x91,
	0xe0, 0x1c, 0x38, 0xb6, 0x88, 0x6c, 0x5b, 0xee, 0xa2, 0xd5, 0xe9, 0x21, 0x79, 0xda, 0xc2, 0xd2,
	0x1a, 0xe7, 0x26, 0x22, 0x04, 0x23, 0x9b, 0x8d, 0xdd, 0x0f, 0xee, 0x38, 0x37, 0x31, 0x16, 0x28,
	0xc1, 0x73, 0xb8, 0x05, 0xed, 0x1d, 0x5b, 0x9a, 0x1a, 0xcb, 0x49, 0x09, 0xfe, 0x04, 0x0d, 0x7a,
	0x44, 0x58, 0xcb, 0x2b, 0xe1, 0x45, 0x9c, 0xcd, 0x6f, 0x9b, 0x79, 0x70, 0x6c, 0x79, 0xaa, 0xd0,
	0xb1, 0xf0, 0x22, 0x4c, 0xe8, 0x27, 0x9a, 0xf4, 0x33, 0x79, 0x7e, 0xef, 0x36, 0xbc, 0x76, 0x60,
	0xc3, 0x08, 0x3c, 0xfc, 0xf1, 0xec, 0x31, 0x86, 0xb7, 0xa6, 0x2f, 0xf6, 0xc3, 0x81, 0xed, 0x47,
	0x80, 0x7e, 0x24, 0x5b, 0xa3, 0x27, 0x12, 0x1f, 0x9a, 0xa9, 0xc0, 0x0a, 0x9f, 0x19, 0xed, 0xd8,
	0x0a, 0xa6, 0xdb, 0x11, 0x5c, 0xa3, 0xff, 0x7d, 0x64, 0xd3, 0x03, 0xf2, 0xa4, 0x30, 0x69, 0x6c,
	0x56, 0x65, 0x2e, 0xe7, 0xc3, 0xda, 0x78, 0xc1, 0xcb, 0x84, 0xad, 0xc6, 0x07, 0x53, 0x98, 0x34,
	0xf4, 0x7a, 0x9c, 0xb9, 0xfc, 0x32, 0x58, 0x67, 0x09, 0xed, 0x91, 0x17, 0xff, 0xcb, 0x54, 0x60,
	0x25, 0x68, 0x2f, 0x52, 0x60, 0x6b, 0x98, 0xdd, 0xbe, 0x9f, 0xbd, 0x18, 0x11, 0xb4, 0x4b, 0x3a,
	0x61, 0x75, 0x40, 0x86, 0x36, 0xc2, 0x29, 0x22, 0x29, 0x40, 0xb1, 0xf5, 0x58, 0x74, 0x6c, 0x1d,
	0x37, 0x0e, 0x7d, 0x43, 0x56, 0xda, 0x2d, 0xaa, 0xb8, 0x50, 0xca, 0x32, 0x8a, 0xbb, 0xb2, 0xdc,
	0xac, 0x52, 0xf5, 0x45, 0x29, 0x4b, 0x0f, 0xc8, 0x66, 0x43, 0x25, 0x42, 0xe6, 0xa0, 0x15, 0x77,
	0x60, 0x6f, 0xc1, 0xb2, 0x0e, 0xc2, 0x9d, 0x68, 0xf6, 0xa2, 0x77, 0x8d, 0x56, 0x58, 0x41, 0xfc,
	0xb6, 0x7c, 0x58, 0x43, 0x0d, 0x6c, 0x03, 0x5b, 0x20, 0x28, 0x5d, 0x06, 0x85, 0xbe, 0x22, 0xcb,
	0x49, 0xa6, 0x85, 0xbd, 0x6b, 0x88, 0x4d, 0x24, 0x96, 0xa2, 0x86, 0x48, 0xef, 0xe0, 0xd7, 0x7e,
	0x9a, 0xf9, 0x9b, 0x3a, 0xd9, 0x93, 0xa6, 0xec, 0x96, 0x99, 0xcc, 0xa1, 0x3a, 0x3a, 0xec, 0x1a,
	0x37, 0xac, 0xc1, 0xde, 0xbd, 0xc7, 0xbf, 0x44, 0x52, 0x0f, 0xba, 0x55, 0x9e, 0x7e, 0x6a, 0xfe,
	0x15, 0xc9, 0x1c, 0xaa, 0x1f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x8b, 0xd6, 0xc2, 0x59,
	0x04, 0x00, 0x00,
}
