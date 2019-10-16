package math

import "math"

// Round left 2 sign after comma
func Round(num float64) float64 {
	return float64(int(num * 100)) / 100
}

// ToFixed returns float with precision
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
