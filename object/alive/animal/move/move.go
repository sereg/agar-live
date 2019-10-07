package move

import (
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive"
)

type Move struct {
	chCrd        object.Crd
	oldDirection object.Crd
	oldDist         float64
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

func (s *Move) SetCrdByDirection(a alive.Alive, direction object.Crd, dist float64, changeDirection bool) {
	if changeDirection || direction != s.oldDirection || s.oldDist != dist{
		s.chCrd.SetCrd(getXYWithLength(a.GetX(), a.GetY(), direction.GetX(), direction.GetY(), dist))
	}
	s.oldDirection = direction
	s.oldDist = dist
	newX := a.GetX() + s.chCrd.GetX()
	newY := a.GetY() + s.chCrd.GetY()
	a.SetCrd(newX, newY)
}
