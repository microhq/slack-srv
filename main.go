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
	app := cli.NewApp()
	app.HideVersion = true
	app.Flags = cmd.Flags
	app.Flags = append(app.Flags, cli.StringFlag{
		Name:   "api_token",
		EnvVar: "API_TOKEN",
		Usage:  "Slack api token",
	})
	app.Before = cmd.Setup
	app.Action = func(c *cli.Context) {
		slack.Token = c.String("api_token")
	}
	app.RunAndExitOnError()
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

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
