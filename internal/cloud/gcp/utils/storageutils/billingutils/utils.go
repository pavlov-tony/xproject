package billingutils

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"time"
)

// NOTE: may be should fix structure
type ServiceBill struct {
	Description string // NOTE: Now field 17, may be should change
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
}

type ServicesBills []ServiceBill

// type GCPTableAttributes [5]string

// Set in appropriate type ServiceBill attributes frow filtered row of billing table in bucket
func (sb *ServiceBill) setAttributes(row []string) {
	res := filterGCPTableRow(row)

	sb.Description = res[0]
	t, err := time.Parse(time.RFC3339, res[1])
	if err != nil {
		log.Fatal(err)
	}
	sb.StartTime = t
	t, err = time.Parse(time.RFC3339, res[2])
	if err != nil {
		log.Fatal(err)
	}
	sb.EndTime = t
	cost, err := strconv.ParseFloat(res[3], 64)
	if err != nil {
		log.Fatal(err)
	}
	sb.Cost = cost
	sb.Currency = res[4]
}

// Select appropriate columns from GCP billing table from bucket
func filterGCPTableRow(row []string) (res [5]string) {
	res[0] = row[17]
	res[1] = row[2]
	res[2] = row[3]
	res[3] = row[11]
	res[4] = row[12]

	return res
}

// NOTE: must free slice before?
// FIXME: handle errors
// TODO: write data into db
// fill ServicesBills by csv.Reader
func (sbs *ServicesBills) Fill(reader *csv.Reader) {
	reader.Read() // first time - read columns names

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		sb := ServiceBill{}
		sb.setAttributes(row)
		*sbs = append(*sbs, sb)
	}

}
