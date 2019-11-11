package crd

import (
	math2 "agar-life/math"
	"fmt"
	"math"
)

//Crd coordinates
type Crd struct {
	X, Y float64
}

//NewCrd return new instance of crd
func NewCrd(x, y float64) Crd {
	return Crd{X: math2.Round(x), Y: math2.Round(y)}
}

//SetX set X
func (c *Crd) SetX(x float64) {
	c.X = math2.Round(x)
}

//GetX it return X
func (c Crd) GetX() float64 {
	return c.X
}

//GetY set X
func (c *Crd) SetY(y float64) {
	c.Y = math2.Round(y)
}

//SetXY set X and Y
func (c *Crd) SetXY(x float64, y float64) {
	if (x == 0 && y == 0) || (x == 1 && y == 1) {
		fmt.Println("ff")
	}
	c.X, c.Y = math2.Round(x), math2.Round(y)
}

//GetY it return Y
func (c Crd) GetY() float64 {
	return c.Y
}

//GetCrd it return crd
func (c Crd) GetCrd() Crd {
	return c
}

func (c *Crd) SetCrd(a Crd) {
	if !math.IsNaN(a.GetX()) && !math.IsNaN(a.GetY()) {
		c.X, c.Y = math2.Round(a.GetX()), math2.Round(a.GetY())
	}
}
