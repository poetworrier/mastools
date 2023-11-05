// Statuses contains the Cloud Run function code that

package statuses

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/imroc/req/v3"
	"github.com/poetworrier/mastools/api/mastodon"
	"github.com/poetworrier/mastools/converters/webhooks"
)

var statusesURL string
var client *req.Client

func init() {
	functions.HTTP("DiscordStatuses", DiscordStatuses)

	statusesURL = os.Getenv("DISCORD_WEBHOOK_STATUSES_URL")
	client = req.C()
}

// TODO: testing... could use httptest?
// DiscordStatuses will proxy webhooks originating from mastodon.
func DiscordStatuses(w http.ResponseWriter, r *http.Request) {
	var m mastodon.StatusWebhook
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil && !writeErr(w, err) {
		return
	}
	converter := webhooks.StatusConverter{}
	d, err := converter.Forward(&m)
	if err != nil && !writeErr(w, err) {
		return
	}
	resp, err := client.R().SetBodyJsonMarshal(d).Post(statusesURL)
	if err != nil && !writeErr(w, err) {
		return
	}
	w.WriteHeader(resp.StatusCode)
}

func writeErr(w http.ResponseWriter, err error) bool {
	w.WriteHeader(400)
	_, err = fmt.Fprint(w, err)
	if err != nil {
		slog.Error("write failed", err)
		return false
	}
	return true
}
