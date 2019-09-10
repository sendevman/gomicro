// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/router/proto"
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

// Empty request
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse is returned by ListNodes
type ListResponse struct {
	// network nodes
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// PeerRequest requests list of peers
type PeerRequest struct {
	// node id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerRequest) Reset()         { *m = PeerRequest{} }
func (m *PeerRequest) String() string { return proto.CompactTextString(m) }
func (*PeerRequest) ProtoMessage()    {}
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *PeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerRequest.Unmarshal(m, b)
}
func (m *PeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerRequest.Marshal(b, m, deterministic)
}
func (m *PeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerRequest.Merge(m, src)
}
func (m *PeerRequest) XXX_Size() int {
	return xxx_messageInfo_PeerRequest.Size(m)
}
func (m *PeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerRequest proto.InternalMessageInfo

func (m *PeerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// PeerResponse is returned by ListPeers
type PeerResponse struct {
	Peers                *Peers   `protobuf:"bytes,1,opt,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerResponse.Unmarshal(m, b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
}
func (m *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(m, src)
}
func (m *PeerResponse) XXX_Size() int {
	return xxx_messageInfo_PeerResponse.Size(m)
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetPeers() *Peers {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Peers are node peers
type Peers struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// node peers
	Peers                []*Node  `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (m *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(m, src)
}
func (m *Peers) XXX_Size() int {
	return xxx_messageInfo_Peers.Size(m)
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Peers) GetPeers() []*Node {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Node is network node
type Node struct {
	// node id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Solicit is sent when soliciting routes from the network nodes
type Solicit struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Solicit) Reset()         { *m = Solicit{} }
func (m *Solicit) String() string { return proto.CompactTextString(m) }
func (*Solicit) ProtoMessage()    {}
func (*Solicit) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{8}
}

func (m *Solicit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solicit.Unmarshal(m, b)
}
func (m *Solicit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solicit.Marshal(b, m, deterministic)
}
func (m *Solicit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solicit.Merge(m, src)
}
func (m *Solicit) XXX_Size() int {
	return xxx_messageInfo_Solicit.Size(m)
}
func (m *Solicit) XXX_DiscardUnknown() {
	xxx_messageInfo_Solicit.DiscardUnknown(m)
}

var xxx_messageInfo_Solicit proto.InternalMessageInfo

func (m *Solicit) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Peer is used to advertise node peers
type Peer struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// node peers
	Peers                []*Peer  `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{9}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Peer) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.network.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.network.ListResponse")
	proto.RegisterType((*PeerRequest)(nil), "go.micro.network.PeerRequest")
	proto.RegisterType((*PeerResponse)(nil), "go.micro.network.PeerResponse")
	proto.RegisterType((*Peers)(nil), "go.micro.network.Peers")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Solicit)(nil), "go.micro.network.Solicit")
	proto.RegisterType((*Peer)(nil), "go.micro.network.Peer")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x4f, 0xc2, 0x30,
	0x10, 0xc7, 0x19, 0x32, 0x17, 0x6e, 0x60, 0x4c, 0x1f, 0x74, 0x21, 0x99, 0x21, 0x7d, 0x22, 0x46,
	0x86, 0x81, 0xf0, 0xa6, 0x4f, 0x3c, 0xf8, 0x42, 0x88, 0x99, 0x5f, 0x40, 0xd8, 0x2e, 0xd8, 0x08,
	0x3b, 0x6c, 0x4b, 0xfc, 0xd6, 0x7e, 0x06, 0xd3, 0x76, 0x44, 0x84, 0xcc, 0x84, 0xc4, 0xb7, 0xde,
	0xdd, 0xff, 0x7e, 0x77, 0x6d, 0xef, 0xa0, 0x5d, 0xa0, 0xfe, 0x24, 0xf9, 0x9e, 0x6c, 0x24, 0x69,
	0x62, 0x97, 0x4b, 0x4a, 0xd6, 0x22, 0x93, 0x94, 0x94, 0xfe, 0xce, 0x68, 0x29, 0xf4, 0xdb, 0x76,
	0x91, 0x64, 0xb4, 0x1e, 0xd8, 0xc8, 0x60, 0x49, 0x7d, 0x77, 0x90, 0xb4, 0xd5, 0x28, 0x07, 0x36,
	0xb3, 0x34, 0x1c, 0x86, 0xb7, 0x21, 0x9c, 0x0a, 0xa5, 0x53, 0xfc, 0xd8, 0xa2, 0xd2, 0xfc, 0x01,
	0x5a, 0xce, 0x54, 0x1b, 0x2a, 0x14, 0xb2, 0x3b, 0xf0, 0x0b, 0xca, 0x51, 0x45, 0x5e, 0xf7, 0xac,
	0x17, 0x0e, 0xaf, 0x92, 0xc3, 0xaa, 0xc9, 0x8c, 0x72, 0x4c, 0x9d, 0x88, 0xc7, 0x10, 0x3e, 0x23,
	0xca, 0x12, 0xc6, 0x2e, 0xa0, 0x2e, 0xf2, 0xc8, 0xeb, 0x7a, 0xbd, 0x66, 0x5a, 0x17, 0x39, 0x7f,
	0x84, 0x96, 0x0b, 0x97, 0xf0, 0x3e, 0xf8, 0x1b, 0x44, 0xa9, 0xac, 0x24, 0x1c, 0x5e, 0x1f, 0xc3,
	0x8d, 0x5c, 0xa5, 0x4e, 0xc5, 0xe7, 0xe0, 0x5b, 0x9b, 0xdd, 0x42, 0xc3, 0xd4, 0x2b, 0xd3, 0xaa,
	0x7a, 0xb2, 0x1a, 0x73, 0x01, 0x57, 0xa3, 0xfe, 0xf7, 0x05, 0x5c, 0x89, 0x7b, 0x68, 0x18, 0xf3,
	0xb0, 0x73, 0x16, 0x41, 0x30, 0xcf, 0x73, 0x89, 0xca, 0x70, 0x8c, 0x73, 0x67, 0xf2, 0x31, 0x04,
	0x13, 0x2a, 0x0a, 0xcc, 0xf4, 0x29, 0x6d, 0xf1, 0x11, 0xf8, 0x93, 0x15, 0x29, 0x3c, 0x29, 0x69,
	0x0c, 0xc1, 0x0b, 0xad, 0x44, 0x26, 0x4e, 0xab, 0xf5, 0x0a, 0x0d, 0xf3, 0x6e, 0xff, 0xfc, 0x6c,
	0xf6, 0x27, 0x9d, 0x68, 0xf8, 0xe5, 0x41, 0x30, 0x73, 0x7e, 0x36, 0x85, 0xa6, 0x99, 0x20, 0xc3,
	0x52, 0x2c, 0x3e, 0xce, 0xdb, 0x9b, 0xb6, 0xce, 0x4d, 0x55, 0xd8, 0x0d, 0x08, 0xaf, 0xed, 0x68,
	0xee, 0xdf, 0xe3, 0x8a, 0x2e, 0xaa, 0x69, 0xfb, 0xe3, 0xc6, 0x6b, 0xec, 0x09, 0xc0, 0xf2, 0xcd,
	0x02, 0x28, 0x16, 0xfd, 0xe8, 0xcb, 0x95, 0xd8, 0x91, 0xe2, 0xa3, 0xc8, 0xef, 0xb6, 0x16, 0xe7,
	0x76, 0x79, 0x46, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xb8, 0x9b, 0x3f, 0x94, 0x03, 0x00,
	0x00,
}
