package pgcln

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// Don't forget add driver importing to main
	_ "github.com/lib/pq"
)

const (
	pgcLogPref = "postgres client"
)

// Env for testing
const (
	EnvDBHost = "APP_DB_PG_HOST"
	EnvDBPort = "APP_DB_PG_PORT"
	EnvDBName = "APP_DB_PG_NAME"
	EnvDBUser = "APP_DB_PG_USER"
	EnvDBPwd  = "APP_DB_PG_PWD"
)

// Config sets database configs
type Config struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
}

// Client implements postgres db client
type Client struct {
	// config
	conf Config

	// db
	idb IDB
}

// Account contains information about GCP user account
type Account struct {
	ID             int
	GcpAccountInfo string
}

// GcpCsvFile contains information about CSV files with billing reports
type GcpCsvFile struct {
	ID          int
	Name        string
	Bucket      string
	TimeCreated time.Time
	AccountID   int
}

// ServiceBill contains relevant information from billing report
type ServiceBill struct {
	ID           int
	LineItem     string
	StartTime    time.Time
	EndTime      time.Time
	Cost         float64
	Currency     string
	ProjectID    string
	Description  string
	GcpCsvFileID int
}

// Accounts is a set of Account
type Accounts []*Account

// GcpCsvFiles is a set of GcpCsvFile
type GcpCsvFiles []*GcpCsvFile

// ServiceBills is a set of ServiceBill
type ServiceBills []*ServiceBill

// New inits client
func New(conf Config) (*Client, error) {
	c := &Client{
		conf: conf,
	}

	// get db
	// FIXME: how to add empty password
	// dbSourceName := fmt.Sprintf(
	// 	"host=%v port=%v dbname=%v user=%v password=%v sslmode=%v",
	// 	conf.Host, conf.Port, conf.DB, conf.User, conf.Password, conf.SSLMode,
	// )
	dbSourceName := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v sslmode=%v",
		conf.Host, conf.Port, conf.DB, conf.User, conf.SSLMode,
	)
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		log.Printf("%v: db open err, %v", pgcLogPref, err)
		return nil, err
	}

	// set db interface
	c.idb = db

	return c, nil
}

// Close releases db resources
func (c *Client) Close() error {
	return c.idb.Close()
}

// Ping tests ping to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}

// SelectAccounts returns all accounts from db
func (c *Client) SelectAccounts() (Accounts, error) {
	rows, err := c.idb.Query("SELECT * FROM xproject.accounts ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	var table Accounts
	var row Account

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.GcpAccountInfo); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// InsertAccount inserts account into db
func (c *Client) InsertAccount(account Account) error {
	if _, err := c.idb.Query("INSERT INTO xproject.accounts VALUES(DEFAULT, $1)",
		account.GcpAccountInfo); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// deleteLastAccount deletes the latest added account from db
func (c *Client) deleteLastAccount() error {
	if _, err := c.idb.Query("DELETE FROM xproject.accounts WHERE id = (SELECT MAX(id) FROM xproject.accounts)"); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// SelectCsvFiles returns all CSV files from db
func (c *Client) SelectCsvFiles() (GcpCsvFiles, error) {
	rows, err := c.idb.Query("SELECT * FROM xproject.gcp_csv_files ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	var table GcpCsvFiles
	var row GcpCsvFile

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.Name, &row.Bucket, &row.TimeCreated, &row.AccountID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// InsertCsvFile inserts CSV file into db
func (c *Client) InsertCsvFile(file GcpCsvFile) error {
	if _, err := c.idb.Query("INSERT INTO xproject.gcp_csv_files VALUES(DEFAULT, $1, $2, $3, $4)",
		file.Name, file.Bucket, file.TimeCreated, file.AccountID); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// deleteLastCsvFile deletes the latest added CSV file from db
func (c *Client) deleteLastCsvFile() error {
	if _, err := c.idb.Query("DELETE FROM xproject.gcp_csv_files WHERE id = (SELECT MAX(id) FROM xproject.gcp_csv_files)"); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// SelectBillsByTime returns bills from db that belong to specified time period
func (c *Client) SelectBillsByTime(start, end time.Time) (ServiceBills, error) {
	if start.After(end) || end.Before(start) {
		return nil, fmt.Errorf("%v: invalid arguments err", pgcLogPref)
	}

	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills WHERE start_time >= $1 AND end_time <= $2 ORDER BY id ASC", start, end)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	var table ServiceBills
	var row ServiceBill

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost, &row.Currency, &row.ProjectID, &row.Description, &row.GcpCsvFileID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// SelectBillsByService returns bills from db that are related to specified GCP service
// If service is an empty string then all bills will be returned
func (c *Client) SelectBillsByService(service string) (ServiceBills, error) {
	service = "%" + service + "%"

	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills WHERE line_item LIKE $1 ORDER BY id ASC", service)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	var table ServiceBills
	var row ServiceBill

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost, &row.Currency, &row.ProjectID, &row.Description, &row.GcpCsvFileID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// SelectLastBill returns the latest added bill from db
func (c *Client) SelectLastBill() (ServiceBill, error) {
	var row ServiceBill

	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return row, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost, &row.Currency, &row.ProjectID, &row.Description, &row.GcpCsvFileID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return row, err
		}
	}

	return row, nil
}

// InsertBill inserts bill into db
func (c *Client) InsertBill(bill ServiceBill) error {
	if _, err := c.idb.Query("INSERT INTO xproject.service_bills VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)",
		bill.LineItem, bill.StartTime, bill.EndTime, bill.Cost, bill.Currency, bill.ProjectID, bill.Description, bill.GcpCsvFileID); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// deleteLastBill deletes the latest added bill from db
func (c *Client) deleteLastBill() error {
	if _, err := c.idb.Query("DELETE FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)"); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}
