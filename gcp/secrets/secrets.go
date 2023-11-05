// Secrets is for interaction with the GCP secrets manager
package secrets

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// WARNING: Do not print the secret in a production environment
//
// AccessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
//
// Snippet based upon https://github.com/GoogleCloudPlatform/golang-samples/blob/968c611f22fca94d82ffcb3ab77b52d51bc4408f/secretmanager/access_secret_version.go#L17-L67
// Copyright 2019 Google LLC - Apache V2
//
// This version is modified from the original to take a [context.Context] and return [(string, error)]
func AccessSecretVersion(ctx context.Context, name string) (string, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	result, err := client.AccessSecretVersion(ctx,
		&secretmanagerpb.AccessSecretVersionRequest{
			Name: name,
		})
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
