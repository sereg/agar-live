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

func (g *stringGrid) Set(x, y float64, i int) {
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

func (g *stringGrid) GetObjInRadius(x, y, vision float64) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := stringKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj
}