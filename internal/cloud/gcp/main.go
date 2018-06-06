package main

import (
	"context"
	"fmt"
	"log"
	"time"
	// TODO: imports must be from src/ - github.com/pavlov-tony/xproject/internal/cloud/gcp/utils/billingutils
	"xproject/internal/cloud/gcp/utils/billingutils"
	"xproject/internal/cloud/gcp/utils/storageutils"

	"cloud.google.com/go/storage"
)

const (
	pkgLogPref = "cloud billing"
)

// TODO: it's not good place for main, if your need test you package use unit tests *_test.go
func main() {
	ctx := context.Background()

	// TODO: don't forget about commenting your code
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: logging must be more informative
		log.Fatalf("%v: new client", pkgLogPref, err)
	}

	objects, err := storageutils.FetchBucketObjects(ctx, client, "churomann-bucket")
	if err != nil {
		log.Fatal(err)
	}

	var serviceBills billingutils.ServicesBills
	serviceBills.FillByObjects(ctx, client, &objects)

	var sum float64
	for _, sb := range serviceBills {
		sum += sb.Cost
	}
	// TODO: use log pkg instead fmt
	fmt.Println("Full cost from bucket:", sum)

	sum = 0
	objects, err = objects.SelectInTimeRange(
		time.Date(2018, time.May, 29, 0, 0, 0, 0, time.Local),
		time.Now())
	if err != nil {
		log.Fatal(err)
	}
	serviceBills.FillByObjects(ctx, client, &objects)
	for _, sb := range serviceBills {
		sum += sb.Cost
	}
	fmt.Println("Full cost in time period:", sum)
}
