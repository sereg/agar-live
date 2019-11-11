package object

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	_const "agar-life/world/const"
)

//Base is basic realization on object interface
type Base struct {
	Color    string
	Size     float64
	ViewSize float64
	GrowSize float64
	Hidden   bool
	crd.Crd
}

//GetColor return Color of point
func (p Base) GetColor() string {
	return p.Color
}

//SetColor sets a Color for the point
func (p *Base) SetColor(color string) {
	p.Color = color
}

//GetSize return Size of point
func (p Base) GetSize() float64 {
	return p.Size
}

//SetSize sets a Size for the point
func (p *Base) SetSize(size float64) {
	newSize := math2.Round(size)
	p.GrowSize = (newSize - p.ViewSize) / _const.GrowTime
	p.Size = newSize
}

//GetViewSize return ViewSize of point
func (p Base) GetViewSize() float64 {
	return p.ViewSize
}

//GetGrowSize return GrowSize of point
func (p Base) GetGrowSize() float64 {
	return p.GrowSize
}

//SetSize sets a Size for the point
func (p *Base) SetViewSize(size float64) {
	if size <= 0 {
		p.ViewSize = p.Size
	} else {
		p.ViewSize = size
	}
}

//GetHidden return is Hidden the point
func (p *Base) GetHidden() bool {
	return p.Hidden
}

//SetHidden set Hidden property
func (p *Base) SetHidden(isHidden bool) {
	p.Hidden = isHidden
}
