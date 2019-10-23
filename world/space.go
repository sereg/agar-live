package world

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/behavior/ai"
	"agar-life/object/alive/animal/species"
	"agar-life/object/alive/plant"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
	"agar-life/world/const"
	"agar-life/world/frame"
	"agar-life/world/grid"
	"math"
	"sort"
	"strconv"
)

type World struct {
	w, h       float64
	animal     frame.Frame
	plant      frame.Frame
	cycle      uint64
	gridPlant  grid.Grid //TODO move grid index to frame
	gridAnimal grid.Grid //TODO test Quadtree https://github.com/JamesMilnerUK/quadtree-go
	resurrect  resurrects
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     frame.NewFrame(countAnimal, w, h),
		plant:      frame.NewFrame(countPlant, w, h),
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(ai.NewAiv1(w, h))
		//el := species.NewBeast(behavior.NewSimple(w, h))
		//gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(10 + float64(math2.Random(0, 40))))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(_const.AliveStartSize))
		world.gridAnimal.Set(el.X(), el.Y(), el.Size(), i)
		world.animal.Set(i, el)
	}
	for i := 0; i < countPlant; i++ {
		var el plant.Plant
		if poison() {
			el = sp.NewPoison()
		} else {
			el = sp.NewPlant()
		}
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)))
		world.gridPlant.Set(el.X(), el.Y(), el.Size(), i)
		world.plant.Set(i, el)
	}
	return world
}

func poison() bool {
	//return false
	return math2.Random(0, 10) == 9
}

func (w *World) Cycle() {
	first := true
	if w.cycle > 0 {
		first = false
		w.plant.SetUpdateState(false)
	}
	removeList := rmList{}
	for i := 0; i < len(w.animal.All()); i++ {
		el := w.animal.Get(i).(animal.Animal)
		if el.GetDead() {
			continue
		}
		idCA, closestAnimal := getClosest(w.gridAnimal, el, w.animal, i)
		idCP, closestPlant := getClosest(w.gridPlant, el, w.plant, -1)
		closestAnimal = w.forIntersect(el, closestAnimal, idCA, &w.animal, &removeList)
		closestPlant = w.forIntersect(el, closestPlant, idCP, &w.plant, &removeList)
		var direction crd.Crd
		split := false
		dist := el.Speed()
		if directionL, speed := el.GetInertia(); speed > 0 {
			direction, dist = directionL, speed
			el.SetCrdByDirection(el, direction, dist, first)
		} else {
			direction, split = el.Action(closestAnimal, closestPlant, w.cycle)
			if split {
				Split(&w.animal, el, direction, w.cycle)
			}
			el.SetCrdByDirection(el, direction, dist, first)
			w.fixNeighborhood(el)
		}
		w.fixLimit(el)
	}
	w.resurrect.resurrect(w.cycle, w.w, w.h)
	w.remove(removeList)
	w.gridAnimal.Reset()
	for i := 0; i < len(w.animal.All()); i++ {
		el := w.animal.Get(i)
		w.gridAnimal.Set(el.X(), el.Y(), el.Size(), i)
	}
	if w.plant.UpdateState() {
		w.gridPlant.Reset()
		for i := 0; i < len(w.plant.All()); i++ {
			el := w.plant.Get(i)
			w.gridPlant.Set(el.X(), el.Y(), el.Size(), i)
		}
	}
	w.cycle++
}

