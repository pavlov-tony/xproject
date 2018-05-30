// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("APP_PROJECT_ID")

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucket, err := client.Buckets(ctx, projectID).Next()
	if err != nil {
		log.Fatal("fatal")
	}

	rc, err := client.Bucket(bucket.Name).Object("test-2018-05-23.csv").NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	csvReader := csv.NewReader(rc)
	res, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res[1][0])
}
