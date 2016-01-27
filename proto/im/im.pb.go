// Code generated by protoc-gen-go.
// source: github.com/micro/slack-srv/proto/im/im.proto
// DO NOT EDIT!

/*
Package go_micro_srv_slack_im is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/im/im.proto

It has these top-level messages:
	Im
	Channel
	CloseRequest
	CloseResponse
	HistoryRequest
	HistoryResponse
	ListRequest
	ListResponse
	MarkRequest
	MarkResponse
	OpenRequest
	OpenResponse
*/
package go_micro_srv_slack_im

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_srv_slack "github.com/micro/slack-srv/proto"

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
const _ = proto.ProtoPackageIsVersion1

type Im struct {
	Id            string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	IsIm          bool   `protobuf:"varint,2,opt,name=is_im" json:"is_im,omitempty"`
	User          string `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Created       int64  `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	IsUserDeleted bool   `protobuf:"varint,5,opt,name=is_user_deleted" json:"is_user_deleted,omitempty"`
}

func (m *Im) Reset()                    { *m = Im{} }
func (m *Im) String() string            { return proto.CompactTextString(m) }
func (*Im) ProtoMessage()               {}
func (*Im) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Channel struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CloseRequest struct {
	Channel string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
}

func (m *CloseRequest) Reset()                    { *m = CloseRequest{} }
func (m *CloseRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()               {}
func (*CloseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CloseResponse struct {
	Ok            bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error         string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	NoOp          bool   `protobuf:"varint,3,opt,name=no_op" json:"no_op,omitempty"`
	AlreadyClosed bool   `protobuf:"varint,4,opt,name=already_closed" json:"already_closed,omitempty"`
}

func (m *CloseResponse) Reset()                    { *m = CloseResponse{} }
func (m *CloseResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()               {}
func (*CloseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type HistoryRequest struct {
	Channel   string  `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	Latest    float64 `protobuf:"fixed64,2,opt,name=latest" json:"latest,omitempty"`
	Oldest    float64 `protobuf:"fixed64,3,opt,name=oldest" json:"oldest,omitempty"`
	Inclusive int64   `protobuf:"varint,4,opt,name=inclusive" json:"inclusive,omitempty"`
	Count     int64   `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
	Unreads   int64   `protobuf:"varint,6,opt,name=unreads" json:"unreads,omitempty"`
}

func (m *HistoryRequest) Reset()                    { *m = HistoryRequest{} }
func (m *HistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*HistoryRequest) ProtoMessage()               {}
func (*HistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type HistoryResponse struct {
	Ok       bool                          `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error    string                        `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Latest   string                        `protobuf:"bytes,3,opt,name=latest" json:"latest,omitempty"`
	Messages []*go_micro_srv_slack.Message `protobuf:"bytes,4,rep,name=messages" json:"messages,omitempty"`
	HasMore  bool                          `protobuf:"varint,5,opt,name=has_more" json:"has_more,omitempty"`
}

func (m *HistoryResponse) Reset()                    { *m = HistoryResponse{} }
func (m *HistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*HistoryResponse) ProtoMessage()               {}
func (*HistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HistoryResponse) GetMessages() []*go_micro_srv_slack.Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Ims   []*Im  `protobuf:"bytes,3,rep,name=ims" json:"ims,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListResponse) GetIms() []*Im {
	if m != nil {
		return m.Ims
	}
	return nil
}

type MarkRequest struct {
	Channel string  `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	Ts      float64 `protobuf:"fixed64,2,opt,name=ts" json:"ts,omitempty"`
}

func (m *MarkRequest) Reset()                    { *m = MarkRequest{} }
func (m *MarkRequest) String() string            { return proto.CompactTextString(m) }
func (*MarkRequest) ProtoMessage()               {}
func (*MarkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type MarkResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *MarkResponse) Reset()                    { *m = MarkResponse{} }
func (m *MarkResponse) String() string            { return proto.CompactTextString(m) }
func (*MarkResponse) ProtoMessage()               {}
func (*MarkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type OpenRequest struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *OpenRequest) Reset()                    { *m = OpenRequest{} }
func (m *OpenRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()               {}
func (*OpenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type OpenResponse struct {
	Ok          bool     `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error       string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Channel     *Channel `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
	NoOp        bool     `protobuf:"varint,4,opt,name=no_op" json:"no_op,omitempty"`
	AlreadyOpen bool     `protobuf:"varint,5,opt,name=already_open" json:"already_open,omitempty"`
}

func (m *OpenResponse) Reset()                    { *m = OpenResponse{} }
func (m *OpenResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()               {}
func (*OpenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *OpenResponse) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func init() {
	proto.RegisterType((*Im)(nil), "go.micro.srv.slack.im.Im")
	proto.RegisterType((*Channel)(nil), "go.micro.srv.slack.im.Channel")
	proto.RegisterType((*CloseRequest)(nil), "go.micro.srv.slack.im.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "go.micro.srv.slack.im.CloseResponse")
	proto.RegisterType((*HistoryRequest)(nil), "go.micro.srv.slack.im.HistoryRequest")
	proto.RegisterType((*HistoryResponse)(nil), "go.micro.srv.slack.im.HistoryResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.slack.im.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.srv.slack.im.ListResponse")
	proto.RegisterType((*MarkRequest)(nil), "go.micro.srv.slack.im.MarkRequest")
	proto.RegisterType((*MarkResponse)(nil), "go.micro.srv.slack.im.MarkResponse")
	proto.RegisterType((*OpenRequest)(nil), "go.micro.srv.slack.im.OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "go.micro.srv.slack.im.OpenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for IM service

type IMClient interface {
	Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error)
	History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error)
	Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error)
}

type iMClient struct {
	c           client.Client
	serviceName string
}

func NewIMClient(serviceName string, c client.Client) IMClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.im"
	}
	return &iMClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *iMClient) Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IM.Close", in)
	out := new(CloseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IM.History", in)
	out := new(HistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IM.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IM.Mark", in)
	out := new(MarkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IM.Open", in)
	out := new(OpenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IM service

type IMHandler interface {
	Close(context.Context, *CloseRequest, *CloseResponse) error
	History(context.Context, *HistoryRequest, *HistoryResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Mark(context.Context, *MarkRequest, *MarkResponse) error
	Open(context.Context, *OpenRequest, *OpenResponse) error
}

func RegisterIMHandler(s server.Server, hdlr IMHandler) {
	s.Handle(s.NewHandler(&IM{hdlr}))
}

type IM struct {
	IMHandler
}

func (h *IM) Close(ctx context.Context, in *CloseRequest, out *CloseResponse) error {
	return h.IMHandler.Close(ctx, in, out)
}

func (h *IM) History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error {
	return h.IMHandler.History(ctx, in, out)
}

func (h *IM) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.IMHandler.List(ctx, in, out)
}

func (h *IM) Mark(ctx context.Context, in *MarkRequest, out *MarkResponse) error {
	return h.IMHandler.Mark(ctx, in, out)
}

func (h *IM) Open(ctx context.Context, in *OpenRequest, out *OpenResponse) error {
	return h.IMHandler.Open(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0x9b, 0x7e, 0x64, 0xb7, 0x69, 0x0b, 0x16, 0x83, 0xd2, 0x49, 0x80, 0x3c, 0x98, 0x06,
	0xa2, 0x89, 0x54, 0x7e, 0xc2, 0x5e, 0x98, 0x44, 0x85, 0x86, 0x90, 0x90, 0x78, 0x89, 0xb2, 0xc4,
	0x6a, 0xad, 0xc6, 0x71, 0xb1, 0x93, 0x49, 0x7b, 0x40, 0xfc, 0x24, 0xfe, 0x22, 0xfe, 0x48, 0xbb,
	0x20, 0xcd, 0x8b, 0xd4, 0x87, 0x5e, 0xdf, 0xe3, 0x73, 0xce, 0xbd, 0x3e, 0x2d, 0x7c, 0x5c, 0xd3,
	0x72, 0x53, 0xdd, 0x84, 0x29, 0x67, 0x11, 0xa3, 0xa9, 0xe0, 0x91, 0xcc, 0x93, 0x74, 0xbb, 0x90,
	0xe2, 0x36, 0xda, 0x09, 0x5e, 0xf2, 0x88, 0x32, 0xf5, 0x09, 0xcd, 0x77, 0x74, 0xb2, 0xe6, 0xa1,
	0x41, 0x85, 0xaa, 0x1f, 0x1a, 0x64, 0x48, 0xd9, 0x7c, 0xd1, 0x4a, 0xa2, 0x3a, 0x8c, 0x17, 0x96,
	0x05, 0xff, 0x80, 0xee, 0x15, 0x43, 0x00, 0x5d, 0x9a, 0xcd, 0x3a, 0x6f, 0x3a, 0x17, 0xc7, 0x68,
	0x0c, 0x7d, 0x2a, 0x63, 0xca, 0x66, 0x5d, 0x55, 0xfa, 0x28, 0x80, 0x5e, 0x25, 0x89, 0x98, 0x79,
	0xa6, 0x39, 0x85, 0x61, 0x2a, 0x48, 0x52, 0x92, 0x6c, 0xd6, 0x53, 0x07, 0x1e, 0x7a, 0x01, 0x53,
	0x85, 0xd6, 0x88, 0x38, 0x23, 0x39, 0xd1, 0x8d, 0xbe, 0xbe, 0x87, 0x4f, 0x60, 0x78, 0xb9, 0x49,
	0x8a, 0x82, 0xe4, 0x4d, 0x76, 0xfc, 0x1a, 0x82, 0xcb, 0x9c, 0x4b, 0xf2, 0x8d, 0xfc, 0xaa, 0x88,
	0x2c, 0x0d, 0xa1, 0x85, 0xd5, 0x80, 0x6b, 0x18, 0xd7, 0x00, 0xb9, 0xe3, 0x85, 0x24, 0xfa, 0x36,
	0xdf, 0x9a, 0xa6, 0xaf, 0xbd, 0x11, 0x21, 0xb8, 0x30, 0xde, 0x8c, 0xd5, 0x82, 0xc7, 0x7c, 0x67,
	0xcc, 0xf9, 0xe8, 0x39, 0x4c, 0x92, 0x5c, 0xb9, 0xcb, 0xee, 0xe2, 0x54, 0x53, 0x58, 0x8f, 0x3e,
	0x2e, 0x61, 0xf2, 0x99, 0xca, 0x92, 0x8b, 0x3b, 0x97, 0x2a, 0x9a, 0xc0, 0x20, 0x57, 0x53, 0xc9,
	0xd2, 0x30, 0x77, 0x74, 0xcd, 0xf3, 0x4c, 0xd7, 0x9e, 0xa9, 0x9f, 0xc2, 0x31, 0x2d, 0xd2, 0xbc,
	0x92, 0xf4, 0x96, 0xd4, 0x93, 0x2b, 0xf1, 0x94, 0x57, 0x45, 0x69, 0xe6, 0xf5, 0x34, 0x65, 0x55,
	0x68, 0x71, 0x39, 0x1b, 0xe8, 0x03, 0xfc, 0x1b, 0xa6, 0x07, 0xd5, 0xf6, 0x51, 0xee, 0x0d, 0xd8,
	0x45, 0x2f, 0xc0, 0x67, 0x44, 0xca, 0x64, 0x4d, 0xa4, 0xd2, 0xf3, 0x2e, 0x46, 0xcb, 0xd3, 0xf0,
	0x81, 0x07, 0x5f, 0x59, 0x0c, 0x7a, 0x02, 0xfe, 0x26, 0x91, 0x31, 0xe3, 0x82, 0xd4, 0xfb, 0x1f,
	0xc3, 0xe8, 0x8b, 0x92, 0xaf, 0x27, 0x56, 0x6b, 0x0d, 0x6c, 0xd9, 0x6e, 0xe5, 0x1c, 0x3c, 0xca,
	0xa4, 0xf2, 0xa1, 0x55, 0x5f, 0x86, 0x0f, 0xc6, 0x2c, 0xbc, 0x62, 0xf8, 0x03, 0x8c, 0x56, 0x89,
	0xd8, 0x3a, 0x77, 0xaa, 0x24, 0x4a, 0x69, 0xf7, 0x89, 0xdf, 0x43, 0x60, 0xb1, 0xad, 0xf2, 0xf8,
	0x14, 0x46, 0x5f, 0x77, 0xa4, 0xd8, 0xd3, 0xee, 0xf3, 0x67, 0xd3, 0xf1, 0x07, 0x02, 0xdb, 0x6c,
	0x1f, 0x23, 0xba, 0xf7, 0xa3, 0x57, 0x3a, 0x5a, 0xbe, 0x72, 0x8c, 0xb2, 0x8f, 0xe9, 0x21, 0x4d,
	0x26, 0x35, 0xe8, 0x19, 0x04, 0xfb, 0x34, 0x71, 0x25, 0x69, 0xd7, 0xba, 0xfc, 0xeb, 0xa9, 0x1f,
	0xcc, 0x0a, 0x7d, 0x87, 0xbe, 0x49, 0x29, 0x3a, 0x73, 0x91, 0x36, 0x42, 0x3e, 0x7f, 0xfb, 0x38,
	0xc8, 0xce, 0x82, 0x8f, 0xd0, 0x4f, 0x18, 0xd6, 0x91, 0x41, 0xef, 0x1c, 0x57, 0xfe, 0x0f, 0xf2,
	0xfc, 0xbc, 0x0d, 0x76, 0xe0, 0xbe, 0x86, 0x9e, 0x0e, 0x00, 0xc2, 0x8e, 0x1b, 0x8d, 0xb0, 0xcc,
	0xcf, 0x1e, 0xc5, 0x34, 0x29, 0xf5, 0xa3, 0x3a, 0x29, 0x1b, 0xe9, 0x70, 0x52, 0x36, 0x53, 0x61,
	0x29, 0xf5, 0xfb, 0x3a, 0x29, 0x1b, 0xc9, 0x70, 0x52, 0x36, 0x03, 0x82, 0x8f, 0x6e, 0x06, 0xe6,
	0x8f, 0xee, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xea, 0xab, 0xc4, 0x5e, 0x05, 0x00,
	0x00,
}