func (w *World) fixNeighborhood(el animal.Animal) {
	var parent animal.Animal
	if parent = el.Parent(); parent == nil {
		return
	}
	for _, v := range parent.Children() {
		var el1 animal.Animal
		if v.ID() == el.ID() {
			el1 = parent
		} else {
			el1 = v
		}
		if el.GlueTime() <= w.cycle && el1.GlueTime() <= w.cycle {
			continue
		}
		dis := geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
		if dis < el.Size()+el1.Size() {
			dist := el.Size() + el1.Size() - dis
			vec := vector.GetVectorByPoint(el.GetCrd(), el1.GetCrd())
			vec.AddAngle(math.Pi)
			direction := vec.GetPointFromVector(el.GetCrd())
			el.SetCrdByDirection(el, direction, dist, true)
		}
	}
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

type remove struct {
	index int
	fr    *frame.Frame
}

type rmList struct {
	list []remove
}

func (r *rmList) add(index int, fr *frame.Frame) {
	r.list = append(r.list, remove{
		index: index,
		fr:    fr,
	})
}

func (r rmList) Len() int           { return len(r.list) }
func (r rmList) Less(i, j int) bool { return r.list[i].index > r.list[j].index }
func (r rmList) Swap(i, j int)      { r.list[i], r.list[j] = r.list[j], r.list[i] }

func (w *World) remove(m rmList) {
	if len(m.list) > 1 {
		sort.Sort(m)
	}
	for i:=0; i < len(m.list); i++ {
		v := m.list[i]
		index, fr := v.index, v.fr
		if el, ok := fr.Get(index).(animal.Animal); ok {
			if el.Parent() == nil && len(el.Children()) == 0 {
				w.resurrect.add(fr, fr.Get(index), w.cycle)
			}
		} else {
			w.resurrect.add(fr, fr.Get(index), w.cycle)
		}
		fr.Delete(index)
	}
}

func (w *World) forIntersect(
	el animal.Animal,
	closest []alive.Alive,
	idInt []int,
	fr *frame.Frame,
	removeList *rmList,
) []alive.Alive {
	removedId := map[int]struct{}{}
	for j := 0; j < len(closest); j++ {
		index := idInt[j]
		el1 := closest[j]
		dis := -1.0
		dist := func() float64 {
			if dis == -1.0 {
				dis = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
			}
			return dis
		}
		died := false
		if el != nil && el1 != nil && !el1.GetDead() {
			if _, ok := removedId[index]; ok {
				died = true
			}
			if !died && (el.Size()/el1.Size() > _const.EatRatio || (el.Group() == el1.Group() &&
				el1.GlueTime() <= w.cycle && el.GlueTime() <= w.cycle)) && !el1.Danger() && dist() < el.Size() {
				died = true
				el.Eat(el1) //TODO change size in 30 cycles
			}
			if !died && el1.Danger() && el1.Size() < el.Size() && dist() < el.Size() {
				if Burst(&w.animal, el, w.cycle) {
					died = true
				}
			}
			if died {
				el1.Die()
				if _, ok := removedId[index]; !ok {
					removeList.add(index, fr)
					removedId[index] = struct{}{}
				}
				fr.SetUpdateState(true)
				closest = alive.Remove(closest, j)
				idInt = removeFromInt(idInt, j)
				j--
			}
		}
		if !died && dist() > el.Size() + _const.GridSize {
			break
		}
		//if !died && el.Group() == el1.Group() {
		//	closest = alive.Remove(closest, j)
		//	idInt = removeFromInt(idInt, j)
		//	j--
		//}
	}
	return closest
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0    // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func getClosest(gr grid.Grid, el animal.Animal, fr frame.Frame, ind int) ([]int, []alive.Alive) {
	idInt := gr.GetObjInRadius(el.X(), el.Y(), el.Vision(), ind)
	closest := make([]alive.Alive, len(idInt))
	j := 0
	for i := 0; i < len(idInt); i++ {
		id := idInt[i]
		closest[j] = fr.Get(id)
		j++
	}
	return idInt, closest
}

func (w *World) fixLimit(el animal.Animal) {
	x, y := el.X(), el.Y()
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
	el.SetCrd(crd.NewCrd(x, y))
}

func NewWorldTest(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     frame.NewFrame(countAnimal, w, h),
		plant:      frame.NewFrame(countPlant, w, h),
	}
	crAnimal := func(i int, x, y float64) {
		//el := species.NewBeast(behavior.NewAiv1(w, h))
		el := species.NewBeast(behavior.NewTestAngel(math.Pi / 2 * -1))
		//el := species.NewBeast(behavior.NewSimple(w, h))
		//gnt.Generate(el, gnt.WorldWH(w, h), gnt.SetGroup("a"+strconv.Itoa(i)), gnt.SetSize(6))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(41), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridAnimal.Set(el.X(), el.Y(), el.Size(), i)
		world.animal.Set(0, el)
	}
	crPlant := func(i int, x, y float64) {
		el := sp.NewPlant()
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridPlant.Set(el.X(), el.Y(), el.Size(), i)
		world.plant.Set(i, el)
	}
	crAnimal(0, 200, 200)
	crPlant(0, 30, 50)
	//crPlant(1, 70, 50)
	return world
}
