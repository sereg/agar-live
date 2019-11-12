package move

import (
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object/alive"
	_const "agar-life/world/const"
)

type Move struct {
	ChCrd        crd.Crd
	OldDirection crd.Crd
	OldDist      float64
	Inertia Inertia
}

type Inertia struct{
	Direction           crd.Crd
	Speed, Acceleration float64
}

func (i *Inertia) SetInertia(direction crd.Crd) {
	i.Direction = direction
	i.Acceleration = _const.SplitDeceleration
	i.Speed = _const.SplitSpeed
}

func (m *Move) SetInertia(direction crd.Crd) {
	m.Inertia.SetInertia(direction)
}

func (m *Move) SetInertiaImport(direction crd.Crd, speed, acceleration float64) {
	m.Inertia.Direction = direction
	m.Inertia.Speed = speed
	m.Inertia.Acceleration = acceleration
}

func (m *Move) GetInertia() (direction crd.Crd, speed float64){
	direction = m.Inertia.Direction
	speed = m.Inertia.Speed
	if m.Inertia.Speed > 0 {
		m.Inertia.Speed -= m.Inertia.Acceleration
	}
	if m.Inertia.Speed < 0 {
		m.Inertia.Speed = 0
	}
	return
}

func (m *Move) GetDirection() crd.Crd {
	return m.OldDirection
}

func (m *Move) SetCrdByDirection(a alive.Alive, direction crd.Crd, dist float64, changeDirection bool) {
	if direction == a.GetCrd() {
		return
	}
	if changeDirection || direction != m.OldDirection || m.OldDist != dist{
		c := vector.GetCrdWithLength(a.GetCrd(), direction, dist)
		xDif, yDif := c.GetX()-a.GetX(), c.GetY()-a.GetY()
		m.ChCrd.SetXY(xDif, yDif)
	}
	m.OldDirection = direction
	m.OldDist = dist
	a.SetXY(a.GetX() + m.ChCrd.GetX(), a.GetY() + m.ChCrd.GetY())
}
