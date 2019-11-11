package behavior

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type follower struct {
	direction       crd.Crd
	Name string
}

func NewFollower() animal.Behavior {
	return &follower{Name: "follower"}
}

func (a *follower) Action(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	if parent := self.GetParent(); parent != nil {
		a.direction.SetCrd(parent.GetCrd())
	}
	return a.direction, false
}