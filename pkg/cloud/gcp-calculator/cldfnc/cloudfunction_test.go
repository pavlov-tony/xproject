package cldfnc

import (
	"math"
	"testing"
)

func Test_Cloudfunction_Calculate(t *testing.T) {
	res1, err := Calculate(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, "29E7-DA93-CA13")

	if err != nil {
		t.Error("Failed to calculate: ", err)
	}

	if math.Round(res1*100)/100 != 7.20 {
		t.Error("Error in calculatiing")
	}

	res2, err := Calculate(ServiceInfo{256, 400, "MHz", 500, 5, "KB", 50000000}, "29E7-DA93-CA13")

	if err != nil {
		t.Error("Failed to calculate: ", err)
	}

	if math.Round(res2*100)/100 != 159.84 {
		t.Error("Error in calculatiing")
	}
}
