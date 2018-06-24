package forecaster

import (
	"testing"
)

func Test_forecast(t *testing.T) {
	testData := []float64{5.3, 6.3, 4.8, 3.8, 3.3}
	// test 1
	ahead := 0
	exp := 6.65
	got, err := forecast(testData, ahead)
	if err != nil {
		t.Error("1) Unexpected division by zero")
	}
	if got != exp {
		t.Errorf("1) Exp %f, Got %f", exp, got)
	}
	//test 2
	ahead = 6
	exp = 2.75
	got, err = forecast(testData, ahead)
	if err != nil {
		t.Error("2) Unexpected division by zero")
	}
	if got != exp {
		t.Errorf("2) Exp %f, Got %f", exp, got)
	}
	// test 3
	testData = []float64{}
	got, err = forecast(testData, ahead)
	if err == nil {
		t.Error("3) Expected division by zero")
	}
	// test 4
	got, err = forecast(testData, ahead)
	if err == nil {
		t.Error("4) Expected division by zero")
	}
}
