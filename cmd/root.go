package cmd

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
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

	cobra.OnInitialize(func() {
		if accessToken == "" {
			if secretName == "" {
				log.Fatal(errors.New("no accessToken or secretName provided "))
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			var err error
			accessToken, err = accessSecretVersion(ctx, secretName)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
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

// TODO: boy would be nice to have a metautil project...
// accessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
// TODO: in a more generalized version this should might take a context
// WARNING: Do not print the secret in a production environment
func accessSecretVersion(ctx context.Context, name string) (string, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %w", err)
	}

	// Verify the data checksum.
	crc32c := crc32.MakeTable(crc32.Castagnoli)
	checksum := int64(crc32.Checksum(result.Payload.Data, crc32c))
	if checksum != *result.Payload.DataCrc32C {
		return "", errors.New("data corruption detected")
	}

	return string(result.Payload.Data), nil
}
