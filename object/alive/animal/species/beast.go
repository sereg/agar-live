package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
)

type beast struct {
	animal.Base
	behavior behavior.Behavior
}

func NewBeast(behavior behavior.Behavior) animal.Animal {
	return &beast{
		behavior: behavior,
	}
}

func (b *beast) Step(animals []alive.Alive, plants []alive.Alive) {
	b.behavior.SetDirection(animal.Animal(b), animals, plants)
}
