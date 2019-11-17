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
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := s.a.GetX(), s.a.GetY(), s.b.GetX(), s.b.GetY(), s1.a.GetX(), s1.a.GetY(), s1.b.GetX(), s1.b.GetY()
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
	a1 := s.Finish().GetY() - s.Start().GetY()
	b1 := s.Start().GetX() - s.Finish().GetX()
	c1 := a1*s.Start().GetX() + b1*s.Start().GetY()

	// Line CD represented as a2x + b2y = c2
	a2 := s1.Finish().GetY() - s1.Start().GetY()
	b2 := s1.Start().GetX() - s1.Finish().GetX()
	c2 := a2*s1.Start().GetX() + b2*s1.Start().GetY()

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
	return crd.NewCrd((s.a.GetX() + s.b.GetX()) / 2, (s.a.GetY() + s.b.GetY()) / 2)
}
