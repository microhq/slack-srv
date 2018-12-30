// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/slack-srv/proto/channels/channels.proto

/*
Package go_micro_srv_slack is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/channels/channels.proto

It has these top-level messages:
	Channel
	ListRequest
	ListResponse
	HistoryRequest
	HistoryResponse
	InfoRequest
	InfoResponse
	MarkRequest
	MarkResponse
	SetPurposeRequest
	SetPurposeResponse
	SetTopicRequest
	SetTopicResponse
*/
package go_micro_srv_slack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/micro/slack-srv/proto"

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

// Client API for Channels service

type ChannelsService interface {
	History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error)
	Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error)
	SetPurpose(ctx context.Context, in *SetPurposeRequest, opts ...client.CallOption) (*SetPurposeResponse, error)
	SetTopic(ctx context.Context, in *SetTopicRequest, opts ...client.CallOption) (*SetTopicResponse, error)
}

type channelsService struct {
	c           client.Client
	serviceName string
}

func ChannelsServiceClient(serviceName string, c client.Client) ChannelsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack"
	}
	return &channelsService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *channelsService) History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.History", in)
	out := new(HistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsService) Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.Info", in)
	out := new(InfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsService) Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.Mark", in)
	out := new(MarkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsService) SetPurpose(ctx context.Context, in *SetPurposeRequest, opts ...client.CallOption) (*SetPurposeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.SetPurpose", in)
	out := new(SetPurposeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsService) SetTopic(ctx context.Context, in *SetTopicRequest, opts ...client.CallOption) (*SetTopicResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Channels.SetTopic", in)
	out := new(SetTopicResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Channels service

type ChannelsHandler interface {
	History(context.Context, *HistoryRequest, *HistoryResponse) error
	Info(context.Context, *InfoRequest, *InfoResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Mark(context.Context, *MarkRequest, *MarkResponse) error
	SetPurpose(context.Context, *SetPurposeRequest, *SetPurposeResponse) error
	SetTopic(context.Context, *SetTopicRequest, *SetTopicResponse) error
}

func RegisterChannelsHandler(s server.Server, hdlr ChannelsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Channels{hdlr}, opts...))
}

type Channels struct {
	ChannelsHandler
}

func (h *Channels) History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error {
	return h.ChannelsHandler.History(ctx, in, out)
}

func (h *Channels) Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error {
	return h.ChannelsHandler.Info(ctx, in, out)
}

func (h *Channels) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ChannelsHandler.List(ctx, in, out)
}

func (h *Channels) Mark(ctx context.Context, in *MarkRequest, out *MarkResponse) error {
	return h.ChannelsHandler.Mark(ctx, in, out)
}

func (h *Channels) SetPurpose(ctx context.Context, in *SetPurposeRequest, out *SetPurposeResponse) error {
	return h.ChannelsHandler.SetPurpose(ctx, in, out)
}

func (h *Channels) SetTopic(ctx context.Context, in *SetTopicRequest, out *SetTopicResponse) error {
	return h.ChannelsHandler.SetTopic(ctx, in, out)
}