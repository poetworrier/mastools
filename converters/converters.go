// Converters create mappings from one type to the next. You might say they're automatic.
package converters

import (
	"github.com/poetworrier/mastools/api/discord"
	"github.com/poetworrier/mastools/api/mastodon"
)

type Convertable interface {
	discord.EmbedWebhook | mastodon.StatusWebhook
}

type Converter[Src Convertable, Dst Convertable] interface {
	Forward(s *Src) (*Dst, error)
	Backward(d *Dst) (*Src, error)
}
