package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx,
		option.WithCredentialsFile("../pppp-91f5204dae0e.json"))
	if err != nil {
		log.Fatal(err)
	}

	it := storageClient.Buckets(ctx, "hale-entry-204416")
	log.Printf("proj-id: %v", "hale-entry-204416")
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}
}
