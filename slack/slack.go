package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	Token string
	erl   = "https://slack.com/api/"
)

func Do(method string, args url.Values) ([]byte, error) {
	u := erl + method
	args.Set("token", Token)

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
