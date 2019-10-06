package world

import (
	"agar-life/math/geom"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
	"agar-life/world/const"
	"agar-life/world/grid"
	"sort"
	"strconv"
)

type Frame struct {
	updateState bool
	el          [][]alive.Alive
}

func (f *Frame) El() [][]alive.Alive{
	return f.el
}

type World struct {
	w, h       float64
	animal     Frame
	plant      Frame
	cycle      uint64
	gridPlant  grid.Grid
	gridAnimal grid.Grid
	resurrect  resurrects
}

func NewWorldTest(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     Frame{el: make([][]alive.Alive, countAnimal)},
		plant:      Frame{el: make([][]alive.Alive, countPlant), updateState: true},
	}
	crAnimal := func(i int, x, y float64) {
		el := species.NewBeast(behavior.NewAiv1(w, h))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.animal.el[0] = []alive.Alive{el}
	}
	crPlant := func(i int, x, y float64) {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.el[i] = []alive.Alive{el}
	}
	crAnimal(0, 50, 50)
	crPlant(0, 30, 50)
	crPlant(1, 70, 50)
	return world
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     Frame{el: make([][]alive.Alive, countAnimal)},
		plant:      Frame{el: make([][]alive.Alive, countPlant), updateState: true},
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(behavior.NewAiv1(w, h))
		//el := species.NewBeast(behavior.NewSimple(w, h))
		//gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6))
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.animal.el[i] = []alive.Alive{el}
	}
	for i := 0; i < countPlant; i++ {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)))
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.el[i] = []alive.Alive{el}
	}
	return world
}

func (w *World) Cycle() {
	if w.cycle > 0 {
		w.plant.updateState = false
	}
	killList:= make(map[int]*Frame)
	for i := 0; i < len(w.animal.el); i++ {
		el := w.animal.el[i][0].(animal.Animal)
		if el.GetDead() {
			continue
		}
		idCA, closestAnimal := getClosest(w.gridAnimal, el, w.animal, i)
		idCP, closestPlant := getClosest(w.gridPlant, el, w.plant, -1)
		for _, an := range w.animal.el[i] {
			el = an.(animal.Animal)
			closestAnimal = w.forIntersect(el, closestAnimal, idCA, &w.animal, killList)
			closestPlant = w.forIntersect(el, closestPlant, idCP, &w.plant, killList)
		}
		el.Step(closestAnimal, closestPlant, w.cycle)
		w.fixLimit(el)
	}
	w.resurrect.resurrect(w.cycle, w.w, w.h)
	w.kill(killList)
	w.gridAnimal.Reset()
	for i := 0; i < len(w.animal.el); i++ {
		el := w.animal.el[i]
		w.gridAnimal.Set(el[0].GetX(), el[0].GetY(), el[0].GetSize(), i)
	}
	w.gridPlant.Reset()
	for i := 0; i < len(w.plant.el); i++ {
		el := w.plant.el[i]
		w.gridPlant.Set(el[0].GetX(), el[0].GetY(), el[0].GetSize(), i)
	}
	w.cycle++
}

func (w *World) fixLimit(el animal.Animal) {
	x, y := el.GetX(), el.GetY()
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
	el.SetCrd(x, y)
}

func (w *World) GetPlant() []alive.Alive {
	var el []alive.Alive
	if !w.plant.updateState {
		return el
	}
	al := make([]alive.Alive, 0, len(w.plant.el))
	for _, v := range w.plant.el {
		al = append(al, v...)
	}
	return al
}

func (w *World) GetAnimal() []alive.Alive {
	al := make([]alive.Alive, 0, len(w.animal.el))
	for _, v := range w.animal.el {
		al = append(al, v...)
	}
	return al
}

func (w *World) kill(m map[int]*Frame) {
	for _, v := range mapKeyToArray(m){
		index, fr := v, m[v]
		w.resurrect.add(fr, fr.el[index][0], w.cycle)
		fr.el = removeFromAliveArray(fr.el, index)
	}
}

func mapKeyToArray(m map[int]*Frame) []int{
	a := make([]int, len(m))
	ind :=0
	for index, _ := range m{
		a[ind] = index
		ind++
	}
	sort.Sort(intSlice(a))
	return a
}

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (w *World) forIntersect(el animal.Animal, closest []alive.Alive, idInt []int, fr *Frame, killList map[int]*Frame) []alive.Alive {
	prev := -1
	indexEl := 0
	for j := 0; j < len(closest); j++ {
		el1 := closest[j]
		dist := func() float64 {
			return geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
		}
		index := idInt[j]
		if index != prev {
			indexEl = 0
		} else {
			indexEl++
		}
		prev = index
		if el != nil && el1 != nil && !el1.GetDead() && el.GetSize()/el1.GetSize() > _const.EatRatio && dist() < el.GetSize() {
			el.Eat(el1)
			if len(fr.el[index]) == 1 {
				el1.Die()
				killList[index] = fr
			} else {
				fr.el[index] = removeFromAlive(fr.el[index], indexEl)
				indexEl--
			}
			fr.updateState = true
			closest = removeFromAlive(closest, j)
			j--
		}
	}
	return closest
}

func getClosest(gr grid.Grid, el animal.Animal, fr Frame, ind int) ([]int, []alive.Alive) {
	idInt := gr.GetObjInRadius(el.GetX(), el.GetY(), el.GetVision(), ind)
	closest := make([]alive.Alive, 0, len(idInt))
	for i := 0; i < len(idInt); i++ {
		id := idInt[i]
		closest = append(closest, fr.el[id]...)
	}
	return idInt, closest
}

func removeFromAliveArray(a [][]alive.Alive, i int) [][]alive.Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func removeFromAlive(a []alive.Alive, i int) []alive.Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0    // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}