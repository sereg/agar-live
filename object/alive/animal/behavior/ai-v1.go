package behavior

import (
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/world/const"
	"math"
	"strconv"
	"strings"
)

type aiV1 struct {
	simple
	mem memory
}

const (
	running = 100
	eating  = 50
	nothing = 9
)

type memory struct {
	valid     bool
	priority  uint8
	validTime uint64
	reason    string
	crd       object.Crd
}

func (m *memory) set(pr uint8, vt uint64, reason string, crd object.Crd) {
	m.valid = true
	m.priority = pr
	m.validTime = vt
	m.reason = reason
	m.crd = crd
}

func (m *memory) check(pr uint8, cycle uint64) (bool, object.Crd) {
	if m.valid && m.validTime < cycle && m.priority >= pr {
		return true, m.crd
	}
	m.reset()
	return false, m.crd
}

func (m *memory) checkByReason(pr uint8, cycle uint64, reason string) (bool, object.Crd) {
	if m.valid && m.validTime > cycle && m.priority >= pr && m.reason == reason {
		return true, m.crd
	}
	m.reset()
	return false, m.crd
}

func (m *memory) reset() {
	m.valid = false
}

func NewAiv1(w, h float64) animal.Behavior {
	simple := simple{w: w, h: h, changeDirection: true}
	return &aiV1{
		simple: simple,
	}
}

func tD(speed, distance float64, cycle uint64) uint64 {
	return uint64(distance/speed*1.1) + cycle
}

type strategy struct {
	priority  uint8
	mem       bool
	condition func() bool
	reason    func() string
	action    func() object.Crd
}

func (a *aiV1) GetDirection() object.Crd {
	return a.direction
}

func (a *aiV1) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (object.Crd, bool) {
	dangerous := dangerous(self, animals)
	plants, poisons := poisons(self, plants)
	var closest alive.Alive
	split := false
	closestFn := func() alive.Alive {
		if len(animals) == 0 && len(plants) == 0 {
			return nil
		}
		closest, split = closestFn(self, animals, plants)
		return closest
	}
	strategies := []strategy{
		strategy{ //running
			priority: running,
			mem:      true,
			condition: func() bool {
				return len(dangerous.obj) > 0
			},
			reason: func() string {
				return strconv.Itoa(len(dangerous.obj)) + "-" + strconv.Itoa(len(poisons))
			},
			action: func() object.Crd {
				sum := vector.GetVectorByPoint(self.GetX(), self.GetY(), self.GetX(), self.GetY())
				for _, v := range dangerous.obj {
					sum = vector.Add(sum, v.vec)
				}
				sum = checkEdge2(sum, self.GetX(), self.GetY(), a.w, a.h, self.Vision())
				x, y := sum.GetPointFromVector(self.GetX(), self.GetY())
				//TODO hide in poison plant if size of them more then object
				return object.NewCrd(x, y)
			},
		},
		strategy{ //running from memory
			priority: running,
			mem:      false,
			condition: func() bool {
				valid, _ := a.mem.check(running, cycle)
				return valid
			},
			action: func() object.Crd {
				_, crd := a.mem.check(running, cycle)
				return crd
			},
		},
		strategy{ //eating
			priority: eating,
			mem:      true,
			condition: func() bool {
				return closestFn() != nil
			},
			reason: func() string {
				return strconv.Itoa(len(animals)) + "-" + strconv.Itoa(len(dangerous.obj)) + "-" + strconv.Itoa(len(poisons))
			},
			action: func() object.Crd {
				//TODO dont send objects in poisonous plants
				return object.NewCrd(closest.GetX(), closest.GetY())
			},
		},
		strategy{ //default
			priority: nothing,
			mem:      false,
			condition: func() bool {
				return true
			},
			action: func() object.Crd {
				crd, _ := a.simple.Direction(self, nil, nil, 0)
				return crd
			},
		},
	}
	for _, strategy := range strategies {
		if strategy.condition() {
			reason := ""
			if strategy.mem {
				reason = strategy.reason()
				if valid, crd := a.mem.checkByReason(strategy.priority, cycle, reason); valid {
					a.direction.SetCrd(crd.GetX(), crd.GetY())
					break
				}
			}
			crd := strategy.action()
			crd = bypass(self, crd, poisons)
			if strategy.mem {
				a.mem.set(running, tD(self.Speed(), self.Vision(), cycle), reason, crd)
			}
			a.direction.SetCrd(crd.GetX(), crd.GetY())
			break
		}
	}
	return a.direction, split
}

