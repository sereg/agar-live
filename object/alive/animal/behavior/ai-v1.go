package behavior

import (
	"agar-life/object/alive/animal"
	"agar-life/object/alive/plant"
)

type aiV1 struct {

}

func NewAiv1() Behavior {
	return &aiV1{

	}
}

func (a *aiV1) SetDirection(animals []animal.Animal, plants []plant.Plant) {

}
