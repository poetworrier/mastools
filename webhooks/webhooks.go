// Webhooks provide json format converters that are useful for webhooks
package webhooks

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/poetworrier/mastools/webhooks/discord"
	"github.com/poetworrier/mastools/webhooks/mastodon"
)

type Convertable interface {
	discord.Webhook | mastodon.MastodonStatus
}

type Converter[Src Convertable, Dst Convertable] interface {
	Forward(s *Src) (*Dst, error)
	Backward(d *Dst) (*Src, error)
}

// Converts mastodon status webhooks into discord embed webhooks
type StatusConverter struct{}

// TODO: if it single stateless method, could these be value receivers?
func (c *StatusConverter) Forward(m *mastodon.MastodonStatus) (*discord.Webhook, error) {
	if m.Object == nil {
		return nil, errors.New("cannot convert nil")
	}
	var fields []discord.Field
	for name, value := range map[string]int{
		"replies":    m.Object.RepliesCount,
		"reblogs":    m.Object.ReblogsCount,
		"favourites": m.Object.FavouritesCount} {

		fields = append(fields, discord.Field{
			Name:   name,
			Value:  strconv.Itoa(value),
			Inline: true,
		})
	}
	return &discord.Webhook{
		Username:  m.Object.Account.Username,
		AvatarURL: m.Object.Account.Avatar,
		Content:   fmt.Sprintf("[view post](%s)", m.Object.URI),
		Embeds: []discord.Embed{
			{
				Author: discord.Author{
					Name:    m.Object.Account.DisplayName,
					URL:     m.Object.Account.URL,
					IconURL: m.Object.Account.Avatar,
				},
				Timestamp: m.Object.CreatedAt,
				Fields:    fields,
			},
		},
	}, nil
}

func (c *StatusConverter) Backward(d *discord.Webhook) (*mastodon.MastodonStatus, error) {
	return nil, errors.ErrUnsupported
}
