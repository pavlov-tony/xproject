package gcptypes

import (
	"time"
)

// Object represents csv objects (files) from gcp
type Object struct {
	Id        int
	Name      string
	Bucket    string
	Created   time.Time
	AccountID int
}

// Objects is a list of Object structures
type Objects []Object

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

// TODO:
type Report struct {
	Object Object
	Bills  ServicesBills
}

// TODO:
type Reports []*Report

// after filters object, select only objects which after t
func (objs Objects) After(t time.Time) (res Objects) {
	for _, o := range objs {
		if o.Created.After(t) {
			res = append(res, o)
		}
	}

	return res
}
