package behavior

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type Simple struct {
	w, h            float64
	direction       crd.Crd
	chCrd           crd.Crd
	changeDirection bool
}

func (s *Simple) SetSize(w, h float64) {
	s.w, s.h = w, h
}

func (s *Simple) SetDir(direction crd.Crd) {
	s.direction = direction
}

func (s *Simple) Dir() crd.Crd {
	return s.direction
}

func (s *Simple) SetChangeDir(changeDirection bool) {
	s.changeDirection = changeDirection
}

func (s Simple) W() float64 {
	return s.w
}

func (s Simple) H() float64 {
	return s.h
}

func NewSimple(w, h float64) Simple {
	s := Simple{
		w: w, h: h,
		changeDirection: true,
	}
	return s
}

func (s *Simple) SetDirection(self animal.Animal, direction crd.Crd) {
	s.direction = vector.GetCrdWithLength(self.GetCrd(), s.direction, self.Vision())
}

func (s *Simple) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	if s.direction.X() == 0 && s.direction.Y() == 0 {
		s.SetDirection(self, crd.NewCrd(float64(math2.Random(0, int(s.w))), float64(math2.Random(0, int(s.h)))))
	}
	if change() || self.X()+(self.Size()) >= s.w {
		s.SetDirection(self, crd.NewCrd(0, self.Y()))
	}
	if change() || self.X()-(self.Size()) <= 0 {
		s.SetDirection(self, crd.NewCrd(s.w, self.Y()))
	}
	if change() || self.Y()+(self.Size()) >= s.h {
		s.SetDirection(self, crd.NewCrd(self.X(), 0))
	}
	if change() || self.Y()-(self.Size()) <= 0 {
		s.SetDirection(self, crd.NewCrd(self.X(), s.h))
	}
	return s.direction, false
}

func change() bool {
	return math2.Random(0, 10000) > 9990
}
