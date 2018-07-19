package approx

import (
	"testing"
)

func Test_LinearTrend(t *testing.T) {
	testData := []float64{5.3, 6.3, 4.8, 3.8, 3.3}
	// test 1
	ahead := 0
	exp := 6.65
	got, err := LinearTrend(testData, ahead)
	if err != nil {
		t.Error("LinearTrend: unexpected division by zero")
	}
	if got != exp {
		t.Errorf("LinearTrend: Exp %f, Got %f", exp, got)
	}
	//test 2
	ahead = 6
	exp = 2.75
	got, err = LinearTrend(testData, ahead)
	if err != nil {
		t.Error("LinearTrend: Unexpected division by zero")
	}
	if got != exp {
		t.Errorf("LinearTrend: Exp %f, Got %f", exp, got)
	}
	// test 3
	testData = []float64{}
	got, err = LinearTrend(testData, ahead)
	if err == nil {
		t.Error("LinearTrend: Expected division by zero")
	}
	// test 4
	got, err = LinearTrend(testData, ahead)
	if err == nil {
		t.Error("LinearTrend: Expected division by zero")
	}
}
