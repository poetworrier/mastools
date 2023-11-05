package cmd

import (
	"log"
	"log/slog"

	"github.com/poetworrier/mastools/api/mastodon"
	"github.com/spf13/cobra"
)

var getTrendingTags = &cobra.Command{
	Use:   "get-trending-tags",
	Short: "Lists trending tags",
	Run: func(cmd *cobra.Command, args []string) {
		loadAccessToken()
		c, closer := mastodon.NewMastodonClient(origin, accessToken, debug)
		defer closer()

		t := mastodon.NewTrends(c)
		s, err := t.ListTags()
		if err != nil {
			log.Fatal(err)
		}
		for i := range s {
			slog.Info("tag", s[i].Name, s[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(getTrendingTags)
}
