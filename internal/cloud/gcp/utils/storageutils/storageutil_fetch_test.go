// TODO: test file usually build as testfilename_test.go
package storageutils

import (
	"context"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/storage"
)

func TestFetchProjectBuckets(t *testing.T) {
	ctx := context.Background()
	projectID := os.Getenv("APP_PROJECT_ID")
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client:", err)
	}

	resBuckets, err := FetchProjectBuckets(ctx, client, projectID)

	expBuckets := []string{"churomann-bucket", "churomann-bucket-2"}
	if !isStringSlicesEqual(resBuckets, expBuckets) {
		t.Error("IsStringSlicesEqual\nresult:", resBuckets, "\nexpected:", expBuckets)
	}
}

func isStringSlicesEqual(sl1 []string, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for _, s := range sl2 {
		if !isStringSliceContains(sl1, s) {
			return false
		}
	}
	return true
}

// IDEA: generic?
func isStringSliceContains(sl []string, str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}

	return false
}

func TestFetchBucketObjects(t *testing.T) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client: ", err)
	}

	resObjects, err := FetchBucketObjects(ctx, client, "churomann-bucket-2")

	expObjects := []Object{
		Object{"test1", "churomann-bucket-2", time.Time{}},
		Object{"test2", "churomann-bucket-2", time.Time{}}}
	if !isObjectSlicesEqual(resObjects, expObjects) {
		t.Error("IsStringSlicesEqual\nresult:", resObjects, "\nexpected:", expObjects)
	}
}

func isObjectSlicesEqual(sl1 []Object, sl2 []Object) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for _, o := range sl2 {
		if !isObjectSliceContains(sl1, o) {
			return false
		}
	}
	return true
}

// IDEA: generic?
func isObjectSliceContains(sl []Object, obj Object) bool {
	for _, o := range sl {
		if o.Name == obj.Name {
			return true
		}
	}
	return false
}
