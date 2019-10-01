package behavior

import (
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/world/const"
)

type aiV1 struct {
	simple
}

func NewAiv1(w, h float64) Behavior {
	simple := simple{w: w, h: h, changeDirection: true}
	return &aiV1{
		simple: simple,
	}
}

func (a *aiV1) SetDirection(self animal.Animal, animals []alive.Alive, plants []alive.Alive) {
	oldDirection := a.direction
	dangerous := dangerous(self, animals)
	if len(dangerous.obj) > 0 {
		sum := vector.GetVectorByPoint(self.GetX(), self.GetY(), self.GetX(), self.GetY())
		for _, v := range dangerous.obj {
			sum = vector.Add(sum, v)
		}
		a.direction.SetCrd(sum.GetPointFromVector(self.GetX(), self.GetY()))
		a.setCrdByDirection(self, oldDirection)
		//println("dangerous")
		return
	}
	var closest alive.Alive
	closestFn := func() alive.Alive {
		closest = getClosest(self, append(animals, plants...))
		return closest
	}
	if (len(animals) == 0 && len(plants) == 0) || closestFn() == nil {
		a.simple.SetDirection(self, nil, nil)
		//println("simple", len(animals), len(plants), closest)
		return
	}
	a.direction.SetCrd(closest.GetX(), closest.GetY())
	a.setCrdByDirection(self, oldDirection)
	//println("pursue")
}

type dangerObj struct {
	obj []vector.Vector
}

func (d *dangerObj) add(x1, y1, x2, y2, vision float64) {
	x2, y2 = getXYWithLength(x1, y1, x2, y2, vision)
	x2, y2 = x1+x2, y1+y2
	vec := vector.GetVectorByPoint(x2, y2, x1, y1)
	for _, v := range d.obj {
		if vector.Compare(v, vec) {
			return
		}
	}
	d.obj = append(d.obj, vec)
}

func dangerous(el animal.Animal, animals []alive.Alive) dangerObj {
	danObj := dangerObj{}
	for i := 0; i < len(animals); i++ {
		el1 := animals[i]
		if el != nil && el1 != nil && !el1.GetDead() && el1.GetSize()/el.GetSize() > _const.EatRatio {
			danObj.add(el.GetX(), el.GetY(), el1.GetX(), el1.GetY(), el.GetVision())
		}
	}
	return danObj
}

func getClosest(el animal.Animal, animals []alive.Alive) alive.Alive {
	var elRes alive.Alive
	dist := 9e+5
	mass := 0.0
	for i := 0; i < len(animals); i++ {
		el1 := animals[i]
		if el1.GetSize() < mass {
			break
		}
		var distRes float64
		distFn := func() float64 {
			distRes = geom.GetDistanceByCrd(el.GetCrd(), el1.GetCrd())
			return distRes
		}
		if el != nil && el1 != nil && !el1.GetDead() &&
			el.GetSize()/el1.GetSize() > _const.EatRatio &&
			mass <= el1.GetSize() && distFn() < dist && distRes < el.GetVision() {
			elRes = el1
			dist = distRes
			mass = el1.GetSize()
		}
	}
	return elRes
}
