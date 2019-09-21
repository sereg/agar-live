package behavior

import (
	"agar-life/object/alive"
)

type aiV1 struct {

}

func NewAiv1() Behavior {
	return &aiV1{

	}
}

func (a *aiV1) SetDirection(animals []alive.Alive, plants []alive.Alive) {

}
