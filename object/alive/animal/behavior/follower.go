package behavior

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type follower struct {
	direction       object.Crd
}

func NewFollower() Behavior {
	return &follower{}
}

func (a *follower) Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd {
	target := animals[0]
	a.direction.SetCrd(target.GetX(), target.GetY())
	return a.direction
}