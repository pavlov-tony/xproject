package main

// sumSize (in bytes) of files in concrete bucket

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
)

type Bucket struct {
	name         string
	storageClass string // TODO: const iota
	location     string // TODO: const iota
}

// TODO: check max size and max int
type BucketObject struct {
	name       string
	size       uint64
	bucketName string
}

func main() {
	ctx := context.Background()

	// TODO: check scope
	cli, err := google.DefaultClient(ctx, storage.CloudPlatformScope)
	if err != nil {
		log.Fatal("1\n\n", err)
	}

	storageService, err := storage.New(cli)
	if err != nil {
		log.Fatal("2\n\n", err)
	}

	name := os.Getenv("APP_PROJECT_ID")
	resp, err := storageService.Buckets.List(name).Context(ctx).Do() // TODO: why should use Context.Do?
	if err != nil {
		log.Fatal("3\n\n", err)
	}

	// fill list of buckets
	buckets := []Bucket{}
	for _, b := range resp.Items {
		buckets = append(buckets, Bucket{b.Name, b.StorageClass, b.Location})
	}

	// fill list of files in bucket
	bucketObjects := []BucketObject{}
	for _, b := range buckets {
		resp, err := storageService.Objects.List(b.name).Context(ctx).Do()
		if err != nil {
			log.Fatal("4\n\n", err)
		}

		for _, o := range resp.Items {
			bucketObjects = append(bucketObjects, BucketObject{o.Name, o.Size, b.name})
		}
	}

	// calc sum size of files in bucket[0]
	var sumSize uint64
	for _, o := range bucketObjects {
		if o.bucketName == buckets[0].name {
			sumSize += o.size
		}
	}

	fmt.Println(sumSize)
}
