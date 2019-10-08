package species

import (
	"agar-life/object/alive/animal"
	"agar-life/object/alive/plant"
)

type poison struct {
	Base
}

func NewBeast(behavior animal.Behavior) plant.Plant {
	return &poison{
		Base: Base{
			danger: true,
			edible: false,
		},
	}
}
