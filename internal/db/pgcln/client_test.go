package pgcln

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/radisvaliullin/cloudbilling/pkg/cloud/gcptypes"
)

// TestGcpAccount tests adding account into db, listing and removing it
func TestGcpAccount(t *testing.T) {
	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgcln, err := New(context.Background(), conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	prevLen := len(accounts)

	testAccount := GcpAccount{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	err = pgcln.AddAccount(testAccount)
	if err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastAccount()

	accounts, err = pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	if len(accounts)-prevLen != 1 {
		t.Fatalf("%v: expected 1 new account, not %v", pgcLogPref, len(accounts)-prevLen)
	}

	if strings.Compare(accounts[len(accounts)-1].GcpAccountInfo, "testInfo") != 0 {
		t.Fatalf("%v: account's info doesn't match the test account", pgcLogPref)
	}

	account, err := pgcln.GetLastAccount()
	if err != nil {
		t.Fatalf("%v: get last account err: %v", pgcLogPref, err)
	}

	if strings.Compare(account.GcpAccountInfo, "testInfo") != 0 {
		t.Fatalf("%v: last account's info doesn't match the test account", pgcLogPref)
	}
}

// TestFileMetadata tests adding file's metadata into db, listing and removing it
func TestFileMetadata(t *testing.T) {
	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgcln, err := New(context.Background(), conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	files, err := pgcln.ListFiles()
	if err != nil {
		t.Fatalf("%v: list files err: %v", pgcLogPref, err)
	}

	prevLen := len(files)

	testAccount := GcpAccount{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	err = pgcln.AddAccount(testAccount)
	if err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastAccount()

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	testFile1 := gcptypes.FileMetadata{
		ID:        1,
		Name:      "testName1",
		Bucket:    "testBucket1",
		Created:   time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID: accounts[len(accounts)-1].ID,
	}

	testFile2 := gcptypes.FileMetadata{
		ID:        2,
		Name:      "testName2",
		Bucket:    "testBucket2",
		Created:   time.Date(2077, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID: accounts[len(accounts)-1].ID,
	}

	err = pgcln.AddFile(testFile1)
	if err != nil {
		t.Fatalf("%v: add file err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastFile()

	err = pgcln.AddFile(testFile2)
	if err != nil {
		t.Fatalf("%v: add file err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastFile()

	files, err = pgcln.ListFiles()
	if err != nil {
		t.Fatalf("%v: list file err: %v", pgcLogPref, err)
	}

	if len(files)-prevLen != 2 {
		t.Fatalf("%v: expected 2 new files, not %v", pgcLogPref, len(files)-prevLen)
	}

	if strings.Compare(files[len(files)-1].Name, "testName2") != 0 {
		t.Fatalf("%v: file's name doesn't match the test file", pgcLogPref)
	}

	lastFile, err := pgcln.GetLastFile()
	if err != nil {
		t.Fatalf("%v: get last file err: %v", pgcLogPref, err)
	}

	if strings.Compare(lastFile.Bucket, "testBucket1") != 0 {
		t.Fatalf("%v: file's bucket doesn't match the last file", pgcLogPref)
	}
}

// TestServiceBill tests all functions that are related to adding,
// listing and removing service bills
func TestServiceBill(t *testing.T) {
	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgcln, err := New(context.Background(), conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	bills, err := pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list all bills err: %v", pgcLogPref, err)
	}

	prevLen := len(bills)

	testAccount := GcpAccount{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	err = pgcln.AddAccount(testAccount)
	if err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastAccount()

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	testFile := gcptypes.FileMetadata{
		ID:        1,
		Name:      "testName",
		Bucket:    "testBucket",
		Created:   time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID: accounts[len(accounts)-1].ID,
	}

	err = pgcln.AddFile(testFile)
	if err != nil {
		t.Fatalf("%v: add file err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastFile()

	files, err := pgcln.ListFiles()
	if err != nil {
		t.Fatalf("%v: list files err: %v", pgcLogPref, err)
	}

	testBill1 := gcptypes.ServiceBill{
		ID:             1,
		LineItem:       "testItem1",
		StartTime:      time.Date(2077, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:        time.Date(2077, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:           123.456,
		Currency:       "testCurrency1",
		ProjectID:      "testProject1",
		Description:    "testDescription1",
		FileMetadataID: files[len(files)-1].ID,
	}

	testBill2 := gcptypes.ServiceBill{
		ID:             2,
		LineItem:       "testItem2",
		StartTime:      time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:        time.Date(2078, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:           456.789,
		Currency:       "testCurrency2",
		ProjectID:      "testProject2",
		Description:    "testDescription2",
		FileMetadataID: files[len(files)-1].ID,
	}

	err = pgcln.AddBill(testBill1)
	if err != nil {
		t.Fatalf("%v: add bill err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastBill()

	err = pgcln.AddBill(testBill2)
	if err != nil {
		t.Fatalf("%v: add bill err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastBill()

	bills, err = pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list all bills err: %v", pgcLogPref, err)
	}

	if len(bills)-prevLen != 2 {
		t.Fatalf("%v: expected 2 new bills, not %v", pgcLogPref, len(bills)-prevLen)
	}

	lastBill, err := pgcln.GetLastBill()
	if err != nil {
		t.Fatalf("%v: get last bill err: %v", pgcLogPref, err)
	}

	if strings.Compare(testBill2.LineItem, lastBill.LineItem) != 0 ||
		strings.Compare(testBill2.ProjectID, lastBill.ProjectID) != 0 {
		t.Fatalf("%v: last bill doesn't match the test bill", pgcLogPref)
	}

	bills, err = pgcln.ListBillsByService("testItem")
	if err != nil {
		t.Fatalf("%v: list bills by service err: %v", pgcLogPref, err)
	}

	if len(bills) != 2 {
		t.Fatalf("%v: expected 2 bills, not %v", pgcLogPref, len(bills))
	}

	bills, err = pgcln.ListBillsByTime(time.Date(2077, 12, 31, 0, 0, 0, 0, time.Local),
		time.Date(2078, 1, 15, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.Fatalf("%v: list bills by time err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(testBill2.Description, bills[0].Description) != 0 {
		t.Fatalf("%v: bill's description doesn't match the test bill", pgcLogPref)
	}

	bills, err = pgcln.ListBillsByProject("testProject1")
	if err != nil {
		t.Fatalf("%v: list bills by project err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(testBill1.Description, bills[0].Description) != 0 {
		t.Fatalf("%v: bill's description doesn't match the test bill", pgcLogPref)
	}
}

// TestReport tests adding report (file's metadata and bills) into db
func TestReport(t *testing.T) {
	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgcln, err := New(context.Background(), conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	files, err := pgcln.ListFiles()
	if err != nil {
		t.Fatalf("%v: list files err: %v", pgcLogPref, err)
	}

	prevLen1 := len(files)

	bills, err := pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	prevLen2 := len(bills)

	testAccount := GcpAccount{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	err = pgcln.AddAccount(testAccount)
	if err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastAccount()

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	testFile := gcptypes.FileMetadata{
		ID:        1,
		Name:      "testName",
		Bucket:    "testBucket",
		Created:   time.Date(2077, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID: accounts[len(accounts)-1].ID,
	}

	testBill1 := gcptypes.ServiceBill{
		ID:             1,
		LineItem:       "testItem1",
		StartTime:      time.Date(2077, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:        time.Date(2077, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:           123.456,
		Currency:       "testCurrency1",
		ProjectID:      "testProject1",
		Description:    "testDescription1",
		FileMetadataID: testFile.ID,
	}

	testBill2 := gcptypes.ServiceBill{
		ID:             2,
		LineItem:       "testItem2",
		StartTime:      time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:        time.Date(2078, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:           456.789,
		Currency:       "testCurrency2",
		ProjectID:      "testProject2",
		Description:    "testDescription2",
		FileMetadataID: testFile.ID,
	}

	report := gcptypes.Report{
		Metadata: testFile,
		Bills:    gcptypes.ServicesBills{&testBill1, &testBill2},
	}

	err = pgcln.AddReportsToAccount(gcptypes.Reports{&report}, accounts[len(accounts)-1].ID)
	if err != nil {
		t.Fatalf("%v: add reports err: %v", pgcLogPref, err)
	}
	defer pgcln.removeLastFile()
	defer pgcln.removeLastBill()
	defer pgcln.removeLastBill()

	files, err = pgcln.ListFiles()
	if err != nil {
		t.Fatalf("%v: list file err: %v", pgcLogPref, err)
	}

	if len(files)-prevLen1 != 1 {
		t.Fatalf("%v: expected 1 new file, not %v", pgcLogPref, len(files)-prevLen1)
	}

	if strings.Compare(files[len(files)-1].Name, "testName") != 0 {
		t.Fatalf("%v: file's name doesn't match the test file", pgcLogPref)
	}

	bills, err = pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if len(bills)-prevLen2 != 2 {
		t.Fatalf("%v: expected 2 new bills, not %v", pgcLogPref, len(bills)-prevLen2)
	}

	if strings.Compare(bills[len(bills)-1].LineItem, "testItem2") != 0 {
		t.Fatalf("%v: bill's line item doesn't match the test bill", pgcLogPref)
	}
}
