package move

import (
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object/alive"
	_const "agar-life/world/const"
)

type Move struct {
	chCrd        crd.Crd
	oldDirection crd.Crd
	oldDist         float64
	inertia
}

type inertia struct{
	direction           crd.Crd
	speed, acceleration float64
}

func (i *inertia) SetInertia(direction crd.Crd) {
	i.direction = direction
	i.acceleration = _const.SplitDeceleration
	i.speed = _const.SplitSpeed
}

func (m *Move) GetInertia() (direction crd.Crd, speed float64){
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

func (m *Move) GetDirection() crd.Crd {
	return m.oldDirection
}

func (m *Move) SetCrdByDirection(a alive.Alive, direction crd.Crd, dist float64, changeDirection bool) {
	if changeDirection || direction != m.oldDirection || m.oldDist != dist{
		c := vector.GetCrdWithLength(a.GetCrd(), direction, dist)
		xDif, yDif := c.X()-a.X(), c.Y()-a.Y()
		m.chCrd.SetXY(xDif, yDif)
	}
	m.oldDirection = direction
	m.oldDist = dist
	a.SetX(a.X() + m.chCrd.X())
	a.SetY(a.Y() + m.chCrd.Y())
}
