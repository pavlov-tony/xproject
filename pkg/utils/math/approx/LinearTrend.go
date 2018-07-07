// Package approx is util for creating linear trend line
package approx

import "fmt"

// LinearTrend implements least squares method
// http://www.ekonomika-st.ru/drugie/metodi/metodi-prognoz-1-5.html
//
// Original function value: yOrig
// Opproximation function:  yAppr = ax + b
//
// F(a, b) = Sum(yOrig[i] - yAppr[i]) -> min
//
// Find min of F(a, b):
// F'(a, b) = 0 =>
// 	1. dF/da = 0
// 	2. dF/db = 0
//
// params:
// data - data
// ahX - steps ahead prediction
func LinearTrend(data []float64, ahX int) (float64, error) {
	var a, b, n, sumY, sumX, sumPow2X, sumXY float64

	n = float64(len(data))

	for i, v := range data {
		x := float64(i + 1)
		sumX += x
		sumY += v
		sumXY += x * v
		sumPow2X += x * x
	}

	if n == 0 || sumX == 0 && sumPow2X == 0 {
		return 0, fmt.Errorf("linear trend: division by zero")
	}

	a = (sumXY - sumX*sumY/n) / (sumPow2X - sumX*sumX/n)
	b = sumY/n - a*sumX/n

	res := a*float64(ahX) + b

	return res, nil
}
