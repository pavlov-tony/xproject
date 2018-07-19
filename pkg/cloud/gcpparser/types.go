package gcpparser

import "time"

// targCols uses for store gcp csv billing file's columns numbers
type targCols struct {
	ColLineItem    int64 `gcpcsv:"Line Item"`
	ColStartTime   int64 `gcpcsv:"Start Time"`
	ColEndTime     int64 `gcpcsv:"End Time"`
	ColCost        int64 `gcpcsv:"Cost"`
	ColCurrency    int64 `gcpcsv:"Currency"`
	ColProjectID   int64 `gcpcsv:"Project ID"`
	ColDescription int64 `gcpcsv:"Description"`
}

// ServiceBill represent a composition of main service fields from gcp billing csv file
// for cost calculation tasks
type ServiceBill struct {
	LineItem    string
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
	ProjectID   string
	Description string
}

// ServicesBills is a set of ServiceBill
type ServicesBills []*ServiceBill
