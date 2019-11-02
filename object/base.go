package object

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	_const "agar-life/world/const"
)

//Base is basic realization on object interface
type Base struct {
	color    string
	size     float64
	viewSize float64
	growSize float64
	hidden   bool
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
	newSize := math2.Round(size)
	p.growSize = (newSize - p.size) / _const.GrowTime
	if (p.growSize > 0 && p.viewSize > p.size) || (p.growSize < 0 && p.viewSize < p.size) {
		p.growSize = (newSize - p.viewSize) / _const.GrowTime
	}
	p.size = newSize
}

//ViewSize return viewSize of point
func (p Base) ViewSize() float64 {
	return p.viewSize
}

//GrowSize return growSize of point
func (p Base) GrowSize() float64 {
	return p.growSize
}

//SetSize sets a size for the point
func (p *Base) SetViewSize(size float64) {
	if size <= 0 {
		p.viewSize = p.size
	} else {
		p.viewSize = size
	}
}

//Hidden return is hidden the point
func (p *Base) Hidden() bool {
	return p.hidden
}

//SetHidden set hidden property
func (p *Base) SetHidden(isHidden bool) {
	p.hidden = isHidden
}
