package object

//Base is basic realization on object interface
type Base struct {
	color  string
	size   float64
	crd    Crd
	hidden bool
}

//GetColor return color of point
func (p Base) GetColor() string {
	return p.color
}

//Color sets a color for the point
func (p *Base) Color(color string) {
	p.color = color
}

//GetSize return size of point
func (p Base) GetSize() float64 {
	return p.size
}

//Size sets a size for the point
func (p *Base) Size(size float64) {
	p.size = size
}

//GetCrd return coordinates
func (p *Base) GetCrd() Crd {
	return p.crd
}

//Crd set coordinates
func (p *Base) Crd(x, y float64) {
	p.crd.x, p.crd.x = x, y
}

//GetHidden return is hidden the point
func (p *Base) GetHidden() bool {
	return p.hidden
}

//Hidden set hidden property
func (p *Base) Hidden(isHidden bool) {
	p.hidden = isHidden
}
