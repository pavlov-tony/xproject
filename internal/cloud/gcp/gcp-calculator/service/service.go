package service

import (
	"fmt"
	"net/http"
	"os"

	b "google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/googleapi/transport"
)

type Service struct {
	skus []b.Sku // how to make b.*(Sku), need type *Sku from b
}

// init or redo slice of Sku
func (s *Service) New(serviceId string) error {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("DEVKEY")},
	}

	billClient, err := b.New(client)
	if err != nil {
		return fmt.Errorf("Failed to get sku list: %v", err)
	}

	l, err := billClient.Services.Skus.List("services/" + serviceId).Do()
	if err != nil {
		return fmt.Errorf("Failed to get sku list: %v", err)
	}

	s.skus = nil

	for i := range l.Skus {
		s.skus = append(s.skus, *(l.Skus[i]))
	}
	if s.skus == nil {
		return fmt.Errorf("Empty sku list")
	}

	return nil
}

// this function returns type PricingInfo from billing api, we took PricingInfo[0] because it's the latest
func (s *Service) GetPriceInfoBySku(id string) (price b.PricingInfo) {

	for _, val := range s.skus {
		if val.SkuId == id {
			price = *val.PricingInfo[0]
			break
		}
	}

	return price
}
