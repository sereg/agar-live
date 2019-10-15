package geom

import (
	"agar-life/math/crd"
	"math"
)

func GetDistanceByCrd(crd1, crd2 crd.Crd) float64 {
	d1 := crd1.X() - crd2.X()
	d2 := crd1.Y() - crd2.Y()
	return math.Sqrt(
		d1*d1 + d2*d2,
	)
}
