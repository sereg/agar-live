package species

import (
	"agar-life/object/alive/plant"
)

type plantX struct {
	Base
}

func NewPlant() plant.Plant{
	p := plantX{}
	p.SetDanger(true)
	p.SetEdible(false)
	return &p
}
