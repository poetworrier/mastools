// Webhooks provide json format converters that are useful for webhooks
package webhooks_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/api/discord"
	"github.com/poetworrier/mastools/api/mastodon"
	"github.com/poetworrier/mastools/converters/webhooks"
	"github.com/poetworrier/mastools/converters/webhooks/testdata"
)

func TestStatusConverter_Forward(t *testing.T) {
	var conv webhooks.StatusConverter

	type args struct {
		mastodon string
	}
	tests := []struct {
		name    string
		conv    webhooks.StatusConverter
		args    args
		want    string
		wantErr bool
	}{
		{
			"test nil fails",
			conv,
			args{
				"",
			},
			"",
			true,
		},
		{
			"test full example",
			conv,
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
			got, err := tt.conv.Forward(m)
			if (err != nil) && !tt.wantErr {
				t.Errorf("StatusConverter.Forward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(*got, *want) {
				t.Errorf("StatusConverter.Forward() = %v, want %v", got, want)
			}
		})
	}
}

func TestStatusConverter_Backward(t *testing.T) {
	var converter webhooks.StatusConverter
	_, err := converter.Backward(nil)
	if !errors.Is(err, errors.ErrUnsupported) {
		t.Error("unexpected error returned", err)
		return
	}
}
