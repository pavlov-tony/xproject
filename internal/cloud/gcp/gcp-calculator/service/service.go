package service

import (
	"log"
	"net/http"
	"os"

	b "google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/googleapi/transport"
)

type Service struct {
	skus []b.Sku // how to make b.*(Sku), need type *Sku from b
}

// init slice of Skus
func (s *Service) New() {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("DEVKEY")},
	}

	billClient, err := b.New(client)

	if err != nil {
		log.Fatal(err)
	}

	l, err := billClient.Services.Skus.List("services/6F81-5844-456A").Do()
	if err != nil {
		log.Fatal(err)
	}

	for i := range l.Skus {
		s.skus = append(s.skus, *(l.Skus[i]))
	}

}

// this function returns type PricingInfo from billing api, we took PricingInfo[0] because it's the latest
func (s *Service) GetPriceInfoBySku(id string) (price b.PricingInfo) {

	for _, val := range s.skus {
		if val.SkuId == id {
			price = *val.PricingInfo[0]
		}
		break
	}

	return price
}
