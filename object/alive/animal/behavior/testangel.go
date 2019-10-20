package behavior

import (
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type testAngel struct {
	direction crd.Crd
	angel     float64
}

func NewTestAngel(angel float64) animal.Behavior {
	return &testAngel{angel: angel}
}

func (a *testAngel) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	vec := vector.GetVectorByPoint(self.GetCrd(), crd.NewCrd(self.X()+self.Vision(), self.Y()))
	vec.SetAngle(a.angel)
	a.direction.SetCrd(vec.GetPointFromVector(self.GetCrd()))

	return a.direction, false
}
