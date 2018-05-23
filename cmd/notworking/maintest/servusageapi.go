package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/serviceusage/v1"
)

func main() {
	// NOTE: How does it work?
	ctx := context.Background()

	//getting client with readonly scope
	cli, err := google.DefaultClient(ctx, serviceusage.CloudPlatformReadOnlyScope)
	if err != nil {
		log.Fatal("context\n", err)
	}

	serviceusageService, err := serviceusage.New(cli)
	if err != nil {
		log.Fatal("service\n", err)
	}

	name := os.Getenv("APP_PROJECT_ID")

	resp, err := serviceusageService.Services.ListEnabled(name).Context(ctx).Do()
	if err != nil {
		log.Fatal("list enabled\n", err)
	}
	for _, s := range resp.Services {
		fmt.Println(s.Name)
		fmt.Println(s.Service.Name)
	}
}
