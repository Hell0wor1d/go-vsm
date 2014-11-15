//	implement code of VSM by Kev7n
//	https://github.com/Hell0wor1d/go-vsm
package util

import (
	"math"
)

//	compute similarity between source dict and target dict.
//	parmas:
//		dt: source dictionary
//		ds: target dictionary
//	return similarity value.
func Compute(ds map[string]float64, dt map[string]float64) float64 {

	if len(ds) < 1 || len(dt) < 1 {
		return .0
	}

	numerator, denominator1, denominator2, temp2 := .0, .0, .0, .0

	for key, value := range ds {
		temp1 := value
		if v, ok := dt[key]; !ok {
			temp2 = 0
		} else {
			temp2 = v
		}
		delete(dt, key)
		numerator += temp1 * temp2
		denominator1 += temp1 * temp1
		denominator2 += temp2 * temp2
	}

	for _, value := range dt {
		denominator2 += value * value
	}

	result := numerator / (math.Sqrt(denominator1 * denominator2))
	return result
}
