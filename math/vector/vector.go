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

//func getXYWithLength(x1, y1, x2, y2, dist float64) (x float64, y float64) {
//	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
//	length := vec.Len()
//	ratio := dist
//	if length > 0 {
//		ratio = dist / length
//	}
//	vec.MultiplyByScalar(ratio)
//	x, y = vec.GetPointFromVector(x2, y2)
//	x, y = x-x2, y-y2
//	return
//}

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
	return math.Atan2(vec.x, vec.y)
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
