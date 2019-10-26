package world

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	_const "agar-life/world/const"
	"agar-life/world/frame"
	"math"

	"agar-life/object/alive/animal"
	gnt "agar-life/world/generate"
)

func Split(fr *frame.Frame, el animal.Animal, direction crd.Crd, cycle uint64) {
	if el.Size() < _const.MinSizeSplit {
		return
	}
	size := math2.Round(el.Size() * _const.Half)
	el.SetSize(size)
	el.SetGlueTime(cycle)
	var parent animal.Animal
	if p := el.Parent(); p != nil {
		parent = p
	} else {
		parent = el
	}
	direction = vector.GetCrdWithLength(el.GetCrd(), direction, _const.SplitDist)
	el.SetInertia(direction)
	alv := species.NewBeast(behavior.NewFollower())
	alv.SetParent(parent)
	alv.SetGlueTime(cycle)
	parent.AddChild(alv)
	gnt.Generate(
		alv,
		gnt.Size(size),
		gnt.Name(el.Group()),
		gnt.Color(el.Color()),
		gnt.Crd(gnt.FixCrd(el.X(), el.Y())),
	)

	fr.Add(alv)
}

func Burst(fr *frame.Frame, el animal.Animal, cycle uint64) bool {
	burstCount := _const.BurstCount
	if _const.SplitMaxCount < (el.Count() + burstCount - 1) {
		burstCount = _const.SplitMaxCount - el.Count()
		if burstCount < 2 {
			return true
		}
	}
	size := math2.Round(el.Size() / float64(burstCount))
	if size < _const.MinSizeAlive {
		burstCount = int(el.Size() / _const.MinSizeAlive)
		if burstCount < 2 {
			return false
		}
		size = math2.Round(el.Size() / float64(burstCount))
	}
	el.SetSize(size)
	addAngel := 2.0 * math.Pi / float64(burstCount)
	vec := vector.GetVectorByPoint(el.GetCrd(), crd.NewCrd(el.X()+_const.SplitDist, el.Y()))
	el.SetInertia(vec.GetPointFromVector(el.GetCrd()))
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
			gnt.Crd(gnt.FixCrd(el.X(), el.Y())),
		)
		vec.AddAngle(addAngel)
		alv.SetInertia(vec.GetPointFromVector(el.GetCrd()))
		fr.Add(alv)
	}
	return true
}
