package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/plant"
)

type plantX struct {
	alive.Base
	danger bool
	edible bool
}

func NewPlant() plant.Plant{
	return &plantX{}
}

func (p plantX)GetDanger() bool{
	return p.danger
}
func (p plantX) GetEdible() bool{
	return p.edible
}
