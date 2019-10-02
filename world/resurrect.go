package world

import (
	"agar-life/object/alive/animal"
	gnt "agar-life/object/generate"
	"agar-life/world/const"
)

type resurrect struct {
	frame       *frame
	index       int
	cycleRevive uint64
}

type resurrects struct {
	r []resurrect
}

func (r *resurrects) add(frame *frame, index int, cycle uint64) {
	r.r = append(r.r, resurrect{frame: frame, index: index, cycleRevive: cycle + _const.ResurrectTime})
}

func (r *resurrects) resurrect(cycle uint64, w, h float64) {
	for i := 0; i < len(r.r); i++ {
		el := r.r[i]
		if el.cycleRevive <= cycle {
			alv := el.frame.el[el.index]
			if _, ok := alv.(animal.Animal); ok {
				gnt.Generate(alv, gnt.WorldWH(w, h), gnt.Size(6))
			} else {
				gnt.Generate(alv, gnt.WorldWH(w, h))
			}
			if el.index != el.frame.deedIndex {
				deedIndex := len(el.frame.el) - el.frame.deedIndex
				el.frame.el[el.index], el.frame.el[deedIndex] = el.frame.el[deedIndex], el.frame.el[el.index]
			}
			el.frame.deedIndex--
			el.frame.updateState = true
			r.r = removeFromResurrect(r.r, i)
			i--
		}
	}
}

func removeFromResurrect(a []resurrect, i int) []resurrect {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a = a[:len(a)-1]
	return a
}