package gcpcln

import (
	"context"
	"os"
	"testing"
)

func Test_Client_NewClient(t *testing.T) {
	ctx := context.Background()
	_, err := NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
}

func Test_Client_BucketsList(t *testing.T) {
	ctx := context.Background()
	cln, err := NewClient(ctx)

	projectID := os.Getenv("APP_PROJECT_ID")

	bs, err := cln.BucketsList(projectID)
	if err != nil {
		t.Error("Expected err == nil, Got: ", err)
	}
	if len(bs) == 0 {
		t.Error("Buckets len = 0, exp > 0")
	}
}

func Test_Cleint_CsvObjctsList(t *testing.T) {
	ctx := context.Background()
	cln, err := NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
	bktName := os.Getenv("APP_PROJECT_BUCKET")
	prefix := ""
	objs, err := cln.CsvObjectsList(bktName, prefix)
	if err != nil {
		t.Error("Failed to fetch buckets")
	}
	if len(objs) == 0 {
		t.Error("Objects len = 0, exp > 0")
	}
}

func Test_Client_CsvObjectContent(t *testing.T) {
	ctx := context.Background()
	cln, err := NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client:", err)
	}
	objCont, err := cln.CsvObjectContent(os.Getenv("APP_PROJECT_BUCKET"),
		os.Getenv("APP_PROJECT_CSV_OBJECT"))
	if err != nil {
		t.Error("Failed to fetch content", err)
	}
	if len(objCont) == 0 {
		t.Error("Object content len == 0, exp > 0")
	}
}
