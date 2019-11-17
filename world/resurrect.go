package world

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/world/const"
	"agar-life/world/frame"
	gnt "agar-life/world/generate"
)

type resurrect struct {
	frame       *frame.Frame
	el          alive.Alive
	cycleRevive uint
}

type resurrects struct {
	r []resurrect
}

func (r *resurrects) add(frame *frame.Frame, el alive.Alive, cycle uint) {
	r.r = append(r.r, resurrect{frame: frame, el: el, cycleRevive: cycle + _const.ResurrectTime})
}

func (r *resurrects) resurrect(cycle uint, w, h float64) {
	for i := 0; i < len(r.r); i++ {
		el := r.r[i]
		if el.cycleRevive <= cycle {
			alv := el.el
			if _, ok := alv.(animal.Animal); ok {
				gnt.Generate(alv, gnt.WorldWH(w, h), gnt.Size(6))
			} else {
				gnt.Generate(alv, gnt.WorldWH(w, h))
			}
			el.frame.Add(alv)
			el.frame.SetUpdateState(true)
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
