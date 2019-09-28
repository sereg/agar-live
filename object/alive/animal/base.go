package animal

import (
	"agar-life/object/alive"
	"math"
)

type Base struct {
	alive.Base
	speed  float64
	vision float64
}

func NewBase() Animal {
	return &Base{
		speed:  0,
		vision: 0,
	}
}

func (b *Base) Size(size float64) {
	b.Base.Size(size)
	b.Speed(StartSpeed - (math.Log(size * SpeedRatio)))
	b.Vision(StartVision + b.GetSize()*(VisionRatio-math.Log(b.GetSize())))
}

func (b *Base) Speed(speed float64) {
	b.speed = speed
}

func (b Base) GetSpeed() float64 {
	return b.speed
}

func (b *Base) Vision(vision float64) {
	b.vision = vision
}

func (b Base) GetVision() float64 {
	return b.vision
}

func (b *Base) Step(animals []alive.Alive, plants []alive.Alive) {
}

func (b *Base) Eat(el alive.Alive) {
	if el.GetDead() {
		return
	}
	b.Size(b.GetSize() + (el.GetSize() * EatIncreaseRation))
}
