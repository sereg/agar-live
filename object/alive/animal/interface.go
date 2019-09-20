package animal

import (
	"agar-life/object/alive"
	"agar-life/object/alive/plant"
)

type Animal interface {
	alive.Alive
	Speed(float64)
	GetSpeed() float64
	Vision(float64)
	GetVision() float64
	Eat(a alive.Alive)
	Step(animals []Animal, plants []plant.Plant)
}
