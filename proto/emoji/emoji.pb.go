// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/slack-srv/proto/emoji/emoji.proto

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
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

func (m *ListResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ListResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

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
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Emoji service

type EmojiClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type emojiClient struct {
	cc *grpc.ClientConn
}

func NewEmojiClient(cc *grpc.ClientConn) EmojiClient {
	return &emojiClient{cc}
}

func (c *emojiClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/go.micro.srv.slack.emoji.Emoji/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Emoji service

type EmojiServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterEmojiServer(s *grpc.Server, srv EmojiServer) {
	s.RegisterService(&_Emoji_serviceDesc, srv)
}

func _Emoji_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.slack.emoji.Emoji/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Emoji_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.slack.emoji.Emoji",
	HandlerType: (*EmojiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Emoji_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/slack-srv/proto/emoji/emoji.proto",
}

func init() { proto.RegisterFile("github.com/micro/slack-srv/proto/emoji/emoji.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x42, 0x10, 0xbd, 0x02, 0x42, 0x16, 0x43, 0xd4, 0x29, 0x8a, 0x04, 0xca, 0xc2,
	0x45, 0x84, 0xa5, 0x62, 0xaf, 0x58, 0x98, 0xbc, 0xb0, 0xd2, 0x46, 0x56, 0x09, 0x69, 0x7a, 0xe5,
	0xec, 0x44, 0xca, 0x2f, 0xe3, 0xef, 0xa1, 0x9c, 0x85, 0x60, 0x41, 0xb0, 0x58, 0x7e, 0xcf, 0x77,
	0xf7, 0x3d, 0x1f, 0x54, 0xdb, 0xc6, 0xbf, 0xf6, 0x1b, 0xac, 0xa9, 0x2b, 0xbb, 0xa6, 0x66, 0x2a,
	0xdd, 0x6e, 0x5d, 0xb7, 0xb7, 0x8e, 0x87, 0xf2, 0xc0, 0xe4, 0xa9, 0xb4, 0x1d, 0xbd, 0x35, 0xe1,
	0x44, 0x71, 0x74, 0xba, 0x25, 0x94, 0x5a, 0x74, 0x3c, 0xa0, 0xd4, 0xa3, 0xbc, 0xe7, 0xe7, 0x30,
	0x7f, 0x6a, 0x9c, 0x37, 0xf6, 0xbd, 0xb7, 0xce, 0xe7, 0x1f, 0x0a, 0xce, 0x82, 0x76, 0x07, 0xda,
	0x3b, 0xab, 0x2f, 0x20, 0xa2, 0x36, 0x55, 0x99, 0x2a, 0x4e, 0x4d, 0x44, 0xad, 0xbe, 0x82, 0xc4,
	0x32, 0x13, 0xa7, 0x51, 0xa6, 0x8a, 0x99, 0x09, 0x42, 0x3f, 0x42, 0x22, 0xe3, 0xd2, 0x38, 0x8b,
	0x8b, 0x79, 0x75, 0x87, 0xbf, 0xf1, 0xf0, 0xe7, 0x70, 0x5c, 0x4d, 0xd6, 0x6a, 0xef, 0x79, 0x34,
	0xa1, 0x7f, 0xb1, 0x04, 0xf8, 0x36, 0xf5, 0x25, 0xc4, 0xad, 0x1d, 0x85, 0x3e, 0x33, 0xd3, 0x75,
	0xc2, 0x0f, 0xeb, 0x5d, 0x6f, 0xbf, 0xf0, 0x22, 0x1e, 0xa2, 0xa5, 0xaa, 0x5e, 0x20, 0x91, 0x4e,
	0xfd, 0x0c, 0xc7, 0x13, 0x44, 0x5f, 0xff, 0x15, 0x42, 0x7e, 0xbc, 0xb8, 0xf9, 0x5f, 0xd6, 0xfc,
	0x68, 0x73, 0x22, 0xbb, 0xbc, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x24, 0xc9, 0x98, 0x01, 0x81,
	0x01, 0x00, 0x00,
}
