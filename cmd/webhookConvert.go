package cmd

import (
	"encoding/json"
	"log"

	"github.com/poetworrier/mastools/api/mastodon"
	"github.com/poetworrier/mastools/converters/webhooks"
	"github.com/spf13/cobra"
)

// webhookConvertCmd represents the webhookConvert command
var webhookConvertCmd = &cobra.Command{
	Use:   "webhook-convert",
	Short: "Converts webhook json from one format to another",
	Run: func(cmd *cobra.Command, args []string) {
		var conv webhooks.StatusConverter
		var m mastodon.StatusWebhook
		err := json.NewDecoder(cmd.InOrStdin()).Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		d, err := conv.Forward(&m)
		if err != nil {
			log.Fatal(err)
		}
		err = json.NewEncoder(cmd.OutOrStdout()).Encode(d)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(webhookConvertCmd)
}
