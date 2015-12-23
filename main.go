package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/micro/slack-srv/handler"
	"github.com/micro/slack-srv/slack"
)

func main() {
	cmd.Flags = append(cmd.Flags, cli.StringFlag{
		Name:   "api_token",
		EnvVar: "API_TOKEN",
		Usage:  "Slack api token",
	})

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		slack.Token = c.String("api_token")
	})

	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.slack"),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Api),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Auth),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Channels),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Chat),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Emoji),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Groups),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.IM),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Reactions),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Team),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Users),
		),
	)

	server.Handle(
		server.NewHandler(
			new(handler.Rtm),
		),
	)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
