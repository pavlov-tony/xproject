package awsutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	caurs "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"

	awsmanifest "github.com/pavlov-tony/xproject/pkg/cloud/awsparser/models/manifest"
)

// FindManifest helps to automagically find the report manifest by
// provided report's name
//
// It requests the available report exporters (with given name) from
// AWS, scans them, and chooses the _first_ with the matching name.
func FindManifest(cfg aws.Config, reportName string) (*caurs.ReportDefinition, error) {
	// us-east-1 is required to recv report definitions (?)
	cfg.Region = endpoints.UsEast1RegionID

	caursClient := caurs.New(cfg)
	params := caurs.DescribeReportDefinitionsInput{}
	req := caursClient.DescribeReportDefinitionsRequest(&params)

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("cannot make request to AWS: '%v'", err)
	}

	for _, repdef := range resp.ReportDefinitions {
		if reportName == *repdef.ReportName {
			return &repdef, nil
		}
	}

	return nil, fmt.Errorf("error: no reports found with name '%v'", reportName)
}

// FindReportByName helps to find report by its definition
//
// TODO: Consider renaming this function
func FindReportByName(cfg *aws.Config, repdef *caurs.ReportDefinition) (*CaurInfo, error) {
	rp, manif, err := getMostRecentReportPath(cfg, repdef)
	if err != nil {
		log.Println("Error while getting the most recent report path:", err)
		return nil, err
	}

	result := &CaurInfo{
		Definition: repdef,
		ReportPath: rp,
		Manifest:   manif,
	}

	return result, nil

}

// getMostRecentReportPath automagically finds the most recent report
// with provided name
func getMostRecentReportPath(cfg *aws.Config, reportDefinition *caurs.ReportDefinition) (string, *awsmanifest.Manifest, error) {

	reportDate, err := getMostRecentReportInterval(cfg, reportDefinition)
	if err != nil {
		return "", nil, fmt.Errorf("error finding report intervals: %v", err)
	}

	rd := reportDefinition
	manifestPath := *rd.S3Prefix + "/" +
		*rd.ReportName + "/" + reportDate + "/" + *rd.ReportName + "-Manifest.json"
	cfg.Region = string(rd.S3Region)
	fmt.Println("manifest path:", manifestPath)

	downloader := s3manager.NewDownloader(*cfg)

	tmpfile, err := ioutil.TempFile("", "caur-manifest_")
	if err != nil {
		log.Println("error openning temp file", err)
		return "", nil, fmt.Errorf("error openning temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = downloader.Download(tmpfile, &s3.GetObjectInput{
		Bucket: reportDefinition.S3Bucket,
		Key:    aws.String(manifestPath),
	})
	if err != nil {
		log.Printf("failed to download manifest file, %v", err)
		return "", nil, fmt.Errorf("failed to download manifest file, %v", err)
	}

	fname := tmpfile.Name()
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("failed to read file:", err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	var manif awsmanifest.Manifest
	err = json.Unmarshal(dat, &manif)

	if manif.ReportKeys != nil {
		// NOTE: report name is a unique identifier, so it is probably ok
		// to get 0 index
		return manif.ReportKeys[0], &manif, nil
	}

	return "", nil, fmt.Errorf("can't get manifest report keys")
}

// getMostRecentReportInterval returns the time interval of the most
// recent report as a string
//
// Actually, the current date can mismatch this interval, even if the
// last update was at that day. So, I decided to scan all files in the
// bucket with known prefix path for the latest time interval.
func getMostRecentReportInterval(cfg *aws.Config, reportDefinition *caurs.ReportDefinition) (string, error) {
	svc := s3.New(*cfg)

	params := &s3.ListObjectsInput{
		Bucket: reportDefinition.S3Bucket,
		Prefix: aws.String(*reportDefinition.S3Prefix + "/" + *reportDefinition.ReportName),
	}

	req := svc.ListObjectsRequest(params)
	resp, err := req.Send()
	if err != nil {
		return "", fmt.Errorf("cannot send request to AWS: %v", err)
	}

	intervalPrefix := regexp.MustCompile(`^\d{6}01\-\d{6}01`)
	uselessPrefix := *reportDefinition.S3Prefix + "/" + *reportDefinition.ReportName + "/"

	var last *string
	for _, key := range resp.Contents {
		trimmed := strings.TrimPrefix(*key.Key, uselessPrefix)
		if intervalPrefix.MatchString(trimmed) {

			// AWS ListObject returns the list of the objects in alphabetical order,
			// so we can simply get the last element
			last = &trimmed
		}
	}

	if last == nil {
		log.Println("Can't find any report related time intervals, probably client has no reports at S3 or hasn't enabled the feature yet")
		return "", fmt.Errorf("cannot find any time intervals")
	}

	interval := intervalPrefix.FindString(*last)

	return interval, nil
}

// DownloadReport returns the file descriptor of the downloaded CSV report
func DownloadReport(cfg *aws.Config, reportInfo *CaurInfo) (*os.File, error) {

	downloader := s3manager.NewDownloader(*cfg)

	f, err := ioutil.TempFile("", "_report")
	if err != nil {
		return nil, fmt.Errorf("failed to create file %q, %v", f.Name(), err)
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: reportInfo.Definition.S3Bucket,
		Key:    aws.String(reportInfo.ReportPath),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}

	fmt.Printf("file downloaded, %d bytes\n", n)

	return f, nil
}