func bypass(el animal.Animal, direction object.Crd, poisons []alive.Alive) object.Crd {
	if len(poisons) == 0 {
		return direction
	}
	vec := getVectorWithLength(el.GetX(), el.GetY(), direction.GetX(), direction.GetY(), el.Size())
	vec.AddAngle(math.Pi / 2)
	p1 := object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY()))
	vec = getVectorWithLength(p1.GetX(), p1.GetY(), el.GetX(), el.GetY(), el.Size()*2)
	p2 := object.NewCrd(vec.GetPointFromVector(p1.GetX(), p1.GetY()))
	vec = getVectorWithLength(direction.GetX(), direction.GetY(), el.GetX(), el.GetY(), el.Size())
	vec.AddAngle(math.Pi / 2)
	p3 := object.NewCrd(vec.GetPointFromVector(direction.GetX(), direction.GetY()))
	vec = getVectorWithLength(p3.GetX(), p3.GetY(), direction.GetX(), direction.GetY(), el.Size()*2)
	p4 := object.NewCrd(vec.GetPointFromVector(p3.GetX(), p3.GetY()))
	dist := 9e+5
	var closestPoison alive.Alive
	intersect := false
	linesR := []geom.Segment{
		geom.NewSegment(p1, p2),
		geom.NewSegment(p2, p3),
		geom.NewSegment(p3, p4),
		geom.NewSegment(p4, p1),
	}
	for _, v := range poisons {
		dis := geom.GetDistanceByCrd(el.GetCrd(), v.GetCrd())
		if dist < dis {
			continue
		}
		linesEl := []geom.Segment{
			geom.NewSegment(object.NewCrd(0, v.GetY()), object.NewCrd(v.GetX(), v.GetY())),
			geom.NewSegment(object.NewCrd(v.GetX(), 0), object.NewCrd(v.GetX(), v.GetY())),
			geom.NewSegment(object.NewCrd(9e+4, v.GetY()), object.NewCrd(v.GetX(), v.GetY())),
			geom.NewSegment(object.NewCrd(v.GetX(), 9e+4), object.NewCrd(v.GetX(), v.GetY())),
		}
		countIntersect := 0
		for _, l := range linesEl {
			inter := false
			for _, r := range linesR {
				if l.Intersection(r) {
					countIntersect++
					inter = true
					break
				}
			}
			if !inter {
				break
			}
		}
		if countIntersect >= 4 {
			dist = dis
			intersect = true
			closestPoison = v
		}
	}
	if intersect && closestPoison != nil {
		vec = getVectorWithLength(closestPoison.GetX(), closestPoison.GetY(), el.GetX(), el.GetY(), closestPoison.Size())
		vec.AddAngle(-1 * math.Pi / 2)
		p1 = object.NewCrd(vec.GetPointFromVector(closestPoison.GetX(), closestPoison.GetY()))
		vec = getVectorWithLength(p1.GetX(), p1.GetY(), p2.GetX(), p2.GetY(), el.Size() * 1.8)
		vec.AddAngle(-1 * math.Pi / 2)
		p1 = object.NewCrd(vec.GetPointFromVector(p1.GetX(), p1.GetY()))
		dis :=  geom.GetDistanceByCrd(el.GetCrd(), p1)
		dis += el.Size() * 3
		vec = getVectorWithLength(el.GetX(), el.GetY(), p1.GetX(), p1.GetY(), dis)
		direction = object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY()))
		//println("set new direction")
	}
	return direction
}

func getVectorWithLength(x1, y1, x2, y2, dist float64) vector.Vector {
	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
	length := vec.Len()
	ratio := dist / length
	vec.MultiplyByScalar(ratio)
	return vec
}

