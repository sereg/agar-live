package vector

import (
	"agar-life/math/crd"
	"math"
)

type Vector struct {
	x, y float64
}

func NewVector(crd crd.Crd) Vector {
	return Vector{x: crd.X(), y: crd.Y()}
}

func GetVectorByPoint(a, b crd.Crd) Vector {
	vec := Vector{}
	vec.x = b.X() - a.X()
	vec.y = b.Y() - a.Y()
	return vec
}

func GetVectorWithLength(a, b crd.Crd, dist float64) Vector {
	vec := GetVectorByPoint(a, b)
	length := vec.Len()
	ratio := dist / length
	vec.MultiplyByScalar(ratio)
	return vec
}

func GetCrdWithLength(a, b crd.Crd, dist float64) crd.Crd {
	vec := GetVectorWithLength(a, b, dist)
	return vec.GetPointFromVector(a)
}

func (vec Vector) Len() float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y)
}

func (vec *Vector) MultiplyByScalar(s float64) {
	vec.x *= s
	vec.y *= s
}

func (vec Vector) MultiplyByVector(s Vector) Vector {
	//vec.x *= s
	//vec.y *= s
	return vec
}

func (vec *Vector) AddAngle(angle float64) {
	l := vec.Len()
	y := math.Cos(vec.GetAngle() + angle) * l
	x := math.Sin(vec.GetAngle() + angle) * l
	if x == l {
		y = 0
	}
	if y == l {
		x = 0
	}
	vec.x = x
	vec.y = y
}

func (vec *Vector) SetAngle(angle float64) {
	l := vec.Len()
	y := math.Cos(angle) * l
	x := math.Sin(angle) * l
	if x == l {
		y = 0
	}
	if y == l {
		x = 0
	}
	vec.x = x
	vec.y = y
}

func (vec Vector) GetPointFromVector(a crd.Crd) crd.Crd {
	xr := vec.x + a.X()
	yr := vec.y + a.Y()
	return crd.NewCrd(xr, yr)
}

func (vec Vector) GetAngle() float64 {
	angel := math.Atan2(vec.x, vec.y)
	if math.IsNaN(angel) {
		angel = math.Pi
	}
	return angel
}

func (vec Vector) GetPerpendicularVector(x float64) Vector {
	y := (-1 * vec.x * x) / vec.y
	return Vector{x: x, y: y}
}

func Add(a, b Vector) Vector {
	a.x += b.x
	a.y += b.y
	return a
}

func Compare(a, b Vector) bool {
	return a.x == b.x && a.y == b.y
}
