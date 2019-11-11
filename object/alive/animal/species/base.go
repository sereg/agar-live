package species

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	_const "agar-life/world/const"
	"agar-life/world/move"
	"encoding/json"
	"math"
)

type Base struct {
	alive.Base
	move.Move
	Speed     float64
	Vision    float64
	CycleGlue uint64
	Parent    animal.Animal `json:"-"`
	Children  []animal.Animal
	Behavior  animal.Behavior
}

type JsonBase struct {
	Base
	Parent bool
}

func (b *Base) MarshalJSON() ([]byte, error) {
	returnObj := JsonBase{}
	returnObj.Base = *b
	if b.Parent != nil {
		returnObj.Parent = true
	}
	return json.Marshal(returnObj)
}

func NewBase() animal.Animal {
	return &Base{
		Speed:  0,
		Vision: 0,
	}
}

func (b Base) Child(i int) animal.Animal {
	if len(b.Children) > i {
		return b.Children[i]
	}
	return nil
}

func (b *Base) AddChild(child animal.Animal) {
	b.Children = append(b.Children, child)
}

func (b *Base) SetCountChildren(count int) {
	b.Children = make([]animal.Animal, 0, count)
}

func (b *Base) SetBehaviour(behavior animal.Behavior) {
	b.Behavior = behavior
}

func (b Base) GetBehaviour() animal.Behavior {
	return b.Behavior
}

func (b Base) GetGlueTime() uint64 {
	return b.CycleGlue
}

func (b Base) Count() int {
	if len(b.GetChildren()) > 0 {
		return len(b.GetChildren()) + 1
	}
	if parent := b.GetParent(); parent != nil {
		return len(parent.GetChildren()) + 1
	}
	return 1
}

func (b *Base) SetGlueTime(cycle uint64) {
	b.CycleGlue = cycle + _const.GlueTime
}

func (b *Base) DeleteChild(id int) {
	if index := getIndexByID(b.Children, id); index != -1 {
		b.Children = animal.Remove(b.Children, index)
	}
}

func getIndexByID(a []animal.Animal, id int) int {
	for k, v := range a {
		if id == v.GetID() {
			return k
		}
	}
	return -1
}

func (b Base) GetChildren() []animal.Animal {
	return b.Children
}

func (b Base) GetParent() animal.Animal {
	return b.Parent
}

func (b *Base) SetParent(a animal.Animal) {
	b.Parent = a
}

func (b *Base) SetSize(size float64) {
	b.Base.SetSize(size)
	b.SetSpeed(reduce(size))
	b.SetVision(_const.StartVision + b.GetSize()*(_const.VisionRatio-math.Log(b.GetSize())))
}

func reduce(i float64) float64 {
	return (_const.StartSpeed - math.Log(i*_const.SpeedRatio)) / 10
}

func (b *Base) SetSpeed(speed float64) {
	if speed <= 0 {
		panic("Speed less than 0")
	}
	b.Speed = speed
}

func (b Base) GetSpeed() float64 {
	return b.Speed
}

func (b *Base) SetVision(vision float64) {
	b.Vision = vision
}

func (b Base) GetVision() float64 {
	return b.Vision
}

func (b *Base) Action(animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool) {
	return crd.Crd{}, false
}

func (b *Base) Direction() crd.Crd {
	return b.Move.GetDirection()
}

func (b *Base) Eat(el alive.Alive) {
	if el.GetDead() {
		return
	}
	eatRation := _const.EatIncreaseRation
	if b.GetGroup() == el.GetGroup() {
		eatRation = _const.EatSelfIncreaseRation
	}
	b.SetSize(b.GetSize() + (el.GetSize() * eatRation))
}
