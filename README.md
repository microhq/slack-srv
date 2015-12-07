# Slack Server ![Slack](https://github.com/micro/slack-srv/blob/master/slack.jpg)

Slack server is slackbot developed as a go-micro service.

Implements bot API methods:

- Api
- Auth
- Channels
- Chat
- Emoji
- Groups
- Reactions
- Users

Incomplete (Someone please PR this)

- Rtm

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Add a new bot in your Slack Team and get your API token

4. Download and start the service

	```shell
	go get github.com/micro/slack-srv
	slack-srv --api_token=YOUR_API_TOKEN
	```

	OR as a docker container

	```shell
	docker run microhq/slack-srv --api_token=YOUR_API_TOKEN --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API
Slack server implements the [bot API](https://api.slack.com/bot-users) as RPC

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
