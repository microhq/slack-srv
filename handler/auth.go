package handler

import (
	"encoding/json"
	"net/url"

	"github.com/micro/go-micro/errors"
	auth "github.com/micro/slack-srv/proto/auth"
	"github.com/micro/slack-srv/slack"
	"golang.org/x/net/context"
)

type Auth struct{}

func (c *Auth) Test(ctx context.Context, req *auth.TestRequest, rsp *auth.TestResponse) error {
	b, err := slack.Do("auth.test", url.Values{})
	if err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
