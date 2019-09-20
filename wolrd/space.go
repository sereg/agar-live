package wolrd

import (
	"agar-life/object/alive"
)

type frame struct {
	deedIndex   int
	updateState bool
	el          []alive.Alive
}

type xy struct {
	x, y int
}

type World struct {
	w, h       float64
	animal     frame
	plant      frame
	cycle      int64
	gridPlant  map[xy]int
	gridAnimal map[xy]int
}

func NewWorld() {

}

func (w *World) Cycle() {
	for i := 0; i < len(w.animal.el)-w.animal.deedIndex; i++ {

	}
	w.cycle++
}
