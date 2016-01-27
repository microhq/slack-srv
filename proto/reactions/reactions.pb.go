// Code generated by protoc-gen-go.
// source: github.com/micro/slack-srv/proto/reactions/reactions.proto
// DO NOT EDIT!

/*
Package go_micro_srv_slack_reactions is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/reactions/reactions.proto

It has these top-level messages:
	Reaction
	Message
	File
	Item
	AddRequest
	AddResponse
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	RemoveRequest
	RemoveResponse
*/
package go_micro_srv_slack_reactions

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

type Reaction struct {
	Name  string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Count int64    `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Users []string `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
}

func (m *Reaction) Reset()                    { *m = Reaction{} }
func (m *Reaction) String() string            { return proto.CompactTextString(m) }
func (*Reaction) ProtoMessage()               {}
func (*Reaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Message   *go_micro_srv_slack.Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Reactions []*Reaction                 `protobuf:"bytes,2,rep,name=reactions" json:"reactions,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetMessage() *go_micro_srv_slack.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Message) GetReactions() []*Reaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

type File struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Created   int64  `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Item struct {
	Type        string      `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Channel     string      `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	File        *File       `protobuf:"bytes,3,opt,name=file" json:"file,omitempty"`
	FileComment *File       `protobuf:"bytes,4,opt,name=file_comment" json:"file_comment,omitempty"`
	Message     *Message    `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Reactions   []*Reaction `protobuf:"bytes,6,rep,name=reactions" json:"reactions,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Item) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Item) GetFileComment() *File {
	if m != nil {
		return m.FileComment
	}
	return nil
}

func (m *Item) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Item) GetReactions() []*Reaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

