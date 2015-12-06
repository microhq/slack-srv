// Code generated by protoc-gen-go.
// source: github.com/micro/slack-srv/proto/emoji/emoji.proto
// DO NOT EDIT!

/*
Package go_micro_srv_slack_emoji is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/emoji/emoji.proto

It has these top-level messages:
	ListRequest
	ListResponse
*/
package go_micro_srv_slack_emoji

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

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListResponse struct {
	Ok    bool              `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string            `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Emoji map[string]string `protobuf:"bytes,3,rep,name=emoji" json:"emoji,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetEmoji() map[string]string {
	if m != nil {
		return m.Emoji
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.slack.emoji.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.srv.slack.emoji.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Emoji service

type EmojiClient interface {
	List(ctx context.Context, in *ListRequest) (*ListResponse, error)
}

type emojiClient struct {
	c           client.Client
	serviceName string
}

func NewEmojiClient(serviceName string, c client.Client) EmojiClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.emoji"
	}
	return &emojiClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *emojiClient) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Emoji.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Emoji service

type EmojiHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterEmojiHandler(s server.Server, hdlr EmojiHandler) {
	s.Handle(s.NewHandler(hdlr))
}

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0xce, 0x49,
	0x4c, 0xce, 0xd6, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcd, 0xcd,
	0xcf, 0xca, 0x84, 0x90, 0x7a, 0x60, 0x11, 0x21, 0x89, 0xf4, 0x7c, 0x3d, 0xb0, 0x5a, 0x3d, 0xa0,
	0x2a, 0x3d, 0xb0, 0x7a, 0x3d, 0xb0, 0xbc, 0x12, 0x2f, 0x17, 0xb7, 0x4f, 0x66, 0x71, 0x49, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0xd2, 0x4c, 0x46, 0x2e, 0x1e, 0x08, 0xbf, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x88, 0x8b, 0x8b, 0x29, 0x3f, 0x5b, 0x82, 0x51, 0x81, 0x51, 0x83, 0x43, 0x88,
	0x97, 0x8b, 0x35, 0xb5, 0xa8, 0x28, 0xbf, 0x48, 0x82, 0x09, 0xc8, 0xe5, 0x14, 0x72, 0x00, 0x72,
	0x41, 0x66, 0x48, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0x19, 0xea, 0xe1, 0xb2, 0x44, 0x0f, 0xd9,
	0x44, 0x3d, 0x57, 0x90, 0x90, 0x6b, 0x5e, 0x49, 0x51, 0xa5, 0x94, 0x0e, 0x17, 0x17, 0x82, 0x27,
	0xc4, 0xcd, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xb6, 0x8b, 0x13, 0x64, 0x57, 0x59, 0x62, 0x4e, 0x69,
	0x2a, 0xc4, 0x2e, 0x2b, 0x26, 0x0b, 0x46, 0xa3, 0x04, 0x2e, 0x56, 0xb0, 0x6a, 0xa1, 0x70, 0x2e,
	0x16, 0x90, 0x89, 0x42, 0xaa, 0x84, 0x6c, 0x04, 0xfb, 0x49, 0x4a, 0x8d, 0x38, 0x87, 0x29, 0x31,
	0x24, 0xb1, 0x81, 0x43, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x74, 0x55, 0xd1, 0x24, 0x63,
	0x01, 0x00, 0x00,
}
