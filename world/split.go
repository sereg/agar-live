package world

import (
	math2 "agar-life/math"
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	_const "agar-life/world/const"
	"agar-life/world/frame"
	"math"

	"agar-life/object/alive/animal"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
)



func Split(fr *frame.Frame, cur int, direction object.Crd) {
	//TODO dont split if size less than _const.MinSizeSplit
	el := fr.Get(cur)
	size := math2.ToFixed(el.Size() / _const.Half, 2)
	el.SetSize(size)
	var alv alive.Alive
	if _, ok := el.(animal.Animal); ok {
		alv = species.NewBeast(behavior.NewFollower())
		gnt.Generate(alv, gnt.Size(size), gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())))
	} else {
		alv = sp.NewPlant()
		gnt.Generate(alv, gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())))
	}
	fr.Add(alv)
}

func Burst(fr *frame.Frame, el animal.Animal, cycle uint64) {
	burstCount := _const.BurstCount
	if _const.SplitMaxCount < (el.Count()+int(burstCount)-1) {
		burstCount = _const.SplitMaxCount - el.Count()
		if burstCount < 2 {
			return
		}
	}
	size := math2.ToFixed(el.Size() / float64(burstCount), 2)
	if size < _const.MinSizeAlive {
		burstCount = int(el.Size() / _const.MinSizeAlive)
		if burstCount < 2 {
			return
		}
		size = math2.ToFixed(el.Size() / float64(burstCount), 2)
	}
	el.SetSize(size) //TODO choose size and count of element based on _const.MinSizeAlive
	addAngel := 2.0 * math.Pi / float64(burstCount)
	vec := vector.GetVectorByPoint(el.GetX(), el.GetY(), el.GetX()+_const.SplitDist, el.GetY())
	el.SetInertia(object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY())))
	el.SetGlueTime(cycle)
	var parent animal.Animal
	if p := el.Parent(); p != nil {
		parent = p
	} else {
		parent = el
	}
	for i := 1; i < burstCount; i++ {
		alv := species.NewBeast(behavior.NewFollower())
		alv.SetParent(parent)
		alv.SetGlueTime(cycle)
		parent.AddChild(alv)
		gnt.Generate(
			alv,
			gnt.Size(size),
			gnt.Name(el.Group()),
			gnt.Color(el.Color()),
			gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())),
		)
		vec.AddAngle(addAngel)
		alv.SetInertia(object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY())))
		fr.Add(alv)
	}
}


