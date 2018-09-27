// +build cloud

package gcpcln

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/pavlov-tony/xproject/internal/db/pgcln"
)

func Test_Client_Fetch(t *testing.T) {
	ctx := context.Background()
	pgCln, err := pgcln.New(ctx, makeConfig())
	if err != nil {
		t.Errorf("Can not init new pg client: %v", err)
	}
	cln, err := NewClient(ctx, pgCln)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	bktName := os.Getenv("GCP_APP_PROJECT_BUCKET")
	prefix := "test"
	dt := time.Second * 10
	// NOTE: tester should add account with id 21 to db
	accountID := 21

	// run method fetch
	cln.Fetch(bktName, prefix, accountID, dt)
}
