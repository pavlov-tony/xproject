package main

import (
	"fmt"
	"os"

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
		_ = err
	}
	client, err := cloudbilling.New(hc)
	if err != nil {
		fmt.Println(err)
		_ = err
	}

	name := "billingAccounts/" + os.Getenv("APP_BILLING_ACCOUNT_ID")

	resp, err := client.BillingAccounts.Projects.List(name).Context(ctx).Do()
	// resp, err := client.Services.List().Context(ctx).Do()
	if err != nil {
		fmt.Println(err)
		_ = err
	}

	fmt.Println(resp.ProjectBillingInfo[1].Name)
	fmt.Println(resp.ProjectBillingInfo[1].ProjectId)
	fmt.Println(resp.ProjectBillingInfo[1].BillingAccountName)
	fmt.Println(resp.ProjectBillingInfo[1].BillingEnabled)

	// fmt.Println
}
