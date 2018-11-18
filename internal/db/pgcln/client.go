package pgcln

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/radisvaliullin/cloudbilling/pkg/cloud/gcptypes"
	// Don't forget add driver importing to main
	// _ "github.com/lib/pq"
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

	// prepared statements
	queries map[string]*sql.Stmt
}

// New inits client
func New(ctx context.Context, conf Config) (*Client, error) {
	c := &Client{
		conf: conf,
	}

	// get db
	dbSourceName := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v sslmode=%v",
		conf.Host, conf.Port, conf.DB, conf.User, conf.Password, conf.SSLMode,
	)
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		log.Printf("%v: db open err, %v", pgcLogPref, err)
		return nil, err
	}

	// set db interface
	c.idb = db

	// prepare queries
	err = c.prepareQueries(ctx)
	if err != nil {
		log.Printf("%v: prepare queries err, %v", pgcLogPref, err)
		return nil, err
	}

	return c, nil
}

// combineAccounts combines rows of query result into GcpAccounts
func combineAccounts(rows *sql.Rows) (GcpAccounts, error) {
	var table GcpAccounts
	var row GcpAccount

	for rows.Next() {
		err := rows.Scan(&row.ID, &row.GcpAccountInfo)
		if err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	err := rows.Err()
	if err != nil {
		log.Printf("%v: db rows err, %v", pgcLogPref, err)
		return nil, err
	}

	return table, nil
}

// combineFiles combines rows of query result into FilesMetadata
func combineFiles(rows *sql.Rows) (gcptypes.FilesMetadata, error) {
	var table gcptypes.FilesMetadata
	var row gcptypes.FileMetadata

	for rows.Next() {
		err := rows.Scan(
			&row.ID, &row.Name, &row.Bucket, &row.Created, &row.AccountID,
		)
		if err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	err := rows.Err()
	if err != nil {
		log.Printf("%v: db rows err, %v", pgcLogPref, err)
		return nil, err
	}

	return table, nil
}

// combineBills combines rows of query result into ServicesBills
func combineBills(rows *sql.Rows) (gcptypes.ServicesBills, error) {
	var table gcptypes.ServicesBills
	var row gcptypes.ServiceBill

	for rows.Next() {
		err := rows.Scan(
			&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost,
			&row.Currency, &row.ProjectID, &row.Description, &row.FileMetadataID,
		)
		if err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	err := rows.Err()
	if err != nil {
		log.Printf("%v: db rows err, %v", pgcLogPref, err)
		return nil, err
	}

	return table, nil
}

// Close releases db resources
func (c *Client) Close() error {
	for _, stmt := range c.queries {
		stmt.Close()
	}

	return c.idb.Close()
}

// Ping tests ping to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}

// ListAccounts returns all accounts from db
func (c *Client) ListAccounts() (GcpAccounts, error) {
	rows, err := c.queries[selAccounts].Query()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineAccounts(rows)
}

// GetLastAccount returns the latest added account from db by max id
func (c *Client) GetLastAccount() (*GcpAccount, error) {
	var row GcpAccount

	err := c.queries[selLastAccount].QueryRow().Scan(
		&row.ID, &row.GcpAccountInfo,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v: db scan err, %v", pgcLogPref, err)
		return nil, err
	}

	return &row, nil
}

// AddAccount adds account into db
func (c *Client) AddAccount(account GcpAccount) error {
	_, err := c.queries[insAccount].Exec(account.GcpAccountInfo)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// removeLastAccount removes the latest added account from db
func (c *Client) removeLastAccount() error {
	_, err := c.queries[delAccount].Exec()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// ListFiles returns all files' metadata from db
func (c *Client) ListFiles() (gcptypes.FilesMetadata, error) {
	rows, err := c.queries[selFiles].Query()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineFiles(rows)
}

// GetLastFile returns last file's metadata from db by time
func (c *Client) GetLastFile() (*gcptypes.FileMetadata, error) {
	var row gcptypes.FileMetadata

	err := c.queries[selLastFile].QueryRow().Scan(
		&row.ID, &row.Name, &row.Bucket, &row.Created, &row.AccountID,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v: db scan err, %v", pgcLogPref, err)
		return nil, err
	}

	return &row, nil
}

// AddFile adds file's metadata into db
func (c *Client) AddFile(file gcptypes.FileMetadata) error {
	_, err := c.queries[insFile].Exec(
		file.Name, file.Bucket, file.Created, file.AccountID,
	)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// removeLastFile removes the latest added file's metadata from db
func (c *Client) removeLastFile() error {
	_, err := c.queries[delFile].Exec()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// ListAllBills returns all bills from db
func (c *Client) ListAllBills() (gcptypes.ServicesBills, error) {
	rows, err := c.queries[selBills].Query()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByTime returns bills from db that are within the specified time period
func (c *Client) ListBillsByTime(start, end time.Time) (gcptypes.ServicesBills, error) {
	if start.After(end) || end.Before(start) {
		return nil, fmt.Errorf("%v: invalid arguments err", pgcLogPref)
	}

	rows, err := c.queries[selBillsByTime].Query(start, end)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByService returns bills from db that are related to specified GCP service
// If service is an empty string then all bills will be returned
func (c *Client) ListBillsByService(service string) (gcptypes.ServicesBills, error) {
	rows, err := c.queries[selBillsByService].Query(service)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByProject returns bills from db that are related to specified GCP project
// If project is an empty string then all bills will be returned
func (c *Client) ListBillsByProject(project string) (gcptypes.ServicesBills, error) {
	rows, err := c.queries[selBillsByProject].Query(project)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// GetLastBill returns the latest added bill from db by time
func (c *Client) GetLastBill() (*gcptypes.ServiceBill, error) {
	var row gcptypes.ServiceBill

	err := c.queries[selLastBill].QueryRow().Scan(
		&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost,
		&row.Currency, &row.ProjectID, &row.Description, &row.FileMetadataID,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}

	return &row, nil
}

// AddBill adds bill into db
func (c *Client) AddBill(bill gcptypes.ServiceBill) error {
	_, err := c.queries[insBill].Exec(
		bill.LineItem, bill.StartTime, bill.EndTime, bill.Cost,
		bill.Currency, bill.ProjectID, bill.Description, bill.FileMetadataID,
	)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// AddReport adds file's metadata and bills into db
func (c *Client) AddReport(report gcptypes.Report) error {
	err := c.AddFile(report.Metadata)
	if err != nil {
		log.Printf("%v: add file err, %v", pgcLogPref, err)
		return err
	}

	files, err := c.ListFiles()
	if err != nil {
		log.Printf("%v: list files err, %v", pgcLogPref, err)
		return err
	}

	for _, bill := range report.Bills {
		bill.FileMetadataID = files[len(files)-1].ID

		err := c.AddBill(*bill)
		if err != nil {
			log.Printf("%v: add bill err, %v", pgcLogPref, err)
			return err
		}
	}

	return nil
}

// AddReportsToAccount adds a set of reports of specified account into db
func (c *Client) AddReportsToAccount(reports gcptypes.Reports, accountID int) error {
	accounts, err := c.ListAccounts()
	if err != nil {
		log.Printf("%v: list accounts err, %v", pgcLogPref, err)
		return err
	}

	accountExists := false

	for _, account := range accounts {
		if account.ID == accountID {
			accountExists = true
			break
		}
	}

	if !accountExists {
		return fmt.Errorf("%v: specified account doesn't exist", pgcLogPref)
	}

	for _, report := range reports {
		report.Metadata.AccountID = accountID

		err := c.AddReport(*report)
		if err != nil {
			log.Printf("%v: add report err, %v", pgcLogPref, err)
			return err
		}
	}

	return nil
}

// removeLastBill removes the latest added bill from db
func (c *Client) removeLastBill() error {
	_, err := c.queries[delBill].Exec()
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}
