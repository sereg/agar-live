package frame

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
)

type Frame struct {
	w, h        float64
	updateState bool
	secID       int
	el          []alive.Alive
}

func NewFrame(count int, w, h float64) Frame {
	return Frame{
		el:          make([]alive.Alive, count),
		w: w, h: h,
		updateState: true,
	}
}

func (f *Frame) Delete(index int) {
	if el, ok := f.el[index].(animal.Animal); ok {
		if parent := el.Parent(); parent != nil {
			parent.DeleteChild(el.ID())
		}
		if children := el.Children(); len(children) > 0 {
			parent := el.Child(0)
			parent.SetParent(nil)
			parent.SetBehaviour(behavior.NewAiv1(f.w, f.h))
			parent.SetCountChildren(len(children))
			for i := 1; i < len(children); i++ {
				el.Child(i).SetParent(parent)
				parent.AddChild(el.Child(i))
			}
		}
	}
	f.el = alive.Remove(f.el, index)
}

func (f Frame) Get(index int) alive.Alive {
	return f.el[index]
}

func (f *Frame) Set(index int, el alive.Alive) {
	el.SetID(f.sec())
	f.el[index] = el
}

func (f Frame) All() []alive.Alive {
	return f.el
}

func (f *Frame) SetUpdateState(state bool) {
	f.updateState = state
}

func (f Frame) UpdateState() bool {
	return f.updateState
}

func (f *Frame) Add(el alive.Alive) {
	el.SetID(f.sec())
	f.el = append(f.el, el)
}

func (f *Frame) sec() (id int) {
	id = f.secID
	f.secID++
	return
}
