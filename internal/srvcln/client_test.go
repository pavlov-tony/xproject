package srvcln

import (
	"testing"
)

func Test_Client_NewClient(t *testing.T) {
	_, err := NewClient("29E7-DA93-CA13")

	if err != nil {
		t.Error("Failed to create client: ", err)
	}
}

func Test_Client_GetPriceInfoBySku(t *testing.T) {
	serv, _ := NewClient("29E7-DA93-CA13")

	_, err := serv.GetPriceInfoBySku("C024-9C10-2A5B")

	if err != nil {
		t.Error("Failed to get sku: ", err)
	}
}

func Test_Client_GetServiceId(t *testing.T) {
	id := "29E7-DA93-CA13"

	serv, _ := NewClient(id)

	tId, err := serv.GetServiceId()

	if err != nil {
		t.Error("Failed to get service id: ", err)
	}

	if tId != id {
		t.Errorf("Different ids: expected: %v, got: %v", id, tId)
	}
}
