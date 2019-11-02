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

func (s Segment) Len() float64 {
	return GetDistanceByCrd(s.Start(), s.Finish())
}

func (s Segment) IntersectionPoint(s1 Segment) (intersect bool, cr crd.Crd) {
	// Line AB represented as a1x + b1y = c1
	a1 := s.Finish().Y() - s.Start().Y()
	b1 := s.Start().X() - s.Finish().X()
	c1 := a1*s.Start().X() + b1*s.Start().Y()

	// Line CD represented as a2x + b2y = c2
	a2 := s1.Finish().Y() - s1.Start().Y()
	b2 := s1.Start().X() - s1.Finish().X()
	c2 := a2*s1.Start().X() + b2*s1.Start().Y()

	determinant := a1*b2 - a2*b1
	if determinant == 0 {
		return
	} else {
		x := (b2*c1 - b1*c2)/determinant
		y := (a1*c2 - a2*c1)/determinant
		return true, crd.NewCrd(x, y)
	}
}

func (s *Segment) SetStart(cr crd.Crd)  {
	s.a = cr
}

func (s *Segment) SetFinish(cr crd.Crd)  {
	s.b = cr
}

func (s Segment) Start() crd.Crd {
	return s.a
}

func (s Segment) Finish() crd.Crd {
	return s.b
}

func (s Segment) MidPoint() crd.Crd {
	return crd.NewCrd((s.a.X() + s.b.X()) / 2, (s.a.Y() + s.b.Y()) / 2)
}
