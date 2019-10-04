package species

import (
	"agar-life/object/alive/plant"
)

type plantX struct {
	Base
}

func NewPlant() plant.Plant{
	return &plantX{
		Base: Base{
			danger: false,
			edible: true,
		},
	}
}
