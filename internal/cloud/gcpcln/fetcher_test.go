package gcpcln

import (
	"context"
	"os"
	"testing"
	"time"
)

func Test_Client_Fetch(t *testing.T) {
	ctx := context.Background()
	cln, err := NewClient(ctx)
	if err != nil {
		t.Error("Failed to create client:", err)
	}

	bktName := os.Getenv("APP_PROJECT_BUCKET")
	prefix := "test"
	dt := time.Second * 10

	cln.Fetch(bktName, prefix, dt)
}
