package object

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

func (c *Crd) Set(x, y float64) {
	c.x, c.y = x, y
}

//Object the base interface for all objects
type Object interface {
	GetColor() string
	Color(string)
	GetSize() float64
	Size(size float64)
	GetCrd() Crd
	Crd(x, y float64)
	GetHidden() bool
	Hidden(bool)
}
