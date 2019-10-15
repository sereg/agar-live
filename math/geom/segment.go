package geom

import "agar-life/math/crd"

type Segment struct{
	a, b crd.Crd
}

func NewSegment(a, b crd.Crd) Segment {
	return Segment{
		a: a, b: b,
	}
}

func (s Segment) Intersection(s1 Segment) bool {
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := s.a.X(), s.a.Y(), s.b.X(), s.b.Y(), s1.a.X(), s1.a.Y(), s1.b.X(), s1.b.Y()
	v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
	v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
	v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
	v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)
	return (v1*v2 < 0) && (v3*v4 < 0)
}
