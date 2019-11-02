package geom

import "math"

func ModuleDegree(f float64) float64 {
	if f < 0 {
		f = math.Pi + math.Pi + f
	}
	return f
}