type AddRequest struct {
	Name        string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	File        string  `protobuf:"bytes,2,opt,name=file" json:"file,omitempty"`
	FileComment string  `protobuf:"bytes,3,opt,name=file_comment" json:"file_comment,omitempty"`
	Channel     string  `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	Timestamp   float64 `protobuf:"fixed64,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AddResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetRequest struct {
	File        string  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	FileComment string  `protobuf:"bytes,2,opt,name=file_comment" json:"file_comment,omitempty"`
	Channel     string  `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
	Timestamp   float64 `protobuf:"fixed64,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Full        bool    `protobuf:"varint,5,opt,name=full" json:"full,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetResponse struct {
	Ok      bool     `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error   string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Type    string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Channel string   `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	Message *Message `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListRequest struct {
	User  string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Full  bool   `protobuf:"varint,2,opt,name=full" json:"full,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Page  int64  `protobuf:"varint,4,opt,name=page" json:"page,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListResponse struct {
	Ok    bool    `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string  `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Items []*Item `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type RemoveRequest struct {
	Name        string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	File        string  `protobuf:"bytes,2,opt,name=file" json:"file,omitempty"`
	FileComment string  `protobuf:"bytes,3,opt,name=file_comment" json:"file_comment,omitempty"`
	Channel     string  `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	Timestamp   float64 `protobuf:"fixed64,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type RemoveResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveResponse) Reset()                    { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()               {}
func (*RemoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Reaction)(nil), "go.micro.srv.slack.reactions.Reaction")
	proto.RegisterType((*Message)(nil), "go.micro.srv.slack.reactions.Message")
	proto.RegisterType((*File)(nil), "go.micro.srv.slack.reactions.File")
	proto.RegisterType((*Item)(nil), "go.micro.srv.slack.reactions.Item")
	proto.RegisterType((*AddRequest)(nil), "go.micro.srv.slack.reactions.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "go.micro.srv.slack.reactions.AddResponse")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.slack.reactions.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.slack.reactions.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.slack.reactions.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.srv.slack.reactions.ListResponse")
	proto.RegisterType((*RemoveRequest)(nil), "go.micro.srv.slack.reactions.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "go.micro.srv.slack.reactions.RemoveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Reactions service

type ReactionsClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error)
}

type reactionsClient struct {
	c           client.Client
	serviceName string
}

func NewReactionsClient(serviceName string, c client.Client) ReactionsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.reactions"
	}
	return &reactionsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *reactionsClient) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Reactions.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionsClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Reactions.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionsClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Reactions.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionsClient) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Reactions.Remove", in)
	out := new(RemoveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reactions service

type ReactionsHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Remove(context.Context, *RemoveRequest, *RemoveResponse) error
}

func RegisterReactionsHandler(s server.Server, hdlr ReactionsHandler) {
	s.Handle(s.NewHandler(&Reactions{hdlr}))
}

type Reactions struct {
	ReactionsHandler
}

func (h *Reactions) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.ReactionsHandler.Add(ctx, in, out)
}

func (h *Reactions) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ReactionsHandler.Get(ctx, in, out)
}

func (h *Reactions) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ReactionsHandler.List(ctx, in, out)
}

func (h *Reactions) Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error {
	return h.ReactionsHandler.Remove(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x1f, 0xf9, 0xf0, 0x38, 0x05, 0xb1, 0xe2, 0x10, 0x05, 0x0e, 0x95, 0x25, 0x50, 0xa0,
	0xad, 0x03, 0x41, 0xaa, 0x80, 0x1b, 0x97, 0x22, 0x24, 0xb8, 0xe4, 0x5c, 0xa9, 0x72, 0x9c, 0x25,
	0xb5, 0x6a, 0x7b, 0x8d, 0x77, 0x5d, 0x89, 0x13, 0x17, 0xee, 0xfc, 0x65, 0x66, 0xc7, 0x71, 0x62,
	0x42, 0xc8, 0x06, 0xa4, 0xde, 0x6c, 0xef, 0xbc, 0x79, 0x6f, 0xde, 0xbc, 0x4d, 0xe0, 0xdd, 0x32,
	0x51, 0xd7, 0xd5, 0x3c, 0x8c, 0x45, 0x36, 0xc9, 0x92, 0xb8, 0x14, 0x13, 0x99, 0x46, 0xf1, 0xcd,
	0x99, 0x2c, 0x6f, 0x27, 0x45, 0x29, 0x94, 0x98, 0x94, 0x3c, 0x8a, 0x55, 0x22, 0x72, 0xb9, 0x79,
	0x0a, 0xe9, 0x84, 0x3d, 0x59, 0x8a, 0x90, 0x30, 0x21, 0x56, 0x87, 0x84, 0x0b, 0xd7, 0x35, 0xa3,
	0x33, 0x63, 0x67, 0x3c, 0xc9, 0x44, 0x5e, 0x37, 0x0b, 0xce, 0xa1, 0x3f, 0x5b, 0x61, 0xd9, 0x00,
	0xdc, 0x3c, 0xca, 0xf8, 0xd0, 0x3a, 0xb6, 0xc6, 0x1e, 0x3b, 0x82, 0x4e, 0x2c, 0xaa, 0x5c, 0x0d,
	0x6d, 0x7c, 0x75, 0xf4, 0x6b, 0x25, 0x79, 0x29, 0x87, 0xce, 0xb1, 0x33, 0xf6, 0x82, 0x12, 0x7a,
	0x9f, 0xb9, 0x94, 0xd1, 0x92, 0xb3, 0x53, 0xe8, 0x65, 0xf5, 0x23, 0x21, 0xfd, 0xe9, 0xe3, 0x70,
	0x87, 0xc2, 0xa6, 0xfa, 0x2d, 0x78, 0x6b, 0xb1, 0xd8, 0xda, 0xc1, 0xfa, 0x67, 0xe1, 0xbe, 0x89,
	0xc2, 0x46, 0x1f, 0x6a, 0x75, 0x2f, 0x92, 0x94, 0x33, 0x00, 0x3b, 0x59, 0xac, 0x54, 0x3e, 0x80,
	0x5e, 0x8c, 0xa5, 0x8a, 0x2f, 0x56, 0x3a, 0x1f, 0x82, 0xa7, 0x12, 0xd4, 0xa3, 0xa2, 0xac, 0x40,
	0xad, 0xf8, 0x29, 0xf8, 0x61, 0x83, 0xfb, 0x51, 0xf1, 0x4c, 0x0f, 0xa8, 0xbe, 0x15, 0xbc, 0x05,
	0xbd, 0x8e, 0xf2, 0x9c, 0xa7, 0x04, 0xf5, 0xd8, 0x4b, 0x70, 0xbf, 0x60, 0x7f, 0x42, 0xf9, 0xd3,
	0x60, 0xbf, 0x2a, 0x52, 0xf2, 0x06, 0x06, 0x1a, 0x71, 0xa5, 0x2d, 0xe5, 0x68, 0x95, 0x7b, 0x30,
	0xf2, 0x7c, 0x63, 0x5a, 0x87, 0x40, 0x4f, 0xf7, 0x83, 0x76, 0xda, 0xd7, 0xfd, 0x27, 0xfb, 0xe6,
	0x00, 0xef, 0x17, 0x8b, 0x19, 0xff, 0x5a, 0xa1, 0x3b, 0x5b, 0xcb, 0x1e, 0xac, 0x46, 0xaf, 0x8d,
	0x78, 0xb4, 0x35, 0x96, 0xb3, 0xed, 0x97, 0x4b, 0x1f, 0x7e, 0xb3, 0x5a, 0x4f, 0x61, 0x05, 0x63,
	0xf0, 0x89, 0x43, 0x16, 0x48, 0x4d, 0x9b, 0x12, 0x37, 0x44, 0xd1, 0xd7, 0x01, 0xe2, 0x65, 0x29,
	0xca, 0x9a, 0x43, 0xab, 0xf9, 0xc0, 0x55, 0x4b, 0x0d, 0xf1, 0x5b, 0x3b, 0xf9, 0xed, 0x6d, 0x7e,
	0xe7, 0x4f, 0x7e, 0x2d, 0xc9, 0xa2, 0x3e, 0x55, 0x9a, 0x92, 0x9a, 0x7e, 0xf0, 0x1d, 0x7c, 0xe2,
	0x30, 0xaa, 0x59, 0x27, 0xe3, 0x2f, 0x93, 0xfe, 0xe7, 0xb6, 0x82, 0x0b, 0xf0, 0x3f, 0x25, 0xb2,
	0x3d, 0xa5, 0xbe, 0x43, 0x2d, 0xcf, 0xb5, 0x56, 0xbb, 0x11, 0x54, 0x5f, 0x37, 0xca, 0xac, 0x3e,
	0x2c, 0x34, 0x9d, 0x4b, 0x09, 0xbe, 0x84, 0x41, 0xdd, 0xc7, 0x3c, 0xc9, 0x2b, 0xe8, 0x24, 0x98,
	0xf5, 0xfa, 0x9e, 0x1a, 0xb3, 0xa8, 0xaf, 0x45, 0xc0, 0xe1, 0x68, 0xc6, 0x33, 0x71, 0xcb, 0xef,
	0x36, 0x1b, 0x27, 0x70, 0xbf, 0xa1, 0x31, 0x8e, 0x31, 0xfd, 0xe9, 0x80, 0xd7, 0x24, 0x57, 0xb2,
	0x4b, 0x70, 0x30, 0x56, 0x6c, 0xbc, 0x7f, 0x98, 0x4d, 0xba, 0x47, 0xcf, 0x0f, 0xa8, 0xac, 0x45,
	0x04, 0xf7, 0x74, 0x77, 0x8c, 0x89, 0xa9, 0xfb, 0x26, 0xad, 0xa6, 0xee, 0xad, 0xcc, 0x61, 0xf7,
	0x2b, 0x70, 0xf5, 0xee, 0x98, 0x01, 0xd4, 0xca, 0xc9, 0xe8, 0xc5, 0x21, 0xa5, 0x6b, 0x02, 0x0e,
	0xdd, 0xda, 0x57, 0x76, 0x62, 0xfa, 0x25, 0x68, 0x2d, 0x79, 0x74, 0x7a, 0x58, 0x71, 0x43, 0x33,
	0xef, 0xd2, 0x1f, 0xc6, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xe7, 0xdb, 0x69, 0xbb,
	0x06, 0x00, 0x00,
}
