// TODO: I prefer shortage names like ubill if this does not bad for readability
// TODO: utils.go - file names must describe content, it cant be general word
// for example u can use fetch.go or others
package billingutils

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"time"
	"xproject/internal/cloud/gcp/utils/storageutils"

	"cloud.google.com/go/storage"
)

// NOTE: may be should fix structure
type ServiceBill struct {
	Description string // NOTE: Now field 17, may be should change
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
}

// NOTE: daily??
type ServicesBills []ServiceBill

// TODO: in pull req can't be commented unused code
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
	// TODO: is there no need check out of range
	res[0] = row[17]
	res[1] = row[2]
	res[2] = row[3]
	res[3] = row[11]
	res[4] = row[12]

	return res
}

// NOTE: must free slice before?
// NOTE: daily
// FIXME: handle errors
// TODO: write data into db
// TODO: fill by objects
// TODO:
// add ServiceBill to ServicesBills  csv.Reader

func (sbs *ServicesBills) addByReader(reader *csv.Reader) {
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

// Fill ServicesBills list by csv object, Clean list every time
func (sbs *ServicesBills) FillByObject(ctx context.Context, client *storage.Client,
	object *storageutils.Object) {
	*sbs = nil

	sbs.addByReader(object.NewCSVReader(ctx, client))
}

// Fill ServicesBills list by csv objects, Clean list every time
func (sbs *ServicesBills) FillByObjects(ctx context.Context, client *storage.Client,
	objects *storageutils.Objects) {
	*sbs = nil

	for _, o := range *objects {
		sbs.addByReader(o.NewCSVReader(ctx, client))
	}
}
