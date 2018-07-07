package gcpparser

import (
	"fmt"
	"strconv"
	"time"
)

// Parse full GCP csv billing file
func Parse(data [][]string) (res ServicesBills, err error) {
	for _, l := range data {
		sb, err := parseLine(l)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		res = append(res, sb)
	}

	return res, nil
}

// Parse row from GCP csv billing file
func parseLine(line []string) (*ServiceBill, error) {
	if len(line) < MaxColNum {
		return nil, fmt.Errorf("parse line: line length < MaxColNum")
	}

	st, err := time.Parse(time.RFC3339, line[ColStartTime])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse StartTime, %v", err)
	}
	et, err := time.Parse(time.RFC3339, line[ColEndTime])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse EndTime, %v", err)
	}
	cst, err := strconv.ParseFloat(line[ColCost], 64)
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse Cost, %v", err)
	}

	sb := ServiceBill{
		Item:    line[ColLineItem],
		Started: st,
		Ended:   et,
		Cost:    cst,
		Curr:    line[ColCurrency],
		ProjID:  line[ColProjectID],
		Descr:   line[ColDescription],
	}

	return &sb, nil
}
