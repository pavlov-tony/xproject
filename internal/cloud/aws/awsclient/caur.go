// Implements the algorithm of fetching the Amazon Cost and Usage Reports

package awsclient

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/pavlov-tony/xproject/internal/cloud/aws/awsutils"
	"github.com/pavlov-tony/xproject/pkg/cloud/awsparser"
	"github.com/pavlov-tony/xproject/pkg/cloud/awsparser/models/reportrow"
)

// CaurMeta describe the requested report
type CaurMeta struct {
	ReportName string
	AwsCfg     *aws.Config
	ReportID   string
}

// NewCaurMeta creates new meta for the requested report
func NewCaurMeta(connectionCfg *aws.Config, reportName string) *CaurMeta {
	return &CaurMeta{
		ReportName: reportName,
		AwsCfg:     connectionCfg,
	}
}

// FetchCAUR actually fetches a Cost and Usage Report from Amazon.
// The report meta is required.
func (client *Client) FetchCAUR() ([]*reportrow.ReportRow, error) {
	manifestInfo, err := awsutils.FindManifest(*client.Meta.AwsCfg, client.Meta.ReportName)
	if err != nil {
		return nil, fmt.Errorf("error finding manifest by report name: %v", err)
	}

	log.Println("manifest info:", manifestInfo)

	reportInfo, err := awsutils.FindReportByName(client.Meta.AwsCfg, manifestInfo)
	if err != nil {
		return nil, fmt.Errorf("error finding report by name: %v", err)
	}

	log.Println("manifest report id:", reportInfo.Manifest.ReportId)

	if reportInfo.Manifest.ReportId == client.Meta.ReportID {
		log.Println("Skip the report due to similarity")
		return nil, nil // Skip if id is similar
	}
	client.Meta.ReportID = reportInfo.Manifest.ReportId

	zippedReportFile, err := awsutils.DownloadReport(client.Meta.AwsCfg, reportInfo)
	if err != nil {
		return nil, fmt.Errorf("error downloading report: %v", err)
	}
	defer zippedReportFile.Close()
	defer os.Remove(zippedReportFile.Name())

	// Unzip

	zr, err := gzip.NewReader(zippedReportFile)
	if err != nil {
		return nil, fmt.Errorf("error reading zipped report: %v", err)
	}

	f2, err := ioutil.TempFile("", "_unzippedreport")
	if err != nil {
		return nil, fmt.Errorf("error creating unzipped file: %v", err)
	}
	defer f2.Close()
	defer os.Remove(f2.Name())

	// TODO: Optimize buffer reallocation: provide slice with
	// predefined size, known at runtime
	data_ := bytes.NewBuffer([]byte{})

	if _, err := io.Copy(data_, zr); err != nil {
		return nil, fmt.Errorf("error writting to unzipped report: %v", err)
	}

	if err := zr.Close(); err != nil {
		return nil, fmt.Errorf("error closing zip reader: %v", err)
	}

	reportReader := awsparser.NewReportReader(strings.NewReader(string(data_.Bytes())))
	return reportReader.ReadAll()
}
