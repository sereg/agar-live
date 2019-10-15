package animal

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
)

type Parent interface {
	Parent() Animal
	SetParent(Animal)
}

type Children interface {
	Child(int) Animal
	AddChild(Animal)
	SetCountChildren(int)
	DeleteChild(int)
	Children() []Animal
}

type Animal interface {
	alive.Alive
	Parent
	Children
	SetSpeed(float64)
	Speed() float64
	SetVision(float64)
	Vision() float64
	SetBehaviour(Behavior)
	Behaviour() Behavior
	Eat(a alive.Alive)
	Action(animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool)
	Direction() crd.Crd
	SetCrdByDirection(a alive.Alive, direction crd.Crd, dist float64, changeDirection bool)
	GetInertia() (direction crd.Crd, speed float64)
	SetInertia(direction crd.Crd)
	Count() int
}

func Remove(a []Animal, i int) []Animal {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}
