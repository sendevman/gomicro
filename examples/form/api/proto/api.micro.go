// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/asim/go-micro/v3/api/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Client API for Form service

type FormService interface {
	// regular form
	Submit(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	// multipart form
	Multipart(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type formService struct {
	c    client.Client
	name string
}

func NewFormService(name string, c client.Client) FormService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "form"
	}
	return &formService{
		c:    c,
		name: name,
	}
}

func (c *formService) Submit(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Form.Submit", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) Multipart(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Form.Multipart", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Form service

type FormHandler interface {
	// regular form
	Submit(context.Context, *proto1.Request, *proto1.Response) error
	// multipart form
	Multipart(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterFormHandler(s server.Server, hdlr FormHandler, opts ...server.HandlerOption) error {
	type form interface {
		Submit(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		Multipart(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Form struct {
		form
	}
	h := &formHandler{hdlr}
	return s.Handle(s.NewHandler(&Form{h}, opts...))
}

type formHandler struct {
	FormHandler
}

func (h *formHandler) Submit(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.FormHandler.Submit(ctx, in, out)
}

func (h *formHandler) Multipart(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.FormHandler.Multipart(ctx, in, out)
}
