// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/slack-srv/proto/groups/groups.proto

/*
Package go_micro_srv_slack_group is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/slack-srv/proto/groups/groups.proto

It has these top-level messages:
	Group
	CloseRequest
	CloseResponse
	HistoryRequest
	HistoryResponse
	InfoRequest
	InfoResponse
	ListRequest
	ListResponse
	MarkRequest
	MarkResponse
	OpenRequest
	OpenResponse
	SetPurposeRequest
	SetPurposeResponse
	SetTopicRequest
	SetTopicResponse
*/
package go_micro_srv_slack_group

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

// Client API for Groups service

type GroupsService interface {
	Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error)
	History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error)
	Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error)
	Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error)
	SetPurpose(ctx context.Context, in *SetPurposeRequest, opts ...client.CallOption) (*SetPurposeResponse, error)
	SetTopic(ctx context.Context, in *SetTopicRequest, opts ...client.CallOption) (*SetTopicResponse, error)
}

type groupsService struct {
	c           client.Client
	serviceName string
}

func GroupsServiceClient(serviceName string, c client.Client) GroupsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.slack.group"
	}
	return &groupsService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *groupsService) Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.Close", in)
	out := new(CloseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.History", in)
	out := new(HistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.Info", in)
	out := new(InfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) Mark(ctx context.Context, in *MarkRequest, opts ...client.CallOption) (*MarkResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.Mark", in)
	out := new(MarkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.Open", in)
	out := new(OpenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) SetPurpose(ctx context.Context, in *SetPurposeRequest, opts ...client.CallOption) (*SetPurposeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.SetPurpose", in)
	out := new(SetPurposeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) SetTopic(ctx context.Context, in *SetTopicRequest, opts ...client.CallOption) (*SetTopicResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Groups.SetTopic", in)
	out := new(SetTopicResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Groups service

type GroupsHandler interface {
	Close(context.Context, *CloseRequest, *CloseResponse) error
	History(context.Context, *HistoryRequest, *HistoryResponse) error
	Info(context.Context, *InfoRequest, *InfoResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Mark(context.Context, *MarkRequest, *MarkResponse) error
	Open(context.Context, *OpenRequest, *OpenResponse) error
	SetPurpose(context.Context, *SetPurposeRequest, *SetPurposeResponse) error
	SetTopic(context.Context, *SetTopicRequest, *SetTopicResponse) error
}

func RegisterGroupsHandler(s server.Server, hdlr GroupsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Groups{hdlr}, opts...))
}

type Groups struct {
	GroupsHandler
}

func (h *Groups) Close(ctx context.Context, in *CloseRequest, out *CloseResponse) error {
	return h.GroupsHandler.Close(ctx, in, out)
}

func (h *Groups) History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error {
	return h.GroupsHandler.History(ctx, in, out)
}

func (h *Groups) Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error {
	return h.GroupsHandler.Info(ctx, in, out)
}

func (h *Groups) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.GroupsHandler.List(ctx, in, out)
}

func (h *Groups) Mark(ctx context.Context, in *MarkRequest, out *MarkResponse) error {
	return h.GroupsHandler.Mark(ctx, in, out)
}

func (h *Groups) Open(ctx context.Context, in *OpenRequest, out *OpenResponse) error {
	return h.GroupsHandler.Open(ctx, in, out)
}

func (h *Groups) SetPurpose(ctx context.Context, in *SetPurposeRequest, out *SetPurposeResponse) error {
	return h.GroupsHandler.SetPurpose(ctx, in, out)
}

func (h *Groups) SetTopic(ctx context.Context, in *SetTopicRequest, out *SetTopicResponse) error {
	return h.GroupsHandler.SetTopic(ctx, in, out)
}