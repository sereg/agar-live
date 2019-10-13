package behavior

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type follower struct {
	direction       object.Crd
}

func NewFollower() animal.Behavior {
	return &follower{}
}


func (a *follower) GetDirection() object.Crd {
	return a.direction
}

func (a *follower) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	if parent := self.Parent(); parent != nil {
		a.direction.SetCrd(parent.GetX(), parent.GetY())
	}
	return a.direction
}