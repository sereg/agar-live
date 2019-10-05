package species

import (
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/plant"
)

type poison struct {
	Base
}

func NewBeast(behavior behavior.Behavior) plant.Plant {
	return &poison{
		Base: Base{
			danger: true,
			edible: false,
		},
	}
}
