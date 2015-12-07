package handler

import (
	"net/url"

	users "github.com/micro/slack-srv/proto/users"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Users struct{}

func (c *Users) GetPresence(ctx context.Context, req *users.GetPresenceRequest, rsp *users.GetPresenceResponse) error {
	vals := url.Values{}
	vals.Set("user", req.User)
	return slack.Respond("users.getPresence", vals, rsp)
}

func (c *Users) SetPresence(ctx context.Context, req *users.SetPresenceRequest, rsp *users.SetPresenceResponse) error {
	vals := url.Values{}
	vals.Set("presence", req.Presence)
	return slack.Respond("users.setPresence", vals, rsp)
}

func (c *Users) Info(ctx context.Context, req *users.InfoRequest, rsp *users.InfoResponse) error {
	vals := url.Values{}
	vals.Set("user", req.User)
	return slack.Respond("users.info", vals, rsp)
}

func (c *Users) List(ctx context.Context, req *users.ListRequest, rsp *users.ListResponse) error {
	vals := url.Values{}
	if req.Presence > 0 {
		vals.Set("presence", "1")
	}
	return slack.Respond("users.list", vals, rsp)
}
