package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/errors"
)

var (
	Token string
	erl   = "https://slack.com/api/"
)

func Do(method string, args url.Values, auth bool) ([]byte, error) {
	u := erl + method
	if auth {
		args.Set("token", Token)
	}
	rsp, err := http.PostForm(u, args)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Respond(method string, args url.Values, rsp interface{}) error {
	b, err := Do(method, args, true)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}

func NoAuthRespond(method string, args url.Values, rsp interface{}) error {
	b, err := Do(method, args, false)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
