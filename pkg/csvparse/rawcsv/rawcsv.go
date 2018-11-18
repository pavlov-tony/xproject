package rawcsv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/csvutils"
	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/errors"
	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/summary"
)

// RawCsv represents the matrix of strings parsed from CSV
type RawCsv struct {
	records [][]string
}

// FromRows creates RawCsv from [][]string
func FromRows(rows [][]string) *RawCsv {
	return &RawCsv{
		records: rows,
	}
}

// FromReader creates RawCsv from io.Reader
func FromReader(in io.Reader) (*RawCsv, error) {
	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can't create RawCsv from io.Reader: %v", err)
	}

	result := &RawCsv{
		records: records,
	}
	return result, nil
}

// Rows returns the containing rows of RawCsv
func (raw *RawCsv) Rows() [][]string {
	return raw.records
}

// Row returns the row of RawCsv by index
func (raw *RawCsv) Row(index int) ([]string, error) {
	if index < len(raw.records) {
		return raw.records[index], nil
	}

	return nil, errors.NewIndexError(index)
}

// ColumnIndexByName returns the number of the provided header name
func (raw *RawCsv) ColumnIndexByName(name string) (int, error) {
	if len(raw.records) == 0 {
		return 0, fmt.Errorf("empty RawCsv")
	}

	header := raw.records[0]

	return csvutils.ColumnIndexByName(header, name)
}

// FilterByNames returns RawCsv with provided header names only
func (raw *RawCsv) FilterByNames(columns []string) (*RawCsv, error) {
	indices := make([]int, len(columns))
	for i, columnName := range columns {
		idx, err := raw.ColumnIndexByName(columnName)
		if err != nil {
			return nil, fmt.Errorf("can't get column index by name \"%v\": %v", columnName, err)
		}
		indices[i] = idx
	}

	return raw.FilterByIndices(indices)
}

// FilterByIndices returns RawCsv with provided header indices only
func (raw *RawCsv) FilterByIndices(indices []int) (*RawCsv, error) {
	rows := raw.Rows()
	if rows == nil {
		return nil, fmt.Errorf("empty RawCsv")
	}

	final := make([][]string, len(rows))
	for i, r := range rows {
		final[i] = make([]string, len(indices))
		for j, c := range indices {
			final[i][j] = r[c]
		}
	}

	return FromRows(final), nil
}

// Summarize returns a Summary of the RawCsv by the provided config
func (raw *RawCsv) Summarize(cfg *summary.Config) (*summary.Summary, error) {
	if cfg.GroupBy == "" || cfg.SumBy == nil {
		return nil, fmt.Errorf("can't produce CSV Summary: GroupBy and SumBy mustn't be empty")
	}

	filtered, err := raw.FilterByNames(append([]string{cfg.GroupBy}, cfg.SumBy...))
	if err != nil {
		return nil, fmt.Errorf("can't filter by names: %v", err)
	}

	inner := make(map[string]([]float64))
	sm := summary.NewSummaryFrom(filtered.Rows()[0], &inner)

	for _, row := range filtered.Rows()[1:] {
		err2 := sm.AddStringValues(row)
		if err2 != nil {
			return nil, fmt.Errorf("can't add string values: %v", err2)
		}
	}

	return sm, nil
}
