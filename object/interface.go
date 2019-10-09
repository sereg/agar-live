package object

import "math"

//Crd coordinates
type Crd struct {
	x, y float64
}

//NewCrd return new instance of crd
func NewCrd(x, y float64) Crd {
	return Crd{x: x, y: y}
}

//X set x
func (c *Crd) X(x float64) {
	c.x = x
}

//GetX it return x
func (c Crd) GetX() float64 {
	return c.x
}

//Y set x
func (c *Crd) Y(y float64) {
	c.y = y
}

//GetY it return y
func (c Crd) GetY() float64 {
	return c.y
}

//GetCrd it return crd
func (c Crd) GetCrd() Crd {
	return c
}

func (c *Crd) SetCrd(x, y float64) {
	if !math.IsNaN(x) && !math.IsNaN(y) {
		c.x, c.y = x, y
	}
}

//Object the base interface for all objects
type Object interface {
	Color() string
	SetColor(string)
	Size() float64
	SetSize(size float64)
	GetCrd() Crd
	SetCrd(x, y float64)
	GetX() float64
	GetY() float64
	Hidden() bool
	SetHidden(bool)
}
