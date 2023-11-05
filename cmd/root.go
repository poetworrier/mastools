package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/poetworrier/mastools/gcp/secrets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, accessToken, origin, secretName string
var debug bool

var rootCmd = &cobra.Command{
	Use:   "mastools",
	Short: "A bunch of mastodon tools",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mastools.yaml)")
	rootCmd.PersistentFlags().StringVar(&accessToken, "accessToken", "", "Mastodon instance OAuth2 access token")
	rootCmd.PersistentFlags().StringVar(&origin, "origin", "https://pebble.social", "Mastodon instance origin URL")
	rootCmd.PersistentFlags().StringVar(&secretName, "secretName", "", "GCP Secret name")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
}

func loadAccessToken() {
	if accessToken == "" {
		// yikes, gotta be a better way to load from flag or viper
		if secretName == "" {
			secretName = viper.GetString("secretName")
			if secretName == "" {
				log.Fatal(errors.New("no accessToken or secretName provided "))

			}
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		var err error
		accessToken, err = secrets.AccessSecretVersion(ctx, secretName)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mastools" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mastools")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
