package vector

import "math"

type Vector struct {
	x, y float64
}

func NewVector(x, y float64) Vector {
	return Vector{x: x, y: y}
}

func (vec Vector) Len() float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y)
}

func (vec *Vector) MultiplyByScalar(s float64) {
	vec.x *= s
	vec.y *= s
}

func GetVectorByPoint(x, y, x2, y2 float64) Vector {
	vec := Vector{}
	vec.x = x2 - x
	vec.y = y2 - y
	return vec
}

func (vec Vector) GetPointFromVector(x, y float64) (float64, float64) {
	xr := vec.x + x
	yr := vec.y + y
	return xr, yr
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
