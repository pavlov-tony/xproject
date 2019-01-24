package cldfnc

import (
	"math"
	"testing"

	ser "github.com/pavlov-tony/xproject/internal/srvcln"
)

func Test_Cloudfunction_Calculate(t *testing.T) {
	res1, err := Calculate(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, "29E7-DA93-CA13")

	if err != nil {
		t.Error("Failed to calculate: ", err)
	}

	if got := math.Round(res1*100) / 100; got != 7.20 {
		t.Errorf("Error in calculatiing: expected: 7.20, got: %v", got)
	}

	res2, err := Calculate(ServiceInfo{256, 400, "MHz", 500, 5, "KB", 50000000}, "29E7-DA93-CA13")

	if err != nil {
		t.Error("Failed to calculate: ", err)
	}

	if got := math.Round(res2*100) / 100; got != 159.84 {
		t.Errorf("Error in calculatiing: expected: 159.84, got: %v", got)
	}
}

func Test_Cloudfunction_MbToGb(t *testing.T) {
	if got := mbToGb(1024); got != 1.0 {
		t.Errorf("Error in calculatiing: expected: 1.0, got: %v", got)
	}
}

func Test_Cloudfunction_KbToGb(t *testing.T) {
	if got := kbToGb(1024 * 1024); got != 1.0 {
		t.Errorf("Error in calculatiing: expected: 1.0, got: %v", got)
	}
}

func Test_Cloudfunction_MhzToGHz(t *testing.T) {
	if got := mhzToGHz(1000); got != 1.0 {
		t.Errorf("Error in calculatiing: expected: 1.0, got: %v", got)
	}
}

func Test_Cloudfunction_MsToS(t *testing.T) {
	if got := msToS(1000); got != 1.0 {
		t.Errorf("Error in calculatiing: expected: 1.0, got: %v", got)
	}
}

func Test_Cloudfunction_NanoToNormal(t *testing.T) {
	if got := nanoToNormal(1000000); got != 0.001 {
		t.Errorf("Error in calculatiing: expected: 0.001, got: %v", got)
	}
}

func Test_Cloudfunction_CalcMem(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("F01C-3EA0-06CD")

	if got := calcMem(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem); got > 0 {
		t.Errorf("Error in calculatiing: expected: 0, got: %v", got)
	}
}

func Test_Cloudfunction_CalcCPU(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("C024-9C10-2A5B")

	if got := calcCPU(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem); got != 400000.0 {
		t.Errorf("Error in calculatiing: expected: 400000.0, got: %v", got)
	}
}

func Test_Cloudfunction_CalcNet(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("4BAF-1AD8-483C")

	if got := calcNet(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem); got > 0 {
		t.Errorf("Error in calculatiing: expected: 0, got: %v", got)
	}
}

func Test_Cloudfunction_CalcInvoc(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("8E10-82EB-6917")

	if got := calcInvoc(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem); got != 8000000.0 {
		t.Errorf("Error in calculatiing: expected: 80000000.0, got: %v", got)
	}
}
