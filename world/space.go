package world

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior/ai"
	"agar-life/object/alive/animal/species"
	"agar-life/object/alive/plant"
	sp "agar-life/object/alive/plant/species"
	"agar-life/world/const"
	"agar-life/world/frame"
	"agar-life/world/frame/grid"
	gnt "agar-life/world/generate"
	"fmt"
	"io"
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
	gridAnimal grid.Grid
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

func NewWorldFromFile(reader io.Reader) World {
	w, h := 1000.0, 1000.0
	countAnimal := 0
	countPlant := 0
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
	return math2.Random(0, 10) > 8
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
		idC, inner, closest := getClosest(w.gridAnimal, el, w.animal, i)
		closestAnimal := append(w.forIntersect(el, closest[:inner], idC[:inner], &w.animal, &removeList), closest[inner:]...)
		idC, inner, closest = getClosest(w.gridPlant, el, w.plant, -1)
		closestPlant := append(w.forIntersect(el, closest[:inner], idC[:inner], &w.plant, &removeList), closest[inner:]...)
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
			direction = w.fixNeighborhood(el, direction)
			el.SetCrdByDirection(el, direction, dist, first)
		}
		w.fixLimit(el)
		grow(el)
		if el.Size() < el.ViewSize() && el.GrowSize() >= 0 {

		}
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

func grow(el animal.Animal) {
	if el.Size() == el.ViewSize() {
		return
	}
	if math.Abs(el.Size() - el.ViewSize()) > math.Abs(el.GrowSize()) * 2    {
		el.SetViewSize(el.ViewSize() + el.GrowSize())
	} else {
		el.SetViewSize(el.Size())
	}
}

func (w *World) fixNeighborhood(el animal.Animal, dir crd.Crd) crd.Crd {
	var parent animal.Animal
	if parent = el.Parent(); parent == nil {
		return dir
	}
	intersected := false
	sum := vector.GetVectorByPoint(el.GetCrd(), el.GetCrd())
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
			intersected = true
			sum = vector.Add(sum, vector.GetVectorWithLength(el1.GetCrd(), el.GetCrd(), el.Size() + el1.Size()))
		}
	}
	if intersected {
		dir = sum.GetPointFromVector(el.GetCrd())
	}
	return dir
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
	for i := 0; i < len(m.list); i++ {
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
	for j := 0; j < len(closest); j++ {
		index := idInt[j]
		el1 := fr.Get(index)
		dis := -1.0
		dist := func() float64 {
			if dis == -1.0 {
				dis = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
			}
			return dis
		}
		died := false
		if el != nil && el1 != nil && !el1.GetDead() {
			if (el.Size()/el1.Size() > _const.EatRatio || (el.Group() == el1.Group() &&
				el1.GlueTime() <= w.cycle && el.GlueTime() <= w.cycle)) && !el1.Danger() && dist() < el.Size() {
				died = true
				el.Eat(el1)
			}
			if !died && el1.Danger() && el1.Size() < el.Size() && dist() < el.Size() {
				if Burst(&w.animal, el, w.cycle) {
					died = true
				}
			}
			if died {
				el1.Die()
				removeList.add(index, fr)
				fr.SetUpdateState(true)
				closest = alive.Remove(closest, j)
				idInt = removeFromInt(idInt, j)
				j--
			}
		}
	}
	return closest
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0    // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func getClosest(gr grid.Grid, el animal.Animal, fr frame.Frame, ind int) ([]int, int, []alive.Alive) {
	idInt, inner := gr.GetObjInRadius(el.X(), el.Y(), el.Vision(), el.Size(), ind)
	closest := make([]alive.Alive, len(idInt))
	j := 0
	for i := 0; i < len(idInt); i++ {
		id := idInt[i]
		closest[j] = fr.Get(id)
		j++
	}
	return idInt, inner, closest
}

func (w *World) fixLimit(el animal.Animal) {
	x, y := el.X(), el.Y()
	if x < 0 {
		x = 1
	}
	if x > w.w {
		x = w.w - 1
	}
	if y < 0 {
		y = 1
	}
	if y > w.h {
		y = w.h - 1
	}
	if (x == 0 && y == 0) || (x == 1 && y == 1) {
		fmt.Println("ff")
	}
	if x != el.X() || y != el.Y() {
		el.SetXY(x, y)
	}
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
	crAnimal := func(i int, x, y, size float64) {
		el := species.NewBeast(ai.NewAiv1(w, h))
		//el := species.NewBeast(behavior.NewTestAngel(math.Pi / 2 * -1))
		//el := species.NewBeast(behavior.NewSimple(w, h))
		//gnt.Generate(el, gnt.WorldWH(w, h), gnt.SetGroup("a"+strconv.Itoa(i)), gnt.SetSize(6))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(size), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridAnimal.Set(el.X(), el.Y(), el.Size(), i)
		world.animal.Set(i, el)
	}
	crPlant := func(i int, x, y float64, poison bool) {
		var el plant.Plant
		if poison {
			el = sp.NewPoison()
		} else {
			el = sp.NewPlant()
		}
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("p"+strconv.Itoa(i)), gnt.Crd(gnt.FixCrd(x, y)))
		world.gridPlant.Set(el.X(), el.Y(), el.Size(), i)
		world.plant.Set(i, el)
	}
	//ratio := 6.0
	//size := 10.0
	//for i :=0; i < countAnimal; i++ {
	//	fmt.Println(size)
	//	size += ratio * float64(i)
	//	crAnimal(i, 50 + float64(i) * float64(i) * 30 + size, 400, size)
	//}
	//crPlant(0, 30, 50, false)
	//crAnimal(0, 110.09, 209.04, 26)
	//crAnimal(0, 20, 20, 12)

	crAnimal(1, 20, 200, 20)
	crAnimal(0, 0, 250, 16)

	//crPlant(1, 140, 220, false)
	//crPlant(3, 170, 200, false)
	//crPlant(4, 140, 180, false)
	//crPlant(1, 70, 50)

	//crAnimal(0, 150, 200, 18)
	crPlant(0, 100, 750, true)
	crPlant(1, 100, 710, false)
	return world
}
