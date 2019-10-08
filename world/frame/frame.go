package frame

import "agar-life/object/alive"

type Frame struct {
	updateState bool
	el          []alive.Alive
}

func NewFrame(count int) Frame {
	return Frame{
		el: make([]alive.Alive, count),
		updateState: true,
	}
}

func (f *Frame) Delete(index int){
	f.el = removeFromAlive(f.el, index)
}

func (f Frame) Get(index int) alive.Alive{
	return f.el[index]
}

func (f *Frame) Set(index int, el alive.Alive){
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
	f.el = append(f.el, el)
}

func removeFromAlive(a []alive.Alive, i int) []alive.Alive {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}