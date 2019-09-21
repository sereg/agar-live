package generate

import (
	"agar-life/math"
	"agar-life/object/alive"
	"strconv"
)

func Generate(el alive.Alive, w, h float64, name string)	 {
	el.Color(getRandomColor())
	el.Size(3)
	el.Crd(
		float64(math.Random(int(0), int(w))),
		float64(math.Random(int(0), int(h))),
	)
	el.Hidden(false)
	el.Name(name)
}

func getRandomColor() string {
	r := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	g := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	b := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	return "#" + r + g + b
}
