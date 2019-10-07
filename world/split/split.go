package split

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	_const "agar-life/world/const"

	"agar-life/object/alive/animal"
	sp "agar-life/object/alive/plant/species"
	//"agar-life/object/alive/animal"
	gnt "agar-life/object/generate"
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
	size := el.GetSize() / _const.Half
	el.Size(size)
	var alv alive.Alive
	if _, ok := el.(animal.Animal); ok {
		alv = species.NewBeast(behavior.NewFollower())
		gnt.Generate(alv, gnt.Size(size), gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())))
	} else {
		alv = sp.NewPlant()
		gnt.Generate(alv, gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())))
	}
	fr.El()[index] = append(fr.El()[index], alv)
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
