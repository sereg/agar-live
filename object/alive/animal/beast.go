package animal

import "agar-life/object/alive"

type Beast struct{
	alive.Base
	speed float64
	vision float64
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

func (b *Beast) Eat(a alive.Alive) {

}
