package world

import (
	math2 "agar-life/math"
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/species"
	_const "agar-life/world/const"
	"agar-life/world/frame"
	"math"

	"agar-life/object/alive/animal"
	gnt "agar-life/object/generate"
)



func Split(fr *frame.Frame, el animal.Animal, direction object.Crd, cycle uint64) {
	if el.Size() < _const.MinSizeSplit {
		return
	}
	println("split")
	size := math2.ToFixed(el.Size() * _const.Half, 2)
	el.SetSize(size)
	el.SetGlueTime(cycle)
	var parent animal.Animal
	if p := el.Parent(); p != nil {
		parent = p
	} else {
		parent = el
	}
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
	vec := getVectorWithLength(el.GetX(), el.GetY(), direction.GetX(), direction.GetY(), _const.SplitDist)
	direction = object.NewCrd(vec.GetPointFromVector(el.GetX(), el.GetY()))
	alv.SetInertia(direction)
	fr.Add(alv)
}

func getVectorWithLength(x1, y1, x2, y2, dist float64) vector.Vector {
	vec := vector.GetVectorByPoint(x1, y1, x2, y2)
	length := vec.Len()
	ratio := dist / length
	vec.MultiplyByScalar(ratio)
	return vec
}

func Burst(fr *frame.Frame, el animal.Animal, cycle uint64) bool {
	burstCount := _const.BurstCount
	if _const.SplitMaxCount < (el.Count()+int(burstCount)-1) {
		burstCount = _const.SplitMaxCount - el.Count()
		if burstCount < 2 {
			return true
		}
	}
	size := math2.ToFixed(el.Size() / float64(burstCount), 2)
	if size < _const.MinSizeAlive {
		burstCount = int(el.Size() / _const.MinSizeAlive)
		if burstCount < 2 {
			return false
		}
		size = math2.ToFixed(el.Size() / float64(burstCount), 2)
	}
	el.SetSize(size)
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
	return true
}


