package animal

import (
	"agar-life/object/alive"
)

type base struct {
	alive.Base
	speed    float64
	vision   float64
}

func Newbase() Animal {
	return &base{
		speed:    0,
		vision:   0,
	}
}

func (b *base) Speed(speed float64) {
	b.speed = speed
}

func (b base) GetSpeed() float64 {
	return b.speed
}

func (b *base) Vision(vision float64) {
	b.vision = vision
}

func (b base) GetVision() float64 {
	return b.vision
}

func (b *base) Step(animals []alive.Alive, plants []alive.Alive) {
}

func (b *base) Eat(a alive.Alive) {
	if a.GetDead() {
		return
	}
}