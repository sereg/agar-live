package behavior

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type simple struct {
	self animal.Animal
}

func NewSimple() Behavior {
	return &simple{

	}
}

func (s *simple) SetDirection(animals []alive.Alive, plants []alive.Alive) {

}
