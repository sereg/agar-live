package animal

import (
	"agar-life/object/alive"
)

type Animal interface {
	alive.Alive
	Speed(float64)
	GetSpeed() float64
	Vision(float64)
	GetVision() float64
	Eat(a alive.Alive)
	Step(animals []alive.Alive, plants []alive.Alive)
}
