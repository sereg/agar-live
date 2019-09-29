package object

//Base is basic realization on object interface
type Base struct {
	color  string
	size   float64
	hidden bool
	Crd
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

//GetHidden return is hidden the point
func (p *Base) GetHidden() bool {
	return p.hidden
}

//Hidden set hidden property
func (p *Base) Hidden(isHidden bool) {
	p.hidden = isHidden
}
