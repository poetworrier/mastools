package testdata

import _ "embed"

//go:embed mastodon_status_example.json
var MastodonStatusExample string

// TODO: enable to generate goldens
// -- go:generate go run . webhook-convert  < webhooks/testdata/mastodon_status_example.json 
//go:embed discord_webhook_example.json
var DiscordWebhookExample string