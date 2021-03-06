// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/slack-srv/proto/rtm/rtm.proto

/*
Package go_micro_srv_slack_rtm is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/rtm/rtm.proto

It has these top-level messages:
	StartRequest
	StartResponse
*/
package go_micro_srv_slack_rtm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Rtm service

type RtmService interface {
	Start(ctx context.Context, in *StartRequest, opts ...client.CallOption) (*StartResponse, error)
}

type rtmService struct {
	c           client.Client
	serviceName string
}

func RtmServiceClient(serviceName string, c client.Client) RtmService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.rtm"
	}
	return &rtmService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rtmService) Start(ctx context.Context, in *StartRequest, opts ...client.CallOption) (*StartResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Rtm.Start", in)
	out := new(StartResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rtm service

type RtmHandler interface {
	Start(context.Context, *StartRequest, *StartResponse) error
}

func RegisterRtmHandler(s server.Server, hdlr RtmHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Rtm{hdlr}, opts...))
}

type Rtm struct {
	RtmHandler
}

func (h *Rtm) Start(ctx context.Context, in *StartRequest, out *StartResponse) error {
	return h.RtmHandler.Start(ctx, in, out)
}
