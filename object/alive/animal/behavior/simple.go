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
		changeDirection: true,
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
	if s.direction.GetX() == self.GetX() && s.direction.GetY() == self.GetY() {
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
	s.setCrdByDirection(self, oldDirection)
}

func (s *simple) setCrdByDirection(a animal.Animal, oldDirection object.Crd) {
	if s.changeDirection || s.direction != oldDirection {
		s.chCrd.SetCrd(getXYWithLength(a.GetX(), a.GetY(), s.direction.GetX(), s.direction.GetY(), a.GetSpeed()))
		s.changeDirection = false
	}
	newX := a.GetX() + s.chCrd.GetX()
	newY := a.GetY() + s.chCrd.GetY()
	a.SetCrd(newX, newY)
}
