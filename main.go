package main

import (
	"log"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/slack-srv/handler"
	"github.com/micro/slack-srv/slack"

	// protos

	api "github.com/micro/slack-srv/proto/api"
	auth "github.com/micro/slack-srv/proto/auth"
	channels "github.com/micro/slack-srv/proto/channels"
	chat "github.com/micro/slack-srv/proto/chat"
	emoji "github.com/micro/slack-srv/proto/emoji"
	groups "github.com/micro/slack-srv/proto/groups"
	im "github.com/micro/slack-srv/proto/im"
	reactions "github.com/micro/slack-srv/proto/reactions"
	rtm "github.com/micro/slack-srv/proto/rtm"
	team "github.com/micro/slack-srv/proto/team"
	users "github.com/micro/slack-srv/proto/users"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.slack"),
		micro.Flags(cli.StringFlag{
			Name:   "api_token",
			EnvVar: "API_TOKEN",
			Usage:  "Slack api token",
		}),
		micro.Action(func(c *cli.Context) {
			slack.Token = c.String("api_token")
		}),
	)

	service.Init()

	api.RegisterApiHandler(service.Server(), new(handler.Api))
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))
	channels.RegisterChannelsHandler(service.Server(), new(handler.Channels))
	chat.RegisterChatHandler(service.Server(), new(handler.Chat))
	emoji.RegisterEmojiHandler(service.Server(), new(handler.Emoji))
	groups.RegisterGroupsHandler(service.Server(), new(handler.Groups))
	im.RegisterIMHandler(service.Server(), new(handler.IM))
	reactions.RegisterReactionsHandler(service.Server(), new(handler.Reactions))
	rtm.RegisterRtmHandler(service.Server(), new(handler.Rtm))
	team.RegisterTeamHandler(service.Server(), new(handler.Team))
	users.RegisterUsersHandler(service.Server(), new(handler.Users))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
