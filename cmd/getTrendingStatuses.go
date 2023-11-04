package cmd

import (
	"log"
	"log/slog"

	"github.com/poetworrier/mastools/api"
	"github.com/spf13/cobra"
)

// getTrendingStatuses represents the listStatuses command
var getTrendingStatuses = &cobra.Command{
	Use:   "get-trending-status",
	Short: "List trending statuses to review",
	Run: func(cmd *cobra.Command, args []string) {
		c, closer := api.NewClient(origin, accessToken, debug)
		defer closer()

		t := api.NewTrends(c)
		s, err := t.ListStatus()
		if err != nil {
			log.Fatal(err)
		}
		for i := range s {
			slog.Info("got statuses", "status", s[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(getTrendingStatuses)
}
