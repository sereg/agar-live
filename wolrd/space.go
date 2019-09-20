package wolrd

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	"agar-life/object/generate"
)

const (
	gridSize = 15
)

type frame struct {
	deedIndex   int
	updateState bool
	el          []alive.Alive
}

type xy struct {
	x, y int
}

type grid struct {
	cellSize float64
	data     map[xy][]int
}

func NewGrid(size float64) grid {
	return grid{
		cellSize: size,
	}
}

func (g *grid) reset() {
	g.data = make(map[xy][]int)
}

func (g *grid) set(x, y float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	if _, ok := g.data[xy{x: xInt, y: yInt}]; ok {
		g.data[xy{x: xInt, y: yInt}] = []int{i}
	} else {
		g.data[xy{x: xInt, y: yInt}] = append(g.data[xy{x: xInt, y: yInt}], i)
	}
}

func (g *grid) getObjInVision(x, y, vision float64) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	var obj []int
	for cx:= ltx; cx < rdx; cx++ {
		for cy:= lty; cy < rdy; cy++ {
			if ob, ok := g.data[xy{cx, cy}]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj
}

type World struct {
	w, h       float64
	animal     frame
	plant      frame
	cycle      int64
	gridPlant  grid
	gridAnimal grid
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	word := World{
		w:          w,
		h:          h,
		gridPlant:  NewGrid(gridSize),
		gridAnimal: NewGrid(gridSize),
		animal:     frame{el: make([]alive.Alive, countAnimal)},
		plant:      frame{el: make([]alive.Alive, countPlant)},
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(behavior.NewSimple())
		generate.Generate(word.animal.el[i], w, h)
		word.gridAnimal.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		word.animal.el[i] = el
	}
	for i := 0; i < countPlant; i++ {
		el := sp.NewPlant()
		generate.Generate(word.plant.el[i], w, h)
		word.gridPlant.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		word.plant.el[i] = el
	}
	return word
}

func (w *World) Cycle() {
	gridPlant := NewGrid(gridSize)
	gridAnimal := NewGrid(gridSize)
	for i := 0; i < len(w.animal.el)-w.animal.deedIndex; i++ {
		el := w.animal.el[i].(animal.Animal)
		if el.GetDead() {
			continue
		}
		idClosestAnimal := w.gridAnimal.getObjInVision(el.GetCrd().GetX(), el.GetCrd().GetY(), el.GetVision())
		closestAnimal := make([]alive.Alive, len(idClosestAnimal))
		for i, id := range idClosestAnimal {
			closestAnimal[i] = w.animal.el[id]
		}
	}
	w.gridAnimal = gridAnimal
	w.gridPlant = gridPlant
	w.cycle++
}
