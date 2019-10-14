package geom

import (
	"agar-life/object"
	"math"
)

func GetDistanceByCrd(crd1, crd2 object.Crd) float64 {
	d1 := crd1.GetX() - crd2.GetX()
	d2 := crd1.GetY() - crd2.GetY()
	return math.Sqrt(
		d1*d1 + d2*d2,
	)
}

type Segment struct{
	a, b object.Crd
}

func NewSegment(a, b object.Crd) Segment {
	return Segment{
		a: a, b: b,
	}
}

func (s Segment) Intersection(s1 Segment) bool {
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := s.a.GetX(), s.a.GetY(), s.b.GetX(), s.b.GetY(), s1.a.GetX(), s1.a.GetY(), s1.b.GetX(), s1.b.GetY()
	v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
	v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
	v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
	v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)
	return (v1*v2 < 0) && (v3*v4 < 0)
}
