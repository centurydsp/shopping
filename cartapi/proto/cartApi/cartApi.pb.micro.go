// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cartApi/cartApi.proto

package cartApi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CartApi service

func NewCartApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CartApi service

type CartApiService interface {
	FindAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type cartApiService struct {
	c    client.Client
	name string
}

func NewCartApiService(name string, c client.Client) CartApiService {
	return &cartApiService{
		c:    c,
		name: name,
	}
}

func (c *cartApiService) FindAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CartApi.FindAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CartApi service

type CartApiHandler interface {
	FindAll(context.Context, *Request, *Response) error
}

func RegisterCartApiHandler(s server.Server, hdlr CartApiHandler, opts ...server.HandlerOption) error {
	type cartApi interface {
		FindAll(ctx context.Context, in *Request, out *Response) error
	}
	type CartApi struct {
		cartApi
	}
	h := &cartApiHandler{hdlr}
	return s.Handle(s.NewHandler(&CartApi{h}, opts...))
}

type cartApiHandler struct {
	CartApiHandler
}

func (h *cartApiHandler) FindAll(ctx context.Context, in *Request, out *Response) error {
	return h.CartApiHandler.FindAll(ctx, in, out)
}
