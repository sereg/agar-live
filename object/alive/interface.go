package alive

import "agar-life/object"

type Alive interface {
	object.Object
	Die()
	Revive()
	GetDead() bool
	Grow()
	Decrease()
	GetName() string
	Name(string)
}
