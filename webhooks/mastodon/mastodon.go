package mastodon

import "github.com/poetworrier/mastools/api"

type StatusWebhook struct {
	Event     string      `json:"event"`
	CreatedAt string      `json:"created_at"` // TODO: this should be time.Time
	Object    *api.Status `json:"object"`
}
