package animal

import (
	"agar-life/object"
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
	GetDirection(animals []alive.Alive, plants []alive.Alive, cycle uint64) (object.Crd, bool)
	Direction() object.Crd
	SetCrdByDirection(a alive.Alive, direction object.Crd, dist float64, changeDirection bool)
	GetInertia() (direction object.Crd, speed float64)
	SetInertia(direction object.Crd)
	Count() int
}

func Remove(a []Animal, i int) []Animal {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}
