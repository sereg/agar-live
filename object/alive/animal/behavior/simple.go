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

func NewSimple(w, h float64) animal.Behavior {
	return &simple{
		w: w, h: h,
		changeDirection: true,
	}
}


func (a *simple) GetDirection() object.Crd {
	return a.direction
}

func (s *simple) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	l := self.GetCrd()
	if change() || l.GetX()+(self.Size()) >= s.w {
		s.direction.X(0)
	}
	if change() || l.GetX()-(self.Size()) <= 0 {
		s.direction.X(s.w)
	}
	if change() || l.GetY()+(self.Size()) >= s.h {
		s.direction.Y( 0)
	}
	if change() || l.GetY()-(self.Size()) <= 0 {
		s.direction.Y(s.h)
	}
	return s.direction
}

func change() bool{
	return math.Random(0, 10000) > 9990
}
