package behavior

import (
	"agar-life/math/vector"
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
	}
}

func getXYWithLength(x1, y1, x2, y2, dist float64) (x float64, y float64) {
	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
	length := vec.Len()
	ratio := dist / length
	vec.MultiplyByScalar(ratio)
	x, y = vec.GetPointFromVector(x2, y2)
	x, y = x-x2, y-y2
	return
}

func (s *simple) SetDirection(self animal.Animal, animals []alive.Alive, plants []alive.Alive) {
	if s.direction.GetX() == self.GetCrd().GetX() && s.direction.GetY() == self.GetCrd().GetY() {
		return
	}
	l := self.GetCrd()
	oldDirection := s.direction
	if l.GetX()+(self.GetSize()) >= s.w {
		s.direction.X(0)
	}
	if l.GetX()-(self.GetSize()) <= 0 {
		s.direction.X(s.w)
	}
	if l.GetY()+(self.GetSize()) >= s.h {
		s.direction.Y( 0)
	}
	if l.GetY()-(self.GetSize()) <= 0 {
		s.direction.Y(s.h)
	}
	if oldDirection != s.direction {
		s.changeDirection = true
	}
	if s.changeDirection {
		s.chCrd.Set(getXYWithLength(l.GetX(), l.GetY(), s.direction.GetX(), s.direction.GetY(), self.GetSpeed()))
		s.changeDirection = false
	}
	newX := l.GetX() + s.chCrd.GetX()
	newY := l.GetY() + s.chCrd.GetY()
	self.Crd(newX, newY)
}
