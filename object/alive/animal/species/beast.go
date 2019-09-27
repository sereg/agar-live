package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
)

type beast struct {
	alive.Base
	speed    float64
	vision   float64
	behavior behavior.Behavior
}

func Newbeast(behavior behavior.Behavior) animal.Animal {
	return &beast{
		speed:    0,
		vision:   0,
		behavior: behavior,
	}
}

func (b *beast) Speed(speed float64) {
	b.speed = speed
}

func (b beast) GetSpeed() float64 {
	return b.speed
}

func (b *beast) Vision(vision float64) {
	b.vision = vision
}

func (b beast) GetVision() float64 {
	return b.vision
}

func (b *beast) Step(animals []alive.Alive, plants []alive.Alive) {
	b.behavior.SetDirection(b, animals, plants)
}

func (b *beast) Eat(a alive.Alive) {
	if a.GetDead() {
		return
	}
}
