package handler

import (
	"net/url"

	api "github.com/micro/slack-srv/proto/api"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Api struct{}

func (c *Api) Test(ctx context.Context, req *api.TestRequest, rsp *api.TestResponse) error {
	vals := url.Values{}
	if len(req.Error) > 0 {
		vals.Set("error", req.Error)
	}

	for k, v := range req.Args {
		vals.Set(k, v)
	}

	return slack.NoAuthRespond("api.test", vals, rsp)
}
