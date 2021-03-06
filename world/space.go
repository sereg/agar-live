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
	"encoding/json"
	"math"
	"sort"
	"strconv"
	"strings"
)

type World struct {
	w, h        float64
	animal      frame.Frame
	plant       frame.Frame
	cycle       uint
	gridPlant   grid.Grid //TODO move grid index to frame
	gridAnimal  grid.Grid
	resurrect   resurrects
	countPlant  uint
	countAnimal uint
	action      bool
}

func NewWorld(countPlant, countAnimal int, w, h float64) World {
	world := World{
		w:           w,
		h:           h,
		gridPlant:   grid.NewArray(_const.GridSize, w, h),
		gridAnimal:  grid.NewArray(_const.GridSize, w, h),
		animal:      frame.NewFrame(countAnimal, w, h),
		plant:       frame.NewFrame(countPlant, w, h),
		countAnimal: uint(countAnimal),
		countPlant:  uint(countPlant),
	}
	for i := 0; i < countAnimal; i++ {
		el := species.NewBeast(ai.NewAiv1(w, h))
		gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(i)), gnt.Size(_const.AliveStartSize))
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
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
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.Set(i, el)
	}
	return world
}

func poison() bool {
	return math2.Random(0, 10) > 8
}

func (w *World) Cycle() {
	first := true
	if w.cycle > 0 && !w.action {
		first = false
		w.plant.SetUpdateState(false)
	} else {
		w.action = false
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
		dist := el.GetSpeed()
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
		if el.GetSize() < el.GetViewSize() && el.GetGrowSize() >= 0 {

		}
	}
	w.resurrect.resurrect(w.cycle, w.w, w.h)
	w.remove(removeList)
	w.gridAnimal.Reset()
	for i := 0; i < len(w.animal.All()); i++ {
		el := w.animal.Get(i)
		w.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
	}
	if w.plant.UpdateState() {
		w.gridPlant.Reset()
		for i := 0; i < len(w.plant.All()); i++ {
			el := w.plant.Get(i)
			w.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		}
	}
	w.cycle++
	w.countAnimal++
}

func grow(el animal.Animal) {
	if el.GetSize() == el.GetViewSize() {
		return
	}
	if math.Abs(el.GetSize()-el.GetViewSize()) > math.Abs(el.GetGrowSize())*2 {
		el.SetViewSize(el.GetViewSize() + el.GetGrowSize())
	} else {
		el.SetViewSize(el.GetSize())
	}
}

func (w *World) fixNeighborhood(el animal.Animal, dir crd.Crd) crd.Crd {
	var parent animal.Animal
	if parent = el.GetParent(); parent == nil {
		return dir
	}
	intersected := false
	sum := vector.GetVectorByPoint(el.GetCrd(), el.GetCrd())
	for _, v := range parent.GetChildren() {
		var el1 animal.Animal
		if v.GetID() == el.GetID() {
			el1 = parent
		} else {
			el1 = v
		}
		if el.GetGlueTime() <= w.cycle && el1.GetGlueTime() <= w.cycle {
			continue
		}
		dis := geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
		if dis < el.GetSize()+el1.GetSize() {
			intersected = true
			sum = vector.Add(sum, vector.GetVectorWithLength(el1.GetCrd(), el.GetCrd(), el.GetSize()+el1.GetSize()))
		}
	}
	if intersected {
		dir = sum.GetPointFromVector(el.GetCrd())
	}
	return dir
}

func (w World) GetPlant() []alive.Alive {
	var el []alive.Alive
	if !w.plant.UpdateState() {
		return el
	}
	return w.plant.All()
}

func (w World) GetCycle() uint {
	return w.cycle
}

type Alive struct {
	ID   int
	Size float64
	Typ  string
}

func (w World) GetID(x, y float64) string {
	id, typ := w.get(x, y)
	if id == -1 {
		return ""
	}
	var el alive.Alive
	if typ == _const.AnimalTypeAlive {
		el = w.animal.Get(id)
	} else {
		el = w.plant.Get(id)
	}
	exp := Alive{ID: el.GetID(), Size: el.GetSize(), Typ: typ}
	jData, err := json.MarshalIndent(exp, "", "\t")
	if err != nil {
		println(err)
		return ""
	}
	return string(jData)
}

func (w World) get(x, y float64) (int, string) {
	radius := float64(_const.PoisonSize)
	cr := crd.NewCrd(x, y)
	for radius < 200 {
		inGrid := func(fr *frame.Frame, gr grid.Grid) int {
			idInt, _ := gr.GetObjInRadius(x, y, radius, radius, -1)
			for _, index := range idInt {
				el1 := fr.Get(index)
				dist := geom.GetDistanceByCrd(cr, el1.GetCrd())
				if dist < el1.GetSize() {
					return index
				}
			}
			return -1
		}
		if index := inGrid(&w.animal, w.gridAnimal); index != -1 {
			return index, _const.AnimalTypeAlive
		}
		if radius == float64(_const.PoisonSize) {
			if index := inGrid(&w.plant, w.gridPlant); index != -1 {
				return index, _const.PlantTypeAlive
			}
		}
		radius += 50
	}
	return -1, ""
}

