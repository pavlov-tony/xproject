package gcpparser

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"
)

func Test_Parse(t *testing.T) {
	f, _ := os.Open("./testdata/test1.csv")
	r := bufio.NewReader(f)
	records, _ := csv.NewReader(r).ReadAll()

	// skip row with headers
	res, err := Parse(records[1:])
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}
	exp := "com.google.cloud/services/cloud-storage/StorageMultiRegionalUsGbsec"
	if res[0].Item != exp {
		t.Errorf("parse:\nexp: %v, res %v", exp, res[0].Item)
	}
}
