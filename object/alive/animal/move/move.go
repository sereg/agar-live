package move

import (
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive"
	_const "agar-life/world/const"
)

type Move struct {
	chCrd        object.Crd
	oldDirection object.Crd
	oldDist         float64
	inertia
}

type inertia struct{
	direction           object.Crd
	speed, acceleration float64
}

func (i *inertia) SetInertia(direction object.Crd) {
	i.direction = direction
	i.acceleration = _const.SplitDeceleration
	i.speed = _const.SplitSpeed
}

func (m *Move) GetInertia() (direction object.Crd, speed float64){
	direction = m.direction
	speed = m.speed
	if m.speed > 0 {
		m.speed -= m.acceleration
	}
	if m.speed < 0 {
		m.speed = 0
	}
	return
}

func getXYWithLength(x1, y1, x2, y2, dist float64) (x float64, y float64) {
	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
	length := vec.Len()
	ratio := dist
	if length > 0 {
		ratio = dist / length
	}
	vec.MultiplyByScalar(ratio)
	x, y = vec.GetPointFromVector(x2, y2)
	x, y = x-x2, y-y2
	return
}

func (m *Move) GetDirection() object.Crd {
	return m.oldDirection
}

func (m *Move) SetCrdByDirection(a alive.Alive, direction object.Crd, dist float64, changeDirection bool) {
	if changeDirection || direction != m.oldDirection || m.oldDist != dist{
		m.chCrd.SetCrd(getXYWithLength(a.GetX(), a.GetY(), direction.GetX(), direction.GetY(), dist))
	}
	m.oldDirection = direction
	m.oldDist = dist
	newX := a.GetX() + m.chCrd.GetX()
	newY := a.GetY() + m.chCrd.GetY()
	a.SetCrd(newX, newY)
}
