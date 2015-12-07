package handler

import (
	"fmt"
	"net/url"

	rtm "github.com/micro/slack-srv/proto/rtm"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Rtm struct{}

func (c *Rtm) Start(ctx context.Context, req *rtm.StartRequest, rsp *rtm.StartResponse) error {
	u := url.Values{}
	u.Set("simple_latest", fmt.Sprintf("%t", req.SimpleLatest))
	u.Set("no_unreads", fmt.Sprintf("%t", req.NoUnreads))
	u.Set("mpim_aware", fmt.Sprintf("%t", req.MpimAware))
	return slack.Respond("rtm.start", u, rsp)
}
