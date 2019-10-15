package crd

import (
	math2 "agar-life/math"
	"math"
)

//Crd coordinates
type Crd struct {
	x, y float64
}

//NewCrd return new instance of crd
func NewCrd(x, y float64) Crd {
	return Crd{x: math2.ToFixed(x, 2), y: math2.ToFixed(y, 2)}
}

//SetX set x
func (c *Crd) SetX(x float64) {
	c.x = math2.ToFixed(x, 2)
}

//GetX it return x
func (c Crd) X() float64 {
	return c.x
}

//Y set x
func (c *Crd) SetY(y float64) {
	c.y = math2.ToFixed(y, 2)
}

//SetXY set x and y
func (c *Crd) SetXY(x float64, y float64) {
	c.x, c.y = math2.ToFixed(x, 2), math2.ToFixed(y, 2)
}

//GetY it return y
func (c Crd) Y() float64 {
	return c.y
}

//GetCrd it return crd
func (c Crd) GetCrd() Crd {
	return c
}

func (c *Crd) SetCrd(a Crd) {
	if !math.IsNaN(a.X()) && !math.IsNaN(a.Y()) {
		c.x, c.y = math2.ToFixed(a.X(), 2), math2.ToFixed(a.Y(), 2)
	}
}
