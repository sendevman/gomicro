// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug.proto

package debug

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

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	// default: ok
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{2}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

type StatsResponse struct {
	// unix timestamp
	Started uint64 `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
	// in seconds
	Uptime uint64 `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// in bytes
	Memory uint64 `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	// num threads
	Threads uint64 `protobuf:"varint,4,opt,name=threads,proto3" json:"threads,omitempty"`
	// total gc in nanoseconds
	Gc                   uint64   `protobuf:"varint,5,opt,name=gc,proto3" json:"gc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{3}
}

func (m *StatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsResponse.Unmarshal(m, b)
}
func (m *StatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsResponse.Marshal(b, m, deterministic)
}
func (m *StatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsResponse.Merge(m, src)
}
func (m *StatsResponse) XXX_Size() int {
	return xxx_messageInfo_StatsResponse.Size(m)
}
func (m *StatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsResponse proto.InternalMessageInfo

func (m *StatsResponse) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *StatsResponse) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *StatsResponse) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *StatsResponse) GetThreads() uint64 {
	if m != nil {
		return m.Threads
	}
	return 0
}

func (m *StatsResponse) GetGc() uint64 {
	if m != nil {
		return m.Gc
	}
	return 0
}

// LogRequest requests service logs
type LogRequest struct {
	// count of records to request
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// relative time in seconds
	// before the current time
	// from which to show logs
	Since int64 `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
	// stream records continuously
	Stream               bool     `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{4}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LogRequest) GetSince() int64 {
	if m != nil {
		return m.Since
	}
	return 0
}

func (m *LogRequest) GetStream() bool {
	if m != nil {
		return m.Stream
	}
	return false
}

// Record is service log record
type Record struct {
	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record value
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// record metadata
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{5}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "HealthResponse")
	proto.RegisterType((*StatsRequest)(nil), "StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "StatsResponse")
	proto.RegisterType((*LogRequest)(nil), "LogRequest")
	proto.RegisterType((*Record)(nil), "Record")
	proto.RegisterMapType((map[string]string)(nil), "Record.MetadataEntry")
}

func init() { proto.RegisterFile("debug.proto", fileDescriptor_8d9d361be58531fb) }

var fileDescriptor_8d9d361be58531fb = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x1c, 0xc5, 0x63, 0x3b, 0x76, 0x92, 0x7f, 0x66, 0x67, 0x88, 0x6d, 0x18, 0xb3, 0x43, 0xd0, 0xc9,
	0x30, 0x10, 0x5b, 0x76, 0x19, 0xdb, 0x75, 0x85, 0x1e, 0x52, 0x28, 0xea, 0x27, 0x50, 0x6c, 0xe1,
	0x84, 0xc6, 0x96, 0x6b, 0xfd, 0x5d, 0xc8, 0xad, 0xd0, 0xaf, 0xd3, 0x0f, 0x59, 0x64, 0x29, 0x4d,
	0x0d, 0xbd, 0xf9, 0xf7, 0xe4, 0xff, 0x7b, 0x92, 0x9e, 0x60, 0x59, 0xca, 0x5d, 0x5f, 0xb1, 0xb6,
	0x53, 0xa8, 0xe8, 0x0a, 0xe2, 0x6b, 0x29, 0x8e, 0xb8, 0xe7, 0xf2, 0xa1, 0x97, 0x1a, 0x69, 0x0e,
	0xc9, 0x59, 0xd0, 0xad, 0x6a, 0xb4, 0x24, 0xdf, 0x20, 0xd2, 0x28, 0xb0, 0xd7, 0xa9, 0xb7, 0xf6,
	0xf2, 0x05, 0x77, 0x44, 0x13, 0xf8, 0x74, 0x87, 0x02, 0xf5, 0x79, 0xf2, 0xd9, 0x83, 0xd8, 0x09,
	0x6e, 0x32, 0x85, 0x99, 0x46, 0xd1, 0xa1, 0x2c, 0x87, 0xd1, 0x29, 0x3f, 0xa3, 0xf1, 0xec, 0x5b,
	0x3c, 0xd4, 0x32, 0xf5, 0x87, 0x05, 0x47, 0x46, 0xaf, 0x65, 0xad, 0xba, 0x53, 0x1a, 0x58, 0xdd,
	0x92, 0x71, 0xc2, 0x7d, 0x27, 0x45, 0xa9, 0xd3, 0xa9, 0x75, 0x72, 0x48, 0x12, 0xf0, 0xab, 0x22,
	0x0d, 0x07, 0xd1, 0xaf, 0x0a, 0x7a, 0x0b, 0xb0, 0x55, 0x95, 0xdb, 0x13, 0xf9, 0x02, 0x61, 0xa1,
	0xfa, 0x06, 0x87, 0xfc, 0x80, 0x5b, 0x30, 0xaa, 0x3e, 0x34, 0x85, 0x0d, 0x0f, 0xb8, 0x05, 0x7b,
	0xce, 0x4e, 0x8a, 0x7a, 0xc8, 0x9e, 0x73, 0x47, 0xf4, 0xc5, 0x83, 0x88, 0xcb, 0x42, 0x75, 0x25,
	0xf9, 0x0e, 0x0b, 0xb3, 0x4d, 0x8d, 0xa2, 0x6e, 0x9d, 0xe5, 0x45, 0x30, 0xb6, 0x8f, 0xe2, 0xd8,
	0x5b, 0xdb, 0x05, 0xb7, 0x40, 0x7e, 0xc1, 0xbc, 0x96, 0x28, 0x4a, 0x81, 0x22, 0x0d, 0xd6, 0x41,
	0xbe, 0xdc, 0x7c, 0x65, 0xd6, 0x8e, 0xdd, 0x38, 0xfd, 0xaa, 0xc1, 0xee, 0xc4, 0xdf, 0x7e, 0xcb,
	0xfe, 0x41, 0x3c, 0x5a, 0x22, 0x9f, 0x21, 0xb8, 0x97, 0x27, 0x77, 0xff, 0xe6, 0xf3, 0xe3, 0xac,
	0xbf, 0xfe, 0x1f, 0x6f, 0xf3, 0xe4, 0x41, 0xf8, 0xdf, 0x34, 0x4c, 0x7e, 0x40, 0x64, 0xab, 0x24,
	0x09, 0x1b, 0x95, 0x9c, 0xad, 0xd8, 0xb8, 0x63, 0x3a, 0x21, 0x39, 0x84, 0x43, 0x79, 0x24, 0x66,
	0xef, 0x5b, 0xcd, 0x12, 0x36, 0xea, 0x94, 0x4e, 0xc8, 0x1a, 0xa6, 0x5b, 0x55, 0x69, 0xb2, 0x64,
	0x97, 0x8b, 0xce, 0x66, 0xee, 0x4c, 0x74, 0xf2, 0xd3, 0xdb, 0x45, 0xc3, 0xdb, 0xfa, 0xfd, 0x1a,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x2d, 0x15, 0xdb, 0x6a, 0x02, 0x00, 0x00,
}
