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

func (g *structGrid) Set(x, y, size float64, i int) {
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

func (g structGrid) GetObjInRadius(x, y, radius float64, size float64, exclude int) ([]int, int) {
	ltx, lty := int((x-radius)/g.cellSize), int((y-radius)/g.cellSize)
	rdx, rdy := int((x+radius)/g.cellSize), int((y+radius)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := structGridKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj, 0
}