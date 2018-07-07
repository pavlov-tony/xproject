package gcpparser

import "time"

// gcp csv billing file's columns numbers
const (
	ColLineItem    = 1
	ColStartTime   = 2
	ColEndTime     = 3
	ColCost        = 14
	ColCurrency    = 15
	ColProjectID   = 17
	ColDescription = 20
	MaxColNum      = 20
)

// Struct represent a composition of main service fields from gcp billing csv file
// for cost calculation tasks
type ServiceBill struct {
	Item    string    // column number: 1
	Started time.Time // 2
	Ended   time.Time // 3
	Cost    float64   // 14
	Curr    string    // 15
	ProjID  string    // 17
	Descr   string    // 20
}

// Set of ServiceBill
type ServicesBills []*ServiceBill
