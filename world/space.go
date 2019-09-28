package world

import (
	"agar-life/math/geom"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
	"strconv"
)

const (
	gridSize = 20
)

type frame struct {
	deedIndex   int
	updateState bool
	el          []alive.Alive
}

type World struct {
	w, h       float64
	animal     frame
	plant      frame
	cycle      int64
	gridPlant  grid
	gridAnimal grid
	resurrect  resurrects
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  NewGrid(gridSize),
		gridAnimal: NewGrid(gridSize),
		animal:     frame{el: make([]alive.Alive, countAnimal)},
		plant:      frame{el: make([]alive.Alive, countPlant), updateState: true},
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(behavior.NewSimple(w, h))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6))
		world.gridAnimal.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		world.animal.el[i] = el
	}
	for i := 0; i < countPlant; i++ {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)))
		world.gridPlant.set(el.GetCrd().GetX(), el.GetCrd().GetY(), i)
		world.plant.el[i] = el
	}
	return world
}

func (w *World) Cycle() {
	if w.cycle > 0 {
		w.plant.updateState = false
	}
	for i := 0; i < len(w.animal.el)-w.animal.deedIndex; i++ {
		el := w.animal.el[i].(animal.Animal)
		if el.GetDead() {
			continue
		}
		idCA, closestAnimal := getClosest(w.gridAnimal, el, w.animal)
		idCP, closestPlant := getClosest(w.gridPlant, el, w.plant)
		idCA, closestAnimal = w.forIntersect(el, closestAnimal, idCA, &w.animal)
		idCP, closestPlant = w.forIntersect(el, closestPlant, idCP, &w.plant)
		el.Step(closestAnimal, closestPlant)
		w.fixLimit(el)
	}
	w.resurrect.resurrect(w.cycle, w.w, w.h)
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

func (w *World) fixLimit(el animal.Animal) {
	x, y := el.GetCrd().GetX(), el.GetCrd().GetY()
	if x < 0 {
		x = 0
	}
	if x > w.w {
		x = w.w
	}
	if y < 0 {
		y = 0
	}
	if y > w.h {
		y = w.h
	}
	el.Crd(x, y)
}

func (w *World) GetPlant() []alive.Alive {
	var el []alive.Alive
	if !w.plant.updateState {
		return el
	}
	return w.plant.el[:len(w.plant.el)-w.plant.deedIndex]
}

func (w *World) GetAnimal() []alive.Alive {
	return w.animal.el[:len(w.animal.el)-w.animal.deedIndex]
}

func (w *World) forIntersect(el animal.Animal, closest []alive.Alive, idInt []int, fr *frame) ([]int, []alive.Alive) {
	for j := 0; j < len(closest); j++ {
		el1 := closest[j]
		dist := func() float64 {
			return geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
		}
		if el != nil && el1 != nil && !el1.GetDead() && el.GetSize()/el1.GetSize() > EatRatio && dist() < el.GetSize() {
			index := idInt[j]
			el.Eat(el1)
			el1.Die()
			deedIndex := len(fr.el) - 1 - fr.deedIndex
			fr.el[index], fr.el[deedIndex] = fr.el[deedIndex], fr.el[index]
			w.resurrect.add(fr, deedIndex, w.cycle)
			fr.deedIndex++
			fr.updateState = true
			closest = removeFromAlive(closest, j)
			idInt = removeFromInt(idInt, j)
			j--
		}
	}
	return idInt, closest
}

func getClosest(gr grid, el animal.Animal, fr frame) ([]int, []alive.Alive) {
	idInt := gr.GetObjInVision(el.GetCrd().GetX(), el.GetCrd().GetY(), el.GetVision())
	lenClosest := len(idInt)
	if len(fr.el) > 0 {
		if _, ok := fr.el[0].(animal.Animal); ok {
			lenClosest--
		}
	}
	closest := make([]alive.Alive, lenClosest)
	left := 0
	for i := 0; i < lenClosest; i++ {
		id := idInt[i-left]
		if el.GetName() == fr.el[id].GetName() {
			left = 1
			continue
		}
		closest[i] = fr.el[id]
	}
	return idInt, closest
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0    // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func removeFromAlive(a []alive.Alive, i int) []alive.Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func getIDByName(a []alive.Alive, name string) int {
	for i, v := range a {
		if v.GetName() == name {
			return i
		}
	}
	return -1
}
