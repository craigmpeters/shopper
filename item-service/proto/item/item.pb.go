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
	Id          int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Price       float64 `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
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

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x35, 0x5d, 0x5b, 0xdd, 0x2d, 0x08, 0x5e, 0x14, 0xcb, 0x7c, 0x29, 0x79, 0x2a, 0x88, 0x11,
	0xe6, 0xf3, 0x40, 0x19, 0x32, 0x7c, 0x8d, 0xbf, 0x60, 0x6b, 0xaf, 0x12, 0x58, 0x9b, 0x7a, 0x13,
	0xf7, 0x47, 0xfd, 0x43, 0x92, 0x4c, 0x41, 0xdc, 0x54, 0xf0, 0x25, 0xb9, 0x27, 0xe7, 0x9c, 0x7b,
	0xf2, 0x05, 0x67, 0x03, 0x5b, 0x6f, 0xaf, 0x8d, 0xa7, 0x2e, 0x0e, 0x2a, 0x62, 0x3c, 0x79, 0xb6,
	0xaa, 0x33, 0x0d, 0x5b, 0xe5, 0x78, 0xa3, 0x02, 0x21, 0xc7, 0x70, 0xa8, 0xe9, 0xe5, 0x95, 0x9c,
	0x97, 0x4f, 0x70, 0xa4, 0xc9, 0x0d, 0xb6, 0x77, 0x84, 0x97, 0x90, 0x06, 0xba, 0x14, 0x95, 0xa8,
	0x8b, 0xe9, 0xb9, 0xda, 0x31, 0xaa, 0x07, 0x4f, 0x9d, 0x8e, 0x22, 0xbc, 0x82, 0x2c, 0xcc, 0xae,
	0x4c, 0xaa, 0xd1, 0x6f, 0xea, 0xad, 0x4a, 0xae, 0x20, 0x0d, 0x10, 0x8f, 0x21, 0x31, 0x6d, 0x4c,
	0xc8, 0x74, 0x62, 0x5a, 0x44, 0x48, 0xfb, 0x65, 0x47, 0x65, 0x52, 0x89, 0x7a, 0xac, 0x63, 0x8d,
	0xa7, 0x90, 0x0d, 0x6c, 0x1a, 0x2a, 0x47, 0x95, 0xa8, 0x85, 0xde, 0x02, 0xac, 0xa0, 0x68, 0xc9,
	0x35, 0x6c, 0x06, 0x6f, 0x6c, 0x5f, 0xa6, 0xd1, 0xf0, 0x75, 0x49, 0xce, 0x20, 0xbb, 0x67, 0xb6,
	0x1c, 0x9a, 0x36, 0xb6, 0xa5, 0x8f, 0x98, 0x58, 0x7f, 0xb7, 0x27, 0x3b, 0xf6, 0xe9, 0x9b, 0x80,
	0x22, 0x6c, 0xf6, 0x91, 0x78, 0x13, 0x02, 0x6f, 0x21, 0x9f, 0x33, 0x2d, 0x3d, 0xe1, 0x4f, 0x87,
	0x9b, 0x5c, 0xec, 0x21, 0x3e, 0xaf, 0x53, 0x1e, 0xe0, 0x0c, 0x46, 0x0b, 0xf2, 0xff, 0xb6, 0xcf,
	0x21, 0x5f, 0x90, 0xbf, 0x5b, 0xaf, 0x71, 0xb2, 0x57, 0x18, 0x5f, 0xf0, 0x8f, 0x26, 0xab, 0x3c,
	0xfe, 0x82, 0x9b, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x90, 0x2e, 0xe6, 0x1e, 0x02, 0x00,
	0x00,
}
