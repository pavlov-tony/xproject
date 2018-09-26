// +build cloud

package gcpcln

import (
	"context"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pavlov-tony/xproject/internal/db/pgcln"
	"github.com/pavlov-tony/xproject/pkg/cloud/gcptypes"
)

func Test_Client_NewClient(t *testing.T) {
	ctx := context.Background()
	// creating new db connection
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}
	// creting new client
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
	// closing db connection
	pgCln.Close()
	// closing client
	err = cln.Close()
	if err != nil {
		t.Errorf("Failed to close client: %v", err)
	}
}

func Test_Client_BucketsList(t *testing.T) {
	ctx := context.Background()
	// creating new db connection
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}
	// creting new client
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Error("Failed to create client:", err)
	}

	projectID := os.Getenv("APP_PROJECT_ID")

	bs, err := cln.BucketsList(projectID)
	if err != nil {
		t.Error("Expected err == nil, Got: ", err)
	}
	if len(bs) == 0 {
		t.Error("Buckets len = 0, exp > 0")
	}

	// closing db connection
	pgCln.Close()

	// closing client
	err = cln.Close()
	if err != nil {
		t.Errorf("Failed to close client: %v", err)
	}
}

func Test_Cleint_CsvObjctsList(t *testing.T) {
	ctx := context.Background()
	// creating new db connection
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}
	// creting new client
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
	bktName := os.Getenv("GCP_APP_PROJECT_BUCKET")
	prefix := ""
	objs, err := cln.CsvObjsList(bktName, prefix)

	if err != nil {
		t.Error("Failed to fetch buckets")
	}
	if len(objs) == 0 {
		t.Error("Objects len = 0, exp > 0")
	}

	// closing db connection
	pgCln.Close()

	// closing client
	err = cln.Close()
	if err != nil {
		t.Errorf("Failed to close client: %v", err)
	}
}

func Test_Client_csvObjectContent(t *testing.T) {
	ctx := context.Background()
	// creating new db connection
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}
	// creting new client
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
	objCont, err := cln.csvObjectContent(os.Getenv("APP_PROJECT_BUCKET"),
		os.Getenv("GCP_APP_PROJECT_CSV_OBJECT"))
	if err != nil {
		t.Error("Failed to fetch content", err)
	}
	if len(objCont) == 0 {
		t.Error("Got: Object content len == 0, Exp: len > 0")
	}

	// closing db connection
	pgCln.Close()

	// closing client
	err = cln.Close()
	if err != nil {
		t.Errorf("Failed to close client: %v", err)
	}
}

func Test_Client_makeReport(t *testing.T) {
	ctx := context.Background()
	// creating new db connection
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}

	// creting new client
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Error("Failed to create client:", err)
	}

	obj := gcptypes.FileMetadata{
		Name:   os.Getenv("GCP_APP_PROJECT_CSV_OBJECT"),
		Bucket: os.Getenv("GCP_APP_PROJECT_BUCKET")}

	rep, err := cln.makeReport(obj)
	if err != nil {
		t.Error("Failed to make report:", err)
	}

	if len(rep.Bills) == 0 {
		t.Error("Got: rep.Bills len == 0, Exp: len > 0")
	}

	// closing db connection
	pgCln.Close()

	// closing client
	err = cln.Close()
	if err != nil {
		t.Errorf("Failed to close client: %v", err)
	}
}

// func Test_Client_MakeReports(t *testing.T) {
// 	ctx := context.Background()
// 	// creating new db connection
// 	pgCln, err := pgcln.New(ctx, makeConfig())
// 	if err != nil {
// 		t.Errorf("Can not init new pg client: %v", err)
// 	}
//
// 	// creting new client
// 	cln, err := NewClient(ctx, pgCln)
// 	if err != nil {
// 		t.Error("Failed to create client:", err)
// 	}
//
// 	bktName := os.Getenv("GCP_APP_PROJECT_BUCKET")
// 	prefix := "test"
// 	objs, err := cln.CsvObjsList(bktName, prefix)
// 	if err != nil {
// 		t.Error("Failed to fetch objs:", err)
// 	}
//
// 	reps, err := cln.MakeReports(objs)
// 	if err != nil {
// 		t.Error("Failed to make reports:", err)
// 	}
//
// 	if len(reps) == 0 {
// 		t.Error("Got: reps len == 0, Exp: len > 0")
// 	}
//
// 	// closing db connection
// 	pgCln.Close()
//
// 	// closing client
// 	err = cln.Close()
// 	if err != nil {
// 		t.Errorf("Failed to close client: %v", err)
// 	}
// }

// makeConfig creates default config for pg client
func makeConfig() pgcln.Config {
	conf := pgcln.Config{
		Host:     os.Getenv(pgcln.EnvDBHost),
		Port:     os.Getenv(pgcln.EnvDBPort),
		DB:       os.Getenv(pgcln.EnvDBName),
		User:     os.Getenv(pgcln.EnvDBUser),
		Password: os.Getenv(pgcln.EnvDBPwd),
		SSLMode:  "disable",
	}
	return conf
}
