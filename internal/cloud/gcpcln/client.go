// Package gcpcln periodically fetches csv from gcp bucket and writes in into db
// Algorithm:
// select client
// select Buckets
// choose bucket
// fetch from bucket obj with prefix since

package gcpcln

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/pavlov-tony/xproject/internal/db/pgcln"
	"github.com/pavlov-tony/xproject/pkg/cloud/gcpparser"
	"github.com/pavlov-tony/xproject/pkg/cloud/gcptypes"
	"google.golang.org/api/iterator"
)

// mime content type
const contentTypeCsv = "text/csv"

// Client for GCP predictions which stores google api as fileds
type Client struct {
	strgCln *storage.Client
	pgCln   *pgcln.Client
	ctx     context.Context
}

type Report struct {
	Object gcptypes.Object
	Bills  gcptypes.ServicesBills
}

type Reports []*Report

// NewClient creates new client for GCP
func NewClient(ctx context.Context) (*Client, error) {
	c := new(Client)
	strgCln, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("new client: %v", err)
	}
	c.ctx = ctx
	c.strgCln = strgCln

	// init new pg client
	conf := pgcln.Config{
		Host:     os.Getenv(pgcln.EnvDBHost),
		Port:     os.Getenv(pgcln.EnvDBPort),
		DB:       os.Getenv(pgcln.EnvDBName),
		User:     os.Getenv(pgcln.EnvDBUser),
		Password: os.Getenv(pgcln.EnvDBPwd),
		SSLMode:  "disable",
	}
	c.pgCln, err = pgcln.New(conf)
	if err != nil {
		log.Fatalf("in fetch pgcln.New: %v", err)
	}

	return c, nil
}

// BucketsList fetches bucket list from project
func (c *Client) BucketsList(projectID string) (buckets []string, err error) {

	it := c.strgCln.Buckets(c.ctx, projectID)
	for {
		b, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("fetch project buckets: %v", err)
		}
		buckets = append(buckets, b.Name)
	}

	return buckets, nil
}

// CsvObjectsList fetches csv ojects list from bucket with prefix
func (c *Client) CsvObjsList(bktName, prefix string) (objs gcptypes.Objects, err error) {

	it := c.strgCln.Bucket(bktName).Objects(c.ctx, &storage.Query{Prefix: prefix})
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("fetch bucket objects: %v", err)
		}
		if o.ContentType == contentTypeCsv {
			objs = append(objs, gcptypes.Object{Name: o.Name, Bucket: o.Bucket, Created: o.Created})
		}
	}

	return objs, nil
}

// CsvObjectContent fetches data from bucket object by bucket name and object name
func (c *Client) csvObjectContent(bktName, objName string) ([][]string, error) {
	r, err := c.strgCln.Bucket(bktName).Object(objName).NewReader(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch object csv data: %v", err)
	}
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("fetch object csv data: %v", err)
	}

	return records, nil
}

// MakeReport creates report from gcp object in cloud for object
func (c *Client) makeReport(obj gcptypes.Object) (*Report, error) {
	data, err := c.csvObjectContent(obj.Bucket, obj.Name)
	if err != nil {
		return nil, fmt.Errorf("can not make report: %v", err)
	}

	sbs, err := gcpparser.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("can not make report: %v", err)
	}
	rep := Report{
		Object: obj,
		Bills:  sbs,
	}

	return &rep, nil
}

// MakeReports creates reports from gcp object in cloud for objects range
func (c *Client) MakeReports(objs gcptypes.Objects) (reps Reports, err error) {
	for _, o := range objs {
		r, err := c.makeReport(o)
		if err != nil {
			return nil, fmt.Errorf("can not make reports: %v", err)
		}
		reps = append(reps, r)
	}

	return reps, nil
}
