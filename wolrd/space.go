package wolrd

import (
	"agar-life/math/geom"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	"agar-life/object/generate"
	"strconv"
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

func (g *grid) GetObjInVision(x, y, vision float64) []int {
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
		generate.Generate(el, w, h, "a" + strconv.Itoa(i))
		word.gridAnimal.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		word.animal.el[i] = el
	}
	for i := 0; i < countPlant; i++ {
		el := sp.NewPlant()
		generate.Generate(el, w, h, "p" + strconv.Itoa(i))
		word.gridPlant.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		word.plant.el[i] = el
	}
	return word
}

func (w *World) Cycle() {
	for i := 0; i < len(w.animal.el)-w.animal.deedIndex; i++ {
		el := w.animal.el[i].(animal.Animal)
		if el.GetDead() {
			continue
		}
		idCA, closestAnimal := getClosest(w.gridAnimal, el, w.animal)
		idCP, closestPlant := getClosest(w.gridPlant, el, w.plant)
		idCA, closestAnimal = forIntersect(el, closestAnimal, w.cycle, idCA, w.animal)
		idCP, closestPlant = forIntersect(el, closestPlant, w.cycle, idCP, w.plant)
		el.Step(closestAnimal, closestPlant)
	}
	w.gridAnimal = NewGrid(gridSize)
	for i := 0; i < len(w.animal.el)-w.animal.deedIndex; i++ {
		el := w.animal.el[i]
		w.gridAnimal.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
	}
	w.gridPlant = NewGrid(gridSize)
	for i := 0; i < len(w.plant.el)-w.plant.deedIndex; i++ {
		el := w.plant.el[i]
		w.gridPlant.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
	}
	w.cycle++
}

func getClosest(gr grid, el animal.Animal, fr frame) ([]int, []alive.Alive){
	idInt := gr.GetObjInVision(el.GetCrd().GetX(), el.GetCrd().GetY(), el.GetVision())
	closest := make([]alive.Alive, len(idInt))
	for i, id := range idInt {
		closest[i] = fr.el[id]
	}
	return idInt, closest
}

func forIntersect(el animal.Animal, closest []alive.Alive, cycle int64, idInt []int, fr frame) ([]int, []alive.Alive) {
	for j := 0; j < len(closest); j++ {
		el1 := closest[j]
		if el.GetName() == el1.GetName() || intersect(el, el1, cycle, idInt[j], fr) {
			closest = removeFromAliveByName(closest, j)
			idInt = removeFromInt(idInt, j)
			j--
		}
	}
	return idInt, closest
}

func removeFromInt(a []int, i int) []int{
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func removeFromAliveByName(a []alive.Alive, i int) []alive.Alive{
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func intersect(el animal.Animal, el1 alive.Alive, cycle int64, index int, container frame) bool {
	if el.GetName() == el1.GetName() {
		return false
	}
	dist := geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
	if dist < el.GetSize() && el.GetSize()/el1.GetSize() > EatRatio {
		el.Eat(el1)
		el1.Die()
		deedIndex := len(container.el) - 1 - container.deedIndex
		container.el[index], container.el[deedIndex] = container.el[deedIndex], container.el[index]
		container.deedIndex++
		container.updateState = true
		return true
	}
	return false
}

func getIDByName(a []alive.Alive, name string) int{
	for i, v := range a {
		if v.GetName() == name {
			return i
		}
	}
	return -1
}
