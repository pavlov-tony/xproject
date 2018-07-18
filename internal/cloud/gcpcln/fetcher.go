package gcpcln

import (
	"log"
	"time"
)

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, prefix string, dt time.Duration) {
	for {
		// pull all csv objects from bucket
		objs, err := c.CsvObjsList(bktName, prefix)
		if err != nil {
			log.Fatalf("in fetch c.CsvObjList: %v", err)
		}

		// select last report from db with its creation time

		lstObject, err := c.pgCln.SelLastGcpCsvObject()
		if err != nil {
			log.Printf("in fetch c.pgCln.SelLstGcpCsvObjectl: %v", err)
		}

		// filter CsvObjList, save only fresh objects
		if lstObject != nil {
			objs = objs.After(lstObject.Created)
		}

		// make Reports from object content
		reps, err := c.MakeReports(objs)
		_ = reps
		if err != nil {
			log.Fatalf("in fetch MakeReports: %v", err)
		}

		write Reports into db
		// c.pgCln.InsReps(reps)
		// c.pgCln.InsObjs(pgcln.Objects(make(Objects, 10)))

		time.Sleep(dt * time.Hour)
	}
}
