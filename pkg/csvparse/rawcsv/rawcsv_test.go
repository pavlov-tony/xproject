package rawcsv

import (
	"strings"
	"testing"

	"github.com/radisvaliullin/cloudbilling/pkg/csvparse/summary"
	"github.com/radisvaliullin/cloudbilling/pkg/utils"
)

func TestRawCsvFromReader(t *testing.T) {
	rawcsv := &RawCsv{
		records: [][]string{
			{"a", "b", "c"},
			{"1", "2", "3"},
			{"4", "5", "6"},
		},
	}

	csvString := `a,b,c
1,2,3
4,5,6
`

	strReader := strings.NewReader(csvString)
	csv, err := FromReader(strReader)
	if err != nil {
		t.Fatalf("Creating RawCsv failed:, %v\n", err)
	}

	if !equalRawCsvs(*csv, *rawcsv) {
		t.Fatal("FromReader produced RawCsv that differs:", csv, "!=", rawcsv)
	}
}

func TestColumnIndexByName(t *testing.T) {
	colsNames := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "Get index by column name 'a'",
			in:   "a",
			out:  0,
		},
		{
			name: "Get index by column name 'b'",
			in:   "b",
			out:  1,
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, cn := range colsNames {
		t.Run(cn.name, func(t *testing.T) {
			idx, err := rawcsv.ColumnIndexByName(cn.in)
			if err != nil {
				t.Fatal("Error while getting column index by name: ", err)
			}
			if idx != cn.out {
				t.Fatal("Found index isn't correct: ", idx, "!=", cn.out)
			}
		})
	}
}

func equalRawCsvs(a, b RawCsv) bool {
	if len(a.Rows()) != len(b.Rows()) {
		return false
	}
	for i, row := range a.Rows() {
		for j, value := range row {
			if value != b.Rows()[i][j] {
				return false
			}
		}
	}

	return true
}

func TestFilterByIndices(t *testing.T) {
	testcases := []struct {
		name string
		in   []int
		out  *RawCsv
	}{
		{
			name: "Filter by column with index 0",
			in:   []int{0},
			out: &RawCsv{
				records: [][]string{{"a"}, {"1"}},
			},
		},
		{
			name: "Filter by column with index 1",
			in:   []int{1},
			out: &RawCsv{
				records: [][]string{{"b"}, {"2"}},
			},
		},
		{
			name: "Filter by columns with indices 0 and 1",
			in:   []int{0, 1},
			out: &RawCsv{
				records: [][]string{{"a", "b"}, {"1", "2"}},
			},
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			csv, err := rawcsv.FilterByIndices(testcase.in)
			if err != nil {
				t.Fatal("Error while filtering by indices: ", err)
			}
			if !equalRawCsvs(*csv, *testcase.out) {
				t.Fatal("Filter isn't correct: ", testcase, "!=", testcase.out)
			}
		})
	}
}

func TestFilterByNames(t *testing.T) {
	testcases := []struct {
		name string
		in   []string
		out  *RawCsv
	}{
		{
			name: "Filter by column with name 'a'",
			in:   []string{"a"},
			out: &RawCsv{
				records: [][]string{{"a"}, {"1"}},
			},
		},
		{
			name: "Filter by column with name 'b'",
			in:   []string{"b"},
			out: &RawCsv{
				records: [][]string{{"b"}, {"2"}},
			},
		},
		{
			name: "Filter by columns 'a' and 'b'",
			in:   []string{"a", "b"},
			out: &RawCsv{
				records: [][]string{{"a", "b"}, {"1", "2"}},
			},
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			csv, err := rawcsv.FilterByNames(testcase.in)
			if err != nil {
				t.Fatal("Error while filtering by names: ", err)
			}
			if !equalRawCsvs(*csv, *testcase.out) {
				t.Fatal("Filter isn't correct: ", testcase, "!=", testcase.out)
			}
		})
	}
}

func TestSummarize(t *testing.T) {
	testcases := []struct {
		name string
		in   *RawCsv
		cfg  *summary.Config
		out  *summary.Summary
	}{
		{
			name: "GroupBy 'a' and SumBy 'b' and 'c'",
			in: FromRows([][]string{
				{"a", "b", "c"},
				{"z", "2", "3"},
				{"z", "2", "3"},
				{"y", "5", "6"},
				{"y", "5", "6"},
			}),
			cfg: summary.NewConfig("a", []string{"b", "c"}),
			out: summary.NewSummaryFrom(
				[]string{"a", "b", "c"},
				&map[string]([]float64){
					"z": []float64{4.0, 6.0},
					"y": []float64{10.0, 12.0},
				},
			),
		},
		{
			name: "GroupBy 'identity/TimeInterval' and SumBy 'lineItem/UnblendedCost'",
			in: FromRows([][]string{
				{"identity/TimeInterval", "lineItem/UnblendedCost"},
				{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "1"},
				{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "2"},
				{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "3"},
				{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "4"},
				{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "10"},
				{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "20"},
				{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "30"},
				{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "40"},
			}),
			cfg: summary.NewConfig("identity/TimeInterval", []string{"lineItem/UnblendedCost"}),
			out: summary.NewSummaryFrom(
				[]string{"identity/TimeInterval", "lineItem/UnblendedCost"},
				&map[string]([]float64){
					"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z": []float64{10},
					"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z": []float64{100},
				},
			),
		},
		{
			name: "GroupBy 'lineItem/ProductCode' and SumBy 'lineItem/UnblendedCost'",
			in: FromRows([][]string{
				{"lineItem/ProductCode", "lineItem/UnblendedCost"},
				{"AmazonEC2", "1"},
				{"AmazonQuickSight", "2"},
				{"AmazonRDS", "3"},
				{"AmazonQuickSight", "4"},
				{"AmazonEC2", "10"},
				{"AmazonEC2", "20"},
				{"AmazonRDS", "30"},
				{"AmazonCloudWatch", "40"},
			}),
			cfg: summary.NewConfig("lineItem/ProductCode", []string{"lineItem/UnblendedCost"}),
			out: summary.NewSummaryFrom(
				[]string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
				&map[string]([]float64){
					"AmazonCloudWatch": []float64{40},
					"AmazonEC2":        []float64{31},
					"AmazonRDS":        []float64{33},
					"AmazonQuickSight": []float64{6},
				},
			),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			scfg := testcase.cfg
			summary, err := testcase.in.Summarize(scfg)
			if err != nil {
				t.Fatal("Error while getting summary:", err)
			}

			if !utils.IsStringSlicesEqual(summary.Headers(), testcase.out.Headers()) {
				t.Fatal("Header differs:", summary.Headers(), "!=", testcase.out.Headers())
			}
			if !utils.IsMapsEqual(*summary.Inner(), *testcase.out.Inner()) {
				t.Fatal("Summary differs:", *summary.Inner(), "!=", *testcase.out.Inner())
			}
		})
	}
}
