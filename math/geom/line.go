package geom

import (
	"agar-life/math/crd"
	"errors"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x:x, y:y}
}

func (l Point) X() float64 {
	return l.x
}

func (l Point) Y() float64 {
	return l.y
}

type Line struct {
	slope float64
	y     float64
}

func NewLine(a, b Point) Line {
	var slope float64
	if (b.x - a.x) == 0 {
		b.x += 0.01
	}
	slope = (b.y - a.y) / (b.x - a.x)
	y := a.y - slope*a.x
	return Line{slope, y}
}

func NewLineCrd(a, b crd.Crd) Line {
	var slope float64
	ax, ay, bx, by := a.X(), a.Y(), b.X(), b.Y()
	if (bx - ax) == 0 {
		bx += 0.01
	}
	slope = (by - ay) / (bx - ax)
	y := ay - slope*ax
	return Line{slope, y}
}

func (l Line) evalX(x float64) float64 {
	return l.slope*x + l.y
}

func (l Line) Intersection(l2 Line) (Point, error) {
	if l.slope == l2.slope {
		return Point{}, errors.New("the lines do not intersect")
	}
	x := (l2.y - l.y) / (l.slope - l2.slope)
	y := l.evalX(x)
	return Point{x, y}, nil
}

func LengthLine(a, b Point) float64 {
	d1 := a.x - b.x
	d2 := a.y - b.y
	return math.Sqrt(
		d1*d1 + d2*d2,
	)
}

func AngleAxisX(a, b Point) float64 {
	return math.Atan2(b.y- a.y, b.x- a.x)
}
