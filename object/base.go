package object

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
)

//Base is basic realization on object interface
type Base struct {
	color  string
	size   float64
	hidden bool
	crd.Crd
}

//Color return color of point
func (p Base) Color() string {
	return p.color
}

//SetColor sets a color for the point
func (p *Base) SetColor(color string) {
	p.color = color
}

//Size return size of point
func (p Base) Size() float64 {
	return p.size
}

//SetSize sets a size for the point
func (p *Base) SetSize(size float64) {
	p.size = math2.Round(size)
}

//Hidden return is hidden the point
func (p *Base) Hidden() bool {
	return p.hidden
}

//SetHidden set hidden property
func (p *Base) SetHidden(isHidden bool) {
	p.hidden = isHidden
}
