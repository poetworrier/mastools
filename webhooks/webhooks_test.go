// Webhooks provide json format converters that are useful for webhooks
package webhooks

import (
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/webhooks/discord"
	"github.com/poetworrier/mastools/webhooks/mastodon"
)

func TestStatusConverter_Forward(t *testing.T) {
	converter := StatusConverter{}

	type args struct {
		m *mastodon.MastodonStatus
	}
	tests := []struct {
		name    string
		c       *StatusConverter
		args    args
		want    *discord.Webhook
		wantErr bool
	}{
		{
			"test nil fails",
			&converter,
			args{
				&mastodon.MastodonStatus{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Forward(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusConverter.Forward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusConverter.Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusConverter_Backward(t *testing.T) {
	type args struct {
		d *discord.Webhook
	}
	tests := []struct {
		name    string
		c       *StatusConverter
		args    args
		want    *mastodon.MastodonStatus
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
