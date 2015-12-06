package handler

import (
	"fmt"
	"net/url"

	im "github.com/micro/slack-srv/proto/im"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type IM struct{}

func (c *IM) Close(ctx context.Context, req *im.CloseRequest, rsp *im.CloseResponse) error {
	u := url.Values{}
	u.Set("channel", req.Channel)
	return slack.Respond("im.close", u, rsp)
}

func (c *IM) History(ctx context.Context, req *im.HistoryRequest, rsp *im.HistoryResponse) error {
	u := url.Values{}
	u.Set("channel", req.Channel)
	if req.Latest > 0 {
		u.Set("latest", fmt.Sprintf("%.6f", req.Latest))
	}
	if req.Oldest > 0 {
		u.Set("oldest", fmt.Sprintf("%.6f", req.Oldest))
	}
	if req.Inclusive > 0 {
		u.Set("inclusive", "1")
	}
	if req.Count > 0 {
		u.Set("count", fmt.Sprintf("%d", req.Count))
	}
	if req.Unreads > 0 {
		u.Set("unreads", "1")
	}
	return slack.Respond("im.history", u, rsp)
}

func (c *IM) List(ctx context.Context, req *im.ListRequest, rsp *im.ListResponse) error {
	return slack.Respond("im.list", url.Values{}, rsp)
}

func (c *IM) Mark(ctx context.Context, req *im.MarkRequest, rsp *im.MarkResponse) error {
	u := url.Values{}
	u.Set("channel", req.Channel)
	if req.Ts > 0 {
		u.Set("ts", fmt.Sprintf("%.6f", req.Ts))
	}
	return slack.Respond("im.mark", u, rsp)
}

func (c *IM) Open(ctx context.Context, req *im.OpenRequest, rsp *im.OpenResponse) error {
	u := url.Values{}
	u.Set("user", req.User)
	return slack.Respond("im.open", u, rsp)
}
