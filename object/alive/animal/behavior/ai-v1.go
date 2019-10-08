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
	eating = 50
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

func tD(speed, distance float64, cycle uint64) uint64{
	return uint64(distance/speed*1.1) + cycle
}

func (a *aiV1) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	dangerous := dangerous(self, animals)
	if len(dangerous.obj) > 0 {
		reason := dangerous.Names()
		if valid, crd := a.mem.checkByReason(running, cycle, reason); valid {
			a.direction.SetCrd(crd.GetX(), crd.GetY())
			return a.direction
		}
		sum := vector.GetVectorByPoint(self.GetX(), self.GetY(), self.GetX(), self.GetY())
		for _, v := range dangerous.obj {
			sum = vector.Add(sum, v.vec)
		}
		sum = checkEdge2(sum, self.GetX(), self.GetY(), a.w, a.h, self.Vision())
		x, y := sum.GetPointFromVector(self.GetX(), self.GetY())
		a.mem.set(running, tD(self.Speed(), self.Vision(), cycle), reason, object.NewCrd(x, y))
		a.direction.SetCrd(x, y)
		return a.direction
	}
	if valid, crd := a.mem.check(eating, cycle); valid {
		a.direction.SetCrd(crd.GetX(), crd.GetY())
		return a.direction
	}
	var closest alive.Alive
	closestFn := func() alive.Alive {
		closest = closestFn(self, animals, plants)
		return closest
	}
	reason := strconv.Itoa(len(animals))+"-"+strconv.Itoa(len(plants))
	if valid, crd := a.mem.checkByReason(eating, cycle, reason); valid {
		a.direction.SetCrd(crd.GetX(), crd.GetY())
		return a.direction
	}
	if (len(animals) == 0 && len(plants) == 0) || closestFn() == nil {
		return a.simple.Direction(self, nil, nil, 0)
	}
	a.mem.set(eating, tD(self.Speed(), self.Vision(), cycle), reason, object.NewCrd(closest.GetX(), closest.GetY()))
	a.direction.SetCrd(closest.GetX(), closest.GetY())
	return a.direction
}

func nameAlive(al []alive.Alive) string {
	names := make([]string, len(al))
	for k, v := range al {
		names[k] = v.Group()
	}
	return strings.Join(names, "")
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
				tmpPoints := geom.NewPoint(interPoint.X() - l1, interPoint.Y())
				checkLeng(tmpPoints, point2)
			}
			if interPoint.X()+l2 < w {
				tmpPoints := geom.NewPoint(interPoint.X() + l2, interPoint.Y())
				checkLeng(tmpPoints, point2)
			}
		}
		if v.side == "y" {
			if interPoint.Y()-l1 > 0 {
				tmpPoints := geom.NewPoint(interPoint.X(), interPoint.Y() - l1)
				checkLeng(tmpPoints, point2)
			}
			if interPoint.Y()+l2 < h {
				tmpPoints := geom.NewPoint(interPoint.X(), interPoint.Y() + l2)
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

func closestFn(self animal.Animal, animals []alive.Alive, plants []alive.Alive) alive.Alive {
	closestFnAn := func() alive.Alive {
		return getClosest(self, animals)
	}
	closestFnPl := func() alive.Alive {
		return getClosest(self, plants)
	}
	if closest := closestFnAn(); closest == nil {
		return closestFnPl()
	} else {
		return closest
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

func getClosest(el animal.Animal, els []alive.Alive) alive.Alive {
	var elRes alive.Alive
	dist := 9e+5
	mass := 0.0
	for i := 0; i < len(els); i++ {
		el1 := els[i]
		var distRes float64
		distFn := func() float64 {
			distRes = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd()) - el.Size()
			return distRes
		}
		if el != nil && el1 != nil && !el1.GetDead() &&
			el.Size()/el1.Size() > _const.EatRatio &&
			mass <= el1.Size() && distFn() < dist && distRes < el.Vision() {
			elRes = el1
			dist = distRes
			mass = el1.Size()
		}
	}
	return elRes
}
