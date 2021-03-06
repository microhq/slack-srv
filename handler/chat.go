package handler

import (
	"fmt"
	"net/url"

	chat "github.com/micro/slack-srv/proto/chat"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Chat struct{}

func (c *Chat) Delete(ctx context.Context, req *chat.DeleteRequest, rsp *chat.DeleteResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	if req.Ts > 0.0 {
		vals.Set("ts", fmt.Sprintf("%.6f", req.Ts))
	}
	return slack.Respond("chat.delete", vals, rsp)
}

func (c *Chat) PostMessage(ctx context.Context, req *chat.PostMessageRequest, rsp *chat.PostMessageResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("text", req.Text)
	vals.Set("username", req.Username)
	vals.Set("as_user", fmt.Sprintf("%t", req.AsUser))
	vals.Set("parse", req.Parse)
	if req.LinkNames > 0 {
		vals.Set("link_names", fmt.Sprintf("%d", req.LinkNames))
	}
	vals.Set("attachments", req.Attachments)
	vals.Set("unfurl_links", fmt.Sprintf("%t", req.UnfurlLinks))
	vals.Set("unfurl_media", fmt.Sprintf("%t", req.UnfurlMedia))
	vals.Set("icon_url", req.IconUrl)
	vals.Set("icon_emoji", req.IconEmoji)

	return slack.Respond("chat.postMessage", vals, rsp)
}

func (c *Chat) Update(ctx context.Context, req *chat.UpdateRequest, rsp *chat.UpdateResponse) error {
	vals := url.Values{}
	vals.Set("channel", req.Channel)
	vals.Set("text", req.Text)
	vals.Set("parse", req.Parse)
	vals.Set("attachments", req.Attachments)
	if req.Ts > 0.0 {
		vals.Set("ts", fmt.Sprintf("%.6f", req.Ts))
	}
	if req.LinkNames > 0 {
		vals.Set("link_names", fmt.Sprintf("%d", req.LinkNames))
	}
	return slack.Respond("chat.update", vals, rsp)
}
