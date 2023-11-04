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
	discord.EmbedWebhook | mastodon.StatusWebhook
}

type Converter[Src Convertable, Dst Convertable] interface {
	Forward(s *Src) (*Dst, error)
	Backward(d *Dst) (*Src, error)
}

// Converts mastodon status webhooks into discord embed webhooks
// TODO: if it's stateless, could it use value receivers?
type StatusConverter struct{}

// It is an error to pass a status webhook with a nil Object field.
func (c *StatusConverter) Forward(m *mastodon.StatusWebhook) (*discord.EmbedWebhook, error) {
	if m == nil || m.Object == nil {
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
	return &discord.EmbedWebhook{
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

func (c *StatusConverter) Backward(d *discord.EmbedWebhook) (*mastodon.StatusWebhook, error) {
	return nil, errors.ErrUnsupported
}
