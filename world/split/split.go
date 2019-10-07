package split

import (
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"

	//"agar-life/object/alive/animal"
	gnt "agar-life/object/generate"
	"agar-life/object/alive/animal"
	"agar-life/world"
)

type split struct {
	Frame     *world.Frame
	index     int
	cycleGlue uint64
}

type splits struct {
	r []split
}

func (r *splits) split(fr *world.Frame, index int, cycle uint64) {
	el := fr.El()[index][0]
	if _, ok := el.(animal.Animal); ok {
		alv := species.NewBeast(behavior.NewAiv1(w, h))
		gnt.Generate(alv, gnt.Size(6), gnt.Crd(gnt.FixCrd(x, y))
	} else {
		gnt.Generate(alv, gnt.WorldWH(w, h))
	}
	for _, v := range fr.El()[index] {

	}
	//r.r = append(r.r, split{Frame: world.Frame, el: el, cycleRevive: cycle + _const.splitTime})
}

func (r *splits) glue(cycle uint64) {
	for i := 0; i < len(r.r); i++ {
		el := r.r[i]
		if el.cycleGlue <= cycle {
			//alv := el.el
			//if _, ok := alv.(animal.Animal); ok {
			//	gnt.Generate(alv, gnt.WorldWH(w, h), gnt.Size(6))
			//} else {
			//	gnt.Generate(alv, gnt.WorldWH(w, h))
			//}
			//el.Frame.el = append(el.Frame.el, []alive.Alive{alv})
			//el.Frame.updateState = true
			r.r = removeFromSplit(r.r, i)
			i--
		}
	}
}

func removeFromSplit(a []split, i int) []split {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a = a[:len(a)-1]
	return a
}
