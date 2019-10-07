package species

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/move"
)

type beast struct {
	animal.Base
	behavior behavior.Behavior
	move.Move
}

func NewBeast(behavior behavior.Behavior) animal.Animal {
	return &beast{
		behavior: behavior,
	}
}

func (b *beast) GetDirection(animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	return b.behavior.Direction(animal.Animal(b), animals, plants, cycle)
}
