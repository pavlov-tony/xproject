package cloudfunction

import (
	"fmt"
	"math"

	ser "../service"
)

const (
	// CPUTime is the amount of time, on average, you expect your function to run each time it is invoked
	CPUTime = "C024-9C10-2A5B"

	// Invocations is the number of times per month that you expect your function to run
	Invocations = "8E10-82EB-6917"

	// MemoryTime is the maximum amount of memory that your function will need
	MemoryTime = "F01C-3EA0-06CD"

	// NetworkEgress is the amount of data transferred out of your function each time it is invoked
	NetworkEgress = "4BAF-1AD8-483C"
)

// ServiceInfo contains parameters of Cloud Functions service
type ServiceInfo struct {
	MemoryUsage  float64
	CPUUsage     float64
	CPUType      string
	Time         float64
	NetworkUsage float64
	NetworkType  string
	Invocations  float64
}

// Calculate prising for Cloud Functions service
// Example: ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, "29E7-DA93-CA13")
func Calculate(s ServiceInfo, serviceID string) (float64, error) {
	var serv ser.Service
	err := serv.New(serviceID)
	if err != nil {
		return 0, fmt.Errorf("Failed calculate price of cloud func: %v", err)
	}

	cpu := calcCPU(s, serv) * (float64(serv.GetPriceInfoBySku(CPUTime).PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(serv.GetPriceInfoBySku(CPUTime).PricingExpression.TieredRates[1].UnitPrice.Nanos))
	mem := calcMem(s, serv) * (float64(serv.GetPriceInfoBySku(MemoryTime).PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(serv.GetPriceInfoBySku(MemoryTime).PricingExpression.TieredRates[1].UnitPrice.Nanos))
	net := calcNet(s, serv) * (float64(serv.GetPriceInfoBySku(NetworkEgress).PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(serv.GetPriceInfoBySku(NetworkEgress).PricingExpression.TieredRates[1].UnitPrice.Nanos))
	inv := calcInvoc(s, serv) * (float64(serv.GetPriceInfoBySku(Invocations).PricingExpression.TieredRates[1].UnitPrice.Units) +
		nanoToNormal(serv.GetPriceInfoBySku(Invocations).PricingExpression.TieredRates[1].UnitPrice.Nanos))

	return cpu + mem + net + inv, nil
}

func mbToGb(v float64) float64 {
	return v / 1024
}

func kbToGb(v float64) float64 {
	return v / 1024 / 1024
}

func mhzToGHz(v float64) float64 {
	return v / 1000
}

func msToS(v float64) float64 {
	return v / 1000
}

func nanoToNormal(v int64) float64 {
	return float64(v) / math.Pow(10, 9)
}

func calcMem(s ServiceInfo, serv ser.Service) float64 {
	v := mbToGb(s.MemoryUsage)*msToS(s.Time)*s.Invocations - serv.GetPriceInfoBySku(MemoryTime).PricingExpression.TieredRates[1].StartUsageAmount

	if v < 0 {
		return 0
	}
	return v
}

func calcCPU(s ServiceInfo, serv ser.Service) float64 {
	var v float64

	if s.CPUType == "MHz" {
		v = mhzToGHz(s.CPUUsage)*msToS(s.Time)*s.Invocations - serv.GetPriceInfoBySku(CPUTime).PricingExpression.TieredRates[1].StartUsageAmount
	} else {
		v = s.CPUUsage*msToS(s.Time)*s.Invocations - serv.GetPriceInfoBySku(CPUTime).PricingExpression.TieredRates[1].StartUsageAmount
	}

	if v < 0 {
		return 0
	}
	return v
}

func calcNet(s ServiceInfo, serv ser.Service) float64 {
	var v float64

	if s.NetworkUsage == 0 {
		return 0
	}

	if s.NetworkType == "KB" {
		v = s.Invocations*kbToGb(s.NetworkUsage) - serv.GetPriceInfoBySku(NetworkEgress).PricingExpression.TieredRates[1].StartUsageAmount
	} else {
		v = s.Invocations*mbToGb(s.NetworkUsage) - serv.GetPriceInfoBySku(NetworkEgress).PricingExpression.TieredRates[1].StartUsageAmount
	}

	if v < 0 {
		return 0
	}
	return v
}

func calcInvoc(s ServiceInfo, serv ser.Service) float64 {
	v := s.Invocations - serv.GetPriceInfoBySku(Invocations).PricingExpression.TieredRates[1].StartUsageAmount

	if v < 0 {
		return 0
	}
	return v
}
