package species

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type beast struct {
	Base
}

func NewBeast(behavior animal.Behavior) animal.Animal {
	b := beast{}
	b.SetBehaviour(behavior)
	return &b
}

func (b *beast) GetDirection(animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	return b.Behaviour().Direction(animal.Animal(b), animals, plants, cycle)
}
