package ai

import (
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/behavior/checkangels"
	"agar-life/world/const"
	"strconv"
	"strings"
)

type aiV1 struct {
	behavior.Simple
	mem memory
}

const (
	running = 100
	eating  = 50
	nothing = 9
)

var (
	top   checkangels.Obstacle
	down  checkangels.Obstacle
	left  checkangels.Obstacle
	right checkangels.Obstacle
)

func NewAiv1(w, h float64) animal.Behavior {
	top = checkangels.NewLine(geom.NewSegment(crd.NewCrd(0, 0), crd.NewCrd(w, 0)))
	down = checkangels.NewLine(geom.NewSegment(crd.NewCrd(0, h), crd.NewCrd(w, h)))
	left = checkangels.NewLine(geom.NewSegment(crd.NewCrd(0, 0), crd.NewCrd(0, h)))
	right = checkangels.NewLine(geom.NewSegment(crd.NewCrd(w, 0), crd.NewCrd(w, h)))
	return &aiV1{
		Simple: behavior.NewSimple(w, h),
	}
}

func tD(speed, distance float64, cycle uint64) uint64 {
	return uint64(distance/speed*0.3) + cycle
}

type strategy struct {
	priority  uint8
	mem       bool
	condition func() bool
	reason    func() string
	action    func() crd.Crd
}

func (a *aiV1) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	dangerous := dangerous(self, animals)
	poisons := poisons(self, plants)
	edges := a.edgeObstacle(self)
	poisonCount := strconv.Itoa(len(poisons))
	edgesCount := strconv.Itoa(len(edges))
	dangerousCount := strconv.Itoa(len(dangerous.obj))
	var dAngeles checkangels.Angels
	dAngelesFn := func() checkangels.Angels {
		if dAngeles.Angel() == 0 {
			dAngeles = checkangels.CheckAngels(self, append(edges, poisons...))
			//dAngeles = checkangels.CheckAngels(self, poisons)
		}
		return dAngeles
	}
	var closest *crd.Crd
	split := false
	closestFn := func() *crd.Crd {
		if len(animals) == 0 && len(plants) == 0 {
			return nil
		}
		closest, split = closestFn(self, animals, plants, dAngeles)
		return closest
	}
	strategies := []strategy{
		{ //running
			priority: running,
			mem:      true,
			condition: func() bool {
				return len(dangerous.obj) > 0
			},
			reason: func() string {
				return dangerousCount + "-" + poisonCount + "-" + edgesCount
			},
			action: func() crd.Crd {
				dAngelesFn()
				sum := vector.GetVectorByPoint(self.GetCrd(), self.GetCrd())
				for _, v := range dangerous.obj {
					sum = vector.Add(sum, v.vec)
				}
				vecAngel := sum.GetAngle()
				reachable, _ := dAngeles.Check(sum.GetAngle(), sum.Len())
				if !reachable {
					sum.SetAngle(dAngeles.ClosestAvailable(vecAngel))
				}
				//sum = checkEdge2(sum, self.GetCrd(), a.W(), a.H(), self.Vision())
				//TODO hide in poison plant if size of them more then object
				return sum.GetPointFromVector(self.GetCrd())
			},
		},
		{ //running from memory
			priority: running,
			mem:      false,
			condition: func() bool {
				valid, _ := a.mem.check(running, cycle)
				return valid
			},
			action: func() crd.Crd {
				_, crdRes := a.mem.check(running, cycle)
				return crdRes
			},
		},
		{ //eating
			priority: eating,
			mem:      true,
			condition: func() bool {
				dAngelesFn()
				return closestFn() != nil
			},
			reason: func() string {
				return strconv.Itoa(len(animals)) + "-" + strconv.Itoa(len(plants)) + "-" + poisonCount + "-" + edgesCount
			},
			action: func() crd.Crd {
				return closest.GetCrd()
			},
		},
		{ //default
			priority: nothing,
			mem:      true,
			condition: func() bool {
				return true
			},
			reason: func() string {
				return strconv.Itoa(len(animals)) + "-" + strconv.Itoa(len(plants)) + "-" + poisonCount + "-" + edgesCount
			},
			action: func() crd.Crd {
				cr, _ := a.Simple.Action(self, nil, nil, 0, dAngeles)
				return cr
			},
		},
	}
	for _, strategy := range strategies {
		if strategy.condition() {
			reason := ""
			if strategy.mem {
				reason = strategy.reason()
				if valid, cr := a.mem.checkByReason(strategy.priority, cycle, reason); valid {
					a.SetDir(cr)
					break
				}
			}
			cr := strategy.action()
			//if cr == crd.NewCrd(240, 420) {
			//	fmt.Println(self.GetCrd())
			//}
			if strategy.mem {
				a.mem.set(strategy.priority, tD(self.Speed(), self.Vision(), cycle), reason, cr)
			}
			a.SetDir(cr)
			break
		}
	}
	//if self.Size() < 80 {
	//	self.SetSize(80)
	//}
	return a.Dir(), split
}

