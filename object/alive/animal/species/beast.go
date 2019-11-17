package species

import (
	"agar-life/math/crd"
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

func (b *beast) Action(animals []alive.Alive, plants []alive.Alive, cycle uint) (crd.Crd, bool) {
	return b.GetBehaviour().Action(animal.Animal(b), animals, plants, cycle)
}
