package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/api"
)

func TestListStatuses(t *testing.T) {
	type args struct {
		name    string
		path    string
		want    string
		wantErr bool
	}
	var tests []args
	for _, path := range []string{"/api/v1/admin/trends/statuses", "/api/v1/admin/trends/tags"} {
		tests = []args{
			{
				"no statuses",
				path,
				"[]",
				false,
			},
			{
				"empty",
				path,
				"",
				true,
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: figure out how to make it work with TLS, Req hates the self-signed cert
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.path {
					w.WriteHeader(400)
					return
				}
				io.WriteString(w, tt.want)
			}))
			defer ts.Close()
			c, cls := api.NewClient(ts.URL, "", false)
			defer cls()

			tr := api.NewTrends(c)
			got, err := tr.ListStatus()
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			}
			if !tt.wantErr {
				var s []api.Status
				err := json.Unmarshal([]byte(tt.want), &s)
				if err != nil {
					t.Error(err)
					return
				}
				if !reflect.DeepEqual(s, got) {
					t.Errorf("want=%v, got=%v", s, got)
				}
			}
		})
	}
}
