package handler

import (
	"net/url"

	team "github.com/micro/slack-srv/proto/team"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Team struct{}

func (c *Team) Info(ctx context.Context, req *team.InfoRequest, rsp *team.InfoResponse) error {
	return slack.Respond("team.info", url.Values{}, rsp)
}
