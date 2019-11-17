package animal

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
)

type Parent interface {
	GetParent() Animal
	SetParent(Animal)
}

type Children interface {
	Child(int) Animal
	AddChild(Animal)
	SetCountChildren(int)
	DeleteChild(int)
	GetChildren() []Animal
}

type Animal interface {
	alive.Alive
	Parent
	Children
	SetSpeed(float64)
	GetSpeed() float64
	SetVision(float64)
	GetVision() float64
	SetBehaviour(Behavior)
	GetBehaviour() Behavior
	Eat(a alive.Alive)
	Action(animals []alive.Alive, plants []alive.Alive, cycle uint) (crd.Crd, bool)
	Direction() crd.Crd
	SetCrdByDirection(a alive.Alive, direction crd.Crd, dist float64, changeDirection bool)
	GetInertia() (direction crd.Crd, speed float64)
	SetInertia(direction crd.Crd)
	SetInertiaImport(direction crd.Crd, speed, acceleration float64)
	Count() int
}

func Remove(a []Animal, i int) []Animal {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}
