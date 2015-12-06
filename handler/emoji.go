package handler

import (
	"encoding/json"
	"net/url"

	"github.com/micro/go-micro/errors"
	emoji "github.com/micro/slack-srv/proto/emoji"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Emoji struct{}

func (c *Emoji) List(ctx context.Context, req *emoji.EmojiListRequest, rsp *emoji.EmojiListResponse) error {
	b, err := slack.Do("emoji.list", url.Values{})
	if err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
