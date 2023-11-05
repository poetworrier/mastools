package mastodon_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/poetworrier/mastools/api/mastodon"
)

func TestList(t *testing.T) {
	const (
		tag = iota
		statuses
	)

	type args struct {
		name    string
		kind    int
		path    string
		want    string
		wantErr bool
	}
	var tests []args
	for _, p := range []struct {
		path string
		kind int
	}{
		{"/api/v1/admin/trends/statuses", statuses},
		{"/api/v1/admin/trends/tags", tag},
	} {
		tests = []args{
			{
				"empty",
				p.kind,
				p.path,
				"[]",
				false,
			},
			{
				"no data",
				p.kind,
				p.path,
				"",
				true,
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name+tt.path, func(t *testing.T) {
			// TODO: figure out how to make it work with TLS, Req hates the self-signed cert
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.path {
					w.WriteHeader(400)
					return
				}
				io.WriteString(w, tt.want)
			}))
			defer ts.Close()
			c, cls := mastodon.NewMastodonClient(ts.URL, "", false)
			defer cls()

			tr := mastodon.NewTrends(c)
			switch tt.kind {
			case tag:
				got, err := tr.ListTags()
				if !tt.wantErr {
					check(t, tt.want, err, &got)
				}
			case statuses:
				got, err := tr.ListStatus()
				if !tt.wantErr {
					check(t, tt.want, err, &got)
				}
			}
		})
	}
}

func check[T any](t *testing.T, want string, err error, got *T) {
	if err != nil {
		t.Error(errors.Join(fmt.Errorf("wanted %q: %w", want, err)))
		return
	}
	if ok, err := unmarshalDeepEqual([]byte(want), got); !ok {
		t.Error(err)
		return
	}
}

func unmarshalDeepEqual[T any](b []byte, got *T) (bool, error) {
	var want T
	err := json.Unmarshal(b, &want)
	if err != nil {
		return false, err
	}
	if !reflect.DeepEqual(want, *got) {
		return false, fmt.Errorf("want=%v, got=%v", want, *got)
	}
	return true, nil
}
