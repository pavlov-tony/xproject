package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/csvutils"
)

// CsvReader is a type for reading a CSV file row-by-row, with the
// filtering option.
type CsvReader struct {
	// filter is a slice of header names to be filtered while reading
	// the CSV file
	filter []string

	// indices are the corresponding indices of the headers
	indices []int

	// csvReader is an inner representation of the string reader
	// See "encoding/csv" for more details
	csvReader *csv.Reader
}

// New returns a new CsvReader
// Note that it skips the header for purpose
func New(r io.Reader, filterHeaders []string) *CsvReader {
	return &CsvReader{
		filter:    filterHeaders,
		indices:   nil,
		csvReader: csv.NewReader(r),
	}
}

// Read is actually reads a row from CSV as a slice of strings
func (r *CsvReader) Read() ([]string, error) {
	if r.indices == nil {
		r.indices = make([]int, len(r.filter))

		header, err := r.csvReader.Read()
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			return nil, fmt.Errorf("can't read header: %v", err)
		}

		// find indices from filter in header
		for i, columnName := range r.filter {
			idx, err := csvutils.ColumnIndexByName(header, columnName)
			if err != nil {
				return nil, fmt.Errorf("can't get column index by name \"%v\": %v", columnName, err)
			}
			r.indices[i] = idx
		}
	}

	record, err := r.csvReader.Read()

	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("can't read row: %v", err)
	}

	filteredRecord := make([]string, len(r.filter))

	for i, idx := range r.indices {
		filteredRecord[i] = record[idx]
	}

	return filteredRecord, nil
}
