// Webhooks provide json format converters that are useful for webhooks
package webhooks

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/webhooks/discord"
	"github.com/poetworrier/mastools/webhooks/mastodon"
	"github.com/poetworrier/mastools/webhooks/testdata"
)

func TestStatusConverter_Forward(t *testing.T) {
	converter := StatusConverter{}

	type args struct {
		mastodon string
	}
	tests := []struct {
		name    string
		c       *StatusConverter
		args    args
		want    string
		wantErr bool
	}{
		{
			"test nil fails",
			&converter,
			args{
				"",
			},
			"",
			true,
		},
		{
			"test full example",
			&converter,
			args{
				testdata.MastodonStatusExample,
			},
			testdata.DiscordWebhookExample,
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want *discord.EmbedWebhook
			if tt.want != "" {
				err := json.Unmarshal([]byte(tt.want), &want)
				if err != nil {
					t.Error(err)
					return
				}
			}

			var m *mastodon.StatusWebhook
			if tt.args.mastodon != "" {
				err := json.Unmarshal([]byte(tt.args.mastodon), &m)
				if err != nil {
					t.Error(err)
					return
				}
			}
			got, err := tt.c.Forward(m)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusConverter.Forward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("StatusConverter.Forward() = %v, want %v", got, want)
			}
		})
	}
}

func TestStatusConverter_Backward(t *testing.T) {
	type args struct {
		d *discord.EmbedWebhook
	}
	tests := []struct {
		name    string
		c       *StatusConverter
		args    args
		want    *mastodon.StatusWebhook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Backward(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusConverter.Backward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusConverter.Backward() = %v, want %v", got, tt.want)
			}
		})
	}
}
