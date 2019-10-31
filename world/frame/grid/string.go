package grid

import "strconv"

type stringGrid struct {
	cellSize     float64
	data   map[string][]int
}

func NewString(size float64, cap int) Grid {
	return &stringGrid{
		cellSize:     size,
		data: make(map[string][]int, cap),
	}
}

func (g *stringGrid) Len() int {
	return len(g.data)
}

func (g *stringGrid) Reset() {

}

func (g *stringGrid) Set(x, y, size float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	key := stringKey(xInt, yInt)
	if _, ok := g.data[key]; ok {
		g.data[key] = append(g.data[key], i)
	} else {
		g.data[key] = []int{i}
	}
}

func stringKey(xInt, yInt int) string {
	return strconv.Itoa(xInt) +"x:"+ strconv.Itoa(yInt)+"y"
}

func (g *stringGrid) GetObjInRadius(x, y, radius float64, size float64, exclude int) ([]int, int) {
	ltx, lty := int((x-radius)/g.cellSize), int((y-radius)/g.cellSize)
	rdx, rdy := int((x+radius)/g.cellSize), int((y+radius)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := stringKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj, 0
}