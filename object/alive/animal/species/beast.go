package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
)

type Beast struct {
	alive.Base
	speed    float64
	vision   float64
	behavior behavior.Behavior
}

func NewBeast(behavior behavior.Behavior) animal.Animal {
	return &Beast{
		speed:    0,
		vision:   0,
		behavior: behavior,
	}
}

func (b *Beast) Speed(speed float64) {
	b.speed = speed
}

func (b Beast) GetSpeed() float64 {
	return b.speed
}

func (b *Beast) Vision(vision float64) {
	b.vision = vision
}

func (b Beast) GetVision() float64 {
	return b.vision
}

func (b *Beast) Step(animals []alive.Alive, plants []alive.Alive) {
	b.behavior.SetDirection(animals, plants)
}

func (b *Beast) Eat(a alive.Alive) {
	if a.GetDead() {
		return
	}
}
