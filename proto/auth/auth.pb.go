// Code generated by protoc-gen-go.
// source: github.com/micro/slack-srv/proto/auth/auth.proto
// DO NOT EDIT!

/*
Package go_micro_srv_slack is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/auth/auth.proto

It has these top-level messages:
	AuthTestRequest
	AuthTestResponse
*/
package go_micro_srv_slack

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

type AuthTestRequest struct {
}

func (m *AuthTestRequest) Reset()                    { *m = AuthTestRequest{} }
func (m *AuthTestRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthTestRequest) ProtoMessage()               {}
func (*AuthTestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AuthTestResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Team   string `protobuf:"bytes,4,opt,name=team" json:"team,omitempty"`
	User   string `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	TeamId string `protobuf:"bytes,6,opt,name=team_id" json:"team_id,omitempty"`
	UserId string `protobuf:"bytes,7,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *AuthTestResponse) Reset()                    { *m = AuthTestResponse{} }
func (m *AuthTestResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthTestResponse) ProtoMessage()               {}
func (*AuthTestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AuthTestRequest)(nil), "go.micro.srv.slack.AuthTestRequest")
	proto.RegisterType((*AuthTestResponse)(nil), "go.micro.srv.slack.AuthTestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Test(ctx context.Context, in *AuthTestRequest) (*AuthTestResponse, error)
}

type authClient struct {
	c client.Client
}

func NewAuthClient(c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	return &authClient{
		c: c,
	}
}

func (c *authClient) Test(ctx context.Context, in *AuthTestRequest) (*AuthTestResponse, error) {
	req := c.c.NewRequest("go.micro.srv.slack", "Auth.Test", in)
	out := new(AuthTestResponse)
	err := c.c.Call(ctx, req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Test(context.Context, *AuthTestRequest, *AuthTestResponse) error
}

func RegisterAuthServer(s server.Server, srv AuthServer) {
	s.Handle(s.NewHandler(srv))
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0x85, 0x30,
	0x10, 0x45, 0x05, 0x0a, 0xe8, 0xa8, 0x41, 0xbb, 0x6a, 0x5c, 0x19, 0x74, 0xe1, 0xc6, 0xd6, 0xe8,
	0x17, 0xf8, 0x0b, 0xea, 0xce, 0x85, 0x01, 0x6c, 0x80, 0x00, 0x0e, 0x4e, 0x5b, 0xe3, 0xe7, 0xdb,
	0xd6, 0x10, 0x13, 0x5f, 0xf2, 0x36, 0x6d, 0xce, 0x9d, 0x93, 0xcc, 0x6d, 0xe1, 0xae, 0x1f, 0xed,
	0xe0, 0x5a, 0xd9, 0xe1, 0xa2, 0x96, 0xb1, 0x23, 0x54, 0x66, 0x6e, 0xba, 0xe9, 0xd6, 0xd0, 0x97,
	0x5a, 0x09, 0x2d, 0xaa, 0xc6, 0xd9, 0x21, 0x1e, 0x32, 0x32, 0xe7, 0x3d, 0xca, 0x68, 0x4a, 0xef,
	0xc8, 0x68, 0xd7, 0xe7, 0x50, 0x3d, 0x7a, 0xe3, 0x45, 0x1b, 0xfb, 0xa4, 0x3f, 0x9d, 0xbf, 0xea,
	0x6f, 0x38, 0xfb, 0x8b, 0xcc, 0x8a, 0x1f, 0x46, 0x73, 0x80, 0x14, 0x27, 0x91, 0x5c, 0x26, 0x37,
	0x87, 0xfc, 0x14, 0x72, 0x4d, 0x84, 0x24, 0x52, 0x8f, 0x47, 0xfc, 0x18, 0x32, 0x47, 0xb3, 0xc8,
	0x22, 0x9c, 0x00, 0xb3, 0xba, 0x59, 0x04, 0xdb, 0xc8, 0x19, 0x4d, 0x22, 0x8f, 0x54, 0x41, 0x19,
	0x66, 0x6f, 0xe3, 0xbb, 0x28, 0xb6, 0x20, 0x8c, 0x43, 0x50, 0x86, 0xe0, 0xfe, 0x15, 0x58, 0xd8,
	0xcc, 0x9f, 0x81, 0x85, 0xed, 0xfc, 0x4a, 0xee, 0x36, 0x96, 0xff, 0xea, 0x5e, 0x5c, 0xef, 0x97,
	0x7e, 0x1f, 0x50, 0x1f, 0xb4, 0x45, 0xfc, 0x84, 0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea,
	0xff, 0x9c, 0xb4, 0x38, 0x01, 0x00, 0x00,
}
