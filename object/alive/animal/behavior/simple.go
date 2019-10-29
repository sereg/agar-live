package behavior

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior/checkangels"
	"math"
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

func (s *Simple) Action(self animal.Animal, animals, plants []alive.Alive, cycle uint64, dAngeles checkangels.Angels) (crd.Crd, bool) {
	if s.direction.X() == 0 && s.direction.Y() == 0 {
		s.SetDirection(self, crd.NewCrd(float64(math2.Random(0, int(s.w))), float64(math2.Random(0, int(s.h)))))
	}
	vec := vector.GetVectorByPoint(self.GetCrd(), s.direction)
	if change() {
		part := 6.0
		addAngel := randomFloat(-1 * math.Pi / part, math.Pi / part)
		vec.AddAngle(addAngel)
	}
	vecAngel := geom.ModuleDegree(vec.GetAngle())
	reachable, _ := dAngeles.Check(vecAngel, vec.Len())
	if !reachable {
		angel := dAngeles.ClosestAvailable(vecAngel)
		vec.SetAngle(angel)
	}
	return vec.GetPointFromVector(self.GetCrd()), false
}

func randomFloat(min, max float64) float64 {
	minInt := int(min * 100)
	maxInt := int(max * 100)
	rand := math2.Random(minInt, maxInt)
	return float64(rand) / 100
}

func change() bool {
	return math2.Random(0, 100) > 90
}
