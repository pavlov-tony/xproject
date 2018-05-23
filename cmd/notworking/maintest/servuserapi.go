package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/serviceuser/v1"
	// "google.golang.org/api/serviceusage/v1"
)

func main() {
	// NOTE: How does it work?
	ctx := context.Background()

	//getting client with readonly scope
	cli, err := google.DefaultClient(ctx, serviceuser.CloudPlatformReadOnlyScope)
	if err != nil {
		log.Fatal("context\n", err)
	}

	serviceuserService, err := serviceuser.New(cli)
	if err != nil {
		log.Fatal("service\n", err)
	}

	name := "projects/" + os.Getenv("APP_PROJECT_ID")

	resp, err := serviceuserService.Projects.Services.List(name).Context(ctx).Do()
	if err != nil {
		log.Fatal("list\n", err)
	}
	_ = resp
	// resp, err := serviceusageService.Services.Search().Context(ctx).Do()
	// resp, err := serviceusageService.Services.ListEnabled(name).Context(ctx).Do()
	// if err != nil {
	// 	log.Fatal("list enabled\n", err)
	// }
	// for _, s := range resp.Services {
	// 	fmt.Println(s.Name)
	// 	fmt.Println(s.Service.Name)
	// }
}
