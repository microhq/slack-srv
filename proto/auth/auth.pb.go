// Code generated by protoc-gen-go.
// source: github.com/micro/slack-srv/proto/auth/auth.proto
// DO NOT EDIT!

/*
Package go_micro_srv_slack_auth is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/auth/auth.proto

It has these top-level messages:
	TestRequest
	TestResponse
*/
package go_micro_srv_slack_auth

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

type TestRequest struct {
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TestResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Team   string `protobuf:"bytes,4,opt,name=team" json:"team,omitempty"`
	User   string `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	TeamId string `protobuf:"bytes,6,opt,name=team_id" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,7,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TestRequest)(nil), "go.micro.srv.slack.auth.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "go.micro.srv.slack.auth.TestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Test(ctx context.Context, in *TestRequest) (*TestResponse, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Test(ctx context.Context, in *TestRequest) (*TestResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Test", in)
	out := new(TestResponse)
	err := c.c.Call(ctx, req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Test(context.Context, *TestRequest, *TestResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler) {
	s.Handle(s.NewHandler(hdlr))
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xdd, 0xdd, 0xec, 0xae, 0x4e, 0x5b, 0x84, 0x5c, 0x0c, 0x9e, 0xa4, 0x28, 0x78, 0x71,
	0x22, 0xfa, 0x0b, 0xfc, 0x0b, 0xa2, 0x47, 0x91, 0xb6, 0x86, 0xb6, 0xb4, 0x75, 0xea, 0x24, 0xf1,
	0xf7, 0x9b, 0xa4, 0x14, 0xbc, 0x88, 0x97, 0x84, 0xef, 0xbd, 0x07, 0xf3, 0x66, 0xe0, 0xbe, 0xed,
	0x5d, 0xe7, 0x6b, 0x6c, 0x68, 0xd2, 0x53, 0xdf, 0x30, 0x69, 0x3b, 0x56, 0xcd, 0x70, 0x67, 0xf9,
	0x5b, 0xcf, 0x4c, 0x8e, 0x74, 0xe5, 0x5d, 0x97, 0x1e, 0x4c, 0x2c, 0x2f, 0x5a, 0xc2, 0x94, 0xc4,
	0x90, 0xc1, 0x94, 0xc6, 0x68, 0x97, 0x05, 0x64, 0x2f, 0xc6, 0xba, 0x67, 0xf3, 0xe5, 0xc3, 0x57,
	0x3a, 0xc8, 0x17, 0xb4, 0x33, 0x7d, 0x5a, 0x23, 0x01, 0xb6, 0x34, 0xa8, 0xcd, 0xd5, 0xe6, 0xf6,
	0x54, 0x16, 0xb0, 0x37, 0xcc, 0xc4, 0x6a, 0x1b, 0xf0, 0x4c, 0x66, 0xb0, 0xf3, 0x3c, 0xaa, 0x5d,
	0x82, 0x1c, 0x84, 0x33, 0xd5, 0xa4, 0xc4, 0x4a, 0xde, 0x1a, 0x56, 0xfb, 0x44, 0xe7, 0x70, 0x8c,
	0xde, 0x7b, 0xff, 0xa1, 0x0e, 0xab, 0x10, 0xed, 0x28, 0x1c, 0xa3, 0xf0, 0xf0, 0x06, 0xe2, 0x29,
	0x94, 0x91, 0xaf, 0x20, 0xe2, 0x74, 0x79, 0x8d, 0x7f, 0xd4, 0xc5, 0x5f, 0x5d, 0x2f, 0x6f, 0xfe,
	0x49, 0x2d, 0x2b, 0x94, 0x27, 0xf5, 0x21, 0xdd, 0xe0, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x20, 0x26, 0xc3, 0x37, 0x01, 0x00, 0x00,
}
