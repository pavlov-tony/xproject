package pgcln

import (
	"fmt"
	"os"
	"testing"
)

func Test_Client_SelLastGcpCsvObject(t *testing.T) {
	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgCln, err := New(conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgCln.Close()

	res, err := pgCln.SelLastGcpCsvObject()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