func closestFn(self animal.Animal, animals []alive.Alive, plants []alive.Alive) (closest alive.Alive, split bool) {
	closestFnAn := func() (closest alive.Alive, split bool) {
		return getClosest(self, animals)
	}
	closestFnPl := func() (closest alive.Alive, split bool) {
		return getClosest(self, plants)
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

func (d *dangerObj) add(x1, y1, x2, y2, vision float64, name string) {
	x2, y2 = getXYWithLength(x1, y1, x2, y2, vision)
	x2, y2 = x1+x2, y1+y2
	vec := vector.GetVectorByPoint(x2, y2, x1, y1)
	for _, v := range d.obj {
		if vector.Compare(v.vec, vec) {
			return
		}
	}
	d.obj = append(d.obj, dp{vec: vec, name: name})
}

func getXYWithLength(x1, y1, x2, y2, dist float64) (x float64, y float64) {
	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
	length := vec.Len()
	ratio := dist / length
	vec.MultiplyByScalar(ratio)
	x, y = vec.GetPointFromVector(x2, y2)
	x, y = x-x2, y-y2
	return
}

func dangerous(el animal.Animal, animals []alive.Alive) dangerObj {
	danObj := dangerObj{}
	for i := 0; i < len(animals); i++ {
		el1 := animals[i]
		if el != nil && el1 != nil && !el1.GetDead() && el1.Size()/el.Size() > _const.EatRatio {
			danObj.add(el.GetX(), el.GetY(), el1.GetX(), el1.GetY(), el.Vision(), el1.Group())
		}
	}
	return danObj
}

func poisons(el animal.Animal, plants []alive.Alive) (food []alive.Alive, poisons []alive.Alive) {
	if el.Size()-_const.MinSizeAlive < _const.MinSizeAlive || el.Count() >= _const.SplitMaxCount {
		return plants, poisons
	}
	food = make([]alive.Alive, 0, len(plants))
	poisons = make([]alive.Alive, 0, len(plants))
	for i := 0; i < len(plants); i++ {
		el1 := plants[i]
		if el == nil || el1 == nil || el1.GetDead() {
			continue
		}
		if el1.Size() < el.Size() && el1.Danger() {
			poisons = append(poisons, el1)
		} else {
			food = append(food, el1)
		}
	}
	return
}

func getClosest(el animal.Animal, els []alive.Alive) (closest alive.Alive, split bool) {
	dist := 9e+5
	mass := 0.0
	for i := 0; i < len(els); i++ {
		el1 := els[i]
		var distRes float64
		distFn := func() float64 {
			distRes = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd()) - el.Size()
			return distRes
		}
		if el != nil && el1 != nil && !el1.GetDead() && !el1.Danger() &&
			el.Size()/el1.Size() > _const.EatRatio &&
			mass <= el1.Size() && distFn() < dist && distRes < el.Vision() {
			closest = el1
			dist = distRes
			if dist < _const.SplitDist && el.Size() > el1.Size() * 2 {
				if _, ok := el1.(animal.Animal); ok {
					split = true
				}
			}
			mass = el1.Size()
		}
	}
	return
}

func checkEdge2(sum vector.Vector, xobj, yobj, w, h, vision float64) vector.Vector {
	ratio := 0.0
	intersectCount := 0
	type inter struct {
		point1, point2 geom.Point
		side           string
	}
	inters := []inter{}
	xn, yn := sum.GetPointFromVector(xobj, yobj)
	repeat := 1
	increase := func(i inter) {
		if ratio == 0 {
			repeat++
		}
		intersectCount++
		ratio = 0.3 * vision
		inters = append(inters, i)
	}
	for repeat > 0 {
		if xn+ratio > w {
			increase(inter{geom.NewPoint(w, 0), geom.NewPoint(w, h), "y"})
		}
		if xn-ratio < 0 {
			increase(inter{geom.NewPoint(0, 0), geom.NewPoint(0, h), "y"})
		}
		if yn+ratio > h {
			increase(inter{geom.NewPoint(0, h), geom.NewPoint(w, h), "x"})
		}
		if yn-ratio < 0 {
			increase(inter{geom.NewPoint(0, 0), geom.NewPoint(w, 0), "x"})
		}
		repeat--
	}
	if intersectCount == 0 {
		return sum
	}
	point1 := geom.NewPoint(xobj, yobj)
	point2 := geom.NewPoint(xn, yn)
	resPoints := geom.Point{}
	diff := 9e+6
	checkLeng := func(point1, point2 geom.Point) {
		len := geom.LengthLine(point1, point2)
		if len < diff {
			resPoints = point1
			diff = len
		}
	}
	for _, v := range inters {
		interPoint, l1, l2 := getNewPoint(point1, point2, v.point1, v.point2)
		if v.side == "x" {
			if interPoint.X()-l1 > 0 {
				tmpPoints := geom.NewPoint(interPoint.X()-l1, interPoint.Y())
				checkLeng(tmpPoints, point2)
			}
			if interPoint.X()+l2 < w {
				tmpPoints := geom.NewPoint(interPoint.X()+l2, interPoint.Y())
				checkLeng(tmpPoints, point2)
			}
		}
		if v.side == "y" {
			if interPoint.Y()-l1 > 0 {
				tmpPoints := geom.NewPoint(interPoint.X(), interPoint.Y()-l1)
				checkLeng(tmpPoints, point2)
			}
			if interPoint.Y()+l2 < h {
				tmpPoints := geom.NewPoint(interPoint.X(), interPoint.Y()+l2)
				checkLeng(tmpPoints, point2)
			}
		}
	}
	return vector.GetVectorByPoint(xobj, yobj, math.Abs(resPoints.X()), math.Abs(resPoints.Y()))
}

func getNewPoint(point1, point2, point3, point4 geom.Point) (geom.Point, float64, float64) {
	l1 := geom.NewLine(point1, point2)
	l2 := geom.NewLine(point3, point4)
	var result geom.Point
	var err error
	if result, err = l1.Intersection(l2); err == nil {
		// fmt.Println(result)
	}
	len := geom.LengthLine(point1, point2)
	lenb := geom.LengthLine(point1, result)
	angle := math.Abs(geom.AngleAxisX(point1, point2))
	len1 := getSide(len, lenb, angle)
	len2 := getSide(len, lenb, math.Pi-angle)
	return result, len1, len2
}

func getSide(len, lenb, angle float64) float64 {
	angleB := lenb * math.Sin(angle) / len
	angleC := math.Pi - (angle + angleB)

	lenC1 := lenb * lenb
	lenC2 := len * len
	lenC3 := 2 * len * lenb * math.Cos(angleC)
	return math.Sqrt(lenC1 + lenC2 - lenC3)
}
