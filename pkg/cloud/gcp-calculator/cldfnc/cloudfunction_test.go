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

func Test_Cloudfunction_MbToGb(t *testing.T) {
	if mbToGb(1024) != 1.0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_KbToGb(t *testing.T) {
	if kbToGb(1024*1024) != 1.0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_MhzToGHz(t *testing.T) {
	if mhzToGHz(1000) != 1.0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_MsToS(t *testing.T) {
	if msToS(1000) != 1.0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_NanoToNormal(t *testing.T) {
	if nanoToNormal(1000000) != 0.001 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_CalcMem(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("F01C-3EA0-06CD")

	if calcMem(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem) > 0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_CalcCPU(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("C024-9C10-2A5B")

	if calcCPU(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem) != 400000.0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_CalcNet(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("4BAF-1AD8-483C")

	if calcNet(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem) > 0 {
		t.Error("Error in calculating")
	}
}

func Test_Cloudfunction_CalcInvoc(t *testing.T) {
	serv, _ := ser.NewClient("29E7-DA93-CA13")
	priceMem, _ := serv.GetPriceInfoBySku("8E10-82EB-6917")

	if calcInvoc(ServiceInfo{128, 200, "MHz", 300, 0, "KB", 10000000}, priceMem) > 80000000.0 {
		t.Error("Error in calculating")
	}
}
