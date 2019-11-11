package behavior

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"math"
)

type testAngel struct {
	direction crd.Crd
	angel     float64
	t float64
}

func NewTestAngel(angel float64) animal.Behavior {
	return &testAngel{angel: angel}
}

func (a *testAngel) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	//vec := vector.GetVectorByPoint(self.GetCrd(), crd.NewCrd(self.GetX()+self.GetVision(), self.GetY()))
	//vec.SetAngle(a.angel)
	//a.direction.SetCrd(vec.GetPointFromVector(self.GetCrd()))
	cx := 400.0
	cy := 200.0
	r := 200.0
	x:=cx+r * math.Cos(a.t)
	y:=cy+r * math.Sin(a.t)
	a.t++
	a.direction.SetXY(x, y)
	return a.direction, false
}
