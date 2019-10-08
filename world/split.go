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
	el := fr.Get(cur)
	size := math2.ToFixed(el.GetSize() / _const.Half, 2)
	el.Size(size)
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

func Burst(fr *frame.Frame, cur int) {
	el := fr.Get(cur).(animal.Animal)
	size := math2.ToFixed(el.GetSize() / _const.BurstCount, 2)
	el.Size(size)
	addAngel := 2.0 * math.Pi / _const.BurstCount
	vec := vector.GetVectorByPoint(el.GetX(), el.GetY(), el.GetX()+_const.SplitDist, el.GetY())
	el.SetInertia(object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY())))
	for i := 0; i < _const.BurstCount; i++ {
		alv := species.NewBeast(behavior.NewFollower())
		gnt.Generate(alv, gnt.Size(size), gnt.Crd(gnt.FixCrd(el.GetX(), el.GetY())))
		vec.AddAngle(addAngel)
		el.SetInertia(object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY())))
		fr.Add(alv)
	}
}


