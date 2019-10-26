package ai

import (
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/world/const"
	"math"
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
	top   = geom.Segment{}
	down  = geom.Segment{}
	left  = geom.Segment{}
	right = geom.Segment{}
)

func NewAiv1(w, h float64) animal.Behavior {
	top = geom.NewSegment(crd.NewCrd(0, 0), crd.NewCrd(w, 0))
	down = geom.NewSegment(crd.NewCrd(0, h), crd.NewCrd(w, h))
	left = geom.NewSegment(crd.NewCrd(0, 0), crd.NewCrd(0, h))
	right = geom.NewSegment(crd.NewCrd(w, 0), crd.NewCrd(w, h))
	return &aiV1{
		Simple: behavior.NewSimple(w, h),
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
	action    func() crd.Crd
}

func (a *aiV1) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	dangerous := dangerous(self, animals)
	poisons := poisons(self, plants)
	poisonCount := len(poisons)
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
		{ //running
			priority: running,
			mem:      true,
			condition: func() bool {
				return len(dangerous.obj) > 0
			},
			reason: func() string {
				return strconv.Itoa(len(dangerous.obj)) + "-" + strconv.Itoa(poisonCount)
			},
			action: func() crd.Crd {
				sum := vector.GetVectorByPoint(self.GetCrd(), self.GetCrd())
				for _, v := range dangerous.obj {
					sum = vector.Add(sum, v.vec)
				}
				sum = checkEdge2(sum, self.GetCrd(), a.W(), a.H(), self.Vision())
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
				return closestFn() != nil
			},
			reason: func() string {
				return strconv.Itoa(len(animals)) + "-" + strconv.Itoa(len(dangerous.obj)) + "-" + strconv.Itoa(poisonCount)
			},
			action: func() crd.Crd {
				//TODO dont send objects in poisonous plants
				return closest.GetCrd()
			},
		},
		{ //default
			priority: nothing,
			mem:      false,
			condition: func() bool {
				return true
			},
			action: func() crd.Crd {
				cr, _ := a.Simple.Action(self, nil, nil, 0)
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
			cr = bypass(self, cr, poisons)
			if strategy.mem {
				a.mem.set(running, tD(self.Speed(), self.Vision(), cycle), reason, cr)
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

type rangeAngels struct{
	dangerous bool
	start, finish float64
	dist float64
}

func (a *aiV1) checkAngels(el animal.Animal, poisons []alive.Alive) []rangeAngels {
	if len(poisons) == 0 {
		return nil
	}
	locPoisons := make([]intersect, len(poisons))
	for k, v := range poisons {
		locPoisons[k] = newPoint(v)
	}
	count := int((el.Vision() * math.Pi * 2) / el.Size())
	for count % 4 != 0 {
		count++
	}
	addAngel := 2.0 * math.Pi / float64(count)
	angel := 0.0
	sift := float64(count / 4.0)
	angelV := angel + addAngel*sift - addAngel
	angelD := angelV * 2
	xd := el.X() + el.Vision()*math.Cos(angelV)
	yd := el.Y() + el.Vision()*math.Sin(angelV)
	vec := vector.GetVectorByPoint(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(xd, yd))
	dirAngel := geom.ModuleDegree(vec.GetAngle())
	rangAng := make([]rangeAngels, 0, len(poisons))
	for i := 0.0; i < float64(count); i++ {
		xs1 := el.X() + el.Size()*math.Cos(angel)
		ys1 := el.Y() + el.Size()*math.Sin(angel)
		xs2 := el.X() + el.Size()*math.Cos(angel)
		ys2 := el.Y() + el.Size()*math.Sin(angel)
		angelV = angel + addAngel*sift
		xf1 := el.X() + el.Vision()*math.Cos(angelV)
		yf1 := el.Y() + el.Vision()*math.Sin(angelV)
		angelV += angelD
		xf2 := el.X() + el.Vision()*math.Cos(angelV)
		yf2 := el.Y() + el.Vision()*math.Sin(angelV)
		line1 := geom.NewSegment(crd.NewCrd(xs1, ys1), crd.NewCrd(xf1, yf1))
		line2 := geom.NewSegment(crd.NewCrd(xf1, yf1), crd.NewCrd(xf2, yf2))
		line3 := geom.NewSegment(crd.NewCrd(xs2, ys2), crd.NewCrd(xf2, yf2))
		for _, v := range locPoisons {
			if intersect, dist := v.check(el.GetCrd(), line1, line2, line3); intersect {
				rangAng = append(rangAng, rangeAngels{
					dangerous: v.dangerous(),
					start:     dirAngel + addAngel,
					finish:    dirAngel - addAngel,
					dist:      dist,
				})
				break
			}
		}
		dirAngel -=addAngel
		angel += addAngel
	}
	return rangAng
}

type intersect interface{
	check(center crd.Crd, lines ...geom.Segment) (bool, float64)
	dangerous() bool
}

type point struct {
	outer []geom.Segment
	inner []geom.Segment
}

func newPoint(el alive.Alive) intersect {
	p := point{
		outer: []geom.Segment{
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X()-el.Size(), el.Y())),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X(), el.Y()-el.Size())),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X()+el.Size(), el.Y())),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X(), el.Y()+el.Size())),
		},
		inner: []geom.Segment{
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(0, el.Y())),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X(), 0)),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(9e+2, el.Y())),
			geom.NewSegment(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(el.X(), 9e+2)),
		},
	}
	return &p
}

func (p *point) dangerous() bool {
	return true
}

