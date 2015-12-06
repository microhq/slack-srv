package handler

import (
	"net/url"

	emoji "github.com/micro/slack-srv/proto/emoji"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Emoji struct{}

func (c *Emoji) List(ctx context.Context, req *emoji.ListRequest, rsp *emoji.ListResponse) error {
	return slack.Respond("emoji.list", url.Values{}, rsp)
}
