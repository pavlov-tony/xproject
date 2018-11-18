package summary

import (
	"testing"

	"github.com/radisvaliullin/cloudbilling/pkg/utils"
)

func TestAddStringValues(t *testing.T) {
	testcases := []struct {
		name     string
		in       [][]string
		expected []*Summary
	}{
		{
			name: "Iteratively add data",
			in: [][]string{
				{"lineItem/ProductCode", "lineItem/UnblendedCost"},
				{"AmazonEC2", "1"},
				{"AmazonQuickSight", "2"},
				{"AmazonEC2", "30"},
			},
			expected: []*Summary{
				NewSummaryFrom(
					[]string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
					nil,
				),
				NewSummaryFrom(
					[]string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
					&map[string]([]float64){
						"AmazonEC2": []float64{1},
					},
				),
				NewSummaryFrom(
					[]string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
					&map[string]([]float64){
						"AmazonEC2":        []float64{1},
						"AmazonQuickSight": []float64{2},
					},
				),
				NewSummaryFrom(
					[]string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
					&map[string]([]float64){
						"AmazonEC2":        []float64{31},
						"AmazonQuickSight": []float64{2},
					},
				),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			inner := make(map[string]([]float64))
			summary := &Summary{
				inner: &inner,
			}

			for i, row := range testcase.in {
				err := summary.AddStringValues(row)
				if err != nil {
					t.Fatal("Error while adding values to summary:", err)
				}

				// Check if the header is correct
				if i == 0 {
					if !utils.IsStringSlicesEqual(summary.Headers(), testcase.expected[i].Headers()) {
						t.Fatal("Header differs:", *summary.Inner(), "!=", *testcase.expected[i].Inner())
					}

					continue
				}

				// Check the entries
				if !utils.IsMapsEqual(*summary.Inner(), *testcase.expected[i].Inner()) {
					t.Fatal("Summary differs:", *summary.Inner(), "!=", *testcase.expected[i].Inner())
				}
			}
		})
	}
}
