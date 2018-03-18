// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/item/item.proto

/*
Package go_micro_srv_item is a generated protocol buffer package.

It is generated from these files:
	proto/item/item.proto

It has these top-level messages:
	Request
	Response
	Item
	Error
*/
package go_micro_srv_item

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Item  *Item   `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Response) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Item struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.item.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.item.Response")
	proto.RegisterType((*Item)(nil), "go.micro.srv.item.Item")
	proto.RegisterType((*Error)(nil), "go.micro.srv.item.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ItemService service

type ItemServiceClient interface {
	Create(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type itemServiceClient struct {
	c           client.Client
	serviceName string
}

func NewItemServiceClient(serviceName string, c client.Client) ItemServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.item"
	}
	return &itemServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *itemServiceClient) Create(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Get(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemService service

type ItemServiceHandler interface {
	Create(context.Context, *Item, *Response) error
	Get(context.Context, *Item, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterItemServiceHandler(s server.Server, hdlr ItemServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ItemService{hdlr}, opts...))
}

type ItemService struct {
	ItemServiceHandler
}

func (h *ItemService) Create(ctx context.Context, in *Item, out *Response) error {
	return h.ItemServiceHandler.Create(ctx, in, out)
}

func (h *ItemService) Get(ctx context.Context, in *Item, out *Response) error {
	return h.ItemServiceHandler.Get(ctx, in, out)
}

func (h *ItemService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.ItemServiceHandler.GetAll(ctx, in, out)
}

func init() { proto.RegisterFile("proto/item/item.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0x7e, 0xb3, 0x5f, 0xaf, 0x9d, 0x05, 0xc1, 0x41, 0x71, 0xa9, 0x97, 0x65, 0x4f, 0x0b, 0x62,
	0x84, 0x7a, 0x2e, 0x54, 0x8a, 0x14, 0xaf, 0xf1, 0x17, 0xd4, 0xec, 0x28, 0x81, 0xee, 0x66, 0x9d,
	0xc4, 0xfe, 0x51, 0xff, 0x90, 0x24, 0xad, 0x20, 0x58, 0x15, 0xbc, 0x24, 0xf3, 0xe4, 0xf9, 0xc8,
	0x30, 0x03, 0x67, 0x23, 0x5b, 0x6f, 0xaf, 0x8d, 0xa7, 0x3e, 0x1e, 0x32, 0x62, 0x3c, 0x79, 0xb6,
	0xb2, 0x37, 0x9a, 0xad, 0x74, 0xbc, 0x95, 0x81, 0x68, 0x26, 0xf0, 0x5f, 0xd1, 0xcb, 0x2b, 0x39,
	0xdf, 0x3c, 0xc1, 0x91, 0x22, 0x37, 0xda, 0xc1, 0x11, 0x5e, 0x42, 0x16, 0xe8, 0x4a, 0xd4, 0xa2,
	0x2d, 0x67, 0xe7, 0xf2, 0x8b, 0x51, 0xde, 0x7b, 0xea, 0x55, 0x14, 0xe1, 0x15, 0xe4, 0xe1, 0x76,
	0x55, 0x52, 0xa7, 0x3f, 0xa9, 0x77, 0xaa, 0x66, 0x01, 0x59, 0x80, 0x78, 0x0c, 0x89, 0xe9, 0xe2,
	0x0f, 0xb9, 0x4a, 0x4c, 0x87, 0x08, 0xd9, 0xb0, 0xee, 0xa9, 0x4a, 0x6a, 0xd1, 0x4e, 0x54, 0xac,
	0xf1, 0x14, 0xf2, 0x91, 0x8d, 0xa6, 0x2a, 0xad, 0x45, 0x2b, 0xd4, 0x0e, 0x34, 0x73, 0xc8, 0xef,
	0x98, 0x2d, 0x07, 0x8b, 0xb6, 0x1d, 0xed, 0x43, 0x62, 0x8d, 0x35, 0x94, 0x1d, 0x39, 0xcd, 0x66,
	0xf4, 0xc6, 0x0e, 0xfb, 0xb4, 0xcf, 0x4f, 0xb3, 0x37, 0x01, 0x65, 0x68, 0xe5, 0x81, 0x78, 0x6b,
	0x34, 0xe1, 0x02, 0x8a, 0x25, 0xd3, 0xda, 0x13, 0x7e, 0xd7, 0xfa, 0xf4, 0xe2, 0x00, 0xf1, 0x31,
	0xac, 0xe6, 0x1f, 0xce, 0x21, 0x5d, 0x91, 0xff, 0xb3, 0x7d, 0x09, 0xc5, 0x8a, 0xfc, 0xed, 0x66,
	0x83, 0xd3, 0x83, 0xc2, 0xb8, 0x9f, 0x5f, 0x42, 0x1e, 0x8b, 0xb8, 0xe3, 0x9b, 0xf7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0xe3, 0x31, 0x92, 0xfc, 0x01, 0x00, 0x00,
}