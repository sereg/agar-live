package animal

import (
	"agar-life/object"
	"agar-life/object/alive"
)

type Animal interface {
	alive.Alive
	Speed(float64)
	GetSpeed() float64
	Vision(float64)
	GetVision() float64
	Eat(a alive.Alive)
	GetDirection(animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd
	SetCrdByDirection(a alive.Alive, direction object.Crd, dist float64, changeDirection bool)
}
