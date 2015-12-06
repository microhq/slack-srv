package handler

import (
	"net/url"

	auth "github.com/micro/slack-srv/proto/auth"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Auth struct{}

func (c *Auth) Test(ctx context.Context, req *auth.TestRequest, rsp *auth.TestResponse) error {
	return slack.Respond("auth.test", url.Values{}, rsp)
}
