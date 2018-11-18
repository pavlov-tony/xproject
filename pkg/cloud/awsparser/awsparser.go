package awsparser

import (
	"fmt"
	"io"

	"github.com/radisvaliullin/cloudbilling/pkg/cloud/awsparser/models/reportrow"
	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/csvreader"
	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/rawcsv"
)

var (
	// NamesToFilter decalared as var because of the nature of the slices,
	// but it is actually a constant array.
	// See the description of the headers at the
	// github.com/radisvaliullin/xproject/pkg/cloud/awsparser/models/reportrow/reportrow.go
	NamesToFilter = []string{
		"identity/LineItemId",
		"identity/TimeInterval", // should br splitted to start and end
		"bill/PayerAccountId",
		"bill/BillingPeriodStartDate",
		"bill/BillingPeriodEndDate",
		"lineItem/LineItemType",
		"lineItem/ProductCode",
		"lineItem/UsageAmount",
		"lineItem/CurrencyCode",
		"lineItem/UnblendedRate",
		"lineItem/UnblendedCost",
		"lineItem/BlendedRate",
		"lineItem/BlendedCost",
		"product/region",
		"product/sku",
		"pricing/publicOnDemandCost",
		"pricing/publicOnDemandRate",
		"pricing/term",
		"pricing/unit",
	}
)

// ReportReader is a reader over the io.Reader
// Each Read() returns the ReportRow which is represents the needed row
// of the AWS Cost And Usage Report
type ReportReader struct {
	reader *csvreader.CsvReader
}

// NewReportReader returns a new ReportReader
func NewReportReader(r io.Reader) *ReportReader {
	csvr := csvreader.New(r, NamesToFilter)
	return &ReportReader{
		reader: csvr,
	}
}

// Read reads from ReportReader by ReportRow
func (r *ReportReader) Read() (*reportrow.ReportRow, error) {
	s, err := r.reader.Read()
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't read: %v", err)
	}

	return reportrow.FromStrings(s)
}

// ReadAll collects all the values
func (rr *ReportReader) ReadAll() ([]*reportrow.ReportRow, error) {
	// Bewcause of the unknown size of the stream, we do not predict the length
	// of the resulting slice.
	result := make([]*reportrow.ReportRow, 0)

	for {
		row, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while reading from ReportReader: %v", err)
		}

		result = append(result, row)
	}

	return result, nil
}

// RawCsvToDbRecords converts RawCsv struct to the slice of ReportRows.
// To read a report line by line use the ReportReader
func RawCsvToReportRows(csv *rawcsv.RawCsv) ([]*reportrow.ReportRow, error) {
	filteredCsv, err := csv.FilterByNames(NamesToFilter)
	if err != nil {
		return nil, fmt.Errorf("can't filter csv by required columns: %v", err)
	}

	result := make([]*reportrow.ReportRow, len(filteredCsv.Rows()))

	for _, row := range filteredCsv.Rows()[1:] {
		rr, err := reportrow.FromStrings(row)
		if err != nil {
			return nil, fmt.Errorf("error while parsing ReportRow: %v", err)
		}

		result = append(result, rr)
	}

	return result, nil
}
