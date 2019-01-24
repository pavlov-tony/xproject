package cldfnc

import (
	"fmt"
	"math"

	ser "github.com/pavlov-tony/xproject/internal/srvcln"

	b "google.golang.org/api/cloudbilling/v1"
)

// ids of skus
const (
	CPUTime       = "C024-9C10-2A5B"
	Invocations   = "8E10-82EB-6917"
	MemoryTime    = "F01C-3EA0-06CD"
	NetworkEgress = "4BAF-1AD8-483C"
)

// struct for cloud function parameters
// example: ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, "29E7-DA93-CA13")
type ServiceInfo struct {
	MemoryUsage  float64
	CpuUsage     float64
	CpuType      string
	Time         float64
	NetworkUsage float64
	NetworkType  string
	Invocations  float64
}

// calculates price of cloud function
func Calculate(s ServiceInfo, serviceId string) (float64, error) {

	serv, err := ser.NewClient(serviceId)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate price of cloud func: %v", err)
	}

	priceCPU, err := serv.GetPriceInfoBySku(CPUTime)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate price of cloud func: %v", err)
	}

	priceMem, err := serv.GetPriceInfoBySku(MemoryTime)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate price of cloud func: %v", err)
	}

	priceNet, err := serv.GetPriceInfoBySku(NetworkEgress)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate price of cloud func: %v", err)
	}

	priceInvoc, err := serv.GetPriceInfoBySku(Invocations)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate price of cloud func: %v", err)
	}

	cpu := calcCPU(s, priceCPU) * (float64(priceCPU.PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(priceCPU.PricingExpression.TieredRates[1].UnitPrice.Nanos))
	mem := calcMem(s, priceMem) * (float64(priceMem.PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(priceMem.PricingExpression.TieredRates[1].UnitPrice.Nanos))
	net := calcNet(s, priceNet) * (float64(priceNet.PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(priceNet.PricingExpression.TieredRates[1].UnitPrice.Nanos))
	inv := calcInvoc(s, priceInvoc) * (float64(priceInvoc.PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(priceInvoc.PricingExpression.TieredRates[1].UnitPrice.Nanos))

	return cpu + mem + net + inv, nil
}

// transform megabytes to gigabytes
func mbToGb(v float64) float64 {
	return v / 1024
}

// transform megabytes to kilobytes
func kbToGb(v float64) float64 {
	return v / 1024 / 1024
}

// transform megahertz to gigahertz
func mhzToGHz(v float64) float64 {
	return v / 1000
}

// transform millieseconds to seconds
func msToS(v float64) float64 {
	return v / 1000
}

// transform nanos to normal view
func nanoToNormal(v int64) float64 {
	return float64(v) / math.Pow(10, 9)
}

// calculates memory price
func calcMem(s ServiceInfo, price b.PricingInfo) float64 {
	v := mbToGb(s.MemoryUsage)*msToS(s.Time)*s.Invocations - price.PricingExpression.TieredRates[1].StartUsageAmount

	if v < 0 {
		return 0
	}
	return v
}

// calculates cpu price
func calcCPU(s ServiceInfo, price b.PricingInfo) float64 {
	var v float64

	if s.CpuType == "MHz" {
		v = mhzToGHz(s.CpuUsage)*msToS(s.Time)*s.Invocations - price.PricingExpression.TieredRates[1].StartUsageAmount
	} else {
		v = s.CpuUsage*msToS(s.Time)*s.Invocations - price.PricingExpression.TieredRates[1].StartUsageAmount
	}

	if v < 0 {
		return 0
	}
	return v
}

// calculates network ingress price
func calcNet(s ServiceInfo, price b.PricingInfo) float64 {
	var v float64

	if s.NetworkUsage == 0 {
		return 0
	}

	if s.NetworkType == "KB" {
		v = s.Invocations*kbToGb(s.NetworkUsage) - price.PricingExpression.TieredRates[1].StartUsageAmount
	} else {
		v = s.Invocations*mbToGb(s.NetworkUsage) - price.PricingExpression.TieredRates[1].StartUsageAmount
	}

	if v < 0 {
		return 0
	}
	return v
}

// calculates price of invocations
func calcInvoc(s ServiceInfo, price b.PricingInfo) float64 {
	v := s.Invocations - price.PricingExpression.TieredRates[1].StartUsageAmount

	if v < 0 {
		return 0
	}
	return v
}
