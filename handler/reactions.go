package handler

import (
	"fmt"
	"net/url"

	reactions "github.com/micro/slack-srv/proto/reactions"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Reactions struct{}

func (c *Reactions) Add(ctx context.Context, req *reactions.AddRequest, rsp *reactions.AddResponse) error {
	vals := url.Values{}
	vals.Set("name", req.Name)
	vals.Set("file", req.File)
	vals.Set("file_comment", req.FileComment)
	vals.Set("channel", req.Channel)
	if req.Timestamp > 0.0 {
		vals.Set("timestamp", fmt.Sprintf("%.6f", req.Timestamp))
	}
	return slack.Respond("reactions.add", vals, rsp)
}

func (c *Reactions) Get(ctx context.Context, req *reactions.GetRequest, rsp *reactions.GetResponse) error {
	vals := url.Values{}
	vals.Set("file", req.File)
	vals.Set("file_comment", req.FileComment)
	vals.Set("channel", req.Channel)
	if req.Timestamp > 0.0 {
		vals.Set("timestamp", fmt.Sprintf("%.6f", req.Timestamp))
	}
	if req.Full {
		vals.Set("full", fmt.Sprintf("%t", req.Full))
	}
	return slack.Respond("reactions.get", vals, rsp)
}

func (c *Reactions) List(ctx context.Context, req *reactions.ListRequest, rsp *reactions.ListResponse) error {
	vals := url.Values{}
	vals.Set("user", req.User)
	if req.Full {
		vals.Set("full", fmt.Sprintf("%t", req.Full))
	}
	if req.Count > 0 {
		vals.Set("count", fmt.Sprintf("%d", req.Count))
	}
	if req.Page > 0 {
		vals.Set("page", fmt.Sprintf("%d", req.Page))
	}
	return slack.Respond("reactions.list", vals, rsp)
}

func (c *Reactions) Remove(ctx context.Context, req *reactions.RemoveRequest, rsp *reactions.RemoveResponse) error {
	vals := url.Values{}
	vals.Set("name", req.Name)
	vals.Set("file", req.File)
	vals.Set("file_comment", req.FileComment)
	vals.Set("channel", req.Channel)
	if req.Timestamp > 0.0 {
		vals.Set("timestamp", fmt.Sprintf("%.6f", req.Timestamp))
	}
	return slack.Respond("reactions.remove", vals, rsp)
}
