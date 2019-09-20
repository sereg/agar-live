package generate

import (
	"agar-life/math"
	"strconv"
)

func Generate() {

}

func getRandomColor() string {
	r := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	g := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	b := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	return "#" + r + g + b
}
