package awsutils

import (
	caur "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	awsmanifest "github.com/radisvaliullin/cloudbilling/pkg/cloud/awsparser/models/manifest"
)

// CaurInfo - useful information about an AWS Cost and Usage Report which is
// being processed
type CaurInfo struct {
	Definition   *caur.ReportDefinition
	ManifestPath string
	ReportPath   string

	Manifest *awsmanifest.Manifest
}

// ReportInfo helps to fetch the Cost and Usage Reports from AWS
type ReportInfo struct {
	BucketName         string
	ReportPrefix       string
	ReportName         string
	ReportDateInterval string
}

// ManifestPath constructs the manifest path from known parts
func (z *ReportInfo) ManifestPath() string {
	return z.BucketName + "/" + z.ReportName + "/" + z.ReportDateInterval +
		"/" + "-Manifest.json"
}
