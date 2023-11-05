package mastodon

type StatusWebhook struct {
	Event     string  `json:"event"`
	CreatedAt string  `json:"created_at"` // TODO: this should be time.Time
	Object    *Status `json:"object"`
}
