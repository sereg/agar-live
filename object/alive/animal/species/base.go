package species

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/move"
	_const "agar-life/world/const"
	"math"
)

type Base struct {
	alive.Base
	move.Move
	speed     float64
	vision    float64
	cycleGlue uint64
	parent    animal.Animal
	children  []animal.Animal
	behavior  animal.Behavior
}

func NewBase() animal.Animal {
	return &Base{
		speed:  0,
		vision: 0,
	}
}

func (b Base) Child(i int) animal.Animal {
	if len(b.children) > i {
		return b.children[i]
	}
	return nil
}

func (b *Base) AddChild(child animal.Animal) {
	b.children = append(b.children, child)
}

func (b *Base) SetCountChildren(count int) {
	b.children = make([]animal.Animal, 0, count)
}

func (b *Base) SetBehaviour(behavior animal.Behavior){
	b.behavior = behavior
}

func (b Base) Behaviour() animal.Behavior {
	return b.behavior
}

func (b Base) GlueTime() uint64 {
	return b.cycleGlue
}

func (b Base) Count() int {
	if len(b.Children()) > 0 {
		return len(b.Children()) + 1
	}
	if parent := b.Parent(); parent != nil {
		return len(parent.Children()) + 1
	}
	return 1
}

func (b *Base) SetGlueTime(cycle uint64) {
	b.cycleGlue = cycle + _const.GlueTime
}

func (b *Base) DeleteChild(id int){
	if index := getIndexByID(b.children, id); index != -1 {
		b.children = animal.Remove(b.children, index)
	}
}

func getIndexByID(a []animal.Animal, id int) int{
	for k, v := range a {
		if id == v.ID() {
			return k
		}
	}
	return -1
}

func (b Base) Children() []animal.Animal {
	return b.children
}

func (b Base) Parent() animal.Animal {
	return b.parent
}

func (b *Base) SetParent(a animal.Animal) {
	b.parent = a
}

func (b *Base) SetSize(size float64) {
	b.Base.SetSize(size)
	b.SetSpeed(reduce(size))
	b.SetVision(_const.StartVision + b.Size()*(_const.VisionRatio-math.Log(b.Size())))
}

func reduce(i float64) float64 {
	return (_const.StartSpeed - math.Log(i*_const.SpeedRatio)) / 10
}

func (b *Base) SetSpeed(speed float64) {
	if speed <= 0 {
		panic("speed less than 0")
	}
	b.speed = speed
}

func (b Base) Speed() float64 {
	return b.speed
}

func (b *Base) SetVision(vision float64) {
	b.vision = vision
}

func (b Base) Vision() float64 {
	return b.vision
}

func (b *Base) GetDirection(animals []alive.Alive, plants []alive.Alive, cycle uint64) (object.Crd, bool) {
	return object.Crd{}, false
}

func (b *Base) Direction() object.Crd {
	return b.Move.GetDirection()
}

func (b *Base) Eat(el alive.Alive) {
	if el.GetDead() {
		return
	}
	eatRation := _const.EatIncreaseRation
	if b.Group() == el.Group() {
		eatRation = _const.EatSelfIncreaseRation
	}
	b.SetSize(b.Size() + (el.Size() * eatRation))
}
