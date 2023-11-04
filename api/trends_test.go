package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/poetworrier/mastools/api"
)

func TestListStatuses(t *testing.T) {
	type args struct {
		name string
		s    string
	}
	tests := []args{
		{
			"no statuses",
			"[]",
		},
		{
			"empty",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: figure out how to make it work with TLS, Req hates the self-signed cert
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/api/v1/admin/trends/statuses" {
					w.WriteHeader(400)
					return
				}
				io.WriteString(w, "[]")
			}))
			defer ts.Close()
			c, cls := api.NewClient(ts.URL, "", false)
			defer cls()

			tr := api.NewTrends(c)
			status, err := tr.ListStatus()
			if err != nil {
				t.Fatal(err)
			}
			if len(status) != 0 {
				t.Errorf("non-empty status: len=%d", len(status))
			}
		})
	}
}
