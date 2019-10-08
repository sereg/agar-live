package world

import (
	"agar-life/math/geom"
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
	"agar-life/world/const"
	"agar-life/world/frame"
	"agar-life/world/grid"
	"sort"
	"strconv"
)

type World struct {
	w, h       float64
	animal     frame.Frame
	plant      frame.Frame
	cycle      uint64
	gridPlant  grid.Grid
	gridAnimal grid.Grid
	resurrect  resurrects
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     frame.NewFrame(countAnimal),
		plant:      frame.NewFrame(countPlant),
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(behavior.NewAiv1(w, h))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6))
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.animal.Set(i, el)
	}
	for i := 0; i < countPlant; i++ {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)))
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.Set(i, el)
	}
	return world
}

func (w *World) Cycle() {
	first := true
	if w.cycle > 0 {
		first = false
		w.plant.SetUpdateState(false)
	}
	removeList := make(map[int]*frame.Frame)
	for i := 0; i < len(w.animal.All()); i++ {
		el := w.animal.Get(i).(animal.Animal)
		if el.GetDead() {
			continue
		}
		idCA, closestAnimal := getClosest(w.gridAnimal, el, w.animal, i)
		idCP, closestPlant := getClosest(w.gridPlant, el, w.plant, -1)
		closestAnimal = w.forIntersect(el, closestAnimal, idCA, &w.animal, removeList)
		closestPlant = w.forIntersect(el, closestPlant, idCP, &w.plant, removeList)
		var direction object.Crd
		dist := el.GetSpeed()
		if directionL, speed := el.GetInertia(); speed > 0 {
			direction, dist = directionL, speed
		} else {
			direction = el.GetDirection(closestAnimal, closestPlant, w.cycle)
		}
		el.SetCrdByDirection(el, direction, dist, first)
		w.fixLimit(el)
	}
	w.resurrect.resurrect(w.cycle, w.w, w.h)
	w.remove(removeList)
	w.gridAnimal.Reset()
	for i := 0; i < len(w.animal.All()); i++ {
		el := w.animal.Get(i)
		w.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
	}
	w.gridPlant.Reset()
	for i := 0; i < len(w.plant.All()); i++ {
		el := w.plant.Get(i)
		w.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
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
	if !w.plant.UpdateState() {
		return el
	}
	return w.plant.All()
}

func (w *World) GetAnimal() []alive.Alive {
	return w.animal.All()
}

func (w *World) remove(m map[int]*frame.Frame) {
	for _, v := range mapKeyToArray(m) {
		index, fr := v, m[v]
		w.resurrect.add(fr, fr.Get(index), w.cycle)
		fr.Delete(index)
	}
}

func mapKeyToArray(m map[int]*frame.Frame) []int {
	a := make([]int, len(m))
	ind := 0
	for index, _ := range m {
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

func (w *World) forIntersect(
	el animal.Animal,
	closest []alive.Alive,
	idInt []int,
	fr *frame.Frame,
	removeList map[int]*frame.Frame,
	) []alive.Alive {

	for j := 0; j < len(closest); j++ {
		el1 := closest[j]
		dist := func() float64 {
			return geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
		}
		index := idInt[j]
		if el != nil && el1 != nil && !el1.GetDead() &&
			(el.GetSize()/el1.GetSize() > _const.EatRatio || el.GetName() == el1.GetName()) &&
			dist() < el.GetSize() {
			el.Eat(el1)
			el1.Die()
			removeList[index] = fr
			fr.SetUpdateState(true)
			closest = removeFromAlive(closest, j)
			j--
		}
	}
	return closest
}

func getClosest(gr grid.Grid, el animal.Animal, fr frame.Frame, ind int) ([]int, []alive.Alive) {
	idInt := gr.GetObjInRadius(el.GetX(), el.GetY(), el.GetVision(), ind)
	closest := make([]alive.Alive, 0, len(idInt))
	for i := 0; i < len(idInt); i++ {
		id := idInt[i]
		closest = append(closest, fr.Get(id))
	}
	return idInt, closest
}

func removeFromAlive(a []alive.Alive, i int) []alive.Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func NewWorldTest(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     frame.NewFrame(countAnimal),
		plant:      frame.NewFrame(countPlant),
	}
	crAnimal := func(i int, x, y float64) {
		el := species.NewBeast(behavior.NewAiv1(w, h))
		//el := species.NewBeast(behavior.NewSimple(w, h))
		//gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(6), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.animal.Set(0, el)
	}
	crPlant := func(i int, x, y float64) {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.Set(i, el)
	}
	crAnimal(0, 50, 50)
	crPlant(0, 30, 50)
	crPlant(1, 70, 50)
	return world
}
