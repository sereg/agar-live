package behavior

import (
	"agar-life/object/alive/animal"
	"agar-life/object/alive/plant"
)

type simple struct {
	self animal.Animal
}

func NewSimple() Behavior {
	return &simple{

	}
}

func (s *simple) SetDirection(animals []animal.Animal, plants []plant.Plant) {

}