type ElType struct {
	Type string
}

func (w *World) GetEl(x, y float64) string {
	id, typ := w.get(x, y)
	if id == -1 {
		return ""
	}
	var el alive.Alive
	if typ == _const.AnimalTypeAlive {
		el = w.animal.Get(id)
	} else {
		el = w.plant.Get(id)
	}
	exp := struct {
		Type string
		El   alive.Alive
	}{typ, el}
	jData, err := json.MarshalIndent(exp, "", "\t")
	if err != nil {
		println(err)
	}
	w.deleteBYIndex(id, typ)
	return string(jData)
}

type animalType struct {
	Type string
	El   Animal
}

type plantType struct {
	Type string
	El   Plant
}

func (w *World) AddFromJSON(data string, x, y float64) string {
	if data == "" {
		return ""
	}
	var typ ElType
	err := json.NewDecoder(strings.NewReader(data)).Decode(&typ)
	if err != nil {
		println(err)
		return ""
	}
	var el1 alive.Alive
	if typ.Type == _const.AnimalTypeAlive {
		var infoEl animalType
		err := json.NewDecoder(strings.NewReader(data)).Decode(&infoEl)
		if err != nil {
			println(err)
			return ""
		}
		an := infoEl.El
		an.X = x
		an.Y = y
		var parent animal.Animal
		if an.Parent != nil {
			if parentID := w.getIndexByID(*an.Parent, _const.AnimalTypeAlive); parentID != -1 {
				parent = w.animal.Get(parentID).(animal.Animal)
			}
		}
		el1 = createAnimalFromJSON(an, w.w, w.w, parent)
		w.animal.Add(el1)
	}
	if typ.Type == _const.PlantTypeAlive {
		var infoEl plantType
		err := json.NewDecoder(strings.NewReader(data)).Decode(&infoEl)
		if err != nil {
			println(err)
			return ""
		}
		pl := infoEl.El
		pl.X = x
		pl.Y = y
		el1 = createPlantFromJSON(pl)
		w.plant.Add(el1)
		w.plant.SetUpdateState(true)
		w.action = true
	}
	exp := struct {
		Type string
		El   alive.Alive
	}{typ.Type, el1}
	jData, err := json.MarshalIndent(exp, "", "\t")
	if err != nil {
		println(err)
	}
	return string(jData)
}

type elInfo struct {
	ID   int
	Size float64
	Type string
}

func (w *World) SetSize(data string) {
	var info elInfo
	err := json.NewDecoder(strings.NewReader(data)).Decode(&info)
	if err != nil {
		println(err)
		return
	}
	if index := w.getIndexByID(info.ID, info.Type); index != -1 {
		var el alive.Alive
		if info.Type == _const.AnimalTypeAlive {
			el = w.animal.Get(index)
		}
		if info.Type == _const.PlantTypeAlive {
			el = w.plant.Get(index)
		}
		if el == nil {
			return
		}
		el.SetSize(info.Size)
	}
}

func (w *World) DeleteByID(id int, typ string) {
	w.deleteBYIndex(w.getIndexByID(id, typ), typ)
}

func (w *World) deleteBYIndex(index int, typ string) {
	if index == -1 {
		return
	}
	if typ == _const.AnimalTypeAlive {
		w.animal.Delete(index)
	} else {
		w.plant.Delete(index)
		w.plant.SetUpdateState(true)
	}
	w.action = true
}

func (w World) getIndexByID(id int, typ string) int {
	getEl := func(fr frame.Frame) int {
		for index, el := range fr.All() {
			if el.GetID() == id {
				return index
			}
		}
		return -1
	}
	if typ == _const.AnimalTypeAlive {
		return getEl(w.animal)
	}
	if typ == _const.PlantTypeAlive {
		return getEl(w.plant)
	}
	return -1
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
			if el.GetParent() == nil && len(el.GetChildren()) == 0 {
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
			if (el.GetSize()/el1.GetSize() > _const.EatRatio || (el.GetGroup() == el1.GetGroup() &&
				el1.GetGlueTime() <= w.cycle && el.GetGlueTime() <= w.cycle)) && !el1.GetDanger() && dist() < el.GetSize() {
				died = true
				el.Eat(el1)
			}
			if !died && el1.GetDanger() && el1.GetSize() < el.GetSize() && dist() < el.GetSize() {
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
	idInt, inner := gr.GetObjInRadius(el.GetX(), el.GetY(), el.GetVision(), el.GetSize(), ind)
	closest := make([]alive.Alive, 0, len(idInt))
	was := map[int]struct{}{}
	j := 0
	for i := 0; i < len(idInt); i++ {
		id := idInt[i]
		if _, ok := was[id]; ok {
			continue
		} else {
			//was[id] = struct{}{}
		}
		closest = append(closest, fr.Get(id))
		j++
	}
	return idInt, inner, closest
}

func (w *World) fixLimit(el animal.Animal) {
	x, y := el.GetX(), el.GetY()
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
	if x != el.GetX() || y != el.GetY() {
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
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), i)
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
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
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
