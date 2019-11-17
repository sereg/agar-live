package object

import "agar-life/math/crd"

//Object the base interface for all objects
type Object interface {
	GetColor() string
	SetColor(string)
	GetSize() float64
	SetSize(size float64)
	GetViewSize() float64
	SetViewSize(size float64)
	GetGrowSize() float64
	GetCrd() crd.Crd
	SetCrd(crd.Crd)
	SetX(float64)
	SetY(float64)
	SetXY(float64, float64)
	GetX() float64
	GetY() float64
	GetHidden() bool
	SetHidden(bool)
}
