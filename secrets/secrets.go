package secrets

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// AccessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
func AccessSecretVersion(name string) string {
	// name := "projects/my-project/secrets/my-secret/versions/5"
	// name := "projects/my-project/secrets/my-secret/versions/latest"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Print(fmt.Errorf("failed to create secretmanager client: %v", err))
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Print(fmt.Errorf("failed to access secret version: %v", err))
	}

	return string(result.Payload.Data)
}

func GetSecret(name string) string {
	// if running locally with different key, grab key from environment
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		// else use secret manager
		apiKey = AccessSecretVersion(name)
	}
	return apiKey
}
