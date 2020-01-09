// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/router/service/proto"
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

// Query is passed in a LookupRequest
type Query struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Router               string   `protobuf:"bytes,4,opt,name=router,proto3" json:"router,omitempty"`
	Network              string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Query) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Query) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Query) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Query) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type ConnectRequest struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ConnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

// PeerRequest requests list of peers
type NodesRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

func (m *NodesRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

// PeerResponse is returned by ListPeers
type NodesResponse struct {
	// return peer topology
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GraphRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphRequest) Reset()         { *m = GraphRequest{} }
func (m *GraphRequest) String() string { return proto.CompactTextString(m) }
func (*GraphRequest) ProtoMessage()    {}
func (*GraphRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *GraphRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphRequest.Unmarshal(m, b)
}
func (m *GraphRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphRequest.Marshal(b, m, deterministic)
}
func (m *GraphRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphRequest.Merge(m, src)
}
func (m *GraphRequest) XXX_Size() int {
	return xxx_messageInfo_GraphRequest.Size(m)
}
func (m *GraphRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GraphRequest proto.InternalMessageInfo

func (m *GraphRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type GraphResponse struct {
	Root                 *Peer    `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphResponse) Reset()         { *m = GraphResponse{} }
func (m *GraphResponse) String() string { return proto.CompactTextString(m) }
func (*GraphResponse) ProtoMessage()    {}
func (*GraphResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *GraphResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphResponse.Unmarshal(m, b)
}
func (m *GraphResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphResponse.Marshal(b, m, deterministic)
}
func (m *GraphResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphResponse.Merge(m, src)
}
func (m *GraphResponse) XXX_Size() int {
	return xxx_messageInfo_GraphResponse.Size(m)
}
func (m *GraphResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GraphResponse proto.InternalMessageInfo

func (m *GraphResponse) GetRoot() *Peer {
	if m != nil {
		return m.Root
	}
	return nil
}

type RoutesRequest struct {
	// filter based on
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesRequest) Reset()         { *m = RoutesRequest{} }
func (m *RoutesRequest) String() string { return proto.CompactTextString(m) }
func (*RoutesRequest) ProtoMessage()    {}
func (*RoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *RoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesRequest.Unmarshal(m, b)
}
func (m *RoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesRequest.Marshal(b, m, deterministic)
}
func (m *RoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesRequest.Merge(m, src)
}
func (m *RoutesRequest) XXX_Size() int {
	return xxx_messageInfo_RoutesRequest.Size(m)
}
func (m *RoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesRequest proto.InternalMessageInfo

func (m *RoutesRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type RoutesResponse struct {
	Routes               []*proto1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoutesResponse) Reset()         { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()    {}
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{8}
}

func (m *RoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesResponse.Unmarshal(m, b)
}
func (m *RoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesResponse.Marshal(b, m, deterministic)
}
func (m *RoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesResponse.Merge(m, src)
}
func (m *RoutesResponse) XXX_Size() int {
	return xxx_messageInfo_RoutesResponse.Size(m)
}
func (m *RoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesResponse proto.InternalMessageInfo

func (m *RoutesResponse) GetRoutes() []*proto1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesRequest) Reset()         { *m = ServicesRequest{} }
func (m *ServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ServicesRequest) ProtoMessage()    {}
func (*ServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{9}
}

func (m *ServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesRequest.Unmarshal(m, b)
}
func (m *ServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesRequest.Marshal(b, m, deterministic)
}
func (m *ServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesRequest.Merge(m, src)
}
func (m *ServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ServicesRequest.Size(m)
}
func (m *ServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesRequest proto.InternalMessageInfo

type ServicesResponse struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesResponse) Reset()         { *m = ServicesResponse{} }
func (m *ServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ServicesResponse) ProtoMessage()    {}
func (*ServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{10}
}

func (m *ServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesResponse.Unmarshal(m, b)
}
func (m *ServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesResponse.Marshal(b, m, deterministic)
}
func (m *ServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesResponse.Merge(m, src)
}
func (m *ServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ServicesResponse.Size(m)
}
func (m *ServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesResponse proto.InternalMessageInfo

func (m *ServicesResponse) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

// Node is network node
type Node struct {
	// node id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// the network
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// associated metadata
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{11}
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

func (m *Node) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
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
	return fileDescriptor_8571034d60397816, []int{12}
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
	return fileDescriptor_8571034d60397816, []int{13}
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
	return fileDescriptor_8571034d60397816, []int{14}
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

// Sync is network sync message
type Sync struct {
	// node address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// node peers
	Peers []*Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	// node routes
	Routes               []*proto1.Route `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Sync) Reset()         { *m = Sync{} }
func (m *Sync) String() string { return proto.CompactTextString(m) }
func (*Sync) ProtoMessage()    {}
func (*Sync) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{15}
}

func (m *Sync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sync.Unmarshal(m, b)
}
func (m *Sync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sync.Marshal(b, m, deterministic)
}
func (m *Sync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sync.Merge(m, src)
}
func (m *Sync) XXX_Size() int {
	return xxx_messageInfo_Sync.Size(m)
}
func (m *Sync) XXX_DiscardUnknown() {
	xxx_messageInfo_Sync.DiscardUnknown(m)
}

var xxx_messageInfo_Sync proto.InternalMessageInfo

func (m *Sync) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Sync) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Sync) GetRoutes() []*proto1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "go.micro.network.Query")
	proto.RegisterType((*ConnectRequest)(nil), "go.micro.network.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "go.micro.network.ConnectResponse")
	proto.RegisterType((*NodesRequest)(nil), "go.micro.network.NodesRequest")
	proto.RegisterType((*NodesResponse)(nil), "go.micro.network.NodesResponse")
	proto.RegisterType((*GraphRequest)(nil), "go.micro.network.GraphRequest")
	proto.RegisterType((*GraphResponse)(nil), "go.micro.network.GraphResponse")
	proto.RegisterType((*RoutesRequest)(nil), "go.micro.network.RoutesRequest")
	proto.RegisterType((*RoutesResponse)(nil), "go.micro.network.RoutesResponse")
	proto.RegisterType((*ServicesRequest)(nil), "go.micro.network.ServicesRequest")
	proto.RegisterType((*ServicesResponse)(nil), "go.micro.network.ServicesResponse")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.network.Node.MetadataEntry")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Peer)(nil), "go.micro.network.Peer")
	proto.RegisterType((*Sync)(nil), "go.micro.network.Sync")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x8d, 0x2c, 0x29, 0x76, 0xa6, 0x91, 0xeb, 0x2e, 0x25, 0x15, 0x3a, 0xb4, 0xee, 0xe2, 0x43,
	0x28, 0x8d, 0x0c, 0x09, 0x85, 0x52, 0xd3, 0x10, 0x08, 0xa5, 0x50, 0x48, 0x48, 0xe5, 0x1f, 0xa8,
	0x62, 0x2d, 0xb6, 0x49, 0xac, 0x75, 0x56, 0xeb, 0x04, 0x5f, 0x7a, 0xed, 0x27, 0xf4, 0x9b, 0xfa,
	0x57, 0x65, 0x77, 0x47, 0x8e, 0x14, 0xcb, 0x22, 0xbe, 0x79, 0x34, 0x6f, 0xde, 0xec, 0xce, 0x7b,
	0xb3, 0x06, 0x2f, 0x65, 0xf2, 0x81, 0x8b, 0x9b, 0x70, 0x2e, 0xb8, 0xe4, 0xa4, 0x33, 0xe6, 0xe1,
	0x6c, 0x3a, 0x12, 0x3c, 0xc4, 0xef, 0xc1, 0x60, 0x3c, 0x95, 0x93, 0xc5, 0x75, 0x38, 0xe2, 0xb3,
	0xbe, 0xce, 0xf4, 0xc7, 0xfc, 0xc8, 0xfc, 0x10, 0x7c, 0x21, 0x99, 0xe8, 0x67, 0x4c, 0xdc, 0x4f,
	0x47, 0xac, 0xaf, 0x19, 0xf0, 0xa3, 0xa1, 0xa3, 0x7f, 0x2c, 0x70, 0x7f, 0x2e, 0x98, 0x58, 0x12,
	0x1f, 0x9a, 0x88, 0xf3, 0xad, 0xae, 0x75, 0xb8, 0x17, 0xe5, 0xa1, 0xca, 0xc4, 0x49, 0x22, 0x58,
	0x96, 0xf9, 0x0d, 0x93, 0xc1, 0x50, 0x65, 0xc6, 0xb1, 0x64, 0x0f, 0xf1, 0xd2, 0xb7, 0x4d, 0x06,
	0x43, 0x72, 0x00, 0xbb, 0xa6, 0x8f, 0xef, 0xe8, 0x04, 0x46, 0xaa, 0x02, 0xcf, 0xed, 0xbb, 0xa6,
	0x02, 0x43, 0x7a, 0x0a, 0xed, 0x73, 0x9e, 0xa6, 0x6c, 0x24, 0x23, 0x76, 0xb7, 0x60, 0x99, 0x24,
	0x1f, 0xc1, 0x4d, 0x79, 0xc2, 0x32, 0xdf, 0xea, 0xda, 0x87, 0x2f, 0x8e, 0x0f, 0xc2, 0xa7, 0x57,
	0x0f, 0x2f, 0x79, 0xc2, 0x22, 0x03, 0xa2, 0xaf, 0xe0, 0xe5, 0xaa, 0x3e, 0x9b, 0xf3, 0x34, 0x63,
	0xb4, 0x07, 0xfb, 0x0a, 0x91, 0xe5, 0x84, 0xaf, 0xc1, 0x4d, 0xd8, 0x5c, 0x4e, 0xf4, 0x05, 0xbd,
	0xc8, 0x04, 0xf4, 0x2b, 0x78, 0x88, 0x32, 0x65, 0x5b, 0xf6, 0xed, 0xc1, 0xfe, 0x77, 0x11, 0xcf,
	0x27, 0xf5, 0x4d, 0x06, 0xe0, 0x21, 0x0a, 0x9b, 0x7c, 0x00, 0x47, 0x70, 0x2e, 0x35, 0xaa, 0xb2,
	0xc7, 0x15, 0x63, 0x22, 0xd2, 0x18, 0x7a, 0x0a, 0x5e, 0xa4, 0xc6, 0xb7, 0xba, 0xc8, 0x11, 0xb8,
	0x77, 0x4a, 0x34, 0xac, 0x7e, 0xb3, 0x5e, 0xad, 0x35, 0x8d, 0x0c, 0x8a, 0x9e, 0x41, 0x3b, 0xaf,
	0xc7, 0xee, 0x21, 0xca, 0x53, 0x71, 0x47, 0xb4, 0x87, 0x2e, 0x40, 0xd9, 0xf4, 0x70, 0x87, 0xc6,
	0x0d, 0xf9, 0x19, 0x68, 0x08, 0x9d, 0xc7, 0x4f, 0x48, 0x1b, 0x40, 0x0b, 0x4d, 0x63, 0x88, 0xf7,
	0xa2, 0x55, 0x4c, 0xff, 0x59, 0xe0, 0xa8, 0xb9, 0x91, 0x36, 0x34, 0xa6, 0x09, 0x7a, 0xac, 0x31,
	0x4d, 0xea, 0xed, 0x95, 0x9b, 0xc5, 0x2e, 0x99, 0x85, 0x9c, 0x41, 0x6b, 0xc6, 0x64, 0x9c, 0xc4,
	0x32, 0xf6, 0x1d, 0x7d, 0x83, 0x5e, 0xb5, 0x4a, 0xe1, 0x05, 0xc2, 0xbe, 0xa5, 0x52, 0x2c, 0xa3,
	0x55, 0x55, 0x30, 0x00, 0xaf, 0x94, 0x22, 0x1d, 0xb0, 0x6f, 0xd8, 0x12, 0xcf, 0xa5, 0x7e, 0x2a,
	0x25, 0xef, 0xe3, 0xdb, 0x05, 0xc3, 0x63, 0x99, 0xe0, 0x4b, 0xe3, 0xb3, 0x45, 0x3f, 0x41, 0x13,
	0xbd, 0xa6, 0x74, 0x54, 0x3e, 0xd8, 0xac, 0xa3, 0xf6, 0x8a, 0xc6, 0xd0, 0x13, 0x70, 0xcf, 0x6f,
	0xb9, 0x11, 0xff, 0xd9, 0x45, 0xbf, 0xc0, 0x51, 0x56, 0xd8, 0xa6, 0x46, 0x39, 0x78, 0xce, 0x98,
	0x50, 0x03, 0xb5, 0x6b, 0xdc, 0x65, 0x40, 0xf4, 0x37, 0x38, 0xc3, 0x65, 0x3a, 0x2a, 0x0a, 0x61,
	0x95, 0x85, 0xd8, 0x8a, 0xaf, 0x60, 0x2e, 0xfb, 0x39, 0xe6, 0x3a, 0xfe, 0x6b, 0x43, 0xf3, 0x12,
	0x85, 0xbd, 0x7a, 0x9c, 0x6c, 0x77, 0xbd, 0x4b, 0xf9, 0x81, 0x08, 0xde, 0xd7, 0x20, 0xf0, 0x09,
	0xd8, 0x21, 0x3f, 0xc0, 0xd5, 0x9b, 0x47, 0xde, 0xae, 0xa3, 0x8b, 0x8b, 0x1b, 0xbc, 0xdb, 0x98,
	0x2f, 0x72, 0xe9, 0xa7, 0xa2, 0x8a, 0xab, 0xf8, 0xd2, 0x54, 0x71, 0x95, 0xde, 0x18, 0xba, 0x43,
	0x2e, 0x60, 0xd7, 0x2c, 0x25, 0xa9, 0x00, 0x97, 0xd6, 0x3d, 0xe8, 0x6e, 0x06, 0xac, 0xe8, 0x86,
	0xd0, 0xca, 0xd7, 0x91, 0x54, 0xcc, 0xe5, 0xc9, 0xf6, 0x06, 0xb4, 0x0e, 0x92, 0x93, 0x5e, 0xef,
	0xea, 0x3f, 0x89, 0x93, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x0e, 0x14, 0x60, 0x84, 0x06,
	0x00, 0x00,
}
