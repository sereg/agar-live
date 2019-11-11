package alive

import "agar-life/object"

type Alive interface {
	object.Object
	Die()
	Revive()
	GetDead() bool
	Grow()
	Decrease()
	GetGroup() string
	SetGroup(string)
	GetID() int
	SetID(int)
	GetGlueTime() uint64
	SetGlueTime(uint64)
	GetDanger() bool
	SetDanger(bool)
	GetEdible() bool
	SetEdible(bool)
}

func Remove(a []Alive, i int) []Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}