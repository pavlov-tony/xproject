package main

import (
	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/cloudbilling/v1"
	// "google.golang.org/api/googleapi"
)

func main() {
	ctx := context.Background()
	hc, err := google.DefaultClient(ctx, cloudbilling.CloudPlatformScope)
	if err != nil {
		fmt.Println(err)
	}
	client, err := cloudbilling.New(hc)
	if err != nil {
		fmt.Println(err)
	}

	// name := "billingAccounts/" + os.Getenv("APP_BILLING_ACCOUNT_ID")

	// resp, err := client.BillingAccounts.Projects.List(name).Context(ctx).Do()
	resp, err := client.Services.List().Context(ctx).Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("billing, services list - %+v", resp)

}
