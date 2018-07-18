package pgcln

import (
	_ "github.com/lib/pq"
)

// // Test_Account tests inserting account into db, selecting and deleting it
// func Test_Account(t *testing.T) {
// 	conf := Config{
// 		Host:     os.Getenv(EnvDBHost),
// 		Port:     os.Getenv(EnvDBPort),
// 		DB:       os.Getenv(EnvDBName),
// 		User:     os.Getenv(EnvDBUser),
// 		Password: os.Getenv(EnvDBPwd),
// 		SSLMode:  "disable",
// 	}
//
// 	pgcln, err := New(conf)
// 	if err != nil {
// 		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
// 	}
// 	defer pgcln.Close()
//
// 	testAccount := Account{
// 		ID:             1,
// 		GcpAccountInfo: "testInfo",
// 	}
//
// 	if err := pgcln.InsertAccount(testAccount); err != nil {
// 		t.Fatalf("%v: insert account err: %v", pgcLogPref, err)
// 	}
//
// 	accounts, err := pgcln.SelectAccounts()
// 	if err != nil {
// 		t.Fatalf("%v: select accounts err: %v", pgcLogPref, err)
// 	}
//
// 	if len(accounts) != 2 {
// 		t.Fatalf("%v: incorrect selection, expected 2 accounts, not %v", pgcLogPref, len(accounts))
// 	}
//
// 	if strings.Compare(accounts[0].GcpAccountInfo, "testInfo") != 0 {
// 		t.Fatalf("%v: selected account's info doesn't match the test account", pgcLogPref)
// 	}
//
// 	if err := pgcln.deleteLastAccount(); err != nil {
// 		t.Fatalf("%v: delete last account err: %v", pgcLogPref, err)
// 	}
//
// 	accounts, err = pgcln.SelectAccounts()
// 	if err != nil {
// 		t.Fatalf("%v: select accounts err: %v", pgcLogPref, err)
// 	}
//
// 	if len(accounts) != 1 {
// 		t.Fatalf("%v: incorrect selection, expected 1 account, not %v", pgcLogPref, len(accounts))
// 	}
// }
//
// // Test_CsvFile tests inserting CSV file into db, selecting and deleting it
// func Test_CsvFile(t *testing.T) {
// 	conf := Config{
// 		Host:     os.Getenv(EnvDBHost),
// 		Port:     os.Getenv(EnvDBPort),
// 		DB:       os.Getenv(EnvDBName),
// 		User:     os.Getenv(EnvDBUser),
// 		Password: os.Getenv(EnvDBPwd),
// 		SSLMode:  "disable",
// 	}
//
// 	pgcln, err := New(conf)
// 	if err != nil {
// 		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
// 	}
// 	defer pgcln.Close()
//
// 	testCsvFile := GcpCsvFile{
// 		ID:          1,
// 		Name:        "testName",
// 		Bucket:      "testBucket",
// 		TimeCreated: time.Now(),
// 		AccountID:   1,
// 	}
//
// 	if err := pgcln.InsertCsvFile(testCsvFile); err != nil {
// 		t.Fatalf("%v: insert csv file err: %v", pgcLogPref, err)
// 	}
//
// 	csvFiles, err := pgcln.SelectCsvFiles()
// 	if err != nil {
// 		t.Fatalf("%v: select csv file err: %v", pgcLogPref, err)
// 	}
//
// 	if len(csvFiles) != 2 {
// 		t.Fatalf("%v: incorrect selection, expected 2 csv files, not %v", pgcLogPref, len(csvFiles))
// 	}
//
// 	if strings.Compare(csvFiles[0].Name, "testName") != 0 {
// 		t.Fatalf("%v: selected csv file's name doesn't match the test csv file", pgcLogPref)
// 	}
//
// 	if err := pgcln.deleteLastCsvFile(); err != nil {
// 		t.Fatalf("%v: delete last csv file err: %v", pgcLogPref, err)
// 	}
//
// 	csvFiles, err = pgcln.SelectCsvFiles()
// 	if err != nil {
// 		t.Fatalf("%v: select csv files err: %v", pgcLogPref, err)
// 	}
//
// 	if len(csvFiles) != 1 {
// 		t.Fatalf("%v: incorrect selection, expected 1 csv file, not %v", pgcLogPref, len(csvFiles))
// 	}
// }
//
// // Test_SelectBillsByTime checks if SelectBillsByTime() returns bills from non-empty db
// func Test_SelectBillsByTime(t *testing.T) {
// 	conf := Config{
// 		Host:     os.Getenv(EnvDBHost),
// 		Port:     os.Getenv(EnvDBPort),
// 		DB:       os.Getenv(EnvDBName),
// 		User:     os.Getenv(EnvDBUser),
// 		Password: os.Getenv(EnvDBPwd),
// 		SSLMode:  "disable",
// 	}
//
// 	pgcln, err := New(conf)
// 	if err != nil {
// 		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
// 	}
// 	defer pgcln.Close()
//
// 	start, err := time.Parse(time.RFC3339, "2000-01-01T00:00:00-07:00")
// 	if err != nil {
// 		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
// 	}
//
// 	end := time.Now()
//
// 	bills, err := pgcln.SelectBillsByTime(start, end)
// 	if err != nil {
// 		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
// 	}
//
// 	if len(bills) == 0 {
// 		t.Fatalf("%v: expected non-empty selection, but no bills were selected", pgcLogPref)
// 	}
//
// 	if strings.Compare(bills[0].LineItem, "test_service") != 0 {
// 		t.Fatalf("%v: selected bill's line item doesn't match 'test_service'", pgcLogPref)
// 	}
// }
//
// // Test_InsertBill_deleteLastBill tests inserting bill into db, selecting it by service and deleting it
// func Test_InsertBill_deleteLastBill(t *testing.T) {
// 	conf := Config{
// 		Host:     os.Getenv(EnvDBHost),
// 		Port:     os.Getenv(EnvDBPort),
// 		DB:       os.Getenv(EnvDBName),
// 		User:     os.Getenv(EnvDBUser),
// 		Password: os.Getenv(EnvDBPwd),
// 		SSLMode:  "disable",
// 	}
//
// 	pgcln, err := New(conf)
// 	if err != nil {
// 		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
// 	}
// 	defer pgcln.Close()
//
// 	testBill := ServiceBill{
// 		ID:           1,
// 		LineItem:     "testItem",
// 		StartTime:    time.Now(),
// 		EndTime:      time.Now(),
// 		Cost:         123.456,
// 		Currency:     "testCurrency",
// 		ProjectID:    "testProject",
// 		Description:  "testDescription",
// 		GcpCsvFileID: 1,
// 	}
//
// 	if err := pgcln.InsertBill(testBill); err != nil {
// 		t.Fatalf("%v: insert bill err: %v", pgcLogPref, err)
// 	}
//
// 	bills, err := pgcln.SelectBillsByService("estIte")
// 	if err != nil {
// 		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
// 	}
//
// 	if len(bills) != 1 {
// 		t.Fatalf("%v: incorrect selection, expected 1 bill, not %v", pgcLogPref, len(bills))
// 	}
//
// 	if strings.Compare(bills[0].LineItem, "testItem") != 0 || strings.Compare(bills[0].ProjectID, "testProject") != 0 {
// 		t.Fatalf("%v: selected bill's line item doesn't match the test bill", pgcLogPref)
// 	}
//
// 	if err := pgcln.deleteLastBill(); err != nil {
// 		t.Fatalf("%v: delete last bill err: %v", pgcLogPref, err)
// 	}
//
// 	bills, err = pgcln.SelectBillsByService("testItem")
// 	if err != nil {
// 		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
// 	}
//
// 	if len(bills) != 0 {
// 		t.Fatalf("%v: incorrect selection, expected 0 bills, not %v", pgcLogPref, len(bills))
// 	}
// }
//
// // Test_SelectLastBill checks if SelectLastBill() returns the last bill based on test data from db
// func Test_SelectLastBill(t *testing.T) {
// 	conf := Config{
// 		Host:     os.Getenv(EnvDBHost),
// 		Port:     os.Getenv(EnvDBPort),
// 		DB:       os.Getenv(EnvDBName),
// 		User:     os.Getenv(EnvDBUser),
// 		Password: os.Getenv(EnvDBPwd),
// 		SSLMode:  "disable",
// 	}
//
// 	pgcln, err := New(conf)
// 	if err != nil {
// 		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
// 	}
// 	defer pgcln.Close()
//
// 	bill, err := pgcln.SelectLastBill()
// 	if err != nil {
// 		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
// 	}
//
// 	if strings.Compare(bill.ProjectID, "test_project") != 0 {
// 		t.Fatalf("%v: selected bill's project id doesn't match 'test_project'", pgcLogPref)
// 	}
// }
