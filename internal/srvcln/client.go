package srvcln

import (
	"fmt"
	"net/http"
	"os"

	b "google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/googleapi/transport"
)

// struct for gcp service with it's list of skus and it's ID
type Service struct {
	skus      []b.Sku
	serviceId string
}

// creates client and gets list of skus for service
func NewClient(serviceId string) (*Service, error) {
	s := new(Service)

	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("DEVKEY")},
	}

	billClient, err := b.New(client)
	if err != nil {
		return nil, fmt.Errorf("Failed to get sku list: %v", err)
	}

	l, err := billClient.Services.Skus.List("services/" + serviceId).Do()
	if err != nil {
		return nil, fmt.Errorf("Failed to get sku list: %v", err)
	}

	s.serviceId = serviceId

	for i := range l.Skus {
		s.skus = append(s.skus, *(l.Skus[i]))
	}
	if s.skus == nil {
		return nil, fmt.Errorf("Empty sku list")
	}

	return s, nil
}

// returns type PricingInfo from billing api for sku, we took PricingInfo[0] because it's the latest
func (s *Service) GetPriceInfoBySku(id string) (price b.PricingInfo, err error) {

	for _, val := range s.skus {
		if val.SkuId == id {
			price = *val.PricingInfo[0]

			return price, nil
		}
	}

	return price, fmt.Errorf("Failed to find sku with id: %v", id)
}

// returns serviceId
func (s *Service) GetServiceId() (string, error) {
	if s.serviceId == "" {
		return "", fmt.Errorf("Empty service id")
	} else {
		return s.serviceId, nil
	}
}
