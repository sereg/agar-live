package object

import "agar-life/math/crd"

//Object the base interface for all objects
type Object interface {
	Color() string
	SetColor(string)
	Size() float64
	SetSize(size float64)
	GetCrd() crd.Crd
	SetCrd(crd.Crd)
	SetX(float64)
	SetY(float64)
	SetXY(float64, float64)
	X() float64
	Y() float64
	Hidden() bool
	SetHidden(bool)
}
