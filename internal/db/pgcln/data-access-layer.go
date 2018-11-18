package pgcln

import (
	"context"
	"database/sql"
	"log"
)

const selAccounts string = "selectFromAccounts"
const selLastAccount string = "selectLastAccount"
const insAccount string = "insertIntoAccounts"
const delAccount string = "deleteFromAccounts"
const selFiles string = "selectFromCsvFiles"
const selLastFile string = "selectLastCsvFile"
const insFile string = "insertIntoCsvFiles"
const delFile string = "deleteFromCsvFiles"
const selBills string = "selectFromBills"
const selBillsByTime string = "selectBillsByTime"
const selBillsByService string = "selectBillsByService"
const selBillsByProject string = "selectBillsByProject"
const selLastBill string = "selectLastBill"
const insBill string = "insertIntoBills"
const delBill string = "deleteFromBills"

func (c *Client) prepareQueries(ctx context.Context) error {
	c.queries = make(map[string]*sql.Stmt)
	var err error

	c.queries[selAccounts], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.accounts ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selLastAccount], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.accounts ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[insAccount], err = c.idb.PrepareContext(ctx,
		"INSERT INTO cloudbilling.accounts VALUES(DEFAULT, $1)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[delAccount], err = c.idb.PrepareContext(ctx,
		"DELETE FROM cloudbilling.accounts WHERE id = (SELECT MAX(id) FROM cloudbilling.accounts)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selFiles], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.gcp_csv_files ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selLastFile], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.gcp_csv_files ORDER BY time_created DESC, id DESC LIMIT 1")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[insFile], err = c.idb.PrepareContext(ctx,
		"INSERT INTO cloudbilling.gcp_csv_files VALUES(DEFAULT, $1, $2, $3, $4)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[delFile], err = c.idb.PrepareContext(ctx,
		"DELETE FROM cloudbilling.gcp_csv_files WHERE id = (SELECT MAX(id) FROM cloudbilling.gcp_csv_files)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selBills], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.service_bills ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selBillsByTime], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.service_bills WHERE start_time >= $1 AND end_time <= $2 ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selBillsByService], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.service_bills WHERE line_item LIKE '%' || $1 || '%' ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selBillsByProject], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.service_bills WHERE project_id LIKE '%' || $1 || '%' ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[selLastBill], err = c.idb.PrepareContext(ctx,
		"SELECT * FROM cloudbilling.service_bills ORDER BY end_time DESC, start_time DESC, id DESC LIMIT 1")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[insBill], err = c.idb.PrepareContext(ctx,
		"INSERT INTO cloudbilling.service_bills VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries[delBill], err = c.idb.PrepareContext(ctx,
		"DELETE FROM cloudbilling.service_bills WHERE id = (SELECT MAX(id) FROM cloudbilling.service_bills)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	return err
}
