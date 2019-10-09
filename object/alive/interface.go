package alive

import "agar-life/object"

type Alive interface {
	object.Object
	Die()
	Revive()
	GetDead() bool
	Grow()
	Decrease()
	Group() string
	SetGroup(string)
	ID() int
	SetID(int)
	GlueTime() uint64
	SetGlueTime(uint64)
	Danger() bool
	SetDanger(bool)
	Edible() bool
	SetEdible(bool)
}

func Remove(a []Alive, i int) []Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}