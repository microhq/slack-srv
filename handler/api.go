package handler

import (
	"encoding/json"
	"net/url"

	"github.com/micro/go-micro/errors"
	api "github.com/micro/slack-srv/proto/api"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Api struct{}

func (c *Api) Test(ctx context.Context, req *api.ApiTestRequest, rsp *api.ApiTestResponse) error {
	vals := url.Values{}
	if len(req.Error) > 0 {
		vals.Set("error", req.Error)
	}

	for k, v := range req.Args {
		vals.Set(k, v)
	}

	b, err := slack.Do("api.test", vals)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
