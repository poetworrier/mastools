// Statuses contains the Cloud Run function code that

package statuses

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/poetworrier/mastools/api/mastodon"
	"github.com/poetworrier/mastools/gcp/secrets"
)

var mastodonAPIKey string
var mastodonOrigin string
var debug bool

func init() {
	functions.HTTP("TrendingStatuses", TrendingStatuses)

	secretName := os.Getenv("MASTODON_API_KEY_SECRET_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var err error
	mastodonAPIKey, err = secrets.AccessSecretVersion(ctx, secretName)
	if err != nil {
		log.Fatal(err)
	}

	mastodonOrigin = os.Getenv("MASTODON_ORIGIN")
	debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}
}

func TrendingStatuses(w http.ResponseWriter, r *http.Request) {
	c, cancel := mastodon.NewMastodonClient(mastodonOrigin, mastodonAPIKey, debug)
	defer cancel()

	tr := mastodon.NewTrends(c)
	st, err := tr.ListStatus()
	if err != nil {
		slog.Error("unable to fetch statuses", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(st)
	if err != nil {
		slog.Error("unable to encode statuses", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
