package main

// taken from vslchnk

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	b "google.golang.org/api/cloudbilling/v1"
)

func main() {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"https://www.googleapis.com/auth/cloud-billing"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)

	computeService, err := b.New(client)
	if err != nil {
		log.Fatal(err)
	}

	l, err := computeService.BillingAccounts.Projects.List("billingAccounts/010A3B-1985BE-B1826F").Do()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(l.ProjectBillingInfo); i++ {
		fmt.Println("info: ", l.ProjectBillingInfo[i])
	}
}
