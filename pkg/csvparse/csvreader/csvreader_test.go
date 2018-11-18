package csvreader

import (
	"io"
	"strings"
	"testing"

	"github.com/radisvaliullin/cloudbilling/pkg/utils"
)

func TestReadRow(t *testing.T) {
	testcases := []struct {
		name    string
		in      string
		columns []string
		out     [][]string
	}{
		{
			name: "Filter by first_name and username",
			in: `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`,
			columns: []string{"first_name", "username"},
			out: [][]string{
				{"Rob", "rob"},
				{"Ken", "ken"},
				{"Robert", "gri"},
			},
		},
		{
			name: "Filter by a, c, and d",
			in: `a,b,c,d,e
1,2,3,4,5
6,7,8,9,0
1,1,1,1,1
`,
			columns: []string{"a", "c", "d"},
			out: [][]string{
				{"1", "3", "4"},
				{"6", "8", "9"},
				{"1", "1", "1"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r := New(strings.NewReader(testcase.in), testcase.columns)

			for _, row := range testcase.out {
				s, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("Error while reading from csv: %v", err)
				}
				if !utils.IsStringSlicesEqual(s, row) {
					t.Fatalf("Rows are not equal: %v != %v, (%v, %v)", s, row, len(s), len(row))
				}
			}
		})
	}
}