func (a *aiV1) edgeObstacle(el animal.Animal) (obstacles []checkangels.Obstacle) {
	if el.X()-el.Vision() < 0 {
		obstacles = append(obstacles, left)
	}
	if el.X()+el.Vision() > a.W() {
		obstacles = append(obstacles, right)
	}
	if el.Y()-el.Vision() < 0 {
		obstacles = append(obstacles, top)
	}
	if el.Y()+el.Vision() > a.H() {
		obstacles = append(obstacles, down)
	}
	return
}

func closestFn(self animal.Animal, animals, plants []alive.Alive, dAngeles checkangels.Angels) (closest *crd.Crd, split bool) {
	closestFnAn := func() (closest *crd.Crd, split bool) {
		return getClosest(self, animals, true, dAngeles)
	}
	closestFnPl := func() (closest *crd.Crd, split bool) {
		return getClosest(self, plants, false, dAngeles)
	}
	if closest, split := closestFnAn(); closest == nil {
		return closestFnPl()
	} else {
		return closest, split
	}
}

type dp struct {
	vec  vector.Vector
	name string
}

type dangerObj struct {
	obj []dp
}

func (d dangerObj) Names() string {
	names := make([]string, len(d.obj))
	for k, v := range d.obj {
		names[k] = v.name
	}
	return strings.Join(names, "")
}

func (d *dangerObj) add(a, b crd.Crd, vision float64, name string) {
	vec := vector.GetVectorWithLength(b, a, vision)
	for _, v := range d.obj {
		if vector.Compare(v.vec, vec) {
			return
		}
	}
	d.obj = append(d.obj, dp{vec: vec, name: name})
}

func dangerous(el animal.Animal, animals []alive.Alive) dangerObj {
	danObj := dangerObj{}
	for i := 0; i < len(animals); i++ {
		el1 := animals[i]
		if el != nil && el1 != nil && el1.Size()/el.Size() > _const.EatRatio && el1.Group() != el.Group() && !el1.GetDead() {
			danObj.add(el.GetCrd(), el1.GetCrd(), el.Vision(), el1.Group())
		}
	}
	return danObj
}

func poisons(el animal.Animal, plants []alive.Alive) (poisons []checkangels.Obstacle) {
	if el.Size()-_const.MinSizeAlive < _const.MinSizeAlive || el.Count() >= _const.SplitMaxCount {
		return poisons
	}
	poisons = make([]checkangels.Obstacle, 0, len(plants)/10)
	for i := 0; i < len(plants); i++ {
		el1 := plants[i]
		if el == nil || el1 == nil || el1.GetDead() {
			continue
		}
		if el1.Size() < el.Size() && el1.Danger() {
			poisons = append(poisons, checkangels.NewPoint(el1))
		}
	}
	return
}

func getClosest(el animal.Animal, els []alive.Alive, animal bool, dAngeles checkangels.Angels) (closest *crd.Crd, split bool) {
	dist := 9e+5
	mass := 0.0
	obstacle := true
	for i := 0; i < len(els); i++ {
		el1 := els[i]
		distRes := -1.0
		var vec vector.Vector
		distFn := func() float64 {
			if distRes == -1.0 {
				vec = vector.GetVectorByPoint(el.GetCrd(), el1.GetCrd())
				distRes = vec.Len() - el.Size()
			}
			return distRes
		}
		if el != nil && el1 != nil && !el1.Danger() &&
			el.Size()/el1.Size() > _const.EatRatio &&
			(mass < el1.Size() || obstacle) && //TODO add equation choice distance or size
			el1.Group() != el.Group() && !el1.GetDead() && distFn() < dist &&
			distFn() < el.Vision() {
			vecAngel := geom.ModuleDegree(vec.GetAngle())
			reachable, dangerous := dAngeles.Check(vecAngel, distRes)
			if !reachable && !obstacle {
				continue
			}
			cr := el1.GetCrd()
			if reachable {
				obstacle = false
			} else if dangerous {
				angel := dAngeles.ClosestAvailable(vecAngel)
				vec.SetAngle(angel)
				cr = vec.GetPointFromVector(el.GetCrd())
			}
			closest = &cr
			dist = distRes
			if dist < _const.SplitDist && (el.Size()*_const.SplitRatio)/el1.Size() > _const.EatRatio && !dangerous && animal {
				split = true
			}
			mass = el1.Size()
		} else {
			if !obstacle && i > 10 && mass > 0 && !animal && distFn() > _const.GridSize {
				return
			}
		}
	}
	return
}