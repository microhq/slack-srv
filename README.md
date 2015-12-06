# Slack Server ![Slack](https://github.com/micro/slack-srv/blob/master/slack.jpg)

Slack server is slackbot developed as a go-micro service.

Work in progress. Currently implements API methods:

- Api
- Auth
- Channels
- Chat
- Emoji
- Groups

## Getting started

1. Add a new bot and get your API token

2. Download the service

```shell
go get github.com/micro/slack-srv
```

3. Start the service with your token
```shell
slack-srv --api_token=YOUR_API_TOKEN
```

Slack server implements the [bot API](https://api.slack.com/bot-users)

```shell
micro query go.micro.srv.slack Channels.List 
{
	"channels": [
		{
			"created": 1.446682059e+09,
			"creator": "...",
			"id": "...",
			"is_member": true,
			"members": [
				...
			],
			"name": "general",
			"num_members": 34,
			"purpose": {
				"value": "This channel is for team-wide communication and announcements. All team members are in this channel."
			},
			"topic": {
				"creator": "...",
				"last_set": 1.44814907e+09,
				"value": "https://micro-services.co"
			}
		},
		{
			"created": 1.446682059e+09,
			"creator": "...",
			"id": "...",
			"members": [
				...
			],
			"name": "random",
			"num_members": 33,
			"purpose": {
				"value": "A place for non-work-related flimflam, faffing, hodge-podge or jibber-jabber you'd prefer to keep out of more focused work-related channels."
			},
			"topic": {
				"value": "Non-work banter and water cooler conversation"
			}
		}
	],
	"ok": true
}
```
