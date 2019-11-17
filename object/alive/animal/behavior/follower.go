package behavior

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

const (
	FollowerName = "follower"
)

type follower struct {
	direction       crd.Crd
	Name string
}

func NewFollower() animal.Behavior {
	return &follower{Name: FollowerName}
}

func (a *follower) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint) (crd.Crd, bool) {
	if parent := self.GetParent(); parent != nil {
		a.direction.SetCrd(parent.GetCrd())
	}
	return a.direction, false
}