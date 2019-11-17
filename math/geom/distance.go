package geom

import (
	"agar-life/math/crd"
	"math"
)

func GetDistanceByCrd(crd1, crd2 crd.Crd) float64 {
	d1 := crd1.GetX() - crd2.GetX()
	d2 := crd1.GetY() - crd2.GetY()
	return math.Sqrt(
		d1*d1 + d2*d2,
	)
}
