package handler

import (
	"fmt"
	"net/url"

	channels "github.com/micro/slack-srv/proto/channels"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Channels struct{}

func (c *Channels) List(ctx context.Context, req *channels.ListRequest, rsp *channels.ListResponse) error {
	vals := url.Values{}
	if req.ExcludeArchived > 0 {
		vals.Set("exclude_archived", "1")
	}
	return slack.Respond("channels.list", vals, rsp)
}

func (c *Channels) History(ctx context.Context, req *channels.HistoryRequest, rsp *channels.HistoryResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	if req.Latest > 0.0 {
		vals.Set("latest", fmt.Sprintf("%.6f", req.Latest))
	}
	if req.Oldest > 0.0 {
		vals.Set("oldest", fmt.Sprintf("%.6f", req.Oldest))
	}
	if req.Inclusive > 0 {
		vals.Set("inclusive", "1")
	}
	if req.Count > 0 {
		vals.Set("count", fmt.Sprintf("%d", req.Count))
	}
	if req.Unreads > 0 {
		vals.Set("unreads", fmt.Sprintf("%d", req.Unreads))
	}
	return slack.Respond("channels.history", vals, rsp)
}

func (c *Channels) Info(ctx context.Context, req *channels.InfoRequest, rsp *channels.InfoResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	return slack.Respond("channels.info", vals, rsp)
}

func (c *Channels) Mark(ctx context.Context, req *channels.MarkRequest, rsp *channels.MarkResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("ts", fmt.Sprintf("%.6f", req.Ts))
	return slack.Respond("channels.mark", vals, rsp)
}

func (c *Channels) SetPurpose(ctx context.Context, req *channels.SetPurposeRequest, rsp *channels.SetPurposeResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("purpose", req.Purpose)
	return slack.Respond("channels.setPurpose", vals, rsp)
}

func (c *Channels) SetTopic(ctx context.Context, req *channels.SetTopicRequest, rsp *channels.SetTopicResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("topic", req.Topic)
	return slack.Respond("channels.setTopic", vals, rsp)
}
