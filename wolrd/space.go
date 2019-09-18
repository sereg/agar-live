package wolrd

import (
	"agar-life/object/alive/animal"
	"agar-life/object/alive/plant"
)

type World struct{
	w, h    float64
	animal animal.Animal
	plant plant.Plant
	cycle   int64
}

func NewWorld() {

}

func (w *World) Cycle() {
	w.cycle++
}
