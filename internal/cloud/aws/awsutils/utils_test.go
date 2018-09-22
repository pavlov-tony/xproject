// +build integration

package awsutils

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	caurs "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

func TestFindingReport(t *testing.T) {
	// FIXME: Consider using the environment variables explicitly
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		t.Fatalf("Failed to load config, %v", err)
	}

	reportName := "test-somereport"

	manifestInfo, err := FindManifest(cfg, reportName)
	if err != nil {
		t.Fatalf("Error finding manifest by report name: %v", err)
	}

	t.Log("manifest:", manifestInfo)

	rep, err := FindReportByName(&cfg, manifestInfo)
	if err != nil {
		t.Fatalf("Failed to get report, %v", err)
	}

	t.Log("report:", rep.ReportPath)
}

func TestGetMostRecentReportInterval(t *testing.T) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		t.Fatalf("Failed to load config, %v", err)
	}

	rd := caurs.ReportDefinition{
		S3Bucket:   aws.String("user-xproject-test-bucket-12345678"),
		S3Prefix:   aws.String("reportprefix"),
		ReportName: aws.String("test-somereport"),
	}
	value, err := getMostRecentReportInterval(&cfg, &rd)
	if err != nil {
		t.Fatal("Error getting most recent report:", err)
	}

	t.Log(value)
}