func (p *point) check(center crd.Crd, lines ...geom.Segment) (bool, float64){
	for _, v := range p.outer {
		for _, line := range lines {
			if v.Intersection(line) {
				return true, geom.GetDistanceByCrd(center, v.Start())
			}
		}
	}
	countIntersect := 0
	for _, v := range p.outer {
		for _, line := range lines {
			if v.Intersection(line) {
				countIntersect++
				break
			}
		}
	}
	if countIntersect >= 3 {
		return true, geom.GetDistanceByCrd(center, p.outer[0].Start())
	}
	return false, 0
}

//func (a *aiV1) segmentEdge(el animal.Animal) (seg []geom.Segment) {
//	if el.X() - el.Vision() < 0 {
//		seg = append(seg, left)
//	}
//	if el.X() + el.Vision() > a.W() {
//		seg = append(seg, right)
//	}
//	if el.Y() - el.Vision() < 0 {
//		seg = append(seg, top)
//	}
//	if el.Y() + el.Vision() > a.H() {
//		seg = append(seg, down)
//	}
//	return
//}

func bypass(el animal.Animal, direction crd.Crd, poisons []alive.Alive) crd.Crd {
	if len(poisons) == 0 {
		return direction
	}
	vec := vector.GetVectorWithLength(el.GetCrd(), direction, el.Size())
	vec.AddAngle(math.Pi / 2)
	p1 := vec.GetPointFromVector(el.GetCrd())
	vec = vector.GetVectorWithLength(p1, el.GetCrd(), el.Size()*2)
	p2 := vec.GetPointFromVector(p1)
	vec = vector.GetVectorWithLength(direction, el.GetCrd(), el.Size())
	vec.AddAngle(math.Pi / 2)
	p3 := vec.GetPointFromVector(direction)
	vec = vector.GetVectorWithLength(p3, direction, el.Size()*2)
	p4 := vec.GetPointFromVector(p3)
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
			geom.NewSegment(crd.NewCrd(0, v.Y()), crd.NewCrd(v.X(), v.Y())),
			geom.NewSegment(crd.NewCrd(v.X(), 0), crd.NewCrd(v.X(), v.Y())),
			geom.NewSegment(crd.NewCrd(9e+4, v.Y()), crd.NewCrd(v.X(), v.Y())),
			geom.NewSegment(crd.NewCrd(v.X(), 9e+4), crd.NewCrd(v.X(), v.Y())),
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
		vec = vector.GetVectorWithLength(closestPoison.GetCrd(), el.GetCrd(), closestPoison.Size())
		vec.AddAngle(-1 * math.Pi / 2)
		p1 = vec.GetPointFromVector(closestPoison.GetCrd())
		vec = vector.GetVectorWithLength(p1, p2, el.Size()*1.8)
		vec.AddAngle(-1 * math.Pi / 2)
		p1 = vec.GetPointFromVector(p1)
		dis := geom.GetDistanceByCrd(el.GetCrd(), p1)
		dis += el.Size() * 3
		vec = vector.GetVectorWithLength(el.GetCrd(), p1, dis)
		direction = vec.GetPointFromVector(el.GetCrd())
		//println("set new direction")
	}
	return direction
}

func closestFn(self animal.Animal, animals []alive.Alive, plants []alive.Alive) (closest alive.Alive, split bool) {
	closestFnAn := func() (closest alive.Alive, split bool) {
		return getClosest(self, animals, true)
	}
	closestFnPl := func() (closest alive.Alive, split bool) {
		return getClosest(self, plants, false)
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

func poisons(el animal.Animal, plants []alive.Alive) (poisons []alive.Alive) {
	if el.Size()-_const.MinSizeAlive < _const.MinSizeAlive || el.Count() >= _const.SplitMaxCount {
		return poisons
	}
	poisons = make([]alive.Alive, 0, len(plants)/10)
	for i := 0; i < len(plants); i++ {
		el1 := plants[i]
		if el == nil || el1 == nil || el1.GetDead() {
			continue
		}
		if el1.Size() < el.Size() && el1.Danger() {
			poisons = append(poisons, el1)
		}
	}
	return
}

func getClosest(el animal.Animal, els []alive.Alive, animal bool) (closest alive.Alive, split bool) {
	dist := 9e+5
	mass := 0.0
	for i := 0; i < len(els); i++ {
		el1 := els[i]
		distRes := -1.0
		distFn := func() float64 {
			if distRes == -1.0 {
				distRes = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd()) - el.Size()
			}
			return distRes
		}
		if el != nil && el1 != nil && !el1.Danger() &&
			el.Size()/el1.Size() > _const.EatRatio &&
			(mass <= el1.Size() || mass > _const.FoodSize) && //TODO add equation choice distance or size
			distFn() < dist && distRes < el.Vision() && el1.Group() != el.Group() && !el1.GetDead() {
			closest = el1
			dist = distRes
			if dist < _const.SplitDist && el.Size() > el1.Size()*2.5 {
				if animal { //TODO check object in dangerous angles
					split = true
				}
			}
			mass = el1.Size()
		} else {
			if mass > 0 && !animal && distFn() > _const.GridSize {
				return
			}
		}
	}
	return
}

func checkEdge2(sum vector.Vector, pos crd.Crd, w, h, vision float64) vector.Vector {
	ratio := 0.0
	intersectCount := 0
	type inter struct {
		point1, point2 geom.Point
		side           string
	}
	inters := []inter{}
	position := sum.GetPointFromVector(pos)
	xn, yn := position.X(), position.Y()
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
	point1 := geom.NewPoint(pos.X(), pos.Y())
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
	return vector.GetVectorByPoint(pos, crd.NewCrd(math.Abs(resPoints.X()), math.Abs(resPoints.Y())))
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
