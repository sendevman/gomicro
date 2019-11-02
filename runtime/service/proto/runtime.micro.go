// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/go-micro/runtime/service/proto/runtime.proto

package go_micro_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Runtime service

type RuntimeService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type runtimeService struct {
	c    client.Client
	name string
}

func NewRuntimeService(name string, c client.Client) RuntimeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.runtime"
	}
	return &runtimeService{
		c:    c,
		name: name,
	}
}

func (c *runtimeService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) error {
	type runtime interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Runtime struct {
		runtime
	}
	h := &runtimeHandler{hdlr}
	return s.Handle(s.NewHandler(&Runtime{h}, opts...))
}

type runtimeHandler struct {
	RuntimeHandler
}

func (h *runtimeHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.RuntimeHandler.Create(ctx, in, out)
}

func (h *runtimeHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.RuntimeHandler.Delete(ctx, in, out)
}

func (h *runtimeHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.RuntimeHandler.Update(ctx, in, out)
}

func (h *runtimeHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.RuntimeHandler.List(ctx, in, out)
}
