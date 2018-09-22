package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Client represents the AWS Client to fetch different kinds of billing reports
type Client struct {
	ctx  context.Context
	Meta *CaurMeta
}

// NewClient creates the new AWS Client
func NewClient(ctx context.Context, awsconfig *aws.Config, reportName string) *Client {
	c := new(Client)

	c.ctx = ctx
	c.Meta = NewCaurMeta(awsconfig, reportName)

	return c
}
