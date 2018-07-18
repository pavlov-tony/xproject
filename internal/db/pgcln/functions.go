package pgcln

import (
	"fmt"

	"github.com/pavlov-tony/xproject/pkg/cloud/gcptypes"
	// Don't forget add driver importing to main
	_ "github.com/lib/pq"
)

func (c *Client) SelLastGcpCsvObject() (*gcptypes.Object, error) {

	var obj gcptypes.Object
	err := c.idb.QueryRow("SELECT * FROM xproject.sel_last_report()").Scan(&obj.Id, &obj.Name, &obj.Bucket, &obj.Created, &obj.AccountID)
	if err != nil {
		return nil, fmt.Errorf("in pgcln SelLastGcpCsvObject: %v", err)
	}

	return &obj, nil
}

func (c *Client) insRep(rep gcptypes.Report) error {
	return nil
}

func (c *Client) InsReps(reps gcptypes.Reports) error {
	// insRep in loop
	return nil
}

func (c *Client) InsObjs(objs gcptypes.Objects) error {
	return nil
}
