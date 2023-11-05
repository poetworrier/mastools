// The api package contains http json bindings for Mastodon
package mastodon_test

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/api/mastodon"
)

func TestStatus_LogValue(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		s    mastodon.Status
		want slog.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.LogValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status.LogValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
