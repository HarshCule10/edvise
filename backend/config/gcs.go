package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

var GCSClient *storage.Client
var BucketName = "eduvise-bucket" // Your GCS bucket name

// ConnectGCS initializes Google Cloud Storage client
func ConnectGCS() {
	ctx := context.Background()

	// Load GCP credentials from the JSON key file
	credFile := "config/eduvise-bucket.json" // Update path if needed
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credFile)

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	GCSClient = client
	fmt.Println("✅ Connected to Google Cloud Storage!")
}
