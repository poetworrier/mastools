package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"log"
	"log/slog"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/poetworrier/mastools/api"
)

var (
	accessToken = flag.String("accessToken", "", "Mastodon instance OAuth2 access token")
	debug       = flag.Bool("debug", false, "Enable debug logging")
	origin    = flag.String("origin", "https://pebble.social", "Mastodon instance origin URL")
	secretName = flag.String("secretName", "", "GCP Secret name")
)

func main() {
	flag.Parse()

	if *accessToken == "" {
		if *secretName == "" {
			log.Fatal(errors.New("no accessToken or secretName provided "))
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second *5)
		defer cancel()

		var err error
		*accessToken, err = accessSecretVersion(ctx, *secretName)
		if err != nil {
			log.Fatal(err)
		}
	}

	 c, closer := api.NewClient(*origin, *accessToken, *debug)
	defer closer()

	t := api.NewTrends(c)
	s, err := t.ListStatus()
	if err != nil {
		log.Fatal(err)
	}
	for i := range s {
		slog.Info("got statuses", "status", s[i])
	}
}

// TODO: boy would be nice to have a metautil project...
// accessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
// TODO: in a more generalized version this should might take a context
// WARNING: Do not print the secret in a production environment
func accessSecretVersion(ctx context.Context,  name string) (string, error) {
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
