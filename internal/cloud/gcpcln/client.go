package gcpcln

import (
	"context"
	"encoding/csv"
	"errors"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// mime content type
const contentTypeCsv = "text/csv"

// Common client for GCP predictions which stores google api as fileds
type Client struct {
	strgCln *storage.Client
	ctx     context.Context
}

// Creating new client
func NewClient(ctx context.Context) (*Client, error) {
	c := new(Client)
	strgCln, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.New("NewClient: " + err.Error())
	}
	c.ctx = ctx
	c.strgCln = strgCln

	return c, nil
}

// fetch bucket list from project
func (c *Client) BucketsList(projectID string) (buckets []string, err error) {

	it := c.strgCln.Buckets(c.ctx, projectID)
	for {
		b, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.New("FetchProjectBuckets: " + err.Error())
		}
		buckets = append(buckets, b.Name)
	}

	return buckets, nil
}

// fetch csv ojects list from bucket with prefix
func (c *Client) CsvObjectsList(bktName, prefix string) (objs []string, err error) {

	it := c.strgCln.Bucket(bktName).Objects(c.ctx, &storage.Query{Prefix: prefix})
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.New("FetchBucketObjects: " + err.Error())
		}
		if o.ContentType == contentTypeCsv {
			objs = append(objs, o.Name)
		}
	}

	return objs, nil
}

// fetch data from bucket object by bucket name and object name
func (c *Client) CsvObjectContent(bktName, objName string) ([][]string, error) {
	r, err := c.strgCln.Bucket(bktName).Object(objName).NewReader(c.ctx)
	if err != nil {
		return nil, errors.New("fetchObjectCSVData: " + err.Error())
	}
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, errors.New("fetchObjectCSVData: " + err.Error())
	}

	return records, nil
}
