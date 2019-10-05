package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/plant"
)

type Base struct {
	alive.Base
	danger bool
	edible bool
}

func NewBase() plant.Plant{
	return &Base{}
}

func (p Base)GetDanger() bool{
	return p.danger
}
func (p Base) GetEdible() bool{
	return p.edible
}
