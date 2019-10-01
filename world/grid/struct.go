package grid

type xy struct {
	x, y int
}

type structGrid struct {
	cellSize     float64
	data         map[xy][]int
}

func NewStruct(size float64, cap int) Grid {
	return &structGrid{
		cellSize:     size,
		data:         make(map[xy][]int, cap),
	}
}

func (g structGrid) Len() int {
	return len(g.data)
}

func (g *structGrid) Reset() {

}

func (g *structGrid) Set(x, y float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	key := structGridKey(xInt, yInt)
	if _, ok := g.data[key]; ok {
		g.data[key] = append(g.data[key], i)
	} else {
		g.data[key] = []int{i}
	}
}

func structGridKey(xInt, yInt int) xy {
	return xy{x: xInt, y: yInt}
}

func (g structGrid) GetObjInRadius(x, y, vision float64, exclude int) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := structGridKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj
}