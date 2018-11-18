package summary

import (
	"fmt"
	"strconv"

	"github.com/radisvaliullin/cloudbilling/pkg/utils"
)

// Summary represents the summary
type Summary struct {
	headers []string
	inner   *map[string]([]float64)
}

// New returns empty Summary
func New() *Summary {
	return &Summary{}
}

// NewSummaryFrom returns the new Summary
func NewSummaryFrom(h []string, m *map[string]([]float64)) *Summary {
	return &Summary{
		headers: h,
		inner:   m,
	}
}

// Inner returns the key-value map of the CsvSummary
func (summary Summary) Inner() *map[string]([]float64) {
	return summary.inner
}

// Headers returns the key-value map of the CsvSummary
func (summary Summary) Headers() []string {
	return summary.headers
}

// AddStringValues takes new data into the account
func (summary *Summary) AddStringValues(row []string) error {
	// If there are no headers, so the first line is headers
	if summary.headers == nil {
		summary.headers = row
		return nil
	}

	// parse the rest of the data
	if (*summary.inner)[row[0]] == nil {
		floats, err := utils.StringSliceToFloat64(row[1:])
		if err != nil {
			return fmt.Errorf("can't parse SumBy to slice of floats: %v", err)
		}
		(*summary.inner)[row[0]] = floats

		return nil
	} else {
		prevRow := (*summary.inner)[row[0]]
		for j, v := range prevRow {
			value, err := strconv.ParseFloat(row[j+1], 64)
			if err != nil {
				return fmt.Errorf("can't parse '%v' to float64: %v", v, err)
			}
			prevRow[j] += value
		}
		(*summary.inner)[row[0]] = prevRow
	}

	return nil
}
