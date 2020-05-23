package main

import "math"

func sqrtNTimes(x float64, n int) float64 {
	result := x
	for i := 0; i < n; i++ {
		result = math.Sqrt(result)
	}
	return result
}