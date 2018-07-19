package gcpparser

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Parse need for parse full GCP csv billing file
func Parse(data [][]string) (res ServicesBills, err error) {
	if data == nil {
		return nil, fmt.Errorf("parse: data == nil")
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("parse: data len == 0")
	}

	// mark header targets
	tg, err := markCols(data[0])
	if err != nil {
		return nil, fmt.Errorf("parse: error in markCols: %v", err)
	}
	// skip headers
	data = data[1:]

	for _, l := range data {
		sb, err := parseLine(l, tg)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		res = append(res, sb)
	}

	return res, nil
}

// parseLine need for parse row from GCP csv billing file
func parseLine(line []string, tg *targCols) (*ServiceBill, error) {

	st, err := time.Parse(time.RFC3339, line[tg.ColStartTime])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse StartTime, %v", err)
	}
	et, err := time.Parse(time.RFC3339, line[tg.ColEndTime])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse EndTime, %v", err)
	}
	cst, err := strconv.ParseFloat(line[tg.ColCost], 64)
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse Cost, %v", err)
	}

	sb := ServiceBill{
		LineItem:    line[tg.ColLineItem],
		StartTime:   st,
		EndTime:     et,
		Cost:        cst,
		Currency:    line[tg.ColCurrency],
		ProjectID:   line[tg.ColProjectID],
		Description: line[tg.ColDescription],
	}

	return &sb, nil
}

// markCols uses for mark main columns in csv file
func markCols(headers []string) (*targCols, error) {

	tg := new(targCols)
	val := reflect.ValueOf(tg).Elem()
	valType := reflect.ValueOf(*tg).Type()
	numField := reflect.ValueOf(*tg).NumField()

	// compare headers with struct fields
	for i := 0; i < numField; i++ {
		for j, h := range headers {
			if valType.Field(i).Tag.Get("gcpcsv") == h {
				val.Field(i).SetInt(int64(j))
				// break if field has found
				break
			}
			// if field has not found it is a error
			if j == len(headers)-1 {
				return nil, fmt.Errorf("can not find field %v",
					valType.Field(i).Name)
			}
		}
	}

	return tg, nil
}
