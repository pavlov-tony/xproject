// Package gcpparser need for parsing gcp billing csv data
// Also this package store
package gcpparser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/radisvaliullin/cloudbilling/pkg/cloud/gcptypes"
)

// necessary columns
const (
	colLineItem    = "Line Item"
	colStartTime   = "Start Time"
	colEndTime     = "End Time"
	colCost        = "Cost"
	colCurrency    = "Currency"
	colProjectID   = "Project ID"
	colDescription = "Description"
)

// array of necessary columns
var necesCols = [...]string{colLineItem, colStartTime, colEndTime, colCost, colCurrency, colProjectID, colDescription}

// Parse need for parse full GCP csv billing file
func Parse(data [][]string) (res gcptypes.ServicesBills, err error) {

	// check if input data == nil
	if data == nil {
		return nil, fmt.Errorf("parse: data == nil")
	}
	// check if input data has not columns
	if len(data) == 0 {
		return nil, fmt.Errorf("parse: data len == 0")
	}

	// mark necessary headers (targets)
	tg, err := markCols(data[0])
	if err != nil {
		return nil, fmt.Errorf("parse: error in markCols: %v", err)
	}

	// skip headers, handling content line by line
	for _, l := range data[1:] {
		sb, err := parseLine(l, tg)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		res = append(res, sb)
	}

	return res, nil
}

// parseLine need for parse row from GCP csv billing file
func parseLine(line []string, tg targCols) (*gcptypes.ServiceBill, error) {

	// parse startTime and endTime
	st, err := time.Parse(time.RFC3339, line[tg[colStartTime]])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse StartTime, %v", err)
	}
	et, err := time.Parse(time.RFC3339, line[tg[colEndTime]])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse EndTime, %v", err)
	}
	// parse cost
	cst, err := strconv.ParseFloat(line[tg[colCost]], 64)
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse Cost, %v", err)
	}

	// create new ServiceBill
	sb := gcptypes.ServiceBill{
		LineItem:    line[tg[colLineItem]],
		StartTime:   st,
		EndTime:     et,
		Cost:        cst,
		Currency:    line[tg[colCurrency]],
		ProjectID:   line[tg[colProjectID]],
		Description: line[tg[colDescription]],
	}

	return &sb, nil
}

// markCols takes headers of billing file and returns numbers of necessary columns
func markCols(headers []string) (targCols, error) {

	res := make(targCols)

	// match headers from file with necessary columns
	for i, h := range headers {
		for _, v := range necesCols {
			if h == v {
				res[v] = i
			}
		}
	}

	// if all necessary columns have not matched
	if len(res) != len(necesCols) {
		return nil, fmt.Errorf("can not mark all columns, marked %v", res)
	}

	return res, nil
}
