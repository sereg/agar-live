package checkangels

import (
	m2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/geom"
	"agar-life/math/vector"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"math"
)

type rangeAngels struct {
	dangerous bool
	dist      float64
}

type keyRange struct {
	start, finish int
}

type mapRange map[keyRange]rangeAngels

type Angels struct {
	rangeAngels mapRange
	angel       int
	first       int
}

func (a Angels) Angel() int {
	return a.angel
}

func (a *Angels) Check(angel, dist float64) (reachable, dangerous bool) {
	if len(a.rangeAngels) == 0 {
		return true, false
	}
	number := int(angel*100) / a.angel
	if number == 0 {
		number = a.first / a.angel
	}
	key := keyRange{number * a.angel, number*a.angel - a.angel}
	if v, ok := a.rangeAngels[key]; ok {
		return dist < v.dist, v.dangerous
	}
	return true, false
}

func (a *Angels) ClosestAvailable(angel float64) (newAngel float64) {
	if len(a.rangeAngels) == 0 {
		return angel
	}
	number := int(angel*100) / a.angel
	if number == 0 {
		number = a.first / a.angel
	}
	angelL := number*a.angel - a.angel
	angelR := number*a.angel + a.angel
	correct := func() {
		if angelR > 628 {
			angelR = 0
		}
		if angelL <= 0 {
			angelL = a.first
		}
	}
	correct()
	for angelR != a.angel*number {
		key := keyRange{angelR + a.angel, angelR}
		if _, ok := a.rangeAngels[key]; !ok {
			return float64(angelR+a.angel) / 100
		}
		key = keyRange{angelL, angelL - a.angel}
		if _, ok := a.rangeAngels[key]; !ok {
			return float64(angelL-a.angel) / 100
		}
		angelL -= a.angel
		angelR += a.angel
		correct()
	}
	return angel
}

func CheckAngels(el animal.Animal, obstacles []Obstacle) (ang Angels) {
	if len(obstacles) == 0 {
		return ang
	}
	count := int((el.Vision() * math.Pi * 2) / el.Size())
	for count%4 != 0 {
		count++
	}
	addAngel := m2.Round(2.0 * math.Pi / float64(count))
	addInrAngel := int(2.0 * math.Pi / float64(count) * 100)
	addAngelD := addAngel * 2
	angel := 0.0
	sift := float64(int(count / 4.0))
	shiftCorrect := 0.0
	angelV := angel + addAngel*sift - shiftCorrect + addAngel
	xd := el.X() + el.Vision()*math.Cos(angelV)
	yd := el.Y() + el.Vision()*math.Sin(angelV)
	vec := vector.GetVectorByPoint(crd.NewCrd(el.X(), el.Y()), crd.NewCrd(xd, yd))
	dirAngel := int(geom.ModuleDegree(vec.GetAngle())*100) / addInrAngel
	dirAngel = dirAngel * addInrAngel
	rangAng := make(mapRange, len(obstacles))
	first := dirAngel + addInrAngel
	for i := 0; i < (count); i++ {
		xs1 := el.X() + el.Size()*math.Cos(angel)
		ys1 := el.Y() + el.Size()*math.Sin(angel)
		angel += math.Pi
		xs2 := el.X() + el.Size()*math.Cos(angel)
		ys2 := el.Y() + el.Size()*math.Sin(angel)
		angel -= math.Pi
		angelV = angel + addAngel*sift - shiftCorrect
		xf1 := el.X() + el.Vision()*math.Cos(angelV)
		yf1 := el.Y() + el.Vision()*math.Sin(angelV)
		angelV += addAngelD
		xf2 := el.X() + el.Vision()*math.Cos(angelV)
		yf2 := el.Y() + el.Vision()*math.Sin(angelV)
		line1 := geom.NewSegment(crd.NewCrd(xs1, ys1), crd.NewCrd(xf1, yf1))
		line2 := geom.NewSegment(crd.NewCrd(xf1, yf1), crd.NewCrd(xf2, yf2))
		line3 := geom.NewSegment(crd.NewCrd(xs2, ys2), crd.NewCrd(xf2, yf2))
		if dirAngel == 0 {
			dirAngel = first
		}
		for _, v := range obstacles {
			if intersect, dist := v.check(el.GetCrd(), el.Size(), line1, line2, line3); intersect {
				rangAng[keyRange{dirAngel, dirAngel - addInrAngel}] = rangeAngels{
					dangerous: v.dangerous(),
					dist:      dist,
				}
				if dirAngel == first || dirAngel - addInrAngel == 0 {
					rangAng[keyRange{first, first - addInrAngel}] = rangeAngels{
						dangerous: v.dangerous(),
						dist:      dist,
					}
				}
				break
			}
		}
		dirAngel -= addInrAngel
		angel += addAngel
	}
	ang.angel = addInrAngel
	ang.rangeAngels = rangAng
	ang.first = first
	return ang
}

type Obstacle interface {
	check(center crd.Crd, size float64, lines ...geom.Segment) (bool, float64)
	dangerous() bool
}

type point struct {
	outer []geom.Segment
	inner []geom.Segment
	size  float64
}

func NewPoint(el alive.Alive) Obstacle {
	p := point{
		size: el.Size(),
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

func (p *point) check(center crd.Crd, size float64, lines ...geom.Segment) (bool, float64) {
	res := func() (bool, float64) {
		dist := geom.GetDistanceByCrd(center, p.outer[0].Start())
		if dist > p.size+size {
			dist -= p.size + size
		} else {
			dist = math.Max(dist-(p.size*0.3+size), 0)
		}
		return true, dist
	}
	for _, v := range p.outer {
		for _, line := range lines {
			if v.Intersection(line) {
				return res()
			}
		}
	}
	countIntersect := 0
	for _, v := range p.inner {
		for _, line := range lines {
			if v.Intersection(line) {
				countIntersect++
				break
			}
		}
	}
	if countIntersect >= 3 {
		return res()
	}
	return false, 0
}

type line struct {
	line geom.Segment
}

func NewLine(l geom.Segment) Obstacle {
	return &line{
		line: l,
	}
}

func (l *line) dangerous() bool {
	return false
}

func (l *line) check(center crd.Crd, size float64, lines ...geom.Segment) (bool, float64) {
	line := geom.NewSegment(center, lines[1].MidPoint())
	if l.line.Intersection(line) {
		_, cr := l.line.IntersectionPoint(line)
		return true, geom.GetDistanceByCrd(center, cr) - size
	}
	return false, 0
}
