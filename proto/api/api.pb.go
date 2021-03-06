// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/slack-srv/proto/api/api.proto

/*
Package go_micro_srv_slack_api is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/api/api.proto

It has these top-level messages:
	TestRequest
	TestResponse
*/
package go_micro_srv_slack_api

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

type TestRequest struct {
	Error string            `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Args  map[string]string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TestRequest) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type TestResponse struct {
	Ok    bool              `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Args  map[string]string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error string            `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TestResponse) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TestResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "go.micro.srv.slack.api.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "go.micro.srv.slack.api.TestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := grpc.Invoke(ctx, "/go.micro.srv.slack.api.Api/Test", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.slack.api.Api/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.slack.api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Api_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/slack-srv/proto/api/api.proto",
}

func init() { proto.RegisterFile("github.com/micro/slack-srv/proto/api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0xce, 0x49,
	0x4c, 0xce, 0xd6, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0xc8,
	0x04, 0x61, 0x3d, 0x30, 0x4f, 0x48, 0x2c, 0x3d, 0x5f, 0x0f, 0xac, 0x4e, 0xaf, 0xb8, 0xa8, 0x4c,
	0x0f, 0xac, 0x56, 0x2f, 0xb1, 0x20, 0x53, 0x69, 0x3e, 0x23, 0x17, 0x77, 0x48, 0x6a, 0x71, 0x49,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x51, 0x51, 0x7e, 0x91,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0xe4, 0xc8, 0xc5, 0x92, 0x58, 0x94, 0x5e,
	0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xab, 0x87, 0xdd, 0x30, 0x3d, 0x24, 0x83, 0xf4,
	0x1c, 0x8b, 0xd2, 0x8b, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xc0, 0x5a, 0xa5, 0xcc, 0xb9, 0x38,
	0xe1, 0x42, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x3b, 0x40, 0x4c, 0x90, 0xbd, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x10, 0x7b, 0xc1, 0x1c, 0x2b, 0x26, 0x0b, 0x46, 0xa5, 0x8d,
	0x8c, 0x5c, 0x3c, 0x10, 0x83, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0xf2,
	0xb3, 0xc1, 0x7a, 0x39, 0x82, 0x98, 0xf2, 0xb3, 0x85, 0x9c, 0x50, 0x1c, 0xa7, 0x87, 0xdf, 0x71,
	0x10, 0x33, 0xd0, 0x5d, 0x87, 0xf0, 0x36, 0x33, 0x92, 0xb7, 0xc9, 0x76, 0xb3, 0x51, 0x14, 0x17,
	0xb3, 0x63, 0x41, 0xa6, 0x50, 0x30, 0x17, 0x0b, 0xc8, 0x56, 0x21, 0x65, 0x22, 0x02, 0x4c, 0x4a,
	0x85, 0x18, 0x87, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x23, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0x18, 0x62, 0xd1, 0x02, 0x02, 0x00, 0x00,
}
