package behavior

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type aiV1 struct {

}

func NewAiv1() Behavior {
	return &aiV1{

	}
}

func (a *aiV1) SetDirection(self animal.Animal, animals []alive.Alive, plants []alive.Alive) {

}
