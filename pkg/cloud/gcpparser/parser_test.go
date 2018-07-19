// Package gcpparser uses for parse raw GCP billing csv file
package gcpparser

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"
)

func Test_Parse(t *testing.T) {
	f, err := os.Open("./testdata/test1.csv")
	if err != nil {
		t.Errorf("Can not open test1.csv: %v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	res, err := Parse(records)
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}

	// compare record with data from "./testdata/test1.csv"
	if len(res) == 0 {
		t.Errorf("Got: Res len < 0, Exp: len > 0")
		t.FailNow()
	}
	exp := "com.google.cloud/services/cloud-storage/StorageMultiRegionalUsGbsec"
	if res[0].LineItem != exp {
		t.Errorf("parse:\nExp: %v, Got: %v", exp, res[0].LineItem)
	}

	// test csv file contains only header
	f, err = os.Open("./testdata/test_only_header.csv")
	if err != nil {
		t.Errorf("Can not open test_only_header.csv: %v", err)
	}

	r = bufio.NewReader(f)
	records, err = csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	res, err = Parse(records)
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}

	if len(res) != 0 {
		t.Errorf("Got: Res len > 0, Exp: len < 0")
		t.FailNow()
	}

	// test csv file contains wrong header
	f, err = os.Open("./testdata/test_wrong_header.csv")
	if err != nil {
		t.Errorf("Can not open test_wrong_header.csv: %v", err)
	}

	r = bufio.NewReader(f)
	records, err = csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("Can not parse file from bucket: %v", err)
	}

	_, err = Parse(records)
	if err == nil {
		t.Error("Parsed file with wrong header")
	}
	// test nil argument
	_, err = Parse(nil)
	if err == nil {
		t.Error("Parse nil argument")
	}
}

func Test_markCols(t *testing.T) {
	f, _ := os.Open("./testdata/test1.csv")
	defer f.Close()
	r := bufio.NewReader(f)
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	tg, err := markCols(records[0])
	if err != nil {
		t.Errorf("can not mark main headers: %v", err)
		t.FailNow()
	}

	// main columns from test csv
	exp := targCols{1, 2, 3, 14, 15, 17, 20}
	if *tg != exp {
		t.Errorf("markCols: Exp: %v, Got: %v", exp, *tg)
	}
}
