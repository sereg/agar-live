package behavior

import (
	"agar-life/math"
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type simple struct {
	w, h            float64
	direction       object.Crd
	chCrd           object.Crd
	changeDirection bool
}

func NewSimple(w, h float64) Behavior {
	return &simple{
		w: w, h: h,
		changeDirection: true,
	}
}

func (s *simple) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	l := self.GetCrd()
	if change() || l.GetX()+(self.GetSize()) >= s.w {
		s.direction.X(0)
	}
	if change() || l.GetX()-(self.GetSize()) <= 0 {
		s.direction.X(s.w)
	}
	if change() || l.GetY()+(self.GetSize()) >= s.h {
		s.direction.Y( 0)
	}
	if change() || l.GetY()-(self.GetSize()) <= 0 {
		s.direction.Y(s.h)
	}
	return s.direction
}

func change() bool{
	return math.Random(0, 10000) > 9990
}
