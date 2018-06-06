// TODO: storageutil.go should to use - storage.go or fetchdata.go or else
package storageutils

import (
	"encoding/csv"
	"errors"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// make reader?
type Object struct {
	Name    string
	Bucket  string
	Created time.Time
}

type Objects []Object

// TODO: testing (need test api?)
// fetching project bucket names
func FetchProjectBuckets(ctx context.Context, client *storage.Client,
	projectID string) (buckets []string, err error) {

	it := client.Buckets(ctx, projectID)
	for {
		b, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return buckets, errors.New("GetBucket() failed\n" + err.Error())
		}
		buckets = append(buckets, b.Name)
	}

	return buckets, nil
}

// fetching objects names with created time from bucket
func FetchBucketObjects(ctx context.Context, client *storage.Client,
	bucketName string) (objects Objects, err error) {

	it := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return objects, errors.New("GetBucketObjects() failed\n" + err.Error())
		}
		objects = append(objects, Object{o.Name, o.Bucket, o.Created})
	}

	return objects, nil
}

// TODO: testing
//CSV Reader for object, works only for CSV objects
func (object Object) NewCSVReader(ctx context.Context, client *storage.Client) *csv.Reader {
	objectReader, err := client.Bucket(object.Bucket).Object(object.Name).NewReader(ctx)
	if err != nil {
		log.Fatal(err) // FIXME: handle error
	}

	return csv.NewReader(objectReader) // TODO:  have no closed yet
}

// FIXME: *Objects?
// Select objects from objects list where name has prefix
func (objects Objects) SelectWithPrefix(prefix string) (result Objects) {
	for _, o := range objects {
		if strings.HasPrefix(o.Name, prefix) {
			result = append(result, o)
		}
	}

	return result
}

// Select objects from objects list where from < created time < to
func (objects Objects) SelectInTimeRange(from, to time.Time) (result Objects, err error) {
	if to.Before(from) {
		return result, errors.New("error: to before from")
	}
	for _, o := range objects {
		if o.Created.After(from) && o.Created.Before(to) {
			result = append(result, o)
		}
	}

	return result, nil
}
