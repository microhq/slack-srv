package handler

import (
	"fmt"
	"net/url"

	groups "github.com/micro/slack-srv/proto/groups"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Groups struct{}

func (c *Groups) Close(ctx context.Context, req *groups.CloseRequest, rsp *groups.CloseResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	return slack.Respond("groups.close", vals, rsp)
}

func (c *Groups) History(ctx context.Context, req *groups.HistoryRequest, rsp *groups.HistoryResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("latest", fmt.Sprintf("%.6f", req.Latest))
	vals.Set("oldest", fmt.Sprintf("%.6f", req.Oldest))
	if req.Inclusive > 0 {
		vals.Set("inclusive", "1")
	}
	if req.Count > 0 {
		vals.Set("count", fmt.Sprintf("%d", req.Count))
	}
	if req.Unreads > 0 {
		vals.Set("unreads", "1")
	}
	return slack.Respond("groups.history", vals, rsp)
}

func (c *Groups) Info(ctx context.Context, req *groups.InfoRequest, rsp *groups.InfoResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	return slack.Respond("groups.info", vals, rsp)
}

func (c *Groups) List(ctx context.Context, req *groups.ListRequest, rsp *groups.ListResponse) error {
	vals := url.Values{}
	if req.ExcludeArchived > 0 {
		vals.Set("exclude_archived", "1")
	}
	return slack.Respond("groups.list", vals, rsp)
}

func (c *Groups) Mark(ctx context.Context, req *groups.MarkRequest, rsp *groups.MarkResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	if req.Ts > 0.0 {
		vals.Set("ts", fmt.Sprintf("%.6f", req.Ts))
	}
	return slack.Respond("groups.mark", vals, rsp)
}

func (c *Groups) Open(ctx context.Context, req *groups.OpenRequest, rsp *groups.OpenResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	return slack.Respond("groups.open", vals, rsp)
}

func (c *Groups) SetPurpose(ctx context.Context, req *groups.SetPurposeRequest, rsp *groups.SetPurposeResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("purpose", req.Purpose)
	return slack.Respond("groups.setPurpose", vals, rsp)
}

func (c *Groups) SetTopic(ctx context.Context, req *groups.SetTopicRequest, rsp *groups.SetTopicResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("topic", req.Topic)
	return slack.Respond("groups.setTopic", vals, rsp)
}
